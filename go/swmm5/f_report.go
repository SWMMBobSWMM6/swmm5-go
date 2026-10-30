// Generated from EPA SWMM 5.2.4 (report.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var SysTime int64

// report.c:93
func report_readOptions(tok []R.CPtr, ntoks int32) int32 {
	var (
		k int32
		j int32
		m int32
		t int32
	)
	_, _, _, _ = k, j, m, t
	k = 0
	j = 0
	m = 0
	t = 0
	if ntoks < 2 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	k = int32(int8(findmatch(tok[0], ReportWords[:])))
	if k < 0 {
		return error_setInpError(ERR_KEYWORD, tok[0])
	}
	if (k < 2) || (k > 4) {
		m = findmatch(tok[1], NoYesWords[:])
		if (m < NO) || (m > YES) {
			return error_setInpError(ERR_KEYWORD, tok[1])
		}
		switch k {
		case 0:
			RptFlags.disabled = m
			return 0
		case 1:
			RptFlags.input = m
			return 0
		case 5:
			RptFlags.continuity = m
			return 0
		case 6:
			RptFlags.flowStats = m
			return 0
		case 7:
			RptFlags.controls = m
			return 0
		case 8:
			RptFlags.averages = m
			return 0
		case 9:
			return 0
		default:
			return error_setInpError(ERR_KEYWORD, tok[1])
		}
	} else {
		switch k {
		case 2:
			m = SUBCATCH
			break
		case 3:
			m = NODE
			break
		case 4:
			m = LINK
			break
		}
	}
	if strcomp(tok[1], L355) != 0 {
		k = NONE
	} else if strcomp(tok[1], L351) != 0 {
		k = ALL
	} else {
		k = SOME
		for t = 1; t < ntoks; t++ {
			j = project_findObject(m, tok[t])
			if j < 0 {
				return error_setInpError(ERR_NAME, tok[t])
			}
			switch m {
			case SUBCATCH:
				Subcatch[j].rptFlag = 1
				break
			case NODE:
				Node[j].rptFlag = 1
				break
			case LINK:
				Link[j].rptFlag = 1
				break
			}
		}
	}
	switch m {
	case SUBCATCH:
		RptFlags.subcatchments = k
		break
	case NODE:
		RptFlags.nodes = k
		break
	case LINK:
		RptFlags.links = k
		break
	}
	return 0
}

// report.c:171
func report_writeLine(line R.CPtr) {
	if Frpt.file != nil {
		R.Fprintf(Frpt.file, L356, line)
	}
}

// report.c:183
func report_writeSysTime() {
	var (
		theTime     R.CPtr
		elapsedTime float64
		endTime     int64
	)
	_, _, _ = theTime, elapsedTime, endTime
	theTime = R.Buf(9)
	elapsedTime = 0
	endTime = 0
	if Frpt.file != nil {
		R.Fprintf(Frpt.file, L357, R.Ctime(SysTime))
		R.Set(&endTime, R.Time())
		R.Fprintf(Frpt.file, L358, R.Ctime(endTime))
		elapsedTime = float64(endTime - SysTime)
		R.Fprintf(Frpt.file, L359)
		if elapsedTime < float64(1.0) {
			R.Fprintf(Frpt.file, L360)
		} else {
			elapsedTime = float64(elapsedTime / (float64(86400)))
			if elapsedTime >= float64(1.0) {
				R.Fprintf(Frpt.file, L361, R.D2I(M.Floor(elapsedTime)))
				elapsedTime = float64(elapsedTime - (M.Floor(elapsedTime)))
			}
			datetime_timeToStr(elapsedTime, theTime)
			R.Fprintf(Frpt.file, L169, theTime)
		}
	}
}

// report.c:220
func report_writeLogo() {
	R.Fprintf(Frpt.file, L362)
	R.Fprintf(Frpt.file, L363)
	R.Fprintf(Frpt.file, L276)
	R.Set(&SysTime, R.Time())
}

// report.c:235
func report_writeTitle() {
	var (
		i         int32
		lineCount int32
	)
	_, _ = i, lineCount
	i = 0
	lineCount = 0
	if ErrorCode != 0 {
		return
	}
	for i = 0; i < 3; i++ {
		if R.Strlen(Title[i]) > 0 {
			report_writeLine((Title[i]))
			lineCount++
		}
	}
}

// report.c:254
func report_writeOptions() {
	var (
		str R.CPtr
	)
	_ = str
	str = R.Buf(80)
	report_writeLine((L0))
	report_writeLine((L227))
	report_writeLine((L364))
	report_writeLine((L227))
	R.Fprintf(Frpt.file, L365, FlowUnitWords[FlowUnits])
	R.Fprintf(Frpt.file, L366)
	R.Fprintf(Frpt.file, L367)
	if (IgnoreRainfall != 0) || (Nobjects[GAGE] == 0) {
		R.Fprintf(Frpt.file, L4)
	} else {
		R.Fprintf(Frpt.file, L5)
	}
	R.Fprintf(Frpt.file, L368)
	if (IgnoreRDII != 0) || (Nobjects[UNITHYD] == 0) {
		R.Fprintf(Frpt.file, L4)
	} else {
		R.Fprintf(Frpt.file, L5)
	}
	R.Fprintf(Frpt.file, L369)
	if (IgnoreSnowmelt != 0) || (Nobjects[SNOWMELT] == 0) {
		R.Fprintf(Frpt.file, L4)
	} else {
		R.Fprintf(Frpt.file, L5)
	}
	R.Fprintf(Frpt.file, L370)
	if (IgnoreGwater != 0) || (Nobjects[AQUIFER] == 0) {
		R.Fprintf(Frpt.file, L4)
	} else {
		R.Fprintf(Frpt.file, L5)
	}
	R.Fprintf(Frpt.file, L371)
	if (IgnoreRouting != 0) || (Nobjects[LINK] == 0) {
		R.Fprintf(Frpt.file, L4)
	} else {
		R.Fprintf(Frpt.file, L5)
		R.Fprintf(Frpt.file, L372)
		if AllowPonding != 0 {
			R.Fprintf(Frpt.file, L5)
		} else {
			R.Fprintf(Frpt.file, L4)
		}
	}
	R.Fprintf(Frpt.file, L373)
	if (IgnoreQuality != 0) || (Nobjects[POLLUT] == 0) {
		R.Fprintf(Frpt.file, L4)
	} else {
		R.Fprintf(Frpt.file, L5)
	}
	if Nobjects[SUBCATCH] > 0 {
		R.Fprintf(Frpt.file, L374, InfilModelWords[InfilModel])
	}
	if Nobjects[LINK] > 0 {
		R.Fprintf(Frpt.file, L375, RouteModelWords[RouteModel])
	}
	if RouteModel == DW {
		R.Fprintf(Frpt.file, L376, SurchargeWords[SurchargeMethod])
	}
	datetime_dateToStr(StartDate, str)
	R.Fprintf(Frpt.file, L377, str)
	datetime_timeToStr(StartTime, str)
	R.Fprintf(Frpt.file, L378, str)
	datetime_dateToStr(EndDate, str)
	R.Fprintf(Frpt.file, L379, str)
	datetime_timeToStr(EndTime, str)
	R.Fprintf(Frpt.file, L378, str)
	R.Fprintf(Frpt.file, L380, StartDryDays)
	datetime_timeToStr(datetime_encodeTime(0, 0, ReportStep), str)
	R.Fprintf(Frpt.file, L381, str)
	if Nobjects[SUBCATCH] > 0 {
		datetime_timeToStr(datetime_encodeTime(0, 0, WetStep), str)
		R.Fprintf(Frpt.file, L382, str)
		datetime_timeToStr(datetime_encodeTime(0, 0, DryStep), str)
		R.Fprintf(Frpt.file, L383, str)
	}
	if Nobjects[LINK] > 0 {
		R.Fprintf(Frpt.file, L384, RouteStep)
		if RouteModel == DW {
			R.Fprintf(Frpt.file, L385)
			if CourantFactor > float64(0.0) {
				R.Fprintf(Frpt.file, L5)
			} else {
				R.Fprintf(Frpt.file, L4)
			}
			R.Fprintf(Frpt.file, L386, MaxTrials)
			R.Fprintf(Frpt.file, L387, NumThreads)
			R.Fprintf(Frpt.file, L388, float64(HeadTol*UCF(LENGTH)))
			if UnitSystem == US {
				R.Fprintf(Frpt.file, L389)
			} else {
				R.Fprintf(Frpt.file, L390)
			}
		}
	}
	report_writeLine((L0))
}

// report.c:355
func report_writeRainStats(i int32, r *TRainStats) {
	var (
		date1 R.CPtr
		date2 R.CPtr
	)
	_, _ = date1, date2
	date1 = R.BufFrom(L391, 12)
	date2 = R.BufFrom(L391, 12)
	if RptFlags.disabled != 0 {
		return
	}
	if i < 0 {
		report_writeLine((L0))
		report_writeLine((L255))
		report_writeLine((L392))
		report_writeLine((L255))
		R.Fprintf(Frpt.file, L393)
		R.Fprintf(Frpt.file, L394)
		R.Fprintf(Frpt.file, L245)
	} else {
		if r.startDate != float64((-693594)) {
			datetime_dateToStr(r.startDate, date1)
		}
		if r.endDate != float64((-693594)) {
			datetime_dateToStr(r.endDate, date2)
		}
		R.Fprintf(Frpt.file, L395, R.CP(Gage[i].staID[:]), date1, date2, (Gage[i].rainInterval / 60), r.periodsRain, r.periodsMissing, r.periodsMalfunc)
	}
	report_writeLine((L0))
}

// report.c:396
func report_writeRdiiStats(rainVol float64, rdiiVol float64) {
	var (
		ratio float64
		ucf1  float64
		ucf2  float64
	)
	_, _, _ = ratio, ucf1, ucf2
	ratio = 0
	ucf1 = 0
	ucf2 = 0
	if RptFlags.disabled != 0 {
		return
	}
	ucf1 = float64(UCF(LENGTH) * UCF(LANDAREA))
	if UnitSystem == US {
		ucf2 = float64(float64(0.64632) / float64(86400))
	} else {
		ucf2 = float64(float64(2.4466) / float64(86400))
	}
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L396)
	if UnitSystem == US {
		R.Fprintf(Frpt.file, L397)
	} else {
		R.Fprintf(Frpt.file, L398)
	}
	R.Fprintf(Frpt.file, L399)
	R.Fprintf(Frpt.file, L400, float64(rainVol*ucf1), float64(rainVol*ucf2))
	R.Fprintf(Frpt.file, L401, float64(rdiiVol*ucf1), float64(rdiiVol*ucf2))
	if rainVol == float64(0.0) {
		ratio = float64(0.0)
	} else {
		ratio = float64(rdiiVol / rainVol)
	}
	R.Fprintf(Frpt.file, L402, ratio)
	report_writeLine((L0))
}

// report.c:441
func report_writeControlActionsHeading() {
	report_writeLine((L0))
	report_writeLine((L255))
	report_writeLine((L403))
	report_writeLine((L255))
	R.Fprintf(Frpt.file, L276)
}

// report.c:452
func report_writeControlAction(aDate float64, linkID R.CPtr, value float64, ruleID R.CPtr) {
	var (
		theDate R.CPtr
		theTime R.CPtr
	)
	_, _ = theDate, theTime
	theDate = R.Buf(12)
	theTime = R.Buf(9)
	if RptFlags.disabled != 0 {
		return
	}
	datetime_dateToStr(aDate, theDate)
	datetime_timeToStr(aDate, theTime)
	R.Fprintf(Frpt.file, L404, theDate, theTime, linkID, value, ruleID)
}

// report.c:480
func report_writeRunoffError(totals *TRunoffTotals, totalArea float64) {
	if Frunoff.mode == USE_FILE {
		report_writeLine((L0))
		R.Fprintf(Frpt.file, L405, R.CP(Frunoff.name[:]))
		report_writeLine((L0))
		return
	}
	if totalArea == float64(0.0) {
		return
	}
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L406)
	if UnitSystem == US {
		R.Fprintf(Frpt.file, L407)
	} else {
		R.Fprintf(Frpt.file, L408)
	}
	R.Fprintf(Frpt.file, L409)
	if totals.initStorage > float64(0.0) {
		R.Fprintf(Frpt.file, L410, float64(float64(totals.initStorage*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.initStorage/totalArea)*UCF(RAINDEPTH)))
	}
	if Nobjects[SNOWMELT] > 0 {
		R.Fprintf(Frpt.file, L411, float64(float64(totals.initSnowCover*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.initSnowCover/totalArea)*UCF(RAINDEPTH)))
	}
	R.Fprintf(Frpt.file, L412, float64(float64(totals.rainfall*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.rainfall/totalArea)*UCF(RAINDEPTH)))
	if totals.runon > float64(0.0) {
		R.Fprintf(Frpt.file, L413, float64(float64(totals.runon*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.runon/totalArea)*UCF(RAINDEPTH)))
	}
	R.Fprintf(Frpt.file, L414, float64(float64(totals.evap*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.evap/totalArea)*UCF(RAINDEPTH)))
	R.Fprintf(Frpt.file, L415, float64(float64(totals.infil*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.infil/totalArea)*UCF(RAINDEPTH)))
	R.Fprintf(Frpt.file, L416, float64(float64(totals.runoff*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.runoff/totalArea)*UCF(RAINDEPTH)))
	if totals.drains > float64(0.0) {
		R.Fprintf(Frpt.file, L417, float64(float64(totals.drains*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.drains/totalArea)*UCF(RAINDEPTH)))
	}
	if Nobjects[SNOWMELT] > 0 {
		R.Fprintf(Frpt.file, L418, float64(float64(totals.snowRemoved*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.snowRemoved/totalArea)*UCF(RAINDEPTH)))
		R.Fprintf(Frpt.file, L419, float64(float64(totals.finalSnowCover*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.finalSnowCover/totalArea)*UCF(RAINDEPTH)))
	}
	R.Fprintf(Frpt.file, L420, float64(float64(totals.finalStorage*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.finalStorage/totalArea)*UCF(RAINDEPTH)))
	R.Fprintf(Frpt.file, L421, totals.pctError)
	report_writeLine((L0))
}

// report.c:577
func report_writeLoadingError(totals []TLoadingTotals) {
	var (
		p1 int32
		p2 int32
	)
	_, _ = p1, p2
	p1 = 0
	p2 = 0
	p1 = 1
	if (5) <= (Nobjects[POLLUT]) {
		p2 = (5)
	} else {
		p2 = (Nobjects[POLLUT])
	}
	for p1 <= Nobjects[POLLUT] {
		report_LoadingErrors((p1 - 1), (p2 - 1), totals)
		p1 = (p2 + 1)
		p2 = (p1 + 4)
		if (p2) <= (Nobjects[POLLUT]) {
			p2 = (p2)
		} else {
			p2 = (Nobjects[POLLUT])
		}
	}
}

// report.c:599
func report_LoadingErrors(p1 int32, p2 int32, totals []TLoadingTotals) {
	var (
		i     int32
		p     int32
		cf    float64
		units R.CPtr
	)
	_, _, _, _ = i, p, cf, units
	i = 0
	p = 0
	cf = float64(1.0)
	units = R.Buf(15)
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L422)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L423, Pollut[p].ID)
	}
	R.Fprintf(Frpt.file, L424)
	for p = p1; p <= p2; p++ {
		i = UnitSystem
		if Pollut[p].units == COUNT {
			i = 2
		}
		sstrncpy(units, LoadUnitsWords[i], 14)
		R.Fprintf(Frpt.file, L423, units)
	}
	R.Fprintf(Frpt.file, L422)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L425)
	}
	R.Fprintf(Frpt.file, L426)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, float64(totals[p].initLoad*cf))
	}
	R.Fprintf(Frpt.file, L428)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, float64(totals[p].buildup*cf))
	}
	R.Fprintf(Frpt.file, L429)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, float64(totals[p].deposition*cf))
	}
	R.Fprintf(Frpt.file, L430)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, float64(totals[p].sweeping*cf))
	}
	R.Fprintf(Frpt.file, L431)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, float64(totals[p].infil*cf))
	}
	R.Fprintf(Frpt.file, L432)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, float64(totals[p].bmpRemoval*cf))
	}
	R.Fprintf(Frpt.file, L433)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, float64(totals[p].runoff*cf))
	}
	R.Fprintf(Frpt.file, L434)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, float64(totals[p].finalLoad*cf))
	}
	R.Fprintf(Frpt.file, L435)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, totals[p].pctError)
	}
	report_writeLine((L0))
}

// report.c:685
func report_writeGwaterError(totals *TGwaterTotals, gwArea float64) {
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L406)
	if UnitSystem == US {
		R.Fprintf(Frpt.file, L436)
	} else {
		R.Fprintf(Frpt.file, L437)
	}
	R.Fprintf(Frpt.file, L409)
	R.Fprintf(Frpt.file, L438, float64(float64(totals.initStorage*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.initStorage/gwArea)*UCF(RAINDEPTH)))
	R.Fprintf(Frpt.file, L439, float64(float64(totals.infil*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.infil/gwArea)*UCF(RAINDEPTH)))
	R.Fprintf(Frpt.file, L440, float64(float64(totals.upperEvap*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.upperEvap/gwArea)*UCF(RAINDEPTH)))
	R.Fprintf(Frpt.file, L441, float64(float64(totals.lowerEvap*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.lowerEvap/gwArea)*UCF(RAINDEPTH)))
	R.Fprintf(Frpt.file, L442, float64(float64(totals.lowerPerc*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.lowerPerc/gwArea)*UCF(RAINDEPTH)))
	R.Fprintf(Frpt.file, L443, float64(float64(totals.gwater*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.gwater/gwArea)*UCF(RAINDEPTH)))
	R.Fprintf(Frpt.file, L420, float64(float64(totals.finalStorage*UCF(LENGTH))*UCF(LANDAREA)), float64(float64(totals.finalStorage/gwArea)*UCF(RAINDEPTH)))
	R.Fprintf(Frpt.file, L421, totals.pctError)
	report_writeLine((L0))
}

// report.c:737
func report_writeFlowError(totals *TRoutingTotals) {
	var (
		ucf1 float64
		ucf2 float64
	)
	_, _ = ucf1, ucf2
	ucf1 = 0
	ucf2 = 0
	ucf1 = float64(UCF(LENGTH) * UCF(LANDAREA))
	if UnitSystem == US {
		ucf2 = float64(float64(0.64632) / float64(86400))
	} else {
		ucf2 = float64(float64(2.4466) / float64(86400))
	}
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L444)
	if UnitSystem == US {
		R.Fprintf(Frpt.file, L445)
	} else {
		R.Fprintf(Frpt.file, L446)
	}
	R.Fprintf(Frpt.file, L447)
	R.Fprintf(Frpt.file, L448, float64(totals.dwInflow*ucf1), float64(totals.dwInflow*ucf2))
	R.Fprintf(Frpt.file, L449, float64(totals.wwInflow*ucf1), float64(totals.wwInflow*ucf2))
	R.Fprintf(Frpt.file, L450, float64(totals.gwInflow*ucf1), float64(totals.gwInflow*ucf2))
	R.Fprintf(Frpt.file, L451, float64(totals.iiInflow*ucf1), float64(totals.iiInflow*ucf2))
	R.Fprintf(Frpt.file, L452, float64(totals.exInflow*ucf1), float64(totals.exInflow*ucf2))
	R.Fprintf(Frpt.file, L453, float64(totals.outflow*ucf1), float64(totals.outflow*ucf2))
	R.Fprintf(Frpt.file, L454, float64(totals.flooding*ucf1), float64(totals.flooding*ucf2))
	R.Fprintf(Frpt.file, L414, float64(totals.evapLoss*ucf1), float64(totals.evapLoss*ucf2))
	R.Fprintf(Frpt.file, L455, float64(totals.seepLoss*ucf1), float64(totals.seepLoss*ucf2))
	R.Fprintf(Frpt.file, L456, float64(totals.initStorage*ucf1), float64(totals.initStorage*ucf2))
	R.Fprintf(Frpt.file, L457, float64(totals.finalStorage*ucf1), float64(totals.finalStorage*ucf2))
	R.Fprintf(Frpt.file, L421, totals.pctError)
	report_writeLine((L0))
}

// report.c:800
func report_writeQualError(QualTotals_2 []TRoutingTotals) {
	var (
		p1 int32
		p2 int32
	)
	_, _ = p1, p2
	p1 = 0
	p2 = 0
	p1 = 1
	if (5) <= (Nobjects[POLLUT]) {
		p2 = (5)
	} else {
		p2 = (Nobjects[POLLUT])
	}
	for p1 <= Nobjects[POLLUT] {
		report_QualErrors((p1 - 1), (p2 - 1), QualTotals_2)
		p1 = (p2 + 1)
		p2 = (p1 + 4)
		if (p2) <= (Nobjects[POLLUT]) {
			p2 = (p2)
		} else {
			p2 = (Nobjects[POLLUT])
		}
	}
}

// report.c:821
func report_QualErrors(p1 int32, p2 int32, QualTotals_2 []TRoutingTotals) {
	var (
		i     int32
		p     int32
		units R.CPtr
	)
	_, _, _ = i, p, units
	i = 0
	p = 0
	units = R.Buf(15)
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L422)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L423, Pollut[p].ID)
	}
	R.Fprintf(Frpt.file, L458)
	for p = p1; p <= p2; p++ {
		i = UnitSystem
		if Pollut[p].units == COUNT {
			i = 2
		}
		sstrncpy(units, LoadUnitsWords[i], 14)
		R.Fprintf(Frpt.file, L423, units)
	}
	R.Fprintf(Frpt.file, L422)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L425)
	}
	R.Fprintf(Frpt.file, L459)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, QualTotals_2[p].dwInflow)
	}
	R.Fprintf(Frpt.file, L460)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, QualTotals_2[p].wwInflow)
	}
	R.Fprintf(Frpt.file, L461)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, QualTotals_2[p].gwInflow)
	}
	R.Fprintf(Frpt.file, L462)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, QualTotals_2[p].iiInflow)
	}
	R.Fprintf(Frpt.file, L463)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, QualTotals_2[p].exInflow)
	}
	R.Fprintf(Frpt.file, L464)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, QualTotals_2[p].outflow)
	}
	R.Fprintf(Frpt.file, L465)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, QualTotals_2[p].flooding)
	}
	R.Fprintf(Frpt.file, L466)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, QualTotals_2[p].seepLoss)
	}
	R.Fprintf(Frpt.file, L467)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, QualTotals_2[p].reacted)
	}
	R.Fprintf(Frpt.file, L468)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, QualTotals_2[p].initStorage)
	}
	R.Fprintf(Frpt.file, L469)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, QualTotals_2[p].finalStorage)
	}
	R.Fprintf(Frpt.file, L435)
	for p = p1; p <= p2; p++ {
		R.Fprintf(Frpt.file, L427, QualTotals_2[p].pctError)
	}
	report_writeLine((L0))
}

// report.c:923
func report_writeMaxStats(maxMassBalErrs []TMaxStats, maxCourantCrit []TMaxStats, nMaxStats int32) {
	var (
		i int32
		j int32
		k int32
	)
	_, _, _ = i, j, k
	i = 0
	j = 0
	k = 0
	if (RouteModel != DW) || (Nobjects[LINK] == 0) {
		return
	}
	if nMaxStats <= 0 {
		return
	}
	if maxMassBalErrs[0].index >= 0 {
		report_writeLine((L0))
		report_writeLine((L470))
		report_writeLine((L471))
		report_writeLine((L470))
		for i = 0; i < nMaxStats; i++ {
			j = maxMassBalErrs[i].index
			if j < 0 {
				continue
			}
			R.Fprintf(Frpt.file, L472, Node[j].ID, maxMassBalErrs[i].value)
		}
		report_writeLine((L0))
	}
	if CourantFactor == float64(0.0) {
		return
	}
	report_writeLine((L0))
	report_writeLine((L473))
	report_writeLine((L474))
	report_writeLine((L473))
	k = 0
	for i = 0; i < nMaxStats; i++ {
		j = maxCourantCrit[i].index
		if j < 0 {
			continue
		}
		k++
		if maxCourantCrit[i].objType == NODE {
			R.Fprintf(Frpt.file, L475, Node[j].ID)
		} else {
			R.Fprintf(Frpt.file, L476, Link[j].ID)
		}
		R.Fprintf(Frpt.file, L477, maxCourantCrit[i].value)
	}
	if k == 0 {
		R.Fprintf(Frpt.file, L478)
	}
	report_writeLine((L0))
}

// report.c:977
func report_writeMaxFlowTurns(flowTurns []TMaxStats, nMaxStats int32) {
	var (
		i int32
		j int32
	)
	_, _ = i, j
	i = 0
	j = 0
	if Nobjects[LINK] == 0 {
		return
	}
	report_writeLine((L0))
	report_writeLine((L479))
	report_writeLine((L480))
	report_writeLine((L479))
	if (nMaxStats <= 0) || (flowTurns[0].index <= 0) {
		R.Fprintf(Frpt.file, L481)
	} else {
		for i = 0; i < nMaxStats; i++ {
			j = flowTurns[i].index
			if j < 0 {
				continue
			}
			R.Fprintf(Frpt.file, L482, Link[j].ID, flowTurns[i].value)
		}
	}
	report_writeLine((L0))
}

// report.c:1011
func report_writeNonconvergedStats(maxNonconverged []TMaxStats, nMaxStats int32) {
	var (
		i int32
		j int32
	)
	_, _ = i, j
	i = 0
	j = 0
	if (Nobjects[NODE] == 0) || (RouteModel != DW) {
		return
	}
	report_writeLine((L0))
	report_writeLine((L483))
	report_writeLine((L484))
	report_writeLine((L483))
	if ((nMaxStats <= 0) || (maxNonconverged[0].index <= 0)) || (maxNonconverged[0].value < float64(0.00005)) {
		R.Fprintf(Frpt.file, L485)
	} else {
		for i = 0; i < nMaxStats; i++ {
			j = maxNonconverged[i].index
			if (j < 0) || (maxNonconverged[i].value <= float64(0.0)) {
				continue
			}
			R.Fprintf(Frpt.file, L472, Node[j].ID, float64(float64(100.0)*maxNonconverged[i].value))
		}
	}
	report_writeLine((L0))
}

// report.c:1037
func report_writeTimeStepStats(timeStepStats *TTimeStepStats) {
	var (
		timeStepCount    float64
		totalRoutingTime float64
		fSteadyState     float64
	)
	_, _, _ = timeStepCount, totalRoutingTime, fSteadyState
	timeStepCount = float64(timeStepStats.timeStepCount)
	totalRoutingTime = 0
	fSteadyState = float64(0.0)
	if (Nobjects[LINK] == 0) || (float64(timeStepStats.timeStepCount) == float64(0.0)) {
		return
	}
	totalRoutingTime = float64(timeStepStats.steadyStateTime + timeStepStats.routingTime)
	if totalRoutingTime > float64(0.0) {
		fSteadyState = float64(float64(float64(100.0)*timeStepStats.steadyStateTime) / totalRoutingTime)
	}
	report_writeLine((L0))
	report_writeLine((L470))
	report_writeLine((L486))
	report_writeLine((L470))
	R.Fprintf(Frpt.file, L487, timeStepStats.minTimeStep)
	R.Fprintf(Frpt.file, L488, float64(timeStepStats.routingTime/timeStepCount))
	R.Fprintf(Frpt.file, L489, timeStepStats.maxTimeStep)
	R.Fprintf(Frpt.file, L490, (R.If[float64](((fSteadyState) <= (float64(100.0))), (fSteadyState), (float64(100.0)))))
	R.Fprintf(Frpt.file, L491, float64(timeStepStats.trialsCount/timeStepCount))
	R.Fprintf(Frpt.file, L492, float64(float64(float64(100.0)*float64(NonConvergeCount))/timeStepCount))
	if (RouteModel == DW) && (CourantFactor > float64(0.0)) {
		report_RouteStepFreq(timeStepStats)
	}
	report_writeLine((L0))
}

// report.c:1085
func report_RouteStepFreq(timeStepStats *TTimeStepStats) {
	var (
		totalSteps float64
		i          int32
	)
	_, _ = totalSteps, i
	totalSteps = float64(0.0)
	i = 0
	for i = 1; i < 6; i++ {
		totalSteps = float64(totalSteps + (float64(timeStepStats.timeStepCounts[i])))
	}
	if totalSteps == float64(0) {
		return
	}
	R.Fprintf(Frpt.file, L493)
	for i = 1; i < 6; i++ {
		R.Fprintf(Frpt.file, L494, timeStepStats.timeStepIntervals[(i-1)], timeStepStats.timeStepIntervals[i], float64(float64(float64(100.0)*float64((timeStepStats.timeStepCounts[i])))/totalSteps))
	}
}

// report.c:1113
func report_writeReport() {
	if ErrorCode != 0 {
		return
	}
	if Nperiods == 0 {
		return
	}
	if (RptFlags.subcatchments != NONE) && (((IgnoreRainfall == 0) || (IgnoreSnowmelt == 0)) || (IgnoreGwater == 0)) {
		report_Subcatchments()
	}
	if (IgnoreRouting == 1) && (IgnoreQuality == 1) {
		return
	}
	if RptFlags.nodes != NONE {
		report_Nodes()
	}
	if RptFlags.links != NONE {
		report_Links()
	}
}

// report.c:1135
func report_Subcatchments() {
	var (
		j           int32
		p           int32
		k           int32
		period      int32
		days        float64
		theDate     R.CPtr
		theTime     R.CPtr
		hasSnowmelt int32
		hasGwater   int32
		hasQuality  int32
	)
	_, _, _, _, _, _, _, _, _, _ = j, p, k, period, days, theDate, theTime, hasSnowmelt, hasGwater, hasQuality
	j = 0
	p = 0
	k = 0
	period = 0
	days = 0
	theDate = R.Buf(12)
	theTime = R.Buf(9)
	hasSnowmelt = R.B2I(((Nobjects[SNOWMELT] > 0) && (!(IgnoreSnowmelt != 0))))
	hasGwater = R.B2I(((Nobjects[AQUIFER] > 0) && (!(IgnoreGwater != 0))))
	hasQuality = R.B2I(((Nobjects[POLLUT] > 0) && (!(IgnoreQuality != 0))))
	if Nobjects[SUBCATCH] == 0 {
		return
	}
	report_writeLine((L0))
	report_writeLine((L479))
	report_writeLine((L495))
	report_writeLine((L479))
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		k = (Subcatch[j].rptFlag - 1)
		if k >= 0 {
			report_SubcatchHeader(Subcatch[j].ID)
			for period = 1; int64(period) <= Nperiods; period++ {
				output_readDateTime(int64(period), &days)
				datetime_dateToStr(days, theDate)
				datetime_timeToStr(days, theTime)
				output_readSubcatchResults(int64(period), k)
				R.Fprintf(Frpt.file, L496, theDate, theTime, float64(SubcatchResults[SUBCATCH_RAINFALL]), float64(float64(float64(SubcatchResults[SUBCATCH_EVAP])/float64(24.0))+float64(SubcatchResults[SUBCATCH_INFIL])), float64(SubcatchResults[SUBCATCH_RUNOFF]))
				if hasSnowmelt != 0 {
					R.Fprintf(Frpt.file, L497, float64(SubcatchResults[SUBCATCH_SNOWDEPTH]))
				}
				if hasGwater != 0 {
					R.Fprintf(Frpt.file, L498, float64(SubcatchResults[SUBCATCH_GW_ELEV]), float64(SubcatchResults[SUBCATCH_GW_FLOW]))
				}
				if hasQuality != 0 {
					for p = 0; p < Nobjects[POLLUT]; p++ {
						R.Fprintf(Frpt.file, L499, float64(SubcatchResults[(SUBCATCH_WASHOFF+p)]))
					}
				}
			}
			report_writeLine((L0))
		}
	}
}

// report.c:1192
func report_SubcatchHeader(id R.CPtr) {
	var (
		i           int32
		hasSnowmelt int32
		hasGwater   int32
		hasQuality  int32
	)
	_, _, _, _ = i, hasSnowmelt, hasGwater, hasQuality
	i = 0
	hasSnowmelt = R.B2I(((Nobjects[SNOWMELT] > 0) && (!(IgnoreSnowmelt != 0))))
	hasGwater = R.B2I(((Nobjects[AQUIFER] > 0) && (!(IgnoreGwater != 0))))
	hasQuality = R.B2I(((Nobjects[POLLUT] > 0) && (!(IgnoreQuality != 0))))
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L500, id)
	report_writeLine((L501))
	if hasSnowmelt > 0 {
		R.Fprintf(Frpt.file, L502)
	}
	if hasGwater != 0 {
		R.Fprintf(Frpt.file, L503)
		R.Fprintf(Frpt.file, L503)
	}
	if hasQuality != 0 {
		for i = 0; i < Nobjects[POLLUT]; i++ {
			R.Fprintf(Frpt.file, L503)
		}
	}
	R.Fprintf(Frpt.file, L504)
	if hasSnowmelt != 0 {
		R.Fprintf(Frpt.file, L505)
	}
	if hasGwater != 0 {
		R.Fprintf(Frpt.file, L506)
	}
	if hasQuality != 0 {
		for i = 0; i < Nobjects[POLLUT]; i++ {
			R.Fprintf(Frpt.file, L507, Pollut[i].ID)
		}
	}
	if UnitSystem == US {
		R.Fprintf(Frpt.file, L508, FlowUnitWords[FlowUnits])
	} else {
		R.Fprintf(Frpt.file, L509, FlowUnitWords[FlowUnits])
	}
	if hasSnowmelt != 0 {
		if UnitSystem == US {
			R.Fprintf(Frpt.file, L510)
		} else {
			R.Fprintf(Frpt.file, L511)
		}
	}
	if hasGwater != 0 {
		if UnitSystem == US {
			R.Fprintf(Frpt.file, L512, FlowUnitWords[FlowUnits])
		} else {
			R.Fprintf(Frpt.file, L513, FlowUnitWords[FlowUnits])
		}
	}
	if hasQuality != 0 {
		for i = 0; i < Nobjects[POLLUT]; i++ {
			R.Fprintf(Frpt.file, L507, QualUnitsWords[Pollut[i].units])
		}
	}
	report_writeLine((L501))
	if hasSnowmelt != 0 {
		R.Fprintf(Frpt.file, L502)
	}
	if hasGwater != 0 {
		R.Fprintf(Frpt.file, L503)
		R.Fprintf(Frpt.file, L503)
	}
	if hasQuality != 0 {
		for i = 0; i < Nobjects[POLLUT]; i++ {
			R.Fprintf(Frpt.file, L503)
		}
	}
}

// report.c:1261
func report_Nodes() {
	var (
		j       int32
		p       int32
		k       int32
		period  int32
		days    float64
		theDate R.CPtr
		theTime R.CPtr
	)
	_, _, _, _, _, _, _ = j, p, k, period, days, theDate, theTime
	j = 0
	p = 0
	k = 0
	period = 0
	days = 0
	theDate = R.Buf(12)
	theTime = R.Buf(9)
	if Nobjects[NODE] == 0 {
		return
	}
	report_writeLine((L0))
	report_writeLine((L514))
	report_writeLine((L515))
	report_writeLine((L514))
	for j = 0; j < Nobjects[NODE]; j++ {
		k = (Node[j].rptFlag - 1)
		if k >= 0 {
			report_NodeHeader(Node[j].ID)
			for period = 1; int64(period) <= Nperiods; period++ {
				output_readDateTime(int64(period), &days)
				datetime_dateToStr(days, theDate)
				datetime_timeToStr(days, theTime)
				output_readNodeResults(int64(period), k)
				R.Fprintf(Frpt.file, L516, theDate, theTime, float64(NodeResults[NODE_INFLOW]), float64(NodeResults[NODE_OVERFLOW]), float64(NodeResults[NODE_DEPTH]), float64(NodeResults[NODE_HEAD]))
				if !(IgnoreQuality != 0) {
					for p = 0; p < Nobjects[POLLUT]; p++ {
						R.Fprintf(Frpt.file, L195, float64(NodeResults[(NODE_QUAL+p)]))
					}
				}
			}
			report_writeLine((L0))
		}
	}
}

// report.c:1305
func report_NodeHeader(id R.CPtr) {
	var (
		i           int32
		lengthUnits R.CPtr
	)
	_, _ = i, lengthUnits
	i = 0
	lengthUnits = R.Buf(9)
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L517, id)
	report_writeLine((L518))
	for i = 0; i < Nobjects[POLLUT]; i++ {
		R.Fprintf(Frpt.file, L503)
	}
	R.Fprintf(Frpt.file, L519)
	if !(IgnoreQuality != 0) {
		for i = 0; i < Nobjects[POLLUT]; i++ {
			R.Fprintf(Frpt.file, L507, Pollut[i].ID)
		}
	}
	if UnitSystem == US {
		sstrncpy(lengthUnits, L520, 8)
	} else {
		sstrncpy(lengthUnits, L521, 8)
	}
	R.Fprintf(Frpt.file, L522, FlowUnitWords[FlowUnits], FlowUnitWords[FlowUnits], lengthUnits, lengthUnits)
	if !(IgnoreQuality != 0) {
		for i = 0; i < Nobjects[POLLUT]; i++ {
			R.Fprintf(Frpt.file, L507, QualUnitsWords[Pollut[i].units])
		}
	}
	report_writeLine((L518))
	if !(IgnoreQuality != 0) {
		for i = 0; i < Nobjects[POLLUT]; i++ {
			R.Fprintf(Frpt.file, L503)
		}
	}
}

// report.c:1339
func report_Links() {
	var (
		j       int32
		p       int32
		k       int32
		period  int32
		days    float64
		theDate R.CPtr
		theTime R.CPtr
	)
	_, _, _, _, _, _, _ = j, p, k, period, days, theDate, theTime
	j = 0
	p = 0
	k = 0
	period = 0
	days = 0
	theDate = R.Buf(12)
	theTime = R.Buf(9)
	if Nobjects[LINK] == 0 {
		return
	}
	report_writeLine((L0))
	report_writeLine((L514))
	report_writeLine((L523))
	report_writeLine((L514))
	for j = 0; j < Nobjects[LINK]; j++ {
		k = (Link[j].rptFlag - 1)
		if k >= 0 {
			report_LinkHeader(Link[j].ID)
			for period = 1; int64(period) <= Nperiods; period++ {
				output_readDateTime(int64(period), &days)
				datetime_dateToStr(days, theDate)
				datetime_timeToStr(days, theTime)
				output_readLinkResults(int64(period), k)
				R.Fprintf(Frpt.file, L516, theDate, theTime, float64(LinkResults[LINK_FLOW]), float64(LinkResults[LINK_VELOCITY]), float64(LinkResults[LINK_DEPTH]), float64(LinkResults[LINK_CAPACITY]))
				if !(IgnoreQuality != 0) {
					for p = 0; p < Nobjects[POLLUT]; p++ {
						R.Fprintf(Frpt.file, L195, float64(LinkResults[(LINK_QUAL+p)]))
					}
				}
			}
			report_writeLine((L0))
		}
	}
}

// report.c:1383
func report_LinkHeader(id R.CPtr) {
	var (
		i int32
	)
	_ = i
	i = 0
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L524, id)
	report_writeLine((L518))
	for i = 0; i < Nobjects[POLLUT]; i++ {
		R.Fprintf(Frpt.file, L503)
	}
	R.Fprintf(Frpt.file, L525)
	if !(IgnoreQuality != 0) {
		for i = 0; i < Nobjects[POLLUT]; i++ {
			R.Fprintf(Frpt.file, L507, Pollut[i].ID)
		}
	}
	if UnitSystem == US {
		R.Fprintf(Frpt.file, L526, FlowUnitWords[FlowUnits])
	} else {
		R.Fprintf(Frpt.file, L527, FlowUnitWords[FlowUnits])
	}
	if !(IgnoreQuality != 0) {
		for i = 0; i < Nobjects[POLLUT]; i++ {
			R.Fprintf(Frpt.file, L528, QualUnitsWords[Pollut[i].units])
		}
	}
	report_writeLine((L518))
	if !(IgnoreQuality != 0) {
		for i = 0; i < Nobjects[POLLUT]; i++ {
			R.Fprintf(Frpt.file, L503)
		}
	}
}

// report.c:1422
func report_writeErrorMsg(code int32, s R.CPtr) {
	if Frpt.file != nil {
		report_writeLine((L0))
		R.Fprintf(Frpt.file, error_getMsg(code, Msg), s)
	}
	ErrorCode = code
	if (ErrorCode <= ERR_INPUT) || (ErrorCode >= ERR_FILE_NAME) {
		R.Snprintf(ErrorMsg, 1024, error_getMsg(ErrorCode, Msg), s)
	}
}

// report.c:1446
func report_writeErrorCode() {
	if Frpt.file != nil {
		if (((ErrorCode >= ERR_MEMORY) && (ErrorCode <= ERR_TIMESTEP)) || ((ErrorCode >= ERR_FILE_NAME) && (ErrorCode <= ERR_OUT_FILE))) || (ErrorCode == ERR_SYSTEM) {
			R.Fprintf(Frpt.file, L169, error_getMsg(ErrorCode, Msg))
		}
	}
}

// report.c:1464
func report_writeInputErrorMsg(k int32, sect int32, line R.CPtr, lineCount int64) {
	if Frpt.file != nil {
		report_writeErrorMsg(k, ErrString)
		if sect < 0 {
			R.Fprintf(Frpt.file, L529, lineCount)
		} else {
			R.Fprintf(Frpt.file, L530, lineCount, SectWords[sect])
		}
		R.Fprintf(Frpt.file, L356, line)
	}
}

// report.c:1485
func report_writeWarningMsg(msg R.CPtr, id R.CPtr) {
	R.Fprintf(Frpt.file, L531, msg, id)
	Warnings++
}

// report.c:1499
func report_writeTseriesErrorMsg(code int32, tseries *TTable) {
	var (
		theDate R.CPtr
		theTime R.CPtr
		x       float64
	)
	_, _, _ = theDate, theTime, x
	theDate = R.Buf(20)
	theTime = R.Buf(20)
	x = 0
	if code == ERR_CURVE_SEQUENCE {
		x = tseries.x2
		datetime_dateToStr(x, theDate)
		datetime_timeToStr(x, theTime)
		report_writeErrorMsg(ERR_TIMESERIES_SEQUENCE, tseries.ID)
		R.Fprintf(Frpt.file, L532, theDate, theTime)
	} else {
		report_writeErrorMsg(code, tseries.ID)
	}
}
