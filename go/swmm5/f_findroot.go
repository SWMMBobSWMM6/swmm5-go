// Generated from EPA SWMM 5.2.4 (findroot.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

// findroot.c:20
func findroot_Newton(x1 float64, x2 float64, rts *float64, xacc float64, func_2 FP0, p any) int32 {
	var (
		j     int32
		n_2   int32
		df    float64
		dx    float64
		dxold float64
		f     float64
		x     float64
		temp  float64
		xhi   float64
		xlo   float64
	)
	_, _, _, _, _, _, _, _, _, _ = j, n_2, df, dx, dxold, f, x, temp, xhi, xlo
	j = 0
	n_2 = 0
	df = 0
	dx = 0
	dxold = 0
	f = 0
	x = 0
	temp = 0
	xhi = 0
	xlo = 0
	x = (*rts)
	xlo = x1
	xhi = x2
	dxold = M.Fabs(float64(x2 - x1))
	dx = dxold
	func_2(x, &f, &df, p)
	n_2++
	for j = 1; j <= 60; j++ {
		if (float64((float64(float64((float64(x-xhi))*df)-f))*(float64(float64((float64(x-xlo))*df)-f))) >= float64(0.0)) || (M.Fabs(float64(float64(2.0)*f)) > M.Fabs(float64(dxold*df))) {
			dxold = dx
			dx = float64(float64(0.5) * (float64(xhi - xlo)))
			x = float64(xlo + dx)
			if xlo == x {
				break
			}
		} else {
			dxold = dx
			dx = float64(f / df)
			temp = x
			x = float64(x - dx)
			if temp == x {
				break
			}
		}
		if M.Fabs(dx) < xacc {
			break
		}
		func_2(x, &f, &df, p)
		n_2++
		if f < float64(0.0) {
			xlo = x
		} else {
			xhi = x
		}
	}
	*rts = x
	if n_2 <= 60 {
		return n_2
	} else {
		return 0
	}
	return 0
}

// findroot.c:91
func findroot_Ridder(x1 float64, x2 float64, xacc float64, func_2 FP1, p any) float64 {
	var (
		j    int32
		ans  float64
		fhi  float64
		flo  float64
		fm   float64
		fnew float64
		s    float64
		xhi  float64
		xlo  float64
		xm   float64
		xnew float64
	)
	_, _, _, _, _, _, _, _, _, _, _ = j, ans, fhi, flo, fm, fnew, s, xhi, xlo, xm, xnew
	j = 0
	ans = 0
	fhi = 0
	flo = 0
	fm = 0
	fnew = 0
	s = 0
	xhi = 0
	xlo = 0
	xm = 0
	xnew = 0
	flo = func_2(x1, p)
	fhi = func_2(x2, p)
	if flo == float64(0.0) {
		return x1
	}
	if fhi == float64(0.0) {
		return x2
	}
	ans = float64(float64(0.5) * (float64(x1 + x2)))
	if ((flo > float64(0.0)) && (fhi < float64(0.0))) || ((flo < float64(0.0)) && (fhi > float64(0.0))) {
		xlo = x1
		xhi = x2
		for j = 1; j <= 60; j++ {
			xm = float64(float64(0.5) * (float64(xlo + xhi)))
			fm = func_2(xm, p)
			s = M.Sqrt(float64(float64(fm*fm) - float64(flo*fhi)))
			if s == float64(0.0) {
				return ans
			}
			xnew = float64(xm + float64((float64(xm-xlo))*(float64(float64((R.If[float64]((flo >= fhi), float64(1.0), (-(float64(1.0)))))*fm)/s))))
			if M.Fabs(float64(xnew-ans)) <= xacc {
				break
			}
			ans = xnew
			fnew = func_2(ans, p)
			if (R.If[float64](((fnew) >= float64(0.0)), M.Fabs(fm), (-(M.Fabs(fm))))) != fm {
				xlo = xm
				flo = fm
				xhi = ans
				fhi = fnew
			} else if (R.If[float64](((fnew) >= float64(0.0)), M.Fabs(flo), (-(M.Fabs(flo))))) != flo {
				xhi = ans
				fhi = fnew
			} else if (R.If[float64](((fnew) >= float64(0.0)), M.Fabs(fhi), (-(M.Fabs(fhi))))) != fhi {
				xlo = ans
				flo = fnew
			} else {
				return ans
			}
			if M.Fabs(float64(xhi-xlo)) <= xacc {
				return ans
			}
		}
		return ans
	}
	return (-(float64(1.0e20)))
}
