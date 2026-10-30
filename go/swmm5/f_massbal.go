// Generated from EPA SWMM 5.2.4 (massbal.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	M "swmm5go/glibm"
)

var MAX_RUNOFF_BALANCE_ERR float64 = float64(10.0)
var MAX_FLOW_BALANCE_ERR float64 = float64(10.0)
var RunoffTotals TRunoffTotals
var LoadingTotals []TLoadingTotals
var GwaterTotals TGwaterTotals
var FlowTotals TRoutingTotals
var QualTotals []TRoutingTotals
var StepFlowTotals TRoutingTotals
var OldStepFlowTotals TRoutingTotals
var StepQualTotals []TRoutingTotals
var NodeInflow []float64
var NodeOutflow []float64
var TotalArea float64

// massbal.c:100
func massbal_open() int32 {
	var (
		j   int32
		n_2 int32
	)
	_, _ = j, n_2
	j = 0
	n_2 = 0
	RunoffError = float64(0.0)
	GwaterError = float64(0.0)
	FlowError = float64(0.0)
	QualError = float64(0.0)
	RunoffTotals.rainfall = float64(0.0)
	RunoffTotals.evap = float64(0.0)
	RunoffTotals.infil = float64(0.0)
	RunoffTotals.runoff = float64(0.0)
	RunoffTotals.runon = float64(0.0)
	RunoffTotals.drains = float64(0.0)
	RunoffTotals.snowRemoved = float64(0.0)
	RunoffTotals.initStorage = float64(0.0)
	RunoffTotals.initSnowCover = float64(0.0)
	TotalArea = float64(0.0)
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		RunoffTotals.initStorage = float64(RunoffTotals.initStorage + (subcatch_getStorage(j)))
		RunoffTotals.initSnowCover = float64(RunoffTotals.initSnowCover + (snow_getSnowCover(j)))
		TotalArea = float64(TotalArea + Subcatch[j].area)
	}
	GwaterTotals.infil = float64(0.0)
	GwaterTotals.upperEvap = float64(0.0)
	GwaterTotals.lowerEvap = float64(0.0)
	GwaterTotals.lowerPerc = float64(0.0)
	GwaterTotals.gwater = float64(0.0)
	GwaterTotals.initStorage = float64(0.0)
	GwaterTotals.finalStorage = float64(0.0)
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		GwaterTotals.initStorage = float64(GwaterTotals.initStorage + (float64(gwater_getVolume(j) * Subcatch[j].area)))
	}
	FlowTotals.dwInflow = float64(0.0)
	FlowTotals.wwInflow = float64(0.0)
	FlowTotals.gwInflow = float64(0.0)
	FlowTotals.iiInflow = float64(0.0)
	FlowTotals.exInflow = float64(0.0)
	FlowTotals.flooding = float64(0.0)
	FlowTotals.outflow = float64(0.0)
	FlowTotals.evapLoss = float64(0.0)
	FlowTotals.seepLoss = float64(0.0)
	FlowTotals.reacted = float64(0.0)
	FlowTotals.initStorage = float64(0.0)
	for j = 0; j < Nobjects[NODE]; j++ {
		FlowTotals.initStorage = float64(FlowTotals.initStorage + Node[j].newVolume)
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		FlowTotals.initStorage = float64(FlowTotals.initStorage + Link[j].newVolume)
	}
	StepFlowTotals = FlowTotals
	LoadingTotals = nil
	QualTotals = nil
	StepQualTotals = nil
	NodeInflow = nil
	NodeOutflow = nil
	n_2 = Nobjects[POLLUT]
	if n_2 > 0 {
		LoadingTotals = make([]TLoadingTotals, int(uint64(n_2)))
		if LoadingTotals == nil {
			report_writeErrorMsg(ERR_MEMORY, L0)
			return ErrorCode
		}
		for j = 0; j < n_2; j++ {
			LoadingTotals[j].initLoad = massbal_getBuildup(j)
			LoadingTotals[j].buildup = float64(0.0)
			LoadingTotals[j].deposition = float64(0.0)
			LoadingTotals[j].sweeping = float64(0.0)
			LoadingTotals[j].infil = float64(0.0)
			LoadingTotals[j].bmpRemoval = float64(0.0)
			LoadingTotals[j].runoff = float64(0.0)
			LoadingTotals[j].finalLoad = float64(0.0)
		}
	}
	if n_2 > 0 {
		QualTotals = make([]TRoutingTotals, int(uint64(n_2)))
		StepQualTotals = make([]TRoutingTotals, int(uint64(n_2)))
		if (QualTotals == nil) || (StepQualTotals == nil) {
			report_writeErrorMsg(ERR_MEMORY, L0)
			return ErrorCode
		}
	}
	for j = 0; j < n_2; j++ {
		QualTotals[j].dwInflow = float64(0.0)
		QualTotals[j].wwInflow = float64(0.0)
		QualTotals[j].gwInflow = float64(0.0)
		QualTotals[j].exInflow = float64(0.0)
		QualTotals[j].flooding = float64(0.0)
		QualTotals[j].outflow = float64(0.0)
		QualTotals[j].evapLoss = float64(0.0)
		QualTotals[j].seepLoss = float64(0.0)
		QualTotals[j].reacted = float64(0.0)
		QualTotals[j].initStorage = massbal_getStoredMass(j)
	}
	massbal_initTimeStepTotals()
	if Nobjects[NODE] > 0 {
		NodeInflow = make([]float64, int(uint64(Nobjects[NODE])))
		if NodeInflow == nil {
			report_writeErrorMsg(ERR_MEMORY, L0)
			return ErrorCode
		}
		NodeOutflow = make([]float64, int(uint64(Nobjects[NODE])))
		if NodeOutflow == nil {
			report_writeErrorMsg(ERR_MEMORY, L0)
			return ErrorCode
		}
		for j = 0; j < Nobjects[NODE]; j++ {
			NodeInflow[j] = Node[j].newVolume
		}
	}
	return ErrorCode
}

// massbal.c:246
func massbal_close() {
	{
		if LoadingTotals != nil {
			_ = LoadingTotals
			LoadingTotals = nil
		}
	}
	{
		if QualTotals != nil {
			_ = QualTotals
			QualTotals = nil
		}
	}
	{
		if StepQualTotals != nil {
			_ = StepQualTotals
			StepQualTotals = nil
		}
	}
	{
		if NodeInflow != nil {
			_ = NodeInflow
			NodeInflow = nil
		}
	}
	{
		if NodeOutflow != nil {
			_ = NodeOutflow
			NodeOutflow = nil
		}
	}
}

// massbal.c:262
func massbal_report() {
	var (
		j      int32
		gwArea float64
	)
	_, _ = j, gwArea
	j = 0
	gwArea = float64(0.0)
	if Nobjects[SUBCATCH] > 0 {
		if (massbal_getRunoffError() > MAX_RUNOFF_BALANCE_ERR) || (RptFlags.continuity == 1) {
			report_writeRunoffError(&RunoffTotals, TotalArea)
		}
		if (Nobjects[POLLUT] > 0) && (!(IgnoreQuality != 0)) {
			if (massbal_getLoadingError() > MAX_RUNOFF_BALANCE_ERR) || (RptFlags.continuity == 1) {
				report_writeLoadingError(LoadingTotals)
			}
		}
	}
	if (Nobjects[AQUIFER] > 0) && (!(IgnoreGwater != 0)) {
		if (massbal_getGwaterError() > MAX_RUNOFF_BALANCE_ERR) || (RptFlags.continuity == 1) {
			for j = 0; j < Nobjects[SUBCATCH]; j++ {
				if Subcatch[j].groundwater != nil {
					gwArea = float64(gwArea + Subcatch[j].area)
				}
			}
			if gwArea > float64(0.0) {
				report_writeGwaterError(&GwaterTotals, gwArea)
			}
		}
	}
	if (Nobjects[NODE] > 0) && (!(IgnoreRouting != 0)) {
		if (massbal_getFlowError() > MAX_FLOW_BALANCE_ERR) || (RptFlags.continuity == 1) {
			report_writeFlowError(&FlowTotals)
		}
		if (Nobjects[POLLUT] > 0) && (!(IgnoreQuality != 0)) {
			if (massbal_getQualError() > MAX_FLOW_BALANCE_ERR) || (RptFlags.continuity == 1) {
				report_writeQualError(QualTotals)
			}
		}
	}
}

// massbal.c:316
func massbal_getBuildup(p int32) float64 {
	var (
		i    int32
		j    int32
		load float64
	)
	_, _, _ = i, j, load
	i = 0
	j = 0
	load = float64(0.0)
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		for i = 0; i < Nobjects[LANDUSE]; i++ {
			load = float64(load + Subcatch[j].landFactor[i].buildup[p])
		}
		load = float64(load + (float64(Subcatch[j].pondedQual[p] * Pollut[p].mcf)))
	}
	return load
}

// massbal.c:339
func massbal_updateRunoffTotals(flowType int32, v float64) {
	switch flowType {
	case RUNOFF_RAINFALL:
		RunoffTotals.rainfall = float64(RunoffTotals.rainfall + v)
		break
	case RUNOFF_EVAP:
		RunoffTotals.evap = float64(RunoffTotals.evap + v)
		break
	case RUNOFF_INFIL:
		RunoffTotals.infil = float64(RunoffTotals.infil + v)
		break
	case RUNOFF_RUNOFF:
		RunoffTotals.runoff = float64(RunoffTotals.runoff + v)
		break
	case RUNOFF_DRAINS:
		RunoffTotals.drains = float64(RunoffTotals.drains + v)
		break
	case RUNOFF_RUNON:
		RunoffTotals.runon = float64(RunoffTotals.runon + v)
		break
	}
}

// massbal.c:360
func massbal_updateGwaterTotals(vInfil float64, vUpperEvap float64, vLowerEvap float64, vLowerPerc float64, vGwater float64) {
	GwaterTotals.infil = float64(GwaterTotals.infil + vInfil)
	GwaterTotals.upperEvap = float64(GwaterTotals.upperEvap + vUpperEvap)
	GwaterTotals.lowerEvap = float64(GwaterTotals.lowerEvap + vLowerEvap)
	GwaterTotals.lowerPerc = float64(GwaterTotals.lowerPerc + vLowerPerc)
	GwaterTotals.gwater = float64(GwaterTotals.gwater + vGwater)
}

// massbal.c:381
func massbal_initTimeStepTotals() {
	var (
		j int32
	)
	_ = j
	j = 0
	OldStepFlowTotals = StepFlowTotals
	StepFlowTotals.dwInflow = float64(0.0)
	StepFlowTotals.wwInflow = float64(0.0)
	StepFlowTotals.gwInflow = float64(0.0)
	StepFlowTotals.iiInflow = float64(0.0)
	StepFlowTotals.exInflow = float64(0.0)
	StepFlowTotals.flooding = float64(0.0)
	StepFlowTotals.outflow = float64(0.0)
	StepFlowTotals.evapLoss = float64(0.0)
	StepFlowTotals.seepLoss = float64(0.0)
	StepFlowTotals.reacted = float64(0.0)
	for j = 0; j < Nobjects[POLLUT]; j++ {
		StepQualTotals[j].dwInflow = float64(0.0)
		StepQualTotals[j].wwInflow = float64(0.0)
		StepQualTotals[j].gwInflow = float64(0.0)
		StepQualTotals[j].iiInflow = float64(0.0)
		StepQualTotals[j].exInflow = float64(0.0)
		StepQualTotals[j].flooding = float64(0.0)
		StepQualTotals[j].outflow = float64(0.0)
		StepQualTotals[j].reacted = float64(0.0)
		StepQualTotals[j].seepLoss = float64(0.0)
		StepQualTotals[j].initStorage = float64(0.0)
		StepQualTotals[j].finalStorage = float64(0.0)
	}
}

// massbal.c:418
func massbal_addInflowFlow(type_2 int32, q float64) {
	switch type_2 {
	case DRY_WEATHER_INFLOW:
		StepFlowTotals.dwInflow = float64(StepFlowTotals.dwInflow + q)
		break
	case WET_WEATHER_INFLOW:
		StepFlowTotals.wwInflow = float64(StepFlowTotals.wwInflow + q)
		break
	case GROUNDWATER_INFLOW:
		StepFlowTotals.gwInflow = float64(StepFlowTotals.gwInflow + q)
		break
	case RDII_INFLOW:
		StepFlowTotals.iiInflow = float64(StepFlowTotals.iiInflow + q)
		break
	case EXTERNAL_INFLOW:
		StepFlowTotals.exInflow = float64(StepFlowTotals.exInflow + q)
		break
	}
}

// massbal.c:438
func massbal_updateLoadingTotals(type_2 int32, p int32, w float64) {
	switch type_2 {
	case BUILDUP_LOAD:
		LoadingTotals[p].buildup = float64(LoadingTotals[p].buildup + w)
		break
	case DEPOSITION_LOAD:
		LoadingTotals[p].deposition = float64(LoadingTotals[p].deposition + w)
		break
	case SWEEPING_LOAD:
		LoadingTotals[p].sweeping = float64(LoadingTotals[p].sweeping + w)
		break
	case INFIL_LOAD:
		LoadingTotals[p].infil = float64(LoadingTotals[p].infil + w)
		break
	case BMP_REMOVAL_LOAD:
		LoadingTotals[p].bmpRemoval = float64(LoadingTotals[p].bmpRemoval + w)
		break
	case RUNOFF_LOAD:
		LoadingTotals[p].runoff = float64(LoadingTotals[p].runoff + w)
		break
	case FINAL_LOAD:
		LoadingTotals[p].finalLoad = float64(LoadingTotals[p].finalLoad + w)
		break
	}
}

// massbal.c:461
func massbal_addInflowQual(type_2 int32, p int32, w float64) {
	if (p < 0) || (p >= Nobjects[POLLUT]) {
		return
	}
	switch type_2 {
	case DRY_WEATHER_INFLOW:
		StepQualTotals[p].dwInflow = float64(StepQualTotals[p].dwInflow + w)
		break
	case WET_WEATHER_INFLOW:
		StepQualTotals[p].wwInflow = float64(StepQualTotals[p].wwInflow + w)
		break
	case GROUNDWATER_INFLOW:
		StepQualTotals[p].gwInflow = float64(StepQualTotals[p].gwInflow + w)
		break
	case EXTERNAL_INFLOW:
		StepQualTotals[p].exInflow = float64(StepQualTotals[p].exInflow + w)
		break
	case RDII_INFLOW:
		StepQualTotals[p].iiInflow = float64(StepQualTotals[p].iiInflow + w)
		break
	}
}

// massbal.c:483
func massbal_addOutflowFlow(q float64, isFlooded int32) {
	if isFlooded != 0 {
		StepFlowTotals.flooding = float64(StepFlowTotals.flooding + q)
	} else {
		StepFlowTotals.outflow = float64(StepFlowTotals.outflow + q)
	}
}

// massbal.c:497
func massbal_addOutflowQual(p int32, w float64, isFlooded int32) {
	if (p < 0) || (p >= Nobjects[POLLUT]) {
		return
	}
	if w >= float64(0.0) {
		if isFlooded != 0 {
			StepQualTotals[p].flooding = float64(StepQualTotals[p].flooding + w)
		} else {
			StepQualTotals[p].outflow = float64(StepQualTotals[p].outflow + w)
		}
	} else {
		StepQualTotals[p].exInflow = float64(StepQualTotals[p].exInflow - w)
	}
}

// massbal.c:517
func massbal_addReactedMass(p int32, w float64) {
	if (p < 0) || (p >= Nobjects[POLLUT]) {
		return
	}
	StepQualTotals[p].reacted = float64(StepQualTotals[p].reacted + w)
}

// massbal.c:531
func massbal_addSeepageLoss(p int32, w float64) {
	if (p < 0) || (p >= Nobjects[POLLUT]) {
		return
	}
	StepQualTotals[p].seepLoss = float64(StepQualTotals[p].seepLoss + w)
}

// massbal.c:545
func massbal_addToFinalStorage(p int32, w float64) {
	if (p < 0) || (p >= Nobjects[POLLUT]) {
		return
	}
	StepQualTotals[p].finalStorage = float64(StepQualTotals[p].finalStorage + w)
}

// massbal.c:559
func massbal_addNodeLosses(evapLoss float64, seepLoss float64) {
	StepFlowTotals.evapLoss = float64(StepFlowTotals.evapLoss + evapLoss)
	StepFlowTotals.seepLoss = float64(StepFlowTotals.seepLoss + seepLoss)
}

// massbal.c:573
func massbal_addLinkLosses(evapLoss float64, seepLoss float64) {
	StepFlowTotals.evapLoss = float64(StepFlowTotals.evapLoss + evapLoss)
	StepFlowTotals.seepLoss = float64(StepFlowTotals.seepLoss + seepLoss)
}

// massbal.c:587
func massbal_updateRoutingTotals(tStep float64) {
	var (
		j int32
	)
	_ = j
	j = 0
	FlowTotals.dwInflow = float64(FlowTotals.dwInflow + (float64(StepFlowTotals.dwInflow * tStep)))
	FlowTotals.wwInflow = float64(FlowTotals.wwInflow + (float64(StepFlowTotals.wwInflow * tStep)))
	FlowTotals.gwInflow = float64(FlowTotals.gwInflow + (float64(StepFlowTotals.gwInflow * tStep)))
	FlowTotals.iiInflow = float64(FlowTotals.iiInflow + (float64(StepFlowTotals.iiInflow * tStep)))
	FlowTotals.exInflow = float64(FlowTotals.exInflow + (float64(StepFlowTotals.exInflow * tStep)))
	FlowTotals.flooding = float64(FlowTotals.flooding + (float64(StepFlowTotals.flooding * tStep)))
	FlowTotals.outflow = float64(FlowTotals.outflow + (float64(StepFlowTotals.outflow * tStep)))
	FlowTotals.evapLoss = float64(FlowTotals.evapLoss + (float64(StepFlowTotals.evapLoss * tStep)))
	FlowTotals.seepLoss = float64(FlowTotals.seepLoss + (float64(StepFlowTotals.seepLoss * tStep)))
	for j = 0; j < Nobjects[POLLUT]; j++ {
		QualTotals[j].dwInflow = float64(QualTotals[j].dwInflow + (float64(StepQualTotals[j].dwInflow * tStep)))
		QualTotals[j].wwInflow = float64(QualTotals[j].wwInflow + (float64(StepQualTotals[j].wwInflow * tStep)))
		QualTotals[j].gwInflow = float64(QualTotals[j].gwInflow + (float64(StepQualTotals[j].gwInflow * tStep)))
		QualTotals[j].iiInflow = float64(QualTotals[j].iiInflow + (float64(StepQualTotals[j].iiInflow * tStep)))
		QualTotals[j].exInflow = float64(QualTotals[j].exInflow + (float64(StepQualTotals[j].exInflow * tStep)))
		QualTotals[j].flooding = float64(QualTotals[j].flooding + (float64(StepQualTotals[j].flooding * tStep)))
		QualTotals[j].outflow = float64(QualTotals[j].outflow + (float64(StepQualTotals[j].outflow * tStep)))
		QualTotals[j].reacted = float64(QualTotals[j].reacted + (float64(StepQualTotals[j].reacted * tStep)))
		QualTotals[j].seepLoss = float64(QualTotals[j].seepLoss + (float64(StepQualTotals[j].seepLoss * tStep)))
		QualTotals[j].finalStorage = float64(QualTotals[j].finalStorage + StepQualTotals[j].finalStorage)
	}
	for j = 0; j < Nobjects[NODE]; j++ {
		NodeInflow[j] = float64(NodeInflow[j] + (float64(Node[j].inflow * tStep)))
		if (Node[j].type_ == OUTFALL) || ((Node[j].degree == 0) && (Node[j].type_ != STORAGE)) {
			NodeOutflow[j] = float64(NodeOutflow[j] + (float64(Node[j].inflow * tStep)))
		} else {
			NodeOutflow[j] = float64(NodeOutflow[j] + (float64(Node[j].outflow * tStep)))
			if Node[j].newVolume <= Node[j].fullVolume {
				NodeOutflow[j] = float64(NodeOutflow[j] + (float64(Node[j].overflow * tStep)))
			}
		}
	}
}

// massbal.c:638
func massbal_getStorage(isFinalStorage int32) float64 {
	var (
		j            int32
		totalStorage float64
		nodeStorage  float64
	)
	_, _, _ = j, totalStorage, nodeStorage
	j = 0
	totalStorage = float64(0.0)
	nodeStorage = 0
	for j = 0; j < Nobjects[NODE]; j++ {
		nodeStorage = Node[j].newVolume
		if isFinalStorage != 0 {
			NodeOutflow[j] = float64(NodeOutflow[j] + nodeStorage)
		}
		totalStorage = float64(totalStorage + nodeStorage)
	}
	if (isFinalStorage != 0) && (RouteModel == SF) {
		return totalStorage
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		totalStorage = float64(totalStorage + Link[j].newVolume)
	}
	return totalStorage
}

// massbal.c:670
func massbal_getSysFlows(f float64, sysFlows []float64) {
	var (
		f1 float64
	)
	_ = f1
	f1 = float64(float64(1.0) - f)
	sysFlows[SYS_DWFLOW] = float64((float64(float64(f1*OldStepFlowTotals.dwInflow) + float64(f*StepFlowTotals.dwInflow))) * UCF(FLOW))
	sysFlows[SYS_GWFLOW] = float64((float64(float64(f1*OldStepFlowTotals.gwInflow) + float64(f*StepFlowTotals.gwInflow))) * UCF(FLOW))
	sysFlows[SYS_IIFLOW] = float64((float64(float64(f1*OldStepFlowTotals.iiInflow) + float64(f*StepFlowTotals.iiInflow))) * UCF(FLOW))
	sysFlows[SYS_EXFLOW] = float64((float64(float64(f1*OldStepFlowTotals.exInflow) + float64(f*StepFlowTotals.exInflow))) * UCF(FLOW))
	sysFlows[SYS_FLOODING] = float64((float64(float64(f1*OldStepFlowTotals.flooding) + float64(f*StepFlowTotals.flooding))) * UCF(FLOW))
	sysFlows[SYS_OUTFLOW] = float64((float64(float64(f1*OldStepFlowTotals.outflow) + float64(f*StepFlowTotals.outflow))) * UCF(FLOW))
	sysFlows[SYS_STORAGE] = float64((float64(float64(f1*OldStepFlowTotals.finalStorage) + float64(f*StepFlowTotals.finalStorage))) * UCF(VOLUME))
}

// massbal.c:696
func massbal_getRunoffError() float64 {
	var (
		j            int32
		totalInflow  float64
		totalOutflow float64
	)
	_, _, _ = j, totalInflow, totalOutflow
	j = 0
	totalInflow = 0
	totalOutflow = 0
	RunoffTotals.finalStorage = float64(0.0)
	RunoffTotals.finalSnowCover = float64(0.0)
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		RunoffTotals.finalStorage = float64(RunoffTotals.finalStorage + (subcatch_getStorage(j)))
		RunoffTotals.finalSnowCover = float64(RunoffTotals.finalSnowCover + (snow_getSnowCover(j)))
	}
	RunoffTotals.snowRemoved = Snow.removed
	totalInflow = float64(float64(float64(RunoffTotals.rainfall+RunoffTotals.runon)+RunoffTotals.initStorage) + RunoffTotals.initSnowCover)
	totalOutflow = float64(float64(float64(float64(float64(float64(RunoffTotals.evap+RunoffTotals.infil)+RunoffTotals.runoff)+RunoffTotals.drains)+RunoffTotals.snowRemoved)+RunoffTotals.finalStorage) + RunoffTotals.finalSnowCover)
	RunoffTotals.pctError = float64(0.0)
	if M.Fabs(float64(totalInflow-totalOutflow)) < float64(1.0) {
		RunoffTotals.pctError = float64(1e-06)
	} else if totalInflow > float64(0.0) {
		RunoffTotals.pctError = float64(float64(100.0) * (float64(float64(1.0) - float64(totalOutflow/totalInflow))))
	} else if totalOutflow > float64(0.0) {
		RunoffTotals.pctError = float64(float64(100.0) * (float64(float64(totalInflow/totalOutflow) - float64(1.0))))
	}
	RunoffError = RunoffTotals.pctError
	return RunoffTotals.pctError
}

// massbal.c:750
func massbal_getLoadingError() float64 {
	var (
		j        int32
		loadIn   float64
		loadOut  float64
		maxError float64
	)
	_, _, _, _ = j, loadIn, loadOut, maxError
	j = 0
	loadIn = 0
	loadOut = 0
	maxError = float64(0.0)
	for j = 0; j < Nobjects[POLLUT]; j++ {
		LoadingTotals[j].finalLoad = float64(LoadingTotals[j].finalLoad + (massbal_getBuildup(j)))
		loadIn = float64(float64(LoadingTotals[j].initLoad+LoadingTotals[j].buildup) + LoadingTotals[j].deposition)
		loadOut = float64(float64(float64(float64(LoadingTotals[j].sweeping+LoadingTotals[j].infil)+LoadingTotals[j].bmpRemoval)+LoadingTotals[j].runoff) + LoadingTotals[j].finalLoad)
		LoadingTotals[j].pctError = float64(0.0)
		if M.Fabs(float64(loadIn-loadOut)) < float64(0.001) {
			LoadingTotals[j].pctError = float64(1e-06)
		} else if loadIn > float64(0.0) {
			LoadingTotals[j].pctError = float64(float64(100.0) * (float64(float64(1.0) - float64(loadOut/loadIn))))
		} else if loadOut > float64(0.0) {
			LoadingTotals[j].pctError = float64(float64(100.0) * (float64(float64(loadIn/loadOut) - float64(1.0))))
		}
		if (maxError) >= (LoadingTotals[j].pctError) {
			maxError = (maxError)
		} else {
			maxError = (LoadingTotals[j].pctError)
		}
		if Pollut[j].units == COUNT {
			if (LoadingTotals[j].initLoad) > float64(0.0) {
				LoadingTotals[j].initLoad = M.Log10((LoadingTotals[j].initLoad))
			} else {
				LoadingTotals[j].initLoad = (LoadingTotals[j].initLoad)
			}
			if (LoadingTotals[j].buildup) > float64(0.0) {
				LoadingTotals[j].buildup = M.Log10((LoadingTotals[j].buildup))
			} else {
				LoadingTotals[j].buildup = (LoadingTotals[j].buildup)
			}
			if (LoadingTotals[j].deposition) > float64(0.0) {
				LoadingTotals[j].deposition = M.Log10((LoadingTotals[j].deposition))
			} else {
				LoadingTotals[j].deposition = (LoadingTotals[j].deposition)
			}
			if (LoadingTotals[j].sweeping) > float64(0.0) {
				LoadingTotals[j].sweeping = M.Log10((LoadingTotals[j].sweeping))
			} else {
				LoadingTotals[j].sweeping = (LoadingTotals[j].sweeping)
			}
			if (LoadingTotals[j].infil) > float64(0.0) {
				LoadingTotals[j].infil = M.Log10((LoadingTotals[j].infil))
			} else {
				LoadingTotals[j].infil = (LoadingTotals[j].infil)
			}
			if (LoadingTotals[j].bmpRemoval) > float64(0.0) {
				LoadingTotals[j].bmpRemoval = M.Log10((LoadingTotals[j].bmpRemoval))
			} else {
				LoadingTotals[j].bmpRemoval = (LoadingTotals[j].bmpRemoval)
			}
			if (LoadingTotals[j].runoff) > float64(0.0) {
				LoadingTotals[j].runoff = M.Log10((LoadingTotals[j].runoff))
			} else {
				LoadingTotals[j].runoff = (LoadingTotals[j].runoff)
			}
			if (LoadingTotals[j].finalLoad) > float64(0.0) {
				LoadingTotals[j].finalLoad = M.Log10((LoadingTotals[j].finalLoad))
			} else {
				LoadingTotals[j].finalLoad = (LoadingTotals[j].finalLoad)
			}
		}
	}
	return maxError
}

// massbal.c:813
func massbal_getGwaterError() float64 {
	var (
		j            int32
		totalInflow  float64
		totalOutflow float64
	)
	_, _, _ = j, totalInflow, totalOutflow
	j = 0
	totalInflow = 0
	totalOutflow = 0
	GwaterTotals.finalStorage = float64(0.0)
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		GwaterTotals.finalStorage = float64(GwaterTotals.finalStorage + (float64(gwater_getVolume(j) * Subcatch[j].area)))
	}
	totalInflow = float64(GwaterTotals.infil + GwaterTotals.initStorage)
	totalOutflow = float64(float64(float64(float64(GwaterTotals.upperEvap+GwaterTotals.lowerEvap)+GwaterTotals.lowerPerc)+GwaterTotals.gwater) + GwaterTotals.finalStorage)
	GwaterTotals.pctError = float64(0.0)
	if M.Fabs(float64(totalInflow-totalOutflow)) < float64(1.0) {
		GwaterTotals.pctError = float64(1e-06)
	} else if totalInflow > float64(0.0) {
		GwaterTotals.pctError = float64(float64(100.0) * (float64(float64(1.0) - float64(totalOutflow/totalInflow))))
	} else if totalOutflow > float64(0.0) {
		GwaterTotals.pctError = float64(float64(100.0) * (float64(float64(totalInflow/totalOutflow) - float64(1.0))))
	}
	GwaterError = GwaterTotals.pctError
	return GwaterTotals.pctError
}

// massbal.c:858
func massbal_getFlowError() float64 {
	var (
		totalInflow  float64
		totalOutflow float64
	)
	_, _ = totalInflow, totalOutflow
	totalInflow = 0
	totalOutflow = 0
	FlowTotals.finalStorage = massbal_getStorage(1)
	totalInflow = float64(float64(FlowTotals.initStorage+FlowTotals.wwInflow) + FlowTotals.iiInflow)
	totalOutflow = float64(float64(float64(float64(FlowTotals.finalStorage+FlowTotals.flooding)+FlowTotals.evapLoss)+FlowTotals.seepLoss) + FlowTotals.reacted)
	if FlowTotals.dwInflow >= float64(0.0) {
		totalInflow = float64(totalInflow + FlowTotals.dwInflow)
	} else {
		totalOutflow = float64(totalOutflow - FlowTotals.dwInflow)
	}
	if FlowTotals.gwInflow >= float64(0.0) {
		totalInflow = float64(totalInflow + FlowTotals.gwInflow)
	} else {
		totalOutflow = float64(totalOutflow - FlowTotals.gwInflow)
	}
	if FlowTotals.exInflow >= float64(0.0) {
		totalInflow = float64(totalInflow + FlowTotals.exInflow)
	} else {
		totalOutflow = float64(totalOutflow - FlowTotals.exInflow)
	}
	if FlowTotals.outflow >= float64(0.0) {
		totalOutflow = float64(totalOutflow + FlowTotals.outflow)
	} else {
		totalInflow = float64(totalInflow - FlowTotals.outflow)
	}
	FlowTotals.pctError = float64(0.0)
	if M.Fabs(float64(totalInflow-totalOutflow)) < float64(1.0) {
		FlowTotals.pctError = float64(1e-06)
	} else if M.Fabs(totalInflow) > float64(0.0) {
		FlowTotals.pctError = float64(float64(100.0) * (float64(float64(1.0) - float64(totalOutflow/totalInflow))))
	} else if M.Fabs(totalOutflow) > float64(0.0) {
		FlowTotals.pctError = float64(float64(100.0) * (float64(float64(totalInflow/totalOutflow) - float64(1.0))))
	}
	FlowError = FlowTotals.pctError
	return FlowTotals.pctError
}

// massbal.c:906
func massbal_getQualError() float64 {
	var (
		p            int32
		maxQualError float64
		totalInflow  float64
		totalOutflow float64
		cf           float64
	)
	_, _, _, _, _ = p, maxQualError, totalInflow, totalOutflow, cf
	p = 0
	maxQualError = float64(0.0)
	totalInflow = 0
	totalOutflow = 0
	cf = 0
	for p = 0; p < Nobjects[POLLUT]; p++ {
		QualTotals[p].finalStorage = float64(QualTotals[p].finalStorage + (massbal_getStoredMass(p)))
		totalInflow = float64(float64(float64(float64(float64(QualTotals[p].dwInflow+QualTotals[p].wwInflow)+QualTotals[p].gwInflow)+QualTotals[p].iiInflow)+QualTotals[p].exInflow) + QualTotals[p].initStorage)
		totalOutflow = float64(float64(float64(float64(QualTotals[p].flooding+QualTotals[p].outflow)+QualTotals[p].reacted)+QualTotals[p].seepLoss) + QualTotals[p].finalStorage)
		QualTotals[p].pctError = float64(0.0)
		if M.Fabs(float64(totalInflow-totalOutflow)) < float64(0.001) {
			QualTotals[p].pctError = float64(1e-06)
		} else if totalInflow > float64(0.0) {
			QualTotals[p].pctError = float64(float64(100.0) * (float64(float64(1.0) - float64(totalOutflow/totalInflow))))
		} else if totalOutflow > float64(0.0) {
			QualTotals[p].pctError = float64(float64(100.0) * (float64(float64(totalInflow/totalOutflow) - float64(1.0))))
		}
		if M.Fabs(QualTotals[p].pctError) > M.Fabs(maxQualError) {
			maxQualError = QualTotals[p].pctError
		}
		cf = float64(28.317)
		if Pollut[p].units == COUNT {
			if (float64(cf * QualTotals[p].dwInflow)) > float64(0.0) {
				QualTotals[p].dwInflow = M.Log10((float64(cf * QualTotals[p].dwInflow)))
			} else {
				QualTotals[p].dwInflow = (float64(cf * QualTotals[p].dwInflow))
			}
			if (float64(cf * QualTotals[p].wwInflow)) > float64(0.0) {
				QualTotals[p].wwInflow = M.Log10((float64(cf * QualTotals[p].wwInflow)))
			} else {
				QualTotals[p].wwInflow = (float64(cf * QualTotals[p].wwInflow))
			}
			if (float64(cf * QualTotals[p].gwInflow)) > float64(0.0) {
				QualTotals[p].gwInflow = M.Log10((float64(cf * QualTotals[p].gwInflow)))
			} else {
				QualTotals[p].gwInflow = (float64(cf * QualTotals[p].gwInflow))
			}
			if (float64(cf * QualTotals[p].iiInflow)) > float64(0.0) {
				QualTotals[p].iiInflow = M.Log10((float64(cf * QualTotals[p].iiInflow)))
			} else {
				QualTotals[p].iiInflow = (float64(cf * QualTotals[p].iiInflow))
			}
			if (float64(cf * QualTotals[p].exInflow)) > float64(0.0) {
				QualTotals[p].exInflow = M.Log10((float64(cf * QualTotals[p].exInflow)))
			} else {
				QualTotals[p].exInflow = (float64(cf * QualTotals[p].exInflow))
			}
			if (float64(cf * QualTotals[p].flooding)) > float64(0.0) {
				QualTotals[p].flooding = M.Log10((float64(cf * QualTotals[p].flooding)))
			} else {
				QualTotals[p].flooding = (float64(cf * QualTotals[p].flooding))
			}
			if (float64(cf * QualTotals[p].outflow)) > float64(0.0) {
				QualTotals[p].outflow = M.Log10((float64(cf * QualTotals[p].outflow)))
			} else {
				QualTotals[p].outflow = (float64(cf * QualTotals[p].outflow))
			}
			if (float64(cf * QualTotals[p].reacted)) > float64(0.0) {
				QualTotals[p].reacted = M.Log10((float64(cf * QualTotals[p].reacted)))
			} else {
				QualTotals[p].reacted = (float64(cf * QualTotals[p].reacted))
			}
			if (float64(cf * QualTotals[p].seepLoss)) > float64(0.0) {
				QualTotals[p].seepLoss = M.Log10((float64(cf * QualTotals[p].seepLoss)))
			} else {
				QualTotals[p].seepLoss = (float64(cf * QualTotals[p].seepLoss))
			}
			if (float64(cf * QualTotals[p].initStorage)) > float64(0.0) {
				QualTotals[p].initStorage = M.Log10((float64(cf * QualTotals[p].initStorage)))
			} else {
				QualTotals[p].initStorage = (float64(cf * QualTotals[p].initStorage))
			}
			if (float64(cf * QualTotals[p].finalStorage)) > float64(0.0) {
				QualTotals[p].finalStorage = M.Log10((float64(cf * QualTotals[p].finalStorage)))
			} else {
				QualTotals[p].finalStorage = (float64(cf * QualTotals[p].finalStorage))
			}
		} else {
			cf = float64(cf * UCF(MASS))
			if Pollut[p].units == UG {
				cf = float64(cf / (float64(1000.0)))
			}
			QualTotals[p].dwInflow = float64(QualTotals[p].dwInflow * cf)
			QualTotals[p].wwInflow = float64(QualTotals[p].wwInflow * cf)
			QualTotals[p].gwInflow = float64(QualTotals[p].gwInflow * cf)
			QualTotals[p].iiInflow = float64(QualTotals[p].iiInflow * cf)
			QualTotals[p].exInflow = float64(QualTotals[p].exInflow * cf)
			QualTotals[p].flooding = float64(QualTotals[p].flooding * cf)
			QualTotals[p].outflow = float64(QualTotals[p].outflow * cf)
			QualTotals[p].reacted = float64(QualTotals[p].reacted * cf)
			QualTotals[p].seepLoss = float64(QualTotals[p].seepLoss * cf)
			QualTotals[p].initStorage = float64(QualTotals[p].initStorage * cf)
			QualTotals[p].finalStorage = float64(QualTotals[p].finalStorage * cf)
		}
	}
	QualError = maxQualError
	return maxQualError
}

// massbal.c:995
func massbal_getStepFlowError() float64 {
	var (
		totalInflow  float64
		totalOutflow float64
	)
	_, _ = totalInflow, totalOutflow
	totalInflow = 0
	totalOutflow = 0
	totalInflow = float64(float64(float64(float64(StepFlowTotals.dwInflow+StepFlowTotals.wwInflow)+StepFlowTotals.gwInflow)+StepFlowTotals.iiInflow) + StepFlowTotals.exInflow)
	totalOutflow = float64(float64(float64(float64(StepFlowTotals.flooding+StepFlowTotals.outflow)+StepFlowTotals.evapLoss)+StepFlowTotals.seepLoss) + StepFlowTotals.reacted)
	if M.Fabs(totalInflow) > float64(0.0) {
		return float64(float64(1.0) - float64(totalOutflow/totalInflow))
	} else if M.Fabs(totalOutflow) > float64(0.0) {
		return float64(float64(totalInflow/totalOutflow) - float64(1.0))
	} else {
		return float64(0.0)
	}
	return 0
}

// massbal.c:1025
func massbal_getStoredMass(p int32) float64 {
	var (
		j          int32
		storedMass float64
	)
	_, _ = j, storedMass
	j = 0
	storedMass = float64(0.0)
	for j = 0; j < Nobjects[NODE]; j++ {
		storedMass = float64(storedMass + (float64(Node[j].newVolume * Node[j].newQual[p])))
	}
	if RouteModel != SF {
		for j = 0; j < Nobjects[LINK]; j++ {
			storedMass = float64(storedMass + (float64(Link[j].newVolume * Link[j].newQual[p])))
		}
	}
	return storedMass
}
