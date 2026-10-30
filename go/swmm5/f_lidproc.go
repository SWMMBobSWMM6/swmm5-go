// Generated from EPA SWMM 5.2.4 (lidproc.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var theLidUnit *TLidUnit
var theLidProc *TLidProc
var lidproc_Tstep float64
var lidproc_EvapRate float64
var lidproc_MaxNativeInfil float64
var SurfaceInflow float64
var SurfaceInfil float64
var SurfaceEvap float64
var SurfaceOutflow float64
var SurfaceVolume float64
var PaveEvap float64
var PavePerc float64
var PaveVolume float64
var SoilEvap float64
var SoilPerc float64
var SoilVolume float64
var StorageInflow float64
var StorageExfil float64
var StorageEvap float64
var StorageDrain float64
var StorageVolume float64
var Xold [4]float64

// lidproc.c:178
func lidproc_initWaterBalance(lidUnit *TLidUnit, initVol float64) {
	lidUnit.waterBalance.inflow = float64(0.0)
	lidUnit.waterBalance.evap = float64(0.0)
	lidUnit.waterBalance.infil = float64(0.0)
	lidUnit.waterBalance.surfFlow = float64(0.0)
	lidUnit.waterBalance.drainFlow = float64(0.0)
	lidUnit.waterBalance.initVol = initVol
	lidUnit.waterBalance.finalVol = initVol
}

// lidproc.c:197
func lidproc_getOutflow(lidUnit *TLidUnit, lidProc *TLidProc, inflow float64, evap float64, infil float64, maxInfil float64, tStep float64, lidEvap *float64, lidInfil *float64, lidDrain *float64) float64 {
	var (
		i         int32
		x         [4]float64
		xOld      [4]float64
		xPrev     [4]float64
		xMin      [4]float64
		xMax      [4]float64
		fOld      [4]float64
		f         [4]float64
		xTol      [4]float64
		omega     float64
		fluxRates FP2
	)
	_, _, _, _, _, _, _, _, _, _, _ = i, x, xOld, xPrev, xMin, xMax, fOld, f, xTol, omega, fluxRates
	i = 0
	x = ([4]float64{})
	xOld = ([4]float64{})
	xPrev = ([4]float64{})
	xMin = ([4]float64{})
	xMax = ([4]float64{})
	fOld = ([4]float64{})
	f = ([4]float64{})
	xTol = [4]float64{float64(0.00328), float64(0.00328), float64(0.00328), float64(0.00328)}
	omega = float64(0.0)
	fluxRates = (nil)
	theLidProc = lidProc
	theLidUnit = lidUnit
	lidproc_EvapRate = evap
	lidproc_MaxNativeInfil = maxInfil
	lidproc_Tstep = tStep
	x[SURF] = theLidUnit.surfaceDepth
	x[SOIL] = theLidUnit.soilMoisture
	x[STOR] = theLidUnit.storageDepth
	x[PAVE] = theLidUnit.paveDepth
	SurfaceVolume = float64(0.0)
	PaveVolume = float64(0.0)
	SoilVolume = float64(0.0)
	StorageVolume = float64(0.0)
	SurfaceInflow = inflow
	SurfaceInfil = float64(0.0)
	SurfaceEvap = float64(0.0)
	SurfaceOutflow = float64(0.0)
	PaveEvap = float64(0.0)
	PavePerc = float64(0.0)
	SoilEvap = float64(0.0)
	SoilPerc = float64(0.0)
	StorageInflow = float64(0.0)
	StorageExfil = float64(0.0)
	StorageEvap = float64(0.0)
	StorageDrain = float64(0.0)
	for i = 0; i < 4; i++ {
		f[i] = float64(0.0)
		fOld[i] = theLidUnit.oldFluxRates[i]
		xMin[i] = float64(0.0)
		xMax[i] = float64(10000000000.0)
		Xold[i] = x[i]
	}
	if theLidProc.lidType == POROUS_PAVEMENT {
		SurfaceInfil = float64(0.0)
	} else if theLidUnit.soilInfil.Ks > float64(0.0) {
		SurfaceInfil = grnampt_getInfil(&theLidUnit.soilInfil, lidproc_Tstep, SurfaceInflow, theLidUnit.surfaceDepth, MOD_GREEN_AMPT)
	} else {
		SurfaceInfil = infil
	}
	if theLidProc.soil.thickness > float64(0.0) {
		xMin[SOIL] = theLidProc.soil.wiltPoint
		xMax[SOIL] = theLidProc.soil.porosity
	}
	if theLidProc.pavement.thickness > float64(0.0) {
		xMax[PAVE] = theLidProc.pavement.thickness
	}
	if theLidProc.storage.thickness > float64(0.0) {
		xMax[STOR] = theLidProc.storage.thickness
	}
	if theLidProc.lidType == GREEN_ROOF {
		xMax[STOR] = theLidProc.drainMat.thickness
	}
	switch theLidProc.lidType {
	case BIO_CELL, RAIN_GARDEN:
		fluxRates = biocellFluxRates
		break
	case GREEN_ROOF:
		fluxRates = greenRoofFluxRates
		break
	case INFIL_TRENCH:
		fluxRates = trenchFluxRates
		break
	case POROUS_PAVEMENT:
		fluxRates = pavementFluxRates
		break
	case RAIN_BARREL:
		fluxRates = barrelFluxRates
		break
	case ROOF_DISCON:
		fluxRates = roofFluxRates
		break
	case VEG_SWALE:
		fluxRates = swaleFluxRates
		omega = float64(0.5)
		break
	default:
		return float64(0.0)
	}
	i = modpuls_solve(4, x[:], xOld[:], xPrev[:], xMin[:], xMax[:], xTol[:], fOld[:], f[:], tStep, omega, fluxRates)
	if (theLidProc.surface.canOverflow != 0) || (theLidUnit.fullWidth == float64(0.0)) {
		SurfaceOutflow = float64(SurfaceOutflow + (getSurfaceOverflowRate(&x[SURF])))
	}
	theLidUnit.surfaceDepth = x[SURF]
	theLidUnit.paveDepth = x[PAVE]
	theLidUnit.soilMoisture = x[SOIL]
	theLidUnit.storageDepth = x[STOR]
	for i = 0; i < 4; i++ {
		theLidUnit.oldFluxRates[i] = f[i]
	}
	*lidEvap = float64(float64(float64(SurfaceEvap+PaveEvap)+SoilEvap) + StorageEvap)
	*lidInfil = StorageExfil
	*lidDrain = StorageDrain
	return SurfaceOutflow
}

// lidproc.c:360
func lidproc_saveResults(lidUnit *TLidUnit, ucfRainfall float64, ucfRainDepth float64) {
	var (
		ucf         float64
		totalEvap   float64
		totalVolume float64
		rptVars     [12]float64
		isDry       int32
		timeStamp   R.CPtr
		elapsedHrs  float64
	)
	_, _, _, _, _, _, _ = ucf, totalEvap, totalVolume, rptVars, isDry, timeStamp, elapsedHrs
	ucf = 0
	totalEvap = 0
	totalVolume = 0
	rptVars = ([12]float64{})
	isDry = 0
	timeStamp = R.Buf(22)
	elapsedHrs = 0
	totalEvap = float64(float64(float64(SurfaceEvap+PaveEvap)+SoilEvap) + StorageEvap)
	totalVolume = float64(float64(float64(SurfaceVolume+PaveVolume)+SoilVolume) + StorageVolume)
	updateWaterBalance(theLidUnit, SurfaceInflow, totalEvap, StorageExfil, SurfaceOutflow, StorageDrain, totalVolume)
	if ((((SurfaceInflow < float64(2.3e-08)) && (SurfaceOutflow < float64(2.3e-08))) && (StorageDrain < float64(2.3e-08))) && (StorageExfil < float64(2.3e-08))) && (totalEvap < float64(2.3e-08)) {
		isDry = 1
	}
	if !(isDry != 0) {
		HasWetLids = 1
	}
	if lidUnit.rptFile != nil {
		ucf = ucfRainfall
		rptVars[SURF_INFLOW] = float64(SurfaceInflow * ucf)
		rptVars[TOTAL_EVAP] = float64(totalEvap * ucf)
		rptVars[SURF_INFIL] = float64(SurfaceInfil * ucf)
		rptVars[PAVE_PERC] = float64(PavePerc * ucf)
		rptVars[SOIL_PERC] = float64(SoilPerc * ucf)
		rptVars[STOR_EXFIL] = float64(StorageExfil * ucf)
		rptVars[SURF_OUTFLOW] = float64(SurfaceOutflow * ucf)
		rptVars[STOR_DRAIN] = float64(StorageDrain * ucf)
		ucf = ucfRainDepth
		rptVars[SURF_DEPTH] = float64(theLidUnit.surfaceDepth * ucf)
		rptVars[PAVE_DEPTH] = float64(theLidUnit.paveDepth * ucf)
		rptVars[SOIL_MOIST] = theLidUnit.soilMoisture
		rptVars[STOR_DEPTH] = float64(theLidUnit.storageDepth * ucf)
		if (!(isDry != 0)) && (theLidUnit.rptFile.wasDry > 1) {
			R.Fprintf(theLidUnit.rptFile.file, L169, R.CP(theLidUnit.rptFile.results[:]))
		}
		elapsedHrs = float64(float64(NewRunoffTime/float64(1000.0)) / float64(3600.0))
		datetime_getTimeStamp(1, getDateTime(NewRunoffTime), 21, timeStamp)
		R.Snprintf(R.CP(theLidUnit.rptFile.results[:]), 256, L304, timeStamp, elapsedHrs, rptVars[0], rptVars[1], rptVars[2], rptVars[3], rptVars[4], rptVars[5], rptVars[6], rptVars[7], rptVars[8], rptVars[9], rptVars[10], rptVars[11])
		if isDry != 0 {
			if theLidUnit.rptFile.wasDry == 0 {
				R.Fprintf(theLidUnit.rptFile.file, L169, R.CP(theLidUnit.rptFile.results[:]))
			}
			theLidUnit.rptFile.wasDry++
		} else {
			R.Fprintf(theLidUnit.rptFile.file, L169, R.CP(theLidUnit.rptFile.results[:]))
			theLidUnit.rptFile.wasDry = 0
		}
	}
}

// lidproc.c:469
func roofFluxRates(x []float64, f []float64) {
	var (
		surfaceDepth float64
	)
	_ = surfaceDepth
	surfaceDepth = x[SURF]
	lidproc_getEvapRates(surfaceDepth, float64(0.0), float64(0.0), float64(0.0), float64(1.0))
	SurfaceVolume = surfaceDepth
	SurfaceInfil = float64(0.0)
	if theLidProc.surface.alpha > float64(0.0) {
		SurfaceOutflow = getSurfaceOutflowRate(surfaceDepth)
	} else {
		getSurfaceOverflowRate(&surfaceDepth)
	}
	if (float64(theLidProc.drain.coeff / UCF(RAINFALL))) <= (SurfaceOutflow) {
		StorageDrain = (float64(theLidProc.drain.coeff / UCF(RAINFALL)))
	} else {
		StorageDrain = (SurfaceOutflow)
	}
	SurfaceOutflow = float64(SurfaceOutflow - StorageDrain)
	f[SURF] = (float64(float64(float64(SurfaceInflow-SurfaceEvap)-StorageDrain) - SurfaceOutflow))
}

// lidproc.c:491
func greenRoofFluxRates(x []float64, f []float64) {
	var (
		surfaceDepth     float64
		soilTheta        float64
		storageDepth     float64
		availVolume      float64
		maxRate          float64
		soilThickness    float64
		storageThickness float64
		soilPorosity     float64
		storageVoidFrac  float64
		soilFieldCap     float64
		soilWiltPoint    float64
	)
	_, _, _, _, _, _, _, _, _, _, _ = surfaceDepth, soilTheta, storageDepth, availVolume, maxRate, soilThickness, storageThickness, soilPorosity, storageVoidFrac, soilFieldCap, soilWiltPoint
	surfaceDepth = 0
	soilTheta = 0
	storageDepth = 0
	availVolume = 0
	maxRate = 0
	soilThickness = theLidProc.soil.thickness
	storageThickness = theLidProc.storage.thickness
	soilPorosity = theLidProc.soil.porosity
	storageVoidFrac = theLidProc.storage.voidFrac
	soilFieldCap = theLidProc.soil.fieldCap
	soilWiltPoint = theLidProc.soil.wiltPoint
	surfaceDepth = x[SURF]
	soilTheta = x[SOIL]
	storageDepth = x[STOR]
	SurfaceVolume = float64(surfaceDepth * theLidProc.surface.voidFrac)
	SoilVolume = float64(soilTheta * soilThickness)
	StorageVolume = float64(storageDepth * storageVoidFrac)
	availVolume = float64(SoilVolume - float64(soilWiltPoint*soilThickness))
	lidproc_getEvapRates(SurfaceVolume, float64(0.0), availVolume, StorageVolume, float64(1.0))
	if soilTheta >= soilPorosity {
		StorageEvap = float64(0.0)
	}
	SoilPerc = getSoilPercRate(soilTheta)
	availVolume = float64((float64(soilTheta - soilFieldCap)) * soilThickness)
	maxRate = float64(float64((R.If[float64](((availVolume) >= (float64(0.0))), (availVolume), (float64(0.0))))/lidproc_Tstep) - SoilEvap)
	if (SoilPerc) <= (maxRate) {
		SoilPerc = (SoilPerc)
	} else {
		SoilPerc = (maxRate)
	}
	if (SoilPerc) >= (float64(0.0)) {
		SoilPerc = (SoilPerc)
	} else {
		SoilPerc = (float64(0.0))
	}
	StorageExfil = float64(0.0)
	StorageDrain = getDrainMatOutflow(storageDepth)
	if (soilTheta >= soilPorosity) && (storageDepth >= storageThickness) {
		if (SoilPerc) <= (StorageDrain) {
			maxRate = (SoilPerc)
		} else {
			maxRate = (StorageDrain)
		}
		SoilPerc = maxRate
		StorageDrain = maxRate
		if (SurfaceInfil) <= (maxRate) {
			SurfaceInfil = (SurfaceInfil)
		} else {
			SurfaceInfil = (maxRate)
		}
	} else {
		maxRate = float64(float64(float64(storageDepth*storageVoidFrac)/lidproc_Tstep) - StorageEvap)
		if storageDepth >= storageThickness {
			maxRate = float64(maxRate + SoilPerc)
		}
		if (maxRate) >= (float64(0.0)) {
			maxRate = (maxRate)
		} else {
			maxRate = (float64(0.0))
		}
		if (StorageDrain) <= (maxRate) {
			StorageDrain = (StorageDrain)
		} else {
			StorageDrain = (maxRate)
		}
		maxRate = float64(float64(float64(float64((float64(storageThickness-storageDepth))*storageVoidFrac)/lidproc_Tstep)+StorageDrain) + StorageEvap)
		if (SoilPerc) <= (maxRate) {
			SoilPerc = (SoilPerc)
		} else {
			SoilPerc = (maxRate)
		}
		maxRate = float64(float64(float64(float64((float64(soilPorosity-soilTheta))*soilThickness)/lidproc_Tstep)+SoilPerc) + SoilEvap)
		if (SurfaceInfil) <= (maxRate) {
			SurfaceInfil = (SurfaceInfil)
		} else {
			SurfaceInfil = (maxRate)
		}
	}
	SurfaceOutflow = getSurfaceOutflowRate(surfaceDepth)
	f[SURF] = float64((float64(float64(float64(SurfaceInflow-SurfaceEvap)-SurfaceInfil) - SurfaceOutflow)) / theLidProc.surface.voidFrac)
	f[SOIL] = float64((float64(float64(SurfaceInfil-SoilEvap) - SoilPerc)) / theLidProc.soil.thickness)
	f[STOR] = float64((float64(float64(SoilPerc-StorageEvap) - StorageDrain)) / theLidProc.storage.voidFrac)
}

// lidproc.c:589
func biocellFluxRates(x []float64, f []float64) {
	var (
		surfaceDepth     float64
		soilTheta        float64
		storageDepth     float64
		availVolume      float64
		maxRate          float64
		soilThickness    float64
		soilPorosity     float64
		soilFieldCap     float64
		soilWiltPoint    float64
		storageThickness float64
		storageVoidFrac  float64
	)
	_, _, _, _, _, _, _, _, _, _, _ = surfaceDepth, soilTheta, storageDepth, availVolume, maxRate, soilThickness, soilPorosity, soilFieldCap, soilWiltPoint, storageThickness, storageVoidFrac
	surfaceDepth = 0
	soilTheta = 0
	storageDepth = 0
	availVolume = 0
	maxRate = 0
	soilThickness = theLidProc.soil.thickness
	soilPorosity = theLidProc.soil.porosity
	soilFieldCap = theLidProc.soil.fieldCap
	soilWiltPoint = theLidProc.soil.wiltPoint
	storageThickness = theLidProc.storage.thickness
	storageVoidFrac = theLidProc.storage.voidFrac
	surfaceDepth = x[SURF]
	soilTheta = x[SOIL]
	storageDepth = x[STOR]
	SurfaceVolume = float64(surfaceDepth * theLidProc.surface.voidFrac)
	SoilVolume = float64(soilTheta * soilThickness)
	StorageVolume = float64(storageDepth * storageVoidFrac)
	availVolume = float64(SoilVolume - float64(soilWiltPoint*soilThickness))
	lidproc_getEvapRates(SurfaceVolume, float64(0.0), availVolume, StorageVolume, float64(1.0))
	if soilTheta >= soilPorosity {
		StorageEvap = float64(0.0)
	}
	SoilPerc = getSoilPercRate(soilTheta)
	availVolume = float64((float64(soilTheta - soilFieldCap)) * soilThickness)
	maxRate = float64(float64((R.If[float64](((availVolume) >= (float64(0.0))), (availVolume), (float64(0.0))))/lidproc_Tstep) - SoilEvap)
	if (SoilPerc) <= (maxRate) {
		SoilPerc = (SoilPerc)
	} else {
		SoilPerc = (maxRate)
	}
	if (SoilPerc) >= (float64(0.0)) {
		SoilPerc = (SoilPerc)
	} else {
		SoilPerc = (float64(0.0))
	}
	StorageExfil = getStorageExfilRate()
	StorageDrain = float64(0.0)
	if theLidProc.drain.coeff > float64(0.0) {
		StorageDrain = getStorageDrainRate(storageDepth, soilTheta, float64(0.0), surfaceDepth)
	}
	if storageThickness == float64(0.0) {
		StorageEvap = float64(0.0)
		if (SoilPerc) <= (StorageExfil) {
			maxRate = (SoilPerc)
		} else {
			maxRate = (StorageExfil)
		}
		SoilPerc = maxRate
		StorageExfil = maxRate
		maxRate = float64(float64(float64(float64((float64(soilPorosity-soilTheta))*soilThickness)/lidproc_Tstep)+SoilPerc) + SoilEvap)
		if (SurfaceInfil) <= (maxRate) {
			SurfaceInfil = (SurfaceInfil)
		} else {
			SurfaceInfil = (maxRate)
		}
	} else {
		if (soilTheta >= soilPorosity) && (storageDepth >= storageThickness) {
			maxRate = float64(StorageExfil + StorageDrain)
			if SoilPerc < maxRate {
				maxRate = SoilPerc
				if maxRate > StorageExfil {
					StorageDrain = float64(maxRate - StorageExfil)
				} else {
					StorageExfil = maxRate
					StorageDrain = float64(0.0)
				}
			} else {
				SoilPerc = maxRate
			}
			if (SurfaceInfil) <= (maxRate) {
				SurfaceInfil = (SurfaceInfil)
			} else {
				SurfaceInfil = (maxRate)
			}
		} else {
			maxRate = float64(float64(SoilPerc-StorageEvap) + float64(float64(storageDepth*storageVoidFrac)/lidproc_Tstep))
			if (StorageExfil) <= (maxRate) {
				StorageExfil = (StorageExfil)
			} else {
				StorageExfil = (maxRate)
			}
			if (StorageExfil) >= (float64(0.0)) {
				StorageExfil = (StorageExfil)
			} else {
				StorageExfil = (float64(0.0))
			}
			if StorageDrain > float64(0.0) {
				maxRate = float64((-StorageExfil) - StorageEvap)
				if storageDepth >= storageThickness {
					maxRate = float64(maxRate + SoilPerc)
				}
				if theLidProc.drain.offset <= storageDepth {
					maxRate = float64(maxRate + (float64(float64((float64(storageDepth-theLidProc.drain.offset))*storageVoidFrac) / lidproc_Tstep)))
				}
				if (maxRate) >= (float64(0.0)) {
					maxRate = (maxRate)
				} else {
					maxRate = (float64(0.0))
				}
				if (StorageDrain) <= (maxRate) {
					StorageDrain = (StorageDrain)
				} else {
					StorageDrain = (maxRate)
				}
			}
			maxRate = float64(float64(float64(StorageExfil+StorageDrain)+StorageEvap) + float64(float64((float64(storageThickness-storageDepth))*storageVoidFrac)/lidproc_Tstep))
			if (SoilPerc) <= (maxRate) {
				SoilPerc = (SoilPerc)
			} else {
				SoilPerc = (maxRate)
			}
			maxRate = float64(float64(float64(float64((float64(soilPorosity-soilTheta))*soilThickness)/lidproc_Tstep)+SoilPerc) + SoilEvap)
			if (SurfaceInfil) <= (maxRate) {
				SurfaceInfil = (SurfaceInfil)
			} else {
				SurfaceInfil = (maxRate)
			}
		}
	}
	SurfaceOutflow = getSurfaceOutflowRate(surfaceDepth)
	f[SURF] = float64((float64(float64(float64(SurfaceInflow-SurfaceEvap)-SurfaceInfil) - SurfaceOutflow)) / theLidProc.surface.voidFrac)
	f[SOIL] = float64((float64(float64(SurfaceInfil-SoilEvap) - SoilPerc)) / theLidProc.soil.thickness)
	if storageThickness == float64(0.0) {
		f[STOR] = float64(0.0)
	} else {
		f[STOR] = float64((float64(float64(float64(SoilPerc-StorageEvap)-StorageExfil) - StorageDrain)) / theLidProc.storage.voidFrac)
	}
}

// lidproc.c:735
func trenchFluxRates(x []float64, f []float64) {
	var (
		surfaceDepth     float64
		storageDepth     float64
		availVolume      float64
		maxRate          float64
		storageThickness float64
		storageVoidFrac  float64
	)
	_, _, _, _, _, _ = surfaceDepth, storageDepth, availVolume, maxRate, storageThickness, storageVoidFrac
	surfaceDepth = 0
	storageDepth = 0
	availVolume = 0
	maxRate = 0
	storageThickness = theLidProc.storage.thickness
	storageVoidFrac = theLidProc.storage.voidFrac
	surfaceDepth = x[SURF]
	storageDepth = x[STOR]
	SurfaceVolume = float64(surfaceDepth * theLidProc.surface.voidFrac)
	SoilVolume = float64(0.0)
	StorageVolume = float64(storageDepth * storageVoidFrac)
	availVolume = float64((float64(storageThickness - storageDepth)) * storageVoidFrac)
	lidproc_getEvapRates(SurfaceVolume, float64(0.0), float64(0.0), StorageVolume, float64(1.0))
	if surfaceDepth > float64(0.0) {
		StorageEvap = float64(0.0)
	}
	StorageInflow = float64(SurfaceInflow + float64(SurfaceVolume/lidproc_Tstep))
	StorageExfil = getStorageExfilRate()
	StorageDrain = float64(0.0)
	if theLidProc.drain.coeff > float64(0.0) {
		StorageDrain = getStorageDrainRate(storageDepth, float64(0.0), float64(0.0), surfaceDepth)
	}
	maxRate = float64(float64(StorageInflow-StorageEvap) + float64(float64(storageDepth*storageVoidFrac)/lidproc_Tstep))
	if (StorageExfil) <= (maxRate) {
		StorageExfil = (StorageExfil)
	} else {
		StorageExfil = (maxRate)
	}
	if (StorageExfil) >= (float64(0.0)) {
		StorageExfil = (StorageExfil)
	} else {
		StorageExfil = (float64(0.0))
	}
	if StorageDrain > float64(0.0) {
		maxRate = float64((-StorageExfil) - StorageEvap)
		if storageDepth >= storageThickness {
			maxRate = float64(maxRate + StorageInflow)
		}
		if theLidProc.drain.offset <= storageDepth {
			maxRate = float64(maxRate + (float64(float64((float64(storageDepth-theLidProc.drain.offset))*storageVoidFrac) / lidproc_Tstep)))
		}
		if (maxRate) >= (float64(0.0)) {
			maxRate = (maxRate)
		} else {
			maxRate = (float64(0.0))
		}
		if (StorageDrain) <= (maxRate) {
			StorageDrain = (StorageDrain)
		} else {
			StorageDrain = (maxRate)
		}
	}
	maxRate = float64(float64(float64(float64(float64((float64(storageThickness-storageDepth))*storageVoidFrac)/lidproc_Tstep)+StorageExfil)+StorageEvap) + StorageDrain)
	if (StorageInflow) <= (maxRate) {
		StorageInflow = (StorageInflow)
	} else {
		StorageInflow = (maxRate)
	}
	SurfaceInfil = StorageInflow
	SurfaceOutflow = getSurfaceOutflowRate(surfaceDepth)
	f[SURF] = float64((float64(float64(float64(SurfaceInflow-SurfaceEvap)-StorageInflow) - SurfaceOutflow)) / theLidProc.surface.voidFrac)
	f[STOR] = float64((float64(float64(float64(StorageInflow-StorageEvap)-StorageExfil) - StorageDrain)) / theLidProc.storage.voidFrac)
	f[SOIL] = float64(0.0)
}

// lidproc.c:823
func pavementFluxRates(x []float64, f []float64) {
	var (
		surfaceDepth     float64
		paveDepth        float64
		soilTheta        float64
		storageDepth     float64
		pervFrac         float64
		storageInflow    float64
		availVolume      float64
		maxRate          float64
		paveVoidFrac     float64
		paveThickness    float64
		soilThickness    float64
		soilPorosity     float64
		soilFieldCap     float64
		soilWiltPoint    float64
		storageThickness float64
		storageVoidFrac  float64
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = surfaceDepth, paveDepth, soilTheta, storageDepth, pervFrac, storageInflow, availVolume, maxRate, paveVoidFrac, paveThickness, soilThickness, soilPorosity
	_, _, _, _ = soilFieldCap, soilWiltPoint, storageThickness, storageVoidFrac
	surfaceDepth = 0
	paveDepth = 0
	soilTheta = 0
	storageDepth = 0
	pervFrac = (float64(float64(1.0) - theLidProc.pavement.impervFrac))
	storageInflow = 0
	availVolume = 0
	maxRate = 0
	paveVoidFrac = float64(theLidProc.pavement.voidFrac * pervFrac)
	paveThickness = theLidProc.pavement.thickness
	soilThickness = theLidProc.soil.thickness
	soilPorosity = theLidProc.soil.porosity
	soilFieldCap = theLidProc.soil.fieldCap
	soilWiltPoint = theLidProc.soil.wiltPoint
	storageThickness = theLidProc.storage.thickness
	storageVoidFrac = theLidProc.storage.voidFrac
	surfaceDepth = x[SURF]
	paveDepth = x[PAVE]
	soilTheta = x[SOIL]
	storageDepth = x[STOR]
	SurfaceVolume = float64(surfaceDepth * theLidProc.surface.voidFrac)
	PaveVolume = float64(paveDepth * paveVoidFrac)
	SoilVolume = float64(soilTheta * soilThickness)
	StorageVolume = float64(storageDepth * storageVoidFrac)
	availVolume = float64(SoilVolume - float64(soilWiltPoint*soilThickness))
	lidproc_getEvapRates(SurfaceVolume, PaveVolume, availVolume, StorageVolume, pervFrac)
	if (paveDepth >= paveThickness) || ((soilThickness > float64(0.0)) && (soilTheta >= soilPorosity)) {
		StorageEvap = float64(0.0)
	}
	SurfaceInfil = float64(SurfaceInflow + (float64(SurfaceVolume / lidproc_Tstep)))
	PavePerc = float64(getPavementPermRate() * pervFrac)
	if (SurfaceInfil) <= (PavePerc) {
		SurfaceInfil = (SurfaceInfil)
	} else {
		SurfaceInfil = (PavePerc)
	}
	maxRate = float64(float64(float64(PaveVolume/lidproc_Tstep)+SurfaceInfil) - PaveEvap)
	if (maxRate) >= (float64(0.0)) {
		maxRate = (maxRate)
	} else {
		maxRate = (float64(0.0))
	}
	if (PavePerc) <= (maxRate) {
		PavePerc = (PavePerc)
	} else {
		PavePerc = (maxRate)
	}
	if soilThickness > float64(0.0) {
		SoilPerc = getSoilPercRate(soilTheta)
		availVolume = float64((float64(soilTheta - soilFieldCap)) * soilThickness)
		maxRate = float64(float64((R.If[float64](((availVolume) >= (float64(0.0))), (availVolume), (float64(0.0))))/lidproc_Tstep) - SoilEvap)
		if (SoilPerc) <= (maxRate) {
			SoilPerc = (SoilPerc)
		} else {
			SoilPerc = (maxRate)
		}
		if (SoilPerc) >= (float64(0.0)) {
			SoilPerc = (SoilPerc)
		} else {
			SoilPerc = (float64(0.0))
		}
	} else {
		SoilPerc = PavePerc
	}
	StorageExfil = getStorageExfilRate()
	StorageDrain = float64(0.0)
	if theLidProc.drain.coeff > float64(0.0) {
		StorageDrain = getStorageDrainRate(storageDepth, soilTheta, paveDepth, surfaceDepth)
	}
	if ((soilThickness == float64(0.0)) && (storageDepth >= storageThickness)) && (paveDepth >= paveThickness) {
		maxRate = float64(float64(StorageEvap+StorageDrain) + StorageExfil)
		if PavePerc > maxRate {
			PavePerc = maxRate
		} else {
			if (StorageExfil) <= (PavePerc) {
				StorageExfil = (StorageExfil)
			} else {
				StorageExfil = (PavePerc)
			}
			StorageDrain = float64(PavePerc - StorageExfil)
		}
		SoilPerc = PavePerc
		if (SurfaceInfil) <= (PavePerc) {
			SurfaceInfil = (SurfaceInfil)
		} else {
			SurfaceInfil = (PavePerc)
		}
	} else if (((soilThickness > float64(0)) && (storageDepth >= storageThickness)) && (soilTheta >= soilPorosity)) && (paveDepth >= paveThickness) {
		maxRate = float64(StorageExfil + StorageDrain)
		if SoilPerc < maxRate {
			maxRate = SoilPerc
		} else {
			if (maxRate) <= (PavePerc) {
				maxRate = (maxRate)
			} else {
				maxRate = (PavePerc)
			}
		}
		if maxRate > StorageExfil {
			StorageDrain = float64(maxRate - StorageExfil)
		} else {
			StorageExfil = maxRate
			StorageDrain = float64(0.0)
		}
		SoilPerc = maxRate
		PavePerc = maxRate
		if (SurfaceInfil) <= (PavePerc) {
			SurfaceInfil = (SurfaceInfil)
		} else {
			SurfaceInfil = (PavePerc)
		}
	} else if ((soilThickness > float64(0.0)) && (storageDepth >= storageThickness)) && (soilTheta >= soilPorosity) {
		maxRate = float64(StorageDrain + StorageExfil)
		if SoilPerc > maxRate {
			SoilPerc = maxRate
		} else {
			if (StorageExfil) <= (SoilPerc) {
				StorageExfil = (StorageExfil)
			} else {
				StorageExfil = (SoilPerc)
			}
			StorageDrain = float64(SoilPerc - StorageExfil)
		}
		if (PavePerc) <= (SoilPerc) {
			PavePerc = (PavePerc)
		} else {
			PavePerc = (SoilPerc)
		}
		availVolume = float64((float64(paveThickness - paveDepth)) * paveVoidFrac)
		maxRate = float64(float64(float64(availVolume/lidproc_Tstep)+PavePerc) + PaveEvap)
		if (SurfaceInfil) <= (maxRate) {
			SurfaceInfil = (SurfaceInfil)
		} else {
			SurfaceInfil = (maxRate)
		}
	} else if ((soilThickness > float64(0.0)) && (paveDepth >= paveThickness)) && (soilTheta >= soilPorosity) {
		if (PavePerc) <= (SoilPerc) {
			PavePerc = (PavePerc)
		} else {
			PavePerc = (SoilPerc)
		}
		SoilPerc = PavePerc
		if (SurfaceInfil) <= (PavePerc) {
			SurfaceInfil = (SurfaceInfil)
		} else {
			SurfaceInfil = (PavePerc)
		}
		if (float64(float64(float64(StorageVolume/lidproc_Tstep)+SoilPerc) - StorageEvap)) >= (float64(0.0)) {
			maxRate = (float64(float64(float64(StorageVolume/lidproc_Tstep)+SoilPerc) - StorageEvap))
		} else {
			maxRate = (float64(0.0))
		}
		if (StorageExfil) <= (maxRate) {
			StorageExfil = (StorageExfil)
		} else {
			StorageExfil = (maxRate)
		}
	} else {
		maxRate = float64(float64(SoilPerc-StorageEvap) + float64(StorageVolume/lidproc_Tstep))
		if (float64(0.0)) >= (maxRate) {
			maxRate = (float64(0.0))
		} else {
			maxRate = (maxRate)
		}
		if (StorageExfil) <= (maxRate) {
			StorageExfil = (StorageExfil)
		} else {
			StorageExfil = (maxRate)
		}
		if StorageDrain > float64(0.0) {
			maxRate = float64((-StorageExfil) - StorageEvap)
			if storageDepth >= storageThickness {
				maxRate = float64(maxRate + SoilPerc)
			}
			if theLidProc.drain.offset <= storageDepth {
				maxRate = float64(maxRate + (float64(float64((float64(storageDepth-theLidProc.drain.offset))*storageVoidFrac) / lidproc_Tstep)))
			}
			if (maxRate) >= (float64(0.0)) {
				maxRate = (maxRate)
			} else {
				maxRate = (float64(0.0))
			}
			if (StorageDrain) <= (maxRate) {
				StorageDrain = (StorageDrain)
			} else {
				StorageDrain = (maxRate)
			}
		}
		availVolume = float64((float64(storageThickness - storageDepth)) * storageVoidFrac)
		maxRate = float64(float64(float64(float64(availVolume/lidproc_Tstep)+StorageEvap)+StorageDrain) + StorageExfil)
		if (maxRate) >= (float64(0.0)) {
			maxRate = (maxRate)
		} else {
			maxRate = (float64(0.0))
		}
		if soilThickness > float64(0.0) {
			if (SoilPerc) <= (maxRate) {
				SoilPerc = (SoilPerc)
			} else {
				SoilPerc = (maxRate)
			}
			maxRate = float64(float64(float64((float64(soilPorosity-soilTheta))*soilThickness)/lidproc_Tstep) + SoilPerc)
		}
		if (PavePerc) <= (maxRate) {
			PavePerc = (PavePerc)
		} else {
			PavePerc = (maxRate)
		}
		availVolume = float64((float64(paveThickness - paveDepth)) * paveVoidFrac)
		maxRate = float64(float64(float64(availVolume/lidproc_Tstep)+PavePerc) + PaveEvap)
		if (SurfaceInfil) <= (maxRate) {
			SurfaceInfil = (SurfaceInfil)
		} else {
			SurfaceInfil = (maxRate)
		}
	}
	SurfaceOutflow = getSurfaceOutflowRate(surfaceDepth)
	f[SURF] = float64(float64(float64(SurfaceInflow-SurfaceEvap)-SurfaceInfil) - SurfaceOutflow)
	f[PAVE] = float64((float64(float64(SurfaceInfil-PaveEvap) - PavePerc)) / paveVoidFrac)
	if theLidProc.soil.thickness > float64(0.0) {
		f[SOIL] = float64((float64(float64(PavePerc-SoilEvap) - SoilPerc)) / soilThickness)
		storageInflow = SoilPerc
	} else {
		f[SOIL] = float64(0.0)
		storageInflow = PavePerc
		SoilPerc = float64(0.0)
	}
	f[STOR] = float64((float64(float64(float64(storageInflow-StorageEvap)-StorageExfil) - StorageDrain)) / storageVoidFrac)
}

// lidproc.c:1061
func swaleFluxRates(x []float64, f []float64) {
	var (
		depth      float64
		topWidth   float64
		botWidth   float64
		length     float64
		surfInflow float64
		surfWidth  float64
		surfArea   float64
		flowArea   float64
		lidArea    float64
		hydRadius  float64
		slope      float64
		volume     float64
		dVdT       float64
		dStore     float64
		xDepth     float64
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = depth, topWidth, botWidth, length, surfInflow, surfWidth, surfArea, flowArea, lidArea, hydRadius, slope, volume
	_, _, _ = dVdT, dStore, xDepth
	depth = 0
	topWidth = 0
	botWidth = 0
	length = 0
	surfInflow = 0
	surfWidth = 0
	surfArea = 0
	flowArea = 0
	lidArea = 0
	hydRadius = 0
	slope = 0
	volume = 0
	dVdT = 0
	dStore = 0
	xDepth = 0
	depth = x[SURF]
	if (depth) <= (theLidProc.surface.thickness) {
		depth = (depth)
	} else {
		depth = (theLidProc.surface.thickness)
	}
	dStore = float64(0.0)
	slope = theLidProc.surface.sideSlope
	topWidth = theLidUnit.fullWidth
	if (topWidth) >= (float64(0.5)) {
		topWidth = (topWidth)
	} else {
		topWidth = (float64(0.5))
	}
	botWidth = float64(topWidth - float64(float64(float64(2.0)*slope)*theLidProc.surface.thickness))
	if botWidth < float64(0.5) {
		botWidth = float64(0.5)
		slope = float64(float64(float64(0.5)*(float64(topWidth-float64(0.5)))) / theLidProc.surface.thickness)
	}
	lidArea = theLidUnit.area
	length = float64(lidArea / topWidth)
	surfWidth = float64(botWidth + float64(float64(float64(2.0)*slope)*depth))
	surfArea = float64(length * surfWidth)
	flowArea = float64((float64(depth * (float64(botWidth + float64(slope*depth))))) * theLidProc.surface.voidFrac)
	volume = float64(length * flowArea)
	surfInflow = float64(SurfaceInflow * lidArea)
	SurfaceEvap = float64(lidproc_EvapRate * surfArea)
	if (SurfaceEvap) <= (float64(volume / lidproc_Tstep)) {
		SurfaceEvap = (SurfaceEvap)
	} else {
		SurfaceEvap = (float64(volume / lidproc_Tstep))
	}
	StorageExfil = float64(SurfaceInfil * surfArea)
	xDepth = float64(depth - dStore)
	if xDepth <= float64(1e-10) {
		SurfaceOutflow = float64(0.0)
	} else {
		flowArea = float64(flowArea - (float64((float64(dStore * (float64(botWidth + float64(slope*dStore))))) * theLidProc.surface.voidFrac)))
		if flowArea < float64(1e-10) {
			SurfaceOutflow = float64(0.0)
		} else {
			botWidth = float64(botWidth + float64(float64(float64(2.0)*dStore)*slope))
			hydRadius = float64(botWidth + float64(float64(float64(2.0)*xDepth)*M.Sqrt(float64(float64(1.0)+float64(slope*slope)))))
			hydRadius = float64(flowArea / hydRadius)
			SurfaceOutflow = float64(float64(theLidProc.surface.alpha*flowArea) * M.Pow(hydRadius, float64(float64(2.0)/float64(3.0))))
		}
	}
	dVdT = float64(float64(float64(surfInflow-SurfaceEvap)-StorageExfil) - SurfaceOutflow)
	if (depth == theLidProc.surface.thickness) && (dVdT > float64(0.0)) {
		SurfaceOutflow = float64(SurfaceOutflow + dVdT)
		dVdT = float64(0.0)
	}
	SurfaceEvap = float64(SurfaceEvap / lidArea)
	StorageExfil = float64(StorageExfil / lidArea)
	SurfaceOutflow = float64(SurfaceOutflow / lidArea)
	f[SURF] = float64(dVdT / surfArea)
	f[SOIL] = float64(0.0)
	f[STOR] = float64(0.0)
	SurfaceVolume = float64(volume / lidArea)
	SoilVolume = float64(0.0)
	StorageVolume = float64(0.0)
}

// lidproc.c:1176
func barrelFluxRates(x []float64, f []float64) {
	var (
		storageDepth float64
		head         float64
		maxValue     float64
	)
	_, _, _ = storageDepth, head, maxValue
	storageDepth = x[STOR]
	head = 0
	maxValue = 0
	SurfaceVolume = float64(0.0)
	SoilVolume = float64(0.0)
	StorageVolume = storageDepth
	SurfaceInfil = float64(0.0)
	SurfaceOutflow = float64(0.0)
	StorageDrain = float64(0.0)
	if (theLidProc.drain.delay == float64(0.0)) || (theLidUnit.dryTime >= theLidProc.drain.delay) {
		head = float64(storageDepth - theLidProc.drain.offset)
		if head > float64(0.0) {
			StorageDrain = getStorageDrainRate(storageDepth, float64(0.0), float64(0.0), float64(0.0))
			maxValue = (float64(head / lidproc_Tstep))
			if (StorageDrain) <= (maxValue) {
				StorageDrain = (StorageDrain)
			} else {
				StorageDrain = (maxValue)
			}
		}
	}
	StorageInflow = SurfaceInflow
	maxValue = float64(float64((float64(theLidProc.storage.thickness-storageDepth))/lidproc_Tstep) + StorageDrain)
	if (StorageInflow) <= (maxValue) {
		StorageInflow = (StorageInflow)
	} else {
		StorageInflow = (maxValue)
	}
	SurfaceInfil = StorageInflow
	f[SURF] = float64(SurfaceInflow - StorageInflow)
	f[STOR] = float64(StorageInflow - StorageDrain)
	f[SOIL] = float64(0.0)
}

// lidproc.c:1226
func getSurfaceOutflowRate(depth float64) float64 {
	var (
		delta   float64
		outflow float64
	)
	_, _ = delta, outflow
	delta = 0
	outflow = 0
	delta = float64(depth - theLidProc.surface.thickness)
	if delta < float64(0.0) {
		return float64(0.0)
	}
	outflow = float64(float64(float64(theLidProc.surface.alpha*M.Pow(delta, float64(float64(5.0)/float64(3.0))))*theLidUnit.fullWidth) / theLidUnit.area)
	if (outflow) <= (float64(delta / lidproc_Tstep)) {
		outflow = (outflow)
	} else {
		outflow = (float64(delta / lidproc_Tstep))
	}
	return outflow
}

// lidproc.c:1251
func getPavementPermRate() float64 {
	var (
		permReduction float64
		clogFactor    float64
		regenDays     float64
	)
	_, _, _ = permReduction, clogFactor, regenDays
	permReduction = float64(0.0)
	clogFactor = theLidProc.pavement.clogFactor
	regenDays = theLidProc.pavement.regenDays
	if clogFactor > float64(0.0) {
		if regenDays > float64(0.0) {
			if float64(float64(OldRunoffTime/float64(1000.0))/float64(86400)) >= theLidUnit.nextRegenDay {
				theLidUnit.volTreated = float64(theLidUnit.volTreated * (float64(float64(1.0) - theLidProc.pavement.regenDegree)))
				theLidUnit.nextRegenDay = float64(theLidUnit.nextRegenDay + regenDays)
			}
		}
		permReduction = float64(theLidUnit.volTreated / clogFactor)
		if (permReduction) <= (float64(1.0)) {
			permReduction = (permReduction)
		} else {
			permReduction = (float64(1.0))
		}
	}
	return float64(theLidProc.pavement.kSat * (float64(float64(1.0) - permReduction)))
}

// lidproc.c:1293
func getSoilPercRate(theta float64) float64 {
	var (
		delta float64
	)
	_ = delta
	delta = 0
	if theta <= theLidProc.soil.fieldCap {
		return float64(0.0)
	}
	delta = float64(theLidProc.soil.porosity - theta)
	return float64(theLidProc.soil.kSat * M.Exp(float64((-delta)*theLidProc.soil.kSlope)))
}

// lidproc.c:1313
func getStorageExfilRate() float64 {
	var (
		infil      float64
		clogFactor float64
	)
	_, _ = infil, clogFactor
	infil = float64(0.0)
	clogFactor = float64(0.0)
	if theLidProc.storage.kSat == float64(0.0) {
		return float64(0.0)
	}
	if lidproc_MaxNativeInfil == float64(0.0) {
		return float64(0.0)
	}
	clogFactor = theLidProc.storage.clogFactor
	if clogFactor > float64(0.0) {
		clogFactor = float64(theLidUnit.waterBalance.inflow / clogFactor)
		if (clogFactor) <= (float64(1.0)) {
			clogFactor = (clogFactor)
		} else {
			clogFactor = (float64(1.0))
		}
	}
	infil = float64(theLidProc.storage.kSat * (float64(float64(1.0) - clogFactor)))
	if (infil) <= (lidproc_MaxNativeInfil) {
		return (infil)
	}
	return (lidproc_MaxNativeInfil)
}

// lidproc.c:1344
func getStorageDrainRate(storageDepth float64, soilTheta float64, paveDepth float64, surfaceDepth float64) float64 {
	var (
		curve            int32
		head             float64
		outflow          float64
		paveThickness    float64
		soilThickness    float64
		soilPorosity     float64
		soilFieldCap     float64
		storageThickness float64
	)
	_, _, _, _, _, _, _, _ = curve, head, outflow, paveThickness, soilThickness, soilPorosity, soilFieldCap, storageThickness
	curve = theLidProc.drain.qCurve
	head = storageDepth
	outflow = float64(0.0)
	paveThickness = theLidProc.pavement.thickness
	soilThickness = theLidProc.soil.thickness
	soilPorosity = theLidProc.soil.porosity
	soilFieldCap = theLidProc.soil.fieldCap
	storageThickness = theLidProc.storage.thickness
	if storageDepth >= storageThickness {
		if soilThickness > float64(0.0) {
			if soilTheta > soilFieldCap {
				head = float64(head + (float64(float64((float64(soilTheta-soilFieldCap))/(float64(soilPorosity-soilFieldCap))) * soilThickness)))
				if soilTheta >= soilPorosity {
					if paveThickness > float64(0.0) {
						head = float64(head + paveDepth)
						if paveDepth >= paveThickness {
							head = float64(head + surfaceDepth)
						}
					} else {
						head = float64(head + surfaceDepth)
					}
				}
			}
		} else if paveThickness > float64(0.0) {
			head = float64(head + paveDepth)
			if paveDepth >= paveThickness {
				head = float64(head + surfaceDepth)
			}
		}
	}
	if (theLidUnit.oldDrainFlow == float64(0.0)) && (head <= theLidProc.drain.hOpen) {
		return float64(0.0)
	}
	if (theLidUnit.oldDrainFlow > float64(0.0)) && (head <= theLidProc.drain.hClose) {
		return float64(0.0)
	}
	head = float64(head - theLidProc.drain.offset)
	if head > float64(1e-10) {
		head = float64(head * (UCF(RAINDEPTH)))
		outflow = float64(theLidProc.drain.coeff * M.Pow(head, theLidProc.drain.expon))
		if curve >= 0 {
			outflow = float64(outflow * (table_lookup(&Curve[curve], head)))
		}
		outflow = float64(outflow / (UCF(RAINFALL)))
	}
	return outflow
}

// lidproc.c:1436
func getDrainMatOutflow(depth float64) float64 {
	var (
		result float64
	)
	_ = result
	result = SoilPerc
	if theLidProc.drainMat.alpha > float64(0.0) {
		result = float64(float64(float64(float64(theLidProc.drainMat.alpha*M.Pow(depth, float64(float64(5.0)/float64(3.0))))*theLidUnit.fullWidth)/theLidUnit.area) * theLidProc.drainMat.voidFrac)
	}
	return result
}

// lidproc.c:1458
func lidproc_getEvapRates(surfaceVol float64, paveVol float64, soilVol float64, storageVol float64, pervFrac float64) {
	var (
		availEvap float64
	)
	_ = availEvap
	availEvap = 0
	availEvap = lidproc_EvapRate
	if (availEvap) <= (float64(surfaceVol / lidproc_Tstep)) {
		SurfaceEvap = (availEvap)
	} else {
		SurfaceEvap = (float64(surfaceVol / lidproc_Tstep))
	}
	if (float64(0.0)) >= (SurfaceEvap) {
		SurfaceEvap = (float64(0.0))
	} else {
		SurfaceEvap = (SurfaceEvap)
	}
	if (float64(0.0)) >= (float64(availEvap - SurfaceEvap)) {
		availEvap = (float64(0.0))
	} else {
		availEvap = (float64(availEvap - SurfaceEvap))
	}
	availEvap = float64(availEvap * pervFrac)
	if SurfaceInfil > float64(0.0) {
		PaveEvap = float64(0.0)
		SoilEvap = float64(0.0)
		StorageEvap = float64(0.0)
	} else {
		if (availEvap) <= (float64(paveVol / lidproc_Tstep)) {
			PaveEvap = (availEvap)
		} else {
			PaveEvap = (float64(paveVol / lidproc_Tstep))
		}
		if (float64(0.0)) >= (float64(availEvap - PaveEvap)) {
			availEvap = (float64(0.0))
		} else {
			availEvap = (float64(availEvap - PaveEvap))
		}
		if (availEvap) <= (float64(soilVol / lidproc_Tstep)) {
			SoilEvap = (availEvap)
		} else {
			SoilEvap = (float64(soilVol / lidproc_Tstep))
		}
		if (float64(0.0)) >= (float64(availEvap - SoilEvap)) {
			availEvap = (float64(0.0))
		} else {
			availEvap = (float64(availEvap - SoilEvap))
		}
		if (availEvap) <= (float64(storageVol / lidproc_Tstep)) {
			StorageEvap = (availEvap)
		} else {
			StorageEvap = (float64(storageVol / lidproc_Tstep))
		}
	}
}

// lidproc.c:1503
func getSurfaceOverflowRate(surfaceDepth *float64) float64 {
	var (
		delta float64
	)
	_ = delta
	delta = float64((*surfaceDepth) - theLidProc.surface.thickness)
	if delta <= float64(0.0) {
		return float64(0.0)
	}
	*surfaceDepth = theLidProc.surface.thickness
	return float64(float64(delta*theLidProc.surface.voidFrac) / lidproc_Tstep)
}

// lidproc.c:1518
func updateWaterBalance(lidUnit *TLidUnit, inflow float64, evap float64, infil float64, surfFlow float64, drainFlow float64, storage float64) {
	lidUnit.volTreated = float64(lidUnit.volTreated + (float64(inflow * lidproc_Tstep)))
	lidUnit.waterBalance.inflow = float64(lidUnit.waterBalance.inflow + (float64(inflow * lidproc_Tstep)))
	lidUnit.waterBalance.evap = float64(lidUnit.waterBalance.evap + (float64(evap * lidproc_Tstep)))
	lidUnit.waterBalance.infil = float64(lidUnit.waterBalance.infil + (float64(infil * lidproc_Tstep)))
	lidUnit.waterBalance.surfFlow = float64(lidUnit.waterBalance.surfFlow + (float64(surfFlow * lidproc_Tstep)))
	lidUnit.waterBalance.drainFlow = float64(lidUnit.waterBalance.drainFlow + (float64(drainFlow * lidproc_Tstep)))
	lidUnit.waterBalance.finalVol = storage
}

// lidproc.c:1544
func modpuls_solve(n_2 int32, x []float64, xOld []float64, xPrev []float64, xMin []float64, xMax []float64, xTol []float64, qOld []float64, q []float64, dt float64, omega float64, derivs FP2) int32 {
	var (
		i        int32
		canStop  int32
		steps    int32
		maxSteps int32
	)
	_, _, _, _ = i, canStop, steps, maxSteps
	i = 0
	canStop = 0
	steps = 1
	maxSteps = 20
	for i = 0; i < n_2; i++ {
		xOld[i] = x[i]
		xPrev[i] = x[i]
	}
	for steps < maxSteps {
		canStop = 1
		derivs(x, q)
		for i = 0; i < n_2; i++ {
			x[i] = float64(xOld[i] + float64((float64(float64(omega*qOld[i])+float64((float64(float64(1.0)-omega))*q[i])))*dt))
			if (x[i]) <= (xMax[i]) {
				x[i] = (x[i])
			} else {
				x[i] = (xMax[i])
			}
			if (x[i]) >= (xMin[i]) {
				x[i] = (x[i])
			} else {
				x[i] = (xMin[i])
			}
			if (omega > float64(0.0)) && (M.Fabs(float64(x[i]-xPrev[i])) > xTol[i]) {
				canStop = 0
			}
			xPrev[i] = x[i]
		}
		if canStop != 0 {
			return steps
		}
		steps++
	}
	return 0
}
