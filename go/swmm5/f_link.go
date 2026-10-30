// Generated from EPA SWMM 5.2.4 (link.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var MIN_DELTA_Z float64 = float64(0.001)

// link.c:137
func link_readParams(j int32, type_2 int32, k int32, tok []R.CPtr, ntoks int32) int32 {
	switch type_2 {
	case CONDUIT:
		return conduit_readParams(j, k, tok, ntoks)
	case PUMP:
		return pump_readParams(j, k, tok, ntoks)
	case ORIFICE:
		return orifice_readParams(j, k, tok, ntoks)
	case WEIR:
		return weir_readParams(j, k, tok, ntoks)
	case OUTLET:
		return outlet_readParams(j, k, tok, ntoks)
	default:
		return 0
	}
	return 0
}

// link.c:162
func link_readXsectParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		i int32
		j int32
		k int32
		x [4]float64
	)
	_, _, _, _ = i, j, k, x
	i = 0
	j = 0
	k = 0
	x = ([4]float64{})
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(LINK, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	k = findmatch(tok[1], XsectTypeWords[:])
	if k < 0 {
		return error_setInpError(ERR_KEYWORD, tok[1])
	}
	if Link[j].type_ == CONDUIT {
		Conduit[Link[j].subIndex].barrels = 1
	}
	Link[j].xsect.culvertCode = 0
	if k == IRREGULAR {
		i = project_findObject(TRANSECT, tok[2])
		if i < 0 {
			return error_setInpError(ERR_NAME, tok[2])
		}
		Link[j].xsect.type_ = k
		Link[j].xsect.transect = i
		return 0
	} else if k == STREET_XSECT {
		i = project_findObject(STREET, tok[2])
		if i < 0 {
			return error_setInpError(ERR_NAME, tok[2])
		}
		Link[j].xsect.type_ = k
		Link[j].xsect.transect = i
		return 0
	} else {
		if ntoks < 6 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		if k == CUSTOM {
			if (!(getDouble(tok[2], &x[0]) != 0)) || (x[0] <= float64(0.0)) {
				return error_setInpError(ERR_NUMBER, tok[2])
			}
			i = project_findObject(CURVE, tok[3])
			if i < 0 {
				return error_setInpError(ERR_NAME, tok[3])
			}
			Link[j].xsect.type_ = k
			Link[j].xsect.transect = i
			Link[j].xsect.yFull = float64(x[0] / UCF(LENGTH))
		} else {
			for i = 2; i <= 5; i++ {
				if !(getDouble(tok[i], &x[(i-2)]) != 0) {
					return error_setInpError(ERR_NUMBER, tok[i])
				}
			}
		}
		if (Link[j].type_ != CONDUIT) && (k == RECT_OPEN) {
			x[2] = float64(0.0)
			x[3] = float64(0.0)
		}
		if !(xsect_setParams(&Link[j].xsect, k, x[:], UCF(LENGTH)) != 0) {
			return error_setInpError(ERR_NUMBER, L0)
		}
		if (Link[j].type_ == CONDUIT) && (ntoks >= 7) {
			i = R.Atoi(tok[6])
			if i <= 0 {
				return error_setInpError(ERR_NUMBER, tok[6])
			} else {
				Conduit[Link[j].subIndex].barrels = int32(int8(i))
			}
		}
		if (Link[j].type_ == CONDUIT) && (ntoks >= 8) {
			i = R.Atoi(tok[7])
			if i < 0 {
				return error_setInpError(ERR_NUMBER, tok[7])
			} else {
				Link[j].xsect.culvertCode = i
			}
		}
	}
	return 0
}

// link.c:271
func link_readLossParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		i        int32
		j        int32
		k        int32
		x        [3]float64
		seepRate float64
	)
	_, _, _, _, _ = i, j, k, x, seepRate
	i = 0
	j = 0
	k = 0
	x = ([3]float64{})
	seepRate = float64(0.0)
	if ntoks < 4 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(LINK, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	for i = 1; i <= 3; i++ {
		if (!(getDouble(tok[i], &x[(i-1)]) != 0)) || (x[(i-1)] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[i])
		}
	}
	k = 0
	if ntoks >= 5 {
		k = findmatch(tok[4], NoYesWords[:])
		if k < 0 {
			return error_setInpError(ERR_KEYWORD, tok[4])
		}
	}
	if ntoks >= 6 {
		if !(getDouble(tok[5], &seepRate) != 0) {
			return error_setInpError(ERR_NUMBER, tok[5])
		}
	}
	Link[j].cLossInlet = x[0]
	Link[j].cLossOutlet = x[1]
	Link[j].cLossAvg = x[2]
	Link[j].hasFlapGate = k
	Link[j].seepRate = float64(seepRate / UCF(RAINFALL))
	return 0
}

// link.c:315
func link_setParams(j int32, type_2 int32, n1 int32, n2 int32, k int32, x []float64) {
	Link[j].node1 = n1
	Link[j].node2 = n2
	Link[j].type_ = type_2
	Link[j].subIndex = k
	Link[j].offset1 = float64(0.0)
	Link[j].offset2 = float64(0.0)
	Link[j].q0 = float64(0.0)
	Link[j].qFull = float64(0.0)
	Link[j].setting = float64(1.0)
	Link[j].targetSetting = float64(1.0)
	Link[j].hasFlapGate = 0
	Link[j].qLimit = float64(0.0)
	Link[j].direction = 1
	switch type_2 {
	case CONDUIT:
		Conduit[k].length = float64(x[0] / UCF(LENGTH))
		Conduit[k].modLength = Conduit[k].length
		Conduit[k].roughness = x[1]
		Link[j].offset1 = float64(x[2] / UCF(LENGTH))
		Link[j].offset2 = float64(x[3] / UCF(LENGTH))
		Link[j].q0 = float64(x[4] / UCF(FLOW))
		Link[j].qLimit = float64(x[5] / UCF(FLOW))
		break
	case PUMP:
		Pump[k].pumpCurve = R.D2I(x[0])
		Link[j].hasFlapGate = 0
		Pump[k].initSetting = x[1]
		Pump[k].yOn = float64(x[2] / UCF(LENGTH))
		Pump[k].yOff = float64(x[3] / UCF(LENGTH))
		Pump[k].xMin = float64(0.0)
		Pump[k].xMax = float64(0.0)
		break
	case ORIFICE:
		Orifice[k].type_ = R.D2I(x[0])
		Link[j].offset1 = float64(x[1] / UCF(LENGTH))
		Link[j].offset2 = Link[j].offset1
		Orifice[k].cDisch = x[2]
		if x[3] > float64(0.0) {
			Link[j].hasFlapGate = 1
		} else {
			Link[j].hasFlapGate = 0
		}
		Orifice[k].orate = float64(x[4] * float64(3600.0))
		break
	case WEIR:
		Weir[k].type_ = R.D2I(x[0])
		Link[j].offset1 = float64(x[1] / UCF(LENGTH))
		Link[j].offset2 = Link[j].offset1
		Weir[k].cDisch1 = x[2]
		if x[3] > float64(0.0) {
			Link[j].hasFlapGate = 1
		} else {
			Link[j].hasFlapGate = 0
		}
		Weir[k].endCon = x[4]
		Weir[k].cDisch2 = x[5]
		Weir[k].canSurcharge = R.D2I(x[6])
		Weir[k].roadWidth = float64(x[7] / UCF(LENGTH))
		Weir[k].roadSurface = R.D2I(x[8])
		Weir[k].cdCurve = R.D2I(x[9])
		break
	case OUTLET:
		Link[j].offset1 = float64(x[0] / UCF(LENGTH))
		Link[j].offset2 = Link[j].offset1
		Outlet[k].qCoeff = x[1]
		Outlet[k].qExpon = x[2]
		Outlet[k].qCurve = R.D2I(x[3])
		if x[4] > float64(0.0) {
			Link[j].hasFlapGate = 1
		} else {
			Link[j].hasFlapGate = 0
		}
		Outlet[k].curveType = R.D2I(x[5])
		xsect_setParams(&Link[j].xsect, DUMMY, (nil), float64(0.0))
		break
	}
}

// link.c:403
func link_validate(j int32) {
	var (
		n_2 int32
	)
	_ = n_2
	n_2 = 0
	if LinkOffsets == ELEV_OFFSET {
		link_convertOffsets(j)
	}
	switch Link[j].type_ {
	case CONDUIT:
		conduit_validate(j, Link[j].subIndex)
		break
	case PUMP:
		pump_validate(j, Link[j].subIndex)
		break
	case ORIFICE:
		orifice_validate(j, Link[j].subIndex)
		break
	case WEIR:
		weir_validate(j, Link[j].subIndex)
		break
	}
	switch Link[j].type_ {
	case ORIFICE, WEIR, OUTLET:
		if float64(Node[Link[j].node1].invertElev+Link[j].offset1) < Node[Link[j].node2].invertElev {
			if RouteModel == DW {
				Link[j].offset1 = float64(Node[Link[j].node2].invertElev - Node[Link[j].node1].invertElev)
				report_writeWarningMsg(L305, Link[j].ID)
			} else {
				report_writeWarningMsg(L306, Link[j].ID)
			}
		}
	}
	if (Link[j].type_ == PUMP) || ((Link[j].type_ == ORIFICE) && (Orifice[Link[j].subIndex].type_ == BOTTOM_ORIFICE)) {
		return
	}
	n_2 = Link[j].node1
	if (Node[n_2].type_ != STORAGE) || (Node[n_2].surDepth > float64(0.0)) {
		if (Node[n_2].fullDepth) >= (float64(Link[j].offset1 + Link[j].xsect.yFull)) {
			Node[n_2].fullDepth = (Node[n_2].fullDepth)
		} else {
			Node[n_2].fullDepth = (float64(Link[j].offset1 + Link[j].xsect.yFull))
		}
	}
	n_2 = Link[j].node2
	if ((Node[n_2].type_ != STORAGE) || (Node[n_2].surDepth > float64(0.0))) && (Link[j].type_ == CONDUIT) {
		if (Node[n_2].fullDepth) >= (float64(Link[j].offset2 + Link[j].xsect.yFull)) {
			Node[n_2].fullDepth = (Node[n_2].fullDepth)
		} else {
			Node[n_2].fullDepth = (float64(Link[j].offset2 + Link[j].xsect.yFull))
		}
	}
}

// link.c:468
func link_convertOffsets(j int32) {
	var (
		elev float64
	)
	_ = elev
	elev = 0
	elev = Node[Link[j].node1].invertElev
	Link[j].offset1 = link_getOffsetHeight(j, Link[j].offset1, elev)
	if Link[j].type_ == CONDUIT {
		elev = Node[Link[j].node2].invertElev
		Link[j].offset2 = link_getOffsetHeight(j, Link[j].offset2, elev)
	} else {
		Link[j].offset2 = Link[j].offset1
	}
}

// link.c:489
func link_getOffsetHeight(j int32, offset float64, elev float64) float64 {
	if (offset <= (-(float64(10000000000.0)))) || (Link[j].type_ == PUMP) {
		return float64(0.0)
	}
	offset = float64(offset - elev)
	if offset >= float64(0.0) {
		return offset
	}
	if offset >= (-MIN_DELTA_Z) {
		return float64(0.0)
	}
	report_writeWarningMsg(L307, Link[j].ID)
	return float64(0.0)
}

// link.c:508
func link_initState(j int32) {
	var (
		p int32
	)
	_ = p
	p = 0
	Link[j].oldFlow = Link[j].q0
	Link[j].newFlow = Link[j].q0
	Link[j].oldDepth = float64(0.0)
	Link[j].newDepth = float64(0.0)
	Link[j].oldVolume = float64(0.0)
	Link[j].newVolume = float64(0.0)
	Link[j].setting = float64(1.0)
	Link[j].targetSetting = float64(1.0)
	Link[j].timeLastSet = StartDate
	Link[j].inletControl = 0
	Link[j].normalFlow = 0
	if Link[j].type_ == CONDUIT {
		conduit_initState(j, Link[j].subIndex)
	}
	if Link[j].type_ == PUMP {
		pump_initState(j, Link[j].subIndex)
	}
	for p = 0; p < Nobjects[POLLUT]; p++ {
		Link[j].oldQual[p] = float64(0.0)
		Link[j].newQual[p] = float64(0.0)
		Link[j].totalLoad[p] = float64(0.0)
	}
}

// link.c:543
func link_getInflow(j int32) float64 {
	if Link[j].setting == float64(0) {
		return float64(0.0)
	}
	switch Link[j].type_ {
	case CONDUIT:
		return conduit_getInflow(j)
	case PUMP:
		return pump_getInflow(j)
	case ORIFICE:
		return orifice_getInflow(j)
	case WEIR:
		return weir_getInflow(j)
	case OUTLET:
		return outlet_getInflow(j)
	default:
		return node_getOutflow(Link[j].node1, j)
	}
	return 0
}

// link.c:564
func link_setOldHydState(j int32) {
	var (
		k int32
	)
	_ = k
	k = 0
	Link[j].oldDepth = Link[j].newDepth
	Link[j].oldFlow = Link[j].newFlow
	Link[j].oldVolume = Link[j].newVolume
	if Link[j].type_ == CONDUIT {
		k = Link[j].subIndex
		Conduit[k].q1Old = Conduit[k].q1
		Conduit[k].q2Old = Conduit[k].q2
	}
}

// link.c:587
func link_setOldQualState(j int32) {
	var (
		p int32
	)
	_ = p
	p = 0
	for p = 0; p < Nobjects[POLLUT]; p++ {
		Link[j].oldQual[p] = Link[j].newQual[p]
		Link[j].newQual[p] = float64(0.0)
	}
}

// link.c:604
func link_setTargetSetting(j int32) {
	var (
		k  int32
		n1 int32
	)
	_, _ = k, n1
	k = 0
	n1 = 0
	if Link[j].type_ == PUMP {
		k = Link[j].subIndex
		n1 = Link[j].node1
		Link[j].targetSetting = Link[j].setting
		if ((Pump[k].yOff > float64(0.0)) && (Link[j].setting > float64(0.0))) && (Node[n1].newDepth < Pump[k].yOff) {
			Link[j].targetSetting = float64(0.0)
		}
		if ((Pump[k].yOn > float64(0.0)) && (Link[j].setting == float64(0.0))) && (Node[n1].newDepth > Pump[k].yOn) {
			Link[j].targetSetting = float64(1.0)
		}
	}
}

// link.c:628
func link_setSetting(j int32, tstep float64) {
	if Link[j].type_ == ORIFICE {
		orifice_setSetting(j, tstep)
	} else if Link[j].type_ == WEIR {
		weir_setSetting(j)
	} else {
		Link[j].setting = Link[j].targetSetting
	}
}

// link.c:643
func link_setFlapGate(j int32, n1 int32, n2 int32, q float64) int32 {
	var (
		n_2 int32
	)
	_ = n_2
	n_2 = (-1)
	if Link[j].hasFlapGate != 0 {
		if float64(q*float64(Link[j].direction)) < float64(0.0) {
			return 1
		}
	}
	if q < float64(0.0) {
		n_2 = n2
	}
	if q > float64(0.0) {
		n_2 = n1
	}
	if ((n_2 >= 0) && (Node[n_2].type_ == OUTFALL)) && (Outfall[Node[n_2].subIndex].hasFlapGate != 0) {
		return 1
	}
	return 0
}

// link.c:674
func link_getResults(j int32, f float64, x []float32) {
	var (
		p   int32
		y_2 float64
		q   float64
		u   float64
		v   float64
		c   float64
		f1  float64
	)
	_, _, _, _, _, _, _ = p, y_2, q, u, v, c, f1
	p = 0
	y_2 = 0
	q = 0
	u = 0
	v = 0
	c = 0
	f1 = float64(float64(1.0) - f)
	y_2 = float64(float64(f1*Link[j].oldDepth) + float64(f*Link[j].newDepth))
	q = float64(float64(f1*Link[j].oldFlow) + float64(f*Link[j].newFlow))
	v = float64(float64(f1*Link[j].oldVolume) + float64(f*Link[j].newVolume))
	u = link_getVelocity(j, q, y_2)
	c = float64(0.0)
	if Link[j].type_ == CONDUIT {
		if Link[j].xsect.type_ != DUMMY {
			c = float64(xsect_getAofY(&Link[j].xsect, y_2) / Link[j].xsect.aFull)
		}
	} else {
		c = Link[j].setting
	}
	if (Link[j].type_ == PUMP) && (float64(Link[j].oldFlow*Link[j].newFlow) == float64(0.0)) {
		if f >= f1 {
			q = Link[j].newFlow
		} else {
			q = Link[j].oldFlow
		}
	}
	y_2 = float64(y_2 * (UCF(LENGTH)))
	v = float64(v * (UCF(VOLUME)))
	q = float64(q * (float64(UCF(FLOW) * float64(Link[j].direction))))
	u = float64(u * (float64(UCF(LENGTH) * float64(Link[j].direction))))
	x[LINK_DEPTH] = float32(float64(y_2))
	x[LINK_FLOW] = float32(float64(q))
	x[LINK_VELOCITY] = float32(float64(u))
	x[LINK_VOLUME] = float32(float64(v))
	x[LINK_CAPACITY] = float32(float64(c))
	if !(IgnoreQuality != 0) {
		for p = 0; p < Nobjects[POLLUT]; p++ {
			c = float64(float64(f1*Link[j].oldQual[p]) + float64(f*Link[j].newQual[p]))
			x[(LINK_QUAL + p)] = float32(float64(c))
		}
	}
}

// link.c:728
func link_setOutfallDepth(j int32) {
	var (
		k     int32
		n_2   int32
		z     float64
		q     float64
		yCrit float64
		yNorm float64
	)
	_, _, _, _, _, _ = k, n_2, z, q, yCrit, yNorm
	k = 0
	n_2 = 0
	z = 0
	q = 0
	yCrit = float64(0.0)
	yNorm = float64(0.0)
	if Node[Link[j].node2].type_ == OUTFALL {
		n_2 = Link[j].node2
		z = Link[j].offset2
	} else if Node[Link[j].node1].type_ == OUTFALL {
		n_2 = Link[j].node1
		z = Link[j].offset1
	} else {
		return
	}
	if Link[j].type_ == CONDUIT {
		k = Link[j].subIndex
		q = M.Fabs(float64(Link[j].newFlow / float64(Conduit[k].barrels)))
		yNorm = link_getYnorm(j, q)
		yCrit = link_getYcrit(j, q)
	}
	node_setOutletDepth(n_2, yNorm, yCrit, z)
}

// link.c:770
func link_getYcrit(j int32, q float64) float64 {
	return xsect_getYcrit(&Link[j].xsect, q)
}

// link.c:783
func link_getYnorm(j int32, q float64) float64 {
	var (
		k   int32
		s   float64
		a_2 float64
		y_2 float64
	)
	_, _, _, _ = k, s, a_2, y_2
	k = 0
	s = 0
	a_2 = 0
	y_2 = 0
	if Link[j].type_ != CONDUIT {
		return float64(0.0)
	}
	if Link[j].xsect.type_ == DUMMY {
		return float64(0.0)
	}
	q = M.Fabs(q)
	k = Link[j].subIndex
	if q > Conduit[k].qMax {
		q = Conduit[k].qMax
	}
	if q <= float64(0.0) {
		return float64(0.0)
	}
	s = float64(q / Conduit[k].beta)
	a_2 = xsect_getAofS(&Link[j].xsect, s)
	y_2 = xsect_getYofA(&Link[j].xsect, a_2)
	return y_2
}

// link.c:808
func link_getLength(j int32) float64 {
	if Link[j].type_ == CONDUIT {
		return conduit_getLength(j)
	}
	return float64(0.0)
}

// link.c:821
func link_getVelocity(j int32, flow float64, depth float64) float64 {
	var (
		area  float64
		veloc float64
		k     int32
	)
	_, _, _ = area, veloc, k
	area = 0
	veloc = float64(0.0)
	k = 0
	if depth <= float64(0.01) {
		return float64(0.0)
	}
	if Link[j].type_ == CONDUIT {
		k = Link[j].subIndex
		flow = float64(flow / (float64(Conduit[k].barrels)))
		area = xsect_getAofY(&Link[j].xsect, depth)
		if area > float64(0.0001) {
			veloc = float64(flow / area)
		}
	}
	return veloc
}

// link.c:847
func link_getFroude(j int32, v float64, y_2 float64) float64 {
	var (
		xsect_2 *TXsect
	)
	_ = xsect_2
	xsect_2 = &Link[j].xsect
	if Link[j].type_ != CONDUIT {
		return float64(0.0)
	}
	if y_2 <= float64(0.0001) {
		return float64(0.0)
	}
	if (!(xsect_isOpen(xsect_2.type_) != 0)) && (float64(xsect_2.yFull-y_2) <= float64(0.0001)) {
		return float64(0.0)
	}
	y_2 = float64(xsect_getAofY(xsect_2, y_2) / xsect_getWofY(xsect_2, y_2))
	return float64(M.Fabs(v) / M.Sqrt(float64(float64(32.2)*y_2)))
}

// link.c:875
func link_getPower(j int32) float64 {
	var (
		n1 int32
		n2 int32
		dh float64
		q  float64
	)
	_, _, _, _ = n1, n2, dh, q
	n1 = Link[j].node1
	n2 = Link[j].node2
	dh = float64((float64(Node[n1].invertElev + Node[n1].newDepth)) - (float64(Node[n2].invertElev + Node[n2].newDepth)))
	q = M.Fabs(Link[j].newFlow)
	return float64(float64(float64(M.Fabs(dh)*q)/float64(8.814)) * float64(0.7457))
}

// link.c:893
func link_getLossRate(j int32, routeModel int32, q float64, tstep float64) float64 {
	if Link[j].type_ == CONDUIT {
		return conduit_getLossRate(j, routeModel, q, tstep)
	} else {
		return float64(0.0)
	}
	return 0
}

// link.c:911
func link_getFullState(a1 float64, a2 float64, aFull float64) int32 {
	if a1 >= aFull {
		if a2 >= aFull {
			return ALL_FULL
		} else {
			return UP_FULL
		}
	}
	if a2 >= aFull {
		return DN_FULL
	}
	return 0
}

// link.c:933
func conduit_readParams(j int32, k int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		n1 int32
		n2 int32
		x  [6]float64
		id R.CPtr
	)
	_, _, _, _ = n1, n2, x, id
	n1 = 0
	n2 = 0
	x = ([6]float64{})
	id = (R.CPtr{})
	if ntoks < 7 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(LINK, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	n1 = project_findObject(NODE, tok[1])
	if n1 < 0 {
		return error_setInpError(ERR_NAME, tok[1])
	}
	n2 = project_findObject(NODE, tok[2])
	if n2 < 0 {
		return error_setInpError(ERR_NAME, tok[2])
	}
	if !(getDouble(tok[3], &x[0]) != 0) {
		return error_setInpError(ERR_NUMBER, tok[3])
	}
	if !(getDouble(tok[4], &x[1]) != 0) {
		return error_setInpError(ERR_NUMBER, tok[4])
	}
	if (LinkOffsets == ELEV_OFFSET) && (tok[5].G(0) == 42) {
		x[2] = (-(float64(10000000000.0)))
	} else if !(getDouble(tok[5], &x[2]) != 0) {
		return error_setInpError(ERR_NUMBER, tok[5])
	}
	if (LinkOffsets == ELEV_OFFSET) && (tok[6].G(0) == 42) {
		x[3] = (-(float64(10000000000.0)))
	} else if !(getDouble(tok[6], &x[3]) != 0) {
		return error_setInpError(ERR_NUMBER, tok[6])
	}
	x[4] = float64(0.0)
	if ntoks >= 8 {
		if !(getDouble(tok[7], &x[4]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[7])
		}
	}
	x[5] = float64(0.0)
	if ntoks >= 9 {
		if !(getDouble(tok[8], &x[5]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[8])
		}
	}
	Link[j].ID = id
	link_setParams(j, CONDUIT, n1, n2, k, x[:])
	return 0
}

// link.c:992
func conduit_validate(j int32, k int32) {
	var (
		aa           float64
		lengthFactor float64
		roughness    float64
		slope        float64
	)
	_, _, _, _ = aa, lengthFactor, roughness, slope
	aa = 0
	lengthFactor = 0
	roughness = 0
	slope = 0
	if (Link[j].xsect.type_ == DUMMY) && (RouteModel == DW) {
		if Node[Link[j].node1].type_ == STORAGE {
			report_writeErrorMsg(ERR_DUMMY_LINK, Node[Link[j].node1].ID)
			return
		}
	}
	if Link[j].xsect.type_ == CUSTOM {
		xsect_setCustomXsectParams(&Link[j].xsect)
	}
	if Link[j].xsect.type_ == IRREGULAR {
		xsect_setIrregXsectParams(&Link[j].xsect)
		Conduit[k].roughness = Transect[Link[j].xsect.transect].roughness
	}
	if Link[j].xsect.type_ == STREET_XSECT {
		xsect_setStreetXsectParams(&Link[j].xsect)
		Conduit[k].roughness = Street[Link[j].xsect.transect].roughness
	}
	if Link[j].xsect.type_ == FORCE_MAIN {
		if ForceMainEqn == D_W {
			Link[j].xsect.rBot = float64(Link[j].xsect.rBot / (UCF(RAINDEPTH)))
		}
		if Link[j].xsect.rBot <= float64(0.0) {
			report_writeErrorMsg(ERR_XSECT, Link[j].ID)
		}
	}
	if Conduit[k].length <= float64(0.0) {
		report_writeErrorMsg(ERR_LENGTH, Link[j].ID)
	}
	if Conduit[k].roughness <= float64(0.0) {
		report_writeErrorMsg(ERR_ROUGHNESS, Link[j].ID)
	}
	if Conduit[k].barrels <= 0 {
		report_writeErrorMsg(ERR_BARRELS, Link[j].ID)
	}
	if Link[j].xsect.type_ != DUMMY {
		if Link[j].xsect.type_ < 0 {
			report_writeErrorMsg(ERR_NO_XSECT, Link[j].ID)
		} else if Link[j].xsect.aFull <= float64(0.0) {
			report_writeErrorMsg(ERR_XSECT, Link[j].ID)
		}
	}
	if ErrorCode != 0 {
		return
	}
	if Link[j].offset1 < float64(0.0) {
		report_writeWarningMsg(L307, Link[j].ID)
		Link[j].offset1 = float64(0.0)
	}
	if Link[j].offset2 < float64(0.0) {
		report_writeWarningMsg(L307, Link[j].ID)
		Link[j].offset2 = float64(0.0)
	}
	if Link[j].xsect.type_ == FILLED_CIRCULAR {
		Link[j].offset1 = float64(Link[j].offset1 + Link[j].xsect.yBot)
		Link[j].offset2 = float64(Link[j].offset2 + Link[j].xsect.yBot)
	}
	slope = conduit_getSlope(j)
	Conduit[k].slope = slope
	if ((RouteModel == DW) && (slope < float64(0.0))) && (Link[j].xsect.type_ != DUMMY) {
		conduit_reverse(j, k)
	}
	roughness = Conduit[k].roughness
	if (RouteModel == DW) && (Link[j].xsect.type_ == FORCE_MAIN) {
		roughness = forcemain_getEquivN(j, k)
	}
	if Link[j].xsect.type_ == IRREGULAR {
		lengthFactor = Transect[Link[j].xsect.transect].lengthFactor
		roughness = float64(roughness * (M.Sqrt(lengthFactor)))
	}
	lengthFactor = float64(1.0)
	if ((RouteModel == DW) && (LengtheningStep > float64(0.0))) && (Link[j].xsect.type_ != DUMMY) {
		lengthFactor = conduit_getLengthFactor(j, k, roughness)
	}
	if lengthFactor != float64(1.0) {
		Conduit[k].modLength = float64(lengthFactor * conduit_getLength(j))
		slope = float64(slope / lengthFactor)
		roughness = float64(roughness / M.Sqrt(lengthFactor))
	}
	if (RouteModel == DW) && (Link[j].xsect.type_ == FORCE_MAIN) {
		Link[j].xsect.sBot = forcemain_getRoughFactor(j, lengthFactor)
	}
	Conduit[k].roughFactor = float64(float64(32.2) * (float64((float64(roughness / float64(1.486))) * (float64(roughness / float64(1.486))))))
	if Link[j].xsect.type_ == DUMMY {
		Conduit[k].beta = float64(0.0)
	} else {
		Conduit[k].beta = float64(float64(float64(1.486)*M.Sqrt(M.Fabs(slope))) / roughness)
	}
	Link[j].qFull = float64(Link[j].xsect.sFull * Conduit[k].beta)
	Conduit[k].qMax = float64(Link[j].xsect.sMax * Conduit[k].beta)
	aa = float64(float64(float64(Conduit[k].beta/M.Sqrt(float64(32.2)))*M.Pow(Link[j].xsect.yFull, float64(0.1666667))) * float64(0.3))
	if aa >= float64(1.0) {
		Conduit[k].superCritical = 1
	} else {
		Conduit[k].superCritical = 0
	}
	if ((Link[j].cLossInlet == float64(0.0)) && (Link[j].cLossOutlet == float64(0.0))) && (Link[j].cLossAvg == float64(0.0)) {
		Conduit[k].hasLosses = 0
	} else {
		Conduit[k].hasLosses = 1
	}
}

// link.c:1158
func conduit_reverse(j int32, k int32) {
	var (
		i     int32
		z     float64
		cLoss float64
	)
	_, _, _ = i, z, cLoss
	i = 0
	z = 0
	cLoss = 0
	i = Link[j].node1
	Link[j].node1 = Link[j].node2
	Link[j].node2 = i
	z = Link[j].offset1
	Link[j].offset1 = Link[j].offset2
	Link[j].offset2 = z
	cLoss = Link[j].cLossInlet
	Link[j].cLossInlet = Link[j].cLossOutlet
	Link[j].cLossOutlet = cLoss
	Conduit[k].slope = (-Conduit[k].slope)
	Link[j].direction *= (-1)
	Link[j].q0 = (-Link[j].q0)
}

// link.c:1195
func conduit_getLength(j int32) float64 {
	var (
		k int32
		t int32
	)
	_, _ = k, t
	k = Link[j].subIndex
	t = 0
	if Link[j].xsect.type_ != IRREGULAR {
		return Conduit[k].length
	}
	t = Link[j].xsect.transect
	if (t < 0) || (t >= Nobjects[TRANSECT]) {
		return Conduit[k].length
	}
	return float64(Conduit[k].length / Transect[t].lengthFactor)
}

// link.c:1217
func conduit_getLengthFactor(j int32, k int32, roughness float64) float64 {
	var (
		ratio float64
		yFull float64
		vFull float64
		tStep float64
	)
	_, _, _, _ = ratio, yFull, vFull, tStep
	ratio = 0
	yFull = 0
	vFull = 0
	tStep = 0
	yFull = Link[j].xsect.yFull
	if xsect_isOpen(Link[j].xsect.type_) != 0 {
		yFull = float64(Link[j].xsect.aFull / xsect_getWofY(&Link[j].xsect, yFull))
	}
	vFull = float64(float64(float64(float64(float64(1.486)/roughness)*Link[j].xsect.sFull)*M.Sqrt(M.Fabs(Conduit[k].slope))) / Link[j].xsect.aFull)
	if LengtheningStep == float64(0.0) {
		tStep = RouteStep
	} else {
		if (RouteStep) <= (LengtheningStep) {
			tStep = (RouteStep)
		} else {
			tStep = (LengtheningStep)
		}
	}
	ratio = float64(float64((float64(M.Sqrt(float64(float64(32.2)*yFull))+vFull))*tStep) / conduit_getLength(j))
	if ratio > float64(1.0) {
		return ratio
	} else {
		return float64(1.0)
	}
	return 0
}

// link.c:1258
func conduit_getSlope(j int32) float64 {
	var (
		elev1  float64
		elev2  float64
		delta  float64
		slope  float64
		length float64
	)
	_, _, _, _, _ = elev1, elev2, delta, slope, length
	elev1 = 0
	elev2 = 0
	delta = 0
	slope = 0
	length = conduit_getLength(j)
	elev1 = float64(Link[j].offset1 + Node[Link[j].node1].invertElev)
	elev2 = float64(Link[j].offset2 + Node[Link[j].node2].invertElev)
	delta = M.Fabs(float64(elev1 - elev2))
	if delta < MIN_DELTA_Z {
		report_writeWarningMsg(L308, Link[j].ID)
		delta = MIN_DELTA_Z
	}
	if delta >= length {
		report_writeWarningMsg(L309, Link[j].ID)
		slope = float64(delta / length)
	} else {
		slope = float64(delta / M.Sqrt(float64((float64((length)*(length)))-(float64((delta)*(delta))))))
	}
	if (MinSlope > float64(0.0)) && (slope < MinSlope) {
		report_writeWarningMsg(L310, Link[j].ID)
		slope = MinSlope
		if (RouteModel == SF) || (RouteModel == KW) {
			return slope
		}
	}
	if elev1 < elev2 {
		slope = (-slope)
	}
	return slope
}

// link.c:1304
func conduit_initState(j int32, k int32) {
	Link[j].newDepth = link_getYnorm(j, float64(Link[j].q0/float64(Conduit[k].barrels)))
	Link[j].oldDepth = Link[j].newDepth
	Conduit[k].evapLossRate = float64(0.0)
	Conduit[k].seepLossRate = float64(0.0)
}

// link.c:1320
func conduit_getInflow(j int32) float64 {
	var (
		qIn float64
	)
	_ = qIn
	qIn = node_getOutflow(Link[j].node1, j)
	if Link[j].qLimit > float64(0.0) {
		if (qIn) <= (Link[j].qLimit) {
			qIn = (qIn)
		} else {
			qIn = (Link[j].qLimit)
		}
	}
	return qIn
}

// link.c:1334
func conduit_getLossRate(j int32, routeModel int32, q float64, tstep float64) float64 {
	var (
		xsect_2       *TXsect
		depth         float64
		length        float64
		width         float64
		topWidth      float64
		evapLossRate  float64
		seepLossRate  float64
		totalLossRate float64
	)
	_, _, _, _, _, _, _, _ = xsect_2, depth, length, width, topWidth, evapLossRate, seepLossRate, totalLossRate
	xsect_2 = nil
	depth = float64(float64(0.5) * (float64(Link[j].oldDepth + Link[j].newDepth)))
	length = 0
	width = 0
	topWidth = 0
	evapLossRate = float64(0.0)
	seepLossRate = float64(0.0)
	totalLossRate = float64(0.0)
	if depth > float64(0.0001) {
		xsect_2 = &Link[j].xsect
		length = conduit_getLength(j)
		if (xsect_isOpen(xsect_2.type_) != 0) && (Evap.rate > float64(0.0)) {
			topWidth = xsect_getWofY(xsect_2, depth)
			evapLossRate = float64(float64(topWidth*length) * Evap.rate)
		}
		if Link[j].seepRate > float64(0.0) {
			if xsect_2.type_ == RECT_CLOSED {
				width = xsect_2.wMax
			} else {
				if depth >= xsect_2.ywMax {
					depth = xsect_2.ywMax
				}
				width = xsect_getWofY(xsect_2, depth)
			}
			seepLossRate = float64(float64(Link[j].seepRate*width) * length)
			seepLossRate = float64(seepLossRate * Adjust.hydconFactor)
		}
		totalLossRate = float64(evapLossRate + seepLossRate)
		if routeModel == DW {
			q = float64(Link[j].newVolume / tstep)
		} else {
			if (q) < float64(0) {
				q = (-(q))
			} else {
				q = (q)
			}
		}
		if totalLossRate > q {
			evapLossRate = float64(float64(evapLossRate*q) / totalLossRate)
			seepLossRate = float64(float64(seepLossRate*q) / totalLossRate)
			totalLossRate = q
		}
	}
	Conduit[Link[j].subIndex].evapLossRate = evapLossRate
	Conduit[Link[j].subIndex].seepLossRate = seepLossRate
	return totalLossRate
}

// link.c:1406
func pump_readParams(j int32, k int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		m  int32
		n1 int32
		n2 int32
		x  [4]float64
		id R.CPtr
	)
	_, _, _, _, _ = m, n1, n2, x, id
	m = 0
	n1 = 0
	n2 = 0
	x = ([4]float64{})
	id = (R.CPtr{})
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(LINK, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	n1 = project_findObject(NODE, tok[1])
	if n1 < 0 {
		return error_setInpError(ERR_NAME, tok[1])
	}
	n2 = project_findObject(NODE, tok[2])
	if n2 < 0 {
		return error_setInpError(ERR_NAME, tok[2])
	}
	x[0] = (-(float64(1.0)))
	if ntoks >= 4 {
		if !(strcomp(tok[3], L275) != 0) {
			m = project_findObject(CURVE, tok[3])
			if m < 0 {
				return error_setInpError(ERR_NAME, tok[3])
			}
			x[0] = float64(m)
		}
	}
	x[1] = float64(1.0)
	if ntoks >= 5 {
		m = findmatch(tok[4], OffOnWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, tok[4])
		}
		x[1] = float64(m)
	}
	x[2] = float64(0.0)
	if ntoks >= 6 {
		if (!(getDouble(tok[5], &x[2]) != 0)) || (x[2] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[5])
		}
	}
	x[3] = float64(0.0)
	if ntoks >= 7 {
		if (!(getDouble(tok[6], &x[3]) != 0)) || (x[3] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[6])
		}
	}
	Link[j].ID = id
	link_setParams(j, PUMP, n1, n2, k, x[:])
	return 0
}

// link.c:1473
func pump_validate(j int32, k int32) {
	var (
		m   int32
		n1  int32
		x   float64
		y_2 float64
	)
	_, _, _, _ = m, n1, x, y_2
	m = 0
	n1 = 0
	x = 0
	y_2 = 0
	Link[j].xsect.yFull = float64(0.0)
	m = Pump[k].pumpCurve
	if m < 0 {
		Pump[k].type_ = IDEAL_PUMP
	} else {
		if (Curve[m].curveType < PUMP1_CURVE) || (Curve[m].curveType > PUMP5_CURVE) {
			report_writeErrorMsg(ERR_NO_CURVE, Link[j].ID)
		} else {
			Pump[k].type_ = (Curve[m].curveType - PUMP1_CURVE)
			if table_getFirstEntry(&Curve[m], &x, &y_2) != 0 {
				Link[j].qFull = y_2
				Pump[k].xMin = x
				Pump[k].xMax = x
				for table_getNextEntry(&Curve[m], &x, &y_2) != 0 {
					if (y_2) >= (Link[j].qFull) {
						Link[j].qFull = (y_2)
					} else {
						Link[j].qFull = (Link[j].qFull)
					}
					Pump[k].xMax = x
				}
			}
			Link[j].qFull = float64(Link[j].qFull / (UCF(FLOW)))
		}
	}
	if (Pump[k].yOn > float64(0.0)) && (Pump[k].yOn <= Pump[k].yOff) {
		report_writeErrorMsg(ERR_PUMP_LIMITS, Link[j].ID)
	}
	if Pump[k].type_ == TYPE1_PUMP {
		n1 = Link[j].node1
		if Node[n1].type_ != STORAGE {
			if (Node[n1].fullVolume) >= (float64(Pump[k].xMax / UCF(VOLUME))) {
				Node[n1].fullVolume = (Node[n1].fullVolume)
			} else {
				Node[n1].fullVolume = (float64(Pump[k].xMax / UCF(VOLUME)))
			}
		}
	}
}

// link.c:1534
func pump_initState(j int32, k int32) {
	Link[j].setting = Pump[k].initSetting
	Link[j].targetSetting = Pump[k].initSetting
}

// link.c:1548
func pump_getInflow(j int32) float64 {
	var (
		k     int32
		m     int32
		n1    int32
		n2    int32
		vol   float64
		depth float64
		head  float64
		qIn   float64
		qIn1  float64
		dh    float64
		s     float64
	)
	_, _, _, _, _, _, _, _, _, _, _ = k, m, n1, n2, vol, depth, head, qIn, qIn1, dh, s
	k = 0
	m = 0
	n1 = 0
	n2 = 0
	vol = 0
	depth = 0
	head = 0
	qIn = 0
	qIn1 = 0
	dh = float64(0.001)
	s = float64(1.0)
	k = Link[j].subIndex
	m = Pump[k].pumpCurve
	n1 = Link[j].node1
	n2 = Link[j].node2
	Link[j].flowClass = NO
	Link[j].setting = Link[j].targetSetting
	if Link[j].setting == float64(0.0) {
		return float64(0.0)
	}
	if Pump[k].type_ == IDEAL_PUMP {
		qIn = float64(Node[n1].inflow + Node[n1].overflow)
	} else {
		switch Curve[m].curveType {
		case PUMP1_CURVE:
			vol = float64(Node[n1].newVolume * UCF(VOLUME))
			qIn = float64(table_intervalLookup(&Curve[m], vol) / UCF(FLOW))
			if (vol < Pump[k].xMin) || (vol > Pump[k].xMax) {
				Link[j].flowClass = YES
			}
			break
		case PUMP2_CURVE:
			depth = float64(Node[n1].newDepth * UCF(LENGTH))
			qIn = float64(table_intervalLookup(&Curve[m], depth) / UCF(FLOW))
			if (depth < Pump[k].xMin) || (depth > Pump[k].xMax) {
				Link[j].flowClass = YES
			}
			break
		case PUMP3_CURVE, PUMP5_CURVE:
			if Curve[m].curveType == PUMP5_CURVE {
				s = Link[j].setting
			}
			head = float64(float64((float64((float64(Node[n2].newDepth+Node[n2].invertElev))-(float64(Node[n1].newDepth+Node[n1].invertElev))))/s) / s)
			head = float64((R.If[float64](((head) >= (float64(0.0))), (head), (float64(0.0)))) * UCF(LENGTH))
			qIn = float64(table_lookup(&Curve[m], head) / UCF(FLOW))
			Link[j].dqdh = float64(float64(float64((-(table_getSlope(&Curve[m], head)))*UCF(LENGTH))/UCF(FLOW)) / s)
			if (head < Pump[k].xMin) || (head > Pump[k].xMax) {
				Link[j].flowClass = YES
			}
			break
		case PUMP4_CURVE:
			depth = Node[n1].newDepth
			qIn = float64(table_lookup(&Curve[m], float64(depth*UCF(LENGTH))) / UCF(FLOW))
			qIn1 = float64(table_lookup(&Curve[m], float64((float64(depth+dh))*UCF(LENGTH))) / UCF(FLOW))
			Link[j].dqdh = float64((float64(qIn1 - qIn)) / dh)
			depth = float64(depth * (UCF(LENGTH)))
			if depth < Pump[k].xMin {
				Link[j].flowClass = DN_DRY
			}
			if depth > Pump[k].xMax {
				Link[j].flowClass = UP_DRY
			}
			break
		default:
			qIn = float64(0.0)
		}
	}
	if qIn < float64(0.0) {
		qIn = float64(0.0)
	}
	return float64(qIn * Link[j].setting)
}

// link.c:1641
func orifice_readParams(j int32, k int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		m  int32
		n1 int32
		n2 int32
		x  [5]float64
		id R.CPtr
	)
	_, _, _, _, _ = m, n1, n2, x, id
	m = 0
	n1 = 0
	n2 = 0
	x = ([5]float64{})
	id = (R.CPtr{})
	if ntoks < 6 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(LINK, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	n1 = project_findObject(NODE, tok[1])
	if n1 < 0 {
		return error_setInpError(ERR_NAME, tok[1])
	}
	n2 = project_findObject(NODE, tok[2])
	if n2 < 0 {
		return error_setInpError(ERR_NAME, tok[2])
	}
	m = findmatch(tok[3], OrificeTypeWords[:])
	if m < 0 {
		return error_setInpError(ERR_KEYWORD, tok[3])
	}
	x[0] = float64(m)
	if (LinkOffsets == ELEV_OFFSET) && (tok[4].G(0) == 42) {
		x[1] = (-(float64(10000000000.0)))
	} else if !(getDouble(tok[4], &x[1]) != 0) {
		return error_setInpError(ERR_NUMBER, tok[4])
	}
	if (!(getDouble(tok[5], &x[2]) != 0)) || (x[2] < float64(0.0)) {
		return error_setInpError(ERR_NUMBER, tok[5])
	}
	x[3] = float64(0.0)
	if ntoks >= 7 {
		m = findmatch(tok[6], NoYesWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, tok[6])
		}
		x[3] = float64(m)
	}
	x[4] = float64(0.0)
	if ntoks >= 8 {
		if (!(getDouble(tok[7], &x[4]) != 0)) || (x[4] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[7])
		}
	}
	Link[j].ID = id
	link_setParams(j, ORIFICE, n1, n2, k, x[:])
	return 0
}

// link.c:1696
func orifice_validate(j int32, k int32) {
	var (
		err int32
	)
	_ = err
	err = 0
	if (Link[j].xsect.type_ != RECT_CLOSED) && (Link[j].xsect.type_ != CIRCULAR) {
		err = ERR_REGULATOR_SHAPE
	}
	if err > 0 {
		report_writeErrorMsg(err, Link[j].ID)
		return
	}
	if Link[j].offset1 < float64(0.0) {
		Link[j].offset1 = float64(0.0)
	}
	orifice_setSetting(j, float64(0.0))
	Orifice[k].length = float64(float64(float64(2.0)*RouteStep) * M.Sqrt(float64(float64(32.2)*Link[j].xsect.yFull)))
	if (float64(200.0)) >= (Orifice[k].length) {
		Orifice[k].length = (float64(200.0))
	} else {
		Orifice[k].length = (Orifice[k].length)
	}
	Orifice[k].surfArea = float64(0.0)
}

// link.c:1729
func orifice_setSetting(j int32, tstep float64) {
	var (
		k     int32
		delta float64
		step  float64
		h     float64
		f     float64
	)
	_, _, _, _, _ = k, delta, step, h, f
	k = Link[j].subIndex
	delta = 0
	step = 0
	h = 0
	f = 0
	if (Orifice[k].orate == float64(0.0)) || (tstep == float64(0.0)) {
		Link[j].setting = Link[j].targetSetting
	} else {
		delta = float64(Link[j].targetSetting - Link[j].setting)
		step = float64(tstep / Orifice[k].orate)
		if float64(step+float64(0.001)) >= M.Fabs(delta) {
			Link[j].setting = Link[j].targetSetting
		} else {
			Link[j].setting = float64(Link[j].setting + (float64(float64((R.If[int32](((delta) < float64(0)), (-1), (1)))) * step)))
		}
	}
	h = float64(Link[j].setting * Link[j].xsect.yFull)
	f = float64(xsect_getAofY(&Link[j].xsect, h) * M.Sqrt(float64(float64(2.0)*float64(32.2))))
	Orifice[k].cOrif = float64(Orifice[k].cDisch * f)
	Orifice[k].cWeir = float64(orifice_getWeirCoeff(j, k, h) * f)
}

// link.c:1766
func orifice_getWeirCoeff(j int32, k int32, h float64) float64 {
	var (
		w      float64
		aOverL float64
	)
	_, _ = w, aOverL
	w = 0
	aOverL = 0
	if Orifice[k].type_ == BOTTOM_ORIFICE {
		if Link[j].xsect.type_ == CIRCULAR {
			aOverL = float64(h / float64(4.0))
		} else {
			w = Link[j].xsect.wMax
			aOverL = float64((float64(h * w)) / (float64(float64(2.0) * (float64(h + w)))))
		}
		h = float64(float64(Orifice[k].cDisch/float64(0.414)) * aOverL)
		Orifice[k].hCrit = h
	} else {
		Orifice[k].hCrit = h
		h = float64(h / float64(2.0))
	}
	return float64(Orifice[k].cDisch * M.Sqrt(h))
}

// link.c:1812
func orifice_getInflow(j int32) float64 {
	var (
		k      int32
		n1     int32
		n2     int32
		head   float64
		h1     float64
		h2     float64
		y1     float64
		dir    float64
		f      float64
		hcrest float64
		hcrown float64
		hmidpt float64
		q      float64
		ratio  float64
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = k, n1, n2, head, h1, h2, y1, dir, f, hcrest, hcrown, hmidpt
	_, _ = q, ratio
	k = 0
	n1 = 0
	n2 = 0
	head = 0
	h1 = 0
	h2 = 0
	y1 = 0
	dir = 0
	f = 0
	hcrest = float64(0.0)
	hcrown = float64(0.0)
	hmidpt = 0
	q = 0
	ratio = 0
	n1 = Link[j].node1
	n2 = Link[j].node2
	k = Link[j].subIndex
	if RouteModel == DW {
		h1 = float64(Node[n1].newDepth + Node[n1].invertElev)
		h2 = float64(Node[n2].newDepth + Node[n2].invertElev)
	} else {
		h1 = float64(Node[n1].newDepth + Node[n1].invertElev)
		h2 = Node[n1].invertElev
	}
	if h1 >= h2 {
		dir = float64(1.0)
	} else {
		dir = (-(float64(1.0)))
	}
	y1 = Node[n1].newDepth
	if dir < float64(0.0) {
		head = h1
		h1 = h2
		h2 = head
		y1 = Node[n2].newDepth
	}
	if Orifice[k].type_ == BOTTOM_ORIFICE {
		hcrest = float64(Node[n1].invertElev + Link[j].offset1)
		if h1 < hcrest {
			head = float64(0.0)
		} else if h2 > hcrest {
			head = float64(h1 - h2)
		} else {
			head = float64(h1 - hcrest)
		}
		f = float64(head / Orifice[k].hCrit)
		if (f) <= (float64(1.0)) {
			f = (f)
		} else {
			f = (float64(1.0))
		}
	} else {
		hcrest = float64(Node[n1].invertElev + Link[j].offset1)
		hcrown = float64(hcrest + float64(Link[j].xsect.yFull*Link[j].setting))
		hmidpt = float64((float64(hcrest + hcrown)) / float64(2.0))
		if (h1 < hcrown) && (hcrown > hcrest) {
			f = float64((float64(h1 - hcrest)) / (float64(hcrown - hcrest)))
		} else {
			f = float64(1.0)
		}
		if f < float64(1.0) {
			head = float64(h1 - hcrest)
		} else if h2 < hmidpt {
			head = float64(h1 - hmidpt)
		} else {
			head = float64(h1 - h2)
		}
	}
	if ((head <= float64(0.0001)) || (y1 <= float64(0.0001))) || (link_setFlapGate(j, n1, n2, dir) != 0) {
		Link[j].newDepth = float64(0.0)
		Link[j].flowClass = DRY
		Orifice[k].surfArea = float64(float64(0.0001) * Orifice[k].length)
		Link[j].dqdh = float64(0.0)
		return float64(0.0)
	}
	Link[j].flowClass = SUBCRITICAL
	if hcrest > h2 {
		if dir == float64(1.0) {
			Link[j].flowClass = DN_CRITICAL
		} else {
			Link[j].flowClass = UP_CRITICAL
		}
	}
	y1 = float64(Link[j].xsect.yFull * Link[j].setting)
	if Orifice[k].type_ == SIDE_ORIFICE {
		Link[j].newDepth = float64(y1 * f)
		Orifice[k].surfArea = float64(xsect_getWofY(&Link[j].xsect, Link[j].newDepth) * Orifice[k].length)
	} else {
		Link[j].newDepth = y1
		Orifice[k].surfArea = xsect_getAofY(&Link[j].xsect, y1)
	}
	q = float64(dir * orifice_getFlow(j, k, head, f, Link[j].hasFlapGate))
	if (f < float64(1.0)) && (h2 > hcrest) {
		ratio = float64((float64(h2 - hcrest)) / (float64(h1 - hcrest)))
		q = float64(q * (M.Pow((float64(float64(1.0) - M.Pow(ratio, float64(1.5)))), float64(0.385))))
	}
	return q
}

// link.c:1938
func orifice_getFlow(j int32, k int32, head float64, f float64, hasFlapGate int32) float64 {
	var (
		area  float64
		q     float64
		veloc float64
		hLoss float64
	)
	_, _, _, _ = area, q, veloc, hLoss
	area = 0
	q = 0
	veloc = 0
	hLoss = 0
	if (head == float64(0.0)) || (f <= float64(0.0)) {
		Link[j].dqdh = float64(0.0)
		return float64(0.0)
	} else if f < float64(1.0) {
		q = float64(Orifice[k].cWeir * M.Pow(f, float64(1.5)))
		Link[j].dqdh = float64(float64(float64(1.5)*q) / (float64(f * Orifice[k].hCrit)))
	} else {
		q = float64(Orifice[k].cOrif * M.Sqrt(head))
		Link[j].dqdh = float64(q / (float64(float64(2.0) * head)))
	}
	if hasFlapGate != 0 {
		area = xsect_getAofY(&Link[j].xsect, float64(Link[j].setting*Link[j].xsect.yFull))
		veloc = float64(q / area)
		hLoss = float64(float64(float64((float64(float64(4.0)/float64(32.2)))*veloc)*veloc) * M.Exp(float64(float64((-(float64(1.15)))*veloc)/M.Sqrt(head))))
		if f < float64(1.0) {
			f = float64(f - float64(hLoss/Orifice[k].hCrit))
			if f < float64(0.0) {
				f = float64(0.0)
			}
		} else {
			head = float64(head - hLoss)
			if head < float64(0.0) {
				head = float64(0.0)
			}
		}
		q = orifice_getFlow(j, k, head, f, 0)
	}
	return q
}

// link.c:2010
func weir_readParams(j int32, k int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		m  int32
		n1 int32
		n2 int32
		x  [10]float64
		id R.CPtr
	)
	_, _, _, _, _ = m, n1, n2, x, id
	m = 0
	n1 = 0
	n2 = 0
	x = ([10]float64{})
	id = (R.CPtr{})
	if ntoks < 6 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(LINK, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	n1 = project_findObject(NODE, tok[1])
	if n1 < 0 {
		return error_setInpError(ERR_NAME, tok[1])
	}
	n2 = project_findObject(NODE, tok[2])
	if n2 < 0 {
		return error_setInpError(ERR_NAME, tok[2])
	}
	m = findmatch(tok[3], WeirTypeWords[:])
	if m < 0 {
		return error_setInpError(ERR_KEYWORD, tok[3])
	}
	x[0] = float64(m)
	if (LinkOffsets == ELEV_OFFSET) && (tok[4].G(0) == 42) {
		x[1] = (-(float64(10000000000.0)))
	} else if !(getDouble(tok[4], &x[1]) != 0) {
		return error_setInpError(ERR_NUMBER, tok[4])
	}
	if (!(getDouble(tok[5], &x[2]) != 0)) || (x[2] < float64(0.0)) {
		return error_setInpError(ERR_NUMBER, tok[5])
	}
	x[3] = float64(0.0)
	x[4] = float64(0.0)
	x[5] = float64(0.0)
	x[6] = float64(1.0)
	x[7] = float64(0.0)
	x[8] = float64(0.0)
	x[9] = (-(float64(1.0)))
	if (ntoks >= 7) && (tok[6].G(0) != 42) {
		m = findmatch(tok[6], NoYesWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, tok[6])
		}
		x[3] = float64(m)
	}
	if (ntoks >= 8) && (tok[7].G(0) != 42) {
		if (!(getDouble(tok[7], &x[4]) != 0)) || (x[4] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[7])
		}
	}
	if (ntoks >= 9) && (tok[8].G(0) != 42) {
		if (!(getDouble(tok[8], &x[5]) != 0)) || (x[5] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[8])
		}
	}
	if (ntoks >= 10) && (tok[9].G(0) != 42) {
		m = findmatch(tok[9], NoYesWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, tok[9])
		}
		x[6] = float64(m)
	}
	if (R.Set(&m, R.D2I(x[0]))) == ROADWAY_WEIR {
		if ntoks >= 11 {
			if (!(getDouble(tok[10], &x[7]) != 0)) || (x[7] < float64(0.0)) {
				return error_setInpError(ERR_NUMBER, tok[10])
			}
		}
		if ntoks >= 12 {
			if strcomp(tok[11], L311) != 0 {
				x[8] = float64(1.0)
			} else if strcomp(tok[11], L312) != 0 {
				x[8] = float64(2.0)
			}
		}
	}
	if (ntoks >= 13) && (tok[12].G(0) != 42) {
		m = project_findObject(CURVE, tok[12])
		if m < 0 {
			return error_setInpError(ERR_NAME, tok[12])
		}
		x[9] = float64(m)
	}
	Link[j].ID = id
	link_setParams(j, WEIR, n1, n2, k, x[:])
	return 0
}

// link.c:2103
func weir_validate(j int32, k int32) {
	var (
		err  int32
		q    float64
		q1   float64
		q2   float64
		head float64
	)
	_, _, _, _, _ = err, q, q1, q2, head
	err = 0
	q = 0
	q1 = 0
	q2 = 0
	head = 0
	switch Weir[k].type_ {
	case TRANSVERSE_WEIR, SIDEFLOW_WEIR, ROADWAY_WEIR:
		if Link[j].xsect.type_ != RECT_OPEN {
			err = ERR_REGULATOR_SHAPE
		}
		Weir[k].slope = float64(0.0)
		break
	case VNOTCH_WEIR:
		if Link[j].xsect.type_ != TRIANGULAR {
			err = ERR_REGULATOR_SHAPE
		} else {
			Weir[k].slope = Link[j].xsect.sBot
		}
		break
	case TRAPEZOIDAL_WEIR:
		if Link[j].xsect.type_ != TRAPEZOIDAL {
			err = ERR_REGULATOR_SHAPE
		} else {
			Weir[k].slope = Link[j].xsect.sBot
		}
		break
	}
	if err > 0 {
		report_writeErrorMsg(err, Link[j].ID)
		return
	}
	if Link[j].offset1 < float64(0.0) {
		Link[j].offset1 = float64(0.0)
	}
	Weir[k].length = float64(float64(float64(2.0)*RouteStep) * M.Sqrt(float64(float64(32.2)*Link[j].xsect.yFull)))
	if (float64(200.0)) >= (Weir[k].length) {
		Weir[k].length = (float64(200.0))
	} else {
		Weir[k].length = (Weir[k].length)
	}
	Weir[k].surfArea = float64(0.0)
	head = Link[j].xsect.yFull
	weir_getFlow(j, k, head, float64(1.0), 0, &q1, &q2)
	q = float64(q1 + q2)
	head = float64(head / float64(2.0))
	Weir[k].cSurcharge = float64(q / M.Sqrt(head))
}

// link.c:2166
func weir_setSetting(j int32) {
	var (
		k  int32
		h  float64
		q  float64
		q1 float64
		q2 float64
	)
	_, _, _, _, _ = k, h, q, q1, q2
	k = Link[j].subIndex
	h = 0
	q = 0
	q1 = 0
	q2 = 0
	Link[j].setting = Link[j].targetSetting
	if !(Weir[k].canSurcharge != 0) {
		return
	}
	if Weir[k].type_ == ROADWAY_WEIR {
		return
	}
	if Link[j].setting == float64(0.0) {
		Weir[k].cSurcharge = float64(0.0)
	} else {
		h = float64(Link[j].setting * Link[j].xsect.yFull)
		weir_getFlow(j, k, h, float64(1.0), 0, &q1, &q2)
		q = float64(q1 + q2)
		h = float64(h / float64(2.0))
		Weir[k].cSurcharge = float64(q / M.Sqrt(h))
	}
}

// link.c:2198
func weir_getInflow(j int32) float64 {
	var (
		n1        int32
		n2        int32
		k         int32
		q1        float64
		q2        float64
		head      float64
		h1        float64
		h2        float64
		hcrest    float64
		hcrown    float64
		y_2       float64
		dir       float64
		ratio     float64
		weirPower [4]float64
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = n1, n2, k, q1, q2, head, h1, h2, hcrest, hcrown, y_2, dir
	_, _ = ratio, weirPower
	n1 = 0
	n2 = 0
	k = 0
	q1 = 0
	q2 = 0
	head = 0
	h1 = 0
	h2 = 0
	hcrest = 0
	hcrown = 0
	y_2 = 0
	dir = 0
	ratio = 0
	weirPower = [4]float64{float64(1.5), float64(float64(5.0) / float64(3.0)), float64(2.5), float64(1.5)}
	n1 = Link[j].node1
	n2 = Link[j].node2
	k = Link[j].subIndex
	if RouteModel == DW {
		h1 = float64(Node[n1].newDepth + Node[n1].invertElev)
		h2 = float64(Node[n2].newDepth + Node[n2].invertElev)
	} else {
		h1 = float64(Node[n1].newDepth + Node[n1].invertElev)
		h2 = Node[n1].invertElev
	}
	if h1 > h2 {
		dir = float64(1.0)
	} else {
		dir = (-(float64(1.0)))
	}
	if dir < float64(0.0) {
		head = h1
		h1 = h2
		h2 = head
	}
	hcrest = float64(Node[n1].invertElev + Link[j].offset1)
	hcrown = float64(hcrest + Link[j].xsect.yFull)
	if Weir[k].type_ == ROADWAY_WEIR {
		return roadway_getInflow(j, dir, hcrest, h1, h2)
	}
	hcrest = float64(hcrest + (float64((float64(float64(1.0) - Link[j].setting)) * Link[j].xsect.yFull)))
	head = float64(h1 - hcrest)
	Link[j].dqdh = float64(0.0)
	if ((head <= float64(0.0001)) || (hcrest >= hcrown)) || (link_setFlapGate(j, n1, n2, dir) != 0) {
		Link[j].newDepth = float64(0.0)
		Link[j].flowClass = DRY
		return float64(0.0)
	}
	Link[j].flowClass = SUBCRITICAL
	if hcrest > h2 {
		if dir == float64(1.0) {
			Link[j].flowClass = DN_CRITICAL
		} else {
			Link[j].flowClass = UP_CRITICAL
		}
	}
	y_2 = float64(Link[j].xsect.yFull - (float64(hcrown - (R.If[float64](((h1) <= (hcrown)), (h1), (hcrown))))))
	Weir[k].surfArea = float64(xsect_getWofY(&Link[j].xsect, y_2) * Weir[k].length)
	if h1 >= hcrown {
		if Weir[k].canSurcharge != 0 {
			y_2 = float64((float64(hcrest + hcrown)) / float64(2.0))
			if h2 < y_2 {
				head = float64(h1 - y_2)
			} else {
				head = float64(h1 - h2)
			}
			y_2 = float64(hcrown - hcrest)
			q1 = weir_getOrificeFlow(j, head, y_2, Weir[k].cSurcharge)
			Link[j].newDepth = y_2
			return float64(dir * q1)
		} else {
			head = float64(hcrown - hcrest)
		}
	}
	weir_getFlow(j, k, head, dir, Link[j].hasFlapGate, &q1, &q2)
	if h2 > hcrest {
		ratio = float64((float64(h2 - hcrest)) / (float64(h1 - hcrest)))
		q1 = float64(q1 * (M.Pow((float64(float64(1.0) - M.Pow(ratio, weirPower[Weir[k].type_]))), float64(0.385))))
		if q2 > float64(0.0) {
			q2 = float64(q2 * (M.Pow((float64(float64(1.0) - M.Pow(ratio, weirPower[VNOTCH_WEIR]))), float64(0.385))))
		}
	}
	if (float64(h1 - hcrest)) <= (Link[j].xsect.yFull) {
		Link[j].newDepth = (float64(h1 - hcrest))
	} else {
		Link[j].newDepth = (Link[j].xsect.yFull)
	}
	return float64(dir * (float64(q1 + q2)))
}

// link.c:2321
func weir_getFlow(j int32, k int32, head float64, dir float64, hasFlapGate int32, q1 *float64, q2 *float64) {
	var (
		length  float64
		h       float64
		y_2     float64
		hLoss   float64
		area    float64
		veloc   float64
		wType   int32
		cdCurve int32
		cDisch1 float64
	)
	_, _, _, _, _, _, _, _, _ = length, h, y_2, hLoss, area, veloc, wType, cdCurve, cDisch1
	length = 0
	h = 0
	y_2 = 0
	hLoss = 0
	area = 0
	veloc = 0
	wType = 0
	cdCurve = Weir[k].cdCurve
	cDisch1 = Weir[k].cDisch1
	*q1 = float64(0.0)
	*q2 = float64(0.0)
	Link[j].dqdh = float64(0.0)
	if head <= float64(0.0) {
		return
	}
	length = float64(Link[j].xsect.wMax * UCF(LENGTH))
	h = float64(head * UCF(LENGTH))
	if cdCurve >= 0 {
		cDisch1 = table_lookup(&Curve[cdCurve], h)
	}
	wType = Weir[k].type_
	if (wType == VNOTCH_WEIR) && (Link[j].setting < float64(1.0)) {
		wType = TRAPEZOIDAL_WEIR
	}
	switch wType {
	case TRANSVERSE_WEIR:
		length = float64(length - (float64(float64(float64(0.1)*Weir[k].endCon) * h)))
		if (length) >= (float64(0.0)) {
			length = (length)
		} else {
			length = (float64(0.0))
		}
		*q1 = float64(float64(cDisch1*length) * M.Pow(h, float64(1.5)))
		break
	case SIDEFLOW_WEIR:
		length = float64(length - (float64(float64(float64(0.1)*Weir[k].endCon) * h)))
		if (length) >= (float64(0.0)) {
			length = (length)
		} else {
			length = (float64(0.0))
		}
		if dir < float64(0.0) {
			*q1 = float64(float64(cDisch1*length) * M.Pow(h, float64(1.5)))
		} else {
			*q1 = float64(float64(cDisch1*M.Pow(length, float64(0.83))) * M.Pow(h, float64(1.67)))
		}
		break
	case VNOTCH_WEIR:
		*q1 = float64(float64(cDisch1*Weir[k].slope) * M.Pow(h, float64(2.5)))
		break
	case TRAPEZOIDAL_WEIR:
		y_2 = float64((float64(float64(1.0) - Link[j].setting)) * Link[j].xsect.yFull)
		length = float64(xsect_getWofY(&Link[j].xsect, y_2) * UCF(LENGTH))
		*q1 = float64(float64(cDisch1*length) * M.Pow(h, float64(1.5)))
		*q2 = float64(float64(Weir[k].cDisch2*Weir[k].slope) * M.Pow(h, float64(2.5)))
	}
	if UnitSystem == SI {
		*q1 = float64((*q1) / (float64(0.028317)))
		*q2 = float64((*q2) / (float64(0.028317)))
	}
	if hasFlapGate != 0 {
		area = weir_getOpenArea(j, head)
		if area > float64(1e-06) {
			veloc = float64((float64((*q1) + (*q2))) / area)
			hLoss = float64(float64(float64((float64(float64(4.0)/float64(32.2)))*veloc)*veloc) * M.Exp(float64(float64((-(float64(1.15)))*veloc)/M.Sqrt(head))))
			head = float64(head - hLoss)
			if head < float64(0.0) {
				head = float64(0.0)
			}
			weir_getFlow(j, k, head, dir, 0, q1, q2)
		}
	}
	Link[j].dqdh = weir_getdqdh(k, dir, head, (*q1), (*q2))
}

// link.c:2432
func weir_getOrificeFlow(j int32, head float64, y_2 float64, cOrif float64) float64 {
	var (
		a_2   float64
		q     float64
		v     float64
		hloss float64
	)
	_, _, _, _ = a_2, q, v, hloss
	a_2 = 0
	q = 0
	v = 0
	hloss = 0
	q = float64(cOrif * M.Sqrt(head))
	if Link[j].hasFlapGate != 0 {
		a_2 = weir_getOpenArea(j, y_2)
		if a_2 > float64(0.0) {
			v = float64(q / a_2)
			hloss = float64(float64(float64((float64(float64(4.0)/float64(32.2)))*v)*v) * M.Exp(float64(float64((-(float64(1.15)))*v)/M.Sqrt(y_2))))
			head = float64(head - hloss)
			if (head) >= (float64(0.0)) {
				head = (head)
			} else {
				head = (float64(0.0))
			}
			q = float64(cOrif * M.Sqrt(head))
		}
	}
	if head > float64(0.0) {
		Link[j].dqdh = float64(q / (float64(float64(2.0) * head)))
	} else {
		Link[j].dqdh = float64(0.0)
	}
	return q
}

// link.c:2467
func weir_getOpenArea(j int32, y_2 float64) float64 {
	var (
		z  float64
		zy float64
	)
	_, _ = z, zy
	z = 0
	zy = 0
	z = float64((float64(float64(1.0) - Link[j].setting)) * Link[j].xsect.yFull)
	zy = float64(z + y_2)
	if (zy) <= (Link[j].xsect.yFull) {
		zy = (zy)
	} else {
		zy = (Link[j].xsect.yFull)
	}
	return float64(xsect_getAofY(&Link[j].xsect, zy) - xsect_getAofY(&Link[j].xsect, z))
}

// link.c:2492
func weir_getdqdh(k int32, dir float64, h float64, q1 float64, q2 float64) float64 {
	var (
		q1h float64
		q2h float64
	)
	_, _ = q1h, q2h
	q1h = 0
	q2h = 0
	if M.Fabs(h) < float64(0.0001) {
		return float64(0.0)
	}
	q1h = M.Fabs(float64(q1 / h))
	q2h = M.Fabs(float64(q2 / h))
	switch Weir[k].type_ {
	case TRANSVERSE_WEIR:
		return float64(float64(1.5) * q1h)
	case SIDEFLOW_WEIR:
		if dir < float64(0.0) {
			return float64(float64(1.5) * q1h)
		} else {
			return float64(float64(1.67) * q1h)
		}
	case VNOTCH_WEIR:
		if q2h == float64(0.0) {
			return float64(float64(2.5) * q1h)
		} else {
			return float64(float64(float64(1.5)*q1h) + float64(float64(2.5)*q2h))
		}
	case TRAPEZOIDAL_WEIR:
		return float64(float64(float64(1.5)*q1h) + float64(float64(2.5)*q2h))
	}
	return float64(0.0)
}

// link.c:2524
func outlet_readParams(j int32, k int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		i   int32
		m   int32
		n_2 int32
		n1  int32
		n2  int32
		x   [6]float64
		id  R.CPtr
		s   R.CPtr
	)
	_, _, _, _, _, _, _, _ = i, m, n_2, n1, n2, x, id, s
	i = 0
	m = 0
	n_2 = 0
	n1 = 0
	n2 = 0
	x = ([6]float64{})
	id = (R.CPtr{})
	s = (R.CPtr{})
	if ntoks < 6 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(LINK, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	n1 = project_findObject(NODE, tok[1])
	if n1 < 0 {
		return error_setInpError(ERR_NAME, tok[1])
	}
	n2 = project_findObject(NODE, tok[2])
	if n2 < 0 {
		return error_setInpError(ERR_NAME, tok[2])
	}
	if (LinkOffsets == ELEV_OFFSET) && (tok[3].G(0) == 42) {
		x[0] = (-(float64(10000000000.0)))
	} else {
		if !(getDouble(tok[3], &x[0]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[3])
		}
		if (LinkOffsets == DEPTH_OFFSET) && (x[0] < float64(0.0)) {
			x[0] = float64(0.0)
		}
	}
	m = findmatch(tok[4], RelationWords[:])
	if m < 0 {
		return error_setInpError(ERR_KEYWORD, tok[4])
	}
	x[1] = float64(0.0)
	x[2] = float64(0.0)
	x[3] = (-(float64(1.0)))
	x[4] = float64(0.0)
	x[5] = float64(NODE_DEPTH)
	s = R.Strtok(tok[4], L313)
	s = R.Strtok((R.CPtr{}), L313)
	if strcomp(s, L314) != 0 {
		x[5] = float64(NODE_HEAD)
	}
	if m == FUNCTIONAL {
		if ntoks < 7 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		if !(getDouble(tok[5], &x[1]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[5])
		}
		if !(getDouble(tok[6], &x[2]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[6])
		}
		n_2 = 7
	} else {
		i = project_findObject(CURVE, tok[5])
		if i < 0 {
			return error_setInpError(ERR_NAME, tok[5])
		}
		x[3] = float64(i)
		n_2 = 6
	}
	if ntoks > n_2 {
		i = findmatch(tok[n_2], NoYesWords[:])
		if i < 0 {
			return error_setInpError(ERR_KEYWORD, tok[n_2])
		}
		x[4] = float64(i)
	}
	Link[j].ID = id
	link_setParams(j, OUTLET, n1, n2, k, x[:])
	return 0
}

// link.c:2608
func outlet_getInflow(j int32) float64 {
	var (
		k      int32
		n1     int32
		n2     int32
		head   float64
		hcrest float64
		h1     float64
		h2     float64
		y1     float64
		dir    float64
	)
	_, _, _, _, _, _, _, _, _ = k, n1, n2, head, hcrest, h1, h2, y1, dir
	k = 0
	n1 = 0
	n2 = 0
	head = 0
	hcrest = 0
	h1 = 0
	h2 = 0
	y1 = 0
	dir = 0
	n1 = Link[j].node1
	n2 = Link[j].node2
	k = Link[j].subIndex
	if RouteModel == DW {
		h1 = float64(Node[n1].newDepth + Node[n1].invertElev)
		h2 = float64(Node[n2].newDepth + Node[n2].invertElev)
	} else {
		h1 = float64(Node[n1].newDepth + Node[n1].invertElev)
		h2 = Node[n1].invertElev
	}
	if h1 >= h2 {
		dir = float64(1.0)
	} else {
		dir = (-(float64(1.0)))
	}
	y1 = Node[n1].newDepth
	if dir < float64(0.0) {
		y1 = h1
		h1 = h2
		h2 = y1
		y1 = Node[n2].newDepth
	}
	hcrest = float64(Node[n1].invertElev + Link[j].offset1)
	if (Outlet[k].curveType == NODE_HEAD) && (RouteModel == DW) {
		head = float64(h1 - (R.If[float64](((h2) >= (hcrest)), (h2), (hcrest))))
	} else {
		head = float64(h1 - hcrest)
	}
	if ((head <= float64(0.0001)) || (y1 <= float64(0.0001))) || (link_setFlapGate(j, n1, n2, dir) != 0) {
		Link[j].newDepth = float64(0.0)
		Link[j].flowClass = DRY
		return float64(0.0)
	}
	Link[j].newDepth = head
	Link[j].flowClass = SUBCRITICAL
	return float64(float64(dir*Link[j].setting) * outlet_getFlow(k, head))
}

// link.c:2672
func outlet_getFlow(k int32, head float64) float64 {
	var (
		m int32
		h float64
	)
	_, _ = m, h
	m = 0
	h = 0
	h = float64(head * UCF(LENGTH))
	m = Outlet[k].qCurve
	if m >= 0 {
		return float64(table_lookup(&Curve[m], h) / UCF(FLOW))
	} else {
		return float64(float64(Outlet[k].qCoeff*M.Pow(h, Outlet[k].qExpon)) / UCF(FLOW))
	}
	return 0
}
