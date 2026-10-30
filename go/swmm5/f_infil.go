// Generated from EPA SWMM 5.2.4 (infil.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var Infil []TInfil
var Fumax float64
var InfilFactor float64

// infil.c:101
func infil_create(n_2 int32) {
	Infil = make([]TInfil, int(uint64(n_2)))
	if Infil == nil {
		ErrorCode = ERR_MEMORY
	}
	InfilFactor = float64(1.0)
	return
}

// infil.c:116
func infil_delete() {
	{
		if Infil != nil {
			_ = Infil
			Infil = nil
		}
	}
}

// infil.c:128
func infil_readParams(m int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		i      int32
		j      int32
		n_2    int32
		status int32
		x      [5]float64
	)
	_, _, _, _, _ = i, j, n_2, status, x
	i = 0
	j = 0
	n_2 = 0
	status = 0
	x = ([5]float64{})
	j = project_findObject(SUBCATCH, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	i = findmatch(tok[(ntoks-1)], InfilModelWords[:])
	if i >= 0 {
		m = i
		ntoks--
	}
	if m == HORTON {
		n_2 = 5
	} else if m == MOD_HORTON {
		n_2 = 5
	} else if m == GREEN_AMPT {
		n_2 = 4
	} else if m == MOD_GREEN_AMPT {
		n_2 = 4
	} else if m == CURVE_NUMBER {
		n_2 = 4
	} else {
		return 0
	}
	if ntoks < n_2 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	for i = 0; i < 5; i++ {
		x[i] = float64(0.0)
	}
	for i = 1; i < n_2; i++ {
		if !(getDouble(tok[i], &x[(i-1)]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[i])
		}
	}
	if ((m == HORTON) || (m == MOD_HORTON)) && (ntoks > n_2) {
		if !(getDouble(tok[n_2], &x[(n_2-1)]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[n_2])
		}
	}
	Subcatch[j].infil = j
	Subcatch[j].infilModel = m
	switch m {
	case HORTON, MOD_HORTON:
		status = horton_setParams(&Infil[j].horton, x[:])
		break
	case GREEN_AMPT, MOD_GREEN_AMPT:
		status = grnampt_setParams(&Infil[j].grnAmpt, x[:])
		break
	case CURVE_NUMBER:
		status = curvenum_setParams(&Infil[j].curveNum, x[:])
		break
	default:
		status = 1
	}
	if !(status != 0) {
		return error_setInpError(ERR_INFIL_PARAMS, L0)
	}
	return 0
}

// infil.c:201
func infil_initState(j int32) {
	switch Subcatch[j].infilModel {
	case HORTON, MOD_HORTON:
		horton_initState(&Infil[j].horton)
		break
	case GREEN_AMPT, MOD_GREEN_AMPT:
		grnampt_initState(&Infil[j].grnAmpt)
		break
	case CURVE_NUMBER:
		curvenum_initState(&Infil[j].curveNum)
		break
	}
}

// infil.c:221
func infil_getState(j int32, x []float64) {
	switch Subcatch[j].infilModel {
	case HORTON, MOD_HORTON:
		horton_getState(&Infil[j].horton, x)
		break
	case GREEN_AMPT, MOD_GREEN_AMPT:
		grnampt_getState(&Infil[j].grnAmpt, x)
		break
	case CURVE_NUMBER:
		curvenum_getState(&Infil[j].curveNum, x)
		break
	}
}

// infil.c:241
func infil_setState(j int32, x []float64) {
	switch Subcatch[j].infilModel {
	case HORTON, MOD_HORTON:
		horton_setState(&Infil[j].horton, x)
		break
	case GREEN_AMPT, MOD_GREEN_AMPT:
		grnampt_setState(&Infil[j].grnAmpt, x)
		break
	case CURVE_NUMBER:
		curvenum_setState(&Infil[j].curveNum, x)
		break
	}
}

// infil.c:262
func infil_setInfilFactor(j int32) {
	var (
		m int32
		p int32
	)
	_, _ = m, p
	m = 0
	p = 0
	InfilFactor = Adjust.hydconFactor
	if j >= 0 {
		p = Subcatch[j].infilPattern
		if (p >= 0) && (Pattern[p].type_ == MONTHLY_PATTERN) {
			m = (datetime_monthOfYear(getDateTime(OldRunoffTime)) - 1)
			InfilFactor = Pattern[p].factor[m]
		}
	}
}

// infil.c:288
func infil_getInfil(j int32, tstep float64, rainfall float64, runon float64, depth float64) float64 {
	switch Subcatch[j].infilModel {
	case HORTON:
		return horton_getInfil(&Infil[j].horton, tstep, float64(rainfall+runon), depth)
	case MOD_HORTON:
		return modHorton_getInfil(&Infil[j].horton, tstep, float64(rainfall+runon), depth)
	case GREEN_AMPT, MOD_GREEN_AMPT:
		return grnampt_getInfil(&Infil[j].grnAmpt, tstep, float64(rainfall+runon), depth, Subcatch[j].infilModel)
	case CURVE_NUMBER:
		depth = float64(depth + (float64(runon * tstep)))
		return curvenum_getInfil(&Infil[j].curveNum, tstep, rainfall, depth)
	default:
		return float64(0.0)
	}
	return 0
}

// infil.c:325
func horton_setParams(infil *THorton, p []float64) int32 {
	var (
		k int32
	)
	_ = k
	k = 0
	for k = 0; k < 5; k++ {
		if p[k] < float64(0.0) {
			return 0
		}
	}
	infil.f0 = float64(p[0] / UCF(RAINFALL))
	infil.fmin = float64(p[1] / UCF(RAINFALL))
	infil.decay = float64(p[2] / float64(3600.0))
	if p[3] == float64(0.0) {
		p[3] = float64(1e-06)
	}
	infil.regen = float64(float64((-(M.Log(float64(float64(1.0) - float64(0.98)))))/p[3]) / float64(86400))
	infil.Fmax = float64(p[4] / UCF(RAINDEPTH))
	if infil.f0 < infil.fmin {
		return 0
	}
	return 1
}

// infil.c:357
func horton_initState(infil *THorton) {
	infil.tp = float64(0.0)
	infil.Fe = float64(0.0)
}

// infil.c:370
func horton_getState(infil *THorton, x []float64) {
	x[0] = infil.tp
	x[1] = infil.Fe
}

// infil.c:376
func horton_setState(infil *THorton, x []float64) {
	infil.tp = x[0]
	infil.Fe = x[1]
}

// infil.c:384
func horton_getInfil(infil *THorton, tstep float64, irate float64, depth float64) float64 {
	var (
		iter int32
		fa   float64
		fp   float64
		Fp   float64
		F1   float64
		t1   float64
		tlim float64
		ex   float64
		kt   float64
		FF   float64
		FF1  float64
		r    float64
		f0   float64
		fmin float64
		Fmax float64
		tp   float64
		df   float64
		kd   float64
		kr   float64
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = iter, fa, fp, Fp, F1, t1, tlim, ex, kt, FF, FF1, r
	_, _, _, _, _, _, _ = f0, fmin, Fmax, tp, df, kd, kr
	iter = 0
	fa = 0
	fp = float64(0.0)
	Fp = 0
	F1 = 0
	t1 = 0
	tlim = 0
	ex = 0
	kt = 0
	FF = 0
	FF1 = 0
	r = 0
	f0 = float64(infil.f0 * InfilFactor)
	fmin = float64(infil.fmin * InfilFactor)
	Fmax = infil.Fmax
	tp = infil.tp
	df = float64(f0 - fmin)
	kd = infil.decay
	kr = float64(infil.regen * Evap.recoveryFactor)
	if ((df < float64(0.0)) || (kd < float64(0.0))) || (kr < float64(0.0)) {
		return float64(0.0)
	}
	if (df == float64(0.0)) || (kd == float64(0.0)) {
		fp = f0
		fa = float64(irate + float64(depth/tstep))
		if fp > fa {
			fp = fa
		}
		if (float64(0.0)) >= (fp) {
			return (float64(0.0))
		}
		return (fp)
	}
	fa = float64(irate + float64(depth/tstep))
	if fa > float64(1e-10) {
		t1 = float64(tp + tstep)
		tlim = float64(float64(16.0) / kd)
		if tp >= tlim {
			Fp = float64(float64(fmin*tp) + float64(df/kd))
			F1 = float64(Fp + float64(fmin*tstep))
		} else {
			Fp = float64(float64(fmin*tp) + float64(float64(df/kd)*(float64(float64(1.0)-M.Exp(float64((-kd)*tp))))))
			F1 = float64(float64(fmin*t1) + float64(float64(df/kd)*(float64(float64(1.0)-M.Exp(float64((-kd)*t1))))))
		}
		fp = float64((float64(F1 - Fp)) / tstep)
		if (fp) >= (fmin) {
			fp = (fp)
		} else {
			fp = (fmin)
		}
		if fp > fa {
			fp = fa
		}
		if t1 > tlim {
			tp = t1
		} else if fp < fa {
			tp = t1
		} else {
			F1 = float64(Fp + float64(fp*tstep))
			tp = float64(tp + float64(tstep/float64(2.0)))
			for iter = 1; iter <= 20; iter++ {
				if (float64(60.0)) <= (float64(kd * tp)) {
					kt = (float64(60.0))
				} else {
					kt = (float64(kd * tp))
				}
				ex = M.Exp((-kt))
				FF = float64(float64(float64(fmin*tp)+float64(float64(df/kd)*(float64(float64(1.0)-ex)))) - F1)
				FF1 = float64(fmin + float64(df*ex))
				r = float64(FF / FF1)
				tp = float64(tp - r)
				if M.Fabs(r) <= float64(float64(0.001)*tstep) {
					break
				}
			}
		}
		if Fmax > float64(0.0) {
			if float64(infil.Fe+float64(fp*tstep)) > Fmax {
				fp = float64((float64(Fmax - infil.Fe)) / tstep)
			}
			if (fp) >= (float64(0.0)) {
				fp = (fp)
			} else {
				fp = (float64(0.0))
			}
			infil.Fe = float64(infil.Fe + (float64(fp * tstep)))
		}
	} else if kr > float64(0.0) {
		r = M.Exp(float64((-kr) * tstep))
		tp = float64(float64(1.0) - M.Exp(float64((-kd)*tp)))
		tp = float64((-(M.Log(float64(float64(1.0) - float64(r*tp))))) / kd)
		if Fmax > float64(0.0) {
			infil.Fe = float64(float64(fmin*tp) + float64((float64(df/kd))*(float64(float64(1.0)-M.Exp(float64((-kd)*tp))))))
		}
	}
	infil.tp = tp
	return fp
}

// infil.c:496
func modHorton_getInfil(infil *THorton, tstep float64, irate float64, depth float64) float64 {
	var (
		f    float64
		fp   float64
		fa   float64
		f0   float64
		fmin float64
		df   float64
		kd   float64
		kr   float64
	)
	_, _, _, _, _, _, _, _ = f, fp, fa, f0, fmin, df, kd, kr
	f = float64(0.0)
	fp = 0
	fa = 0
	f0 = float64(infil.f0 * InfilFactor)
	fmin = float64(infil.fmin * InfilFactor)
	df = float64(f0 - fmin)
	kd = infil.decay
	kr = float64(infil.regen * Evap.recoveryFactor)
	if ((df < float64(0.0)) || (kd < float64(0.0))) || (kr < float64(0.0)) {
		return float64(0.0)
	}
	if (df == float64(0.0)) || (kd == float64(0.0)) {
		fp = f0
		fa = float64(irate + float64(depth/tstep))
		if fp > fa {
			fp = fa
		}
		if (float64(0.0)) >= (fp) {
			return (float64(0.0))
		}
		return (fp)
	}
	fa = float64(irate + float64(depth/tstep))
	if fa > float64(1e-10) {
		if (infil.Fmax > float64(0.0)) && (infil.Fe >= infil.Fmax) {
			return float64(0.0)
		}
		fp = float64(f0 - float64(kd*infil.Fe))
		if (fp) >= (fmin) {
			fp = (fp)
		} else {
			fp = (fmin)
		}
		if (fa) <= (fp) {
			f = (fa)
		} else {
			f = (fp)
		}
		infil.Fe = float64(infil.Fe + (float64((R.If[float64](((float64(f - fmin)) >= (float64(0.0))), (float64(f - fmin)), (float64(0.0)))) * tstep)))
		if infil.Fmax > float64(0.0) {
			if (infil.Fe) >= (infil.Fmax) {
				infil.Fe = (infil.Fe)
			} else {
				infil.Fe = (infil.Fmax)
			}
		}
	} else if kr > float64(0.0) {
		infil.Fe = float64(infil.Fe * (M.Exp(float64((-kr) * tstep))))
		if (infil.Fe) >= (float64(0.0)) {
			infil.Fe = (infil.Fe)
		} else {
			infil.Fe = (float64(0.0))
		}
	}
	return f
}

// infil.c:559
func grnampt_getParams(j int32, p []float64) {
	p[0] = float64(Infil[j].grnAmpt.S * UCF(RAINDEPTH))
	p[1] = float64(Infil[j].grnAmpt.Ks * UCF(RAINFALL))
	p[2] = Infil[j].grnAmpt.IMDmax
}

// infil.c:574
func grnampt_setParams(infil *TGrnAmpt, p []float64) int32 {
	var (
		ksat float64
	)
	_ = ksat
	ksat = 0
	if (((p[0] < float64(0.0)) || (p[1] <= float64(0.0))) || (p[2] < float64(0.0))) || (p[2] > float64(1.0)) {
		return 0
	}
	infil.S = float64(p[0] / UCF(RAINDEPTH))
	infil.Ks = float64(p[1] / UCF(RAINFALL))
	infil.IMDmax = p[2]
	ksat = float64(float64(infil.Ks*float64(12.0)) * float64(3600.0))
	infil.Lu = float64(float64(float64(4.0)*M.Sqrt(ksat)) / float64(12.0))
	return 1
}

// infil.c:597
func grnampt_initState(infil *TGrnAmpt) {
	if infil == nil {
		return
	}
	infil.IMD = infil.IMDmax
	infil.Fu = float64(0.0)
	infil.F = float64(0.0)
	infil.Sat = 0
	infil.T = float64(0.0)
}

// infil.c:612
func grnampt_getState(infil *TGrnAmpt, x []float64) {
	x[0] = infil.IMD
	x[1] = infil.F
	x[2] = infil.Fu
	x[3] = float64(infil.Sat)
	x[4] = infil.T
}

// infil.c:621
func grnampt_setState(infil *TGrnAmpt, x []float64) {
	infil.IMD = x[0]
	infil.F = x[1]
	infil.Fu = x[2]
	infil.Sat = int32(int8(R.D2I(x[3])))
	infil.T = x[4]
}

// infil.c:632
func grnampt_getInfil(infil *TGrnAmpt, tstep float64, irate float64, depth float64, modelType int32) float64 {
	Fumax = float64(float64(infil.IMDmax*infil.Lu) * M.Sqrt(InfilFactor))
	infil.T = float64(infil.T - tstep)
	if infil.Sat != 0 {
		return grnampt_getSatInfil(infil, tstep, irate, depth)
	} else {
		return grnampt_getUnsatInfil(infil, tstep, irate, depth, modelType)
	}
	return 0
}

// infil.c:660
func grnampt_getUnsatInfil(infil *TGrnAmpt, tstep float64, irate float64, depth float64, modelType int32) float64 {
	var (
		ia float64
		c1 float64
		F2 float64
		dF float64
		Fs float64
		kr float64
		ts float64
		ks float64
		lu float64
	)
	_, _, _, _, _, _, _, _, _ = ia, c1, F2, dF, Fs, kr, ts, ks, lu
	ia = 0
	c1 = 0
	F2 = 0
	dF = 0
	Fs = 0
	kr = 0
	ts = 0
	ks = float64(infil.Ks * InfilFactor)
	lu = float64(infil.Lu * M.Sqrt(InfilFactor))
	ia = float64(irate + float64(depth/tstep))
	if ia < float64(1e-10) {
		ia = float64(0.0)
	}
	if ia == float64(0.0) {
		if infil.Fu <= float64(0.0) {
			return float64(0.0)
		}
		kr = float64(float64(lu/float64(90000.0)) * Evap.recoveryFactor)
		dF = float64(float64(kr*Fumax) * tstep)
		infil.F = float64(infil.F - dF)
		infil.Fu = float64(infil.Fu - dF)
		if infil.Fu <= float64(0.0) {
			infil.Fu = float64(0.0)
			infil.F = float64(0.0)
			infil.IMD = infil.IMDmax
			return float64(0.0)
		}
		if infil.T <= float64(0.0) {
			infil.IMD = float64((float64(Fumax - infil.Fu)) / lu)
			infil.F = float64(0.0)
		}
		return float64(0.0)
	}
	if ia <= ks {
		dF = float64(ia * tstep)
		infil.F = float64(infil.F + dF)
		infil.Fu = float64(infil.Fu + dF)
		if (infil.Fu) <= (Fumax) {
			infil.Fu = (infil.Fu)
		} else {
			infil.Fu = (Fumax)
		}
		if (modelType == GREEN_AMPT) && (infil.T <= float64(0.0)) {
			infil.IMD = float64((float64(Fumax - infil.Fu)) / lu)
			infil.F = float64(0.0)
		}
		return ia
	}
	infil.T = float64(float64(float64(5400.0)/lu) / Evap.recoveryFactor)
	Fs = float64(float64(float64(ks*(float64(infil.S+depth)))*infil.IMD) / (float64(ia - ks)))
	if infil.F > Fs {
		infil.Sat = 1
		return grnampt_getSatInfil(infil, tstep, irate, depth)
	}
	if float64(infil.F+float64(ia*tstep)) < Fs {
		dF = float64(ia * tstep)
		infil.F = float64(infil.F + dF)
		infil.Fu = float64(infil.Fu + dF)
		if (infil.Fu) <= (Fumax) {
			infil.Fu = (infil.Fu)
		} else {
			infil.Fu = (Fumax)
		}
		return ia
	}
	ts = float64(tstep - float64((float64(Fs-infil.F))/ia))
	if ts <= float64(0.0) {
		ts = float64(0.0)
	}
	c1 = float64((float64(infil.S + depth)) * infil.IMD)
	F2 = grnampt_getF2(Fs, c1, ks, ts)
	if F2 > float64(Fs+float64(ia*ts)) {
		F2 = float64(Fs + float64(ia*ts))
	}
	dF = float64(F2 - infil.F)
	infil.F = F2
	infil.Fu = float64(infil.Fu + dF)
	if (infil.Fu) <= (Fumax) {
		infil.Fu = (infil.Fu)
	} else {
		infil.Fu = (Fumax)
	}
	infil.Sat = 1
	return float64(dF / tstep)
}

// infil.c:767
func grnampt_getSatInfil(infil *TGrnAmpt, tstep float64, irate float64, depth float64) float64 {
	var (
		ia float64
		c1 float64
		dF float64
		F2 float64
		ks float64
		lu float64
	)
	_, _, _, _, _, _ = ia, c1, dF, F2, ks, lu
	ia = 0
	c1 = 0
	dF = 0
	F2 = 0
	ks = float64(infil.Ks * InfilFactor)
	lu = float64(infil.Lu * M.Sqrt(InfilFactor))
	ia = float64(irate + float64(depth/tstep))
	if ia < float64(1e-10) {
		return float64(0.0)
	}
	infil.T = float64(float64(float64(5400.0)/lu) / Evap.recoveryFactor)
	c1 = float64((float64(infil.S + depth)) * infil.IMD)
	F2 = grnampt_getF2(infil.F, c1, ks, tstep)
	dF = float64(F2 - infil.F)
	if dF > float64(ia*tstep) {
		dF = float64(ia * tstep)
		infil.Sat = 0
	}
	infil.F = float64(infil.F + dF)
	infil.Fu = float64(infil.Fu + dF)
	if (infil.Fu) <= (Fumax) {
		infil.Fu = (infil.Fu)
	} else {
		infil.Fu = (Fumax)
	}
	return float64(dF / tstep)
}

// infil.c:813
func grnampt_getF2(f1 float64, c1 float64, ks float64, ts float64) float64 {
	var (
		i     int32
		f2    float64
		f2min float64
		df2   float64
		c2    float64
	)
	_, _, _, _, _ = i, f2, f2min, df2, c2
	i = 0
	f2 = f1
	f2min = 0
	df2 = 0
	c2 = 0
	f2min = float64(f1 + float64(ks*ts))
	if c1 == float64(0.0) {
		return f2min
	}
	if (ts < float64(10.0)) && (f1 > float64(float64(0.01)*c1)) {
		f2 = float64(f1 + float64(float64(ks*(float64(float64(1.0)+float64(c1/f1))))*ts))
		if (f2) >= (f2min) {
			return (f2)
		}
		return (f2min)
	}
	c2 = float64(float64(c1*M.Log(float64(f1+c1))) - float64(ks*ts))
	for i = 1; i <= 20; i++ {
		df2 = float64((float64(float64(float64(f2-f1)-float64(c1*M.Log(float64(f2+c1)))) + c2)) / (float64(float64(1.0) - float64(c1/(float64(f2+c1))))))
		if M.Fabs(df2) < float64(0.00001) {
			if (f2) >= (f2min) {
				return (f2)
			}
			return (f2min)
		}
		f2 = float64(f2 - df2)
	}
	return f2min
}

// infil.c:861
func curvenum_setParams(infil *TCurveNum, p []float64) int32 {
	if p[0] < float64(10.0) {
		p[0] = float64(10.0)
	}
	if p[0] > float64(99.0) {
		p[0] = float64(99.0)
	}
	infil.Smax = float64((float64(float64(float64(1000.0)/p[0]) - float64(10.0))) / float64(12.0))
	if infil.Smax < float64(0.0) {
		return 0
	}
	if p[2] > float64(0.0) {
		infil.regen = float64(float64(1.0) / (float64(p[2] * float64(86400))))
	} else {
		return 0
	}
	infil.Tmax = float64(float64(0.06) / infil.regen)
	return 1
}

// infil.c:888
func curvenum_initState(infil *TCurveNum) {
	infil.S = infil.Smax
	infil.P = float64(0.0)
	infil.F = float64(0.0)
	infil.T = float64(0.0)
	infil.Se = infil.Smax
	infil.f = float64(0.0)
}

// infil.c:903
func curvenum_getState(infil *TCurveNum, x []float64) {
	x[0] = infil.S
	x[1] = infil.P
	x[2] = infil.F
	x[3] = infil.T
	x[4] = infil.Se
	x[5] = infil.f
}

// infil.c:913
func curvenum_setState(infil *TCurveNum, x []float64) {
	infil.S = x[0]
	infil.P = x[1]
	infil.F = x[2]
	infil.T = x[3]
	infil.Se = x[4]
	infil.f = x[5]
}

// infil.c:925
func curvenum_getInfil(infil *TCurveNum, tstep float64, irate float64, depth float64) float64 {
	var (
		F1 float64
		f1 float64
		fa float64
	)
	_, _, _ = F1, f1, fa
	F1 = 0
	f1 = float64(0.0)
	fa = float64(irate + float64(depth/tstep))
	if irate > float64(1e-10) {
		if infil.T >= infil.Tmax {
			infil.P = float64(0.0)
			infil.F = float64(0.0)
			infil.f = float64(0.0)
			infil.Se = infil.S
		}
		infil.T = float64(0.0)
		infil.P = float64(infil.P + (float64(irate * tstep)))
		F1 = float64(infil.P * (float64(float64(1.0) - float64(infil.P/(float64(infil.P+infil.Se))))))
		f1 = float64((float64(F1 - infil.F)) / tstep)
		if (f1 < float64(0.0)) || (infil.S <= float64(0.0)) {
			f1 = float64(0.0)
		}
	} else {
		if (depth > float64(0.004167)) && (infil.S > float64(0.0)) {
			f1 = infil.f
			if float64(f1*tstep) > infil.S {
				f1 = float64(infil.S / tstep)
			}
		} else {
			infil.T = float64(infil.T + tstep)
		}
	}
	if f1 > float64(0.0) {
		if (f1) <= (fa) {
			f1 = (f1)
		} else {
			f1 = (fa)
		}
		if (f1) >= (float64(0.0)) {
			f1 = (f1)
		} else {
			f1 = (float64(0.0))
		}
		infil.F = float64(infil.F + (float64(f1 * tstep)))
		if infil.regen > float64(0.0) {
			infil.S = float64(infil.S - (float64(f1 * tstep)))
			if infil.S < float64(0.0) {
				infil.S = float64(0.0)
			}
		}
	} else {
		infil.S = float64(infil.S + (float64(float64(float64(infil.regen*infil.Smax)*tstep) * Evap.recoveryFactor)))
		if infil.S > infil.Smax {
			infil.S = infil.Smax
		}
	}
	infil.f = f1
	return f1
}
