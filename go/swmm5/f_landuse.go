// Generated from EPA SWMM 5.2.4 (landuse.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

// landuse.c:51
func landuse_readParams(j int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		id R.CPtr
	)
	_ = id
	id = (R.CPtr{})
	if ntoks < 1 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(LANDUSE, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	Landuse[j].ID = id
	if ntoks > 1 {
		if ntoks < 4 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		if !(getDouble(tok[1], &Landuse[j].sweepInterval) != 0) {
			return error_setInpError(ERR_NUMBER, tok[1])
		}
		if !(getDouble(tok[2], &Landuse[j].sweepRemoval) != 0) {
			return error_setInpError(ERR_NUMBER, tok[2])
		}
		if !(getDouble(tok[3], &Landuse[j].sweepDays0) != 0) {
			return error_setInpError(ERR_NUMBER, tok[3])
		}
	} else {
		Landuse[j].sweepInterval = float64(0.0)
		Landuse[j].sweepRemoval = float64(0.0)
		Landuse[j].sweepDays0 = float64(0.0)
	}
	if (Landuse[j].sweepRemoval < float64(0.0)) || (Landuse[j].sweepRemoval > float64(1.0)) {
		return error_setInpError(ERR_NUMBER, tok[2])
	}
	return 0
}

// landuse.c:92
func landuse_readPollutParams(j int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		i        int32
		k        int32
		coPollut int32
		snowFlag int32
		x        [4]float64
		coFrac   float64
		cDWF     float64
		cInit    float64
		id       R.CPtr
	)
	_, _, _, _, _, _, _, _, _ = i, k, coPollut, snowFlag, x, coFrac, cDWF, cInit, id
	i = 0
	k = 0
	coPollut = 0
	snowFlag = 0
	x = ([4]float64{})
	coFrac = 0
	cDWF = 0
	cInit = 0
	id = (R.CPtr{})
	if ntoks < 6 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(POLLUT, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	k = findmatch(tok[1], QualUnitsWords[:])
	if k < 0 {
		return error_setInpError(ERR_KEYWORD, tok[1])
	}
	for i = 2; i <= 4; i++ {
		if (!(getDouble(tok[i], &x[(i-2)]) != 0)) || (x[(i-2)] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[i])
		}
	}
	if !(getDouble(tok[5], &x[3]) != 0) {
		return error_setInpError(ERR_NUMBER, tok[5])
	}
	snowFlag = 0
	coPollut = (-1)
	coFrac = float64(0.0)
	cDWF = float64(0.0)
	cInit = float64(0.0)
	if ntoks >= 7 {
		snowFlag = findmatch(tok[6], NoYesWords[:])
		if snowFlag < 0 {
			return error_setInpError(ERR_KEYWORD, tok[6])
		}
	}
	if ntoks >= 9 {
		if !(strcomp(tok[7], L275) != 0) {
			coPollut = project_findObject(POLLUT, tok[7])
			if coPollut < 0 {
				return error_setInpError(ERR_NAME, tok[7])
			}
			if (!(getDouble(tok[8], &coFrac) != 0)) || (coFrac < float64(0.0)) {
				return error_setInpError(ERR_NUMBER, tok[8])
			}
		}
	}
	if ntoks >= 10 {
		if (!(getDouble(tok[9], &cDWF) != 0)) || (cDWF < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[9])
		}
	}
	if ntoks >= 11 {
		if (!(getDouble(tok[10], &cInit) != 0)) || (cInit < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[9])
		}
	}
	Pollut[j].ID = id
	Pollut[j].units = k
	if Pollut[j].units == MG {
		Pollut[j].mcf = UCF(MASS)
	} else if Pollut[j].units == UG {
		Pollut[j].mcf = float64(UCF(MASS) / float64(1000.0))
	} else {
		Pollut[j].mcf = float64(1.0)
	}
	Pollut[j].pptConcen = x[0]
	Pollut[j].gwConcen = x[1]
	Pollut[j].rdiiConcen = x[2]
	Pollut[j].kDecay = float64(x[3] / float64(86400))
	Pollut[j].snowOnly = snowFlag
	Pollut[j].coPollut = coPollut
	Pollut[j].coFraction = coFrac
	Pollut[j].dwfConcen = cDWF
	Pollut[j].initConcen = cInit
	return 0
}

// landuse.c:190
func landuse_readBuildupParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		i    int32
		j    int32
		k    int32
		n_2  int32
		p    int32
		c    [3]float64
		tmax float64
	)
	_, _, _, _, _, _, _ = i, j, k, n_2, p, c, tmax
	i = 0
	j = 0
	k = 0
	n_2 = 0
	p = 0
	c = [3]float64{float64(0), float64(0), float64(0)}
	tmax = 0
	if ntoks < 3 {
		return 0
	}
	j = project_findObject(LANDUSE, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	p = project_findObject(POLLUT, tok[1])
	if p < 0 {
		return error_setInpError(ERR_NAME, tok[1])
	}
	k = findmatch(tok[2], BuildupTypeWords[:])
	if k < 0 {
		return error_setInpError(ERR_KEYWORD, tok[2])
	}
	Landuse[j].buildupFunc[p].funcType = k
	if k > NO_BUILDUP {
		if ntoks < 7 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		if k != EXTERNAL_BUILDUP {
			for i = 0; i < 3; i++ {
				if (!(getDouble(tok[(i+3)], &c[i]) != 0)) || (c[i] < float64(0.0)) {
					return error_setInpError(ERR_NUMBER, tok[(i+3)])
				}
			}
		}
		n_2 = findmatch(tok[6], NormalizerWords[:])
		if n_2 < 0 {
			return error_setInpError(ERR_KEYWORD, tok[6])
		}
		Landuse[j].buildupFunc[p].normalizer = n_2
	}
	switch Landuse[j].buildupFunc[p].funcType {
	case POWER_BUILDUP:
		if (c[2] > float64(0.0)) && ((c[2] < float64(0.01)) || (c[2] > float64(10.0))) {
			return error_setInpError(ERR_KEYWORD, tok[5])
		}
		if float64(c[1]*c[2]) == float64(0.0) {
			tmax = float64(0.0)
		} else if float64(M.Log10(c[0])/c[2]) > float64(3.5) {
			tmax = float64(3650.0)
		} else {
			tmax = M.Pow(float64(c[0]/c[1]), float64(float64(1.0)/c[2]))
		}
		break
	case EXPON_BUILDUP:
		if c[1] == float64(0.0) {
			tmax = float64(0.0)
		} else {
			tmax = float64((-(M.Log(float64(0.001)))) / c[1])
		}
		break
	case SATUR_BUILDUP:
		tmax = float64(float64(1000.0) * c[2])
		break
	case EXTERNAL_BUILDUP:
		if (!(getDouble(tok[3], &c[0]) != 0)) || (c[0] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[3])
		}
		if (!(getDouble(tok[4], &c[1]) != 0)) || (c[1] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[3])
		}
		n_2 = project_findObject(TSERIES, tok[5])
		if n_2 < 0 {
			return error_setInpError(ERR_NAME, tok[4])
		}
		Tseries[n_2].refersTo = EXTERNAL_BUILDUP
		c[2] = float64(n_2)
		tmax = float64(0.0)
		break
	default:
		tmax = float64(0.0)
	}
	Landuse[j].buildupFunc[p].coeff[0] = c[0]
	Landuse[j].buildupFunc[p].coeff[1] = c[1]
	Landuse[j].buildupFunc[p].coeff[2] = c[2]
	Landuse[j].buildupFunc[p].maxDays = tmax
	return 0
}

// landuse.c:281
func landuse_readWashoffParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		i      int32
		j      int32
		p      int32
		func_2 int32
		x      [4]float64
	)
	_, _, _, _, _ = i, j, p, func_2, x
	i = 0
	j = 0
	p = 0
	func_2 = 0
	x = ([4]float64{})
	if ntoks < 3 {
		return 0
	}
	for i = 0; i < 4; i++ {
		x[i] = float64(0.0)
	}
	func_2 = NO_WASHOFF
	j = project_findObject(LANDUSE, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	p = project_findObject(POLLUT, tok[1])
	if p < 0 {
		return error_setInpError(ERR_NAME, tok[1])
	}
	if ntoks > 2 {
		func_2 = findmatch(tok[2], WashoffTypeWords[:])
		if func_2 < 0 {
			return error_setInpError(ERR_KEYWORD, tok[2])
		}
		if func_2 != NO_WASHOFF {
			if ntoks < 5 {
				return error_setInpError(ERR_ITEMS, L0)
			}
			if !(getDouble(tok[3], &x[0]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[3])
			}
			if !(getDouble(tok[4], &x[1]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[4])
			}
			if ntoks >= 6 {
				if !(getDouble(tok[5], &x[2]) != 0) {
					return error_setInpError(ERR_NUMBER, tok[5])
				}
			}
			if ntoks >= 7 {
				if !(getDouble(tok[6], &x[3]) != 0) {
					return error_setInpError(ERR_NUMBER, tok[6])
				}
			}
		}
	}
	if x[0] < float64(0.0) {
		return error_setInpError(ERR_NUMBER, tok[3])
	}
	if (x[1] < (-(float64(10.0)))) || (x[1] > float64(10.0)) {
		return error_setInpError(ERR_NUMBER, tok[4])
	}
	if (x[2] < float64(0.0)) || (x[2] > float64(100.0)) {
		return error_setInpError(ERR_NUMBER, tok[5])
	}
	if (x[3] < float64(0.0)) || (x[3] > float64(100.0)) {
		return error_setInpError(ERR_NUMBER, tok[6])
	}
	if func_2 == EXPON_WASHOFF {
		x[0] = float64(x[0] / (float64(3600.0)))
	}
	if func_2 == RATING_WASHOFF {
		x[0] = float64(x[0] * (M.Pow(UCF(FLOW), x[1])))
	}
	if func_2 == EMC_WASHOFF {
		x[0] = float64(x[0] * (float64(28.317)))
	}
	Landuse[j].washoffFunc[p].funcType = func_2
	Landuse[j].washoffFunc[p].coeff = x[0]
	Landuse[j].washoffFunc[p].expon = x[1]
	Landuse[j].washoffFunc[p].sweepEffic = float64(x[2] / float64(100.0))
	Landuse[j].washoffFunc[p].bmpEffic = float64(x[3] / float64(100.0))
	return 0
}

// landuse.c:355
func landuse_getInitBuildup(landFactor []TLandFactor, initBuildup []float64, area float64, curb float64) {
	var (
		i               int32
		p               int32
		startDrySeconds float64
		f               float64
		fArea           float64
		fCurb           float64
		buildup         float64
	)
	_, _, _, _, _, _, _ = i, p, startDrySeconds, f, fArea, fCurb, buildup
	i = 0
	p = 0
	startDrySeconds = 0
	f = 0
	fArea = 0
	fCurb = 0
	buildup = 0
	startDrySeconds = float64(StartDryDays * float64(86400))
	for i = 0; i < Nobjects[LANDUSE]; i++ {
		landFactor[i].lastSwept = float64(StartDateTime - Landuse[i].sweepDays0)
		f = landFactor[i].fraction
		fArea = float64(float64(f*area) * UCF(LANDAREA))
		fCurb = float64(f * curb)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			buildup = float64(0.0)
			if initBuildup[p] > float64(0.0) {
				buildup = float64(initBuildup[p] * fArea)
			} else {
				buildup = landuse_getBuildup(i, p, fArea, fCurb, buildup, startDrySeconds)
			}
			landFactor[i].buildup[p] = buildup
		}
	}
}

// landuse.c:411
func landuse_getBuildup(i int32, p int32, area float64, curb float64, buildup float64, tStep float64) float64 {
	var (
		n_2     int32
		days    float64
		perUnit float64
	)
	_, _, _ = n_2, days, perUnit
	n_2 = 0
	days = 0
	perUnit = 0
	if (Landuse[i].buildupFunc[p].funcType == NO_BUILDUP) || (tStep == float64(0.0)) {
		return buildup
	}
	n_2 = Landuse[i].buildupFunc[p].normalizer
	perUnit = float64(1.0)
	if n_2 == PER_AREA {
		perUnit = area
	}
	if n_2 == PER_CURB {
		perUnit = curb
	}
	if perUnit == float64(0.0) {
		return float64(0.0)
	}
	if Landuse[i].buildupFunc[p].funcType == EXTERNAL_BUILDUP {
		return float64(landuse_getExternalBuildup(i, p, float64(buildup/perUnit), tStep) * perUnit)
	}
	days = landuse_getBuildupDays(i, p, float64(buildup/perUnit))
	days = float64(days + (float64(tStep / float64(86400))))
	return float64(landuse_getBuildupMass(i, p, days) * perUnit)
}

// landuse.c:458
func landuse_getBuildupDays(i int32, p int32, buildup float64) float64 {
	var (
		c0 float64
		c1 float64
		c2 float64
	)
	_, _, _ = c0, c1, c2
	c0 = Landuse[i].buildupFunc[p].coeff[0]
	c1 = Landuse[i].buildupFunc[p].coeff[1]
	c2 = Landuse[i].buildupFunc[p].coeff[2]
	if buildup == float64(0.0) {
		return float64(0.0)
	}
	if buildup >= c0 {
		return Landuse[i].buildupFunc[p].maxDays
	}
	switch Landuse[i].buildupFunc[p].funcType {
	case POWER_BUILDUP:
		if float64(c1*c2) == float64(0.0) {
			return float64(0.0)
		} else {
			return M.Pow((float64(buildup / c1)), (float64(float64(1.0) / c2)))
		}
	case EXPON_BUILDUP:
		if float64(c0*c1) == float64(0.0) {
			return float64(0.0)
		} else {
			return float64((-(M.Log(float64(float64(1.0) - float64(buildup/c0))))) / c1)
		}
	case SATUR_BUILDUP:
		if c0 == float64(0.0) {
			return float64(0.0)
		} else {
			return float64(float64(buildup*c2) / (float64(c0 - buildup)))
		}
	default:
		return float64(0.0)
	}
	return 0
}

// landuse.c:494
func landuse_getBuildupMass(i int32, p int32, days float64) float64 {
	var (
		b  float64
		c0 float64
		c1 float64
		c2 float64
	)
	_, _, _, _ = b, c0, c1, c2
	b = 0
	c0 = Landuse[i].buildupFunc[p].coeff[0]
	c1 = Landuse[i].buildupFunc[p].coeff[1]
	c2 = Landuse[i].buildupFunc[p].coeff[2]
	if days == float64(0.0) {
		return float64(0.0)
	}
	if days >= Landuse[i].buildupFunc[p].maxDays {
		return c0
	}
	switch Landuse[i].buildupFunc[p].funcType {
	case POWER_BUILDUP:
		b = float64(c1 * M.Pow(days, c2))
		if b > c0 {
			b = c0
		}
		break
	case EXPON_BUILDUP:
		b = float64(c0 * (float64(float64(1.0) - M.Exp(float64((-days)*c1)))))
		break
	case SATUR_BUILDUP:
		b = float64(float64(days*c0) / (float64(c2 + days)))
		break
	default:
		b = float64(0.0)
	}
	return b
}

// landuse.c:532
func landuse_getAvgBmpEffic(j int32, p int32) float64 {
	var (
		i int32
		r float64
	)
	_, _ = i, r
	i = 0
	r = float64(0.0)
	for i = 0; i < Nobjects[LANDUSE]; i++ {
		r = float64(r + (float64(Subcatch[j].landFactor[i].fraction * Landuse[i].washoffFunc[p].bmpEffic)))
	}
	return r
}

// landuse.c:553
func landuse_getWashoffLoad(i int32, p int32, area float64, landFactor []TLandFactor, runoff float64, vOutflow float64) float64 {
	var (
		landuseArea float64
		buildup     float64
		washoffQual float64
		washoffLoad float64
		bmpRemoval  float64
	)
	_, _, _, _, _ = landuseArea, buildup, washoffQual, washoffLoad, bmpRemoval
	landuseArea = 0
	buildup = 0
	washoffQual = 0
	washoffLoad = 0
	bmpRemoval = 0
	buildup = landFactor[i].buildup[p]
	landuseArea = float64(landFactor[i].fraction * area)
	washoffQual = landuse_getWashoffQual(i, p, buildup, runoff, landuseArea)
	washoffLoad = float64(float64(float64(float64(washoffQual*vOutflow)*landuseArea)/area) * Pollut[p].mcf)
	if (Landuse[i].buildupFunc[p].funcType != NO_BUILDUP) || (buildup > washoffLoad) {
		if (washoffLoad) <= (buildup) {
			washoffLoad = (washoffLoad)
		} else {
			washoffLoad = (buildup)
		}
		buildup = float64(buildup - washoffLoad)
		landFactor[i].buildup[p] = buildup
	} else {
		massbal_updateLoadingTotals(BUILDUP_LOAD, p, washoffLoad)
		landFactor[i].buildup[p] = float64(0.0)
	}
	bmpRemoval = float64(Landuse[i].washoffFunc[p].bmpEffic * washoffLoad)
	if bmpRemoval > float64(0.0) {
		massbal_updateLoadingTotals(BMP_REMOVAL_LOAD, p, bmpRemoval)
		washoffLoad = float64(washoffLoad - bmpRemoval)
	}
	return float64(washoffLoad / Pollut[p].mcf)
}

// landuse.c:612
func landuse_getWashoffQual(i int32, p int32, buildup float64, runoff float64, area float64) float64 {
	var (
		cWashoff float64
		coeff    float64
		expon    float64
		func_2   int32
	)
	_, _, _, _ = cWashoff, coeff, expon, func_2
	cWashoff = float64(0.0)
	coeff = Landuse[i].washoffFunc[p].coeff
	expon = Landuse[i].washoffFunc[p].expon
	func_2 = Landuse[i].washoffFunc[p].funcType
	if (func_2 == NO_WASHOFF) || (runoff == float64(0.0)) {
		return float64(0.0)
	}
	if (Landuse[i].buildupFunc[p].funcType != NO_BUILDUP) && (buildup == float64(0.0)) {
		return float64(0.0)
	}
	if func_2 == EXPON_WASHOFF {
		cWashoff = float64(float64(float64(coeff*M.Pow(float64(runoff*UCF(RAINFALL)), expon))*buildup) / Pollut[p].mcf)
		cWashoff = float64(cWashoff / (float64(runoff * area)))
	} else if func_2 == RATING_WASHOFF {
		cWashoff = float64(coeff * M.Pow(float64(runoff*area), float64(expon-float64(1.0))))
	} else if func_2 == EMC_WASHOFF {
		cWashoff = coeff
	}
	return cWashoff
}

// landuse.c:665
func landuse_getCoPollutLoad(p int32, washoff []float64) float64 {
	var (
		k int32
		w float64
	)
	_, _ = k, w
	k = 0
	w = 0
	k = Pollut[p].coPollut
	if k >= 0 {
		w = float64(Pollut[p].coFraction * washoff[k])
		massbal_updateLoadingTotals(BUILDUP_LOAD, p, float64(w*Pollut[p].mcf))
		return w
	}
	return float64(0.0)
}

// landuse.c:693
func landuse_getExternalBuildup(i int32, p int32, buildup float64, tStep float64) float64 {
	var (
		maxBuildup float64
		sf         float64
		ts         int32
		rate       float64
	)
	_, _, _, _ = maxBuildup, sf, ts, rate
	maxBuildup = Landuse[i].buildupFunc[p].coeff[0]
	sf = Landuse[i].buildupFunc[p].coeff[1]
	ts = R.D2I(M.Floor(Landuse[i].buildupFunc[p].coeff[2]))
	rate = float64(0.0)
	if NewRunoffTime == float64(0.0) {
		return float64(0.0)
	}
	if ts >= 0 {
		rate = float64(sf * table_tseriesLookup(&Tseries[ts], getDateTime(NewRunoffTime), 0))
	}
	buildup = float64(buildup + float64(float64(rate*tStep)/float64(86400)))
	if (buildup) <= (maxBuildup) {
		buildup = (buildup)
	} else {
		buildup = (maxBuildup)
	}
	return buildup
}
