// Generated from EPA SWMM 5.2.4 (kinwave.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

var WX float64 = float64(0.6)
var WT float64 = float64(0.6)
var EPSIL float64 = float64(0.001)
var Beta1 float64
var C1 float64
var C2 float64
var Afull float64
var Qfull float64
var pXsect *TXsect

// kinwave.c:57
func kinwave_execute(j int32, qinflow *float64, qoutflow *float64, tStep float64) int32 {
	var (
		k      int32
		result int32
		dxdt   float64
		dq     float64
		ain    float64
		aout   float64
		qin    float64
		qout   float64
		a1     float64
		a2     float64
		q1     float64
		q2     float64
		q3     float64
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = k, result, dxdt, dq, ain, aout, qin, qout, a1, a2, q1, q2
	_ = q3
	k = 0
	result = 1
	dxdt = 0
	dq = 0
	ain = 0
	aout = 0
	qin = 0
	qout = 0
	a1 = 0
	a2 = 0
	q1 = 0
	q2 = 0
	q3 = 0
	*qoutflow = (*qinflow)
	if Link[j].type_ != CONDUIT {
		return result
	}
	if Link[j].xsect.type_ == DUMMY {
		return result
	}
	pXsect = &Link[j].xsect
	Qfull = Link[j].qFull
	Afull = Link[j].xsect.aFull
	k = Link[j].subIndex
	Beta1 = float64(Conduit[k].beta / Qfull)
	q1 = float64(Conduit[k].q1 / Qfull)
	q2 = float64(Conduit[k].q2 / Qfull)
	qin = float64(float64((*qinflow)/float64(Conduit[k].barrels)) / Qfull)
	q3 = float64(link_getLossRate(j, KW, float64(qin*Qfull), tStep) / Qfull)
	a1 = float64(Conduit[k].a1 / Afull)
	a2 = float64(Conduit[k].a2 / Afull)
	if qin >= float64(1.0) {
		ain = float64(1.0)
	} else {
		ain = float64(xsect_getAofS(pXsect, float64(qin/Beta1)) / Afull)
	}
	if (qin <= float64(1e-06)) && (q2 <= float64(1e-06)) {
		qout = float64(0.0)
		aout = float64(0.0)
	} else {
		dxdt = float64(float64(float64(link_getLength(j)/tStep)*Afull) / Qfull)
		dq = float64(q2 - q1)
		C1 = float64(float64(dxdt*WT) / WX)
		C2 = float64((float64(float64(1.0) - WT)) * (float64(ain - a1)))
		C2 = float64(C2 - float64(WT*a2))
		C2 = float64(float64(C2*dxdt) / WX)
		C2 = float64(float64(C2+float64(float64((float64(float64(1.0)-WX))/WX)*dq)) - qin)
		C2 = float64(C2 + float64(q3/WX))
		aout = a2
		result = solveContinuity(qin, ain, &aout)
		if result == (-1) {
			report_writeErrorMsg(ERR_KINWAVE, Link[j].ID)
			return 1
		}
		if result <= 0 {
			result = 1
		}
		qout = float64(Beta1 * xsect_getSofA(pXsect, float64(aout*Afull)))
		if qin > float64(1.0) {
			qin = float64(1.0)
		}
	}
	Conduit[k].q1 = float64(qin * Qfull)
	Conduit[k].a1 = float64(ain * Afull)
	Conduit[k].q2 = float64(qout * Qfull)
	Conduit[k].a2 = float64(aout * Afull)
	Conduit[k].fullState = link_getFullState(Conduit[k].a1, Conduit[k].a2, Afull)
	*qinflow = float64(Conduit[k].q1 * float64(Conduit[k].barrels))
	*qoutflow = float64(Conduit[k].q2 * float64(Conduit[k].barrels))
	return result
}

// kinwave.c:170
func solveContinuity(qin float64, ain float64, aout *float64) int32 {
	var (
		n_2  int32
		aLo  float64
		aHi  float64
		aTmp float64
		fLo  float64
		fHi  float64
		tol  float64
	)
	_, _, _, _, _, _, _ = n_2, aLo, aHi, aTmp, fLo, fHi, tol
	n_2 = 0
	aLo = 0
	aHi = 0
	aTmp = 0
	fLo = 0
	fHi = 0
	tol = EPSIL
	aHi = float64(1.0)
	fHi = float64(float64(float64(1.0)+C1) + C2)
	aLo = float64(xsect_getAmax(pXsect) / Afull)
	if aLo < aHi {
		fLo = float64(float64((float64(Beta1*pXsect.sMax))+(float64(C1*aLo))) + C2)
	} else {
		fLo = fHi
	}
	if float64(fHi*fLo) > float64(0.0) {
		aHi = aLo
		fHi = fLo
		aLo = float64(0.0)
		fLo = C2
	}
	if float64(fHi*fLo) <= float64(0.0) {
		if ((*aout) < aLo) || ((*aout) > aHi) {
			*aout = float64(float64(0.5) * (float64(aLo + aHi)))
		}
		if fLo > fHi {
			aTmp = aLo
			aLo = aHi
			aHi = aTmp
		}
		n_2 = findroot_Newton(aLo, aHi, aout, tol, evalContinuity, nil)
		if n_2 <= 0 {
			n_2 = (-1)
		}
	} else if fLo < float64(0.0) {
		if qin > float64(1.0) {
			*aout = ain
		} else {
			*aout = float64(1.0)
		}
		n_2 = (-2)
	} else if fLo > float64(0) {
		*aout = float64(0.0)
		n_2 = (-3)
	} else {
		n_2 = (-1)
	}
	return n_2
}

// kinwave.c:261
func evalContinuity(a_2 float64, f *float64, df *float64, p any) {
	*f = float64(float64((float64(Beta1*xsect_getSofA(pXsect, float64(a_2*Afull))))+(float64(C1*a_2))) + C2)
	*df = float64((float64(float64(Beta1*Afull) * xsect_getdSdA(pXsect, float64(a_2*Afull)))) + C1)
}
