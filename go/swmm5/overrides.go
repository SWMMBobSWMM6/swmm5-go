// overrides.go - hand-ported pieces of EPA SWMM 5.2.4 that use C pointer arithmetic the translator does
// not map: odesolve.c (whole file), mempool.c (its pool becomes plain byte buffers) and six helpers from
// swmm5.c and project.c. Each keeps the C statement order and floating-point expressions exactly (every
// product is wrapped in float64(...) so no Go compiler may fuse it into an FMA); pow is glibc's.
// Ported from the Java port's Overrides.java. Written by Claude (Anthropic); see AI-USE.md.
package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

// OdeDerivs is odesolve's derivative callback: derivs(x, y, dydx).
type OdeDerivs func(float64, []float64, []float64)

const MAXFNAME_ = 259

func sstrncpy(dest R.CPtr, src R.CPtr, n uint64) uint64 {
	offset := int32(0)
	if n > 0 {
		for src.G(offset) != 0 {
			if uint64(offset) == n {
				break
			}
			dest.S(offset, src.G(offset))
			offset++
		}
	}
	dest.S(offset, 0)
	return R.Strlen(dest)
}

func sstrcat(dest R.CPtr, src R.CPtr, size uint64) uint64 {
	offset := int32(R.Strlen(dest))
	srcIndex := int32(0)
	for src.G(srcIndex) != 0 {
		dest.S(offset, src.G(srcIndex))
		offset++
		srcIndex++
		if uint64(offset) == size-1 {
			break
		}
	}
	dest.S(offset, 0)
	return R.Strlen(dest)
}

func ovIsRelative(fname R.CPtr) bool {
	if !R.Strchr(fname, ':').IsNull() {
		return false
	}
	if fname.G(0) == '\\' {
		return false
	}
	if fname.G(0) == '/' {
		return false
	}
	return true
}

func getAbsolutePath(fname R.CPtr, absPath R.CPtr, size uint64) {
	if fname.IsNull() || R.Strlen(fname) == 0 {
		return
	}
	if ovIsRelative(fname) {
		full := R.RealpathOf(R.Str(fname)) // realpath(fname, absPath)
		m := len(full)
		if room := len(absPath.B) - int(absPath.O) - 1; m > room {
			m = room
		}
		for i := 0; i < m; i++ {
			absPath.S(int32(i), int32(full[i]))
		}
		absPath.S(int32(m), 0)
	} else {
		sstrncpy(absPath, fname, R.Strlen(fname))
	}
	endOfDir := R.Strrchr(absPath, '/')
	if !endOfDir.IsNull() {
		endOfDir.S(1, 0)
	}
}

func getTempFileName(fname R.CPtr) R.CPtr {
	sstrncpy(fname, R.Lit("swmmXXXXXX"), MAXFNAME_)
	R.Mkstemp(fname)
	return fname
}

func project_createMatrix(nrows int32, ncols int32) [][]float64 {
	if nrows <= 0 || ncols <= 0 {
		return nil
	}
	a := make([][]float64, nrows)
	for i := range a {
		a[i] = make([]float64, ncols)
	}
	return a
}

func project_freeMatrix(a [][]float64) {}

// mempool.c: object IDs live in ordinary byte buffers
var ovPool alloc_handle_t

func AllocInit() *alloc_handle_t                           { return &ovPool }
func Alloc(size int64) R.CPtr                              { return R.Buf(size) }
func AllocSetPool(newpool *alloc_handle_t) *alloc_handle_t { return newpool }
func AllocReset()                                          {}
func AllocFreePool()                                       {}

// ------------------------------------------------------------ odesolve.c (Cash-Karp Runge-Kutta)

const (
	ovMAXSTP = 10000
	ovTINY   = 1.0e-30
	ovSAFETY = 0.9
	ovPGROW  = -0.2
	ovPSHRNK = -0.25
	ovERRCON = 1.89e-4
)

var (
	ovNmax                                                                   int32
	ovY, ovYscal, ovDydx, ovYerr, ovYtemp, ovAk2, ovAk3, ovAk4, ovAk5, ovAk6 []float64
)

func odesolve_open(n int32) int32 {
	ovNmax = 0
	mk := func() []float64 { return make([]float64, n) }
	ovY, ovYscal, ovDydx, ovYerr, ovYtemp = mk(), mk(), mk(), mk(), mk()
	ovAk2, ovAk3, ovAk4, ovAk5, ovAk6 = mk(), mk(), mk(), mk(), mk()
	ovNmax = n
	return 1
}

func odesolve_close() { ovNmax = 0 }

func odesolve_integrate(ystart []float64, n int32, x1 float64, x2 float64, eps float64, h1 float64, derivs OdeDerivs) int32 {
	var hdid, hnext float64
	x := x1
	h := h1
	if ovNmax < n {
		return 1
	}
	for i := int32(0); i < n; i++ {
		ovY[i] = ystart[i]
	}
	for nstp := 1; nstp <= ovMAXSTP; nstp++ {
		derivs(x, ovY, ovDydx)
		for i := int32(0); i < n; i++ {
			ovYscal[i] = float64(float64(M.Fabs(ovY[i])+M.Fabs(float64(ovDydx[i]*h))) + float64(ovTINY))
		}
		if float64(float64(float64(x+h)-x2)*float64(float64(x+h)-x1)) > 0.0 {
			h = float64(x2 - x)
		}
		errcode := ovRkqs(&x, n, h, eps, &hdid, &hnext, derivs)
		if errcode != 0 {
			break
		}
		if float64(float64(x-x2)*float64(x2-x1)) >= 0.0 {
			for i := int32(0); i < n; i++ {
				ystart[i] = ovY[i]
			}
			return 0
		}
		if M.Fabs(hnext) <= 0.0 {
			return 2
		}
		h = hnext
	}
	return 3
}

func ovRkqs(x *float64, n int32, htry float64, eps float64, hdid *float64, hnext *float64, derivs OdeDerivs) int32 {
	xold := *x
	h := htry
	for {
		ovRkck(xold, n, h, derivs)
		errmax := 0.0
		for i := int32(0); i < n; i++ {
			err := M.Fabs(float64(ovYerr[i] / ovYscal[i]))
			if err > errmax {
				errmax = err
			}
		}
		errmax = float64(errmax / eps)
		if errmax > 1.0 {
			htemp := float64(float64(float64(ovSAFETY)*h) * M.Pow(errmax, ovPSHRNK))
			if h >= 0 {
				if htemp > float64(0.1*h) {
					h = htemp
				} else {
					h = float64(0.1 * h)
				}
			} else {
				if htemp < float64(0.1*h) {
					h = htemp
				} else {
					h = float64(0.1 * h)
				}
			}
			xnew := float64(xold + h)
			if xnew == xold {
				return 2
			}
			continue
		}
		if errmax > ovERRCON {
			*hnext = float64(float64(float64(ovSAFETY)*h) * M.Pow(errmax, ovPGROW))
		} else {
			*hnext = float64(5.0 * h)
		}
		*hdid = h
		*x = float64(*x + h)
		for i := int32(0); i < n; i++ {
			ovY[i] = ovYtemp[i]
		}
		return 0
	}
}

// The Cash-Karp constants. As in C, each is a double; the quotients are rounded once at compile time
// (Go evaluates constant expressions exactly and rounds the typed result, which equals C's single
// correctly rounded division of two exactly representable integers).
const (
	ovA2, ovA3, ovA4, ovA5, ovA6      float64 = 0.2, 0.3, 0.6, 1.0, 0.875
	ovB21, ovB31, ovB32               float64 = 0.2, 3.0 / 40.0, 9.0 / 40.0
	ovB41, ovB42, ovB43               float64 = 0.3, -0.9, 1.2
	ovB51, ovB52, ovB53, ovB54        float64 = -11.0 / 54.0, 2.5, -70.0 / 27.0, 35.0 / 27.0
	ovB61, ovB62, ovB63, ovB64, ovB65 float64 = 1631.0 / 55296.0, 175.0 / 512.0, 575.0 / 13824.0, 44275.0 / 110592.0, 253.0 / 4096.0
	ovC1, ovC3, ovC4, ovC6            float64 = 37.0 / 378.0, 250.0 / 621.0, 125.0 / 594.0, 512.0 / 1771.0
	ovDC5                             float64 = -277.00 / 14336.0
)

// dc1 = c1 - 2825/27648 etc. are C double subtractions of two rounded doubles: computed at run time.
var (
	ovDC1 = float64(ovC1 - float64(2825.0/27648.0))
	ovDC3 = float64(ovC3 - float64(18575.0/48384.0))
	ovDC4 = float64(ovC4 - float64(13525.0/55296.0))
	ovDC6 = float64(ovC6 - 0.25)
)

func ovRkck(x float64, n int32, h float64, derivs OdeDerivs) {
	y, dydx, yt := ovY, ovDydx, ovYtemp
	ak2, ak3, ak4, ak5, ak6 := ovAk2, ovAk3, ovAk4, ovAk5, ovAk6
	for i := int32(0); i < n; i++ {
		yt[i] = float64(y[i] + float64(float64(ovB21*h)*dydx[i]))
	}
	derivs(float64(x+float64(ovA2*h)), yt, ak2)
	for i := int32(0); i < n; i++ {
		yt[i] = float64(y[i] + float64(h*float64(float64(ovB31*dydx[i])+float64(ovB32*ak2[i]))))
	}
	derivs(float64(x+float64(ovA3*h)), yt, ak3)
	for i := int32(0); i < n; i++ {
		yt[i] = float64(y[i] + float64(h*float64(float64(float64(ovB41*dydx[i])+float64(ovB42*ak2[i]))+float64(ovB43*ak3[i]))))
	}
	derivs(float64(x+float64(ovA4*h)), yt, ak4)
	for i := int32(0); i < n; i++ {
		yt[i] = float64(y[i] + float64(h*float64(float64(float64(float64(ovB51*dydx[i])+float64(ovB52*ak2[i]))+
			float64(ovB53*ak3[i]))+float64(ovB54*ak4[i]))))
	}
	derivs(float64(x+float64(ovA5*h)), yt, ak5)
	for i := int32(0); i < n; i++ {
		yt[i] = float64(y[i] + float64(h*float64(float64(float64(float64(float64(ovB61*dydx[i])+float64(ovB62*ak2[i]))+
			float64(ovB63*ak3[i]))+float64(ovB64*ak4[i]))+float64(ovB65*ak5[i]))))
	}
	derivs(float64(x+float64(ovA6*h)), yt, ak6)
	for i := int32(0); i < n; i++ {
		yt[i] = float64(y[i] + float64(h*float64(float64(float64(float64(ovC1*dydx[i])+float64(ovC3*ak3[i]))+
			float64(ovC4*ak4[i]))+float64(ovC6*ak6[i]))))
	}
	for i := int32(0); i < n; i++ {
		ovYerr[i] = float64(h * float64(float64(float64(float64(float64(ovDC1*dydx[i])+float64(ovDC3*ak3[i]))+
			float64(ovDC4*ak4[i]))+float64(ovDC5*ak5[i]))+float64(ovDC6*ak6[i])))
	}
}
