// Generated from EPA SWMM 5.2.4 (rain.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

var RainStats TRainStats
var Condition int32
var TimeOffset int32
var DataOffset int32
var ValueOffset int32
var RainType int32
var Interval int32
var UnitsFactor float64
var RainAccum float32
var StationID R.CPtr
var AccumStartDate float64
var PreviousDate float64
var GageIndex int32
var hasStationName int32

// rain.c:118
func rain_open() {
	var (
		i     int32
		count int32
	)
	_, _ = i, count
	i = 0
	count = 0
	count = 0
	for i = 0; i < Nobjects[GAGE]; i++ {
		if Gage[i].dataSource == RAIN_FILE {
			count++
		}
	}
	Frain.file = (nil)
	if count == 0 {
		Frain.mode = NO_FILE
	} else {
		switch Frain.mode {
		case SCRATCH_FILE:
			getTempFileName(R.CP(Frain.name[:]))
			if (R.Set(&Frain.file, R.Fopen(R.CP(Frain.name[:]), L166))) == nil {
				report_writeErrorMsg(ERR_RAIN_FILE_SCRATCH, L0)
				return
			}
			break
		case USE_FILE:
			if (R.Set(&Frain.file, R.Fopen(R.CP(Frain.name[:]), L165))) == nil {
				report_writeErrorMsg(ERR_RAIN_FILE_OPEN, R.CP(Frain.name[:]))
				return
			}
			break
		case SAVE_FILE:
			if (R.Set(&Frain.file, R.Fopen(R.CP(Frain.name[:]), L166))) == nil {
				report_writeErrorMsg(ERR_RAIN_FILE_OPEN, R.CP(Frain.name[:]))
				return
			}
			break
		}
	}
	if (Frain.mode == SCRATCH_FILE) || (Frain.mode == SAVE_FILE) {
		createRainFile(count)
	}
	if Frain.mode != NO_FILE {
		initRainFile()
	}
	rdii_openRdii()
}

// rain.c:184
func rain_close() {
	if Frain.file != nil {
		R.Fclose(Frain.file)
		if Frain.mode == SCRATCH_FILE {
			R.Remove(R.CP(Frain.name[:]))
		}
	}
	Frain.file = (nil)
	rdii_closeRdii()
}

// rain.c:202
func createRainFile(count int32) {
	var (
		i         int32
		k         int32
		kount     int32
		filePos1  int32
		filePos2  int32
		filePos3  int32
		interval  int32
		dummy     int32
		staID     R.CPtr
		fileStamp R.CPtr
	)
	_, _, _, _, _, _, _, _, _, _ = i, k, kount, filePos1, filePos2, filePos3, interval, dummy, staID, fileStamp
	i = 0
	k = 0
	kount = count
	filePos1 = 0
	filePos2 = 0
	filePos3 = 0
	interval = 0
	dummy = (-1)
	staID = R.Buf(1025)
	fileStamp = R.BufFrom(L327, 11)
	if (ErrorCode != 0) || (!(Frain.file != nil)) {
		return
	}
	R.FwriteC(fileStamp, R.Strlen(fileStamp), Frain.file)
	R.Fwrite1(kount, 4, Frain.file)
	filePos1 = int32(R.Ftell(Frain.file))
	if count > 0 {
		report_writeRainStats((-1), &RainStats)
	}
	R.Strcpy(staID, L33)
	for i = 0; i < count; i++ {
		R.FwriteC(staID, uint64((1024 + 1)), Frain.file)
		for k = 1; k <= 3; k++ {
			R.Fwrite1(dummy, 4, Frain.file)
		}
	}
	filePos2 = int32(R.Ftell(Frain.file))
	for i = 0; i < Nobjects[GAGE]; i++ {
		if (ErrorCode != 0) || (Gage[i].dataSource != RAIN_FILE) {
			continue
		}
		if rainFileConflict(i) != 0 {
			break
		}
		R.Fseek(Frain.file, int64(filePos2), 0)
		if addGageToRainFile(i) != 0 {
			filePos3 = int32(R.Ftell(Frain.file))
			R.Fseek(Frain.file, int64(filePos1), 0)
			sstrncpy(staID, R.CP(Gage[i].staID[:]), 1024)
			interval = Interval
			R.FwriteC(staID, uint64((1024 + 1)), Frain.file)
			R.Fwrite1(interval, 4, Frain.file)
			R.Fwrite1(filePos2, 4, Frain.file)
			R.Fwrite1(filePos3, 4, Frain.file)
			filePos1 = int32(R.Ftell(Frain.file))
			filePos2 = filePos3
			report_writeRainStats(i, &RainStats)
		}
	}
	if ErrorCode != 0 {
		R.Fclose(Frain.file)
		Frain.file = (nil)
		R.Remove(R.CP(Frain.name[:]))
	}
}

// rain.c:278
func rainFileConflict(i int32) int32 {
	var (
		j     int32
		staID R.CPtr
		fname R.CPtr
	)
	_, _, _ = j, staID, fname
	j = 0
	staID = R.CP(Gage[i].staID[:])
	fname = R.CP(Gage[i].fname[:])
	for j = 0; j < i; j++ {
		if (strcomp(R.CP(Gage[j].staID[:]), staID) != 0) && (!(strcomp(R.CP(Gage[j].fname[:]), fname) != 0)) {
			report_writeErrorMsg(ERR_RAIN_FILE_CONFLICT, Gage[i].ID)
			return 1
		}
	}
	return 0
}

// rain.c:302
func addGageToRainFile(i int32) int32 {
	var (
		f          *R.CFile
		fileFormat int32
		hdrLines   int32
	)
	_, _, _ = f, fileFormat, hdrLines
	f = nil
	fileFormat = 0
	hdrLines = 0
	StationID = (R.CPtr{})
	if (R.Set(&f, R.Fopen(R.CP(Gage[i].fname[:]), L13))) == nil {
		report_writeErrorMsg(ERR_RAIN_FILE_DATA, R.CP(Gage[i].fname[:]))
	} else {
		fileFormat = findFileFormat(f, i, &hdrLines)
		if fileFormat == UNKNOWN_FORMAT {
			report_writeErrorMsg(ERR_RAIN_FILE_FORMAT, R.CP(Gage[i].fname[:]))
		} else {
			GageIndex = i
			readFile(f, fileFormat, hdrLines, Gage[i].startFileDate, Gage[i].endFileDate)
		}
		R.Fclose(f)
	}
	if ErrorCode != 0 {
		return 0
	} else {
		return 1
	}
	return 0
}

// rain.c:341
func initRainFile() {
	var (
		fileStamp R.CPtr
		fStamp    R.CPtr
		i         int32
		kount     int32
		filePos   int64
	)
	_, _, _, _, _ = fileStamp, fStamp, i, kount, filePos
	fileStamp = R.BufFrom(L327, 11)
	fStamp = R.BufFrom(L327, 11)
	i = 0
	kount = 0
	filePos = 0
	if (ErrorCode != 0) || (!(Frain.file != nil)) {
		return
	}
	R.Rewind(Frain.file)
	R.FreadC(fStamp, R.Strlen(fileStamp), Frain.file)
	if R.Strcmp(fStamp, fileStamp) != 0 {
		report_writeErrorMsg(ERR_RAIN_IFACE_FORMAT, L0)
		return
	}
	R.Fread1(&kount, 4, Frain.file)
	filePos = R.Ftell(Frain.file)
	for i = 0; i < Nobjects[GAGE]; i++ {
		if (ErrorCode != 0) || (Gage[i].dataSource != RAIN_FILE) {
			continue
		}
		R.Fseek(Frain.file, filePos, 0)
		if (!(findGageInFile(i, kount) != 0)) || (Gage[i].startFilePos == Gage[i].endFilePos) {
			report_writeErrorMsg(ERR_RAIN_FILE_GAGE, Gage[i].ID)
		}
	}
}

// rain.c:385
func findGageInFile(i int32, kount int32) int32 {
	var (
		k        int32
		interval int32
		filePos1 int32
		filePos2 int32
		staID    R.CPtr
	)
	_, _, _, _, _ = k, interval, filePos1, filePos2, staID
	k = 0
	interval = 0
	filePos1 = 0
	filePos2 = 0
	staID = R.BufFrom(L0, 1025)
	for k = 1; k <= kount; k++ {
		R.FreadC(staID, uint64((1024 + 1)), Frain.file)
		R.Fread1(&interval, 4, Frain.file)
		R.Fread1(&filePos1, 4, Frain.file)
		R.Fread1(&filePos2, 4, Frain.file)
		if R.Strcmp(staID, R.CP(Gage[i].staID[:])) == 0 {
			Gage[i].rainType = RAINFALL_VOLUME
			Gage[i].rainInterval = interval
			Gage[i].startFilePos = int64(filePos1)
			Gage[i].endFilePos = int64(filePos2)
			Gage[i].currentFilePos = Gage[i].startFilePos
			return 1
		}
	}
	return 0
}

// rain.c:420
func findFileFormat(f *R.CFile, i int32, hdrLines *int32) int32 {
	var (
		fileFormat int32
		lineCount  int32
		maxCount   int32
		n_2        int32
		div        int32
		sn2        int64
		recdType   R.CPtr
		elemType   R.CPtr
		coopID     R.CPtr
		line       R.CPtr
		year       int32
		month      int32
		day        int32
		hour       int32
		minute     int32
		elem       int32
		x          float32
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = fileFormat, lineCount, maxCount, n_2, div, sn2, recdType, elemType, coopID, line, year, month
	_, _, _, _, _ = day, hour, minute, elem, x
	fileFormat = 0
	lineCount = 0
	maxCount = 5
	n_2 = 0
	div = 0
	sn2 = 0
	recdType = R.Buf(4)
	elemType = R.Buf(5)
	coopID = R.BufFrom(L0, 6)
	line = R.Buf(1024)
	year = 0
	month = 0
	day = 0
	hour = 0
	minute = 0
	elem = 0
	x = 0
	fileFormat = UNKNOWN_FORMAT
	hasStationName = 0
	UnitsFactor = float64(1.0)
	Interval = 0
	*hdrLines = 0
	for lineCount = 1; lineCount <= maxCount; lineCount++ {
		if (R.Fgets(line, 1024, f)).IsNull() {
			return fileFormat
		}
		n_2 = R.Sscanf(line, L328, &sn2, &div, elemType)
		if n_2 == 3 {
			Interval = getNWSInterval(elemType)
			TimeOffset = Interval
			if Interval > 0 {
				fileFormat = NWS_SPACE_DELIMITED
				break
			}
		}
		n_2 = R.Sscanf(line.Add(37), L329, &div, elemType, recdType, &year)
		if n_2 == 4 {
			Interval = getNWSInterval(elemType)
			TimeOffset = Interval
			if Interval > 0 {
				fileFormat = NWS_SPACE_DELIMITED
				hasStationName = 1
				break
			}
		}
		n_2 = R.Sscanf(line, L330, &sn2, &div, elemType)
		if n_2 == 3 {
			Interval = getNWSInterval(elemType)
			TimeOffset = Interval
			if Interval > 0 {
				fileFormat = NWS_COMMA_DELIMITED
				break
			}
		}
		n_2 = R.Sscanf(line.Add(37), L331, &div, elemType, recdType, &year)
		if n_2 == 4 {
			Interval = getNWSInterval(elemType)
			TimeOffset = Interval
			if Interval > 0 {
				fileFormat = NWS_COMMA_DELIMITED
				hasStationName = 1
				break
			}
		}
		n_2 = R.Sscanf(line, L332, recdType, &sn2, &div, elemType)
		if n_2 == 4 {
			Interval = getNWSInterval(elemType)
			TimeOffset = Interval
			if Interval > 0 {
				fileFormat = NWS_TAPE
				break
			}
		}
		n_2 = R.Sscanf(line, L333, coopID, &sn2)
		if (n_2 == 2) && (R.Strcmp(coopID, L334) == 0) {
			fileFormat = findNWSOnlineFormat(f, line)
			break
		}
		n_2 = R.Sscanf(line, L335, &sn2, &year, &month, &day, &elem)
		if n_2 == 5 {
			if (elem == 123) && (R.Strlen(line) >= 185) {
				fileFormat = AES_HLY
				Interval = 3600
				TimeOffset = Interval
				UnitsFactor = float64(float64(1.0) / float64(25.4))
				break
			}
		}
		n_2 = R.Sscanf(line, L336, &sn2, &year, &month, &day, &elem)
		if n_2 == 5 {
			if (elem == 159) && (R.Strlen(line) >= 691) {
				fileFormat = CMC_FIF
				Interval = 900
			} else if (elem == 123) && (R.Strlen(line) >= 186) {
				fileFormat = CMC_HLY
				Interval = 3600
			}
			if (fileFormat == CMC_FIF) || (fileFormat == CMC_HLY) {
				TimeOffset = Interval
				UnitsFactor = float64(float64(1.0) / float64(25.4))
				break
			}
		}
		if parseStdLine(line, &year, &month, &day, &hour, &minute, &x) != 0 {
			fileFormat = STD_SPACE_DELIMITED
			RainType = Gage[i].rainType
			Interval = Gage[i].rainInterval
			if Gage[i].rainUnits == SI {
				UnitsFactor = float64(float64(1.0) / float64(25.4))
			}
			TimeOffset = 0
			StationID = R.CP(Gage[i].staID[:])
			break
		}
		(*hdrLines)++
	}
	if fileFormat != UNKNOWN_FORMAT {
		Gage[i].rainInterval = Interval
	}
	return fileFormat
}

// rain.c:584
func findNWSOnlineFormat(f *R.CFile, line R.CPtr) int32 {
	var (
		n_2        int32
		fileFormat int32
		str        R.CPtr
	)
	_, _, _ = n_2, fileFormat, str
	n_2 = 0
	fileFormat = UNKNOWN_FORMAT
	str = (R.CPtr{})
	R.Rewind(f)
	R.Fgets(line, 1024, f)
	if !(R.Set(&str, R.Strstr(line, L337))).IsNull() {
		Interval = 3600
		TimeOffset = Interval
		ValueOffset = int32((int64(str.O - line.O)))
		fileFormat = NWS_ONLINE_60
	} else if !(R.Set(&str, R.Strstr(line, L338))).IsNull() {
		Interval = 900
		TimeOffset = Interval
		ValueOffset = int32((int64(str.O - line.O)))
		fileFormat = NWS_ONLINE_15
	} else {
		return UNKNOWN_FORMAT
	}
	for n_2 = 1; n_2 <= 5; n_2++ {
		if (R.Fgets(line, 1024, f)).IsNull() {
			return UNKNOWN_FORMAT
		}
		if (R.Strstr(line, L334)).IsNull() {
			continue
		}
		str = R.Strrchr(line, 58)
		if str.IsNull() {
			return UNKNOWN_FORMAT
		}
		n_2 = int32((int64(str.O - line.O)))
		DataOffset = (n_2 - 11)
		return fileFormat
	}
	return UNKNOWN_FORMAT
}

// rain.c:641
func getNWSInterval(elemType R.CPtr) int32 {
	if R.Strcmp(elemType, L337) == 0 {
		return 3600
	} else if R.Strcmp(elemType, L338) == 0 {
		return 900
	} else if R.Strcmp(elemType, L339) == 0 {
		return 900
	} else {
		return 0
	}
	return 0
}

// rain.c:656
func readFile(f *R.CFile, fileFormat int32, hdrLines int32, day1 float64, day2 float64) {
	var (
		line R.CPtr
		i    int32
		n_2  int32
	)
	_, _, _ = line, i, n_2
	line = R.Buf(1024)
	i = 0
	n_2 = 0
	R.Rewind(f)
	RainStats.startDate = float64((-693594))
	RainStats.endDate = float64((-693594))
	RainStats.periodsRain = 0
	RainStats.periodsMissing = 0
	RainStats.periodsMalfunc = 0
	RainAccum = float32(float64(float64(0.0)))
	AccumStartDate = float64((-693594))
	PreviousDate = float64((-693594))
	for i = 1; i <= hdrLines; i++ {
		if (R.Fgets(line, 1024, f)).IsNull() {
			return
		}
	}
	for !(R.Fgets(line, 1024, f)).IsNull() {
		switch fileFormat {
		case STD_SPACE_DELIMITED:
			n_2 = readStdLine(line, day1, day2)
			break
		case NWS_TAPE, NWS_SPACE_DELIMITED, NWS_COMMA_DELIMITED, NWS_ONLINE_60, NWS_ONLINE_15:
			n_2 = readNWSLine(line, fileFormat, day1, day2)
			break
		case AES_HLY, CMC_FIF, CMC_HLY:
			n_2 = readCMCLine(line, fileFormat, day1, day2)
			break
		default:
			n_2 = (-1)
			break
		}
		if n_2 < 0 {
			break
		}
	}
}

// rain.c:717
func readNWSLine(line R.CPtr, fileFormat int32, day1 float64, day2 float64) int32 {
	var (
		flag1      int32
		flag2      int32
		isMissing  int32
		date1      float64
		result     int64
		k          int32
		y_2        int32
		m          int32
		d          int32
		n_2        int32
		hour       int32
		minute     int32
		v          int64
		x          float32
		lineLength int32
		nameLength int32
		__c1       R.CPtr
		__c2       R.CPtr
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = flag1, flag2, isMissing, date1, result, k, y_2, m, d, n_2, hour, minute
	_, _, _, _, _, _ = v, x, lineLength, nameLength, __c1, __c2
	flag1 = 0
	flag2 = 0
	isMissing = 0
	date1 = 0
	result = 1
	k = 0
	y_2 = 0
	m = 0
	d = 0
	n_2 = 0
	hour = 0
	minute = 0
	v = 0
	x = 0
	lineLength = (int32(R.Strlen(line)) - 1)
	nameLength = 0
	switch fileFormat {
	case NWS_TAPE:
		if lineLength <= 30 {
			return 0
		}
		if R.Sscanf(line.Add(17), L340, &y_2, &m, &d, &n_2) < 4 {
			return 0
		}
		k = 30
		break
	case NWS_SPACE_DELIMITED:
		if hasStationName != 0 {
			nameLength = 31
		}
		if lineLength <= (28 + nameLength) {
			return 0
		}
		k = (18 + nameLength)
		if R.Sscanf(line.Add(k), L341, &y_2, &m, &d) < 3 {
			return 0
		}
		k = (k + 10)
		break
	case NWS_COMMA_DELIMITED:
		if lineLength <= 28 {
			return 0
		}
		if R.Sscanf(line.Add(18), L342, &y_2, &m, &d) < 3 {
			return 0
		}
		k = 28
		break
	case NWS_ONLINE_60, NWS_ONLINE_15:
		if lineLength <= (DataOffset + 23) {
			return 0
		}
		if R.Sscanf(line.Add(DataOffset), L27, &y_2, &m, &d) < 3 {
			return 0
		}
		k = (DataOffset + 8)
		break
	default:
		return 0
	}
	date1 = datetime_encodeDate(y_2, m, d)
	if (day1 != float64((-693594))) && (date1 < day1) {
		return 0
	}
	if (day2 != float64((-693594))) && (date1 > day2) {
		return (-1)
	}
	for k < lineLength {
		flag1 = 0
		flag2 = 0
		v = 99999
		hour = 25
		minute = 0
		switch fileFormat {
		case NWS_TAPE:
			n_2 = R.Sscanf(line.Add(k), L343, &hour, &minute, &v, &flag1, &flag2)
			k += 12
			break
		case NWS_SPACE_DELIMITED:
			n_2 = R.Sscanf(line.Add(k), L344, &hour, &minute, &v, &flag1, &flag2)
			k += 16
			break
		case NWS_COMMA_DELIMITED:
			n_2 = R.Sscanf(line.Add(k), L345, &hour, &minute, &v, &flag1, &flag2)
			k += 16
			break
		case NWS_ONLINE_60, NWS_ONLINE_15:
			n_2 = R.Sscanf(line.Add(k), L346, &hour, &minute)
			__c1 = R.Buf(2)
			__c1.S(0, flag1)
			__c2 = R.Buf(2)
			__c2.S(0, flag1)
			n_2 += (readNwsOnlineValue(line.Add(ValueOffset), &v, __c2))
			flag1 = __c1.G(0)
			flag1 = __c2.G(0)
			if hour == 0 {
				hour = 24
				date1 = float64(date1 - (float64(1.0)))
			}
			k += lineLength
			break
		default:
			n_2 = 0
		}
		if (n_2 < 3) || (hour >= 25) {
			break
		}
		setCondition(flag1)
		if ((Condition == DELETED_PERIOD) || (Condition == MISSING_PERIOD)) || (flag1 == 77) {
			isMissing = 1
		} else if v >= 9999 {
			isMissing = 1
		} else {
			isMissing = 0
		}
		if flag1 == 97 {
			AccumStartDate = float64(date1 + datetime_encodeTime(hour, minute, 0))
		} else if flag1 == 65 {
			saveAccumRainfall(date1, hour, minute, v)
		} else {
			x = float32(float32(v) / float32(100.0))
			if (x > float32(0)) || (isMissing != 0) {
				saveRainfall(date1, hour, minute, x, isMissing)
			}
		}
		if ((flag1 == 65) || (flag1 == 125)) || (flag1 == 93) {
			Condition = 0
		}
	}
	return int32(result)
}

// rain.c:862
func readNwsOnlineValue(s R.CPtr, v *int64, flag R.CPtr) int32 {
	var (
		n_2 int32
		x   float32
	)
	_, _ = n_2, x
	n_2 = 0
	x = float32(99.98999786376953)
	if !(R.Strchr(s, 46)).IsNull() {
		n_2 = R.Sscanf(s, L347, &x, flag)
		*v = R.D2L(float64((float32(float32(float32(100.0)*x) + float32(0.5)))))
	} else {
		n_2 = R.Sscanf(s, L348, v, flag)
	}
	return n_2
}

// rain.c:891
func setCondition(flag int32) {
	switch flag {
	case 97, 65:
		Condition = ACCUMULATED_PERIOD
		break
	case 123, 125:
		Condition = DELETED_PERIOD
		break
	case 91, 93:
		Condition = MISSING_PERIOD
		break
	default:
		Condition = NO_CONDITION
	}
}

// rain.c:914
func readCMCLine(line R.CPtr, fileFormat int32, day1 float64, day2 float64) int32 {
	var (
		flag      int32
		isMissing int32
		date1     float64
		sn        int64
		v         int64
		col       int32
		j         int32
		jMax      int32
		elem      int32
		y_2       int32
		m         int32
		d         int32
		hour      int32
		minute    int32
		x         float32
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = flag, isMissing, date1, sn, v, col, j, jMax, elem, y_2, m, d
	_, _, _ = hour, minute, x
	flag = 0
	isMissing = 0
	date1 = 0
	sn = 0
	v = 0
	col = 0
	j = 0
	jMax = 0
	elem = 0
	y_2 = 0
	m = 0
	d = 0
	hour = 0
	minute = 0
	x = 0
	if fileFormat == AES_HLY {
		if R.Sscanf(line, L335, &sn, &y_2, &m, &d, &elem) < 5 {
			return 0
		}
		if y_2 < 100 {
			y_2 = (y_2 + 2000)
		} else {
			y_2 = (y_2 + 1000)
		}
		col = 17
	} else {
		if R.Sscanf(line, L336, &sn, &y_2, &m, &d, &elem) < 5 {
			return 0
		}
		col = 18
	}
	date1 = datetime_encodeDate(y_2, m, d)
	if (day1 != float64((-693594))) && (date1 < day1) {
		return 0
	}
	if (day2 != float64((-693594))) && (date1 > day2) {
		return (-1)
	}
	if (fileFormat == AES_HLY) && (elem != 123) {
		return 0
	} else if (fileFormat == CMC_FIF) && (elem != 159) {
		return 0
	} else if (fileFormat == CMC_HLY) && (elem != 123) {
		return 0
	}
	hour = 0
	minute = 0
	jMax = 24
	if fileFormat == CMC_FIF {
		jMax = 96
	}
	for j = 1; j <= jMax; j++ {
		if R.Sscanf(line.Add(col), L349, &v, &flag) < 2 {
			return 0
		}
		col += 7
		if v == -99999 {
			isMissing = 1
		} else {
			isMissing = 0
		}
		x = float32(float64((float64(float64(float64(v)/float64(10.0)) / float64(25.4)))))
		if (x > float32(0)) || (isMissing != 0) {
			saveRainfall(date1, hour, minute, x, isMissing)
		}
		if fileFormat == CMC_FIF {
			minute += 15
			if minute == 60 {
				minute = 0
				hour++
			}
		} else {
			hour++
		}
	}
	return 1
}

// rain.c:994
func readStdLine(line R.CPtr, day1 float64, day2 float64) int32 {
	var (
		date1  float64
		date2  float64
		year   int32
		month  int32
		day    int32
		hour   int32
		minute int32
		x      float32
	)
	_, _, _, _, _, _, _, _ = date1, date2, year, month, day, hour, minute, x
	date1 = 0
	date2 = 0
	year = 0
	month = 0
	day = 0
	hour = 0
	minute = 0
	x = 0
	if !(parseStdLine(line, &year, &month, &day, &hour, &minute, &x) != 0) {
		return 0
	}
	date1 = datetime_encodeDate(year, month, day)
	if (day1 != float64((-693594))) && (date1 < day1) {
		return 0
	}
	if (day2 != float64((-693594))) && (date1 > day2) {
		return (-1)
	}
	date2 = float64(date1 + datetime_encodeTime(hour, minute, 0))
	if date2 <= PreviousDate {
		report_writeErrorMsg(ERR_RAIN_FILE_SEQUENCE, R.CP(Gage[GageIndex].fname[:]))
		report_writeLine(line)
		return (-1)
	}
	PreviousDate = date2
	switch RainType {
	case RAINFALL_INTENSITY:
		x = float32(float32(x*float32(Interval)) / float32(3600.0))
		break
	case CUMULATIVE_RAINFALL:
		if x >= RainAccum {
			x = float32(x - RainAccum)
			RainAccum = float32(RainAccum + x)
		} else {
			RainAccum = x
		}
		break
	}
	x = float32(x * (float32(float64(UnitsFactor))))
	saveRainfall(date1, hour, minute, x, 0)
	return 1
}

// rain.c:1052
func parseStdLine(line R.CPtr, year *int32, month *int32, day *int32, hour *int32, minute *int32, value *float32) int32 {
	var (
		n_2   int32
		token R.CPtr
	)
	_, _ = n_2, token
	n_2 = 0
	token = R.Buf(1024)
	n_2 = R.Sscanf(line, L350, token, year, month, day, hour, minute, value)
	if n_2 < 7 {
		return 0
	}
	if (!StationID.IsNull()) && (!(strcomp(token, StationID) != 0)) {
		return 0
	}
	return 1
}

// rain.c:1078
func saveAccumRainfall(date1 float64, hour int32, minute int32, v int64) {
	var (
		date2 float64
		n_2   int32
		j     int32
		x     float32
	)
	_, _, _, _ = date2, n_2, j, x
	date2 = 0
	n_2 = 0
	j = 0
	x = 0
	if AccumStartDate == float64((-693594)) {
		return
	}
	date2 = float64(date1 + datetime_encodeTime(hour, minute, 0))
	n_2 = int32(((datetime_timeDiff(date2, AccumStartDate) / int64(Interval)) + 1))
	if v == 99999 {
		RainStats.periodsMissing += (int64(n_2))
		return
	}
	RainStats.periodsRain += (int64(n_2))
	x = float32(float32(float32(v)/float32(n_2)) / float32(100.0))
	if x > float32(0.0) {
		date2 = datetime_addSeconds(AccumStartDate, float64((-TimeOffset)))
		if RainStats.startDate == float64((-693594)) {
			RainStats.startDate = date2
		}
		for j = 0; j < n_2; j++ {
			R.Fwrite1(date2, 8, Frain.file)
			R.Fwrite1(x, 4, Frain.file)
			date2 = datetime_addSeconds(date2, float64(Interval))
			RainStats.endDate = date2
		}
	}
	AccumStartDate = float64((-693594))
}

// rain.c:1133
func saveRainfall(date1 float64, hour int32, minute int32, x float32, isMissing int32) {
	var (
		date2   float64
		seconds int32
	)
	_, _ = date2, seconds
	date2 = 0
	seconds = 0
	if isMissing != 0 {
		RainStats.periodsMissing++
	} else {
		RainStats.periodsRain++
	}
	if !(isMissing != 0) {
		seconds = (((3600 * hour) + (60 * minute)) - TimeOffset)
		date2 = datetime_addSeconds(date1, float64(seconds))
		R.Fwrite1(date2, 8, Frain.file)
		R.Fwrite1(x, 4, Frain.file)
		if RainStats.startDate == float64((-693594)) {
			RainStats.startDate = date2
		}
		RainStats.endDate = date2
	}
}
