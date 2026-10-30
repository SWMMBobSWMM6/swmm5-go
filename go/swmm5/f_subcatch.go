// Generated from EPA SWMM 5.2.4 (subcatch.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var MCOEFF float64 = float64(1.49)
var MEXP float64 = float64(1.6666667)
var ODETOL float64 = float64(0.0001)
var Vevap float64
var Vpevap float64
var Vinfil float64
var Vinflow float64
var Voutflow float64
var VlidIn float64
var VlidInfil float64
var VlidOut float64
var VlidDrain float64
var VlidReturn float64
var theSubarea *TSubarea
var Dstore float64
var Alpha float64
var RunoffRoutingWords = [4]R.CPtr{L663, L2, L3, (R.CPtr{})}

// subcatch.c:119
func subcatch_readParams(j int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		i  int32
		k  int32
		m  int32
		id R.CPtr
		x  [9]float64
	)
	_, _, _, _, _ = i, k, m, id, x
	i = 0
	k = 0
	m = 0
	id = (R.CPtr{})
	x = ([9]float64{})
	if ntoks < 8 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(SUBCATCH, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	k = project_findObject(GAGE, tok[1])
	if k < 0 {
		return error_setInpError(ERR_NAME, tok[1])
	}
	x[0] = float64(k)
	m = project_findObject(NODE, tok[2])
	x[1] = float64(m)
	m = project_findObject(SUBCATCH, tok[2])
	x[2] = float64(m)
	if (x[1] < float64(0.0)) && (x[2] < float64(0.0)) {
		return error_setInpError(ERR_NAME, tok[2])
	}
	for i = 3; i < 8; i++ {
		if (!(getDouble(tok[i], &x[i]) != 0)) || (x[i] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[i])
		}
	}
	x[8] = float64((-1))
	if ntoks > 8 {
		k = project_findObject(SNOWMELT, tok[8])
		if k < 0 {
			return error_setInpError(ERR_NAME, tok[8])
		}
		x[8] = float64(k)
	}
	Subcatch[j].ID = id
	Subcatch[j].gage = R.D2I(x[0])
	Subcatch[j].outNode = R.D2I(x[1])
	Subcatch[j].outSubcatch = R.D2I(x[2])
	Subcatch[j].area = float64(x[3] / UCF(LANDAREA))
	Subcatch[j].fracImperv = float64((func() float64 {
		if (x[4]) <= (float64(100.0)) {
			return (x[4])
		}
		return (float64(100.0))
	}()) / float64(100.0))
	Subcatch[j].width = float64(x[5] / UCF(LENGTH))
	Subcatch[j].slope = float64(x[6] / float64(100.0))
	Subcatch[j].curbLength = x[7]
	Subcatch[j].nPervPattern = (-1)
	Subcatch[j].dStorePattern = (-1)
	Subcatch[j].infilPattern = (-1)
	if x[8] >= float64(0) {
		if !(snow_createSnowpack(j, R.D2I(x[8])) != 0) {
			return error_setInpError(ERR_MEMORY, L0)
		}
	}
	return 0
}

// subcatch.c:196
func subcatch_readSubareaParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		i int32
		j int32
		k int32
		m int32
		x [7]float64
	)
	_, _, _, _, _ = i, j, k, m, x
	i = 0
	j = 0
	k = 0
	m = 0
	x = ([7]float64{})
	if ntoks < 7 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(SUBCATCH, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	for i = 0; i < 5; i++ {
		if (!(getDouble(tok[(i+1)], &x[i]) != 0)) || (x[i] < float64(0.0)) {
			return error_setInpError(ERR_NAME, tok[(i+1)])
		}
	}
	m = findmatch(tok[6], RunoffRoutingWords[:])
	if m < 0 {
		return error_setInpError(ERR_KEYWORD, tok[6])
	}
	x[5] = float64(m)
	x[6] = float64(1.0)
	if ntoks >= 8 {
		if ((!(getDouble(tok[7], &x[6]) != 0)) || (x[6] < float64(0.0))) || (x[6] > float64(100.0)) {
			return error_setInpError(ERR_NUMBER, tok[7])
		}
		x[6] = float64(x[6] / (float64(100.0)))
	}
	Subcatch[j].subArea[IMPERV0].N = x[0]
	Subcatch[j].subArea[IMPERV1].N = x[0]
	Subcatch[j].subArea[PERV].N = x[1]
	Subcatch[j].subArea[IMPERV0].dStore = float64(0.0)
	Subcatch[j].subArea[IMPERV1].dStore = float64(x[2] / UCF(RAINDEPTH))
	Subcatch[j].subArea[PERV].dStore = float64(x[3] / UCF(RAINDEPTH))
	Subcatch[j].subArea[IMPERV0].fArea = float64(float64(Subcatch[j].fracImperv*x[4]) / float64(100.0))
	Subcatch[j].subArea[IMPERV1].fArea = float64(Subcatch[j].fracImperv * (float64(float64(1.0) - float64(x[4]/float64(100.0)))))
	Subcatch[j].subArea[PERV].fArea = (float64(float64(1.0) - Subcatch[j].fracImperv))
	for i = IMPERV0; i <= PERV; i++ {
		Subcatch[j].subArea[i].routeTo = TO_OUTLET
		Subcatch[j].subArea[i].fOutlet = float64(1.0)
	}
	k = R.D2I(x[5])
	if (Subcatch[j].fracImperv == float64(0.0)) || (Subcatch[j].fracImperv == float64(1.0)) {
		k = TO_OUTLET
	}
	if (k == TO_IMPERV) && (Subcatch[j].fracImperv != 0) {
		Subcatch[j].subArea[PERV].routeTo = k
		Subcatch[j].subArea[PERV].fOutlet = float64(float64(1.0) - x[6])
	}
	if k == TO_PERV {
		Subcatch[j].subArea[IMPERV0].routeTo = k
		Subcatch[j].subArea[IMPERV1].routeTo = k
		Subcatch[j].subArea[IMPERV0].fOutlet = float64(float64(1.0) - x[6])
		Subcatch[j].subArea[IMPERV1].fOutlet = float64(float64(1.0) - x[6])
	}
	return 0
}

// subcatch.c:284
func subcatch_readLanduseParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		j int32
		k int32
		m int32
		f float64
	)
	_, _, _, _ = j, k, m, f
	j = 0
	k = 0
	m = 0
	f = 0
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(SUBCATCH, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	for k = 2; k <= ntoks; k = (k + 2) {
		m = project_findObject(LANDUSE, tok[(k-1)])
		if m < 0 {
			return error_setInpError(ERR_NAME, tok[(k-1)])
		}
		if (k + 1) > ntoks {
			return error_setInpError(ERR_ITEMS, L0)
		}
		if !(getDouble(tok[k], &f) != 0) {
			return error_setInpError(ERR_NUMBER, tok[k])
		}
		Subcatch[j].landFactor[m].fraction = float64(f / float64(100.0))
	}
	return 0
}

// subcatch.c:324
func subcatch_readInitBuildup(tok []R.CPtr, ntoks int32) int32 {
	var (
		j int32
		k int32
		m int32
		x float64
	)
	_, _, _, _ = j, k, m, x
	j = 0
	k = 0
	m = 0
	x = 0
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(SUBCATCH, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	for k = 2; k <= ntoks; k = (k + 2) {
		m = project_findObject(POLLUT, tok[(k-1)])
		if m < 0 {
			return error_setInpError(ERR_NAME, tok[(k-1)])
		}
		if (k + 1) > ntoks {
			return error_setInpError(ERR_ITEMS, L0)
		}
		if !(getDouble(tok[k], &x) != 0) {
			return error_setInpError(ERR_NUMBER, tok[k])
		}
		Subcatch[j].initBuildup[m] = x
	}
	return 0
}

// subcatch.c:364
func subcatch_validate(j int32) {
	var (
		i          int32
		area       float64
		nonLidArea float64
	)
	_, _, _ = i, area, nonLidArea
	i = 0
	area = 0
	nonLidArea = Subcatch[j].area
	if (Subcatch[j].outNode >= 0) && (Subcatch[j].outSubcatch >= 0) {
		report_writeErrorMsg(ERR_SUBCATCH_OUTLET, Subcatch[j].ID)
	}
	gwater_validate(j)
	nonLidArea = float64(nonLidArea - Subcatch[j].lidArea)
	for i = IMPERV0; i <= PERV; i++ {
		if i == PERV {
			area = float64((float64(float64(1.0) - Subcatch[j].fracImperv)) * nonLidArea)
		} else {
			area = float64(Subcatch[j].fracImperv * nonLidArea)
		}
		Subcatch[j].subArea[i].alpha = float64(0.0)
		if (area > float64(0.0)) && (Subcatch[j].subArea[i].N > float64(0.0)) {
			Subcatch[j].subArea[i].alpha = float64(float64(float64(float64(MCOEFF*Subcatch[j].width)/area)*M.Sqrt(Subcatch[j].slope)) / Subcatch[j].subArea[i].N)
		}
	}
	i = Subcatch[j].gage
	if i >= 0 {
		Gage[i].isUsed = 1
	}
}

// subcatch.c:419
func subcatch_initState(j int32) {
	var (
		i int32
	)
	_ = i
	i = 0
	Subcatch[j].rainfall = float64(0.0)
	Subcatch[j].oldRunoff = float64(0.0)
	Subcatch[j].newRunoff = float64(0.0)
	Subcatch[j].oldSnowDepth = float64(0.0)
	Subcatch[j].newSnowDepth = float64(0.0)
	Subcatch[j].runon = float64(0.0)
	Subcatch[j].evapLoss = float64(0.0)
	Subcatch[j].infilLoss = float64(0.0)
	if Subcatch[j].infil == j {
		infil_initState(j)
	}
	if Subcatch[j].groundwater != nil {
		gwater_initState(j)
	}
	if Subcatch[j].snowpack != nil {
		snow_initSnowpack(j)
	}
	for i = IMPERV0; i <= PERV; i++ {
		Subcatch[j].subArea[i].depth = float64(0.0)
		Subcatch[j].subArea[i].inflow = float64(0.0)
		Subcatch[j].subArea[i].runoff = float64(0.0)
	}
	surfqual_initState(j)
}

// subcatch.c:457
func subcatch_setOldState(j int32) {
	var (
		i int32
	)
	_ = i
	i = 0
	Subcatch[j].oldRunoff = Subcatch[j].newRunoff
	Subcatch[j].oldSnowDepth = Subcatch[j].newSnowDepth
	for i = IMPERV0; i <= PERV; i++ {
		Subcatch[j].subArea[i].inflow = float64(0.0)
	}
	for i = 0; i < Nobjects[POLLUT]; i++ {
		Subcatch[j].oldQual[i] = Subcatch[j].newQual[i]
		Subcatch[j].newQual[i] = float64(0.0)
	}
	lid_setOldGroupState(j)
}

// subcatch.c:481
func subcatch_getFracPerv(j int32) float64 {
	var (
		fracPerv float64
	)
	_ = fracPerv
	fracPerv = float64(float64(1.0) - Subcatch[j].fracImperv)
	if Subcatch[j].lidArea > float64(0.0) {
		fracPerv = float64((float64(float64(fracPerv*(float64(Subcatch[j].area-Subcatch[j].lidArea))) + lid_getPervArea(j))) / Subcatch[j].area)
		if (fracPerv) <= (float64(1.0)) {
			fracPerv = (fracPerv)
		} else {
			fracPerv = (float64(1.0))
		}
	}
	return fracPerv
}

// subcatch.c:502
func subcatch_getStorage(j int32) float64 {
	var (
		i int32
		v float64
	)
	_, _ = i, v
	i = 0
	v = float64(0.0)
	for i = IMPERV0; i <= PERV; i++ {
		v = float64(v + (float64(Subcatch[j].subArea[i].depth * Subcatch[j].subArea[i].fArea)))
	}
	return float64(float64(v*(float64(Subcatch[j].area-Subcatch[j].lidArea))) + lid_getStoredVolume(j))
}

// subcatch.c:523
func subcatch_getRunon(j int32) {
	var (
		k        int32
		p        int32
		q        float64
		q1       float64
		q2       float64
		pervArea float64
	)
	_, _, _, _, _, _ = k, p, q, q1, q2, pervArea
	k = 0
	p = 0
	q = 0
	q1 = 0
	q2 = 0
	pervArea = 0
	k = Subcatch[j].outSubcatch
	q = Subcatch[j].oldRunoff
	if (k >= 0) && (k != j) {
		subcatch_addRunonFlow(k, q)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			Subcatch[k].newQual[p] = float64(Subcatch[k].newQual[p] + (float64(float64(q*Subcatch[j].oldQual[p]) * float64(28.317))))
		}
	}
	if Subcatch[j].lidArea > float64(0.0) {
		lid_addDrainRunon(j)
	}
	if (Subcatch[j].fracImperv < float64(1.0)) && (Subcatch[j].subArea[IMPERV0].routeTo == TO_PERV) {
		q1 = float64(Subcatch[j].subArea[IMPERV0].runoff * Subcatch[j].subArea[IMPERV0].fArea)
		q2 = float64(Subcatch[j].subArea[IMPERV1].runoff * Subcatch[j].subArea[IMPERV1].fArea)
		q = float64(q1 + q2)
		Subcatch[j].subArea[PERV].inflow = float64(Subcatch[j].subArea[PERV].inflow + (float64(float64(q*(float64(float64(1.0)-Subcatch[j].subArea[IMPERV0].fOutlet))) / Subcatch[j].subArea[PERV].fArea)))
	}
	if ((Subcatch[j].fracImperv > float64(0.0)) && (Subcatch[j].subArea[PERV].routeTo == TO_IMPERV)) && (Subcatch[j].subArea[IMPERV1].fArea > float64(0.0)) {
		q = Subcatch[j].subArea[PERV].runoff
		Subcatch[j].subArea[IMPERV1].inflow = float64(Subcatch[j].subArea[IMPERV1].inflow + (float64(float64(float64(q*(float64(float64(1.0)-Subcatch[j].subArea[PERV].fOutlet)))*Subcatch[j].subArea[PERV].fArea) / Subcatch[j].subArea[IMPERV1].fArea)))
	}
	if (Subcatch[j].lidArea > float64(0.0)) && (Subcatch[j].fracImperv < float64(1.0)) {
		pervArea = float64(Subcatch[j].subArea[PERV].fArea * (float64(Subcatch[j].area - Subcatch[j].lidArea)))
		q = lid_getFlowToPerv(j)
		if pervArea > float64(0.0) {
			Subcatch[j].subArea[PERV].inflow = float64(Subcatch[j].subArea[PERV].inflow + (float64(q / pervArea)))
		}
	}
}

// subcatch.c:600
func subcatch_addRunonFlow(k int32, q float64) {
	var (
		i          int32
		nonLidArea float64
	)
	_, _ = i, nonLidArea
	i = 0
	nonLidArea = 0
	if Subcatch[k].area <= float64(0.0) {
		return
	}
	nonLidArea = float64(Subcatch[k].area - Subcatch[k].lidArea)
	if nonLidArea > float64(0.0) {
		q = float64(q / nonLidArea)
	} else {
		q = float64(q / Subcatch[k].area)
	}
	Subcatch[k].runon = float64(Subcatch[k].runon + q)
	for i = IMPERV0; i <= PERV; i++ {
		Subcatch[k].subArea[i].inflow = float64(Subcatch[k].subArea[i].inflow + q)
	}
}

// subcatch.c:629
func subcatch_getRunoff(j int32, tStep float64) float64 {
	var (
		i             int32
		nonLidArea    float64
		area          float64
		netPrecip     [3]float64
		vRain         float64
		vRunon        float64
		vOutflow      float64
		runoff        float64
		evapRate      float64
		subAreaRunoff float64
		vImpervRunoff float64
		vPervRunoff   float64
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = i, nonLidArea, area, netPrecip, vRain, vRunon, vOutflow, runoff, evapRate, subAreaRunoff, vImpervRunoff, vPervRunoff
	i = 0
	nonLidArea = 0
	area = 0
	netPrecip = ([3]float64{})
	vRain = 0
	vRunon = float64(0.0)
	vOutflow = float64(0.0)
	runoff = float64(0.0)
	evapRate = float64(0.0)
	subAreaRunoff = 0
	vImpervRunoff = float64(0.0)
	vPervRunoff = float64(0.0)
	Vevap = float64(0.0)
	Vpevap = float64(0.0)
	Vinfil = float64(0.0)
	Voutflow = float64(0.0)
	VlidIn = float64(0.0)
	VlidInfil = float64(0.0)
	VlidOut = float64(0.0)
	VlidDrain = float64(0.0)
	VlidReturn = float64(0.0)
	nonLidArea = float64(Subcatch[j].area - Subcatch[j].lidArea)
	vRunon = float64(float64(Subcatch[j].runon*tStep) * nonLidArea)
	Vinflow = float64(vRunon + float64(subcatch_getDepth(j)*nonLidArea))
	if nonLidArea == float64(0.0) {
		vRunon = float64(float64(Subcatch[j].runon*tStep) * Subcatch[j].area)
	}
	getNetPrecip(j, netPrecip[:], tStep)
	if (Evap.dryOnly != 0) && (Subcatch[j].rainfall > float64(0.0)) {
		evapRate = float64(0.0)
	} else {
		evapRate = Evap.rate
	}
	infil_setInfilFactor(j)
	if nonLidArea > float64(0.0) {
		for i = IMPERV0; i <= PERV; i++ {
			area = float64(nonLidArea * Subcatch[j].subArea[i].fArea)
			Subcatch[j].subArea[i].runoff = getSubareaRunoff(j, i, area, netPrecip[i], evapRate, tStep)
			subAreaRunoff = float64(Subcatch[j].subArea[i].runoff * area)
			if i == PERV {
				vPervRunoff = float64(subAreaRunoff * tStep)
			} else {
				vImpervRunoff = float64(vImpervRunoff + (float64(subAreaRunoff * tStep)))
			}
			runoff = float64(runoff + subAreaRunoff)
		}
	}
	if Subcatch[j].lidArea > float64(0.0) {
		lid_getRunoff(j, tStep)
	}
	if (!(IgnoreGwater != 0)) && (Subcatch[j].groundwater != nil) {
		gwater_getGroundwater(j, Vpevap, float64(Vinfil+VlidInfil), tStep)
	}
	area = Subcatch[j].area
	Subcatch[j].evapLoss = float64(float64(Vevap/tStep) / area)
	Subcatch[j].infilLoss = float64(float64((float64(Vinfil+VlidInfil))/tStep) / area)
	vOutflow = float64(float64(Voutflow-VlidIn) + VlidOut)
	Subcatch[j].newRunoff = float64(vOutflow / tStep)
	vRain = float64(float64(Subcatch[j].rainfall*tStep) * area)
	stats_updateSubcatchStats(j, vRain, vRunon, Vevap, float64(Vinfil+VlidInfil), vImpervRunoff, vPervRunoff, float64(vOutflow+VlidDrain), float64(Subcatch[j].newRunoff+float64(VlidDrain/tStep)))
	if (Subcatch[j].outNode == (-1)) && (Subcatch[j].outSubcatch != j) {
		vOutflow = float64(0.0)
	}
	massbal_updateRunoffTotals(RUNOFF_RAINFALL, vRain)
	massbal_updateRunoffTotals(RUNOFF_EVAP, Vevap)
	massbal_updateRunoffTotals(RUNOFF_INFIL, float64(Vinfil+VlidInfil))
	massbal_updateRunoffTotals(RUNOFF_RUNOFF, vOutflow)
	return float64(runoff / area)
}

// subcatch.c:758
func getNetPrecip(j int32, netPrecip []float64, tStep float64) {
	var (
		i        int32
		k        int32
		rainfall float64
		snowfall float64
	)
	_, _, _, _ = i, k, rainfall, snowfall
	i = 0
	k = 0
	rainfall = float64(0.0)
	snowfall = float64(0.0)
	k = Subcatch[j].gage
	if k >= 0 {
		gage_getPrecip(k, &rainfall, &snowfall)
	}
	Subcatch[j].rainfall = float64(rainfall + snowfall)
	if (Subcatch[j].snowpack != nil) && (!(IgnoreSnowmelt != 0)) {
		Subcatch[j].newSnowDepth = snow_getSnowMelt(j, rainfall, snowfall, tStep, netPrecip)
	} else {
		for i = IMPERV0; i <= PERV; i++ {
			netPrecip[i] = float64(rainfall + snowfall)
		}
	}
}

// subcatch.c:798
func subcatch_getDepth(j int32) float64 {
	var (
		i     int32
		fArea float64
		depth float64
	)
	_, _, _ = i, fArea, depth
	i = 0
	fArea = 0
	depth = float64(0.0)
	for i = IMPERV0; i <= PERV; i++ {
		fArea = Subcatch[j].subArea[i].fArea
		if fArea > float64(0.0) {
			depth = float64(depth + (float64(Subcatch[j].subArea[i].depth * fArea)))
		}
	}
	return depth
}

// subcatch.c:820
func subcatch_getWtdOutflow(j int32, f float64) float64 {
	if Subcatch[j].area == float64(0.0) {
		return float64(0.0)
	}
	return float64(float64((float64(float64(1.0)-f))*Subcatch[j].oldRunoff) + float64(f*Subcatch[j].newRunoff))
}

// subcatch.c:834
func subcatch_getResults(j int32, f float64, x []float32) {
	var (
		p      int32
		k      int32
		f1     float64
		z      float64
		runoff float64
		gw     *TGroundwater
	)
	_, _, _, _, _, _ = p, k, f1, z, runoff, gw
	p = 0
	k = 0
	f1 = float64(float64(1.0) - f)
	z = 0
	runoff = 0
	gw = nil
	k = Subcatch[j].gage
	if k >= 0 {
		x[SUBCATCH_RAINFALL] = float32(float64(Gage[k].reportRainfall))
	} else {
		x[SUBCATCH_RAINFALL] = float32(0.0)
	}
	z = float64((float64(float64(f1*Subcatch[j].oldSnowDepth) + float64(f*Subcatch[j].newSnowDepth))) * UCF(RAINDEPTH))
	x[SUBCATCH_SNOWDEPTH] = float32(float64(z))
	x[SUBCATCH_EVAP] = float32(float64((float64(Subcatch[j].evapLoss * UCF(EVAPRATE)))))
	x[SUBCATCH_INFIL] = float32(float64((float64(Subcatch[j].infilLoss * UCF(RAINFALL)))))
	runoff = float64(float64(f1*Subcatch[j].oldRunoff) + float64(f*Subcatch[j].newRunoff))
	if Subcatch[j].lidArea > float64(0.0) {
		runoff = float64(runoff + (float64(float64(f1*lid_getDrainFlow(j, PREVIOUS)) + float64(f*lid_getDrainFlow(j, CURRENT)))))
	}
	if runoff < float64(float64(2.31481e-08)*Subcatch[j].area) {
		runoff = float64(0.0)
	}
	x[SUBCATCH_RUNOFF] = float32(float64((float64(runoff * UCF(FLOW)))))
	gw = Subcatch[j].groundwater
	if gw != nil {
		z = float64(float64((float64(float64(f1*gw.oldFlow)+float64(f*gw.newFlow)))*Subcatch[j].area) * UCF(FLOW))
		x[SUBCATCH_GW_FLOW] = float32(float64(z))
		z = float64((float64(gw.bottomElev + gw.lowerDepth)) * UCF(LENGTH))
		x[SUBCATCH_GW_ELEV] = float32(float64(z))
		z = gw.theta
		x[SUBCATCH_SOIL_MOIST] = float32(float64(z))
	} else {
		x[SUBCATCH_GW_FLOW] = float32(0.0)
		x[SUBCATCH_GW_ELEV] = float32(0.0)
		x[SUBCATCH_SOIL_MOIST] = float32(0.0)
	}
	if !(IgnoreQuality != 0) {
		for p = 0; p < Nobjects[POLLUT]; p++ {
			if runoff == float64(0.0) {
				z = float64(0.0)
			} else {
				z = float64(float64(f1*Subcatch[j].oldQual[p]) + float64(f*Subcatch[j].newQual[p]))
			}
			x[(SUBCATCH_WASHOFF + p)] = float32(float64(z))
		}
	}
}

// subcatch.c:907
func getSubareaRunoff(j int32, i int32, area float64, precip float64, evap float64, tStep float64) float64 {
	var (
		tRunoff      float64
		surfMoisture float64
		surfEvap     float64
		infil        float64
		runoff       float64
		subarea      *TSubarea
	)
	_, _, _, _, _, _ = tRunoff, surfMoisture, surfEvap, infil, runoff, subarea
	tRunoff = 0
	surfMoisture = 0
	surfEvap = 0
	infil = float64(0.0)
	runoff = float64(0.0)
	subarea = nil
	if area == float64(0.0) {
		return float64(0.0)
	}
	subarea = &Subcatch[j].subArea[i]
	tRunoff = tStep
	surfMoisture = float64(subarea.depth / tStep)
	if (surfMoisture) <= (evap) {
		surfEvap = (surfMoisture)
	} else {
		surfEvap = (evap)
	}
	if i == PERV {
		infil = getSubareaInfil(j, subarea, precip, tStep)
	}
	subarea.inflow = float64(subarea.inflow + precip)
	surfMoisture = float64(surfMoisture + subarea.inflow)
	Vinflow = float64(Vinflow + (float64(float64(precip*area) * tStep)))
	Vevap = float64(Vevap + (float64(float64(surfEvap*area) * tStep)))
	if i == PERV {
		Vpevap = float64(Vpevap + Vevap)
	}
	Vinfil = float64(Vinfil + (float64(float64(infil*area) * tStep)))
	Alpha = subarea.alpha
	Dstore = subarea.dStore
	adjustSubareaParams(i, j)
	if float64(surfEvap+infil) >= surfMoisture {
		subarea.depth = float64(0.0)
	} else {
		subarea.inflow = float64(subarea.inflow - (float64(surfEvap + infil)))
		updatePondedDepth(subarea, &tRunoff)
	}
	runoff = findSubareaRunoff(subarea, tRunoff)
	Voutflow = float64(Voutflow + (float64(float64(float64(subarea.fOutlet*runoff)*area) * tStep)))
	return runoff
}

// subcatch.c:985
func getSubareaInfil(j int32, subarea *TSubarea, precip float64, tStep float64) float64 {
	var (
		infil float64
	)
	_ = infil
	infil = float64(0.0)
	infil = infil_getInfil(j, tStep, precip, subarea.inflow, subarea.depth)
	if (!(IgnoreGwater != 0)) && (Subcatch[j].groundwater != nil) {
		if (infil) <= (float64(Subcatch[j].groundwater.maxInfilVol / tStep)) {
			infil = (infil)
		} else {
			infil = (float64(Subcatch[j].groundwater.maxInfilVol / tStep))
		}
	}
	return infil
}

// subcatch.c:1013
func findSubareaRunoff(subarea *TSubarea, tRunoff float64) float64 {
	var (
		xDepth float64
		runoff float64
	)
	_, _ = xDepth, runoff
	xDepth = float64(subarea.depth - Dstore)
	runoff = float64(0.0)
	if xDepth > float64(1e-10) {
		if subarea.N > float64(0.0) {
			runoff = float64(Alpha * M.Pow(xDepth, MEXP))
		} else {
			runoff = float64(xDepth / tRunoff)
			subarea.depth = Dstore
		}
	} else {
		runoff = float64(0.0)
	}
	return runoff
}

// subcatch.c:1048
func updatePondedDepth(subarea *TSubarea, dt *float64) {
	var (
		ix float64
		dx float64
		tx float64
	)
	_, _, _ = ix, dx, tx
	ix = subarea.inflow
	dx = 0
	tx = (*dt)
	if float64(subarea.depth+float64(ix*tx)) <= Dstore {
		subarea.depth = float64(subarea.depth + (float64(ix * tx)))
	} else {
		dx = float64(Dstore - subarea.depth)
		if (dx > float64(0.0)) && (ix > float64(0.0)) {
			tx = float64(tx - (float64(dx / ix)))
			subarea.depth = Dstore
		}
		if (Alpha > float64(0.0)) && (tx > float64(0.0)) {
			theSubarea = subarea
			odesolve_integrate(R.One(&subarea.depth), 1, float64(0), tx, ODETOL, tx, getDdDt)
		} else {
			if tx < float64(0.0) {
				tx = float64(0.0)
			}
			subarea.depth = float64(subarea.depth + (float64(ix * tx)))
		}
	}
	if subarea.depth < float64(0.0) {
		subarea.depth = float64(0.0)
	}
	*dt = tx
}

// subcatch.c:1101
func getDdDt(t float64, d []float64, dddt []float64) {
	var (
		ix float64
		rx float64
	)
	_, _ = ix, rx
	ix = theSubarea.inflow
	rx = float64(d[0] - Dstore)
	if rx < float64(0.0) {
		rx = float64(0.0)
	} else {
		rx = float64(Alpha * M.Pow(rx, MEXP))
	}
	dddt[0] = float64(ix - rx)
}

// subcatch.c:1125
func adjustSubareaParams(i int32, j int32) {
	var (
		p int32
		m int32
		f float64
	)
	_, _, _ = p, m, f
	p = 0
	m = 0
	f = 0
	if i == PERV {
		p = Subcatch[j].dStorePattern
		if (p >= 0) && (Pattern[p].type_ == MONTHLY_PATTERN) {
			m = (datetime_monthOfYear(getDateTime(OldRunoffTime)) - 1)
			f = Pattern[p].factor[m]
			if f >= float64(0.0) {
				Dstore = float64(Dstore * f)
			}
		}
		p = Subcatch[j].nPervPattern
		if (p >= 0) && (Pattern[p].type_ == MONTHLY_PATTERN) {
			m = (datetime_monthOfYear(getDateTime(OldRunoffTime)) - 1)
			f = Pattern[p].factor[m]
			if f <= float64(0.0) {
				Alpha = float64(0.0)
			} else {
				Alpha = float64(Alpha / f)
			}
		}
	}
}
