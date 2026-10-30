// Generated from EPA SWMM 5.2.4 (stats.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var TimeStepStats TTimeStepStats
var MaxMassBalErrs [5]TMaxStats
var MaxCourantCrit [5]TMaxStats
var MaxFlowTurns [5]TMaxStats
var MaxNonConverged [5]TMaxStats
var SysOutfallFlow float64
var SubcatchStats []TSubcatchStats
var NodeStats []TNodeStats
var LinkStats []TLinkStats
var StorageStats []TStorageStats
var OutfallStats []TOutfallStats
var PumpStats []TPumpStats
var MaxOutfallFlow float64
var MaxRunoffFlow float64
var RoutingTimeSpan float64

// stats.c:102
func stats_open() int32 {
	var (
		j              int32
		k              int32
		timeStepDelta  float64
		logMaxTimeStep float64
		logMinTimeStep float64
	)
	_, _, _, _, _ = j, k, timeStepDelta, logMaxTimeStep, logMinTimeStep
	j = 0
	k = 0
	timeStepDelta = 0
	logMaxTimeStep = 0
	logMinTimeStep = 0
	NodeStats = nil
	LinkStats = nil
	StorageStats = nil
	OutfallStats = nil
	PumpStats = nil
	SubcatchStats = nil
	if Nobjects[SUBCATCH] > 0 {
		SubcatchStats = make([]TSubcatchStats, int(uint64(Nobjects[SUBCATCH])))
		if !(SubcatchStats != nil) {
			report_writeErrorMsg(ERR_MEMORY, L0)
			return ErrorCode
		}
		for j = 0; j < Nobjects[SUBCATCH]; j++ {
			SubcatchStats[j].precip = float64(0.0)
			SubcatchStats[j].runon = float64(0.0)
			SubcatchStats[j].evap = float64(0.0)
			SubcatchStats[j].infil = float64(0.0)
			SubcatchStats[j].runoff = float64(0.0)
			SubcatchStats[j].maxFlow = float64(0.0)
			SubcatchStats[j].impervRunoff = float64(0.0)
			SubcatchStats[j].pervRunoff = float64(0.0)
		}
		for j = 0; j < Nobjects[SUBCATCH]; j++ {
			if Subcatch[j].groundwater == nil {
				continue
			}
			Subcatch[j].groundwater.stats.avgUpperMoist = float64(0.0)
			Subcatch[j].groundwater.stats.avgWaterTable = float64(0.0)
			Subcatch[j].groundwater.stats.infil = float64(0.0)
			Subcatch[j].groundwater.stats.latFlow = float64(0.0)
			Subcatch[j].groundwater.stats.deepFlow = float64(0.0)
			Subcatch[j].groundwater.stats.evap = float64(0.0)
			Subcatch[j].groundwater.stats.maxFlow = float64(0.0)
			Subcatch[j].groundwater.stats.finalUpperMoist = float64(0.0)
			Subcatch[j].groundwater.stats.finalWaterTable = float64(0.0)
		}
	}
	if Nobjects[LINK] > 0 {
		NodeStats = make([]TNodeStats, int(uint64(Nobjects[NODE])))
		LinkStats = make([]TLinkStats, int(uint64(Nobjects[LINK])))
		if (!(NodeStats != nil)) || (!(LinkStats != nil)) {
			report_writeErrorMsg(ERR_MEMORY, L0)
			return ErrorCode
		}
	}
	if NodeStats != nil {
		for j = 0; j < Nobjects[NODE]; j++ {
			NodeStats[j].avgDepth = float64(0.0)
			NodeStats[j].maxDepth = float64(0.0)
			NodeStats[j].maxDepthDate = StartDateTime
			NodeStats[j].maxRptDepth = float64(0.0)
			NodeStats[j].volFlooded = float64(0.0)
			NodeStats[j].timeFlooded = float64(0.0)
			NodeStats[j].timeSurcharged = float64(0.0)
			NodeStats[j].timeCourantCritical = float64(0.0)
			NodeStats[j].totLatFlow = float64(0.0)
			NodeStats[j].maxLatFlow = float64(0.0)
			NodeStats[j].maxInflow = float64(0.0)
			NodeStats[j].maxOverflow = float64(0.0)
			NodeStats[j].maxPondedVol = float64(0.0)
			NodeStats[j].nonConvergedCount = 0
			NodeStats[j].maxInflowDate = StartDateTime
			NodeStats[j].maxOverflowDate = StartDateTime
		}
	}
	if LinkStats != nil {
		for j = 0; j < Nobjects[LINK]; j++ {
			LinkStats[j].maxFlow = float64(0.0)
			LinkStats[j].maxVeloc = float64(0.0)
			LinkStats[j].maxDepth = float64(0.0)
			LinkStats[j].maxStreetFilled = float64(0.0)
			LinkStats[j].timeSurcharged = float64(0.0)
			LinkStats[j].timeFullUpstream = float64(0.0)
			LinkStats[j].timeFullDnstream = float64(0.0)
			LinkStats[j].timeFullFlow = float64(0.0)
			LinkStats[j].timeCapacityLimited = float64(0.0)
			LinkStats[j].timeCourantCritical = float64(0.0)
			for k = 0; k < MAX_FLOW_CLASSES; k++ {
				LinkStats[j].timeInFlowClass[k] = float64(0.0)
			}
			LinkStats[j].flowTurns = 0
			LinkStats[j].flowTurnSign = 0
		}
	}
	if Nnodes[STORAGE] > 0 {
		StorageStats = make([]TStorageStats, int(uint64(Nnodes[STORAGE])))
		if !(StorageStats != nil) {
			report_writeErrorMsg(ERR_MEMORY, L0)
			return ErrorCode
		} else {
			for k = 0; k < Nobjects[NODE]; k++ {
				if Node[k].type_ != STORAGE {
					continue
				}
				j = Node[k].subIndex
				StorageStats[j].initVol = Node[k].newVolume
				StorageStats[j].avgVol = float64(0.0)
				StorageStats[j].maxVol = float64(0.0)
				StorageStats[j].maxFlow = float64(0.0)
				StorageStats[j].evapLosses = float64(0.0)
				StorageStats[j].exfilLosses = float64(0.0)
				StorageStats[j].maxVolDate = StartDateTime
			}
		}
	}
	if Nnodes[OUTFALL] > 0 {
		OutfallStats = make([]TOutfallStats, int(uint64(Nnodes[OUTFALL])))
		if !(OutfallStats != nil) {
			report_writeErrorMsg(ERR_MEMORY, L0)
			return ErrorCode
		} else {
			for j = 0; j < Nnodes[OUTFALL]; j++ {
				OutfallStats[j].avgFlow = float64(0.0)
				OutfallStats[j].maxFlow = float64(0.0)
				OutfallStats[j].totalPeriods = 0
				if Nobjects[POLLUT] > 0 {
					OutfallStats[j].totalLoad = make([]float64, int(uint64(Nobjects[POLLUT])))
					if !(OutfallStats[j].totalLoad != nil) {
						report_writeErrorMsg(ERR_MEMORY, L0)
						return ErrorCode
					}
					for k = 0; k < Nobjects[POLLUT]; k++ {
						OutfallStats[j].totalLoad[k] = float64(0.0)
					}
				} else {
					OutfallStats[j].totalLoad = nil
				}
			}
		}
	}
	if Nlinks[PUMP] > 0 {
		PumpStats = make([]TPumpStats, int(uint64(Nlinks[PUMP])))
		if !(PumpStats != nil) {
			report_writeErrorMsg(ERR_MEMORY, L0)
			return ErrorCode
		} else {
			for j = 0; j < Nlinks[PUMP]; j++ {
				PumpStats[j].utilized = float64(0.0)
				PumpStats[j].minFlow = float64(0.0)
				PumpStats[j].avgFlow = float64(0.0)
				PumpStats[j].maxFlow = float64(0.0)
				PumpStats[j].volume = float64(0.0)
				PumpStats[j].energy = float64(0.0)
				PumpStats[j].startUps = 0
				PumpStats[j].offCurveLow = float64(0.0)
				PumpStats[j].offCurveHigh = float64(0.0)
			}
		}
	}
	MaxRunoffFlow = float64(0.0)
	MaxOutfallFlow = float64(0.0)
	TimeStepStats.maxTimeStep = float64(0.0)
	TimeStepStats.minTimeStep = RouteStep
	TimeStepStats.routingTime = float64(0.0)
	TimeStepStats.trialsCount = float64(0.0)
	TimeStepStats.steadyStateTime = float64(0.0)
	TimeStepStats.timeStepCount = 0
	logMaxTimeStep = M.Log10(RouteStep)
	logMinTimeStep = M.Log10(MinRouteStep)
	timeStepDelta = float64((float64(logMaxTimeStep - logMinTimeStep)) / float64((6 - 1)))
	TimeStepStats.timeStepIntervals[0] = RouteStep
	for j = 1; j < 6; j++ {
		TimeStepStats.timeStepIntervals[j] = M.Pow(float64(10.0), float64(logMaxTimeStep-float64(float64(j)*timeStepDelta)))
		TimeStepStats.timeStepCounts[j] = 0
	}
	TimeStepStats.timeStepIntervals[(6 - 1)] = MinRouteStep
	RoutingTimeSpan = float64(0.0)
	return 0
}

// stats.c:318
func stats_close() {
	var (
		j int32
	)
	_ = j
	j = 0
	{
		if SubcatchStats != nil {
			_ = SubcatchStats
			SubcatchStats = nil
		}
	}
	{
		if NodeStats != nil {
			_ = NodeStats
			NodeStats = nil
		}
	}
	{
		if LinkStats != nil {
			_ = LinkStats
			LinkStats = nil
		}
	}
	{
		if StorageStats != nil {
			_ = StorageStats
			StorageStats = nil
		}
	}
	if OutfallStats != nil {
		for j = 0; j < Nnodes[OUTFALL]; j++ {
			if OutfallStats[j].totalLoad != nil {
				_ = OutfallStats[j].totalLoad
				OutfallStats[j].totalLoad = nil
			}
		}
		{
			if OutfallStats != nil {
				_ = OutfallStats
				OutfallStats = nil
			}
		}
	}
	{
		if PumpStats != nil {
			_ = PumpStats
			PumpStats = nil
		}
	}
}

// stats.c:342
func stats_report() {
	if (Nobjects[LINK] > 0) && (RouteModel != NO_ROUTING) {
		stats_findMaxStats()
		if (!(RptFlags.disabled != 0)) && (RptFlags.flowStats != 0) {
			report_writeMaxStats(MaxMassBalErrs[:], MaxCourantCrit[:], 5)
			report_writeMaxFlowTurns(MaxFlowTurns[:], 5)
			report_writeNonconvergedStats(MaxNonConverged[:], 5)
			report_writeTimeStepStats(&TimeStepStats)
		}
	}
	if !(RptFlags.disabled != 0) {
		statsrpt_writeReport()
	}
}

// stats.c:369
func stats_updateSubcatchStats(j int32, rainVol float64, runonVol float64, evapVol float64, infilVol float64, impervVol float64, pervVol float64, runoffVol float64, runoff float64) {
	SubcatchStats[j].precip = float64(SubcatchStats[j].precip + rainVol)
	SubcatchStats[j].runon = float64(SubcatchStats[j].runon + runonVol)
	SubcatchStats[j].evap = float64(SubcatchStats[j].evap + evapVol)
	SubcatchStats[j].infil = float64(SubcatchStats[j].infil + infilVol)
	SubcatchStats[j].runoff = float64(SubcatchStats[j].runoff + runoffVol)
	if (SubcatchStats[j].maxFlow) >= (runoff) {
		SubcatchStats[j].maxFlow = (SubcatchStats[j].maxFlow)
	} else {
		SubcatchStats[j].maxFlow = (runoff)
	}
	SubcatchStats[j].impervRunoff = float64(SubcatchStats[j].impervRunoff + impervVol)
	SubcatchStats[j].pervRunoff = float64(SubcatchStats[j].pervRunoff + pervVol)
}

// stats.c:399
func stats_updateGwaterStats(j int32, infil float64, evap float64, latFlow float64, deepFlow float64, theta float64, waterTable float64, tStep float64) {
	Subcatch[j].groundwater.stats.infil = float64(Subcatch[j].groundwater.stats.infil + (float64(infil * tStep)))
	Subcatch[j].groundwater.stats.evap = float64(Subcatch[j].groundwater.stats.evap + (float64(evap * tStep)))
	Subcatch[j].groundwater.stats.latFlow = float64(Subcatch[j].groundwater.stats.latFlow + (float64(latFlow * tStep)))
	Subcatch[j].groundwater.stats.deepFlow = float64(Subcatch[j].groundwater.stats.deepFlow + (float64(deepFlow * tStep)))
	Subcatch[j].groundwater.stats.avgUpperMoist = float64(Subcatch[j].groundwater.stats.avgUpperMoist + (float64(theta * tStep)))
	Subcatch[j].groundwater.stats.avgWaterTable = float64(Subcatch[j].groundwater.stats.avgWaterTable + (float64(waterTable * tStep)))
	Subcatch[j].groundwater.stats.finalUpperMoist = theta
	Subcatch[j].groundwater.stats.finalWaterTable = waterTable
	if M.Fabs(latFlow) > M.Fabs(Subcatch[j].groundwater.stats.maxFlow) {
		Subcatch[j].groundwater.stats.maxFlow = latFlow
	}
}

// stats.c:419
func stats_updateMaxRunoff() {
	var (
		j         int32
		sysRunoff float64
	)
	_, _ = j, sysRunoff
	j = 0
	sysRunoff = float64(0.0)
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		sysRunoff = float64(sysRunoff + Subcatch[j].newRunoff)
	}
	if (MaxRunoffFlow) >= (sysRunoff) {
		MaxRunoffFlow = (MaxRunoffFlow)
	} else {
		MaxRunoffFlow = (sysRunoff)
	}
}

// stats.c:435
func stats_updateMaxNodeDepth(j int32, depth float64) {
	if NodeStats != nil {
		if (NodeStats[j].maxRptDepth) >= (depth) {
			NodeStats[j].maxRptDepth = (NodeStats[j].maxRptDepth)
		} else {
			NodeStats[j].maxRptDepth = (depth)
		}
	}
}

// stats.c:449
func stats_updateFlowStats(tStep float64, aDate float64) {
	var (
		j int32
	)
	_ = j
	j = 0
	if aDate < ReportStart {
		return
	}
	SysOutfallFlow = float64(0.0)
	{
		for j = 0; j < Nobjects[NODE]; j++ {
			stats_updateNodeStats(j, tStep, aDate)
		}
		for j = 0; j < Nobjects[LINK]; j++ {
			stats_updateLinkStats(j, tStep, aDate)
		}
	}
	ReportStepCount++
	RoutingTimeSpan = float64(RoutingTimeSpan + tStep)
	if (MaxOutfallFlow) >= (SysOutfallFlow) {
		MaxOutfallFlow = (MaxOutfallFlow)
	} else {
		MaxOutfallFlow = (SysOutfallFlow)
	}
}

// stats.c:484
func stats_updateTimeStepStats(tStep float64, trialsCount int32, steadyState int32) {
	var (
		j int32
	)
	_ = j
	j = 0
	TimeStepStats.steadyStateTime = float64(TimeStepStats.steadyStateTime + (float64(float64(steadyState) * tStep)))
	if steadyState == 0 {
		if OldRoutingTime > float64(0) {
			if (TimeStepStats.minTimeStep) <= (tStep) {
				TimeStepStats.minTimeStep = (TimeStepStats.minTimeStep)
			} else {
				TimeStepStats.minTimeStep = (tStep)
			}
			for j = 1; j < 6; j++ {
				if tStep >= TimeStepStats.timeStepIntervals[j] {
					TimeStepStats.timeStepCounts[j]++
					break
				}
			}
		}
		if (TimeStepStats.maxTimeStep) >= (tStep) {
			TimeStepStats.maxTimeStep = (TimeStepStats.maxTimeStep)
		} else {
			TimeStepStats.maxTimeStep = (tStep)
		}
		TimeStepStats.routingTime = float64(TimeStepStats.routingTime + tStep)
		TimeStepStats.timeStepCount++
		TimeStepStats.trialsCount = float64(TimeStepStats.trialsCount + (float64(trialsCount)))
	}
}

// stats.c:522
func stats_updateCriticalTimeCount(node int32, link int32) {
	if node >= 0 {
		NodeStats[node].timeCourantCritical = float64(NodeStats[node].timeCourantCritical + (float64(1.0)))
	} else if link >= 0 {
		LinkStats[link].timeCourantCritical = float64(LinkStats[link].timeCourantCritical + (float64(1.0)))
	}
}

// stats.c:536
func stats_updateConvergenceStats(node int32, converged int32) {
	if converged == 0 {
		NodeStats[node].nonConvergedCount++
	}
}

// stats.c:543
func stats_updateNodeStats(j int32, tStep float64, aDate float64) {
	var (
		k         int32
		p         int32
		newVolume float64
		newDepth  float64
		canPond   int32
	)
	_, _, _, _, _ = k, p, newVolume, newDepth, canPond
	k = 0
	p = 0
	newVolume = Node[j].newVolume
	newDepth = Node[j].newDepth
	canPond = R.B2I(((AllowPonding != 0) && (Node[j].pondedArea > float64(0.0))))
	NodeStats[j].avgDepth = float64(NodeStats[j].avgDepth + newDepth)
	if newDepth > NodeStats[j].maxDepth {
		NodeStats[j].maxDepth = newDepth
		NodeStats[j].maxDepthDate = aDate
	}
	if Node[j].type_ != OUTFALL {
		if (newVolume > Node[j].fullVolume) || (Node[j].overflow > float64(0.0)) {
			NodeStats[j].timeFlooded = float64(NodeStats[j].timeFlooded + tStep)
			NodeStats[j].volFlooded = float64(NodeStats[j].volFlooded + (float64(Node[j].overflow * tStep)))
			if canPond != 0 {
				if (NodeStats[j].maxPondedVol) >= (float64(newVolume - Node[j].fullVolume)) {
					NodeStats[j].maxPondedVol = (NodeStats[j].maxPondedVol)
				} else {
					NodeStats[j].maxPondedVol = (float64(newVolume - Node[j].fullVolume))
				}
			}
		}
		if RouteModel == DW {
			if ((Node[j].type_ != STORAGE) || (Node[j].surDepth > float64(0.0))) && (float64(float64(newDepth+Node[j].invertElev)+float64(0.0001)) >= Node[j].crownElev) {
				NodeStats[j].timeSurcharged = float64(NodeStats[j].timeSurcharged + tStep)
			}
		}
	}
	if Node[j].type_ == STORAGE {
		k = Node[j].subIndex
		StorageStats[k].avgVol = float64(StorageStats[k].avgVol + newVolume)
		StorageStats[k].evapLosses = float64(StorageStats[k].evapLosses + Storage[Node[j].subIndex].evapLoss)
		StorageStats[k].exfilLosses = float64(StorageStats[k].exfilLosses + Storage[Node[j].subIndex].exfilLoss)
		if (newVolume) <= (Node[j].fullVolume) {
			newVolume = (newVolume)
		} else {
			newVolume = (Node[j].fullVolume)
		}
		if newVolume > StorageStats[k].maxVol {
			StorageStats[k].maxVol = newVolume
			StorageStats[k].maxVolDate = aDate
		}
		if (StorageStats[k].maxFlow) >= (Node[j].outflow) {
			StorageStats[k].maxFlow = (StorageStats[k].maxFlow)
		} else {
			StorageStats[k].maxFlow = (Node[j].outflow)
		}
	}
	if Node[j].type_ == OUTFALL {
		k = Node[j].subIndex
		if Node[j].inflow >= float64(0.001) {
			OutfallStats[k].avgFlow = float64(OutfallStats[k].avgFlow + Node[j].inflow)
			if (OutfallStats[k].maxFlow) >= (Node[j].inflow) {
				OutfallStats[k].maxFlow = (OutfallStats[k].maxFlow)
			} else {
				OutfallStats[k].maxFlow = (Node[j].inflow)
			}
			OutfallStats[k].totalPeriods++
		}
		for p = 0; p < Nobjects[POLLUT]; p++ {
			OutfallStats[k].totalLoad[p] = float64(OutfallStats[k].totalLoad[p] + (float64(float64(Node[j].inflow*Node[j].newQual[p]) * tStep)))
		}
		SysOutfallFlow = float64(SysOutfallFlow + Node[j].inflow)
	}
	NodeStats[j].totLatFlow = float64(NodeStats[j].totLatFlow + (float64(float64((float64(Node[j].oldLatFlow+Node[j].newLatFlow))*float64(0.5)) * tStep)))
	if M.Fabs(Node[j].newLatFlow) > M.Fabs(NodeStats[j].maxLatFlow) {
		NodeStats[j].maxLatFlow = Node[j].newLatFlow
	}
	if Node[j].inflow > NodeStats[j].maxInflow {
		NodeStats[j].maxInflow = Node[j].inflow
		NodeStats[j].maxInflowDate = aDate
	}
	if Node[j].overflow > NodeStats[j].maxOverflow {
		NodeStats[j].maxOverflow = Node[j].overflow
		NodeStats[j].maxOverflowDate = aDate
	}
}

// stats.c:647
func stats_updateLinkStats(j int32, tStep float64, aDate float64) {
	var (
		k  int32
		q  float64
		v  float64
		dq float64
	)
	_, _, _, _ = k, q, v, dq
	k = 0
	q = 0
	v = 0
	dq = 0
	dq = float64(Link[j].newFlow - Link[j].oldFlow)
	q = M.Fabs(Link[j].newFlow)
	if q > LinkStats[j].maxFlow {
		LinkStats[j].maxFlow = q
		LinkStats[j].maxFlowDate = aDate
	}
	v = link_getVelocity(j, q, Link[j].newDepth)
	if v > LinkStats[j].maxVeloc {
		LinkStats[j].maxVeloc = v
	}
	if Link[j].newDepth > LinkStats[j].maxDepth {
		LinkStats[j].maxDepth = Link[j].newDepth
	}
	if Link[j].type_ == PUMP {
		if q >= Link[j].qFull {
			LinkStats[j].timeFullFlow = float64(LinkStats[j].timeFullFlow + tStep)
		}
		if q > float64(0.001) {
			k = Link[j].subIndex
			if (PumpStats[k].minFlow) <= (q) {
				PumpStats[k].minFlow = (PumpStats[k].minFlow)
			} else {
				PumpStats[k].minFlow = (q)
			}
			PumpStats[k].maxFlow = LinkStats[j].maxFlow
			PumpStats[k].avgFlow = float64(PumpStats[k].avgFlow + q)
			PumpStats[k].volume = float64(PumpStats[k].volume + (float64(q * tStep)))
			PumpStats[k].utilized = float64(PumpStats[k].utilized + tStep)
			PumpStats[k].energy = float64(PumpStats[k].energy + (float64(float64(link_getPower(j)*tStep) / float64(3600.0))))
			if Link[j].flowClass == DN_DRY {
				PumpStats[k].offCurveLow = float64(PumpStats[k].offCurveLow + tStep)
			}
			if Link[j].flowClass == UP_DRY {
				PumpStats[k].offCurveHigh = float64(PumpStats[k].offCurveHigh + tStep)
			}
			if Link[j].oldFlow < float64(0.001) {
				PumpStats[k].startUps++
			}
			PumpStats[k].totalPeriods++
			LinkStats[j].timeSurcharged = float64(LinkStats[j].timeSurcharged + tStep)
			LinkStats[j].timeFullUpstream = float64(LinkStats[j].timeFullUpstream + tStep)
			LinkStats[j].timeFullDnstream = float64(LinkStats[j].timeFullDnstream + tStep)
		}
	} else if Link[j].type_ == CONDUIT {
		if Link[j].normalFlow != 0 {
			LinkStats[j].timeNormalFlow = float64(LinkStats[j].timeNormalFlow + tStep)
		}
		if Link[j].inletControl != 0 {
			LinkStats[j].timeInletControl = float64(LinkStats[j].timeInletControl + tStep)
		}
		k = Link[j].flowClass
		if (k >= 0) && (k < MAX_FLOW_CLASSES) {
			LinkStats[j].timeInFlowClass[k] = float64(LinkStats[j].timeInFlowClass[k] + tStep)
		}
		k = Link[j].subIndex
		if q >= float64(Link[j].qFull*float64(Conduit[k].barrels)) {
			LinkStats[j].timeFullFlow = float64(LinkStats[j].timeFullFlow + tStep)
		}
		if Conduit[k].capacityLimited != 0 {
			LinkStats[j].timeCapacityLimited = float64(LinkStats[j].timeCapacityLimited + tStep)
		}
		switch Conduit[k].fullState {
		case ALL_FULL:
			LinkStats[j].timeSurcharged = float64(LinkStats[j].timeSurcharged + tStep)
			LinkStats[j].timeFullUpstream = float64(LinkStats[j].timeFullUpstream + tStep)
			LinkStats[j].timeFullDnstream = float64(LinkStats[j].timeFullDnstream + tStep)
			break
		case UP_FULL:
			LinkStats[j].timeFullUpstream = float64(LinkStats[j].timeFullUpstream + tStep)
			break
		case DN_FULL:
			LinkStats[j].timeFullDnstream = float64(LinkStats[j].timeFullDnstream + tStep)
		}
		if Link[j].xsect.type_ == STREET_XSECT {
			if (LinkStats[j].maxStreetFilled) >= (street_getExtentFilled(j)) {
				LinkStats[j].maxStreetFilled = (LinkStats[j].maxStreetFilled)
			} else {
				LinkStats[j].maxStreetFilled = (street_getExtentFilled(j))
			}
		}
	}
	k = LinkStats[j].flowTurnSign
	if (dq) < float64(0) {
		LinkStats[j].flowTurnSign = (-1)
	} else {
		LinkStats[j].flowTurnSign = (1)
	}
	if (M.Fabs(dq) > float64(0.001)) && ((k * LinkStats[j].flowTurnSign) < 0) {
		LinkStats[j].flowTurns++
	}
}

// stats.c:756
func stats_findMaxStats() {
	var (
		j         int32
		x         float64
		z         float64
		stepCount float64
	)
	_, _, _, _ = j, x, z, stepCount
	j = 0
	x = 0
	z = 0
	stepCount = 0
	for j = 0; j < 5; j++ {
		MaxMassBalErrs[j].objType = NODE
		MaxMassBalErrs[j].index = (-1)
		MaxMassBalErrs[j].value = (-(float64(1.0)))
		MaxCourantCrit[j].index = (-1)
		MaxCourantCrit[j].value = (-(float64(1.0)))
		MaxFlowTurns[j].index = (-1)
		MaxFlowTurns[j].value = (-(float64(1.0)))
		MaxNonConverged[j].index = (-1)
		MaxNonConverged[j].value = float64(0.0)
	}
	if ReportStepCount > 2 {
		stepCount = float64(ReportStepCount)
		z = float64(float64(100.0) / (float64(float64(float64(2.0)/float64(3.0)) * (float64(stepCount - float64(2.0))))))
		for j = 0; j < Nobjects[LINK]; j++ {
			x = float64(float64(LinkStats[j].flowTurns) * z)
			stats_updateMaxStats(MaxFlowTurns[:], LINK, j, x)
		}
	}
	for j = 0; j < Nobjects[NODE]; j++ {
		if Node[j].degree <= 0 {
			continue
		}
		if NodeInflow[j] <= float64(0.1) {
			continue
		}
		if NodeInflow[j] > float64(0.0) {
			x = float64(float64(1.0) - float64(NodeOutflow[j]/NodeInflow[j]))
		} else if NodeOutflow[j] > float64(0.0) {
			x = (-(float64(1.0)))
		} else {
			x = float64(0.0)
		}
		stats_updateMaxStats(MaxMassBalErrs[:], NODE, j, float64(float64(100.0)*x))
	}
	stepCount = float64(TimeStepStats.timeStepCount)
	if RouteModel == DW {
		for j = 0; j < Nobjects[NODE]; j++ {
			stats_updateMaxStats(MaxNonConverged[:], NODE, j, float64(float64(NodeStats[j].nonConvergedCount)/stepCount))
		}
	}
	if (RouteModel != DW) || (CourantFactor == float64(0.0)) {
		return
	}
	if stepCount == float64(0) {
		return
	}
	for j = 0; j < Nobjects[NODE]; j++ {
		x = float64(NodeStats[j].timeCourantCritical / stepCount)
		stats_updateMaxStats(MaxCourantCrit[:], NODE, j, float64(float64(100.0)*x))
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		x = float64(LinkStats[j].timeCourantCritical / stepCount)
		stats_updateMaxStats(MaxCourantCrit[:], LINK, j, float64(float64(100.0)*x))
	}
}

// stats.c:841
func stats_updateMaxStats(maxStats []TMaxStats, i int32, j int32, x float64) {
	var (
		k         int32
		maxStats1 TMaxStats
		maxStats2 TMaxStats
	)
	_, _, _ = k, maxStats1, maxStats2
	k = 0
	maxStats1 = (TMaxStats{})
	maxStats2 = (TMaxStats{})
	maxStats1.objType = i
	maxStats1.index = j
	maxStats1.value = x
	for k = 0; k < 5; k++ {
		if M.Fabs(maxStats1.value) > M.Fabs(maxStats[k].value) {
			maxStats2 = maxStats[k]
			maxStats[k] = maxStats1
			maxStats1 = maxStats2
		}
	}
}
