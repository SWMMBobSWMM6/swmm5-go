// Generated from EPA SWMM 5.2.4 (snow.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

// snow.c:68
func snow_readMeltParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		i   int32
		j   int32
		k   int32
		m   int32
		n_2 int32
		x   [7]float64
	)
	_, _, _, _, _, _ = i, j, k, m, n_2, x
	i = 0
	j = 0
	k = 0
	m = 0
	n_2 = 0
	x = ([7]float64{})
	if ntoks < 8 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(SNOWMELT, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	if Snowmelt[j].ID.IsNull() {
		Snowmelt[j].ID = project_findID(SNOWMELT, tok[0])
	}
	k = findmatch(tok[1], SnowmeltWords[:])
	if k < 0 {
		return error_setInpError(ERR_KEYWORD, tok[1])
	}
	n_2 = 7
	if k == SNOW_REMOVAL {
		n_2 = 6
	}
	if ntoks < (n_2 + 2) {
		return error_setInpError(ERR_ITEMS, L0)
	}
	for i = 0; i < 7; i++ {
		x[i] = float64(0.0)
	}
	for i = 0; i < n_2; i++ {
		if !(getDouble(tok[(i+2)], &x[i]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[(i+2)])
		}
	}
	if k == SNOW_REMOVAL {
		x[6] = (-(float64(1.0)))
		if ntoks >= 9 {
			m = project_findObject(SUBCATCH, tok[8])
			if m < 0 {
				return error_setInpError(ERR_NAME, tok[8])
			}
			x[6] = float64(m)
		}
	}
	setMeltParams(j, k, x[:])
	return 0
}

// snow.c:126
func snow_createSnowpack(j int32, k int32) int32 {
	var (
		snowpack *TSnowpack
	)
	_ = snowpack
	snowpack = nil
	snowpack = new(TSnowpack)
	if !(snowpack != nil) {
		return 0
	}
	Subcatch[j].snowpack = snowpack
	snowpack.snowmeltIndex = k
	return 1
}

// snow.c:144
func snow_initSnowpack(j int32) {
	var (
		i         int32
		k         int32
		f         float64
		snowDepth float64
		snowpack  *TSnowpack
	)
	_, _, _, _, _ = i, k, f, snowDepth, snowpack
	i = 0
	k = 0
	f = 0
	snowDepth = float64(0.0)
	snowpack = nil
	snowpack = Subcatch[j].snowpack
	if snowpack == nil {
		return
	}
	k = Subcatch[j].snowpack.snowmeltIndex
	f = Snowmelt[k].snn
	snowpack.fArea[SNOW_PLOWABLE] = float64(f * Subcatch[j].fracImperv)
	snowpack.fArea[SNOW_IMPERV] = float64((float64(float64(1.0) - f)) * Subcatch[j].fracImperv)
	snowpack.fArea[SNOW_PERV] = float64(float64(1.0) - Subcatch[j].fracImperv)
	for i = SNOW_PLOWABLE; i <= SNOW_PERV; i++ {
		if snowpack.fArea[i] > float64(0.0) {
			snowpack.wsnow[i] = Snowmelt[k].wsnow[i]
			snowpack.fw[i] = Snowmelt[k].fwnow[i]
		} else {
			snowpack.wsnow[i] = float64(0.0)
			snowpack.fw[i] = float64(0.0)
		}
		snowpack.coldc[i] = float64(0.0)
		snowpack.ati[i] = Snowmelt[k].tbase[i]
		snowpack.awe[i] = float64(1.0)
		snowDepth = float64(snowDepth + (float64(snowpack.wsnow[i] * snowpack.fArea[i])))
	}
	Subcatch[j].newSnowDepth = snowDepth
}

// snow.c:193
func snow_initSnowmelt(j int32) {
	var (
		i int32
		k int32
	)
	_, _ = i, k
	i = 0
	k = 0
	for i = 0; i < 3; i++ {
		Snowmelt[j].snn = float64(0.0)
		Snowmelt[j].si[i] = float64(0.0)
		Snowmelt[j].dhmin[i] = float64(0.0)
		Snowmelt[j].dhmax[i] = float64(0.0)
		Snowmelt[j].tbase[i] = float64(0.0)
		Snowmelt[j].fwfrac[i] = float64(0.0)
		Snowmelt[j].wsnow[i] = float64(0.0)
		Snowmelt[j].fwnow[i] = float64(0.0)
		Snowmelt[j].weplow = float64(1.0e6)
		for k = 0; k < 5; k++ {
			Snowmelt[j].sfrac[k] = float64(0.0)
		}
		Snowmelt[j].toSubcatch = (-1)
	}
}

// snow.c:219
func snow_validateSnowmelt(j int32) {
	var (
		k   int32
		err int32
		sum float64
	)
	_, _, _ = k, err, sum
	k = 0
	err = 0
	sum = float64(0.0)
	for k = SNOW_PLOWABLE; k <= SNOW_PERV; k++ {
		if Snowmelt[j].dhmin[k] > Snowmelt[j].dhmax[k] {
			err = 1
		}
		if (Snowmelt[j].fwfrac[k] < float64(0.0)) || (Snowmelt[j].fwfrac[k] > float64(1.0)) {
			err = 1
		}
	}
	if (Snowmelt[j].snn < float64(0.0)) || (Snowmelt[j].snn > float64(1.0)) {
		err = 1
	}
	for k = 0; k < 5; k++ {
		sum = float64(sum + Snowmelt[j].sfrac[k])
	}
	if sum > float64(1.01) {
		err = 1
	}
	if err != 0 {
		report_writeErrorMsg(ERR_SNOWPACK_PARAMS, Snowmelt[j].ID)
	}
}

// snow.c:251
func snow_getState(i int32, j int32, x []float64) {
	var (
		snowpack *TSnowpack
	)
	_ = snowpack
	snowpack = Subcatch[i].snowpack
	if snowpack == nil {
		return
	}
	x[0] = snowpack.wsnow[j]
	x[1] = snowpack.fw[j]
	x[2] = snowpack.coldc[j]
	x[3] = snowpack.ati[j]
	x[4] = snowpack.awe[j]
}

// snow.c:270
func snow_setState(i int32, j int32, x []float64) {
	var (
		snowpack *TSnowpack
	)
	_ = snowpack
	snowpack = Subcatch[i].snowpack
	if snowpack == nil {
		return
	}
	snowpack.wsnow[j] = x[0]
	snowpack.fw[j] = x[1]
	snowpack.coldc[j] = x[2]
	snowpack.ati[j] = x[3]
	snowpack.awe[j] = x[4]
}

// snow.c:290
func setMeltParams(j int32, k int32, x []float64) {
	var (
		i int32
	)
	_ = i
	i = 0
	if (k >= SNOW_PLOWABLE) && (k <= SNOW_PERV) {
		Snowmelt[j].dhmin[k] = float64(float64(x[0]*UCF(TEMPERATURE)) / UCF(RAINFALL))
		Snowmelt[j].dhmax[k] = float64(float64(x[1]*UCF(TEMPERATURE)) / UCF(RAINFALL))
		Snowmelt[j].tbase[k] = x[2]
		if UnitSystem == SI {
			Snowmelt[j].tbase[k] = float64(float64((float64(float64(9.0)/float64(5.0)))*Snowmelt[j].tbase[k]) + float64(32.0))
		}
		Snowmelt[j].fwfrac[k] = x[3]
		Snowmelt[j].wsnow[k] = float64(x[4] / UCF(RAINDEPTH))
		if (x[5]) <= (float64(x[3] * x[4])) {
			x[5] = (x[5])
		} else {
			x[5] = (float64(x[3] * x[4]))
		}
		Snowmelt[j].fwnow[k] = float64(x[5] / UCF(RAINDEPTH))
		if k == SNOW_PLOWABLE {
			Snowmelt[j].snn = x[6]
		} else {
			Snowmelt[j].si[k] = float64(x[6] / UCF(RAINDEPTH))
		}
	} else if k == SNOW_REMOVAL {
		Snowmelt[j].weplow = float64(x[0] / UCF(RAINDEPTH))
		for i = 0; i <= 4; i++ {
			Snowmelt[j].sfrac[i] = x[(i + 1)]
		}
		if x[6] >= float64(0.0) {
			Snowmelt[j].toSubcatch = R.D2I((float64(x[6] + float64(0.01))))
		} else {
			Snowmelt[j].toSubcatch = (-1)
		}
	}
}

// snow.c:341
func snow_setMeltCoeffs(j int32, s float64) {
	var (
		k int32
	)
	_ = k
	k = 0
	for k = SNOW_PLOWABLE; k <= SNOW_PERV; k++ {
		Snowmelt[j].dhm[k] = float64(float64(0.5) * (float64(float64(Snowmelt[j].dhmax[k]*(float64(float64(1.0)+s))) + float64(Snowmelt[j].dhmin[k]*(float64(float64(1.0)-s))))))
	}
}

// snow.c:360
func snow_plowSnow(j int32, tStep float64) {
	var (
		i          int32
		k          int32
		m          int32
		rainfall   float64
		snowfall   float64
		exc        float64
		f          float64
		sfracTotal float64
		snowpack   *TSnowpack
	)
	_, _, _, _, _, _, _, _, _ = i, k, m, rainfall, snowfall, exc, f, sfracTotal, snowpack
	i = 0
	k = 0
	m = 0
	rainfall = 0
	snowfall = 0
	exc = 0
	f = 0
	sfracTotal = 0
	snowpack = nil
	snowpack = Subcatch[j].snowpack
	if !(snowpack != nil) {
		return
	}
	gage_getPrecip(Subcatch[j].gage, &rainfall, &snowfall)
	for i = SNOW_PLOWABLE; i <= SNOW_PERV; i++ {
		if snowpack.fArea[i] > float64(0.0) {
			snowpack.wsnow[i] = float64(snowpack.wsnow[i] + (float64(snowfall * tStep)))
			snowpack.imelt[i] = float64(0.0)
		}
	}
	if snowpack.fArea[SNOW_PLOWABLE] > float64(0.0) {
		k = snowpack.snowmeltIndex
		if snowpack.wsnow[SNOW_PLOWABLE] >= Snowmelt[k].weplow {
			exc = snowpack.wsnow[SNOW_PLOWABLE]
			f = float64(snowpack.fArea[SNOW_PLOWABLE] * (float64(Subcatch[j].area - Subcatch[j].lidArea)))
			Snow.removed = float64(Snow.removed + (float64(float64(Snowmelt[k].sfrac[0]*exc) * f)))
			sfracTotal = Snowmelt[k].sfrac[0]
			if snowpack.fArea[SNOW_IMPERV] > float64(0.0) {
				f = float64(snowpack.fArea[SNOW_PLOWABLE] / snowpack.fArea[SNOW_IMPERV])
				snowpack.wsnow[SNOW_IMPERV] = float64(snowpack.wsnow[SNOW_IMPERV] + (float64(float64(Snowmelt[k].sfrac[1]*exc) * f)))
				sfracTotal = float64(sfracTotal + Snowmelt[k].sfrac[1])
			}
			if snowpack.fArea[SNOW_PERV] > float64(0.0) {
				f = float64(snowpack.fArea[SNOW_PLOWABLE] / snowpack.fArea[SNOW_PERV])
				snowpack.wsnow[SNOW_PERV] = float64(snowpack.wsnow[SNOW_PERV] + (float64(float64(Snowmelt[k].sfrac[2]*exc) * f)))
				sfracTotal = float64(sfracTotal + Snowmelt[k].sfrac[2])
			}
			snowpack.imelt[SNOW_PLOWABLE] = float64(float64(Snowmelt[k].sfrac[3]*exc) / tStep)
			sfracTotal = float64(sfracTotal + Snowmelt[k].sfrac[3])
			if Snowmelt[k].sfrac[4] > float64(0.0) {
				m = Snowmelt[k].toSubcatch
				if Subcatch[m].snowpack != nil {
					f = Subcatch[m].snowpack.fArea[SNOW_PERV]
				} else {
					f = float64(0.0)
				}
				if f > float64(0.0) {
					f = float64(snowpack.fArea[SNOW_PLOWABLE] / f)
					Subcatch[m].snowpack.wsnow[SNOW_PERV] = float64(Subcatch[m].snowpack.wsnow[SNOW_PERV] + (float64(float64(Snowmelt[k].sfrac[4]*exc) * f)))
					sfracTotal = float64(sfracTotal + Snowmelt[k].sfrac[4])
				}
			}
			if (sfracTotal) <= (float64(1.0)) {
				sfracTotal = (sfracTotal)
			} else {
				sfracTotal = (float64(1.0))
			}
			snowpack.wsnow[SNOW_PLOWABLE] = float64(exc * (float64(float64(1.0) - sfracTotal)))
		}
	}
}

// snow.c:458
func snow_getSnowMelt(j int32, rainfall float64, snowfall float64, tStep float64, netPrecip []float64) float64 {
	var (
		i            int32
		rmelt        float64
		smelt        float64
		asc          float64
		snowDepth    float64
		impervPrecip float64
		snowpack     *TSnowpack
	)
	_, _, _, _, _, _, _ = i, rmelt, smelt, asc, snowDepth, impervPrecip, snowpack
	i = 0
	rmelt = 0
	smelt = 0
	asc = 0
	snowDepth = float64(0.0)
	impervPrecip = 0
	snowpack = nil
	snowpack = Subcatch[j].snowpack
	rmelt = getRainmelt(rainfall)
	for i = SNOW_PLOWABLE; i <= SNOW_PERV; i++ {
		if snowpack.wsnow[i] <= float64(float64(0.001)/float64(12.0)) {
			asc = float64(0.0)
			smelt = float64(0.0)
			snowpack.imelt[i] = float64(snowpack.imelt[i] + (float64((float64(snowpack.wsnow[i] + snowpack.fw[i])) / tStep)))
			snowpack.wsnow[i] = float64(0.0)
			snowpack.fw[i] = float64(0.0)
			snowpack.coldc[i] = float64(0.0)
		} else {
			asc = getArealDepletion(snowpack, i, snowfall, tStep)
			smelt = meltSnowpack(snowpack, i, rmelt, asc, snowfall, tStep)
			smelt = routeSnowmelt(snowpack, i, smelt, asc, rainfall, tStep)
		}
		netPrecip[i] = float64(float64(smelt+snowpack.imelt[i]) + float64(rainfall*(float64(float64(1.0)-asc))))
		snowDepth = float64(snowDepth + (float64(snowpack.wsnow[i] * snowpack.fArea[i])))
	}
	if Subcatch[j].fracImperv > float64(0.0) {
		impervPrecip = float64((float64(float64(netPrecip[SNOW_PLOWABLE]*snowpack.fArea[SNOW_PLOWABLE]) + float64(netPrecip[SNOW_IMPERV]*snowpack.fArea[SNOW_IMPERV]))) / Subcatch[j].fracImperv)
		netPrecip[IMPERV0] = impervPrecip
		netPrecip[IMPERV1] = impervPrecip
	}
	return snowDepth
}

// snow.c:532
func snow_getSnowCover(j int32) float64 {
	var (
		i         int32
		snowCover float64
		snowpack  *TSnowpack
	)
	_, _, _ = i, snowCover, snowpack
	i = 0
	snowCover = float64(0.0)
	snowpack = nil
	snowpack = Subcatch[j].snowpack
	if !(snowpack != nil) {
		return float64(0.0)
	}
	for i = SNOW_PLOWABLE; i <= SNOW_PERV; i++ {
		snowCover = float64(snowCover + (float64((float64(snowpack.wsnow[i] + snowpack.fw[i])) * snowpack.fArea[i])))
	}
	return float64(snowCover * (float64(Subcatch[j].area - Subcatch[j].lidArea)))
}

// snow.c:555
func getArealDepletion(snowpack *TSnowpack, i int32, snowfall float64, tStep float64) float64 {
	var (
		k     int32
		asc   float64
		si    float64
		awesi float64
		awe   float64
		sba   float64
		sbws  float64
	)
	_, _, _, _, _, _, _ = k, asc, si, awesi, awe, sba, sbws
	k = 0
	asc = 0
	si = 0
	awesi = 0
	awe = 0
	sba = 0
	sbws = 0
	if i == SNOW_PLOWABLE {
		return float64(1.0)
	}
	k = snowpack.snowmeltIndex
	si = Snowmelt[k].si[i]
	if (si == float64(0.0)) || (snowpack.wsnow[i] >= si) {
		snowpack.awe[i] = float64(1.0)
		return float64(1.0)
	}
	if snowpack.wsnow[i] == float64(0.0) {
		snowpack.awe[i] = float64(1.0)
		return float64(0.0)
	}
	if snowfall > float64(0.0) {
		awe = float64((float64(snowpack.wsnow[i] - float64(snowfall*tStep))) / si)
		if (awe) >= (float64(0.0)) {
			awe = (awe)
		} else {
			awe = (float64(0.0))
		}
		sba = getArealSnowCover(i, awe)
		sbws = float64(awe + float64((float64(float64(float64(0.75)*snowfall)*tStep))/si))
		if (sbws) <= (float64(1.0)) {
			sbws = (sbws)
		} else {
			sbws = (float64(1.0))
		}
		snowpack.awe[i] = awe
		snowpack.sba[i] = sba
		snowpack.sbws[i] = sbws
		return float64(1.0)
	} else {
		awe = snowpack.awe[i]
		sba = snowpack.sba[i]
		sbws = snowpack.sbws[i]
		awesi = float64(snowpack.wsnow[i] / si)
		if awesi < snowpack.awe[i] {
			snowpack.awe[i] = float64(1.0)
			asc = getArealSnowCover(i, awesi)
		} else if awesi >= snowpack.sbws[i] {
			asc = float64(1.0)
		} else {
			asc = float64(sba + float64(float64((float64(float64(1.0)-sba))/(float64(sbws-awe)))*(float64(awesi-awe))))
		}
		return asc
	}
	return 0
}

// snow.c:636
func getArealSnowCover(i int32, awesi float64) float64 {
	var (
		k    int32
		m    int32
		asc1 float64
		asc2 float64
	)
	_, _, _, _ = k, m, asc1, asc2
	k = 0
	m = 0
	asc1 = 0
	asc2 = 0
	if i == SNOW_IMPERV {
		k = 0
	} else if i == SNOW_PERV {
		k = 1
	} else {
		return float64(1.0)
	}
	if awesi <= float64(0.0) {
		return float64(0.0)
	}
	if awesi >= float64(0.9999) {
		return float64(1.0)
	}
	m = R.D2I((float64(float64(awesi*float64(10.0)) + float64(0.00001))))
	asc1 = Snow.adc[k][m]
	if m >= 9 {
		asc2 = float64(1.0)
	} else {
		asc2 = Snow.adc[k][(m + 1)]
	}
	return float64(asc1 + float64(float64((float64(asc2-asc1))/float64(0.1))*(float64(awesi-float64(float64(0.1)*float64(float32(m)))))))
}

// snow.c:674
func meltSnowpack(snowpack *TSnowpack, i int32, rmelt float64, asc float64, snowfall float64, tStep float64) float64 {
	var (
		k        int32
		smelt    float64
		ccFactor float64
	)
	_, _, _ = k, smelt, ccFactor
	k = 0
	smelt = 0
	ccFactor = 0
	k = snowpack.snowmeltIndex
	if rmelt > float64(0.0) {
		smelt = rmelt
	} else if Temp.ta >= Snowmelt[k].tbase[i] {
		smelt = float64(Snowmelt[k].dhm[i] * (float64(Temp.ta - Snowmelt[k].tbase[i])))
	} else {
		updateColdContent(snowpack, i, asc, snowfall, tStep)
		return float64(0.0)
	}
	smelt = float64(smelt * asc)
	ccFactor = float64(float64(tStep*Snow.rnm) * asc)
	smelt = reduceColdContent(snowpack, i, smelt, ccFactor)
	snowpack.ati[i] = Snowmelt[k].tbase[i]
	return smelt
}

// snow.c:720
func getRainmelt(rainfall float64) float64 {
	var (
		uadj  float64
		t1    float64
		t2    float64
		t3    float64
		smelt float64
	)
	_, _, _, _, _ = uadj, t1, t2, t3, smelt
	uadj = 0
	t1 = 0
	t2 = 0
	t3 = 0
	smelt = 0
	rainfall = float64(rainfall * float64(43200.0))
	if rainfall > float64(0.02) {
		uadj = float64(float64(0.006) * Wind.ws)
		t1 = float64(Temp.ta - float64(32.0))
		t2 = float64(float64(float64(7.5)*Temp.gamma) * uadj)
		t3 = float64(float64(float64(8.5)*uadj) * (float64(Temp.ea - float64(0.18))))
		smelt = float64(float64(t1*(float64(float64(float64(0.001167)+t2)+float64(float64(0.007)*rainfall)))) + t3)
		return float64(smelt / float64(43200.0))
	} else {
		return float64(0.0)
	}
	return 0
}

// snow.c:746
func updateColdContent(snowpack *TSnowpack, i int32, asc float64, snowfall float64, tStep float64) {
	var (
		k     int32
		ati   float64
		cc    float64
		ccMax float64
		tipm  float64
	)
	_, _, _, _, _ = k, ati, cc, ccMax, tipm
	k = 0
	ati = 0
	cc = 0
	ccMax = 0
	tipm = 0
	ati = snowpack.ati[i]
	cc = snowpack.coldc[i]
	if float64(snowfall*float64(43200.0)) > float64(0.02) {
		ati = Temp.ta
	} else {
		tipm = float64(float64(1.0) - M.Pow(float64(float64(1.0)-Snow.tipm), float64(tStep/(float64(float64(6.0)*float64(3600.0))))))
		ati = float64(ati + (float64(tipm * (float64(Temp.ta - ati)))))
	}
	k = snowpack.snowmeltIndex
	if (ati) <= (Snowmelt[k].tbase[i]) {
		ati = (ati)
	} else {
		ati = (Snowmelt[k].tbase[i])
	}
	cc = float64(cc + (float64(float64(float64(float64(Snow.rnm*Snowmelt[k].dhm[i])*(float64(ati-Temp.ta)))*tStep) * asc)))
	if (cc) >= (float64(0.0)) {
		cc = (cc)
	} else {
		cc = (float64(0.0))
	}
	ccMax = float64(float64(float64(snowpack.wsnow[i]*float64(0.007))/float64(12.0)) * (float64(Snowmelt[k].tbase[i] - ati)))
	if (cc) <= (ccMax) {
		cc = (cc)
	} else {
		cc = (ccMax)
	}
	snowpack.coldc[i] = cc
	snowpack.ati[i] = ati
}

// snow.c:799
func reduceColdContent(snowpack *TSnowpack, i int32, smelt float64, ccFactor float64) float64 {
	var (
		cc float64
	)
	_ = cc
	cc = 0
	cc = snowpack.coldc[i]
	if float64(smelt*ccFactor) > cc {
		smelt = float64(smelt - (float64(cc / ccFactor)))
		cc = float64(0.0)
	} else {
		cc = float64(cc - (float64(smelt * ccFactor)))
		smelt = float64(0.0)
	}
	snowpack.coldc[i] = cc
	return smelt
}

// snow.c:828
func routeSnowmelt(snowpack *TSnowpack, i int32, smelt float64, asc float64, rainfall float64, tStep float64) float64 {
	var (
		k     int32
		vmelt float64
	)
	_, _ = k, vmelt
	k = 0
	vmelt = 0
	k = snowpack.snowmeltIndex
	vmelt = float64(smelt * tStep)
	if (vmelt) <= (snowpack.wsnow[i]) {
		vmelt = (vmelt)
	} else {
		vmelt = (snowpack.wsnow[i])
	}
	snowpack.wsnow[i] = float64(snowpack.wsnow[i] - vmelt)
	snowpack.fw[i] = float64(snowpack.fw[i] + (float64(vmelt + float64(float64(rainfall*tStep)*asc))))
	vmelt = float64(snowpack.fw[i] - float64(Snowmelt[k].fwfrac[i]*snowpack.wsnow[i]))
	if (vmelt) >= (float64(0.0)) {
		vmelt = (vmelt)
	} else {
		vmelt = (float64(0.0))
	}
	snowpack.fw[i] = float64(snowpack.fw[i] - vmelt)
	return float64(vmelt / tStep)
}
