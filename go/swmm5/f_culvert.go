// Generated from EPA SWMM 5.2.4 (culvert.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	M "swmm5go/glibm"
)

var MAX_CULVERT_CODE int32 = 57
var Params = [58][5]float64{[5]float64{float64(0.0), float64(0.0), float64(0.0), float64(0.0), float64(0.00)}, [5]float64{float64(1.0), float64(0.0098), float64(2.00), float64(0.0398), float64(0.67)}, [5]float64{float64(1.0), float64(0.0018), float64(2.00), float64(0.0292), float64(0.74)}, [5]float64{float64(1.0), float64(0.0045), float64(2.00), float64(0.0317), float64(0.69)}, [5]float64{float64(1.0), float64(0.0078), float64(2.00), float64(0.0379), float64(0.69)}, [5]float64{float64(1.0), float64(0.0210), float64(1.33), float64(0.0463), float64(0.75)}, [5]float64{float64(1.0), float64(0.0340), float64(1.50), float64(0.0553), float64(0.54)}, [5]float64{float64(1.0), float64(0.0018), float64(2.50), float64(0.0300), float64(0.74)}, [5]float64{float64(1.0), float64(0.0018), float64(2.50), float64(0.0243), float64(0.83)}, [5]float64{float64(1.0), float64(0.026), float64(1.0), float64(0.0347), float64(0.81)}, [5]float64{float64(1.0), float64(0.061), float64(0.75), float64(0.0400), float64(0.80)}, [5]float64{float64(1.0), float64(0.061), float64(0.75), float64(0.0423), float64(0.82)}, [5]float64{float64(2.0), float64(0.510), float64(0.667), float64(0.0309), float64(0.80)}, [5]float64{float64(2.0), float64(0.486), float64(0.667), float64(0.0249), float64(0.83)}, [5]float64{float64(2.0), float64(0.515), float64(0.667), float64(0.0375), float64(0.79)}, [5]float64{float64(2.0), float64(0.495), float64(0.667), float64(0.0314), float64(0.82)}, [5]float64{float64(2.0), float64(0.486), float64(0.667), float64(0.0252), float64(0.865)}, [5]float64{float64(2.0), float64(0.545), float64(0.667), float64(0.04505), float64(0.73)}, [5]float64{float64(2.0), float64(0.533), float64(0.667), float64(0.0425), float64(0.705)}, [5]float64{float64(2.0), float64(0.522), float64(0.667), float64(0.0402), float64(0.68)}, [5]float64{float64(2.0), float64(0.498), float64(0.667), float64(0.0327), float64(0.75)}, [5]float64{float64(2.0), float64(0.497), float64(0.667), float64(0.0339), float64(0.803)}, [5]float64{float64(2.0), float64(0.493), float64(0.667), float64(0.0361), float64(0.806)}, [5]float64{float64(2.0), float64(0.495), float64(0.667), float64(0.0386), float64(0.71)}, [5]float64{float64(2.0), float64(0.497), float64(0.667), float64(0.0302), float64(0.835)}, [5]float64{float64(2.0), float64(0.495), float64(0.667), float64(0.0252), float64(0.881)}, [5]float64{float64(2.0), float64(0.493), float64(0.667), float64(0.0227), float64(0.887)}, [5]float64{float64(1.0), float64(0.0083), float64(2.00), float64(0.0379), float64(0.69)}, [5]float64{float64(1.0), float64(0.0145), float64(1.75), float64(0.0419), float64(0.64)}, [5]float64{float64(1.0), float64(0.0340), float64(1.50), float64(0.0496), float64(0.57)}, [5]float64{float64(1.0), float64(0.0100), float64(2.00), float64(0.0398), float64(0.67)}, [5]float64{float64(1.0), float64(0.0018), float64(2.50), float64(0.0292), float64(0.74)}, [5]float64{float64(1.0), float64(0.0045), float64(2.00), float64(0.0317), float64(0.69)}, [5]float64{float64(1.0), float64(0.0100), float64(2.00), float64(0.0398), float64(0.67)}, [5]float64{float64(1.0), float64(0.0018), float64(2.50), float64(0.0292), float64(0.74)}, [5]float64{float64(1.0), float64(0.0095), float64(2.00), float64(0.0317), float64(0.69)}, [5]float64{float64(1.0), float64(0.0083), float64(2.00), float64(0.0379), float64(0.69)}, [5]float64{float64(1.0), float64(0.0300), float64(1.00), float64(0.0463), float64(0.75)}, [5]float64{float64(1.0), float64(0.0340), float64(1.50), float64(0.0496), float64(0.57)}, [5]float64{float64(1.0), float64(0.0300), float64(1.50), float64(0.0496), float64(0.57)}, [5]float64{float64(1.0), float64(0.0088), float64(2.00), float64(0.0368), float64(0.68)}, [5]float64{float64(1.0), float64(0.0030), float64(2.00), float64(0.0269), float64(0.77)}, [5]float64{float64(1.0), float64(0.0300), float64(1.50), float64(0.0496), float64(0.57)}, [5]float64{float64(1.0), float64(0.0088), float64(2.00), float64(0.0368), float64(0.68)}, [5]float64{float64(1.0), float64(0.0030), float64(2.00), float64(0.0269), float64(0.77)}, [5]float64{float64(1.0), float64(0.0083), float64(2.00), float64(0.0379), float64(0.69)}, [5]float64{float64(1.0), float64(0.0300), float64(1.00), float64(0.0473), float64(0.75)}, [5]float64{float64(1.0), float64(0.0340), float64(1.50), float64(0.0496), float64(0.57)}, [5]float64{float64(2.0), float64(0.534), float64(0.555), float64(0.0196), float64(0.90)}, [5]float64{float64(2.0), float64(0.519), float64(0.640), float64(0.0210), float64(0.90)}, [5]float64{float64(2.0), float64(0.536), float64(0.622), float64(0.0368), float64(0.83)}, [5]float64{float64(2.0), float64(0.5035), float64(0.719), float64(0.0478), float64(0.80)}, [5]float64{float64(2.0), float64(0.547), float64(0.800), float64(0.0598), float64(0.75)}, [5]float64{float64(2.0), float64(0.475), float64(0.667), float64(0.0179), float64(0.97)}, [5]float64{float64(2.0), float64(0.560), float64(0.667), float64(0.0446), float64(0.85)}, [5]float64{float64(2.0), float64(0.560), float64(0.667), float64(0.0378), float64(0.87)}, [5]float64{float64(2.0), float64(0.500), float64(0.667), float64(0.0446), float64(0.65)}, [5]float64{float64(2.0), float64(0.500), float64(0.667), float64(0.0378), float64(0.71)}}

// culvert.c:174
func culvert_getInflow(j int32, q0 float64, h float64) float64 {
	var (
		code      int32
		k         int32
		condition int32
		y_2       float64
		y1        float64
		y2        float64
		q         float64
		culvert   TCulvert
	)
	_, _, _, _, _, _, _, _ = code, k, condition, y_2, y1, y2, q, culvert
	code = 0
	k = 0
	condition = 0
	y_2 = 0
	y1 = 0
	y2 = 0
	q = 0
	culvert = (TCulvert{})
	if Link[j].type_ != CONDUIT {
		return q0
	}
	culvert.xsect = &Link[j].xsect
	code = culvert.xsect.culvertCode
	if (code <= 0) || (code > MAX_CULVERT_CODE) {
		return q0
	}
	k = Link[j].subIndex
	culvert.yFull = culvert.xsect.yFull
	culvert.ad = float64(culvert.xsect.aFull * M.Sqrt(culvert.yFull))
	switch code {
	case 5, 37, 46:
		culvert.scf = float64((-(float64(7.0))) * Conduit[k].slope)
		break
	default:
		culvert.scf = float64(float64(0.5) * Conduit[k].slope)
	}
	y_2 = float64(h - (float64(Node[Link[j].node1].invertElev + Link[j].offset1)))
	y2 = float64(culvert.yFull * (float64(float64(float64(float64(16.0)*Params[code][C])+Params[code][Y]) - culvert.scf)))
	if y_2 >= y2 {
		q = getSubmergedFlow(code, y_2, &culvert)
		condition = 2
	} else {
		y1 = float64(float64(0.95) * culvert.yFull)
		if y_2 <= y1 {
			q = getUnsubmergedFlow(code, y_2, &culvert)
			condition = 1
		} else {
			q = getTransitionFlow(code, y_2, y1, y2, &culvert)
			condition = 0
		}
	}
	if q < q0 {
		Link[j].inletControl = 1
		Link[j].dqdh = culvert.dQdH
		return q
	} else {
		return q0
	}
	return 0
}

// culvert.c:257
func getUnsubmergedFlow(code int32, h float64, culvert *TCulvert) float64 {
	var (
		arg float64
		q   float64
	)
	_, _ = arg, q
	arg = 0
	q = 0
	culvert.kk = Params[code][K]
	culvert.mm = Params[code][M_]
	arg = float64(float64(h/culvert.yFull) / culvert.kk)
	if Params[code][FORM] == float64(1.0) {
		q = getForm1Flow(h, culvert)
	} else {
		q = float64(culvert.ad * M.Pow(arg, float64(float64(1.0)/culvert.mm)))
	}
	culvert.dQdH = float64(float64(q/h) / culvert.mm)
	return q
}

// culvert.c:288
func getSubmergedFlow(code int32, h float64, culvert *TCulvert) float64 {
	var (
		cc  float64
		yy  float64
		arg float64
		q   float64
	)
	_, _, _, _ = cc, yy, arg, q
	cc = Params[code][C]
	yy = Params[code][Y]
	arg = float64((float64(float64(float64(h/culvert.yFull)-yy) + culvert.scf)) / cc)
	q = 0
	if arg <= float64(0.0) {
		culvert.dQdH = float64(0.0)
		return float64(10000000000.0)
	}
	q = float64(M.Sqrt(arg) * culvert.ad)
	culvert.dQdH = float64(float64(float64(float64(float64(0.5)*q)/arg)/culvert.yFull) / cc)
	return q
}

// culvert.c:316
func getTransitionFlow(code int32, h float64, h1 float64, h2 float64, culvert *TCulvert) float64 {
	var (
		q1 float64
		q2 float64
		q  float64
	)
	_, _, _ = q1, q2, q
	q1 = getUnsubmergedFlow(code, h1, culvert)
	q2 = getSubmergedFlow(code, h2, culvert)
	q = float64(q1 + float64(float64((float64(q2-q1))*(float64(h-h1)))/(float64(h2-h1))))
	culvert.dQdH = float64((float64(q2 - q1)) / (float64(h2 - h1)))
	return q
}

// culvert.c:339
func getForm1Flow(h float64, culvert *TCulvert) float64 {
	culvert.hPlus = float64(float64(h/culvert.yFull) + culvert.scf)
	findroot_Ridder(float64(float64(0.01)*h), h, float64(0.001), form1Eqn, culvert)
	return culvert.qc
}

// culvert.c:363
func form1Eqn(yc float64, p any) float64 {
	var (
		ac      float64
		wc      float64
		yh      float64
		culvert *TCulvert
	)
	_, _, _, _ = ac, wc, yh, culvert
	ac = 0
	wc = 0
	yh = 0
	culvert = p.(*TCulvert)
	ac = xsect_getAofY(culvert.xsect, yc)
	wc = xsect_getWofY(culvert.xsect, yc)
	yh = float64(ac / wc)
	culvert.qc = float64(ac * M.Sqrt(float64(float64(32.2)*yh)))
	return float64(float64(float64(culvert.hPlus-float64(yc/culvert.yFull))-float64(float64(yh/float64(2.0))/culvert.yFull)) - float64(culvert.kk*M.Pow(float64(culvert.qc/culvert.ad), culvert.mm)))
}
