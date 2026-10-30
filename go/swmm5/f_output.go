// Generated from EPA SWMM 5.2.4 (output.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var IDStartPos int64
var InputStartPos int64
var OutputStartPos int64
var BytesPerPeriod int64
var NumSubcatchVars int32
var NumNodeVars int32
var NumLinkVars int32
var NumSubcatch int32
var NumNodes int32
var NumLinks int32
var NumPolluts int32
var SysResults [15]float32
var AvgLinkResults []TAvgResults
var AvgNodeResults []TAvgResults
var output_Nsteps int32
var SubcatchResults []float32
var NodeResults []float32
var LinkResults []float32

// output.c:121
func output_open() int32 {
	var (
		j          int32
		m          int32
		k          int32
		x          float32
		z          float64
		numResults int64
	)
	_, _, _, _, _, _ = j, m, k, x, z, numResults
	j = 0
	m = 0
	k = 0
	x = 0
	z = 0
	numResults = 0
	output_openOutFile()
	if ErrorCode != 0 {
		return ErrorCode
	}
	if IgnoreQuality != 0 {
		NumPolluts = 0
	} else {
		NumPolluts = Nobjects[POLLUT]
	}
	NumSubcatchVars = ((9 - 1) + NumPolluts)
	NumNodeVars = ((7 - 1) + NumPolluts)
	NumLinkVars = ((6 - 1) + NumPolluts)
	NumSubcatch = 0
	NumNodes = 0
	NumLinks = 0
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		if Subcatch[j].rptFlag != 0 {
			NumSubcatch++
		}
	}
	for j = 0; j < Nobjects[NODE]; j++ {
		if Node[j].rptFlag != 0 {
			NumNodes++
		}
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		if Link[j].rptFlag != 0 {
			NumLinks++
		}
	}
	numResults = ((((int64(NumSubcatch) * int64(NumSubcatchVars)) + (int64(NumNodes) * int64(NumNodeVars))) + (int64(NumLinks) * int64(NumLinkVars))) + 15)
	BytesPerPeriod = int64((8 + (uint64(numResults) * 4)))
	Nperiods = 0
	SubcatchResults = nil
	NodeResults = nil
	LinkResults = nil
	SubcatchResults = make([]float32, int(uint64(NumSubcatchVars)))
	NodeResults = make([]float32, int(uint64(NumNodeVars)))
	LinkResults = make([]float32, int(uint64(NumLinkVars)))
	if ((!(SubcatchResults != nil)) || (!(NodeResults != nil))) || (!(LinkResults != nil)) {
		report_writeErrorMsg(ERR_MEMORY, L0)
		return ErrorCode
	}
	AvgNodeResults = nil
	AvgLinkResults = nil
	if (RptFlags.averages != 0) && (!(output_openAvgResults() != 0)) {
		report_writeErrorMsg(ERR_MEMORY, L0)
		return ErrorCode
	}
	R.Fseek(Fout.file, 0, 0)
	k = 516114522
	R.Fwrite1(k, 4, Fout.file)
	k = 52004
	R.Fwrite1(k, 4, Fout.file)
	k = FlowUnits
	R.Fwrite1(k, 4, Fout.file)
	k = NumSubcatch
	R.Fwrite1(k, 4, Fout.file)
	k = NumNodes
	R.Fwrite1(k, 4, Fout.file)
	k = NumLinks
	R.Fwrite1(k, 4, Fout.file)
	k = NumPolluts
	R.Fwrite1(k, 4, Fout.file)
	IDStartPos = R.Ftell(Fout.file)
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		if Subcatch[j].rptFlag != 0 {
			output_saveID(Subcatch[j].ID, Fout.file)
		}
	}
	for j = 0; j < Nobjects[NODE]; j++ {
		if Node[j].rptFlag != 0 {
			output_saveID(Node[j].ID, Fout.file)
		}
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		if Link[j].rptFlag != 0 {
			output_saveID(Link[j].ID, Fout.file)
		}
	}
	for j = 0; j < NumPolluts; j++ {
		output_saveID(Pollut[j].ID, Fout.file)
	}
	for j = 0; j < NumPolluts; j++ {
		k = Pollut[j].units
		R.Fwrite1(k, 4, Fout.file)
	}
	InputStartPos = R.Ftell(Fout.file)
	k = 1
	R.Fwrite1(k, 4, Fout.file)
	k = INPUT_AREA
	R.Fwrite1(k, 4, Fout.file)
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		if !(Subcatch[j].rptFlag != 0) {
			continue
		}
		SubcatchResults[0] = float32(float64((float64(Subcatch[j].area * UCF(LANDAREA)))))
		R.Fwrite(SubcatchResults[0:], 1, Fout.file)
	}
	k = 3
	R.Fwrite1(k, 4, Fout.file)
	k = INPUT_TYPE_CODE
	R.Fwrite1(k, 4, Fout.file)
	k = INPUT_INVERT
	R.Fwrite1(k, 4, Fout.file)
	k = INPUT_MAX_DEPTH
	R.Fwrite1(k, 4, Fout.file)
	for j = 0; j < Nobjects[NODE]; j++ {
		if !(Node[j].rptFlag != 0) {
			continue
		}
		k = Node[j].type_
		NodeResults[0] = float32(float64((float64(Node[j].invertElev * UCF(LENGTH)))))
		NodeResults[1] = float32(float64((float64(Node[j].fullDepth * UCF(LENGTH)))))
		R.Fwrite1(k, 4, Fout.file)
		R.Fwrite(NodeResults, 2, Fout.file)
	}
	k = 5
	R.Fwrite1(k, 4, Fout.file)
	k = INPUT_TYPE_CODE
	R.Fwrite1(k, 4, Fout.file)
	k = INPUT_OFFSET
	R.Fwrite1(k, 4, Fout.file)
	k = INPUT_OFFSET
	R.Fwrite1(k, 4, Fout.file)
	k = INPUT_MAX_DEPTH
	R.Fwrite1(k, 4, Fout.file)
	k = INPUT_LENGTH
	R.Fwrite1(k, 4, Fout.file)
	for j = 0; j < Nobjects[LINK]; j++ {
		if !(Link[j].rptFlag != 0) {
			continue
		}
		k = Link[j].type_
		if k == PUMP {
			for m = 0; m < 4; m++ {
				LinkResults[m] = float32(0.0)
			}
		} else {
			LinkResults[0] = float32(float64((float64(Link[j].offset1 * UCF(LENGTH)))))
			LinkResults[1] = float32(float64((float64(Link[j].offset2 * UCF(LENGTH)))))
			if Link[j].direction < 0 {
				x = LinkResults[0]
				LinkResults[0] = LinkResults[1]
				LinkResults[1] = x
			}
			if k == OUTLET {
				LinkResults[2] = float32(0.0)
			} else {
				LinkResults[2] = float32(float64((float64(Link[j].xsect.yFull * UCF(LENGTH)))))
			}
			if k == CONDUIT {
				m = Link[j].subIndex
				LinkResults[3] = float32(float64((float64(Conduit[m].length * UCF(LENGTH)))))
			} else {
				LinkResults[3] = float32(0.0)
			}
		}
		R.Fwrite1(k, 4, Fout.file)
		R.Fwrite(LinkResults, 4, Fout.file)
	}
	k = NumSubcatchVars
	R.Fwrite1(k, 4, Fout.file)
	k = SUBCATCH_RAINFALL
	R.Fwrite1(k, 4, Fout.file)
	k = SUBCATCH_SNOWDEPTH
	R.Fwrite1(k, 4, Fout.file)
	k = SUBCATCH_EVAP
	R.Fwrite1(k, 4, Fout.file)
	k = SUBCATCH_INFIL
	R.Fwrite1(k, 4, Fout.file)
	k = SUBCATCH_RUNOFF
	R.Fwrite1(k, 4, Fout.file)
	k = SUBCATCH_GW_FLOW
	R.Fwrite1(k, 4, Fout.file)
	k = SUBCATCH_GW_ELEV
	R.Fwrite1(k, 4, Fout.file)
	k = SUBCATCH_SOIL_MOIST
	R.Fwrite1(k, 4, Fout.file)
	for j = 0; j < NumPolluts; j++ {
		k = (SUBCATCH_WASHOFF + j)
		R.Fwrite1(k, 4, Fout.file)
	}
	k = NumNodeVars
	R.Fwrite1(k, 4, Fout.file)
	k = NODE_DEPTH
	R.Fwrite1(k, 4, Fout.file)
	k = NODE_HEAD
	R.Fwrite1(k, 4, Fout.file)
	k = NODE_VOLUME
	R.Fwrite1(k, 4, Fout.file)
	k = NODE_LATFLOW
	R.Fwrite1(k, 4, Fout.file)
	k = NODE_INFLOW
	R.Fwrite1(k, 4, Fout.file)
	k = NODE_OVERFLOW
	R.Fwrite1(k, 4, Fout.file)
	for j = 0; j < NumPolluts; j++ {
		k = (NODE_QUAL + j)
		R.Fwrite1(k, 4, Fout.file)
	}
	k = NumLinkVars
	R.Fwrite1(k, 4, Fout.file)
	k = LINK_FLOW
	R.Fwrite1(k, 4, Fout.file)
	k = LINK_DEPTH
	R.Fwrite1(k, 4, Fout.file)
	k = LINK_VELOCITY
	R.Fwrite1(k, 4, Fout.file)
	k = LINK_VOLUME
	R.Fwrite1(k, 4, Fout.file)
	k = LINK_CAPACITY
	R.Fwrite1(k, 4, Fout.file)
	for j = 0; j < NumPolluts; j++ {
		k = (LINK_QUAL + j)
		R.Fwrite1(k, 4, Fout.file)
	}
	k = 15
	R.Fwrite1(k, 4, Fout.file)
	for k = 0; k < 15; k++ {
		R.Fwrite1(k, 4, Fout.file)
	}
	z = float64(float64(ReportStep) / float64(86400.0))
	if float64(StartDateTime+z) > ReportStart {
		z = StartDateTime
	} else {
		z = float64(M.Floor(float64((float64(ReportStart-StartDateTime))/z)) - float64(1.0))
		z = float64(StartDateTime + float64(float64(z*float64(ReportStep))/float64(86400.0)))
	}
	R.Fwrite1(z, 8, Fout.file)
	k = ReportStep
	if R.Fwrite1(k, 4, Fout.file) < 1 {
		report_writeErrorMsg(ERR_OUT_WRITE, L0)
		return ErrorCode
	}
	OutputStartPos = R.Ftell(Fout.file)
	return ErrorCode
}

// output.c:427
func output_openOutFile() {
	if Fout.file != nil {
		R.Fclose(Fout.file)
	} else if R.Strlen(R.CP(Fout.name[:])) != 0 {
		Fout.mode = SAVE_FILE
	} else {
		Fout.mode = SCRATCH_FILE
		getTempFileName(R.CP(Fout.name[:]))
	}
	if (R.Set(&Fout.file, R.Fopen(R.CP(Fout.name[:]), L166))) == nil {
		writecon(L318)
		ErrorCode = ERR_OUT_FILE
	}
}

// output.c:457
func output_saveResults(reportTime float64) {
	var (
		i          int32
		reportDate float64
		date       float64
	)
	_, _, _ = i, reportDate, date
	i = 0
	reportDate = getDateTime(reportTime)
	date = 0
	if reportDate < ReportStart {
		return
	}
	for i = 0; i < 15; i++ {
		SysResults[i] = float32(0.0)
	}
	date = reportDate
	R.Fwrite1(date, 8, Fout.file)
	if Nobjects[SUBCATCH] > 0 {
		output_saveSubcatchResults(reportTime, Fout.file)
	}
	if RptFlags.averages != 0 {
		output_saveAvgResults(Fout.file)
	} else {
		if Nobjects[NODE] > 0 {
			output_saveNodeResults(reportTime, Fout.file)
		}
		if Nobjects[LINK] > 0 {
			output_saveLinkResults(reportTime, Fout.file)
		}
	}
	SysResults[SYS_FLOODING] = float32(float64((float64(StepFlowTotals.flooding * UCF(FLOW)))))
	SysResults[SYS_OUTFLOW] = float32(float64((float64(StepFlowTotals.outflow * UCF(FLOW)))))
	SysResults[SYS_DWFLOW] = float32(float64((float64(StepFlowTotals.dwInflow * UCF(FLOW)))))
	SysResults[SYS_GWFLOW] = float32(float64((float64(StepFlowTotals.gwInflow * UCF(FLOW)))))
	SysResults[SYS_IIFLOW] = float32(float64((float64(StepFlowTotals.iiInflow * UCF(FLOW)))))
	SysResults[SYS_EXFLOW] = float32(float64((float64(StepFlowTotals.exInflow * UCF(FLOW)))))
	SysResults[SYS_INFLOW] = float32(float32(float32(float32(SysResults[SYS_RUNOFF]+SysResults[SYS_DWFLOW])+SysResults[SYS_GWFLOW])+SysResults[SYS_IIFLOW]) + SysResults[SYS_EXFLOW])
	R.Fwrite(SysResults[:], 15, Fout.file)
	if (Foutflows.mode == SAVE_FILE) && (!(IgnoreRouting != 0)) {
		iface_saveOutletResults(reportDate, Foutflows.file)
	}
	Nperiods++
}

// output.c:515
func output_end() {
	var (
		k int32
	)
	_ = k
	k = 0
	R.Fwrite1(IDStartPos, 4, Fout.file)
	R.Fwrite1(InputStartPos, 4, Fout.file)
	R.Fwrite1(OutputStartPos, 4, Fout.file)
	k = int32(Nperiods)
	R.Fwrite1(k, 4, Fout.file)
	k = ErrorCode
	R.Fwrite1(k, 4, Fout.file)
	k = 516114522
	if R.Fwrite1(k, 4, Fout.file) < 1 {
		report_writeErrorMsg(ERR_OUT_WRITE, L0)
	}
}

// output.c:539
func output_close() {
	{
		if SubcatchResults != nil {
			_ = SubcatchResults
			SubcatchResults = nil
		}
	}
	{
		if NodeResults != nil {
			_ = NodeResults
			NodeResults = nil
		}
	}
	{
		if LinkResults != nil {
			_ = LinkResults
			LinkResults = nil
		}
	}
	output_closeAvgResults()
}

// output.c:554
func output_saveID(id R.CPtr, file *R.CFile) {
	var (
		n_2 int32
	)
	_ = n_2
	n_2 = int32(R.Strlen(id))
	R.Fwrite1(n_2, 4, file)
	R.FwriteC(id, uint64(n_2), file)
}

// output.c:569
func output_saveSubcatchResults(reportTime float64, file *R.CFile) {
	var (
		j          int32
		f          float64
		area       float64
		totalArea  float32
		reportDate float64
	)
	_, _, _, _, _ = j, f, area, totalArea, reportDate
	j = 0
	f = 0
	area = 0
	totalArea = float32(0.0)
	reportDate = getDateTime(reportTime)
	for j = 0; j < Nobjects[GAGE]; j++ {
		gage_setReportRainfall(j, reportDate)
	}
	f = float64((float64(reportTime - OldRunoffTime)) / (float64(NewRunoffTime - OldRunoffTime)))
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		subcatch_getResults(j, f, SubcatchResults)
		if Subcatch[j].rptFlag != 0 {
			R.Fwrite(SubcatchResults, uint64(NumSubcatchVars), file)
		}
		area = float64(Subcatch[j].area * UCF(LANDAREA))
		totalArea = float32(totalArea + (float32(float64(area))))
		SysResults[SYS_RAINFALL] = float32(SysResults[SYS_RAINFALL] + (float32(float64((float64(float64(SubcatchResults[SUBCATCH_RAINFALL]) * area))))))
		SysResults[SYS_SNOWDEPTH] = float32(SysResults[SYS_SNOWDEPTH] + (float32(float64((float64(float64(SubcatchResults[SUBCATCH_SNOWDEPTH]) * area))))))
		SysResults[SYS_EVAP] = float32(SysResults[SYS_EVAP] + (float32(float64((float64(float64(SubcatchResults[SUBCATCH_EVAP]) * area))))))
		if Subcatch[j].groundwater != nil {
			SysResults[SYS_EVAP] = float32(SysResults[SYS_EVAP] + (float32(float64((float64(float64(Subcatch[j].groundwater.evapLoss*UCF(EVAPRATE)) * area))))))
		}
		SysResults[SYS_INFIL] = float32(SysResults[SYS_INFIL] + (float32(float64((float64(float64(SubcatchResults[SUBCATCH_INFIL]) * area))))))
		SysResults[SYS_RUNOFF] = float32(SysResults[SYS_RUNOFF] + SubcatchResults[SUBCATCH_RUNOFF])
	}
	if float64(totalArea) > float64(0.0) {
		SysResults[SYS_EVAP] = float32(SysResults[SYS_EVAP] / totalArea)
		SysResults[SYS_RAINFALL] = float32(SysResults[SYS_RAINFALL] / totalArea)
		SysResults[SYS_SNOWDEPTH] = float32(SysResults[SYS_SNOWDEPTH] / totalArea)
		SysResults[SYS_INFIL] = float32(SysResults[SYS_INFIL] / totalArea)
	}
	if UnitSystem == SI {
		f = float64((float64(float64(5.0) / float64(9.0))) * (float64(Temp.ta - float64(32.0))))
	} else {
		f = Temp.ta
	}
	SysResults[SYS_TEMPERATURE] = float32(float64(f))
	f = float64(Evap.rate * UCF(EVAPRATE))
	SysResults[SYS_PET] = float32(float64(f))
}

// output.c:636
func output_saveNodeResults(reportTime float64, file *R.CFile) {
	var (
		j int32
		f float64
	)
	_, _ = j, f
	j = 0
	f = float64((float64(reportTime - OldRoutingTime)) / (float64(NewRoutingTime - OldRoutingTime)))
	for j = 0; j < Nobjects[NODE]; j++ {
		node_getResults(j, f, NodeResults)
		if Node[j].rptFlag != 0 {
			R.Fwrite(NodeResults, uint64(NumNodeVars), file)
		}
		stats_updateMaxNodeDepth(j, float64(NodeResults[NODE_DEPTH]))
		SysResults[SYS_STORAGE] = float32(SysResults[SYS_STORAGE] + NodeResults[NODE_VOLUME])
	}
}

// output.c:666
func output_saveLinkResults(reportTime float64, file *R.CFile) {
	var (
		j int32
		f float64
		z float64
	)
	_, _, _ = j, f, z
	j = 0
	f = 0
	z = 0
	f = float64((float64(reportTime - OldRoutingTime)) / (float64(NewRoutingTime - OldRoutingTime)))
	for j = 0; j < Nobjects[LINK]; j++ {
		if Link[j].rptFlag != 0 {
			link_getResults(j, f, LinkResults)
			R.Fwrite(LinkResults, uint64(NumLinkVars), file)
		}
		z = float64((float64(float64((float64(float64(1.0)-f))*Link[j].oldVolume) + float64(f*Link[j].newVolume))) * UCF(VOLUME))
		SysResults[SYS_STORAGE] = float32(SysResults[SYS_STORAGE] + (float32(float64(z))))
	}
}

// output.c:699
func output_readDateTime(period int64, days *float64) {
	var (
		p       int64
		bytePos int64
	)
	_, _ = p, bytePos
	p = period
	bytePos = (OutputStartPos + ((p - 1) * BytesPerPeriod))
	R.Fseek(Fout.file, bytePos, 0)
	*days = float64((-693594))
	R.Fread1(days, 8, Fout.file)
}

// output.c:716
func output_readSubcatchResults(period int64, index int32) {
	var (
		offset  int64
		p       int64
		bytePos int64
	)
	_, _, _ = offset, p, bytePos
	offset = int64((index * NumSubcatchVars))
	p = period
	bytePos = int64(((uint64((OutputStartPos + ((p - 1) * BytesPerPeriod))) + 8) + (uint64(offset) * 4)))
	R.Fseek(Fout.file, bytePos, 0)
	R.Fread(SubcatchResults, uint64(NumSubcatchVars), Fout.file)
}

// output.c:735
func output_readNodeResults(period int64, index int32) {
	var (
		offset  int64
		p       int64
		bytePos int64
	)
	_, _, _ = offset, p, bytePos
	offset = int64(((NumSubcatch * NumSubcatchVars) + (index * NumNodeVars)))
	p = period
	bytePos = int64(((uint64((OutputStartPos + ((p - 1) * BytesPerPeriod))) + 8) + (uint64(offset) * 4)))
	R.Fseek(Fout.file, bytePos, 0)
	R.Fread(NodeResults, uint64(NumNodeVars), Fout.file)
}

// output.c:753
func output_readLinkResults(period int64, index int32) {
	var (
		offset  int64
		p       int64
		bytePos int64
	)
	_, _, _ = offset, p, bytePos
	offset = int64((((NumSubcatch * NumSubcatchVars) + (NumNodes * NumNodeVars)) + (index * NumLinkVars)))
	p = period
	bytePos = int64(((uint64((OutputStartPos + ((p - 1) * BytesPerPeriod))) + 8) + (uint64(offset) * 4)))
	R.Fseek(Fout.file, bytePos, 0)
	R.Fread(LinkResults, uint64(NumLinkVars), Fout.file)
	R.Fread(SysResults[:], 15, Fout.file)
}

// output.c:774
func output_openAvgResults() int32 {
	var (
		i int32
	)
	_ = i
	i = 0
	AvgNodeResults = make([]TAvgResults, int(uint64(NumNodes)))
	if AvgNodeResults == nil {
		return 0
	}
	for i = 0; i < NumNodes; i++ {
		AvgNodeResults[i].xAvg = nil
	}
	AvgLinkResults = make([]TAvgResults, int(uint64(NumLinks)))
	if AvgLinkResults == nil {
		output_closeAvgResults()
		return 0
	}
	for i = 0; i < NumLinks; i++ {
		AvgLinkResults[i].xAvg = nil
	}
	for i = 0; i < NumNodes; i++ {
		AvgNodeResults[i].xAvg = make([]float32, int(uint64(NumNodeVars)))
		if AvgNodeResults[i].xAvg == nil {
			output_closeAvgResults()
			return 0
		}
	}
	for i = 0; i < NumLinks; i++ {
		AvgLinkResults[i].xAvg = make([]float32, int(uint64(NumLinkVars)))
		if AvgLinkResults[i].xAvg == nil {
			output_closeAvgResults()
			return 0
		}
	}
	return 1
}

// output.c:820
func output_closeAvgResults() {
	var (
		i int32
	)
	_ = i
	i = 0
	if AvgNodeResults != nil {
		for i = 0; i < NumNodes; i++ {
			if AvgNodeResults[i].xAvg != nil {
				_ = AvgNodeResults[i].xAvg
				AvgNodeResults[i].xAvg = nil
			}
		}
		{
			if AvgNodeResults != nil {
				_ = AvgNodeResults
				AvgNodeResults = nil
			}
		}
	}
	if AvgLinkResults != nil {
		for i = 0; i < NumLinks; i++ {
			if AvgLinkResults[i].xAvg != nil {
				_ = AvgLinkResults[i].xAvg
				AvgLinkResults[i].xAvg = nil
			}
		}
		{
			if AvgLinkResults != nil {
				_ = AvgLinkResults
				AvgLinkResults = nil
			}
		}
	}
}

// output.c:839
func output_initAvgResults() {
	var (
		i int32
		j int32
	)
	_, _ = i, j
	i = 0
	j = 0
	output_Nsteps = 0
	for i = 0; i < NumNodes; i++ {
		for j = 0; j < NumNodeVars; j++ {
			AvgNodeResults[i].xAvg[j] = float32(float64(float64(0.0)))
		}
	}
	for i = 0; i < NumLinks; i++ {
		for j = 0; j < NumLinkVars; j++ {
			AvgLinkResults[i].xAvg[j] = float32(float64(float64(0.0)))
		}
	}
}

// output.c:857
func output_updateAvgResults() {
	var (
		i    int32
		j    int32
		k    int32
		sign int32
	)
	_, _, _, _ = i, j, k, sign
	i = 0
	j = 0
	k = 0
	sign = 0
	k = 0
	for i = 0; i < Nobjects[NODE]; i++ {
		if !(Node[i].rptFlag != 0) {
			continue
		}
		node_getResults(i, float64(1.0), NodeResults)
		for j = 0; j < NumNodeVars; j++ {
			AvgNodeResults[k].xAvg[j] = float32(AvgNodeResults[k].xAvg[j] + NodeResults[j])
		}
		k++
	}
	k = 0
	for i = 0; i < Nobjects[LINK]; i++ {
		if !(Link[i].rptFlag != 0) {
			continue
		}
		link_getResults(i, float64(1.0), LinkResults)
		if (LinkResults[LINK_FLOW]) < float32(0) {
			sign = (-1)
		} else {
			sign = (1)
		}
		for j = 0; j < NumLinkVars; j++ {
			if j == LINK_CAPACITY {
				if Link[i].type_ == CONDUIT {
					AvgLinkResults[k].xAvg[j] = float32(AvgLinkResults[k].xAvg[j] + LinkResults[j])
				} else {
					AvgLinkResults[k].xAvg[j] = float32(LinkResults[j] * float32((output_Nsteps + 1)))
				}
			} else {
				AvgLinkResults[k].xAvg[j] = float32(AvgLinkResults[k].xAvg[j] + LinkResults[j])
			}
		}
		k++
	}
	output_Nsteps++
}

// output.c:911
func output_saveAvgResults(file *R.CFile) {
	var (
		i int32
		j int32
	)
	_, _ = i, j
	i = 0
	j = 0
	for i = 0; i < NumNodes; i++ {
		for j = 0; j < NumNodeVars; j++ {
			NodeResults[j] = float32(AvgNodeResults[i].xAvg[j] / float32(output_Nsteps))
		}
		R.Fwrite(NodeResults, uint64(NumNodeVars), file)
	}
	for i = 0; i < Nobjects[NODE]; i++ {
		stats_updateMaxNodeDepth(i, float64(Node[i].newDepth*UCF(LENGTH)))
		SysResults[SYS_STORAGE] = float32(SysResults[SYS_STORAGE] + (float32(float64((float64(Node[i].newVolume * UCF(VOLUME)))))))
	}
	for i = 0; i < NumLinks; i++ {
		for j = 0; j < NumLinkVars; j++ {
			LinkResults[j] = float32(AvgLinkResults[i].xAvg[j] / float32(output_Nsteps))
		}
		R.Fwrite(LinkResults, uint64(NumLinkVars), file)
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		SysResults[SYS_STORAGE] = float32(SysResults[SYS_STORAGE] + (float32(float64((float64(Link[i].newVolume * UCF(VOLUME)))))))
	}
	output_initAvgResults()
}
