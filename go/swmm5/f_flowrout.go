// Generated from EPA SWMM 5.2.4 (flowrout.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var flowrout_OMEGA float64 = float64(0.55)
var MAXITER int32 = 10
var STOPTOL float64 = float64(0.005)

// flowrout.c:75
func flowrout_init(routingModel int32) {
	if routingModel == DW {
		validateGeneralLayout()
		dynwave_init()
		if Fhotstart1.mode == NO_FILE {
			initNodeDepths()
			initLinkDepths()
		}
	} else {
		validateTreeLayout()
	}
	initNodes()
	initLinks(routingModel)
}

// flowrout.c:107
func flowrout_close(routingModel int32) {
	if routingModel == DW {
		dynwave_close()
	}
}

// flowrout.c:119
func flowrout_getRoutingStep(routingModel int32, fixedStep float64) float64 {
	if routingModel == DW {
		return dynwave_getRoutingStep(fixedStep)
	}
	return fixedStep
}

// flowrout.c:136
func flowrout_execute(links []int32, routingModel int32, tStep float64) int32 {
	var (
		i     int32
		j     int32
		n1    int32
		qin   float64
		qout  float64
		steps float64
	)
	_, _, _, _, _, _ = i, j, n1, qin, qout, steps
	i = 0
	j = 0
	n1 = 0
	qin = 0
	qout = 0
	steps = 0
	if ErrorCode != 0 {
		return 0
	}
	for j = 0; j < Nobjects[NODE]; j++ {
		Node[j].updated = 0
		Node[j].overflow = float64(0.0)
		if (Node[j].type_ != STORAGE) && (Node[j].newVolume > Node[j].fullVolume) {
			Node[j].overflow = float64((float64(Node[j].newVolume - Node[j].fullVolume)) / tStep)
		}
	}
	if routingModel == DW {
		return dynwave_execute(tStep)
	}
	steps = float64(0.0)
	for i = 0; i < Nobjects[LINK]; i++ {
		j = links[i]
		n1 = Link[j].node1
		if Node[n1].type_ == STORAGE {
			updateStorageState(n1, i, links, tStep)
		}
		qin = getLinkInflow(j, tStep)
		if routingModel == SF {
			steps = float64(steps + (float64(steadyflow_execute(j, &qin, &qout, tStep))))
		} else {
			steps = float64(steps + (float64(kinwave_execute(j, &qin, &qout, tStep))))
		}
		Link[j].newFlow = qout
		Node[Link[j].node1].outflow = float64(Node[Link[j].node1].outflow + qin)
		Node[Link[j].node2].inflow = float64(Node[Link[j].node2].inflow + qout)
	}
	if Nobjects[LINK] > 0 {
		steps = float64(steps / (float64(Nobjects[LINK])))
	}
	for j = 0; j < Nobjects[NODE]; j++ {
		setNewNodeState(j, tStep)
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		setNewLinkState(j)
	}
	return R.D2I((float64(steps + float64(0.5))))
}

// flowrout.c:203
func validateTreeLayout() {
	var (
		j int32
	)
	_ = j
	j = 0
	for j = 0; j < Nobjects[NODE]; j++ {
		switch Node[j].type_ {
		case DIVIDER:
			if Node[j].degree > 2 {
				report_writeErrorMsg(ERR_DIVIDER, Node[j].ID)
			}
			break
		case OUTFALL:
			if Node[j].degree > 0 {
				report_writeErrorMsg(ERR_OUTFALL, Node[j].ID)
			}
			break
		case STORAGE:
			break
		default:
			if Node[j].degree > 1 {
				report_writeErrorMsg(ERR_MULTI_OUTLET, Node[j].ID)
			}
		}
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		switch Link[j].type_ {
		case CONDUIT:
			if (Conduit[Link[j].subIndex].slope < float64(0.0)) && (Link[j].xsect.type_ != DUMMY) {
				report_writeErrorMsg(ERR_SLOPE, Link[j].ID)
			}
			break
		case ORIFICE, WEIR, OUTLET:
			if Node[Link[j].node1].type_ != STORAGE {
				report_writeErrorMsg(ERR_REGULATOR, Link[j].ID)
			}
		}
	}
}

// flowrout.c:274
func validateGeneralLayout() {
	var (
		i           int32
		j           int32
		outletCount int32
	)
	_, _, _ = i, j, outletCount
	i = 0
	j = 0
	outletCount = 0
	for i = 0; i < Nobjects[NODE]; i++ {
		Node[i].inflow = float64(0.0)
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		i = Link[j].node1
		if Node[i].type_ != OUTFALL {
			i = Link[j].node2
		}
		Node[i].inflow = float64(Node[i].inflow + (float64(1.0)))
		if ((Link[j].type_ == CONDUIT) && (Link[j].xsect.type_ == DUMMY)) || ((Link[j].type_ == PUMP) && (Pump[Link[j].subIndex].type_ == IDEAL_PUMP)) {
			i = Link[j].node1
			if Link[j].direction < 0 {
				i = Link[j].node2
			}
			if Node[i].degree > 1 {
				report_writeErrorMsg(ERR_DUMMY_LINK, Node[i].ID)
			}
		}
	}
	for i = 0; i < Nobjects[NODE]; i++ {
		if Node[i].type_ == OUTFALL {
			if (Node[i].degree + R.D2I(Node[i].inflow)) > 1 {
				report_writeErrorMsg(ERR_OUTFALL, Node[i].ID)
			} else {
				outletCount++
			}
		}
	}
	if outletCount == 0 {
		report_writeErrorMsg(ERR_NO_OUTLETS, L0)
	}
	for i = 0; i < Nobjects[NODE]; i++ {
		if Node[i].inflow == float64(0.0) {
			Node[i].degree = (-Node[i].degree)
		}
		Node[i].inflow = float64(0.0)
	}
}

// flowrout.c:337
func initNodeDepths() {
	var (
		i   int32
		n_2 int32
		y_2 float64
	)
	_, _, _ = i, n_2, y_2
	i = 0
	n_2 = 0
	y_2 = 0
	for i = 0; i < Nobjects[NODE]; i++ {
		Node[i].inflow = float64(0.0)
		Node[i].outflow = float64(0.0)
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		if Link[i].newDepth > float64(0.0001) {
			y_2 = float64(Link[i].newDepth + Link[i].offset1)
		} else {
			y_2 = float64(0.0)
		}
		n_2 = Link[i].node1
		Node[n_2].inflow = float64(Node[n_2].inflow + y_2)
		Node[n_2].outflow = float64(Node[n_2].outflow + (float64(1.0)))
		n_2 = Link[i].node2
		Node[n_2].inflow = float64(Node[n_2].inflow + y_2)
		Node[n_2].outflow = float64(Node[n_2].outflow + (float64(1.0)))
	}
	for i = 0; i < Nobjects[NODE]; i++ {
		if Node[i].type_ == OUTFALL {
			continue
		}
		if Node[i].type_ == STORAGE {
			continue
		}
		if Node[i].initDepth > float64(0.0) {
			continue
		}
		if Node[i].outflow > float64(0.0) {
			Node[i].newDepth = float64(Node[i].inflow / Node[i].outflow)
		}
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		link_setOutfallDepth(i)
	}
}

// flowrout.c:389
func initLinkDepths() {
	var (
		i   int32
		y_2 float64
		y1  float64
		y2  float64
	)
	_, _, _, _ = i, y_2, y1, y2
	i = 0
	y_2 = 0
	y1 = 0
	y2 = 0
	for i = 0; i < Nobjects[LINK]; i++ {
		if Link[i].type_ == CONDUIT {
			if Link[i].q0 != float64(0.0) {
				continue
			}
			y1 = float64(Node[Link[i].node1].newDepth - Link[i].offset1)
			if (y1) >= (float64(0.0)) {
				y1 = (y1)
			} else {
				y1 = (float64(0.0))
			}
			if (y1) <= (Link[i].xsect.yFull) {
				y1 = (y1)
			} else {
				y1 = (Link[i].xsect.yFull)
			}
			y2 = float64(Node[Link[i].node2].newDepth - Link[i].offset2)
			if (y2) >= (float64(0.0)) {
				y2 = (y2)
			} else {
				y2 = (float64(0.0))
			}
			if (y2) <= (Link[i].xsect.yFull) {
				y2 = (y2)
			} else {
				y2 = (Link[i].xsect.yFull)
			}
			y_2 = float64(float64(0.5) * (float64(y1 + y2)))
			if (y_2) >= (float64(0.0001)) {
				y_2 = (y_2)
			} else {
				y_2 = (float64(0.0001))
			}
			Link[i].newDepth = y_2
		}
	}
}

// flowrout.c:425
func initNodes() {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 0; i < Nobjects[NODE]; i++ {
		Node[i].inflow = Node[i].newLatFlow
		Node[i].outflow = float64(0.0)
		Node[i].newVolume = float64(0.0)
		if ((AllowPonding != 0) && (Node[i].pondedArea > float64(0.0))) && (Node[i].newDepth > Node[i].fullDepth) {
			Node[i].newVolume = float64(Node[i].fullVolume + float64((float64(Node[i].newDepth-Node[i].fullDepth))*Node[i].pondedArea))
		} else {
			Node[i].newVolume = node_getVolume(i, Node[i].newDepth)
		}
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		if Link[i].newFlow >= float64(0.0) {
			Node[Link[i].node1].outflow = float64(Node[Link[i].node1].outflow + Link[i].newFlow)
			Node[Link[i].node2].inflow = float64(Node[Link[i].node2].inflow + Link[i].newFlow)
		} else {
			Node[Link[i].node1].inflow = float64(Node[Link[i].node1].inflow - Link[i].newFlow)
			Node[Link[i].node2].outflow = float64(Node[Link[i].node2].outflow - Link[i].newFlow)
		}
	}
}

// flowrout.c:472
func initLinks(routingModel int32) {
	var (
		i int32
		k int32
	)
	_, _ = i, k
	i = 0
	k = 0
	for i = 0; i < Nobjects[LINK]; i++ {
		if routingModel == SF {
			Link[i].newFlow = float64(0.0)
		} else if Link[i].type_ == CONDUIT {
			k = Link[i].subIndex
			Conduit[k].q1 = float64(Link[i].newFlow / float64(Conduit[k].barrels))
			Conduit[k].q2 = Conduit[k].q1
			Conduit[k].a1 = xsect_getAofY(&Link[i].xsect, Link[i].newDepth)
			Conduit[k].a2 = Conduit[k].a1
			{
				Link[i].newVolume = float64(float64(Conduit[k].a1*link_getLength(i)) * float64(Conduit[k].barrels))
			}
			Link[i].oldVolume = Link[i].newVolume
		}
	}
}

// flowrout.c:511
func getLinkInflow(j int32, dt float64) float64 {
	var (
		n1 int32
		q  float64
	)
	_, _ = n1, q
	n1 = Link[j].node1
	q = 0
	if ((Link[j].type_ == CONDUIT) || (Link[j].type_ == PUMP)) || (Node[n1].type_ == STORAGE) {
		q = link_getInflow(j)
	} else {
		q = float64(0.0)
	}
	return node_getMaxOutflow(n1, q, dt)
}

// flowrout.c:531
func updateStorageState(i int32, j int32, links []int32, dt float64) {
	var (
		iter    int32
		stopped int32
		vFixed  float64
		v2      float64
		d1      float64
		d2      float64
	)
	_, _, _, _, _, _ = iter, stopped, vFixed, v2, d1, d2
	iter = 0
	stopped = 0
	vFixed = 0
	v2 = 0
	d1 = 0
	d2 = 0
	if Node[i].type_ != STORAGE {
		return
	}
	if Node[i].updated != 0 {
		return
	}
	vFixed = float64(Node[i].oldVolume + float64(float64(float64(0.5)*(float64(float64(Node[i].oldNetInflow+Node[i].inflow)-Node[i].outflow)))*dt))
	d1 = Node[i].newDepth
	iter = 1
	stopped = 0
	for (iter < MAXITER) && (!(stopped != 0)) {
		v2 = float64(vFixed - float64(float64(float64(0.5)*getStorageOutflow(i, j, links, dt))*dt))
		if (float64(0.0)) >= (v2) {
			v2 = (float64(0.0))
		} else {
			v2 = (v2)
		}
		Node[i].overflow = float64(0.0)
		if v2 > Node[i].fullVolume {
			Node[i].overflow = float64((float64(v2 - (func() float64 {
				if (Node[i].oldVolume) >= (Node[i].fullVolume) {
					return (Node[i].oldVolume)
				}
				return (Node[i].fullVolume)
			}()))) / dt)
			if Node[i].overflow < float64(0.0001) {
				Node[i].overflow = float64(0.0)
			}
			if (!(AllowPonding != 0)) || (Node[i].pondedArea == float64(0.0)) {
				v2 = Node[i].fullVolume
			}
		}
		Node[i].newVolume = v2
		d2 = node_getDepth(i, v2)
		Node[i].newDepth = d2
		d2 = float64(float64((float64(float64(1.0)-flowrout_OMEGA))*d1) + float64(flowrout_OMEGA*d2))
		if M.Fabs(float64(d2-d1)) <= STOPTOL {
			stopped = 1
		}
		d1 = d2
		iter++
	}
	Node[i].updated = 1
}

// flowrout.c:605
func getStorageOutflow(i int32, j int32, links []int32, dt float64) float64 {
	var (
		k       int32
		m       int32
		outflow float64
	)
	_, _, _ = k, m, outflow
	k = 0
	m = 0
	outflow = float64(0.0)
	for k = j; k < Nobjects[LINK]; k++ {
		m = links[k]
		if Link[m].node1 != i {
			break
		}
		outflow = float64(outflow + (getLinkInflow(m, dt)))
	}
	return outflow
}

// flowrout.c:629
func setNewNodeState(j int32, dt float64) {
	var (
		canPond      int32
		newNetInflow float64
	)
	_, _ = canPond, newNetInflow
	canPond = 0
	newNetInflow = 0
	if Node[j].type_ == STORAGE {
		if Node[j].updated == 0 {
			updateStorageState(j, Nobjects[LINK], (nil), dt)
		}
		return
	}
	newNetInflow = float64(float64(Node[j].inflow-Node[j].outflow) - Node[j].losses)
	Node[j].newVolume = float64(Node[j].oldVolume + float64(newNetInflow*dt))
	if Node[j].newVolume < float64(0.0001) {
		Node[j].newVolume = float64(0.0)
	}
	Node[j].overflow = float64(0.0)
	canPond = R.B2I(((AllowPonding != 0) && (Node[j].pondedArea > float64(0.0))))
	if Node[j].newVolume > Node[j].fullVolume {
		Node[j].overflow = float64((float64(Node[j].newVolume - (func() float64 {
			if (Node[j].oldVolume) >= (Node[j].fullVolume) {
				return (Node[j].oldVolume)
			}
			return (Node[j].fullVolume)
		}()))) / dt)
		if Node[j].overflow < float64(0.0001) {
			Node[j].overflow = float64(0.0)
		}
		if !(canPond != 0) {
			Node[j].newVolume = Node[j].fullVolume
		}
	}
	Node[j].newDepth = node_getDepth(j, Node[j].newVolume)
}

// flowrout.c:673
func setNewLinkState(j int32) {
	var (
		k   int32
		a_2 float64
		y1  float64
		y2  float64
	)
	_, _, _, _ = k, a_2, y1, y2
	k = 0
	a_2 = 0
	y1 = 0
	y2 = 0
	Link[j].newDepth = float64(0.0)
	Link[j].newVolume = float64(0.0)
	if Link[j].type_ == CONDUIT {
		k = Link[j].subIndex
		a_2 = float64(float64(0.5) * (float64(Conduit[k].a1 + Conduit[k].a2)))
		Link[j].newVolume = float64(float64(a_2*link_getLength(j)) * float64(Conduit[k].barrels))
		y1 = xsect_getYofA(&Link[j].xsect, Conduit[k].a1)
		y2 = xsect_getYofA(&Link[j].xsect, Conduit[k].a2)
		Link[j].newDepth = float64(float64(0.5) * (float64(y1 + y2)))
		updateNodeDepth(Link[j].node1, float64(y1+Link[j].offset1))
		updateNodeDepth(Link[j].node2, float64(y2+Link[j].offset2))
		if Conduit[k].a1 >= Link[j].xsect.aFull {
			Conduit[k].capacityLimited = 1
			Conduit[k].fullState = ALL_FULL
		} else {
			Conduit[k].capacityLimited = 0
			Conduit[k].fullState = 0
		}
	}
}

// flowrout.c:717
func updateNodeDepth(i int32, y_2 float64) {
	if Node[i].type_ == STORAGE {
		return
	}
	if ((Node[i].type_ != OUTFALL) && (Node[i].degree > 0)) && (Node[i].overflow > float64(0.0)) {
		y_2 = Node[i].fullDepth
	}
	if Node[i].newDepth < y_2 {
		Node[i].newDepth = y_2
		if (Node[i].fullDepth > float64(0.0)) && (y_2 > Node[i].fullDepth) {
			Node[i].newDepth = Node[i].fullDepth
		}
	}
}

// flowrout.c:748
func steadyflow_execute(j int32, qin *float64, qout *float64, tStep float64) int32 {
	var (
		k int32
		s float64
		q float64
	)
	_, _, _ = k, s, q
	k = 0
	s = 0
	q = 0
	if Link[j].type_ == CONDUIT {
		k = Link[j].subIndex
		q = float64((*qin) / float64(Conduit[k].barrels))
		if Link[j].xsect.type_ == DUMMY {
			Conduit[k].a1 = float64(0.0)
		} else {
			q = float64(q - (link_getLossRate(j, SF, q, tStep)))
			if q > Link[j].qFull {
				q = Link[j].qFull
				Conduit[k].a1 = Link[j].xsect.aFull
				*qin = float64(q * float64(Conduit[k].barrels))
			} else {
				s = float64(q / Conduit[k].beta)
				Conduit[k].a1 = xsect_getAofS(&Link[j].xsect, s)
			}
		}
		Conduit[k].a2 = Conduit[k].a1
		Conduit[k].q1Old = Conduit[k].q1
		Conduit[k].q2Old = Conduit[k].q2
		Conduit[k].q1 = q
		Conduit[k].q2 = q
		*qout = float64(q * float64(Conduit[k].barrels))
	} else {
		*qout = (*qin)
	}
	return 1
}
