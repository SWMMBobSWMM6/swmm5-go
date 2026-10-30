// Generated from EPA SWMM 5.2.4 (input.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

var MAXERRS int32 = 100
var Tok [40]R.CPtr
var Ntokens int32
var Mobjects [18]int32
var Mnodes [5]int32
var Mlinks [5]int32
var Mevents int32

// input.c:69
func input_countObjects() int32 {
	var (
		line      R.CPtr
		wLine     R.CPtr
		tok       R.CPtr
		sect      int32
		newsect   int32
		errcode   int32
		errsum    int32
		i         int32
		lineCount int64
	)
	_, _, _, _, _, _, _, _, _ = line, wLine, tok, sect, newsect, errcode, errsum, i, lineCount
	line = R.Buf(1025)
	wLine = R.Buf(1025)
	tok = (R.CPtr{})
	sect = (-1)
	newsect = 0
	errcode = 0
	errsum = 0
	i = 0
	lineCount = 0
	if ErrorCode != 0 {
		return ErrorCode
	}
	error_setInpError(0, L0)
	for i = 0; i < MAX_OBJ_TYPES; i++ {
		Nobjects[i] = 0
	}
	for i = 0; i < 5; i++ {
		Nnodes[i] = 0
	}
	for i = 0; i < 5; i++ {
		Nlinks[i] = 0
	}
	controls_init()
	for !(R.Fgets(line, 1024, Finp.file)).IsNull() {
		lineCount++
		sstrncpy(wLine, line, 1024)
		tok = R.Strtok(wLine, L184)
		if tok.IsNull() {
			continue
		}
		if tok.G(0) == 59 {
			continue
		}
		if tok.G(0) == 91 {
			newsect = findmatch(tok, SectWords[:])
			if newsect >= 0 {
				sect = newsect
				continue
			} else {
				sect = (-1)
				errcode = ERR_KEYWORD
			}
		}
		if sect == s_OPTION {
			errcode = readOption(line)
		} else if sect >= 0 {
			errcode = addObject(sect, tok)
		}
		if errcode != 0 {
			report_writeInputErrorMsg(errcode, sect, line, lineCount)
			errsum++
			if errsum >= MAXERRS {
				break
			}
		}
	}
	if errsum > 0 {
		ErrorCode = ERR_INPUT
	}
	return ErrorCode
}

// input.c:141
func input_readData() int32 {
	var (
		line       R.CPtr
		wLine      R.CPtr
		comment    R.CPtr
		sect       int32
		newsect    int32
		inperr     int32
		errsum     int32
		lineLength int32
		i          int32
		lineCount  int64
	)
	_, _, _, _, _, _, _, _, _, _ = line, wLine, comment, sect, newsect, inperr, errsum, lineLength, i, lineCount
	line = R.Buf(1025)
	wLine = R.Buf(1025)
	comment = (R.CPtr{})
	sect = 0
	newsect = 0
	inperr = 0
	errsum = 0
	lineLength = 0
	i = 0
	lineCount = 0
	if ErrorCode != 0 {
		return ErrorCode
	}
	error_setInpError(0, L0)
	for i = 0; i < MAX_OBJ_TYPES; i++ {
		Mobjects[i] = 0
	}
	for i = 0; i < 5; i++ {
		Mnodes[i] = 0
	}
	for i = 0; i < 5; i++ {
		Mlinks[i] = 0
	}
	Mevents = 0
	for i = 0; i < Nobjects[TSERIES]; i++ {
		Tseries[i].lastDate = float64(StartDate + StartTime)
	}
	sect = 0
	errsum = 0
	R.Rewind(Finp.file)
	for !(R.Fgets(line, 1024, Finp.file)).IsNull() {
		lineCount++
		sstrncpy(wLine, line, 1024)
		Ntokens = getTokens(wLine)
		if Ntokens == 0 {
			continue
		}
		if Tok[0].G(0) == 59 {
			continue
		}
		lineLength = int32(R.Strlen(line))
		if lineLength >= 1024 {
			comment = R.Strchr(line, 59)
			if !comment.IsNull() {
				lineLength = int32((int64(comment.O - line.O)))
			}
			if lineLength >= 1024 {
				inperr = ERR_LINE_LENGTH
				report_writeInputErrorMsg(inperr, sect, line, lineCount)
				errsum++
			}
		}
		if Tok[0].G(0) == 91 {
			newsect = findmatch(Tok[0], SectWords[:])
			if newsect >= 0 {
				if sect == s_TRANSECT {
					transect_validate((Nobjects[TRANSECT] - 1))
				}
				sect = newsect
				continue
			} else {
				inperr = error_setInpError(ERR_KEYWORD, Tok[0])
				report_writeInputErrorMsg(inperr, sect, line, lineCount)
				errsum++
				break
			}
		} else {
			inperr = parseLine(sect, line)
			if inperr > 0 {
				errsum++
				if errsum > MAXERRS {
					report_writeLine(L201)
				} else {
					report_writeInputErrorMsg(inperr, sect, line, lineCount)
				}
			}
		}
		if errsum > MAXERRS {
			break
		}
	}
	if errsum > 0 {
		ErrorCode = ERR_INPUT
	}
	return ErrorCode
}

// input.c:251
func addObject(objType int32, id R.CPtr) int32 {
	var (
		errcode int32
	)
	_ = errcode
	errcode = 0
	switch objType {
	case s_RAINGAGE:
		if !(project_addObject(GAGE, id, Nobjects[GAGE]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[GAGE]++
		break
	case s_SUBCATCH:
		if !(project_addObject(SUBCATCH, id, Nobjects[SUBCATCH]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[SUBCATCH]++
		break
	case s_AQUIFER:
		if !(project_addObject(AQUIFER, id, Nobjects[AQUIFER]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[AQUIFER]++
		break
	case s_UNITHYD:
		if project_findObject(UNITHYD, id) < 0 {
			if !(project_addObject(UNITHYD, id, Nobjects[UNITHYD]) != 0) {
				errcode = error_setInpError(ERR_DUP_NAME, id)
			}
			Nobjects[UNITHYD]++
		}
		break
	case s_SNOWMELT:
		if project_findObject(SNOWMELT, id) < 0 {
			if !(project_addObject(SNOWMELT, id, Nobjects[SNOWMELT]) != 0) {
				errcode = error_setInpError(ERR_DUP_NAME, id)
			}
			Nobjects[SNOWMELT]++
		}
		break
	case s_JUNCTION:
		if !(project_addObject(NODE, id, Nobjects[NODE]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[NODE]++
		Nnodes[JUNCTION]++
		break
	case s_OUTFALL:
		if !(project_addObject(NODE, id, Nobjects[NODE]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[NODE]++
		Nnodes[OUTFALL]++
		break
	case s_STORAGE:
		if !(project_addObject(NODE, id, Nobjects[NODE]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[NODE]++
		Nnodes[STORAGE]++
		break
	case s_DIVIDER:
		if !(project_addObject(NODE, id, Nobjects[NODE]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[NODE]++
		Nnodes[DIVIDER]++
		break
	case s_CONDUIT:
		if !(project_addObject(LINK, id, Nobjects[LINK]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[LINK]++
		Nlinks[CONDUIT]++
		break
	case s_PUMP:
		if !(project_addObject(LINK, id, Nobjects[LINK]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[LINK]++
		Nlinks[PUMP]++
		break
	case s_ORIFICE:
		if !(project_addObject(LINK, id, Nobjects[LINK]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[LINK]++
		Nlinks[ORIFICE]++
		break
	case s_WEIR:
		if !(project_addObject(LINK, id, Nobjects[LINK]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[LINK]++
		Nlinks[WEIR]++
		break
	case s_OUTLET:
		if !(project_addObject(LINK, id, Nobjects[LINK]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[LINK]++
		Nlinks[OUTLET]++
		break
	case s_POLLUTANT:
		if !(project_addObject(POLLUT, id, Nobjects[POLLUT]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[POLLUT]++
		break
	case s_LANDUSE:
		if !(project_addObject(LANDUSE, id, Nobjects[LANDUSE]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[LANDUSE]++
		break
	case s_PATTERN:
		if project_findObject(TIMEPATTERN, id) < 0 {
			if !(project_addObject(TIMEPATTERN, id, Nobjects[TIMEPATTERN]) != 0) {
				errcode = error_setInpError(ERR_DUP_NAME, id)
			}
			Nobjects[TIMEPATTERN]++
		}
		break
	case s_CURVE:
		if project_findObject(CURVE, id) < 0 {
			if !(project_addObject(CURVE, id, Nobjects[CURVE]) != 0) {
				errcode = error_setInpError(ERR_DUP_NAME, id)
			}
			Nobjects[CURVE]++
			id = R.Strtok((R.CPtr{}), L184)
			if findmatch(id, CurveTypeWords[:]) == SHAPE_CURVE {
				Nobjects[SHAPE]++
			}
		}
		break
	case s_TIMESERIES:
		if project_findObject(TSERIES, id) < 0 {
			if !(project_addObject(TSERIES, id, Nobjects[TSERIES]) != 0) {
				errcode = error_setInpError(ERR_DUP_NAME, id)
			}
			Nobjects[TSERIES]++
		}
		break
	case s_CONTROL:
		if match(id, L202) != 0 {
			Nobjects[CONTROL]++
		} else {
			controls_addToCount(id)
		}
		break
	case s_TRANSECT:
		if match(id, L203) != 0 {
			id = R.Strtok((R.CPtr{}), L184)
			if !id.IsNull() {
				if !(project_addObject(TRANSECT, id, Nobjects[TRANSECT]) != 0) {
					errcode = error_setInpError(ERR_DUP_NAME, id)
				}
				Nobjects[TRANSECT]++
			}
		}
		break
	case s_LID_CONTROL:
		if project_findObject(LID, id) < 0 {
			if !(project_addObject(LID, id, Nobjects[LID]) != 0) {
				errcode = error_setInpError(ERR_DUP_NAME, id)
			}
			Nobjects[LID]++
		}
		break
	case s_EVENT:
		NumEvents++
		break
	case s_STREET:
		if !(project_addObject(STREET, id, Nobjects[STREET]) != 0) {
			errcode = error_setInpError(ERR_DUP_NAME, id)
		}
		Nobjects[STREET]++
		break
	case s_INLET:
		if project_findObject(INLET, id) < 0 {
			if !(project_addObject(INLET, id, Nobjects[INLET]) != 0) {
				errcode = error_setInpError(ERR_DUP_NAME, id)
			}
			Nobjects[INLET]++
		}
		break
	}
	return errcode
}

// input.c:464
func parseLine(sect int32, line R.CPtr) int32 {
	var (
		j   int32
		err int32
	)
	_, _ = j, err
	j = 0
	err = 0
	switch sect {
	case s_TITLE:
		return readTitle(line)
	case s_RAINGAGE:
		j = Mobjects[GAGE]
		err = gage_readParams(j, Tok[:], Ntokens)
		Mobjects[GAGE]++
		return err
	case s_TEMP:
		return climate_readParams(Tok[:], Ntokens)
	case s_EVAP:
		return climate_readEvapParams(Tok[:], Ntokens)
	case s_ADJUST:
		return climate_readAdjustments(Tok[:], Ntokens)
	case s_SUBCATCH:
		j = Mobjects[SUBCATCH]
		err = subcatch_readParams(j, Tok[:], Ntokens)
		Mobjects[SUBCATCH]++
		return err
	case s_SUBAREA:
		return subcatch_readSubareaParams(Tok[:], Ntokens)
	case s_INFIL:
		return infil_readParams(InfilModel, Tok[:], Ntokens)
	case s_AQUIFER:
		j = Mobjects[AQUIFER]
		err = gwater_readAquiferParams(j, Tok[:], Ntokens)
		Mobjects[AQUIFER]++
		return err
	case s_GROUNDWATER:
		return gwater_readGroundwaterParams(Tok[:], Ntokens)
	case s_GWF:
		return gwater_readFlowExpression(Tok[:], Ntokens)
	case s_SNOWMELT:
		return snow_readMeltParams(Tok[:], Ntokens)
	case s_JUNCTION:
		return readNode(JUNCTION)
	case s_OUTFALL:
		return readNode(OUTFALL)
	case s_STORAGE:
		return readNode(STORAGE)
	case s_DIVIDER:
		return readNode(DIVIDER)
	case s_CONDUIT:
		return readLink(CONDUIT)
	case s_PUMP:
		return readLink(PUMP)
	case s_ORIFICE:
		return readLink(ORIFICE)
	case s_WEIR:
		return readLink(WEIR)
	case s_OUTLET:
		return readLink(OUTLET)
	case s_XSECTION:
		return link_readXsectParams(Tok[:], Ntokens)
	case s_TRANSECT:
		return transect_readParams(&Mobjects[TRANSECT], Tok[:], Ntokens)
	case s_LOSSES:
		return link_readLossParams(Tok[:], Ntokens)
	case s_POLLUTANT:
		j = Mobjects[POLLUT]
		err = landuse_readPollutParams(j, Tok[:], Ntokens)
		Mobjects[POLLUT]++
		return err
	case s_LANDUSE:
		j = Mobjects[LANDUSE]
		err = landuse_readParams(j, Tok[:], Ntokens)
		Mobjects[LANDUSE]++
		return err
	case s_BUILDUP:
		return landuse_readBuildupParams(Tok[:], Ntokens)
	case s_WASHOFF:
		return landuse_readWashoffParams(Tok[:], Ntokens)
	case s_COVERAGE:
		return subcatch_readLanduseParams(Tok[:], Ntokens)
	case s_INFLOW:
		return inflow_readExtInflow(Tok[:], Ntokens)
	case s_DWF:
		return inflow_readDwfInflow(Tok[:], Ntokens)
	case s_PATTERN:
		return inflow_readDwfPattern(Tok[:], Ntokens)
	case s_RDII:
		return rdii_readRdiiInflow(Tok[:], Ntokens)
	case s_UNITHYD:
		return rdii_readUnitHydParams(Tok[:], Ntokens)
	case s_LOADING:
		return subcatch_readInitBuildup(Tok[:], Ntokens)
	case s_TREATMENT:
		return treatmnt_readExpression(Tok[:], Ntokens)
	case s_CURVE:
		return table_readCurve(Tok[:], Ntokens)
	case s_TIMESERIES:
		return table_readTimeseries(Tok[:], Ntokens)
	case s_CONTROL:
		return readControl(Tok[:], Ntokens)
	case s_REPORT:
		return report_readOptions(Tok[:], Ntokens)
	case s_FILE:
		return iface_readFileParams(Tok[:], Ntokens)
	case s_LID_CONTROL:
		return lid_readProcParams(Tok[:], Ntokens)
	case s_LID_USAGE:
		return lid_readGroupParams(Tok[:], Ntokens)
	case s_EVENT:
		return readEvent(Tok[:], Ntokens)
	case s_STREET:
		return street_readParams(Tok[:], Ntokens)
	case s_INLET:
		return inlet_readDesignParams(Tok[:], Ntokens)
	case s_INLET_USAGE:
		return inlet_readUsageParams(Tok[:], Ntokens)
	default:
		return 0
	}
	return 0
}

// input.c:637
func readControl(tok []R.CPtr, ntoks int32) int32 {
	var (
		index   int32
		keyword int32
	)
	_, _ = index, keyword
	index = 0
	keyword = 0
	if ntoks < 2 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	if match(tok[0], L29) != 0 {
		return controls_addVariable(tok, ntoks)
	}
	if match(tok[0], L30) != 0 {
		return controls_addExpression(tok, ntoks)
	}
	keyword = findmatch(tok[0], RuleKeyWords[:])
	if keyword < 0 {
		return error_setInpError(ERR_KEYWORD, tok[0])
	}
	if keyword == 0 {
		if !(project_addObject(CONTROL, tok[1], Mobjects[CONTROL]) != 0) {
			return error_setInpError(ERR_DUP_NAME, Tok[1])
		}
		Mobjects[CONTROL]++
	}
	index = (Mobjects[CONTROL] - 1)
	if index < 0 {
		return error_setInpError(ERR_RULE, L0)
	}
	return controls_addRuleClause(index, keyword, Tok[:], Ntokens)
}

// input.c:680
func readOption(line R.CPtr) int32 {
	Ntokens = getTokens(line)
	if Ntokens < 2 {
		return 0
	}
	return project_readOption(Tok[0], Tok[1])
}

// input.c:694
func readTitle(line R.CPtr) int32 {
	var (
		i   int32
		n_2 int32
	)
	_, _ = i, n_2
	i = 0
	n_2 = 0
	for i = 0; i < 3; i++ {
		if R.Strlen(Title[i]) == 0 {
			n_2 = int32(R.Strlen(line))
			if line.G((n_2 - 1)) == 10 {
				line.S((n_2 - 1), 32)
			}
			sstrncpy(Title[i], line, 1024)
			break
		}
	}
	return 0
}

// input.c:721
func readNode(type_2 int32) int32 {
	var (
		j   int32
		k   int32
		err int32
	)
	_, _, _ = j, k, err
	j = Mobjects[NODE]
	k = Mnodes[type_2]
	err = node_readParams(j, type_2, k, Tok[:], Ntokens)
	Mobjects[NODE]++
	Mnodes[type_2]++
	return err
}

// input.c:738
func readLink(type_2 int32) int32 {
	var (
		j   int32
		k   int32
		err int32
	)
	_, _, _ = j, k, err
	j = Mobjects[LINK]
	k = Mlinks[type_2]
	err = link_readParams(j, type_2, k, Tok[:], Ntokens)
	Mobjects[LINK]++
	Mlinks[type_2]++
	return err
}

// input.c:755
func readEvent(tok []R.CPtr, ntoks int32) int32 {
	var (
		x [4]float64
	)
	_ = x
	x = ([4]float64{})
	if ntoks < 4 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	if !(datetime_strToDate(tok[0], &x[0]) != 0) {
		return error_setInpError(ERR_DATETIME, tok[0])
	}
	if !(datetime_strToTime(tok[1], &x[1]) != 0) {
		return error_setInpError(ERR_DATETIME, tok[1])
	}
	if !(datetime_strToDate(tok[2], &x[2]) != 0) {
		return error_setInpError(ERR_DATETIME, tok[2])
	}
	if !(datetime_strToTime(tok[3], &x[3]) != 0) {
		return error_setInpError(ERR_DATETIME, tok[3])
	}
	Event[Mevents].start = float64(x[0] + x[1])
	Event[Mevents].end = float64(x[2] + x[3])
	if Event[Mevents].start >= Event[Mevents].end {
		return error_setInpError(ERR_DATETIME, L204)
	}
	Mevents++
	return 0
}

// input.c:779
func findmatch(s R.CPtr, keyword []R.CPtr) int32 {
	var (
		i int32
	)
	_ = i
	i = 0
	for !keyword[i].IsNull() {
		if match(s, keyword[i]) != 0 {
			return (i)
		}
		i++
	}
	return (-1)
}

// input.c:798
func match(str R.CPtr, substr R.CPtr) int32 {
	var (
		i int32
		j int32
		k int32
	)
	_, _, _ = i, j, k
	i = 0
	j = 0
	k = 0
	if !(substr.G(0) != 0) {
		return (0)
	}
	for k = 0; str.G(k) != 0; k++ {
		if str.G(k) != 32 {
			break
		}
	}
	i = k
	j = 0
	for ; substr.G(j) != 0; func() { i++; j++ }() {
		if (!(str.G(i) != 0)) || ((func() int32 {
			if ((str.G(i)) >= 97) && ((str.G(i)) <= 122) {
				return ((str.G(i)) & (^32))
			}
			return (str.G(i))
		}()) != (func() int32 {
			if ((substr.G(j)) >= 97) && ((substr.G(j)) <= 122) {
				return ((substr.G(j)) & (^32))
			}
			return (substr.G(j))
		}())) {
			return (0)
		}
	}
	return (1)
}

// input.c:828
func getInt(s R.CPtr, y_2 *int32) int32 {
	var (
		x float64
	)
	_ = x
	x = 0
	if getDouble(s, &x) != 0 {
		if x < float64(0.0) {
			x = float64(x - (float64(0.01)))
		} else {
			x = float64(x + (float64(0.01)))
		}
		*y_2 = R.D2I(x)
		return 1
	}
	*y_2 = 0
	return 0
}

// input.c:850
func getFloat(s R.CPtr, y_2 *float32) int32 {
	var (
		endptr R.CPtr
	)
	_ = endptr
	endptr = (R.CPtr{})
	*y_2 = float32(float64(R.Strtod(s, &endptr)))
	if endptr.G(0) > 0 {
		return (0)
	}
	return (1)
}

// input.c:866
func getDouble(s R.CPtr, y_2 *float64) int32 {
	var (
		endptr R.CPtr
	)
	_ = endptr
	endptr = (R.CPtr{})
	*y_2 = R.Strtod(s, &endptr)
	if endptr.G(0) > 0 {
		return (0)
	}
	return (1)
}

// input.c:882
func getTokens(s R.CPtr) int32 {
	var (
		len_2 int32
		n_2   int32
		m     int32
		c     R.CPtr
	)
	_, _, _, _ = len_2, n_2, m, c
	len_2 = 0
	n_2 = 0
	m = 0
	c = (R.CPtr{})
	for n_2 = 0; n_2 < 40; n_2++ {
		Tok[n_2] = (R.CPtr{})
	}
	n_2 = 0
	c = R.Strchr(s, 59)
	if !c.IsNull() {
		c.S(0, 0)
	}
	len_2 = int32(R.Strlen(s))
	for (len_2 > 0) && (n_2 < 40) {
		m = int32(R.Strcspn(s, L184))
		if m == 0 {
			s = s.Add(1)
		} else {
			if s.G(0) == 34 {
				s = s.Add(1)
				len_2--
				m = int32(R.Strcspn(s, L205))
			}
			s.S(m, 0)
			Tok[n_2] = s
			n_2++
			s = s.Add((m + 1))
		}
		len_2 -= (m + 1)
	}
	return n_2
}
