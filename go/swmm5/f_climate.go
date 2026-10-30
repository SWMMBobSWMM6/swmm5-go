// Generated from EPA SWMM 5.2.4 (climate.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var MAXCLIMATEVARS int32 = 4
var MAXDAYSPERMONTH int32 = 32
var ClimateVarWords = [6]R.CPtr{L22, L23, L7, L24, L25, (R.CPtr{})}
var TempUnitsWords = [4]R.CPtr{L653, L654, L655, (R.CPtr{})}
var Tmin float64
var Tmax float64
var Trng float64
var Trng1 float64
var Tave float64
var Hrsr float64
var Hrss float64
var Hrday float64
var Dhrdy float64
var Dydif float64
var LastDay float64
var Tma TMovAve
var NextEvapDate float64
var NextEvapRate float64
var FileFormat int32
var FileYear int32
var FileMonth int32
var FileDay int32
var FileLastDay int32
var FileElapsedDays int32
var FileValue [4]float64
var FileData [4][32]float64
var FileLine = R.Buf(1025)
var FileFieldPos [4]int32
var FileDateFieldPos int32
var FileWindType int32
var FileTempUnits int32

// climate.c:153
func climate_readParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		i     int32
		j     int32
		k     int32
		x     [6]float64
		y_2   float64
		aDate float64
		fname R.CPtr
	)
	_, _, _, _, _, _, _ = i, j, k, x, y_2, aDate, fname
	i = 0
	j = 0
	k = 0
	x = ([6]float64{})
	y_2 = 0
	aDate = 0
	fname = R.Buf(260)
	k = findmatch(tok[0], TempKeyWords[:])
	if k < 0 {
		return error_setInpError(ERR_KEYWORD, tok[0])
	}
	switch k {
	case 0:
		if ntoks < 2 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		i = project_findObject(TSERIES, tok[1])
		if i < 0 {
			return error_setInpError(ERR_NAME, tok[1])
		}
		Temp.dataSource = TSERIES_TEMP
		Temp.tSeries = i
		Tseries[i].refersTo = TSERIES_TEMP
		break
	case 1:
		if ntoks < 2 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		Temp.dataSource = FILE_TEMP
		Fclimate.mode = USE_FILE
		sstrncpy(fname, tok[1], 259)
		sstrncpy(R.CP(Fclimate.name[:]), addAbsolutePath(fname), 259)
		Temp.fileStartDate = float64((-693594))
		if ntoks > 2 {
			if tok[2].G(0) != 42 {
				if !(datetime_strToDate(tok[2], &aDate) != 0) {
					return error_setInpError(ERR_DATETIME, tok[2])
				}
				Temp.fileStartDate = aDate
			}
		}
		FileTempUnits = DEG_F
		if UnitSystem == SI {
			FileTempUnits = DEG_C
		}
		if ntoks > 3 {
			i = findmatch(tok[3], TempUnitsWords[:])
			if i < 0 {
				return error_setInpError(ERR_KEYWORD, tok[3])
			}
			FileTempUnits = i
		}
		break
	case 2:
		if strcomp(tok[1], L1) != 0 {
			Wind.type_ = FILE_WIND
		} else {
			if ntoks < 14 {
				return error_setInpError(ERR_ITEMS, L0)
			}
			Wind.type_ = MONTHLY_WIND
			for i = 0; i < 12; i++ {
				if !(getDouble(tok[(i+2)], &y_2) != 0) {
					return error_setInpError(ERR_NUMBER, tok[(i+2)])
				}
				Wind.aws[i] = y_2
			}
		}
		break
	case 3:
		if ntoks < 7 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		for i = 1; i < 7; i++ {
			if !(getDouble(tok[i], &x[(i-1)]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[i])
			}
		}
		if UnitSystem == SI {
			x[0] = float64(float64(float64(float64(9.0)/float64(5.0))*x[0]) + float64(32.0))
		}
		Snow.snotmp = x[0]
		Snow.tipm = x[1]
		Snow.rnm = x[2]
		Temp.elev = float64(x[3] / UCF(LENGTH))
		Temp.anglat = x[4]
		Temp.dtlong = float64(x[5] / float64(60.0))
		break
	case 4:
		if ntoks < 12 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		if match(tok[1], L2) != 0 {
			i = 0
		} else if match(tok[1], L3) != 0 {
			i = 1
		} else {
			return error_setInpError(ERR_KEYWORD, tok[1])
		}
		for j = 0; j < 10; j++ {
			if ((!(getDouble(tok[(j+2)], &y_2) != 0)) || (y_2 < float64(0.0))) || (y_2 > float64(1.0)) {
				return error_setInpError(ERR_NUMBER, tok[(j+2)])
			}
			Snow.adc[i][j] = y_2
		}
		break
	}
	return 0
}

// climate.c:285
func climate_readEvapParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		i int32
		k int32
		x float64
	)
	_, _, _ = i, k, x
	i = 0
	k = 0
	x = 0
	k = findmatch(tok[0], EvapTypeWords[:])
	if k < 0 {
		return error_setInpError(ERR_KEYWORD, tok[0])
	}
	if k == RECOVERY {
		if ntoks < 2 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		i = project_findObject(TIMEPATTERN, tok[1])
		if i < 0 {
			return error_setInpError(ERR_NAME, tok[1])
		}
		Evap.recoveryPattern = i
		return 0
	}
	if k == DRYONLY {
		if ntoks < 2 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		if strcomp(tok[1], L4) != 0 {
			Evap.dryOnly = 0
		} else if strcomp(tok[1], L5) != 0 {
			Evap.dryOnly = 1
		} else {
			return error_setInpError(ERR_KEYWORD, tok[1])
		}
		return 0
	}
	Evap.type_ = k
	if (k != TEMPERATURE_EVAP) && (ntoks < 2) {
		return error_setInpError(ERR_ITEMS, L0)
	}
	switch k {
	case CONSTANT_EVAP:
		if !(getDouble(tok[1], &x) != 0) {
			return error_setInpError(ERR_NUMBER, tok[1])
		}
		for i = 0; i < 12; i++ {
			Evap.monthlyEvap[i] = x
		}
		break
	case MONTHLY_EVAP:
		if ntoks < 13 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		for i = 0; i < 12; i++ {
			if !(getDouble(tok[(i+1)], &Evap.monthlyEvap[i]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[(i+1)])
			}
		}
		break
	case TIMESERIES_EVAP:
		i = project_findObject(TSERIES, tok[1])
		if i < 0 {
			return error_setInpError(ERR_NAME, tok[1])
		}
		Evap.tSeries = i
		Tseries[i].refersTo = TIMESERIES_EVAP
		break
	case FILE_EVAP:
		if ntoks > 1 {
			if ntoks < 13 {
				return error_setInpError(ERR_ITEMS, L0)
			}
			for i = 0; i < 12; i++ {
				if !(getDouble(tok[(i+1)], &Evap.panCoeff[i]) != 0) {
					return error_setInpError(ERR_NUMBER, tok[(i+1)])
				}
			}
		}
		break
	}
	return 0
}

// climate.c:377
func climate_readAdjustments(tok []R.CPtr, ntoks int32) int32 {
	var (
		i int32
		j int32
	)
	_, _ = i, j
	i = 0
	j = 0
	if ntoks == 1 {
		return 0
	}
	if match(tok[0], L6) != 0 {
		if ntoks < 13 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		for i = 1; i < 13; i++ {
			if !(getDouble(tok[i], &Adjust.temp[(i-1)]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[i])
			}
		}
		return 0
	}
	if match(tok[0], L7) != 0 {
		if ntoks < 13 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		for i = 1; i < 13; i++ {
			if !(getDouble(tok[i], &Adjust.evap[(i-1)]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[i])
			}
		}
		return 0
	}
	if match(tok[0], L8) != 0 {
		if ntoks < 13 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		for i = 1; i < 13; i++ {
			if !(getDouble(tok[i], &Adjust.rain[(i-1)]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[i])
			}
		}
		return 0
	}
	if match(tok[0], L9) != 0 {
		if ntoks < 13 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		for i = 1; i < 13; i++ {
			if !(getDouble(tok[i], &Adjust.hydcon[(i-1)]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[i])
			}
			if Adjust.hydcon[(i-1)] <= float64(0.0) {
				Adjust.hydcon[(i - 1)] = float64(1.0)
			}
		}
		return 0
	}
	if match(tok[0], L10) != 0 {
		if ntoks < 3 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		i = project_findObject(SUBCATCH, tok[1])
		if i < 0 {
			return error_setInpError(ERR_NAME, tok[1])
		}
		j = project_findObject(TIMEPATTERN, tok[2])
		if j < 0 {
			return error_setInpError(ERR_NAME, tok[2])
		}
		Subcatch[i].nPervPattern = j
		return 0
	}
	if match(tok[0], L11) != 0 {
		if ntoks < 3 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		i = project_findObject(SUBCATCH, tok[1])
		if i < 0 {
			return error_setInpError(ERR_NAME, tok[1])
		}
		j = project_findObject(TIMEPATTERN, tok[2])
		if j < 0 {
			return error_setInpError(ERR_NAME, tok[2])
		}
		Subcatch[i].dStorePattern = j
		return 0
	}
	if match(tok[0], L12) != 0 {
		if ntoks < 3 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		i = project_findObject(SUBCATCH, tok[1])
		if i < 0 {
			return error_setInpError(ERR_NAME, tok[1])
		}
		j = project_findObject(TIMEPATTERN, tok[2])
		if j < 0 {
			return error_setInpError(ERR_NAME, tok[2])
		}
		Subcatch[i].infilPattern = j
		return 0
	}
	return error_setInpError(ERR_KEYWORD, tok[0])
}

// climate.c:480
func climate_validate() {
	var (
		i   int32
		a_2 float64
		z   float64
		pa  float64
	)
	_, _, _, _ = i, a_2, z, pa
	i = 0
	a_2 = 0
	z = 0
	pa = 0
	if ((Wind.type_ == FILE_WIND) || (Evap.type_ == FILE_EVAP)) || (Evap.type_ == TEMPERATURE_EVAP) {
		if Fclimate.mode == NO_FILE {
			report_writeErrorMsg(ERR_NO_CLIMATE_FILE, L0)
		}
	}
	if Fclimate.mode == USE_FILE {
		climate_openFile()
	}
	if (((Snow.tipm < float64(0.0)) || (Snow.tipm > float64(1.0))) || (Snow.rnm < float64(0.0))) || (Snow.rnm > float64(1.0)) {
		report_writeErrorMsg(ERR_SNOWMELT_PARAMS, L0)
	}
	a_2 = Temp.anglat
	if (a_2 <= (-(float64(89.99)))) || (a_2 >= float64(89.99)) {
		report_writeErrorMsg(ERR_SNOWMELT_PARAMS, L0)
	} else {
		Temp.tanAnglat = M.Tan(float64(float64(a_2*float64(3.141592654)) / float64(180.0)))
	}
	z = float64(Temp.elev / float64(1000.0))
	if z <= float64(0.0) {
		pa = float64(29.9)
	} else {
		pa = float64(float64(float64(29.9)-float64(float64(1.02)*z)) + float64(float64(0.0032)*M.Pow(z, float64(2.4))))
	}
	Temp.gamma = float64(float64(0.000359) * pa)
	for i = 0; i < 12; i++ {
		if UnitSystem == SI {
			Adjust.temp[i] = float64(Adjust.temp[i] * (float64(float64(9.0) / float64(5.0))))
		}
		Adjust.evap[i] = float64(Adjust.evap[i] / (UCF(EVAPRATE)))
	}
}

// climate.c:531
func climate_openFile() {
	var (
		i   int32
		m   int32
		y_2 int32
	)
	_, _, _ = i, m, y_2
	i = 0
	m = 0
	y_2 = 0
	if (R.Set(&Fclimate.file, R.Fopen(R.CP(Fclimate.name[:]), L13))) == nil {
		report_writeErrorMsg(ERR_CLIMATE_FILE_OPEN, R.CP(Fclimate.name[:]))
		return
	}
	FileValue[TMIN] = Temp.ta
	FileValue[TMAX] = Temp.ta
	FileValue[EVAP] = float64(0.0)
	FileValue[WIND] = float64(0.0)
	FileFormat = getFileFormat()
	if FileFormat == UNKNOWN_FORMAT {
		report_writeErrorMsg(ERR_CLIMATE_FILE_READ, R.CP(Fclimate.name[:]))
		return
	}
	R.Rewind(Fclimate.file)
	sstrncpy(FileLine, L0, 0)
	if Temp.fileStartDate == float64((-693594)) {
		datetime_decodeDate(StartDate, &FileYear, &FileMonth, &FileDay)
	} else {
		datetime_decodeDate(Temp.fileStartDate, &FileYear, &FileMonth, &FileDay)
	}
	for !(R.Feof(Fclimate.file) != 0) {
		sstrncpy(FileLine, L0, 0)
		readFileLine(&y_2, &m)
		if (y_2 == FileYear) && (m == FileMonth) {
			break
		}
	}
	if R.Feof(Fclimate.file) != 0 {
		report_writeErrorMsg(ERR_CLIMATE_END_OF_FILE, R.CP(Fclimate.name[:]))
		return
	}
	if !(ErrorCode != 0) {
		FileElapsedDays = 0
		FileLastDay = datetime_daysPerMonth(FileYear, FileMonth)
		readFileValues()
		for i = TMIN; i <= WIND; i++ {
			if FileData[i][FileDay] == (-(float64(10000000000.0))) {
				continue
			}
			FileValue[i] = FileData[i][FileDay]
		}
	}
}

// climate.c:598
func climate_initState() {
	LastDay = float64((-693594))
	Temp.tmax = (-(float64(10000000000.0)))
	Snow.removed = float64(0.0)
	NextEvapDate = StartDate
	NextEvapRate = float64(0.0)
	if (Evap.type_ == TIMESERIES_EVAP) && (Evap.tSeries >= 0) {
		table_getFirstEntry(&Tseries[Evap.tSeries], &NextEvapDate, &NextEvapRate)
		if NextEvapDate < StartDate {
			setNextEvapDate(StartDate)
		}
		Evap.rate = float64(NextEvapRate / UCF(EVAPRATE))
		setNextEvapDate(NextEvapDate)
	}
	if Evap.type_ == TEMPERATURE_EVAP {
		Tma.maxCount = int32((56 / 8))
		Tma.count = 0
		Tma.front = 0
		Tma.tAve = float64(0.0)
		Tma.tRng = float64(0.0)
	}
}

// climate.c:641
func climate_setState(theDate float64) {
	if Fclimate.mode == USE_FILE {
		updateFileValues(theDate)
	}
	if Temp.dataSource != NO_TEMP {
		setTemp(theDate)
	}
	setEvap(theDate)
	setWind(theDate)
	Adjust.rainFactor = Adjust.rain[(datetime_monthOfYear(theDate) - 1)]
	Adjust.hydconFactor = Adjust.hydcon[(datetime_monthOfYear(theDate) - 1)]
	setNextEvapDate(theDate)
}

// climate.c:659
func climate_getNextEvapDate() float64 {
	return NextEvapDate
}

// climate.c:671
func setNextEvapDate(theDate float64) {
	var (
		yr  int32
		mon int32
		day int32
		k   int32
		d   float64
		e   float64
	)
	_, _, _, _, _, _ = yr, mon, day, k, d, e
	yr = 0
	mon = 0
	day = 0
	k = 0
	d = 0
	e = 0
	if NextEvapDate > theDate {
		return
	}
	switch Evap.type_ {
	case CONSTANT_EVAP:
		NextEvapDate = float64(theDate + float64(365.0))
		break
	case MONTHLY_EVAP:
		datetime_decodeDate(theDate, &yr, &mon, &day)
		if mon == 12 {
			mon = 1
			yr++
		} else {
			mon++
		}
		NextEvapDate = datetime_encodeDate(yr, mon, 1)
		break
	case TIMESERIES_EVAP:
		k = Evap.tSeries
		if k >= 0 {
			NextEvapDate = float64(theDate + float64(365.0))
			for (table_getNextEntry(&Tseries[k], &d, &e) != 0) && (d <= EndDateTime) {
				if d >= theDate {
					NextEvapDate = d
					NextEvapRate = e
					break
				}
			}
		}
		break
	case FILE_EVAP:
		NextEvapDate = float64(M.Floor(theDate) + float64(1.0))
		break
	default:
		NextEvapDate = float64(theDate + float64(365.0))
	}
}

// climate.c:734
func updateFileValues(theDate float64) {
	var (
		i         int32
		deltaDays int32
	)
	_, _ = i, deltaDays
	i = 0
	deltaDays = 0
	deltaDays = R.D2I((float64(M.Floor(theDate) - M.Floor(StartDateTime))))
	if deltaDays > FileElapsedDays {
		FileElapsedDays++
		FileDay++
		if FileDay > FileLastDay {
			FileMonth++
			if FileMonth > 12 {
				FileMonth = 1
				FileYear++
			}
			readFileValues()
			FileDay = 1
			FileLastDay = datetime_daysPerMonth(FileYear, FileMonth)
		}
		for i = TMIN; i <= WIND; i++ {
			if FileData[i][FileDay] == (-(float64(10000000000.0))) {
				continue
			}
			FileValue[i] = FileData[i][FileDay]
		}
	}
}

// climate.c:782
func setTemp(theDate float64) {
	var (
		j      int32
		k      int32
		mon    int32
		day    int32
		theDay float64
		hour   float64
		tmp    float64
	)
	_, _, _, _, _, _, _ = j, k, mon, day, theDay, hour, tmp
	j = 0
	k = 0
	mon = 0
	day = 0
	theDay = 0
	hour = 0
	tmp = 0
	mon = datetime_monthOfYear(theDate)
	theDay = M.Floor(theDate)
	if theDay > LastDay {
		day = datetime_dayOfYear(theDate)
		if Temp.dataSource == FILE_TEMP {
			Tmin = float64(FileValue[TMIN] + Adjust.temp[(mon-1)])
			Tmax = float64(FileValue[TMAX] + Adjust.temp[(mon-1)])
			if Tmin > Tmax {
				tmp = Tmin
				Tmin = Tmax
				Tmax = tmp
			}
			updateTempTimes(day)
			if Evap.type_ == TEMPERATURE_EVAP {
				updateTempMoveAve(Tmin, Tmax)
				FileValue[EVAP] = getTempEvap(day, Tma.tAve, Tma.tRng)
			}
		}
		Snow.season = M.Sin(float64(float64(0.0172615) * (float64(float64(day) - float64(81.0)))))
		for j = 0; j < Nobjects[SNOWMELT]; j++ {
			snow_setMeltCoeffs(j, Snow.season)
		}
		LastDay = theDate
	}
	if Temp.dataSource == FILE_TEMP {
		hour = float64((float64(theDate - theDay)) * float64(24.0))
		if hour < Hrsr {
			Temp.ta = float64(Tmin + float64(float64(Trng1/float64(2.0))*M.Sin(float64(float64(float64(3.141592654)/Dydif)*(float64(Hrsr-hour))))))
		} else if (hour >= Hrsr) && (hour <= Hrss) {
			Temp.ta = float64(Tave + float64(Trng*M.Sin(float64(float64(float64(3.141592654)/Dhrdy)*(float64(Hrday-hour))))))
		} else {
			Temp.ta = float64(Tmax - float64(Trng*M.Sin(float64(float64(float64(3.141592654)/Dydif)*(float64(hour-Hrss))))))
		}
	}
	if Temp.dataSource == TSERIES_TEMP {
		k = Temp.tSeries
		if k >= 0 {
			Temp.ta = table_tseriesLookup(&Tseries[k], theDate, 1)
			if UnitSystem == SI {
				Temp.ta = float64(float64((float64(float64(9.0)/float64(5.0)))*Temp.ta) + float64(32.0))
			}
			Temp.ta = float64(Temp.ta + (Adjust.temp[(mon - 1)]))
		}
	}
	Temp.ea = float64(float64(8.1175e6) * M.Exp(float64((-(float64(7701.544)))/(float64(Temp.ta+float64(405.0265))))))
}

// climate.c:872
func setEvap(theDate float64) {
	var (
		k   int32
		mon int32
	)
	_, _ = k, mon
	k = 0
	mon = datetime_monthOfYear(theDate)
	switch Evap.type_ {
	case CONSTANT_EVAP:
		Evap.rate = float64(Evap.monthlyEvap[0] / UCF(EVAPRATE))
		break
	case MONTHLY_EVAP:
		Evap.rate = float64(Evap.monthlyEvap[(mon-1)] / UCF(EVAPRATE))
		break
	case TIMESERIES_EVAP:
		if theDate >= NextEvapDate {
			Evap.rate = float64(NextEvapRate / UCF(EVAPRATE))
		}
		break
	case FILE_EVAP:
		Evap.rate = float64(FileValue[EVAP] / UCF(EVAPRATE))
		Evap.rate = float64(Evap.rate * (Evap.panCoeff[(mon - 1)]))
		break
	case TEMPERATURE_EVAP:
		Evap.rate = float64(FileValue[EVAP] / UCF(EVAPRATE))
		break
	default:
		Evap.rate = float64(0.0)
	}
	Evap.rate = float64(Evap.rate + (Adjust.evap[(mon - 1)]))
	Evap.recoveryFactor = float64(1.0)
	k = Evap.recoveryPattern
	if (k >= 0) && (Pattern[k].type_ == MONTHLY_PATTERN) {
		Evap.recoveryFactor = Pattern[k].factor[(mon - 1)]
	}
}

// climate.c:923
func setWind(theDate float64) {
	var (
		yr  int32
		mon int32
		day int32
	)
	_, _, _ = yr, mon, day
	yr = 0
	mon = 0
	day = 0
	switch Wind.type_ {
	case MONTHLY_WIND:
		datetime_decodeDate(theDate, &yr, &mon, &day)
		Wind.ws = float64(Wind.aws[(mon-1)] / UCF(WINDSPEED))
		break
	case FILE_WIND:
		Wind.ws = FileValue[WIND]
		break
	default:
		Wind.ws = float64(0.0)
	}
}

// climate.c:949
func updateTempTimes(day int32) {
	var (
		decl  float64
		hrang float64
		arg   float64
	)
	_, _, _ = decl, hrang, arg
	decl = 0
	hrang = 0
	arg = 0
	decl = float64(float64(0.40928) * M.Cos(float64(float64(0.017202)*(float64(float64(172.0)-float64(day))))))
	arg = float64((-(M.Tan(decl))) * Temp.tanAnglat)
	if arg <= (-(float64(1.0))) {
		arg = float64(3.141592654)
	} else if arg >= float64(1.0) {
		arg = float64(0.0)
	} else {
		arg = M.Acos(arg)
	}
	hrang = float64(float64(3.8197) * arg)
	Hrsr = float64(float64(float64(12.0)-hrang) + Temp.dtlong)
	Hrss = float64(float64(float64(float64(12.0)+hrang)+Temp.dtlong) - float64(3.0))
	Dhrdy = float64(Hrsr - Hrss)
	Dydif = float64(float64(float64(24.0)+Hrsr) - Hrss)
	Hrday = float64((float64(Hrsr + Hrss)) / float64(2.0))
	Tave = float64((float64(Tmin + Tmax)) / float64(2.0))
	Trng = float64((float64(Tmax - Tmin)) / float64(2.0))
	if Temp.tmax == (-(float64(10000000000.0))) {
		Trng1 = float64(Tmax - Tmin)
	} else {
		Trng1 = float64(Temp.tmax - Tmin)
	}
	Temp.tmax = Tmax
}

// climate.c:981
func getTempEvap(day int32, tave float64, trng float64) float64 {
	var (
		a_2   float64
		ta    float64
		tr    float64
		lamda float64
		dr    float64
		phi   float64
		del   float64
		omega float64
		ra    float64
		e     float64
	)
	_, _, _, _, _, _, _, _, _, _ = a_2, ta, tr, lamda, dr, phi, del, omega, ra, e
	a_2 = float64(float64(float64(2.0)*float64(3.141592654)) / float64(365.0))
	ta = float64(float64((float64(tave-float64(32.0)))*float64(5.0)) / float64(9.0))
	tr = float64(float64(trng*float64(5.0)) / float64(9.0))
	lamda = float64(float64(2.50) - float64(float64(0.002361)*ta))
	dr = float64(float64(1.0) + float64(float64(0.033)*M.Cos(float64(a_2*float64(day)))))
	phi = float64(float64(float64(Temp.anglat*float64(2.0))*float64(3.141592654)) / float64(360.0))
	del = float64(float64(0.4093) * M.Sin(float64(a_2*(float64(float64(284.0)+float64(day))))))
	omega = M.Acos(float64((-(M.Tan(phi))) * M.Tan(del)))
	ra = float64(float64(float64(37.6)*dr) * (float64(float64(float64(omega*M.Sin(phi))*M.Sin(del)) + float64(float64(M.Cos(phi)*M.Cos(del))*M.Sin(omega)))))
	e = float64(float64(float64(float64(float64(0.0023)*ra)/lamda)*M.Sqrt(tr)) * (float64(ta + float64(17.8))))
	if e < float64(0.0) {
		e = float64(0.0)
	}
	if UnitSystem == US {
		e = float64(e / (float64(25.4)))
	}
	return e
}

// climate.c:1010
func getFileFormat() int32 {
	var (
		recdType R.CPtr
		elemType R.CPtr
		filler   R.CPtr
		staID    R.CPtr
		s        R.CPtr
		line     R.CPtr
		y_2      int32
		m        int32
		d        int32
		n_2      int32
	)
	_, _, _, _, _, _, _, _, _, _ = recdType, elemType, filler, staID, s, line, y_2, m, d, n_2
	recdType = R.BufFrom(L0, 4)
	elemType = R.BufFrom(L0, 4)
	filler = R.BufFrom(L0, 5)
	staID = R.Buf(80)
	s = R.Buf(80)
	line = R.Buf(1024)
	y_2 = 0
	m = 0
	d = 0
	n_2 = 0
	if (R.Fgets(line, 1024, Fclimate.file)).IsNull() {
		return UNKNOWN_FORMAT
	}
	sstrncpy(recdType, line, 3)
	sstrncpy(filler, line.Add(23), 4)
	if (R.Strcmp(recdType, L14) == 0) && (R.Strcmp(filler, L15) == 0) {
		return TD3200
	}
	if R.Strlen(line) >= 233 {
		sstrncpy(elemType, line.Add(13), 3)
		n_2 = R.Atoi(elemType)
		if ((n_2 == 1) || (n_2 == 2)) || (n_2 == 151) {
			return DLY0204
		}
	}
	n_2 = R.Sscanf(line, L16, staID, &y_2, &m, &d, s)
	if n_2 == 5 {
		return USER_PREPARED
	}
	if isGhcndFormat(line) != 0 {
		return GHCND
	}
	return UNKNOWN_FORMAT
}

// climate.c:1055
func readFileLine(y_2 *int32, m *int32) {
	for R.Strlen(FileLine) == 0 {
		if (R.Fgets(FileLine, 1024, Fclimate.file)).IsNull() {
			return
		}
		if FileLine.G(0) == 10 {
			FileLine.S(0, 0)
		}
	}
	switch FileFormat {
	case USER_PREPARED:
		readUserFileLine(y_2, m)
		break
	case TD3200:
		readTD3200FileLine(y_2, m)
		break
	case DLY0204:
		readDLY0204FileLine(y_2, m)
		break
	case GHCND:
		readGhcndFileLine(y_2, m)
		break
	}
}

// climate.c:1082
func readUserFileLine(y_2 *int32, m *int32) {
	var (
		n_2   int32
		staID R.CPtr
	)
	_, _ = n_2, staID
	n_2 = 0
	staID = R.Buf(80)
	n_2 = R.Sscanf(FileLine, L17, staID, y_2, m)
	if n_2 < 3 {
		report_writeErrorMsg(ERR_CLIMATE_FILE_READ, R.CP(Fclimate.name[:]))
	}
}

// climate.c:1101
func readTD3200FileLine(y_2 *int32, m *int32) {
	var (
		recdType R.CPtr
		year     R.CPtr
		month    R.CPtr
	)
	_, _, _ = recdType, year, month
	recdType = R.BufFrom(L0, 4)
	year = R.BufFrom(L0, 5)
	month = R.BufFrom(L0, 3)
	if R.Strlen(FileLine) < 30 {
		report_writeErrorMsg(ERR_CLIMATE_FILE_READ, R.CP(Fclimate.name[:]))
		return
	}
	sstrncpy(recdType, FileLine, 3)
	if R.Strcmp(recdType, L14) != 0 {
		report_writeErrorMsg(ERR_CLIMATE_FILE_READ, R.CP(Fclimate.name[:]))
		return
	}
	sstrncpy(year, FileLine.Add(17), 4)
	sstrncpy(month, FileLine.Add(21), 2)
	*y_2 = R.Atoi(year)
	*m = R.Atoi(month)
}

// climate.c:1137
func readDLY0204FileLine(y_2 *int32, m *int32) {
	var (
		year  R.CPtr
		month R.CPtr
	)
	_, _ = year, month
	year = R.BufFrom(L0, 5)
	month = R.BufFrom(L0, 3)
	if R.Strlen(FileLine) < 16 {
		report_writeErrorMsg(ERR_CLIMATE_FILE_READ, R.CP(Fclimate.name[:]))
		return
	}
	sstrncpy(year, FileLine.Add(7), 4)
	sstrncpy(month, FileLine.Add(11), 2)
	*y_2 = R.Atoi(year)
	*m = R.Atoi(month)
}

// climate.c:1164
func readFileValues() {
	var (
		i   int32
		j   int32
		y_2 int32
		m   int32
	)
	_, _, _, _ = i, j, y_2, m
	i = 0
	j = 0
	y_2 = 0
	m = 0
	for i = 0; i < MAXCLIMATEVARS; i++ {
		for j = 0; j < MAXDAYSPERMONTH; j++ {
			FileData[i][j] = (-(float64(10000000000.0)))
		}
	}
	for !(ErrorCode != 0) {
		if R.Feof(Fclimate.file) != 0 {
			return
		}
		readFileLine(&y_2, &m)
		if (y_2 > FileYear) || (m > FileMonth) {
			return
		}
		switch FileFormat {
		case USER_PREPARED:
			parseUserFileLine()
			break
		case TD3200:
			parseTD3200FileLine()
			break
		case DLY0204:
			parseDLY0204FileLine()
			break
		case GHCND:
			parseGhcndFileLine()
			break
		}
		sstrncpy(FileLine, L0, 0)
	}
}

// climate.c:1201
func parseUserFileLine() {
	var (
		n_2   int32
		y_2   int32
		m     int32
		d     int32
		staID R.CPtr
		s0    R.CPtr
		s1    R.CPtr
		s2    R.CPtr
		s3    R.CPtr
		x     float64
	)
	_, _, _, _, _, _, _, _, _, _ = n_2, y_2, m, d, staID, s0, s1, s2, s3, x
	n_2 = 0
	y_2 = 0
	m = 0
	d = 0
	staID = R.Buf(80)
	s0 = R.BufFrom(L0, 80)
	s1 = R.BufFrom(L0, 80)
	s2 = R.BufFrom(L0, 80)
	s3 = R.BufFrom(L0, 80)
	x = 0
	n_2 = R.Sscanf(FileLine, L18, staID, &y_2, &m, &d, s0, s1, s2, s3)
	if n_2 < 4 {
		return
	}
	if (d < 1) || (d > 31) {
		return
	}
	if (R.Strlen(s0) > 0) && (s0.G(0) != 42) {
		x = R.Atof(s0)
		if UnitSystem == SI {
			x = float64(float64(float64(float64(9.0)/float64(5.0))*x) + float64(32.0))
		}
		FileData[TMAX][d] = x
	}
	if (R.Strlen(s1) > 0) && (s1.G(0) != 42) {
		x = R.Atof(s1)
		if UnitSystem == SI {
			x = float64(float64(float64(float64(9.0)/float64(5.0))*x) + float64(32.0))
		}
		FileData[TMIN][d] = x
	}
	if (R.Strlen(s2) > 0) && (s2.G(0) != 42) {
		FileData[EVAP][d] = R.Atof(s2)
	}
	if (R.Strlen(s3) > 0) && (s3.G(0) != 42) {
		FileData[WIND][d] = R.Atof(s3)
	}
}

// climate.c:1249
func parseTD3200FileLine() {
	var (
		i     int32
		param R.CPtr
	)
	_, _ = i, param
	i = 0
	param = R.BufFrom(L0, 5)
	sstrncpy(param, FileLine.Add(11), 4)
	for i = 0; i < MAXCLIMATEVARS; i++ {
		if R.Strcmp(param, ClimateVarWords[i]) == 0 {
			setTD3200FileValues(i)
		}
	}
}

// climate.c:1271
func setTD3200FileValues(i int32) {
	var (
		valCount R.CPtr
		day      R.CPtr
		sign     R.CPtr
		value    R.CPtr
		flag2    R.CPtr
		x        float64
		nValues  int32
		j        int32
		k        int32
		d        int32
	)
	_, _, _, _, _, _, _, _, _, _ = valCount, day, sign, value, flag2, x, nValues, j, k, d
	valCount = R.BufFrom(L0, 4)
	day = R.BufFrom(L0, 3)
	sign = R.BufFrom(L0, 2)
	value = R.BufFrom(L0, 6)
	flag2 = R.BufFrom(L0, 2)
	x = 0
	nValues = 0
	j = 0
	k = 0
	d = 0
	sstrncpy(valCount, FileLine.Add(27), 3)
	nValues = R.Atoi(valCount)
	if int32(R.Strlen(FileLine)) >= ((12 * nValues) + 30) {
		for j = 0; j < nValues; j++ {
			k = (30 + (j * 12))
			sstrncpy(day, FileLine.Add(k), 2)
			sstrncpy(sign, FileLine.Add((k + 4)), 1)
			sstrncpy(value, FileLine.Add((k + 5)), 5)
			sstrncpy(flag2, FileLine.Add((k + 11)), 1)
			d = R.Atoi(day)
			if (((R.Strcmp(value, L19) != 0) && ((flag2.G(0) == 48) || (flag2.G(0) == 49))) && (d > 0)) && (d <= 31) {
				x = R.Atof(value)
				if sign.G(0) == 45 {
					x = (-x)
				}
				if i == EVAP {
					x = float64(x / (float64(100.0)))
					if UnitSystem == SI {
						x = float64(x * (float64(25.4)))
					}
				}
				if i == WIND {
					x = float64(x / (float64(24.0)))
				}
				FileData[i][d] = x
			}
		}
	}
}

// climate.c:1336
func parseDLY0204FileLine() {
	var (
		j     int32
		k     int32
		p     int32
		param R.CPtr
		sign  R.CPtr
		value R.CPtr
		code  R.CPtr
		x     float64
	)
	_, _, _, _, _, _, _, _ = j, k, p, param, sign, value, code, x
	j = 0
	k = 0
	p = 0
	param = R.BufFrom(L0, 4)
	sign = R.BufFrom(L0, 2)
	value = R.BufFrom(L0, 6)
	code = R.BufFrom(L0, 2)
	x = 0
	sstrncpy(param, FileLine.Add(13), 3)
	p = R.Atoi(param)
	if p == 1 {
		p = TMAX
	} else if p == 2 {
		p = TMIN
	} else if p == 151 {
		p = EVAP
	} else {
		return
	}
	if R.Strlen(FileLine) < 233 {
		return
	}
	k = 16
	for j = 1; j <= 31; j++ {
		sstrncpy(sign, FileLine.Add(k), 1)
		sstrncpy(value, FileLine.Add((k + 1)), 5)
		sstrncpy(code, FileLine.Add((k + 6)), 1)
		k += 7
		if (R.Strcmp(value, L19) != 0) && (R.Strcmp(value, L20) != 0) {
			switch p {
			case TMAX, TMIN:
				x = float64(R.Atof(value) / float64(10.0))
				if sign.G(0) == 45 {
					x = (-x)
				}
				x = float64(float64(float64(float64(9.0)/float64(5.0))*x) + float64(32.0))
				break
			case EVAP:
				x = float64(R.Atof(value) / float64(10.0))
				if UnitSystem == US {
					x = float64(x / (float64(25.4)))
				}
				break
			default:
				return
			}
			FileData[p][j] = x
		}
	}
}

// climate.c:1401
func isGhcndFormat(line R.CPtr) int32 {
	var (
		i   int32
		ptr R.CPtr
	)
	_, _ = i, ptr
	i = 0
	ptr = (R.CPtr{})
	ptr = R.Strstr(line, L21)
	if ptr.IsNull() {
		return 0
	}
	FileDateFieldPos = int32((int64(ptr.O - line.O)))
	for i = TMIN; i <= WIND; i++ {
		FileFieldPos[i] = (-1)
	}
	ptr = R.Strstr(line, L22)
	if !ptr.IsNull() {
		FileFieldPos[TMIN] = int32((int64(ptr.O - line.O)))
	}
	ptr = R.Strstr(line, L23)
	if !ptr.IsNull() {
		FileFieldPos[TMAX] = int32((int64(ptr.O - line.O)))
	}
	ptr = R.Strstr(line, L7)
	if !ptr.IsNull() {
		FileFieldPos[EVAP] = int32((int64(ptr.O - line.O)))
	}
	FileWindType = WDMV
	ptr = R.Strstr(line, L24)
	if ptr.IsNull() {
		FileWindType = AWND
		ptr = R.Strstr(line, L25)
	}
	if !ptr.IsNull() {
		FileFieldPos[WIND] = int32((int64(ptr.O - line.O)))
	}
	for i = TMIN; i <= WIND; i++ {
		if FileFieldPos[i] >= 0 {
			return 1
		}
	}
	return 0
}

// climate.c:1445
func readGhcndFileLine(y_2 *int32, m *int32) {
	var (
		n_2 int32
	)
	_ = n_2
	n_2 = R.Sscanf(FileLine.Add(FileDateFieldPos), L26, y_2, m)
	if n_2 != 2 {
		*y_2 = (-99999)
		*m = (-99999)
	}
}

// climate.c:1463
func parseGhcndFileLine() {
	var (
		y_2 int32
		m   int32
		d   int32
		n_2 int32
		i   int32
		v   float64
	)
	_, _, _, _, _, _ = y_2, m, d, n_2, i, v
	y_2 = 0
	m = 0
	d = 0
	n_2 = 0
	i = 0
	v = 0
	n_2 = R.Sscanf(FileLine.Add(FileDateFieldPos), L27, &y_2, &m, &d)
	if n_2 < 3 {
		return
	}
	if (d < 1) || (d > 31) {
		return
	}
	for i = TMIN; i <= WIND; i++ {
		if FileFieldPos[i] >= 0 {
			if R.Sscanf(FileLine.Add(FileFieldPos[i]), L28, &v) > 0 {
				if M.Fabs(v) < float64(9999.0) {
					FileData[i][d] = convertGhcndValue(i, v)
				}
			}
		}
	}
}

// climate.c:1496
func convertGhcndValue(var_2 int32, v float64) float64 {
	switch var_2 {
	case TMIN, TMAX:
		switch FileTempUnits {
		case DEG_C10:
			return float64(float64(float64(float64(v/float64(10.0))*float64(9.0))/float64(5.0)) + float64(32.0))
		case DEG_C:
			return float64(float64(float64(v*float64(9.0))/float64(5.0)) + float64(32.0))
		default:
			return v
		}
		fallthrough
	case EVAP:
		switch FileTempUnits {
		case DEG_C10:
			v = float64(v / (float64(10.0)))
			if UnitSystem == US {
				v = float64(v / (float64(25.4)))
			}
			return v
		case DEG_C:
			if UnitSystem == US {
				v = float64(v / (float64(25.4)))
			}
			return v
		default:
			if UnitSystem == SI {
				v = float64(v * (float64(25.4)))
			}
			return v
		}
		fallthrough
	case WIND:
		switch FileTempUnits {
		case DEG_C10:
			if FileWindType == WDMV {
				return float64(float64(v*float64(0.62137)) / float64(24.0))
			} else {
				return float64(float64(float64(float64(v/float64(10.0))/float64(1000.0))*float64(0.62137)) * float64(3600.0))
			}
		case DEG_C:
			if FileWindType == WDMV {
				return float64(float64(v*float64(0.62137)) / float64(24.0))
			} else {
				return float64(float64(float64(v/float64(1000.0))*float64(0.62137)) * float64(3600.0))
			}
		default:
			if FileWindType == WDMV {
				return float64(v / float64(24.0))
			} else {
				return v
			}
		}
		fallthrough
	default:
		return v
	}
	return 0
}

// climate.c:1569
func updateTempMoveAve(tmin float64, tmax float64) {
	var (
		ta    float64
		tr    float64
		kount int32
		count float64
	)
	_, _, _, _ = ta, tr, kount, count
	ta = 0
	tr = 0
	kount = Tma.count
	count = float64(kount)
	ta = float64((float64(tmin + tmax)) / float64(2.0))
	tr = M.Fabs(float64(tmax - tmin))
	if kount == Tma.maxCount {
		Tma.tAve = float64((float64(float64(float64(Tma.tAve*count)+ta) - Tma.ta[Tma.front])) / count)
		Tma.tRng = float64((float64(float64(float64(Tma.tRng*count)+tr) - Tma.tr[Tma.front])) / count)
		Tma.ta[Tma.front] = ta
		Tma.tr[Tma.front] = tr
		Tma.front++
		if float64(Tma.front) == count {
			Tma.front = 0
		}
	} else {
		Tma.tAve = float64((float64(float64(Tma.tAve*count) + ta)) / (float64(count + float64(1))))
		Tma.tRng = float64((float64(float64(Tma.tRng*count) + tr)) / (float64(count + float64(1))))
		Tma.ta[Tma.front] = ta
		Tma.tr[Tma.front] = tr
		Tma.count++
		Tma.front++
		if Tma.count == Tma.maxCount {
			Tma.front = 0
		}
	}
}
