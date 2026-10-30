// Generated from EPA SWMM 5.2.4 (gage.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var OneSecond float64 = float64(1.1574074e-5)

// gage.c:58
func gage_readParams(j int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		k     int32
		err   int32
		id    R.CPtr
		fname R.CPtr
		staID R.CPtr
		x     [7]float64
	)
	_, _, _, _, _, _ = k, err, id, fname, staID, x
	k = 0
	err = 0
	id = (R.CPtr{})
	fname = R.Buf(260)
	staID = R.Buf(1025)
	x = ([7]float64{})
	if ntoks < 2 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(GAGE, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	x[0] = (-(float64(1.0)))
	x[1] = float64(1.0)
	x[2] = float64(3600.0)
	x[3] = float64(1.0)
	x[4] = float64((-693594))
	x[5] = float64((-693594))
	x[6] = float64(0.0)
	fname.S(0, 0)
	staID.S(0, 0)
	if ntoks < 5 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	k = findmatch(tok[4], GageDataWords[:])
	if k == RAIN_TSERIES {
		err = readGageSeriesFormat(tok, ntoks, x[:])
	} else if k == RAIN_FILE {
		if ntoks < 8 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		sstrncpy(fname, tok[5], 259)
		sstrncpy(staID, tok[6], 1024)
		err = readGageFileFormat(tok, ntoks, x[:])
	} else {
		return error_setInpError(ERR_KEYWORD, tok[4])
	}
	if err > 0 {
		return err
	}
	Gage[j].ID = id
	Gage[j].tSeries = R.D2I(x[0])
	Gage[j].rainType = R.D2I(x[1])
	Gage[j].rainInterval = R.D2I(x[2])
	Gage[j].snowFactor = x[3]
	Gage[j].rainUnits = R.D2I(x[6])
	if Gage[j].tSeries >= 0 {
		Gage[j].dataSource = RAIN_TSERIES
	} else {
		Gage[j].dataSource = RAIN_FILE
	}
	if Gage[j].dataSource == RAIN_FILE {
		sstrncpy(R.CP(Gage[j].fname[:]), addAbsolutePath(fname), 259)
		sstrncpy(R.CP(Gage[j].staID[:]), staID, 1024)
		Gage[j].startFileDate = x[4]
		Gage[j].endFileDate = x[5]
	}
	Gage[j].unitsFactor = float64(1.0)
	Gage[j].coGage = (-1)
	Gage[j].isUsed = 0
	return 0
}

// gage.c:133
func readGageSeriesFormat(tok []R.CPtr, ntoks int32, x []float64) int32 {
	var (
		m     int32
		ts    int32
		aTime float64
	)
	_, _, _ = m, ts, aTime
	m = 0
	ts = 0
	aTime = 0
	if ntoks < 6 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	m = findmatch(tok[1], RainTypeWords[:])
	if m < 0 {
		return error_setInpError(ERR_KEYWORD, tok[1])
	}
	x[1] = float64(m)
	if getDouble(tok[2], &x[2]) != 0 {
		x[2] = M.Floor(float64(float64(x[2]*float64(3600)) + float64(0.5)))
	} else if datetime_strToTime(tok[2], &aTime) != 0 {
		x[2] = M.Floor(float64(float64(aTime*float64(86400)) + float64(0.5)))
	} else {
		return error_setInpError(ERR_DATETIME, tok[2])
	}
	if x[2] <= float64(0.0) {
		return error_setInpError(ERR_DATETIME, tok[2])
	}
	if !(getDouble(tok[3], &x[3]) != 0) {
		return error_setInpError(ERR_DATETIME, tok[3])
	}
	ts = project_findObject(TSERIES, tok[5])
	if ts < 0 {
		return error_setInpError(ERR_NAME, tok[5])
	}
	x[0] = float64(ts)
	sstrncpy(tok[2], L0, 0)
	return 0
}

// gage.c:168
func readGageFileFormat(tok []R.CPtr, ntoks int32, x []float64) int32 {
	var (
		m     int32
		u     int32
		aDate float64
		aTime float64
	)
	_, _, _, _ = m, u, aDate, aTime
	m = 0
	u = 0
	aDate = 0
	aTime = 0
	m = findmatch(tok[1], RainTypeWords[:])
	if m < 0 {
		return error_setInpError(ERR_KEYWORD, tok[1])
	}
	x[1] = float64(m)
	if getDouble(tok[2], &x[2]) != 0 {
		x[2] = float64(x[2] * (float64(3600)))
	} else if datetime_strToTime(tok[2], &aTime) != 0 {
		x[2] = M.Floor(float64(float64(aTime*float64(86400)) + float64(0.5)))
	} else {
		return error_setInpError(ERR_DATETIME, tok[2])
	}
	if x[2] <= float64(0.0) {
		return error_setInpError(ERR_DATETIME, tok[2])
	}
	if !(getDouble(tok[3], &x[3]) != 0) {
		return error_setInpError(ERR_NUMBER, tok[3])
	}
	u = findmatch(tok[7], RainUnitsWords[:])
	if u < 0 {
		return error_setInpError(ERR_KEYWORD, tok[7])
	}
	x[6] = float64(u)
	if (ntoks > 8) && (tok[8].G(0) != 42) {
		if !(datetime_strToDate(tok[8], &aDate) != 0) {
			return error_setInpError(ERR_DATETIME, tok[8])
		}
		x[4] = float64(float32(float64(aDate)))
	}
	return 0
}

// gage.c:209
func gage_validate(j int32) {
	var (
		i            int32
		k            int32
		gageInterval int32
	)
	_, _, _ = i, k, gageInterval
	i = 0
	k = 0
	gageInterval = 0
	if Gage[j].dataSource == RAIN_TSERIES {
		if !(Gage[j].isUsed != 0) {
			return
		}
		k = Gage[j].tSeries
		for i = 0; i < j; i++ {
			if ((Gage[i].dataSource == RAIN_TSERIES) && (Gage[i].tSeries == k)) && (Gage[i].isUsed != 0) {
				Gage[j].coGage = i
				if Gage[j].rainType != Gage[i].rainType {
					report_writeErrorMsg(ERR_RAIN_GAGE_FORMAT, Gage[j].ID)
				}
				return
			}
		}
		if Tseries[k].refersTo >= 0 {
			report_writeErrorMsg(ERR_RAIN_GAGE_TSERIES, Gage[j].ID)
		}
		gageInterval = R.D2I((M.Floor(float64(float64(Tseries[k].dxMin*float64(86400)) + float64(0.5)))))
		if (gageInterval > 0) && (Gage[j].rainInterval > gageInterval) {
			report_writeErrorMsg(ERR_RAIN_GAGE_INTERVAL, Gage[j].ID)
		}
		if Gage[j].rainInterval < gageInterval {
			report_writeWarningMsg(L156, Gage[j].ID)
		}
		if Gage[j].rainInterval < WetStep {
			report_writeWarningMsg(L157, Gage[j].ID)
			WetStep = Gage[j].rainInterval
		}
	}
}

// gage.c:270
func gage_initState(j int32) {
	Gage[j].rainfall = float64(0.0)
	Gage[j].apiRainfall = (-(float64(10000000000.0)))
	Gage[j].reportRainfall = float64(0.0)
	if IgnoreRainfall != 0 {
		return
	}
	if Gage[j].dataSource == RAIN_FILE {
		Gage[j].currentFilePos = Gage[j].startFilePos
		if UnitSystem == SI {
			Gage[j].unitsFactor = float64(25.4)
		}
	}
	if getFirstRainfall(j) != 0 {
		Gage[j].endDate = datetime_addSeconds(Gage[j].startDate, float64(Gage[j].rainInterval))
		if Gage[j].startDate > StartDateTime {
			Gage[j].nextDate = Gage[j].startDate
			Gage[j].nextRainfall = Gage[j].rainfall
			Gage[j].startDate = StartDateTime
			Gage[j].endDate = Gage[j].nextDate
			Gage[j].rainfall = float64(0.0)
		} else if !(getNextRainfall(j) != 0) {
			Gage[j].nextDate = float64((-693594))
		}
	} else {
		Gage[j].startDate = float64((-693594))
	}
	initPastRain(j)
}

// gage.c:323
func gage_setState(j int32, t float64) {
	if Gage[j].isUsed == 0 {
		return
	}
	if IgnoreRainfall != 0 {
		Gage[j].rainfall = float64(0.0)
		return
	}
	if Gage[j].coGage >= 0 {
		Gage[j].rainfall = Gage[Gage[j].coGage].rainfall
		return
	}
	if Gage[j].apiRainfall != (-(float64(10000000000.0))) {
		Gage[j].rainfall = Gage[j].apiRainfall
		return
	}
	t = float64(t + OneSecond)
	for {
		if Gage[j].startDate == float64((-693594)) {
			Gage[j].rainfall = float64(0.0)
			return
		}
		if t < Gage[j].startDate {
			Gage[j].rainfall = float64(0.0)
			return
		}
		if t < Gage[j].endDate {
			return
		}
		if Gage[j].nextDate == float64((-693594)) {
			Gage[j].rainfall = float64(0.0)
			return
		}
		if t < Gage[j].nextDate {
			Gage[j].rainfall = float64(0.0)
			return
		}
		Gage[j].startDate = Gage[j].nextDate
		Gage[j].endDate = datetime_addSeconds(Gage[j].startDate, float64(Gage[j].rainInterval))
		Gage[j].rainfall = Gage[j].nextRainfall
		if !(getNextRainfall(j) != 0) {
			Gage[j].nextDate = float64((-693594))
		}
	}
}

// gage.c:406
func initPastRain(j int32) {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 0; i <= 48; i++ {
		Gage[j].pastRain[i] = float64(0.0)
	}
	Gage[j].pastInterval = 0
}

// gage.c:417
func gage_updatePastRain(j int32, tStep int32) {
	var (
		i int32
		t int32
		r float64
	)
	_, _, _ = i, t, r
	i = 0
	t = 0
	r = 0
	r = float64(Gage[j].rainfall / float64(3600.0))
	for tStep > 0 {
		t = (3600 - Gage[j].pastInterval)
		if tStep > t {
			Gage[j].pastRain[0] = float64(Gage[j].pastRain[0] + (float64(float64(t) * r)))
			for i = 48; i > 0; i-- {
				Gage[j].pastRain[i] = Gage[j].pastRain[(i - 1)]
			}
			Gage[j].pastInterval = 0
			Gage[j].pastRain[0] = float64(0.0)
			tStep -= t
		} else {
			Gage[j].pastRain[0] = float64(Gage[j].pastRain[0] + (float64(float64(tStep) * r)))
			Gage[j].pastInterval += tStep
			tStep = 0
		}
	}
}

// gage.c:468
func gage_getPastRain(j int32, n_2 int32) float64 {
	var (
		i      int32
		result float64
	)
	_, _ = i, result
	i = 0
	result = float64(0.0)
	if (n_2 < 1) || (n_2 > 48) {
		return float64(0.0)
	}
	for i = 1; i <= n_2; i++ {
		result = float64(result + Gage[j].pastRain[i])
	}
	return result
}

// gage.c:486
func gage_getNextRainDate(j int32, aDate float64) float64 {
	if Gage[j].isUsed == 0 {
		return aDate
	}
	aDate = float64(aDate + OneSecond)
	if aDate < Gage[j].startDate {
		return Gage[j].startDate
	}
	if aDate < Gage[j].endDate {
		return Gage[j].endDate
	}
	return Gage[j].nextDate
}

// gage.c:503
func gage_getPrecip(j int32, rainfall *float64, snowfall *float64) float64 {
	*rainfall = float64(0.0)
	*snowfall = float64(0.0)
	if (!(IgnoreSnowmelt != 0)) && (Temp.ta <= Snow.snotmp) {
		*snowfall = float64(float64(Gage[j].rainfall*Gage[j].snowFactor) / UCF(RAINFALL))
	} else {
		*rainfall = float64(Gage[j].rainfall / UCF(RAINFALL))
	}
	return float64((*rainfall) + (*snowfall))
}

// gage.c:524
func gage_setReportRainfall(j int32, reportDate float64) {
	var (
		result float64
	)
	_ = result
	result = 0
	if Gage[j].coGage >= 0 {
		Gage[j].reportRainfall = Gage[Gage[j].coGage].reportRainfall
		return
	}
	if Gage[j].apiRainfall != (-(float64(10000000000.0))) {
		Gage[j].reportRainfall = Gage[j].apiRainfall
		return
	}
	reportDate = float64(reportDate + OneSecond)
	if reportDate < Gage[j].endDate {
		result = Gage[j].rainfall
	} else if reportDate < Gage[j].nextDate {
		result = float64(0.0)
	} else {
		result = Gage[j].nextRainfall
	}
	Gage[j].reportRainfall = result
}

// gage.c:567
func getFirstRainfall(j int32) int32 {
	var (
		k      int32
		vFirst float32
		rFirst float64
	)
	_, _, _ = k, vFirst, rFirst
	k = 0
	vFirst = 0
	rFirst = 0
	Gage[j].startDate = float64((-693594))
	Gage[j].rainfall = float64(0.0)
	Gage[j].rainAccum = float64(0)
	if Gage[j].dataSource == RAIN_FILE {
		if (Frain.file != nil) && (Gage[j].endFilePos > Gage[j].startFilePos) {
			R.Fseek(Frain.file, Gage[j].startFilePos, 0)
			R.Fread1(&Gage[j].startDate, 8, Frain.file)
			R.Fread1(&vFirst, 4, Frain.file)
			Gage[j].currentFilePos = R.Ftell(Frain.file)
			Gage[j].rainfall = convertRainfall(j, float64(vFirst))
			return 1
		}
		return 0
	} else {
		k = Gage[j].tSeries
		if k >= 0 {
			if table_getFirstEntry(&Tseries[k], &Gage[j].startDate, &rFirst) != 0 {
				Gage[j].rainfall = convertRainfall(j, rFirst)
				return 1
			}
		}
		return 0
	}
	return 0
}

// gage.c:624
func getNextRainfall(j int32) int32 {
	var (
		k     int32
		vNext float32
		rNext float64
	)
	_, _, _ = k, vNext, rNext
	k = 0
	vNext = 0
	rNext = 0
	Gage[j].nextRainfall = float64(0.0)
	for __dw1 := true; __dw1 || (rNext == float64(0.0)); __dw1 = false {
		if Gage[j].dataSource == RAIN_FILE {
			if (Frain.file != nil) && (Gage[j].currentFilePos < Gage[j].endFilePos) {
				R.Fseek(Frain.file, Gage[j].currentFilePos, 0)
				R.Fread1(&Gage[j].nextDate, 8, Frain.file)
				R.Fread1(&vNext, 4, Frain.file)
				Gage[j].currentFilePos = R.Ftell(Frain.file)
				rNext = convertRainfall(j, float64(vNext))
			} else {
				return 0
			}
		} else {
			k = Gage[j].tSeries
			if k >= 0 {
				if !(table_getNextEntry(&Tseries[k], &Gage[j].nextDate, &rNext) != 0) {
					return 0
				}
				rNext = convertRainfall(j, rNext)
			} else {
				return 0
			}
		}
	}
	Gage[j].nextRainfall = rNext
	return 1
}

// gage.c:674
func convertRainfall(j int32, r float64) float64 {
	var (
		r1 float64
	)
	_ = r1
	r1 = 0
	switch Gage[j].rainType {
	case RAINFALL_INTENSITY:
		r1 = r
		break
	case RAINFALL_VOLUME:
		r1 = float64(float64(r/float64(Gage[j].rainInterval)) * float64(3600.0))
		break
	case CUMULATIVE_RAINFALL:
		if r < Gage[j].rainAccum {
			r1 = float64(float64(r/float64(Gage[j].rainInterval)) * float64(3600.0))
		} else {
			r1 = float64(float64((float64(r-Gage[j].rainAccum))/float64(Gage[j].rainInterval)) * float64(3600.0))
		}
		Gage[j].rainAccum = r
		break
	default:
		r1 = r
	}
	return float64(float64(r1*Gage[j].unitsFactor) * Adjust.rainFactor)
}
