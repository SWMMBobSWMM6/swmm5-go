// Generated from EPA SWMM 5.2.4 (routing.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var SortedLinks []int32
var NextEvent int32
var BetweenEvents int32
var NewRuleTime float64

// routing.c:93
func routing_open() int32 {
	if !(treatmnt_open() != 0) {
		return ErrorCode
	}
	SortedLinks = nil
	if Nobjects[LINK] > 0 {
		SortedLinks = make([]int32, int(uint64(Nobjects[LINK])))
		if !(SortedLinks != nil) {
			report_writeErrorMsg(ERR_MEMORY, L0)
			return ErrorCode
		}
		toposort_sortLinks(SortedLinks)
		if ErrorCode != 0 {
			return ErrorCode
		}
	}
	iface_openRoutingFiles()
	flowrout_init(RouteModel)
	if Fhotstart1.mode == NO_FILE {
		qualrout_init()
	}
	if NumEvents > 0 {
		sortEvents()
	}
	NextEvent = 0
	BetweenEvents = R.B2I((NumEvents > 0))
	NewRuleTime = float64(0.0)
	return ErrorCode
}

// routing.c:134
func routing_close(routingModel int32) {
	iface_closeRoutingFiles()
	flowrout_close(routingModel)
	treatmnt_close()
	{
		if SortedLinks != nil {
			_ = SortedLinks
			SortedLinks = nil
		}
	}
}

// routing.c:152
func routing_getRoutingStep(routingModel int32, fixedStep float64) float64 {
	var (
		date1           float64
		date2           float64
		nextTime        float64
		routingStep     float64
		nextRuleTime    float64
		nextRoutingTime float64
	)
	_, _, _, _, _, _ = date1, date2, nextTime, routingStep, nextRuleTime, nextRoutingTime
	date1 = 0
	date2 = 0
	nextTime = 0
	routingStep = float64(0.0)
	nextRuleTime = 0
	nextRoutingTime = 0
	if Nobjects[LINK] == 0 {
		return fixedStep
	}
	if (NumEvents > 0) && (BetweenEvents != 0) {
		if (NewRunoffTime) <= (ReportTime) {
			nextTime = (NewRunoffTime)
		} else {
			nextTime = (ReportTime)
		}
		date1 = getDateTime(NewRoutingTime)
		date2 = getDateTime(nextTime)
		if (date2 > date1) && (date2 < Event[NextEvent].start) {
			routingStep = float64((float64(nextTime - NewRoutingTime)) / float64(1000.0))
		} else {
			date1 = getDateTime(float64(NewRoutingTime + float64(float64(1000.0)*fixedStep)))
			if date1 < Event[NextEvent].start {
				return fixedStep
			}
		}
	}
	if routingStep == float64(0.0) {
		routingStep = flowrout_getRoutingStep(routingModel, fixedStep)
	}
	if RuleStep > 0 {
		nextRuleTime = float64(NewRuleTime + float64(float64(1000.0)*float64(RuleStep)))
		nextRoutingTime = float64(NewRoutingTime + float64(float64(1000.0)*routingStep))
		if nextRoutingTime >= nextRuleTime {
			routingStep = float64((float64(nextRuleTime - NewRoutingTime)) / float64(1000.0))
		}
	}
	return routingStep
}

// routing.c:203
func routing_execute(routingModel int32, routingStep float64) {
	var (
		trialsCount   int32
		actionCount   int32
		inSteadyState int32
		currentDate   float64
		stepFlowError float64
	)
	_, _, _, _, _ = trialsCount, actionCount, inSteadyState, currentDate, stepFlowError
	trialsCount = 1
	actionCount = 0
	inSteadyState = 1
	currentDate = 0
	stepFlowError = 0
	if ErrorCode != 0 {
		return
	}
	massbal_updateRoutingTotals(float64(routingStep / float64(2.0)))
	currentDate = getDateTime(NewRoutingTime)
	actionCount = evaluateControlRules(currentDate, routingStep)
	stepFlowError = massbal_getStepFlowError()
	massbal_initTimeStepTotals()
	initSystemInflows()
	BetweenEvents = isBetweenEvents(currentDate)
	if BetweenEvents == 0 {
		addSystemInflows(currentDate, routingStep)
		inlet_findCapturedFlows(routingStep)
		inSteadyState = isInSteadyState(actionCount, stepFlowError)
		if inSteadyState == 0 {
			trialsCount = routeFlow(routingModel, routingStep)
		}
		if (Nobjects[POLLUT] > 0) && (!(IgnoreQuality != 0)) {
			inlet_adjustQualInflows()
			qualrout_execute(routingStep)
		}
		removeSystemOutflows(routingStep)
		inlet_adjustQualOutflows()
		if Nobjects[LINK] > 0 {
			stats_updateFlowStats(routingStep, getDateTime(NewRoutingTime))
			stats_updateTimeStepStats(routingStep, trialsCount, inSteadyState)
		}
	}
	massbal_updateRoutingTotals(float64(routingStep / float64(2.0)))
}

// routing.c:269
func evaluateControlRules(currentDate float64, routingStep float64) int32 {
	var (
		j           int32
		actionCount int32
	)
	_, _ = j, actionCount
	j = 0
	actionCount = 0
	for j = 0; j < Nobjects[LINK]; j++ {
		link_setTargetSetting(j)
	}
	if (RuleStep == 0) || (M.Fabs(float64(NewRoutingTime-NewRuleTime)) < float64(1.0)) {
		controls_evaluate(currentDate, float64(currentDate-StartDateTime), float64(routingStep/float64(86400)))
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		if Link[j].targetSetting != Link[j].setting {
			if float64(Link[j].targetSetting*Link[j].setting) == float64(0.0) {
				Link[j].timeLastSet = currentDate
			}
			link_setSetting(j, routingStep)
			actionCount++
		}
	}
	OldRoutingTime = NewRoutingTime
	NewRoutingTime = float64(NewRoutingTime + float64(float64(1000.0)*routingStep))
	if M.Fabs(float64(NewRoutingTime-(float64(NewRuleTime+float64(float64(1000.0)*float64(RuleStep)))))) < float64(1) {
		NewRuleTime = float64(NewRuleTime + (float64(float64(1000.0) * float64(RuleStep))))
	}
	return actionCount
}

// routing.c:312
func initSystemInflows() {
	var (
		j int32
	)
	_ = j
	j = 0
	if Nobjects[POLLUT] > 0 {
		for j = 0; j < Nobjects[NODE]; j++ {
			node_setOldQualState(j)
		}
		for j = 0; j < Nobjects[LINK]; j++ {
			link_setOldQualState(j)
		}
	}
	infil_setInfilFactor((-1))
	for j = 0; j < Nobjects[NODE]; j++ {
		Node[j].oldLatFlow = Node[j].newLatFlow
		Node[j].newLatFlow = float64(0.0)
	}
}

// routing.c:337
func isBetweenEvents(currentDate float64) int32 {
	if NumEvents == 0 {
		return 0
	}
	if currentDate > Event[NextEvent].end {
		NextEvent++
		return 1
	} else if currentDate >= Event[NextEvent].start {
		return 0
	}
	return 1
}

// routing.c:359
func addSystemInflows(currentDate float64, routingStep float64) {
	var (
		j int32
	)
	_ = j
	j = 0
	for j = 0; j < Nobjects[NODE]; j++ {
		Node[j].losses = node_getLosses(j, routingStep)
	}
	addExternalInflows(currentDate)
	addDryWeatherInflows(currentDate)
	addWetWeatherInflows(OldRoutingTime)
	addGroundwaterInflows(OldRoutingTime)
	addLidDrainInflows(OldRoutingTime)
	addRdiiInflows(currentDate)
	addIfaceInflows(currentDate)
	for j = 0; j < Nobjects[NODE]; j++ {
		if (float64(0.0)) >= (Node[j].newLatFlow) {
			Node[j].qualInflow = (float64(0.0))
		} else {
			Node[j].qualInflow = (Node[j].newLatFlow)
		}
	}
}

// routing.c:383
func isInSteadyState(actionCount int32, stepFlowError float64) int32 {
	if SkipSteadyState != 0 {
		if (((OldRoutingTime == float64(0.0)) || (actionCount > 0)) || (M.Fabs(stepFlowError) > SysFlowTol)) || (inflowHasChanged() != 0) {
			return 0
		} else {
			return 1
		}
	}
	return 0
}

// routing.c:399
func routeFlow(routingModel int32, routingStep float64) int32 {
	var (
		j         int32
		stepCount int32
	)
	_, _ = j, stepCount
	j = 0
	stepCount = 1
	for j = 0; j < Nobjects[LINK]; j++ {
		link_setOldHydState(j)
	}
	for j = 0; j < Nobjects[NODE]; j++ {
		node_setOldHydState(j)
	}
	for j = 0; j < Nobjects[NODE]; j++ {
		node_initFlows(j, routingStep)
	}
	if Nobjects[LINK] > 0 {
		stepCount = flowrout_execute(SortedLinks, routingModel, routingStep)
	}
	return stepCount
}

// routing.c:425
func removeSystemOutflows(routingStep float64) {
	removeStorageLosses(routingStep)
	removeConduitLosses()
	removeOutflows(routingStep)
}

// routing.c:435
func addExternalInflows(currentDate float64) {
	var (
		j      int32
		p      int32
		q      float64
		w      float64
		inflow *TExtInflow
	)
	_, _, _, _, _ = j, p, q, w, inflow
	j = 0
	p = 0
	q = 0
	w = 0
	inflow = nil
	for j = 0; j < Nobjects[NODE]; j++ {
		q = Node[j].apiExtInflow
		inflow = Node[j].extInflow
		for inflow != nil {
			if inflow.type_ == FLOW_INFLOW {
				q = float64(q + (inflow_getExtInflow(inflow, currentDate)))
				break
			} else {
				inflow = inflow.next
			}
		}
		if M.Fabs(q) < float64(1e-05) {
			q = float64(0.0)
		}
		Node[j].newLatFlow = float64(Node[j].newLatFlow + q)
		if q >= float64(0.0) {
			massbal_addInflowFlow(EXTERNAL_INFLOW, q)
		} else {
			massbal_addOutflowFlow((-q), 0)
			continue
		}
		if (Node[j].type_ == OUTFALL) && (Node[j].oldNetInflow < float64(0.0)) {
			q = float64(q - Node[j].oldNetInflow)
		}
		inflow = Node[j].extInflow
		for inflow != nil {
			if inflow.type_ != FLOW_INFLOW {
				p = inflow.param
				w = inflow_getExtInflow(inflow, currentDate)
				if inflow.type_ == CONCEN_INFLOW {
					w = float64(w * q)
				}
				Node[j].newQual[p] = float64(Node[j].newQual[p] + w)
				massbal_addInflowQual(EXTERNAL_INFLOW, p, w)
			}
			inflow = inflow.next
		}
	}
}

// routing.c:498
func addDryWeatherInflows(currentDate float64) {
	var (
		j      int32
		p      int32
		month  int32
		day    int32
		hour   int32
		q      float64
		w      float64
		inflow *TDwfInflow
	)
	_, _, _, _, _, _, _, _ = j, p, month, day, hour, q, w, inflow
	j = 0
	p = 0
	month = 0
	day = 0
	hour = 0
	q = 0
	w = 0
	inflow = nil
	month = (datetime_monthOfYear(currentDate) - 1)
	day = (datetime_dayOfWeek(currentDate) - 1)
	hour = datetime_hourOfDay(currentDate)
	for j = 0; j < Nobjects[NODE]; j++ {
		inflow = Node[j].dwfInflow
		if !(inflow != nil) {
			continue
		}
		q = float64(0.0)
		for inflow != nil {
			if inflow.param < 0 {
				q = inflow_getDwfInflow(inflow, month, day, hour)
				break
			}
			inflow = inflow.next
		}
		if M.Fabs(q) < float64(1e-05) {
			q = float64(0.0)
		}
		Node[j].newLatFlow = float64(Node[j].newLatFlow + q)
		massbal_addInflowFlow(DRY_WEATHER_INFLOW, q)
		if q <= float64(0.0) {
			continue
		}
		for p = 0; p < Nobjects[POLLUT]; p++ {
			if Pollut[p].dwfConcen > float64(0.0) {
				w = float64(q * Pollut[p].dwfConcen)
				Node[j].newQual[p] = float64(Node[j].newQual[p] + w)
				massbal_addInflowQual(DRY_WEATHER_INFLOW, p, w)
			}
		}
		inflow = Node[j].dwfInflow
		for inflow != nil {
			if inflow.param >= 0 {
				p = inflow.param
				w = float64(q * inflow_getDwfInflow(inflow, month, day, hour))
				Node[j].newQual[p] = float64(Node[j].newQual[p] + w)
				massbal_addInflowQual(DRY_WEATHER_INFLOW, p, w)
				if Pollut[p].dwfConcen > float64(0.0) {
					w = float64(q * Pollut[p].dwfConcen)
					Node[j].newQual[p] = float64(Node[j].newQual[p] - w)
					massbal_addInflowQual(DRY_WEATHER_INFLOW, p, (-w))
				}
			}
			inflow = inflow.next
		}
	}
}

// routing.c:579
func addWetWeatherInflows(routingTime float64) {
	var (
		i int32
		j int32
		p int32
		q float64
		w float64
		f float64
	)
	_, _, _, _, _, _ = i, j, p, q, w, f
	i = 0
	j = 0
	p = 0
	q = 0
	w = 0
	f = 0
	if Nobjects[SUBCATCH] == 0 {
		return
	}
	f = float64((float64(routingTime - OldRunoffTime)) / (float64(NewRunoffTime - OldRunoffTime)))
	if f < float64(0.0) {
		f = float64(0.0)
	}
	if f > float64(1.0) {
		f = float64(1.0)
	}
	for i = 0; i < Nobjects[SUBCATCH]; i++ {
		j = Subcatch[i].outNode
		if j >= 0 {
			q = subcatch_getWtdOutflow(i, f)
			Node[j].newLatFlow = float64(Node[j].newLatFlow + q)
			massbal_addInflowFlow(WET_WEATHER_INFLOW, q)
			for p = 0; p < Nobjects[POLLUT]; p++ {
				w = surfqual_getWtdWashoff(i, p, f)
				Node[j].newQual[p] = float64(Node[j].newQual[p] + w)
				massbal_addInflowQual(WET_WEATHER_INFLOW, p, w)
			}
		}
	}
}

// routing.c:621
func addGroundwaterInflows(routingTime float64) {
	var (
		i  int32
		j  int32
		p  int32
		q  float64
		w  float64
		f  float64
		gw *TGroundwater
	)
	_, _, _, _, _, _, _ = i, j, p, q, w, f, gw
	i = 0
	j = 0
	p = 0
	q = 0
	w = 0
	f = 0
	gw = nil
	if Nobjects[SUBCATCH] == 0 {
		return
	}
	f = float64((float64(routingTime - OldRunoffTime)) / (float64(NewRunoffTime - OldRunoffTime)))
	if f < float64(0.0) {
		f = float64(0.0)
	}
	if f > float64(1.0) {
		f = float64(1.0)
	}
	for i = 0; i < Nobjects[SUBCATCH]; i++ {
		gw = Subcatch[i].groundwater
		if gw != nil {
			j = gw.node
			if j >= 0 {
				q = float64((float64(float64((float64(float64(1.0)-f))*(gw.oldFlow)) + float64(f*(gw.newFlow)))) * Subcatch[i].area)
				if M.Fabs(q) < float64(1e-05) {
					continue
				}
				Node[j].newLatFlow = float64(Node[j].newLatFlow + q)
				massbal_addInflowFlow(GROUNDWATER_INFLOW, q)
				if q > float64(0.0) {
					for p = 0; p < Nobjects[POLLUT]; p++ {
						w = float64(q * Pollut[p].gwConcen)
						Node[j].newQual[p] = float64(Node[j].newQual[p] + w)
						massbal_addInflowQual(GROUNDWATER_INFLOW, p, w)
					}
				}
			}
		}
	}
}

// routing.c:674
func addLidDrainInflows(routingTime float64) {
	var (
		j int32
		f float64
	)
	_, _ = j, f
	j = 0
	f = 0
	if Nobjects[SUBCATCH] == 0 {
		return
	}
	f = float64((float64(routingTime - OldRunoffTime)) / (float64(NewRunoffTime - OldRunoffTime)))
	if f < float64(0.0) {
		f = float64(0.0)
	}
	if f > float64(1.0) {
		f = float64(1.0)
	}
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		if (Subcatch[j].area > float64(0.0)) && (Subcatch[j].lidArea > float64(0.0)) {
			lid_addDrainInflow(j, f)
		}
	}
}

// routing.c:698
func addRdiiInflows(currentDate float64) {
	var (
		i            int32
		j            int32
		p            int32
		q            float64
		w            float64
		numRdiiNodes int32
	)
	_, _, _, _, _, _ = i, j, p, q, w, numRdiiNodes
	i = 0
	j = 0
	p = 0
	q = 0
	w = 0
	numRdiiNodes = 0
	numRdiiNodes = rdii_getNumRdiiFlows(currentDate)
	for i = 0; i < numRdiiNodes; i++ {
		rdii_getRdiiFlow(i, &j, &q)
		if j < 0 {
			continue
		}
		if M.Fabs(q) < float64(1e-05) {
			continue
		}
		Node[j].newLatFlow = float64(Node[j].newLatFlow + q)
		massbal_addInflowFlow(RDII_INFLOW, q)
		if q > float64(0.0) {
			for p = 0; p < Nobjects[POLLUT]; p++ {
				w = float64(q * Pollut[p].rdiiConcen)
				Node[j].newQual[p] = float64(Node[j].newQual[p] + w)
				massbal_addInflowQual(RDII_INFLOW, p, w)
			}
		}
	}
}

// routing.c:736
func addIfaceInflows(currentDate float64) {
	var (
		i             int32
		j             int32
		p             int32
		q             float64
		w             float64
		numIfaceNodes int32
	)
	_, _, _, _, _, _ = i, j, p, q, w, numIfaceNodes
	i = 0
	j = 0
	p = 0
	q = 0
	w = 0
	numIfaceNodes = 0
	if Finflows.mode != USE_FILE {
		return
	}
	numIfaceNodes = iface_getNumIfaceNodes(currentDate)
	for i = 0; i < numIfaceNodes; i++ {
		j = iface_getIfaceNode(i)
		if j < 0 {
			continue
		}
		q = iface_getIfaceFlow(i)
		if M.Fabs(q) < float64(1e-05) {
			continue
		}
		Node[j].newLatFlow = float64(Node[j].newLatFlow + q)
		massbal_addInflowFlow(EXTERNAL_INFLOW, q)
		if q > float64(0.0) {
			for p = 0; p < Nobjects[POLLUT]; p++ {
				w = float64(q * iface_getIfaceQual(i, p))
				Node[j].newQual[p] = float64(Node[j].newQual[p] + w)
				massbal_addInflowQual(EXTERNAL_INFLOW, p, w)
			}
		}
	}
}

// routing.c:776
func inflowHasChanged() int32 {
	var (
		j    int32
		diff float64
		qOld float64
		qNew float64
	)
	_, _, _, _ = j, diff, qOld, qNew
	j = 0
	diff = 0
	qOld = 0
	qNew = 0
	for j = 0; j < Nobjects[NODE]; j++ {
		qOld = Node[j].oldLatFlow
		qNew = Node[j].newLatFlow
		if M.Fabs(qOld) > float64(1e-06) {
			diff = float64((float64(qNew / qOld)) - float64(1.0))
		} else if M.Fabs(qNew) > float64(1e-06) {
			diff = float64(1.0)
		} else {
			diff = float64(0.0)
		}
		if M.Fabs(diff) > LatFlowTol {
			return 1
		}
		if (Node[j].type_ == OUTFALL) || (Node[j].degree == 0) {
			qOld = Node[j].oldFlowInflow
			qNew = Node[j].inflow
			if M.Fabs(qOld) > float64(1e-06) {
				diff = float64((float64(qNew / qOld)) - float64(1.0))
			} else if M.Fabs(qNew) > float64(1e-06) {
				diff = float64(1.0)
			} else {
				diff = float64(0.0)
			}
			if M.Fabs(diff) > LatFlowTol {
				return 1
			}
		}
	}
	return 0
}

// routing.c:812
func removeStorageLosses(tStep float64) {
	var (
		i         int32
		evapLoss  float64
		exfilLoss float64
	)
	_, _, _ = i, evapLoss, exfilLoss
	i = 0
	evapLoss = float64(0.0)
	exfilLoss = float64(0.0)
	for i = 0; i < Nobjects[NODE]; i++ {
		if Node[i].type_ == STORAGE {
			evapLoss = float64(evapLoss + Storage[Node[i].subIndex].evapLoss)
			exfilLoss = float64(exfilLoss + Storage[Node[i].subIndex].exfilLoss)
		}
	}
	massbal_addNodeLosses(float64(evapLoss/tStep), float64(exfilLoss/tStep))
}

// routing.c:841
func removeConduitLosses() {
	var (
		i        int32
		k        int32
		barrels  float64
		evapLoss float64
		seepLoss float64
	)
	_, _, _, _, _ = i, k, barrels, evapLoss, seepLoss
	i = 0
	k = 0
	barrels = 0
	evapLoss = float64(0.0)
	seepLoss = float64(0.0)
	for i = 0; i < Nobjects[LINK]; i++ {
		if Link[i].type_ == CONDUIT {
			k = Link[i].subIndex
			barrels = float64(Conduit[k].barrels)
			evapLoss = float64(evapLoss + (float64(Conduit[k].evapLossRate * barrels)))
			seepLoss = float64(seepLoss + (float64(Conduit[k].seepLossRate * barrels)))
		}
	}
	massbal_addLinkLosses(evapLoss, seepLoss)
}

// routing.c:872
func removeOutflows(tStep float64) {
	var (
		i         int32
		p         int32
		k         int32
		isFlooded int32
		q         float64
		w         float64
		v         float64
	)
	_, _, _, _, _, _, _ = i, p, k, isFlooded, q, w, v
	i = 0
	p = 0
	k = 0
	isFlooded = 0
	q = 0
	w = 0
	v = 0
	for i = 0; i < Nobjects[NODE]; i++ {
		if (Node[i].type_ == OUTFALL) && (Node[i].inflow > float64(0.0)) {
			k = Node[i].subIndex
			if Outfall[k].routeTo >= 0 {
				v = float64(Node[i].inflow * tStep)
				Outfall[k].vRouted = float64(Outfall[k].vRouted + v)
				for p = 0; p < Nobjects[POLLUT]; p++ {
					Outfall[k].wRouted[p] = float64(Outfall[k].wRouted[p] + (float64(Node[i].newQual[p] * v)))
				}
			}
		}
		q = node_getSystemOutflow(i, &isFlooded)
		if q > float64(0.0) {
			massbal_addOutflowFlow(q, isFlooded)
			for p = 0; p < Nobjects[POLLUT]; p++ {
				w = float64(q * Node[i].newQual[p])
				massbal_addOutflowQual(p, w, isFlooded)
			}
		} else {
			massbal_addInflowFlow(EXTERNAL_INFLOW, (-q))
		}
		q = Node[i].newLatFlow
		if q < float64(0.0) {
			for p = 0; p < Nobjects[POLLUT]; p++ {
				w = float64((-q) * Node[i].newQual[p])
				massbal_addOutflowQual(p, w, 0)
			}
		}
	}
}

// routing.c:929
func sortEvents() {
	var (
		i    int32
		j    int32
		gap  int32
		temp TEvent
	)
	_, _, _, _ = i, j, gap, temp
	i = 0
	j = 0
	gap = 0
	temp = (TEvent{})
	for gap = (NumEvents / 2); gap >= 1; gap /= 2 {
		for i = gap; i < NumEvents; i += gap {
			temp = Event[i]
			j = (i - gap)
			for (j >= 0) && (Event[j].start > temp.start) {
				Event[(j + gap)] = Event[j]
				j -= gap
			}
			if j != (i - gap) {
				Event[(j + gap)] = temp
			}
		}
	}
	for i = 0; i < (NumEvents - 1); i++ {
		if Event[i].end > (Event[(i + 1)]).start {
			Event[i].end = (Event[(i + 1)]).start
		}
	}
}
