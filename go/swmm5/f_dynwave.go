// Generated from EPA SWMM 5.2.4 (dynwave.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var MINTIMESTEP float64 = float64(0.001)
var dynwave_OMEGA float64 = float64(0.5)
var DEFAULT_SURFAREA float64 = float64(12.566)
var DEFAULT_HEADTOL float64 = float64(0.005)
var EXTRAN_CROWN_CUTOFF float64 = float64(0.96)
var SLOT_CROWN_CUTOFF float64 = float64(0.985257)
var DEFAULT_MAXTRIALS int32 = 8
var VariableStep float64
var Xnode []TXnode
var Omega float64
var Steps int32

// dynwave.c:117
func dynwave_init() {
	var (
		i int32
		j int32
		z float64
	)
	_, _, _ = i, j, z
	i = 0
	j = 0
	z = 0
	VariableStep = float64(0.0)
	Xnode = make([]TXnode, int(uint64(Nobjects[NODE])))
	if Xnode == nil {
		report_writeErrorMsg(ERR_MEMORY, L45)
		return
	}
	for i = 0; i < Nobjects[NODE]; i++ {
		Xnode[i].newSurfArea = float64(0.0)
		Xnode[i].oldSurfArea = float64(0.0)
		Node[i].crownElev = Node[i].invertElev
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		j = Link[i].node1
		z = float64(float64(Node[j].invertElev+Link[i].offset1) + Link[i].xsect.yFull)
		if (Node[j].crownElev) >= (z) {
			Node[j].crownElev = (Node[j].crownElev)
		} else {
			Node[j].crownElev = (z)
		}
		j = Link[i].node2
		z = float64(float64(Node[j].invertElev+Link[i].offset2) + Link[i].xsect.yFull)
		if (Node[j].crownElev) >= (z) {
			Node[j].crownElev = (Node[j].crownElev)
		} else {
			Node[j].crownElev = (z)
		}
		Link[i].flowClass = DRY
		Link[i].dqdh = float64(0.0)
	}
	if SurchargeMethod == SLOT {
		CrownCutoff = SLOT_CROWN_CUTOFF
	} else {
		CrownCutoff = EXTRAN_CROWN_CUTOFF
	}
}

// dynwave.c:165
func dynwave_close() {
	{
		if Xnode != nil {
			_ = Xnode
			Xnode = nil
		}
	}
}

// dynwave.c:177
func dynwave_validate() {
	if MinRouteStep > RouteStep {
		MinRouteStep = RouteStep
	}
	if MinRouteStep < MINTIMESTEP {
		MinRouteStep = MINTIMESTEP
	}
	if MinSurfArea == float64(0.0) {
		MinSurfArea = DEFAULT_SURFAREA
	} else {
		MinSurfArea = float64(MinSurfArea / (float64(UCF(LENGTH) * UCF(LENGTH))))
	}
	if HeadTol == float64(0.0) {
		HeadTol = DEFAULT_HEADTOL
	} else {
		HeadTol = float64(HeadTol / (UCF(LENGTH)))
	}
	if MaxTrials == 0 {
		MaxTrials = DEFAULT_MAXTRIALS
	}
}

// dynwave.c:195
func dynwave_getRoutingStep(fixedStep float64) float64 {
	if CourantFactor == float64(0.0) {
		return fixedStep
	}
	if fixedStep < MINTIMESTEP {
		return fixedStep
	}
	if VariableStep == float64(0.0) {
		VariableStep = MinRouteStep
	} else {
		VariableStep = getVariableStep(fixedStep)
	}
	VariableStep = float64(M.Floor(float64(float64(1000.0)*VariableStep)) / float64(1000.0))
	return VariableStep
}

// dynwave.c:224
func dynwave_execute(tStep float64) int32 {
	var (
		converged int32
	)
	_ = converged
	converged = 0
	if ErrorCode != 0 {
		return 0
	}
	Steps = 0
	converged = 0
	Omega = dynwave_OMEGA
	initRoutingStep()
	for Steps < MaxTrials {
		initNodeStates()
		findLinkFlows(tStep)
		converged = findNodeDepths(tStep)
		Steps++
		if Steps > 1 {
			if converged != 0 {
				break
			}
			findBypassedLinks()
		}
	}
	if !(converged != 0) {
		updateConvergenceStats()
	}
	findLimitedLinks()
	return Steps
}

// dynwave.c:266
func updateConvergenceStats() {
	var (
		i int32
	)
	_ = i
	i = 0
	NonConvergeCount++
	for i = 0; i < Nobjects[NODE]; i++ {
		stats_updateConvergenceStats(i, Xnode[i].converged)
	}
}

// dynwave.c:276
func initRoutingStep() {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 0; i < Nobjects[NODE]; i++ {
		Xnode[i].converged = 0
		Xnode[i].dYdT = float64(0.0)
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		Link[i].bypassed = 0
		Link[i].surfArea1 = float64(0.0)
		Link[i].surfArea2 = float64(0.0)
	}
	for i = 0; i < Nlinks[CONDUIT]; i++ {
		Conduit[i].a2 = Conduit[i].a1
	}
}

// dynwave.c:297
func initNodeStates() {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 0; i < Nobjects[NODE]; i++ {
		if AllowPonding != 0 {
			Xnode[i].newSurfArea = node_getPondedArea(i, Node[i].newDepth)
		} else {
			Xnode[i].newSurfArea = node_getSurfArea(i, Node[i].newDepth)
		}
		Node[i].inflow = float64(0.0)
		Node[i].outflow = Node[i].losses
		if Node[i].newLatFlow >= float64(0.0) {
			Node[i].inflow = float64(Node[i].inflow + Node[i].newLatFlow)
		} else {
			Node[i].outflow = float64(Node[i].outflow - Node[i].newLatFlow)
		}
		Xnode[i].sumdqdh = float64(0.0)
	}
}

// dynwave.c:335
func findBypassedLinks() {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 0; i < Nobjects[LINK]; i++ {
		if (Xnode[Link[i].node1].converged != 0) && (Xnode[Link[i].node2].converged != 0) {
			Link[i].bypassed = 1
		} else {
			Link[i].bypassed = 0
		}
	}
}

// dynwave.c:349
func findLimitedLinks() {
	var (
		j  int32
		n1 int32
		n2 int32
		k  int32
		h1 float64
		h2 float64
	)
	_, _, _, _, _, _ = j, n1, n2, k, h1, h2
	j = 0
	n1 = 0
	n2 = 0
	k = 0
	h1 = 0
	h2 = 0
	for j = 0; j < Nobjects[LINK]; j++ {
		if !(isTrueConduit(j) != 0) {
			continue
		}
		k = Link[j].subIndex
		Conduit[k].capacityLimited = 0
		if Conduit[k].a1 >= Link[j].xsect.aFull {
			n1 = Link[j].node1
			n2 = Link[j].node2
			h1 = float64(Node[n1].newDepth + Node[n1].invertElev)
			h2 = float64(Node[n2].newDepth + Node[n2].invertElev)
			if (float64(h1 - h2)) > float64(M.Fabs(Conduit[k].slope)*Conduit[k].length) {
				Conduit[k].capacityLimited = 1
			}
		}
	}
}

// dynwave.c:382
func findLinkFlows(dt float64) {
	var (
		i int32
	)
	_ = i
	i = 0
	{
		for i = 0; i < Nobjects[LINK]; i++ {
			if (isTrueConduit(i) != 0) && (!(Link[i].bypassed != 0)) {
				dwflow_findConduitFlow(i, Steps, Omega, dt)
			}
		}
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		if isTrueConduit(i) != 0 {
			updateNodeFlows(i)
		}
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		if !(isTrueConduit(i) != 0) {
			if !(Link[i].bypassed != 0) {
				findNonConduitFlow(i, dt)
			}
			updateNodeFlows(i)
		}
	}
}

// dynwave.c:416
func isTrueConduit(j int32) int32 {
	return R.B2I(((Link[j].type_ == CONDUIT) && (Link[j].xsect.type_ != DUMMY)))
}

// dynwave.c:423
func findNonConduitFlow(i int32, dt float64) {
	var (
		qLast float64
		qNew  float64
	)
	_, _ = qLast, qNew
	qLast = 0
	qNew = 0
	qLast = Link[i].newFlow
	Link[i].dqdh = float64(0.0)
	qNew = link_getInflow(i)
	if Link[i].type_ == PUMP {
		qNew = getModPumpFlow(i, qNew, dt)
	}
	findNonConduitSurfArea(i)
	if (Steps > 0) && (Link[i].type_ != PUMP) {
		qNew = float64(float64((float64(float64(1.0)-Omega))*qLast) + float64(Omega*qNew))
		if float64(qNew*qLast) < float64(0.0) {
			qNew = float64(float64(0.001) * float64((R.If[int32](((qNew) < float64(0)), (-1), (1)))))
		}
	}
	Link[i].newFlow = qNew
}

// dynwave.c:458
func getModPumpFlow(i int32, q float64, dt float64) float64 {
	var (
		j             int32
		k             int32
		newNetInflow  float64
		netFlowVolume float64
		y_2           float64
	)
	_, _, _, _, _ = j, k, newNetInflow, netFlowVolume, y_2
	j = Link[i].node1
	k = Link[i].subIndex
	newNetInflow = 0
	netFlowVolume = 0
	y_2 = 0
	if q == float64(0.0) {
		return q
	}
	if Node[j].type_ == STORAGE {
		return node_getMaxOutflow(j, q, dt)
	}
	switch Pump[k].type_ {
	case TYPE1_PUMP:
		return node_getMaxOutflow(j, q, dt)
	case TYPE2_PUMP, TYPE4_PUMP, TYPE3_PUMP:
		newNetInflow = float64(float64(Node[j].inflow-Node[j].outflow) - q)
		netFlowVolume = float64(float64(float64(0.5)*(float64(Node[j].oldNetInflow+newNetInflow))) * dt)
		y_2 = float64(Node[j].oldDepth + float64(netFlowVolume/Xnode[j].newSurfArea))
		if y_2 <= float64(0.0) {
			return Node[j].inflow
		}
	}
	return q
}

// dynwave.c:503
func findNonConduitSurfArea(i int32) {
	if Link[i].type_ == ORIFICE {
		Link[i].surfArea1 = float64(Orifice[Link[i].subIndex].surfArea / float64(2.0))
	} else {
		Link[i].surfArea1 = float64(0.0)
	}
	Link[i].surfArea2 = Link[i].surfArea1
	if (Link[i].flowClass == UP_CRITICAL) || (Node[Link[i].node1].type_ == STORAGE) {
		Link[i].surfArea1 = float64(0.0)
	}
	if (Link[i].flowClass == DN_CRITICAL) || (Node[Link[i].node2].type_ == STORAGE) {
		Link[i].surfArea2 = float64(0.0)
	}
}

// dynwave.c:528
func updateNodeFlows(i int32) {
	var (
		k               int32
		barrels         int32
		n1              int32
		n2              int32
		q               float64
		conduitLossRate float64
	)
	_, _, _, _, _, _ = k, barrels, n1, n2, q, conduitLossRate
	k = 0
	barrels = 1
	n1 = Link[i].node1
	n2 = Link[i].node2
	q = Link[i].newFlow
	conduitLossRate = float64(0.0)
	if q >= float64(0.0) {
		Node[n1].outflow = float64(Node[n1].outflow + q)
		Node[n2].inflow = float64(Node[n2].inflow + q)
	} else {
		Node[n1].inflow = float64(Node[n1].inflow - q)
		Node[n2].outflow = float64(Node[n2].outflow - q)
	}
	if Link[i].type_ == CONDUIT {
		k = Link[i].subIndex
		barrels = Conduit[k].barrels
		conduitLossRate = float64((float64(Conduit[k].evapLossRate + Conduit[k].seepLossRate)) * float64(barrels))
		if conduitLossRate > float64(0.0) {
			if (Node[n1].type_ != OUTFALL) && (Node[n2].type_ != OUTFALL) {
				conduitLossRate = float64(conduitLossRate / (float64(2.0)))
			}
			if Node[n1].type_ != OUTFALL {
				Node[n1].outflow = float64(Node[n1].outflow + conduitLossRate)
			}
			if Node[n2].type_ != OUTFALL {
				Node[n2].outflow = float64(Node[n2].outflow + conduitLossRate)
			}
		}
	}
	Xnode[Link[i].node1].newSurfArea = float64(Xnode[Link[i].node1].newSurfArea + (float64(Link[i].surfArea1 * float64(barrels))))
	Xnode[Link[i].node2].newSurfArea = float64(Xnode[Link[i].node2].newSurfArea + (float64(Link[i].surfArea2 * float64(barrels))))
	Xnode[Link[i].node1].sumdqdh = float64(Xnode[Link[i].node1].sumdqdh + Link[i].dqdh)
	if Link[i].type_ == PUMP {
		k = Link[i].subIndex
		if Pump[k].type_ != TYPE4_PUMP {
			Xnode[n2].sumdqdh = float64(Xnode[n2].sumdqdh + Link[i].dqdh)
		}
	} else {
		Xnode[n2].sumdqdh = float64(Xnode[n2].sumdqdh + Link[i].dqdh)
	}
}

// dynwave.c:593
func findNodeDepths(dt float64) int32 {
	var (
		i    int32
		yOld float64
	)
	_, _ = i, yOld
	i = 0
	yOld = float64(0.0)
	for i = 0; i < Nobjects[LINK]; i++ {
		link_setOutfallDepth(i)
	}
	{
		for i = 0; i < Nobjects[NODE]; i++ {
			if Node[i].type_ == OUTFALL {
				continue
			}
			yOld = Node[i].newDepth
			setNodeDepth(i, dt)
			Xnode[i].converged = 1
			if M.Fabs(float64(yOld-Node[i].newDepth)) > HeadTol {
				Xnode[i].converged = 0
			}
		}
	}
	for i = 0; i < Nobjects[NODE]; i++ {
		if Node[i].type_ == OUTFALL {
			continue
		}
		if Xnode[i].converged == 0 {
			return 0
		}
	}
	return 1
}

// dynwave.c:636
func setNodeDepth(i int32, dt float64) {
	var (
		canPond      int32
		isPonded     int32
		isSurcharged int32
		dQ           float64
		dV           float64
		dy           float64
		yMax         float64
		yOld         float64
		yLast        float64
		yNew         float64
		yCrown       float64
		surfArea     float64
		denom        float64
		corr         float64
		f            float64
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = canPond, isPonded, isSurcharged, dQ, dV, dy, yMax, yOld, yLast, yNew, yCrown, surfArea
	_, _, _ = denom, corr, f
	canPond = 0
	isPonded = 0
	isSurcharged = 0
	dQ = 0
	dV = 0
	dy = 0
	yMax = 0
	yOld = 0
	yLast = 0
	yNew = 0
	yCrown = 0
	surfArea = 0
	denom = 0
	corr = 0
	f = 0
	canPond = R.B2I(((AllowPonding != 0) && (Node[i].pondedArea > float64(0.0))))
	isPonded = R.B2I(((canPond != 0) && (Node[i].newDepth > Node[i].fullDepth)))
	yCrown = float64(Node[i].crownElev - Node[i].invertElev)
	yOld = Node[i].oldDepth
	yLast = Node[i].newDepth
	Node[i].overflow = float64(0.0)
	surfArea = Xnode[i].newSurfArea
	if (surfArea) >= (MinSurfArea) {
		surfArea = (surfArea)
	} else {
		surfArea = (MinSurfArea)
	}
	dQ = float64(Node[i].inflow - Node[i].outflow)
	dV = float64(float64(float64(0.5)*(float64(Node[i].oldNetInflow+dQ))) * dt)
	if SurchargeMethod == EXTRAN {
		if isPonded != 0 {
			isSurcharged = 0
		} else if Node[i].type_ == STORAGE {
			isSurcharged = R.B2I(((Node[i].surDepth > float64(0.0)) && (yLast > Node[i].fullDepth)))
		} else {
			isSurcharged = R.B2I(((yCrown > float64(0.0)) && (yLast > yCrown)))
		}
	}
	if !(isSurcharged != 0) {
		dy = float64(dV / surfArea)
		yNew = float64(yOld + dy)
		if !(isPonded != 0) {
			Xnode[i].oldSurfArea = surfArea
		}
		if Steps > 0 {
			yNew = float64(float64((float64(float64(1.0)-Omega))*yLast) + float64(Omega*yNew))
		}
		if (isPonded != 0) && (yNew < Node[i].fullDepth) {
			yNew = float64(Node[i].fullDepth - float64(0.0001))
		}
	} else {
		corr = float64(1.0)
		if Node[i].degree < 0 {
			corr = float64(0.6)
		}
		denom = Xnode[i].sumdqdh
		if yLast < float64(float64(1.25)*yCrown) {
			f = float64((float64(yLast - yCrown)) / yCrown)
			denom = float64(denom + (float64((float64(float64(Xnode[i].oldSurfArea/dt) - Xnode[i].sumdqdh)) * M.Exp(float64((-(float64(15.0)))*f)))))
		}
		if denom == float64(0.0) {
			dy = float64(0.0)
		} else {
			dy = float64(float64(corr*dQ) / denom)
		}
		yNew = float64(yLast + dy)
		if yNew < yCrown {
			yNew = float64(yCrown - float64(0.0001))
		}
		if (canPond != 0) && (yNew > Node[i].fullDepth) {
			yNew = float64(Node[i].fullDepth + float64(0.0001))
		}
	}
	if yNew < float64(0) {
		yNew = float64(0.0)
	}
	yMax = Node[i].fullDepth
	if canPond == 0 {
		yMax = float64(yMax + Node[i].surDepth)
	}
	if yNew > yMax {
		yNew = getFloodedDepth(i, canPond, dV, yNew, yMax, dt)
	} else {
		Node[i].newVolume = node_getVolume(i, yNew)
	}
	Xnode[i].dYdT = float64(M.Fabs(float64(yNew-yOld)) / dt)
	Node[i].newDepth = yNew
}

// dynwave.c:766
func getFloodedDepth(i int32, canPond int32, dV float64, yNew float64, yMax float64, dt float64) float64 {
	if canPond == 0 {
		Node[i].overflow = float64(dV / dt)
		Node[i].newVolume = Node[i].fullVolume
		yNew = yMax
	} else {
		if (float64(Node[i].oldVolume + dV)) >= (Node[i].fullVolume) {
			Node[i].newVolume = (float64(Node[i].oldVolume + dV))
		} else {
			Node[i].newVolume = (Node[i].fullVolume)
		}
		Node[i].overflow = float64((float64(Node[i].newVolume - (func() float64 {
			if (Node[i].oldVolume) >= (Node[i].fullVolume) {
				return (Node[i].oldVolume)
			}
			return (Node[i].fullVolume)
		}()))) / dt)
	}
	if Node[i].overflow < float64(0.0001) {
		Node[i].overflow = float64(0.0)
	}
	return yNew
}

// dynwave.c:799
func getVariableStep(maxStep float64) float64 {
	var (
		minLink  int32
		minNode  int32
		tMin     float64
		tMinLink float64
		tMinNode float64
	)
	_, _, _, _, _ = minLink, minNode, tMin, tMinLink, tMinNode
	minLink = (-1)
	minNode = (-1)
	tMin = 0
	tMinLink = 0
	tMinNode = 0
	tMin = maxStep
	tMinLink = getLinkStep(tMin, &minLink)
	tMinNode = getNodeStep(tMinLink, &minNode)
	tMin = tMinLink
	if tMinNode < tMin {
		tMin = tMinNode
		minLink = (-1)
	}
	stats_updateCriticalTimeCount(minNode, minLink)
	if tMin < MinRouteStep {
		tMin = MinRouteStep
	}
	return tMin
}

// dynwave.c:836
func getLinkStep(tMin float64, minLink *int32) float64 {
	var (
		i     int32
		k     int32
		q     float64
		t     float64
		tLink float64
	)
	_, _, _, _, _ = i, k, q, t, tLink
	i = 0
	k = 0
	q = 0
	t = 0
	tLink = tMin
	for i = 0; i < Nobjects[LINK]; i++ {
		if Link[i].type_ == CONDUIT {
			k = Link[i].subIndex
			q = float64(M.Fabs(Link[i].newFlow) / float64(Conduit[k].barrels))
			if ((q <= float64(0.0001)) || (Conduit[k].a1 <= float64(0.0001))) || (Link[i].froude <= float64(0.01)) {
				continue
			}
			t = float64(float64(Link[i].newVolume/float64(Conduit[k].barrels)) / q)
			t = float64(float64(t*Conduit[k].modLength) / link_getLength(i))
			t = float64(float64(float64(t*Link[i].froude)/(float64(float64(1.0)+Link[i].froude))) * CourantFactor)
			if t < tLink {
				tLink = t
				*minLink = i
			}
		}
	}
	return tLink
}

// dynwave.c:881
func getNodeStep(tMin float64, minNode *int32) float64 {
	var (
		i        int32
		maxDepth float64
		dYdT     float64
		t1       float64
		tNode    float64
	)
	_, _, _, _, _ = i, maxDepth, dYdT, t1, tNode
	i = 0
	maxDepth = 0
	dYdT = 0
	t1 = 0
	tNode = tMin
	for i = 0; i < Nobjects[NODE]; i++ {
		if Node[i].type_ == OUTFALL {
			continue
		}
		if Node[i].newDepth <= float64(0.0001) {
			continue
		}
		if float64(Node[i].newDepth+float64(0.0001)) >= float64(Node[i].crownElev-Node[i].invertElev) {
			continue
		}
		maxDepth = float64((float64(Node[i].crownElev - Node[i].invertElev)) * float64(0.25))
		if maxDepth < float64(0.0001) {
			continue
		}
		dYdT = Xnode[i].dYdT
		if dYdT < float64(0.0001) {
			continue
		}
		t1 = float64(maxDepth / dYdT)
		if t1 < tNode {
			tNode = t1
			*minNode = i
		}
	}
	return tNode
}
