// Generated from EPA SWMM 5.2.4 (forcmain.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	M "swmm5go/glibm"
)

var VISCOS float64 = float64(1.1e-5)

// forcmain.c:37
func forcemain_getEquivN(j int32, k int32) float64 {
	var (
		xsect_2 TXsect
		f       float64
		d       float64
	)
	_, _, _ = xsect_2, f, d
	xsect_2 = Link[j].xsect
	f = 0
	d = xsect_2.yFull
	switch ForceMainEqn {
	case H_W:
		return float64(float64(float64(1.067)/xsect_2.rBot) * M.Pow(float64(d/Conduit[k].slope), float64(0.04)))
	case D_W:
		f = forcemain_getFricFactor(xsect_2.rBot, float64(d/float64(4.0)), float64(1.0e12))
		return float64(M.Sqrt(float64(f/float64(185.0))) * M.Pow(d, (float64(float64(1.0)/float64(6.0)))))
	}
	return Conduit[k].roughness
}

// forcmain.c:64
func forcemain_getRoughFactor(j int32, lengthFactor float64) float64 {
	var (
		xsect_2 TXsect
		r       float64
	)
	_, _ = xsect_2, r
	xsect_2 = Link[j].xsect
	r = 0
	switch ForceMainEqn {
	case H_W:
		r = float64(float64(float64(1.318)*xsect_2.rBot) * M.Pow(lengthFactor, float64(0.54)))
		return float64(float64(32.2) / M.Pow(r, float64(1.852)))
	case D_W:
		return float64(float64(float64(1.0)/float64(8.0)) / lengthFactor)
	}
	return float64(0.0)
}

// forcmain.c:88
func forcemain_getFricSlope(j int32, v float64, hrad float64) float64 {
	var (
		re      float64
		f       float64
		xsect_2 TXsect
	)
	_, _, _ = re, f, xsect_2
	re = 0
	f = 0
	xsect_2 = Link[j].xsect
	switch ForceMainEqn {
	case H_W:
		return float64(float64(xsect_2.sBot*M.Pow(v, float64(0.852))) / M.Pow(hrad, float64(1.1667)))
	case D_W:
		re = forcemain_getReynolds(v, hrad)
		f = forcemain_getFricFactor(xsect_2.rBot, hrad, re)
		return float64(float64(float64(f*xsect_2.sBot)*v) / hrad)
	}
	return float64(0.0)
}

// forcmain.c:117
func forcemain_getReynolds(v float64, hrad float64) float64 {
	return float64(float64(float64(float64(4.0)*hrad)*v) / VISCOS)
}

// forcmain.c:130
func forcemain_getFricFactor(e float64, hrad float64, re float64) float64 {
	var (
		f float64
	)
	_ = f
	f = 0
	if re < float64(10.0) {
		re = float64(10.0)
	}
	if re <= float64(2000.0) {
		f = float64(float64(64.0) / re)
	} else if re < float64(4000.0) {
		f = forcemain_getFricFactor(e, hrad, float64(4000.0))
		f = float64(float64(0.032) + float64(float64((float64(f-float64(0.032)))*(float64(re-float64(2000.0))))/float64(2000.0)))
	} else {
		f = float64(float64(e/float64(3.7)) / (float64(float64(4.0) * hrad)))
		if re < float64(1.0e10) {
			f = float64(f + (float64(float64(5.74) / M.Pow(re, float64(0.9)))))
		}
		f = M.Log10(f)
		f = float64(float64(float64(0.25)/f) / f)
	}
	return f
}
