// Generated from EPA SWMM 5.2.4 (statsrpt.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var FlowFmt = R.Buf(6)
var Vcf float64

// statsrpt.c:84
func statsrpt_writeReport() {
	if (FlowUnits == MGD) || (FlowUnits == CMS) {
		sstrncpy(FlowFmt, L534, 5)
	} else {
		sstrncpy(FlowFmt, L535, 5)
	}
	if UnitSystem == US {
		Vcf = float64(float64(7.48) / float64(1.0e6))
	} else {
		Vcf = float64(float64(28.317) / float64(1.0e6))
	}
	if Nobjects[SUBCATCH] > 0 {
		if ((!(IgnoreRainfall != 0)) || ((Nobjects[SNOWMELT] > 0) && (!(IgnoreSnowmelt != 0)))) || ((Nobjects[AQUIFER] > 0) && (!(IgnoreGwater != 0))) {
			writeSubcatchRunoff()
			lid_writeWaterBalance()
			if !(IgnoreGwater != 0) {
				writeGroundwater()
			}
			if (Nobjects[POLLUT] > 0) && (!(IgnoreQuality != 0)) {
				writeSubcatchLoads()
			}
		}
	}
	if (Nobjects[LINK] > 0) && (!(IgnoreRouting != 0)) {
		writeNodeDepths()
		writeNodeFlows()
		if RouteModel == DW {
			writeNodeSurcharge()
		}
		writeNodeFlooding()
		writeStorageVolumes()
		writeOutfallLoads()
		inlet_writeStatsReport()
		writeLinkFlows()
		writeFlowClass()
		writeLinkSurcharge()
		writePumpFlows()
		if (Nobjects[POLLUT] > 0) && (!(IgnoreQuality != 0)) {
			writeLinkLoads()
		}
	}
}

// statsrpt.c:133
func writeSubcatchRunoff() {
	var (
		j   int32
		a_2 float64
		x   float64
		r   float64
	)
	_, _, _, _ = j, a_2, x, r
	j = 0
	a_2 = 0
	x = 0
	r = 0
	if Nobjects[SUBCATCH] == 0 {
		return
	}
	report_writeLine((L0))
	report_writeLine((L473))
	report_writeLine((L536))
	report_writeLine((L473))
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L537)
	if UnitSystem == US {
		R.Fprintf(Frpt.file, L538, VolUnitsWords[UnitSystem], FlowUnitWords[FlowUnits])
	} else {
		R.Fprintf(Frpt.file, L539, VolUnitsWords[UnitSystem], FlowUnitWords[FlowUnits])
	}
	R.Fprintf(Frpt.file, L540)
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		a_2 = Subcatch[j].area
		if a_2 == float64(0.0) {
			continue
		}
		R.Fprintf(Frpt.file, L541, Subcatch[j].ID)
		x = float64(SubcatchStats[j].precip * UCF(RAINDEPTH))
		R.Fprintf(Frpt.file, L542, float64(x/a_2))
		x = float64(SubcatchStats[j].runon * UCF(RAINDEPTH))
		R.Fprintf(Frpt.file, L542, float64(x/a_2))
		x = float64(SubcatchStats[j].evap * UCF(RAINDEPTH))
		R.Fprintf(Frpt.file, L542, float64(x/a_2))
		x = float64(SubcatchStats[j].infil * UCF(RAINDEPTH))
		R.Fprintf(Frpt.file, L542, float64(x/a_2))
		x = float64(SubcatchStats[j].impervRunoff * UCF(RAINDEPTH))
		R.Fprintf(Frpt.file, L542, float64(x/a_2))
		x = float64(SubcatchStats[j].pervRunoff * UCF(RAINDEPTH))
		R.Fprintf(Frpt.file, L542, float64(x/a_2))
		x = float64(SubcatchStats[j].runoff * UCF(RAINDEPTH))
		R.Fprintf(Frpt.file, L542, float64(x/a_2))
		x = float64(SubcatchStats[j].runoff * Vcf)
		R.Fprintf(Frpt.file, L543, x)
		x = float64(SubcatchStats[j].maxFlow * UCF(FLOW))
		R.Fprintf(Frpt.file, L544, x)
		r = float64(SubcatchStats[j].precip + SubcatchStats[j].runon)
		if r > float64(0.0) {
			r = float64(SubcatchStats[j].runoff / r)
		}
		R.Fprintf(Frpt.file, L545, r)
	}
	report_writeLine((L0))
}

// statsrpt.c:190
func writeGroundwater() {
	var (
		i            int32
		j            int32
		count        int32
		totalSeconds float64
		x            [9]float64
	)
	_, _, _, _, _ = i, j, count, totalSeconds, x
	i = 0
	j = 0
	count = 0
	totalSeconds = float64(NewRunoffTime / float64(1000.0))
	x = ([9]float64{})
	if Nobjects[SUBCATCH] == 0 {
		return
	}
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		if Subcatch[j].groundwater != nil {
			count++
		}
	}
	if count == 0 {
		return
	}
	report_writeLine((L0))
	report_writeLine((L188))
	report_writeLine((L546))
	report_writeLine((L188))
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L547)
	if UnitSystem == US {
		R.Fprintf(Frpt.file, L548, FlowUnitWords[FlowUnits])
	} else {
		R.Fprintf(Frpt.file, L549, FlowUnitWords[FlowUnits])
	}
	R.Fprintf(Frpt.file, L550)
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		if (Subcatch[j].area == float64(0.0)) || (Subcatch[j].groundwater == nil) {
			continue
		}
		R.Fprintf(Frpt.file, L541, Subcatch[j].ID)
		x[0] = float64(Subcatch[j].groundwater.stats.infil * UCF(RAINDEPTH))
		x[1] = float64(Subcatch[j].groundwater.stats.evap * UCF(RAINDEPTH))
		x[2] = float64(Subcatch[j].groundwater.stats.deepFlow * UCF(RAINDEPTH))
		x[3] = float64(Subcatch[j].groundwater.stats.latFlow * UCF(RAINDEPTH))
		x[4] = float64(float64(Subcatch[j].groundwater.stats.maxFlow*UCF(FLOW)) * Subcatch[j].area)
		x[5] = float64(Subcatch[j].groundwater.stats.avgUpperMoist / totalSeconds)
		x[6] = float64(float64(Subcatch[j].groundwater.stats.avgWaterTable*UCF(LENGTH)) / totalSeconds)
		x[7] = Subcatch[j].groundwater.stats.finalUpperMoist
		x[8] = float64(Subcatch[j].groundwater.stats.finalWaterTable * UCF(LENGTH))
		for i = 0; i < 9; i++ {
			R.Fprintf(Frpt.file, L544, x[i])
		}
	}
	report_writeLine((L0))
}

// statsrpt.c:245
func writeSubcatchLoads() {
	var (
		i            int32
		j            int32
		p            int32
		x            float64
		totals       []float64
		units        R.CPtr
		subcatchLine R.CPtr
		pollutLine   R.CPtr
	)
	_, _, _, _, _, _, _, _ = i, j, p, x, totals, units, subcatchLine, pollutLine
	i = 0
	j = 0
	p = 0
	x = 0
	totals = nil
	units = R.Buf(15)
	subcatchLine = R.BufFrom(L551, 21)
	pollutLine = R.BufFrom(L552, 15)
	totals = make([]float64, int(uint64(Nobjects[POLLUT])))
	if totals != nil {
		report_writeLine((L0))
		report_writeLine((L553))
		report_writeLine((L554))
		report_writeLine((L553))
		report_writeLine((L0))
		R.Fprintf(Frpt.file, L356, subcatchLine)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			R.Fprintf(Frpt.file, L169, pollutLine)
		}
		R.Fprintf(Frpt.file, L555)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			R.Fprintf(Frpt.file, L423, Pollut[p].ID)
		}
		R.Fprintf(Frpt.file, L556)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			i = UnitSystem
			if Pollut[p].units == COUNT {
				i = 2
			}
			sstrncpy(units, LoadUnitsWords[i], 14)
			R.Fprintf(Frpt.file, L423, units)
			totals[p] = float64(0.0)
		}
		R.Fprintf(Frpt.file, L356, subcatchLine)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			R.Fprintf(Frpt.file, L169, pollutLine)
		}
		for j = 0; j < Nobjects[SUBCATCH]; j++ {
			R.Fprintf(Frpt.file, L541, Subcatch[j].ID)
			for p = 0; p < Nobjects[POLLUT]; p++ {
				x = Subcatch[j].totalLoad[p]
				totals[p] = float64(totals[p] + x)
				if Pollut[p].units == COUNT {
					if (x) > float64(0.0) {
						x = M.Log10((x))
					} else {
						x = (x)
					}
				}
				R.Fprintf(Frpt.file, L427, x)
			}
		}
		R.Fprintf(Frpt.file, L356, subcatchLine)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			R.Fprintf(Frpt.file, L169, pollutLine)
		}
		R.Fprintf(Frpt.file, L557)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			x = totals[p]
			if Pollut[p].units == COUNT {
				if (x) > float64(0.0) {
					x = M.Log10((x))
				} else {
					x = (x)
				}
			}
			R.Fprintf(Frpt.file, L427, x)
		}
		_ = totals
		report_writeLine((L0))
	}
}

// statsrpt.c:310
func writeNodeDepths() {
	var (
		j    int32
		days int32
		hrs  int32
		mins int32
	)
	_, _, _, _ = j, days, hrs, mins
	j = 0
	days = 0
	hrs = 0
	mins = 0
	if Nobjects[LINK] == 0 {
		return
	}
	report_writeLine((L0))
	report_writeLine((L558))
	report_writeLine((L559))
	report_writeLine((L558))
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L560)
	if UnitSystem == US {
		R.Fprintf(Frpt.file, L561)
	} else {
		R.Fprintf(Frpt.file, L562)
	}
	R.Fprintf(Frpt.file, L563)
	for j = 0; j < Nobjects[NODE]; j++ {
		R.Fprintf(Frpt.file, L541, Node[j].ID)
		R.Fprintf(Frpt.file, L564, NodeTypeWords[Node[j].type_])
		getElapsedTime(NodeStats[j].maxDepthDate, &days, &hrs, &mins)
		R.Fprintf(Frpt.file, L565, float64(float64(NodeStats[j].avgDepth/float64(ReportStepCount))*UCF(LENGTH)), float64(NodeStats[j].maxDepth*UCF(LENGTH)), float64((float64(NodeStats[j].maxDepth+Node[j].invertElev))*UCF(LENGTH)), days, hrs, mins, NodeStats[j].maxRptDepth)
	}
	report_writeLine((L0))
}

// statsrpt.c:353
func writeNodeFlows() {
	var (
		j     int32
		days1 int32
		hrs1  int32
		mins1 int32
	)
	_, _, _, _ = j, days1, hrs1, mins1
	j = 0
	days1 = 0
	hrs1 = 0
	mins1 = 0
	report_writeLine((L0))
	report_writeLine((L188))
	report_writeLine((L566))
	report_writeLine((L188))
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L567, FlowUnitWords[FlowUnits], FlowUnitWords[FlowUnits], VolUnitsWords[UnitSystem], VolUnitsWords[UnitSystem])
	R.Fprintf(Frpt.file, L568)
	for j = 0; j < Nobjects[NODE]; j++ {
		R.Fprintf(Frpt.file, L541, Node[j].ID)
		R.Fprintf(Frpt.file, L569, NodeTypeWords[Node[j].type_])
		getElapsedTime(NodeStats[j].maxInflowDate, &days1, &hrs1, &mins1)
		R.Fprintf(Frpt.file, FlowFmt, float64(NodeStats[j].maxLatFlow*UCF(FLOW)))
		R.Fprintf(Frpt.file, FlowFmt, float64(NodeStats[j].maxInflow*UCF(FLOW)))
		R.Fprintf(Frpt.file, L570, days1, hrs1, mins1)
		R.Fprintf(Frpt.file, L571, float64(NodeStats[j].totLatFlow*Vcf))
		R.Fprintf(Frpt.file, L571, float64(NodeInflow[j]*Vcf))
		if M.Fabs(NodeOutflow[j]) < float64(1.0) {
			R.Fprintf(Frpt.file, L572, float64(float64((float64(NodeInflow[j]-NodeOutflow[j]))*Vcf)*float64(1.0e6)), VolUnitsWords2[UnitSystem])
		} else {
			R.Fprintf(Frpt.file, L573, float64(float64((float64(NodeInflow[j]-NodeOutflow[j]))/NodeOutflow[j])*float64(100.0)))
		}
	}
	report_writeLine((L0))
}

// statsrpt.c:403
func writeNodeSurcharge() {
	var (
		j   int32
		n_2 int32
		t   float64
		d1  float64
		d2  float64
	)
	_, _, _, _, _ = j, n_2, t, d1, d2
	j = 0
	n_2 = 0
	t = 0
	d1 = 0
	d2 = 0
	report_writeLine((L0))
	report_writeLine((L574))
	report_writeLine((L575))
	report_writeLine((L574))
	report_writeLine((L0))
	for j = 0; j < Nobjects[NODE]; j++ {
		if Node[j].type_ == OUTFALL {
			continue
		}
		if NodeStats[j].timeSurcharged == float64(0.0) {
			continue
		}
		if (float64(0.01)) >= (float64(NodeStats[j].timeSurcharged / float64(3600.0))) {
			t = (float64(0.01))
		} else {
			t = (float64(NodeStats[j].timeSurcharged / float64(3600.0)))
		}
		if n_2 == 0 {
			report_writeLine((L576))
			R.Fprintf(Frpt.file, L577)
			if UnitSystem == US {
				R.Fprintf(Frpt.file, L578)
			} else {
				R.Fprintf(Frpt.file, L579)
			}
			R.Fprintf(Frpt.file, L580)
			n_2 = 1
		}
		R.Fprintf(Frpt.file, L541, Node[j].ID)
		R.Fprintf(Frpt.file, L569, NodeTypeWords[Node[j].type_])
		d1 = float64(float64(NodeStats[j].maxDepth+Node[j].invertElev) - Node[j].crownElev)
		if d1 < float64(0.0) {
			d1 = float64(0.0)
		}
		d2 = float64(Node[j].fullDepth - NodeStats[j].maxDepth)
		if d2 < float64(0.0) {
			d2 = float64(0.0)
		}
		R.Fprintf(Frpt.file, L581, t, float64(d1*UCF(LENGTH)), float64(d2*UCF(LENGTH)))
	}
	if n_2 == 0 {
		report_writeLine((L582))
	}
	report_writeLine((L0))
}

// statsrpt.c:449
func writeNodeFlooding() {
	var (
		j    int32
		n_2  int32
		days int32
		hrs  int32
		mins int32
		t    float64
	)
	_, _, _, _, _, _ = j, n_2, days, hrs, mins, t
	j = 0
	n_2 = 0
	days = 0
	hrs = 0
	mins = 0
	t = 0
	report_writeLine((L0))
	report_writeLine((L255))
	report_writeLine((L583))
	report_writeLine((L255))
	report_writeLine((L0))
	for j = 0; j < Nobjects[NODE]; j++ {
		if Node[j].type_ == OUTFALL {
			continue
		}
		if NodeStats[j].timeFlooded == float64(0.0) {
			continue
		}
		if (float64(0.01)) >= (float64(NodeStats[j].timeFlooded / float64(3600.0))) {
			t = (float64(0.01))
		} else {
			t = (float64(NodeStats[j].timeFlooded / float64(3600.0)))
		}
		if n_2 == 0 {
			report_writeLine((L584))
			R.Fprintf(Frpt.file, L585)
			if RouteModel == DW {
				R.Fprintf(Frpt.file, L586)
			} else {
				R.Fprintf(Frpt.file, L587)
			}
			R.Fprintf(Frpt.file, L588, FlowUnitWords[FlowUnits], VolUnitsWords[UnitSystem])
			if RouteModel == DW {
				R.Fprintf(Frpt.file, L589, PondingUnitsWords[UnitSystem])
			} else if UnitSystem == US {
				R.Fprintf(Frpt.file, L590)
			} else {
				R.Fprintf(Frpt.file, L591)
			}
			R.Fprintf(Frpt.file, L592)
			n_2 = 1
		}
		R.Fprintf(Frpt.file, L541, Node[j].ID)
		R.Fprintf(Frpt.file, L593, t)
		R.Fprintf(Frpt.file, FlowFmt, float64(NodeStats[j].maxOverflow*UCF(FLOW)))
		getElapsedTime(NodeStats[j].maxOverflowDate, &days, &hrs, &mins)
		R.Fprintf(Frpt.file, L594, days, hrs, mins)
		R.Fprintf(Frpt.file, L573, float64(NodeStats[j].volFlooded*Vcf))
		if RouteModel == DW {
			R.Fprintf(Frpt.file, L195, float64((float64(NodeStats[j].maxDepth-Node[j].fullDepth))*UCF(LENGTH)))
		} else {
			R.Fprintf(Frpt.file, L195, float64(float64(NodeStats[j].maxPondedVol/float64(1000.0))*UCF(VOLUME)))
		}
	}
	if n_2 == 0 {
		report_writeLine((L595))
	}
	report_writeLine((L0))
}

// statsrpt.c:508
func writeStorageVolumes() {
	var (
		j           int32
		k           int32
		days        int32
		hrs         int32
		mins        int32
		avgVol      float64
		maxVol      float64
		pctAvgVol   float64
		pctMaxVol   float64
		pctEvapLoss float64
		pctSeepLoss float64
	)
	_, _, _, _, _, _, _, _, _, _, _ = j, k, days, hrs, mins, avgVol, maxVol, pctAvgVol, pctMaxVol, pctEvapLoss, pctSeepLoss
	j = 0
	k = 0
	days = 0
	hrs = 0
	mins = 0
	avgVol = 0
	maxVol = 0
	pctAvgVol = 0
	pctMaxVol = 0
	pctEvapLoss = 0
	pctSeepLoss = 0
	if Nnodes[STORAGE] > 0 {
		report_writeLine((L0))
		report_writeLine((L574))
		report_writeLine((L596))
		report_writeLine((L574))
		report_writeLine((L0))
		R.Fprintf(Frpt.file, L597)
		if UnitSystem == US {
			R.Fprintf(Frpt.file, L598)
		} else {
			R.Fprintf(Frpt.file, L599)
		}
		R.Fprintf(Frpt.file, L600, FlowUnitWords[FlowUnits])
		R.Fprintf(Frpt.file, L601)
		for j = 0; j < Nobjects[NODE]; j++ {
			if Node[j].type_ != STORAGE {
				continue
			}
			k = Node[j].subIndex
			R.Fprintf(Frpt.file, L541, Node[j].ID)
			avgVol = float64(StorageStats[k].avgVol / float64(ReportStepCount))
			maxVol = StorageStats[k].maxVol
			pctMaxVol = float64(0.0)
			pctAvgVol = float64(0.0)
			if Node[j].fullVolume > float64(0.0) {
				pctAvgVol = float64(float64(avgVol/Node[j].fullVolume) * float64(100.0))
				pctMaxVol = float64(float64(maxVol/Node[j].fullVolume) * float64(100.0))
			}
			pctEvapLoss = float64(0.0)
			pctSeepLoss = float64(0.0)
			if NodeInflow[j] > float64(0.0) {
				pctEvapLoss = float64(float64(StorageStats[k].evapLosses/NodeInflow[j]) * float64(100.0))
				pctSeepLoss = float64(float64(StorageStats[k].exfilLosses/NodeInflow[j]) * float64(100.0))
			}
			R.Fprintf(Frpt.file, L602, float64(float64(avgVol*UCF(VOLUME))/float64(1000.0)), pctAvgVol, pctEvapLoss, pctSeepLoss, float64(float64(maxVol*UCF(VOLUME))/float64(1000.0)), pctMaxVol)
			getElapsedTime(StorageStats[k].maxVolDate, &days, &hrs, &mins)
			R.Fprintf(Frpt.file, L603, days, hrs, mins)
			R.Fprintf(Frpt.file, FlowFmt, float64(StorageStats[k].maxFlow*UCF(FLOW)))
		}
		report_writeLine((L0))
	}
}

// statsrpt.c:575
func writeOutfallLoads() {
	var (
		units        R.CPtr
		i            int32
		j            int32
		k            int32
		p            int32
		x            float64
		outfallCount float64
		flowCount    float64
		flowSum      float64
		freqSum      float64
		volSum       float64
		totals       []float64
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = units, i, j, k, p, x, outfallCount, flowCount, flowSum, freqSum, volSum, totals
	units = R.Buf(15)
	i = 0
	j = 0
	k = 0
	p = 0
	x = 0
	outfallCount = 0
	flowCount = 0
	flowSum = 0
	freqSum = 0
	volSum = 0
	totals = nil
	if Nnodes[OUTFALL] > 0 {
		totals = make([]float64, int(uint64(Nobjects[POLLUT])))
		for p = 0; p < Nobjects[POLLUT]; p++ {
			totals[p] = float64(0.0)
		}
		flowSum = float64(0.0)
		freqSum = float64(0.0)
		volSum = float64(0.0)
		report_writeLine((L0))
		report_writeLine((L604))
		report_writeLine((L605))
		report_writeLine((L604))
		report_writeLine((L0))
		R.Fprintf(Frpt.file, L606)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			R.Fprintf(Frpt.file, L552)
		}
		R.Fprintf(Frpt.file, L607)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			R.Fprintf(Frpt.file, L608)
		}
		R.Fprintf(Frpt.file, L609)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			R.Fprintf(Frpt.file, L423, Pollut[p].ID)
		}
		R.Fprintf(Frpt.file, L610, FlowUnitWords[FlowUnits], FlowUnitWords[FlowUnits], VolUnitsWords[UnitSystem])
		for p = 0; p < Nobjects[POLLUT]; p++ {
			i = UnitSystem
			if Pollut[p].units == COUNT {
				i = 2
			}
			sstrncpy(units, LoadUnitsWords[i], 14)
			R.Fprintf(Frpt.file, L423, units)
		}
		R.Fprintf(Frpt.file, L606)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			R.Fprintf(Frpt.file, L552)
		}
		for j = 0; j < Nobjects[NODE]; j++ {
			if Node[j].type_ != OUTFALL {
				continue
			}
			k = Node[j].subIndex
			flowCount = float64(OutfallStats[k].totalPeriods)
			R.Fprintf(Frpt.file, L541, Node[j].ID)
			x = float64(float64(float64(100.0)*flowCount) / float64(ReportStepCount))
			R.Fprintf(Frpt.file, L611, x)
			freqSum = float64(freqSum + x)
			if flowCount > float64(0) {
				x = float64(float64(OutfallStats[k].avgFlow*UCF(FLOW)) / flowCount)
			} else {
				x = float64(0.0)
			}
			flowSum = float64(flowSum + x)
			R.Fprintf(Frpt.file, L33)
			R.Fprintf(Frpt.file, FlowFmt, x)
			R.Fprintf(Frpt.file, L33)
			R.Fprintf(Frpt.file, FlowFmt, float64(OutfallStats[k].maxFlow*UCF(FLOW)))
			R.Fprintf(Frpt.file, L573, float64(NodeInflow[j]*Vcf))
			volSum = float64(volSum + NodeInflow[j])
			for p = 0; p < Nobjects[POLLUT]; p++ {
				x = float64(float64(OutfallStats[k].totalLoad[p]*float64(28.317)) * Pollut[p].mcf)
				totals[p] = float64(totals[p] + x)
				if Pollut[p].units == COUNT {
					if (x) > float64(0.0) {
						x = M.Log10((x))
					} else {
						x = (x)
					}
				}
				R.Fprintf(Frpt.file, L427, x)
			}
		}
		outfallCount = float64(Nnodes[OUTFALL])
		R.Fprintf(Frpt.file, L606)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			R.Fprintf(Frpt.file, L552)
		}
		R.Fprintf(Frpt.file, L612, float64(freqSum/outfallCount))
		R.Fprintf(Frpt.file, FlowFmt, flowSum)
		R.Fprintf(Frpt.file, L33)
		R.Fprintf(Frpt.file, FlowFmt, float64(MaxOutfallFlow*UCF(FLOW)))
		R.Fprintf(Frpt.file, L573, float64(volSum*Vcf))
		for p = 0; p < Nobjects[POLLUT]; p++ {
			x = totals[p]
			if Pollut[p].units == COUNT {
				if (x) > float64(0.0) {
					x = M.Log10((x))
				} else {
					x = (x)
				}
			}
			R.Fprintf(Frpt.file, L427, x)
		}
		report_writeLine((L0))
		_ = totals
	}
}

// statsrpt.c:691
func writeLinkFlows() {
	var (
		j         int32
		k         int32
		days      int32
		hrs       int32
		mins      int32
		v         float64
		fullDepth float64
	)
	_, _, _, _, _, _, _ = j, k, days, hrs, mins, v, fullDepth
	j = 0
	k = 0
	days = 0
	hrs = 0
	mins = 0
	v = 0
	fullDepth = 0
	if Nobjects[LINK] == 0 {
		return
	}
	report_writeLine((L0))
	report_writeLine((L235))
	report_writeLine((L613))
	report_writeLine((L235))
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L614)
	if UnitSystem == US {
		R.Fprintf(Frpt.file, L615, FlowUnitWords[FlowUnits])
	} else {
		R.Fprintf(Frpt.file, L616, FlowUnitWords[FlowUnits])
	}
	R.Fprintf(Frpt.file, L617)
	for j = 0; j < Nobjects[LINK]; j++ {
		k = Link[j].subIndex
		R.Fprintf(Frpt.file, L541, Link[j].ID)
		if Link[j].xsect.type_ == DUMMY {
			R.Fprintf(Frpt.file, L618)
		} else if Link[j].xsect.type_ == IRREGULAR {
			R.Fprintf(Frpt.file, L619)
		} else {
			R.Fprintf(Frpt.file, L620, LinkTypeWords[Link[j].type_])
		}
		getElapsedTime(LinkStats[j].maxFlowDate, &days, &hrs, &mins)
		R.Fprintf(Frpt.file, FlowFmt, float64(LinkStats[j].maxFlow*UCF(FLOW)))
		R.Fprintf(Frpt.file, L570, days, hrs, mins)
		if (Link[j].type_ == PUMP) && (Link[j].qFull > float64(0.0)) {
			R.Fprintf(Frpt.file, L621)
			R.Fprintf(Frpt.file, L622, float64(LinkStats[j].maxFlow/Link[j].qFull))
			continue
		}
		if Link[j].xsect.type_ == DUMMY {
			continue
		}
		if Link[j].type_ == OUTLET {
			continue
		}
		if Link[j].type_ == CONDUIT {
			v = float64(LinkStats[j].maxVeloc * UCF(LENGTH))
			if v > float64(50.0) {
				R.Fprintf(Frpt.file, L623)
			} else {
				R.Fprintf(Frpt.file, L624, v)
			}
			R.Fprintf(Frpt.file, L622, float64(float64(LinkStats[j].maxFlow/Link[j].qFull)/float64(Conduit[k].barrels)))
		} else {
			R.Fprintf(Frpt.file, L625)
		}
		fullDepth = Link[j].xsect.yFull
		if (Link[j].type_ == ORIFICE) && (Orifice[k].type_ == BOTTOM_ORIFICE) {
			fullDepth = float64(0.0)
		}
		if fullDepth > float64(0.0) {
			R.Fprintf(Frpt.file, L622, float64(LinkStats[j].maxDepth/fullDepth))
		} else {
			R.Fprintf(Frpt.file, L626)
		}
	}
	report_writeLine((L0))
}

// statsrpt.c:778
func writeFlowClass() {
	var (
		i            int32
		j            int32
		k            int32
		totalSeconds float64
	)
	_, _, _, _ = i, j, k, totalSeconds
	i = 0
	j = 0
	k = 0
	totalSeconds = RoutingTimeSpan
	if RouteModel != DW {
		return
	}
	report_writeLine((L0))
	report_writeLine((L473))
	report_writeLine((L627))
	report_writeLine((L473))
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L628)
	for j = 0; j < Nobjects[LINK]; j++ {
		if Link[j].type_ != CONDUIT {
			continue
		}
		if Link[j].xsect.type_ == DUMMY {
			continue
		}
		k = Link[j].subIndex
		R.Fprintf(Frpt.file, L541, Link[j].ID)
		R.Fprintf(Frpt.file, L629, float64(Conduit[k].modLength/Conduit[k].length))
		for i = 0; i < MAX_FLOW_CLASSES; i++ {
			R.Fprintf(Frpt.file, L630, R.Set(&LinkStats[j].timeInFlowClass[i], float64(LinkStats[j].timeInFlowClass[i]/totalSeconds)))
		}
		R.Fprintf(Frpt.file, L630, float64(LinkStats[j].timeNormalFlow/totalSeconds))
		R.Fprintf(Frpt.file, L630, float64(LinkStats[j].timeInletControl/totalSeconds))
	}
	report_writeLine((L0))
}

// statsrpt.c:821
func writeLinkSurcharge() {
	var (
		i   int32
		j   int32
		n_2 int32
		t   [5]float64
	)
	_, _, _, _ = i, j, n_2, t
	i = 0
	j = 0
	n_2 = 0
	t = ([5]float64{})
	report_writeLine((L0))
	report_writeLine((L470))
	report_writeLine((L631))
	report_writeLine((L470))
	report_writeLine((L0))
	for j = 0; j < Nobjects[LINK]; j++ {
		if (Link[j].type_ != CONDUIT) || (Link[j].xsect.type_ == DUMMY) {
			continue
		}
		t[0] = float64(LinkStats[j].timeSurcharged / float64(3600.0))
		t[1] = float64(LinkStats[j].timeFullUpstream / float64(3600.0))
		t[2] = float64(LinkStats[j].timeFullDnstream / float64(3600.0))
		t[3] = float64(LinkStats[j].timeFullFlow / float64(3600.0))
		if float64(float64(float64(t[0]+t[1])+t[2])+t[3]) == float64(0.0) {
			continue
		}
		t[4] = float64(LinkStats[j].timeCapacityLimited / float64(3600.0))
		for i = 0; i < 5; i++ {
			if (float64(0.01)) >= (t[i]) {
				t[i] = (float64(0.01))
			} else {
				t[i] = (t[i])
			}
		}
		if n_2 == 0 {
			R.Fprintf(Frpt.file, L632)
			n_2 = 1
		}
		R.Fprintf(Frpt.file, L541, Link[j].ID)
		R.Fprintf(Frpt.file, L633, t[0], t[1], t[2], t[3], t[4])
	}
	if n_2 == 0 {
		report_writeLine((L634))
	}
	report_writeLine((L0))
}

// statsrpt.c:862
func writePumpFlows() {
	var (
		j            int32
		k            int32
		avgFlow      float64
		pctUtilized  float64
		pctOffCurve1 float64
		pctOffCurve2 float64
		totalSeconds float64
	)
	_, _, _, _, _, _, _ = j, k, avgFlow, pctUtilized, pctOffCurve1, pctOffCurve2, totalSeconds
	j = 0
	k = 0
	avgFlow = 0
	pctUtilized = 0
	pctOffCurve1 = 0
	pctOffCurve2 = 0
	totalSeconds = RoutingTimeSpan
	if Nlinks[PUMP] == 0 {
		return
	}
	report_writeLine((L0))
	report_writeLine((L221))
	report_writeLine((L635))
	report_writeLine((L221))
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L636, FlowUnitWords[FlowUnits], FlowUnitWords[FlowUnits], FlowUnitWords[FlowUnits], VolUnitsWords[UnitSystem])
	for j = 0; j < Nobjects[LINK]; j++ {
		if Link[j].type_ != PUMP {
			continue
		}
		k = Link[j].subIndex
		R.Fprintf(Frpt.file, L541, Link[j].ID)
		pctUtilized = float64(float64(PumpStats[k].utilized/totalSeconds) * float64(100.0))
		avgFlow = PumpStats[k].avgFlow
		if PumpStats[k].totalPeriods > 0 {
			avgFlow = float64(avgFlow / (float64(PumpStats[k].totalPeriods)))
		}
		R.Fprintf(Frpt.file, L637, pctUtilized, PumpStats[k].startUps, float64(PumpStats[k].minFlow*UCF(FLOW)), float64(avgFlow*UCF(FLOW)), float64(PumpStats[k].maxFlow*UCF(FLOW)), float64(PumpStats[k].volume*Vcf), PumpStats[k].energy)
		pctOffCurve1 = PumpStats[k].offCurveLow
		pctOffCurve2 = PumpStats[k].offCurveHigh
		if PumpStats[k].utilized > float64(0.0) {
			pctOffCurve1 = float64(float64(pctOffCurve1/PumpStats[k].utilized) * float64(100.0))
			pctOffCurve2 = float64(float64(pctOffCurve2/PumpStats[k].utilized) * float64(100.0))
		}
		R.Fprintf(Frpt.file, L638, pctOffCurve1, pctOffCurve2)
	}
	report_writeLine((L0))
}

// statsrpt.c:916
func writeLinkLoads() {
	var (
		i          int32
		j          int32
		p          int32
		x          float64
		units      R.CPtr
		linkLine   R.CPtr
		pollutLine R.CPtr
	)
	_, _, _, _, _, _, _ = i, j, p, x, units, linkLine, pollutLine
	i = 0
	j = 0
	p = 0
	x = 0
	units = R.Buf(15)
	linkLine = R.BufFrom(L551, 21)
	pollutLine = R.BufFrom(L552, 15)
	report_writeLine((L0))
	report_writeLine((L473))
	report_writeLine((L639))
	report_writeLine((L473))
	report_writeLine((L0))
	R.Fprintf(Frpt.file, L356, linkLine)
	for p = 0; p < Nobjects[POLLUT]; p++ {
		R.Fprintf(Frpt.file, L169, pollutLine)
	}
	R.Fprintf(Frpt.file, L555)
	for p = 0; p < Nobjects[POLLUT]; p++ {
		R.Fprintf(Frpt.file, L423, Pollut[p].ID)
	}
	R.Fprintf(Frpt.file, L640)
	for p = 0; p < Nobjects[POLLUT]; p++ {
		i = UnitSystem
		if Pollut[p].units == COUNT {
			i = 2
		}
		sstrncpy(units, LoadUnitsWords[i], 14)
		R.Fprintf(Frpt.file, L423, units)
	}
	R.Fprintf(Frpt.file, L356, linkLine)
	for p = 0; p < Nobjects[POLLUT]; p++ {
		R.Fprintf(Frpt.file, L169, pollutLine)
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		R.Fprintf(Frpt.file, L541, Link[j].ID)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			x = float64(float64(Link[j].totalLoad[p]*float64(28.317)) * Pollut[p].mcf)
			if Pollut[p].units == COUNT {
				if (x) > float64(0.0) {
					x = M.Log10((x))
				} else {
					x = (x)
				}
			}
			if x < float64(10000.0) {
				R.Fprintf(Frpt.file, L427, x)
			} else {
				R.Fprintf(Frpt.file, L641, x)
			}
		}
	}
	report_writeLine((L0))
}
