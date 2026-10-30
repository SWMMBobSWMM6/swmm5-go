// Generated from EPA SWMM 5.2.4 (surfqual.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

// surfqual.c:65
func surfqual_initState(j int32) {
	var (
		p int32
	)
	_ = p
	p = 0
	for p = 0; p < Nobjects[POLLUT]; p++ {
		Subcatch[j].oldQual[p] = float64(0.0)
		Subcatch[j].newQual[p] = float64(0.0)
		Subcatch[j].pondedQual[p] = float64(0.0)
	}
	landuse_getInitBuildup(Subcatch[j].landFactor, Subcatch[j].initBuildup, Subcatch[j].area, Subcatch[j].curbLength)
}

// surfqual.c:89
func surfqual_getBuildup(j int32, tStep float64) {
	var (
		i          int32
		p          int32
		f          float64
		area       float64
		curb       float64
		oldBuildup float64
		newBuildup float64
	)
	_, _, _, _, _, _, _ = i, p, f, area, curb, oldBuildup, newBuildup
	i = 0
	p = 0
	f = 0
	area = 0
	curb = 0
	oldBuildup = 0
	newBuildup = 0
	for i = 0; i < Nobjects[LANDUSE]; i++ {
		f = Subcatch[j].landFactor[i].fraction
		if f == float64(0.0) {
			continue
		}
		area = float64(float64(f*Subcatch[j].area) * UCF(LANDAREA))
		curb = float64(f * Subcatch[j].curbLength)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			if (Pollut[p].snowOnly != 0) && (Subcatch[j].newSnowDepth < float64(float64(0.001)/float64(12.0))) {
				continue
			}
			oldBuildup = Subcatch[j].landFactor[i].buildup[p]
			newBuildup = landuse_getBuildup(i, p, area, curb, oldBuildup, tStep)
			if (newBuildup) >= (oldBuildup) {
				newBuildup = (newBuildup)
			} else {
				newBuildup = (oldBuildup)
			}
			Subcatch[j].landFactor[i].buildup[p] = newBuildup
			massbal_updateLoadingTotals(BUILDUP_LOAD, p, (float64(newBuildup - oldBuildup)))
		}
	}
}

// surfqual.c:137
func surfqual_sweepBuildup(j int32, aDate float64) {
	var (
		i          int32
		p          int32
		oldBuildup float64
		newBuildup float64
	)
	_, _, _, _ = i, p, oldBuildup, newBuildup
	i = 0
	p = 0
	oldBuildup = 0
	newBuildup = 0
	if (Subcatch[j].snowpack != nil) && (Subcatch[j].snowpack.wsnow[IMPERV0] > float64(0.004167)) {
		return
	}
	for i = 0; i < Nobjects[LANDUSE]; i++ {
		if Subcatch[j].landFactor[i].fraction == float64(0.0) {
			continue
		}
		if Landuse[i].sweepInterval == float64(0.0) {
			continue
		}
		if float64(aDate-Subcatch[j].landFactor[i].lastSwept) >= Landuse[i].sweepInterval {
			Subcatch[j].landFactor[i].lastSwept = aDate
			for p = 0; p < Nobjects[POLLUT]; p++ {
				oldBuildup = Subcatch[j].landFactor[i].buildup[p]
				newBuildup = float64(oldBuildup * (float64(float64(1.0) - float64(Landuse[i].sweepRemoval*Landuse[i].washoffFunc[p].sweepEffic))))
				if (oldBuildup) <= (newBuildup) {
					newBuildup = (oldBuildup)
				} else {
					newBuildup = (newBuildup)
				}
				if (float64(0.0)) >= (newBuildup) {
					newBuildup = (float64(0.0))
				} else {
					newBuildup = (newBuildup)
				}
				Subcatch[j].landFactor[i].buildup[p] = newBuildup
				massbal_updateLoadingTotals(SWEEPING_LOAD, p, float64(oldBuildup-newBuildup))
			}
		}
	}
}

// surfqual.c:192
func surfqual_getWashoff(j int32, runoff float64, tStep float64) {
	var (
		p          int32
		hasOutflow int32
		cOut       float64
		massLoad   float64
		vLidRain   float64
		vLidRunon  float64
		vSurfOut   float64
		vOut1      float64
		vOut2      float64
		area       float64
	)
	_, _, _, _, _, _, _, _, _, _ = p, hasOutflow, cOut, massLoad, vLidRain, vLidRunon, vSurfOut, vOut1, vOut2, area
	p = 0
	hasOutflow = 0
	cOut = 0
	massLoad = 0
	vLidRain = 0
	vLidRunon = 0
	vSurfOut = 0
	vOut1 = 0
	vOut2 = 0
	area = 0
	area = Subcatch[j].area
	if (Nobjects[POLLUT] == 0) || (area == float64(0.0)) {
		return
	}
	for p = 0; p < Nobjects[POLLUT]; p++ {
		OutflowLoad[p] = float64(0.0)
	}
	findWashoffLoads(j, runoff)
	findPondedLoads(j, tStep)
	findLidLoads(j, tStep)
	vLidRain = float64(float64(Subcatch[j].rainfall*Subcatch[j].lidArea) * tStep)
	vLidRunon = float64(0.0)
	if area == Subcatch[j].lidArea {
		vLidRunon = float64(float64(Subcatch[j].runon*area) * tStep)
	}
	vOut1 = float64(float64(Voutflow+vLidRain) + vLidRunon)
	vSurfOut = float64(Subcatch[j].newRunoff * tStep)
	vOut2 = float64(vSurfOut + VlidDrain)
	hasOutflow = R.B2I((vOut2 > float64(float64(float64(2.31481e-08)*area)*tStep)))
	for p = 0; p < Nobjects[POLLUT]; p++ {
		cOut = float64(0.0)
		if (vOut1 > float64(0.0)) && (hasOutflow != 0) {
			cOut = float64(OutflowLoad[p] / vOut1)
		}
		if Subcatch[j].lidArea > float64(0.0) {
			massLoad = float64(float64(cOut*(float64(vOut1-vOut2))) * Pollut[p].mcf)
			if massLoad > float64(0.0) {
				massbal_updateLoadingTotals(BMP_REMOVAL_LOAD, p, massLoad)
			}
		}
		massLoad = float64(float64(cOut*vOut2) * Pollut[p].mcf)
		Subcatch[j].totalLoad[p] = float64(Subcatch[j].totalLoad[p] + massLoad)
		if (Subcatch[j].outNode >= 0) || (Subcatch[j].outSubcatch == j) {
			massLoad = float64(float64(cOut*vSurfOut) * Pollut[p].mcf)
			massbal_updateLoadingTotals(RUNOFF_LOAD, p, massLoad)
		}
		Subcatch[j].newQual[p] = float64(cOut / float64(28.317))
	}
	if Subcatch[j].lidArea > float64(0.0) {
		lid_addDrainLoads(j, Subcatch[j].newQual, tStep)
	}
}

// surfqual.c:298
func surfqual_getWtdWashoff(j int32, p int32, f float64) float64 {
	return float64(float64(float64((float64(float64(1.0)-f))*Subcatch[j].oldRunoff)*Subcatch[j].oldQual[p]) + float64(float64(f*Subcatch[j].newRunoff)*Subcatch[j].newQual[p]))
}

// surfqual.c:313
func findPondedLoads(j int32, tStep float64) {
	var (
		p          int32
		cPonded    float64
		wPonded    float64
		bmpRemoval float64
		vRain      float64
		wRain      float64
		wRunon     float64
		wInfil     float64
		wOutflow   float64
		fullArea   float64
		nonLidArea float64
	)
	_, _, _, _, _, _, _, _, _, _, _ = p, cPonded, wPonded, bmpRemoval, vRain, wRain, wRunon, wInfil, wOutflow, fullArea, nonLidArea
	p = 0
	cPonded = 0
	wPonded = 0
	bmpRemoval = 0
	vRain = 0
	wRain = 0
	wRunon = 0
	wInfil = 0
	wOutflow = 0
	fullArea = 0
	nonLidArea = 0
	if Subcatch[j].area == Subcatch[j].lidArea {
		return
	}
	fullArea = Subcatch[j].area
	nonLidArea = float64(fullArea - Subcatch[j].lidArea)
	vRain = float64(float64(Subcatch[j].rainfall*nonLidArea) * tStep)
	for p = 0; p < Nobjects[POLLUT]; p++ {
		wRain = float64(float64(Pollut[p].pptConcen*float64(28.317)) * vRain)
		massbal_updateLoadingTotals(DEPOSITION_LOAD, p, float64(wRain*Pollut[p].mcf))
		if Vinflow == float64(0.0) {
			massbal_updateLoadingTotals(FINAL_LOAD, p, float64(Subcatch[j].pondedQual[p]*Pollut[p].mcf))
			Subcatch[j].pondedQual[p] = float64(0.0)
		} else {
			wRunon = float64(Subcatch[j].newQual[p] * tStep)
			wPonded = float64(float64(Subcatch[j].pondedQual[p]+wRain) + wRunon)
			cPonded = float64(wPonded / Vinflow)
			wInfil = float64(cPonded * Vinfil)
			if (wInfil) <= (wPonded) {
				wInfil = (wInfil)
			} else {
				wInfil = (wPonded)
			}
			massbal_updateLoadingTotals(INFIL_LOAD, p, float64(wInfil*Pollut[p].mcf))
			wPonded = float64(wPonded - wInfil)
			wOutflow = float64(cPonded * Voutflow)
			if (wOutflow) <= (wPonded) {
				wOutflow = (wOutflow)
			} else {
				wOutflow = (wPonded)
			}
			wPonded = float64(wPonded - wOutflow)
			bmpRemoval = float64(landuse_getAvgBmpEffic(j, p) * wOutflow)
			massbal_updateLoadingTotals(BMP_REMOVAL_LOAD, p, float64(bmpRemoval*Pollut[p].mcf))
			wOutflow = float64(wOutflow - bmpRemoval)
			Subcatch[j].pondedQual[p] = float64(float64(cPonded*subcatch_getDepth(j)) * nonLidArea)
			OutflowLoad[p] = float64(OutflowLoad[p] + wOutflow)
		}
	}
}

// surfqual.c:390
func findWashoffLoads(j int32, runoff float64) {
	var (
		i    int32
		p    int32
		k    int32
		w    float64
		area float64
	)
	_, _, _, _, _ = i, p, k, w, area
	i = 0
	p = 0
	k = 0
	w = 0
	area = Subcatch[j].area
	if runoff < float64(2.31481e-08) {
		return
	}
	for i = 0; i < Nobjects[LANDUSE]; i++ {
		if Subcatch[j].landFactor[i].fraction > float64(0.0) {
			for p = 0; p < Nobjects[POLLUT]; p++ {
				OutflowLoad[p] = float64(OutflowLoad[p] + (landuse_getWashoffLoad(i, p, area, Subcatch[j].landFactor, runoff, Voutflow)))
			}
		}
	}
	for p = 0; p < Nobjects[POLLUT]; p++ {
		k = Pollut[p].coPollut
		if k >= 0 {
			w = float64(Pollut[p].coFraction * OutflowLoad[k])
			massbal_updateLoadingTotals(BUILDUP_LOAD, p, float64(w*Pollut[p].mcf))
			OutflowLoad[p] = float64(OutflowLoad[p] + w)
		}
	}
}

// surfqual.c:443
func findLidLoads(j int32, tStep float64) {
	var (
		p         int32
		useRunon  int32
		lidArea   float64
		vLidRain  float64
		wLidRain  float64
		wLidRunon float64
	)
	_, _, _, _, _, _ = p, useRunon, lidArea, vLidRain, wLidRain, wLidRunon
	p = 0
	useRunon = 0
	lidArea = 0
	vLidRain = 0
	wLidRain = 0
	wLidRunon = 0
	lidArea = Subcatch[j].lidArea
	if lidArea == float64(0.0) {
		return
	}
	vLidRain = float64(float64(Subcatch[j].rainfall*lidArea) * tStep)
	useRunon = R.B2I((lidArea == Subcatch[j].area))
	for p = 0; p < Nobjects[POLLUT]; p++ {
		wLidRain = float64(float64(Pollut[p].pptConcen*vLidRain) * float64(28.317))
		massbal_updateLoadingTotals(DEPOSITION_LOAD, p, float64(wLidRain*Pollut[p].mcf))
		if useRunon != 0 {
			wLidRunon = float64(Subcatch[j].newQual[p] * tStep)
		} else {
			wLidRunon = float64(0.0)
		}
		OutflowLoad[p] = float64(OutflowLoad[p] + (float64(wLidRain + wLidRunon)))
	}
}
