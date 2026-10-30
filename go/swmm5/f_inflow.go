// Generated from EPA SWMM 5.2.4 (inflow.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

// inflow.c:41
func inflow_readExtInflow(tok []R.CPtr, ntoks int32) int32 {
	var (
		j        int32
		param    int32
		type_2   int32
		tseries  int32
		basePat  int32
		cf       float64
		sf       float64
		baseline float64
	)
	_, _, _, _, _, _, _, _ = j, param, type_2, tseries, basePat, cf, sf, baseline
	j = 0
	param = 0
	type_2 = CONCEN_INFLOW
	tseries = (-1)
	basePat = (-1)
	cf = float64(1.0)
	sf = float64(1.0)
	baseline = float64(0.0)
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(NODE, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	param = project_findObject(POLLUT, tok[1])
	if param < 0 {
		if match(tok[1], L183) != 0 {
			param = (-1)
		} else {
			return error_setInpError(ERR_NAME, tok[1])
		}
	}
	if R.Strlen(tok[2]) > 0 {
		tseries = project_findObject(TSERIES, tok[2])
		if tseries < 0 {
			return error_setInpError(ERR_NAME, tok[2])
		}
		Tseries[tseries].refersTo = EXTERNAL_INFLOW
	}
	if param == (-1) {
		type_2 = FLOW_INFLOW
		cf = float64(float64(1.0) / UCF(FLOW))
	}
	if (ntoks >= 4) && (param > (-1)) {
		if match(tok[3], L185) != 0 {
			type_2 = CONCEN_INFLOW
		} else if match(tok[3], L186) != 0 {
			type_2 = MASS_INFLOW
		} else {
			return error_setInpError(ERR_KEYWORD, tok[3])
		}
		if (ntoks >= 5) && (type_2 == MASS_INFLOW) {
			if !(getDouble(tok[4], &cf) != 0) {
				return error_setInpError(ERR_NUMBER, tok[4])
			}
			if cf <= float64(0.0) {
				return error_setInpError(ERR_NUMBER, tok[4])
			}
		}
	}
	if ntoks >= 6 {
		if !(getDouble(tok[5], &sf) != 0) {
			return error_setInpError(ERR_NUMBER, tok[5])
		}
	}
	if ntoks >= 7 {
		if !(getDouble(tok[6], &baseline) != 0) {
			return error_setInpError(ERR_NUMBER, tok[6])
		}
	}
	if ntoks >= 8 {
		basePat = project_findObject(TIMEPATTERN, tok[7])
		if basePat < 0 {
			return error_setInpError(ERR_NAME, tok[7])
		}
	}
	if type_2 == MASS_INFLOW {
		cf = float64(cf / (float64(28.317)))
	}
	return (inflow_setExtInflow(j, param, type_2, tseries, basePat, cf, baseline, sf))
}

// inflow.c:138
func inflow_setExtInflow(j int32, param int32, type_2 int32, tseries int32, basePat int32, cf float64, baseline float64, sf float64) int32 {
	var (
		inflow *TExtInflow
	)
	_ = inflow
	inflow = nil
	inflow = Node[j].extInflow
	for inflow != nil {
		if inflow.param == param {
			break
		}
		inflow = inflow.next
	}
	if inflow == nil {
		inflow = new(TExtInflow)
		if inflow == nil {
			return error_setInpError(ERR_MEMORY, L0)
		}
		inflow.next = Node[j].extInflow
		Node[j].extInflow = inflow
	}
	inflow.param = param
	inflow.type_ = type_2
	inflow.tSeries = tseries
	inflow.cFactor = cf
	inflow.sFactor = sf
	inflow.baseline = baseline
	inflow.basePat = basePat
	return 0
}

// inflow.c:187
func inflow_deleteExtInflows(j int32) {
	var (
		inflow1 *TExtInflow
		inflow2 *TExtInflow
	)
	_, _ = inflow1, inflow2
	inflow1 = nil
	inflow2 = nil
	inflow1 = Node[j].extInflow
	for inflow1 != nil {
		inflow2 = inflow1.next
		_ = inflow1
		inflow1 = inflow2
	}
}

// inflow.c:207
func inflow_getExtInflow(inflow *TExtInflow, aDate float64) float64 {
	var (
		month int32
		day   int32
		hour  int32
		p     int32
		k     int32
		cf    float64
		sf    float64
		blv   float64
		tsv   float64
	)
	_, _, _, _, _, _, _, _, _ = month, day, hour, p, k, cf, sf, blv, tsv
	month = 0
	day = 0
	hour = 0
	p = inflow.basePat
	k = inflow.tSeries
	cf = inflow.cFactor
	sf = inflow.sFactor
	blv = inflow.baseline
	tsv = float64(0.0)
	if p >= 0 {
		month = (datetime_monthOfYear(aDate) - 1)
		day = (datetime_dayOfWeek(aDate) - 1)
		hour = datetime_hourOfDay(aDate)
		blv = float64(blv * (getPatternFactor(p, month, day, hour)))
	}
	if k >= 0 {
		tsv = float64(table_tseriesLookup(&Tseries[k], aDate, 0) * sf)
	}
	return float64(cf * (float64(tsv + blv)))
}

// inflow.c:237
func inflow_readDwfInflow(tok []R.CPtr, ntoks int32) int32 {
	var (
		i      int32
		j      int32
		k      int32
		m      int32
		pats   [4]int32
		x      float64
		inflow *TDwfInflow
	)
	_, _, _, _, _, _, _ = i, j, k, m, pats, x, inflow
	i = 0
	j = 0
	k = 0
	m = 0
	pats = ([4]int32{})
	x = 0
	inflow = nil
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(NODE, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	k = project_findObject(POLLUT, tok[1])
	if k < 0 {
		if match(tok[1], L183) != 0 {
			k = (-1)
		} else {
			return error_setInpError(ERR_NAME, tok[1])
		}
	}
	if !(getDouble(tok[2], &x) != 0) {
		return error_setInpError(ERR_NUMBER, tok[2])
	}
	if k == (-1) {
		x = float64(x / (UCF(FLOW)))
	}
	for i = 0; i < 4; i++ {
		pats[i] = (-1)
	}
	for i = 3; i < 7; i++ {
		if i >= ntoks {
			break
		}
		if R.Strlen(tok[i]) == 0 {
			continue
		}
		m = project_findObject(TIMEPATTERN, tok[i])
		if m < 0 {
			return error_setInpError(ERR_NAME, tok[i])
		}
		pats[(i - 3)] = m
	}
	inflow = Node[j].dwfInflow
	for inflow != nil {
		if inflow.param == k {
			break
		}
		inflow = inflow.next
	}
	if inflow == nil {
		inflow = new(TDwfInflow)
		if inflow == nil {
			return error_setInpError(ERR_MEMORY, L0)
		}
		inflow.next = Node[j].dwfInflow
		Node[j].dwfInflow = inflow
	}
	inflow.param = k
	inflow.avgValue = x
	for i = 0; i < 4; i++ {
		inflow.patterns[i] = pats[i]
	}
	return 0
}

// inflow.c:311
func inflow_deleteDwfInflows(j int32) {
	var (
		inflow1 *TDwfInflow
		inflow2 *TDwfInflow
	)
	_, _ = inflow1, inflow2
	inflow1 = nil
	inflow2 = nil
	inflow1 = Node[j].dwfInflow
	for inflow1 != nil {
		inflow2 = inflow1.next
		_ = inflow1
		inflow1 = inflow2
	}
}

// inflow.c:331
func inflow_initDwfInflow(inflow *TDwfInflow) {
	var (
		i          int32
		p          int32
		tmpPattern [4]int32
	)
	_, _, _ = i, p, tmpPattern
	i = 0
	p = 0
	tmpPattern = ([4]int32{})
	for i = 0; i < 4; i++ {
		tmpPattern[i] = (-1)
	}
	for i = 0; i < 4; i++ {
		p = inflow.patterns[i]
		if p >= 0 {
			tmpPattern[Pattern[p].type_] = p
		}
	}
	for i = 0; i < 4; i++ {
		inflow.patterns[i] = tmpPattern[i]
	}
}

// inflow.c:361
func inflow_getDwfInflow(inflow *TDwfInflow, month int32, day int32, hour int32) float64 {
	var (
		p1 int32
		p2 int32
		f  float64
	)
	_, _, _ = p1, p2, f
	p1 = 0
	p2 = 0
	f = float64(1.0)
	p1 = inflow.patterns[MONTHLY_PATTERN]
	if p1 >= 0 {
		f = float64(f * (getPatternFactor(p1, month, day, hour)))
	}
	p1 = inflow.patterns[DAILY_PATTERN]
	if p1 >= 0 {
		f = float64(f * (getPatternFactor(p1, month, day, hour)))
	}
	p1 = inflow.patterns[HOURLY_PATTERN]
	p2 = inflow.patterns[WEEKEND_PATTERN]
	if p2 >= 0 {
		if (day == 0) || (day == 6) {
			f = float64(f * (getPatternFactor(p2, month, day, hour)))
		} else if p1 >= 0 {
			f = float64(f * (getPatternFactor(p1, month, day, hour)))
		}
	} else if p1 >= 0 {
		f = float64(f * (getPatternFactor(p1, month, day, hour)))
	}
	return float64(f * inflow.avgValue)
}

// inflow.c:394
func inflow_initDwfPattern(j int32) {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 0; i < 24; i++ {
		Pattern[j].factor[i] = float64(1.0)
	}
	Pattern[j].count = 0
	Pattern[j].type_ = (-1)
	Pattern[j].ID = (R.CPtr{})
}

// inflow.c:410
func inflow_readDwfPattern(tok []R.CPtr, ntoks int32) int32 {
	var (
		i   int32
		j   int32
		k   int32
		n_2 int32
	)
	_, _, _, _ = i, j, k, n_2
	i = 0
	j = 0
	k = 0
	n_2 = 1
	if ntoks < 2 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(TIMEPATTERN, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	if Pattern[j].ID.IsNull() {
		Pattern[j].ID = project_findID(TIMEPATTERN, tok[0])
		k = findmatch(tok[1], PatternTypeWords[:])
		if k < 0 {
			return error_setInpError(ERR_KEYWORD, tok[1])
		}
		Pattern[j].type_ = k
		n_2 = 2
	}
	for (ntoks > n_2) && (Pattern[j].count < 24) {
		i = Pattern[j].count
		if !(getDouble(tok[n_2], &Pattern[j].factor[i]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[n_2])
		}
		Pattern[j].count++
		n_2++
	}
	return 0
}

// inflow.c:456
func getPatternFactor(p int32, month int32, day int32, hour int32) float64 {
	switch Pattern[p].type_ {
	case MONTHLY_PATTERN:
		if (month >= 0) && (month < 12) {
			return Pattern[p].factor[month]
		}
		break
	case DAILY_PATTERN:
		if (day >= 0) && (day < 7) {
			return Pattern[p].factor[day]
		}
		break
	case HOURLY_PATTERN:
		if (hour >= 0) && (hour < 24) {
			return Pattern[p].factor[hour]
		}
		break
	case WEEKEND_PATTERN:
		if (day == 0) || (day == 6) {
			if (hour >= 0) && (hour < 24) {
				return Pattern[p].factor[hour]
			}
		}
		break
	}
	return float64(1.0)
}
