// Generated from EPA SWMM 5.2.4 (dwflow.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var MAXVELOCITY float64 = float64(50.0)

// dwflow.c:57
func dwflow_findConduitFlow(j int32, steps int32, omega float64, dt float64) {
	var (
		k        int32
		n1       int32
		n2       int32
		z1       float64
		z2       float64
		h1       float64
		h2       float64
		y1       float64
		y2       float64
		a1       float64
		a2       float64
		r1       float64
		yMid     float64
		rMid     float64
		aMid     float64
		aWtd     float64
		rWtd     float64
		qLast    float64
		qOld     float64
		aOld     float64
		v        float64
		rho      float64
		sigma    float64
		length   float64
		wSlot    float64
		dq1      float64
		dq2      float64
		dq3      float64
		dq4      float64
		dq5      float64
		dq6      float64
		denom    float64
		q        float64
		barrels  float64
		xsect_2  *TXsect
		isFull   int32
		isClosed int32
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = k, n1, n2, z1, z2, h1, h2, y1, y2, a1, a2, r1
	_, _, _, _, _, _, _, _, _, _, _, _ = yMid, rMid, aMid, aWtd, rWtd, qLast, qOld, aOld, v, rho, sigma, length
	_, _, _, _, _, _, _, _, _, _, _, _ = wSlot, dq1, dq2, dq3, dq4, dq5, dq6, denom, q, barrels, xsect_2, isFull
	_ = isClosed
	k = 0
	n1 = 0
	n2 = 0
	z1 = 0
	z2 = 0
	h1 = 0
	h2 = 0
	y1 = 0
	y2 = 0
	a1 = 0
	a2 = 0
	r1 = 0
	yMid = 0
	rMid = 0
	aMid = 0
	aWtd = 0
	rWtd = 0
	qLast = 0
	qOld = 0
	aOld = 0
	v = 0
	rho = 0
	sigma = 0
	length = 0
	wSlot = 0
	dq1 = 0
	dq2 = 0
	dq3 = 0
	dq4 = 0
	dq5 = 0
	dq6 = 0
	denom = 0
	q = 0
	barrels = 0
	xsect_2 = &Link[j].xsect
	isFull = 0
	isClosed = 0
	if Link[j].setting == float64(0) {
		isClosed = 1
	}
	k = Link[j].subIndex
	barrels = float64(Conduit[k].barrels)
	qOld = float64(Link[j].oldFlow / barrels)
	qLast = Conduit[k].q1
	Conduit[k].evapLossRate = float64(0.0)
	Conduit[k].seepLossRate = float64(0.0)
	n1 = Link[j].node1
	n2 = Link[j].node2
	z1 = float64(Node[n1].invertElev + Link[j].offset1)
	z2 = float64(Node[n2].invertElev + Link[j].offset2)
	h1 = float64(Node[n1].newDepth + Node[n1].invertElev)
	h2 = float64(Node[n2].newDepth + Node[n2].invertElev)
	if (h1) >= (z1) {
		h1 = (h1)
	} else {
		h1 = (z1)
	}
	if (h2) >= (z2) {
		h2 = (h2)
	} else {
		h2 = (z2)
	}
	y1 = float64(h1 - z1)
	y2 = float64(h2 - z2)
	if (y1) >= (float64(0.0001)) {
		y1 = (y1)
	} else {
		y1 = (float64(0.0001))
	}
	if (y2) >= (float64(0.0001)) {
		y2 = (y2)
	} else {
		y2 = (float64(0.0001))
	}
	if SurchargeMethod != SLOT {
		if (y1) <= (xsect_2.yFull) {
			y1 = (y1)
		} else {
			y1 = (xsect_2.yFull)
		}
		if (y2) <= (xsect_2.yFull) {
			y2 = (y2)
		} else {
			y2 = (xsect_2.yFull)
		}
	}
	aOld = Conduit[k].a2
	if (aOld) >= (float64(0.0001)) {
		aOld = (aOld)
	} else {
		aOld = (float64(0.0001))
	}
	length = Conduit[k].modLength
	findSurfArea(j, qLast, length, &h1, &h2, &y1, &y2)
	wSlot = getSlotWidth(xsect_2, y1)
	a1 = dwflow_getArea(xsect_2, y1, wSlot)
	r1 = getHydRad(xsect_2, y1)
	wSlot = getSlotWidth(xsect_2, y2)
	a2 = dwflow_getArea(xsect_2, y2, wSlot)
	yMid = float64(float64(0.5) * (float64(y1 + y2)))
	wSlot = getSlotWidth(xsect_2, yMid)
	aMid = dwflow_getArea(xsect_2, yMid, wSlot)
	rMid = getHydRad(xsect_2, yMid)
	if (y1 >= xsect_2.yFull) && (y2 >= xsect_2.yFull) {
		isFull = 1
	}
	if ((((Link[j].flowClass == DRY) || (Link[j].flowClass == UP_DRY)) || (Link[j].flowClass == DN_DRY)) || (isClosed != 0)) || (aMid <= float64(0.0001)) {
		Conduit[k].a1 = float64(float64(0.5) * (float64(a1 + a2)))
		Conduit[k].q1 = float64(0.0)
		Conduit[k].q2 = float64(0.0)
		Link[j].dqdh = float64(float64(float64(float64(float64(32.2)*dt)*aMid)/length) * barrels)
		Link[j].froude = float64(0.0)
		if (yMid) <= (Link[j].xsect.yFull) {
			Link[j].newDepth = (yMid)
		} else {
			Link[j].newDepth = (Link[j].xsect.yFull)
		}
		Link[j].newVolume = float64(float64(Conduit[k].a1*link_getLength(j)) * barrels)
		Link[j].newFlow = float64(0.0)
		return
	}
	v = float64(qLast / aMid)
	if M.Fabs(v) > MAXVELOCITY {
		v = float64(MAXVELOCITY * float64((R.If[int32](((qLast) < float64(0)), (-1), (1)))))
	}
	Link[j].froude = link_getFroude(j, v, yMid)
	if (Link[j].flowClass == SUBCRITICAL) && (Link[j].froude > float64(1.0)) {
		Link[j].flowClass = SUPCRITICAL
	}
	if Link[j].froude <= float64(0.5) {
		sigma = float64(1.0)
	} else if Link[j].froude >= float64(1.0) {
		sigma = float64(0.0)
	} else {
		sigma = float64(float64(2.0) * (float64(float64(1.0) - Link[j].froude)))
	}
	rho = float64(1.0)
	if ((!(isFull != 0)) && (qLast > float64(0.0))) && (h1 >= h2) {
		rho = sigma
	}
	aWtd = float64(a1 + float64((float64(aMid-a1))*rho))
	rWtd = float64(r1 + float64((float64(rMid-r1))*rho))
	if InertDamping == NO_DAMPING {
		sigma = float64(1.0)
	} else if InertDamping == FULL_DAMPING {
		sigma = float64(0.0)
	}
	if (isFull != 0) && (!(xsect_isOpen(xsect_2.type_) != 0)) {
		sigma = float64(0.0)
	}
	if (xsect_2.type_ == FORCE_MAIN) && (isFull != 0) {
		dq1 = float64(dt * forcemain_getFricSlope(j, M.Fabs(v), rMid))
	} else {
		dq1 = float64(float64(float64(dt*Conduit[k].roughFactor)/M.Pow(rWtd, float64(1.33333))) * M.Fabs(v))
	}
	dq2 = float64(float64(float64(float64(dt*float64(32.2))*aWtd)*(float64(h2-h1))) / length)
	dq3 = float64(0.0)
	dq4 = float64(0.0)
	if sigma > float64(0.0) {
		dq3 = float64(float64(float64(float64(2.0)*v)*(float64(aMid-aOld))) * sigma)
		dq4 = float64(float64(float64(float64(float64(dt*v)*v)*(float64(a2-a1)))/length) * sigma)
	}
	dq5 = float64(0.0)
	if Conduit[k].hasLosses != 0 {
		dq5 = float64(float64(float64(findLocalLosses(j, a1, a2, aMid, qLast)/float64(2.0))/length) * dt)
	}
	dq6 = float64(float64(float64(float64(link_getLossRate(j, DW, qLast, dt)*float64(2.5))*dt)*v) / link_getLength(j))
	denom = float64(float64(float64(1.0)+dq1) + dq5)
	q = float64((float64(float64(float64(float64(qOld-dq2)+dq3)+dq4) + dq6)) / denom)
	Link[j].dqdh = float64(float64(float64(float64(float64(float64(float64(1.0)/denom)*float64(32.2))*dt)*aWtd)/length) * barrels)
	Link[j].inletControl = 0
	Link[j].normalFlow = 0
	if q > float64(0.0) {
		if (xsect_2.culvertCode > 0) && (!(isFull != 0)) {
			q = culvert_getInflow(j, q, h1)
		} else if ((NormalFlowLtd != NEITHER) && (y1 < Link[j].xsect.yFull)) && ((Link[j].flowClass == SUBCRITICAL) || (Link[j].flowClass == SUPCRITICAL)) {
			q = checkNormalFlow(j, q, y1, y2, a1, r1)
		}
	}
	if steps > 0 {
		q = float64(float64((float64(float64(1.0)-omega))*qLast) + float64(omega*q))
		if float64(q*qLast) < float64(0.0) {
			q = float64(float64(0.001) * float64((R.If[int32](((q) < float64(0)), (-1), (1)))))
		}
	}
	if Link[j].qLimit > float64(0.0) {
		if M.Fabs(q) > Link[j].qLimit {
			q = float64(float64((R.If[int32](((q) < float64(0)), (-1), (1)))) * Link[j].qLimit)
		}
	}
	if link_setFlapGate(j, n1, n2, q) != 0 {
		q = float64(0.0)
	}
	if (q > float64(0.0001)) && (Node[n1].newDepth <= float64(0.0001)) {
		q = float64(0.0001)
	}
	if (q < (-(float64(0.0001)))) && (Node[n2].newDepth <= float64(0.0001)) {
		q = (-(float64(0.0001)))
	}
	Conduit[k].a1 = aMid
	Conduit[k].q1 = q
	Conduit[k].q2 = q
	if (yMid) <= (xsect_2.yFull) {
		Link[j].newDepth = (yMid)
	} else {
		Link[j].newDepth = (xsect_2.yFull)
	}
	aMid = float64((float64(a1 + a2)) / float64(2.0))
	Conduit[k].fullState = link_getFullState(a1, a2, xsect_2.aFull)
	Link[j].newVolume = float64(float64(aMid*link_getLength(j)) * barrels)
	Link[j].newFlow = float64(q * barrels)
}

// dwflow.c:297
func getFlowClass(j int32, q float64, h1 float64, h2 float64, y1 float64, y2 float64, yC *float64, yN *float64, fasnh *float64) int32 {
	var (
		n1        int32
		n2        int32
		flowClass int32
		ycMin     float64
		ycMax     float64
		z1        float64
		z2        float64
	)
	_, _, _, _, _, _, _ = n1, n2, flowClass, ycMin, ycMax, z1, z2
	n1 = 0
	n2 = 0
	flowClass = 0
	ycMin = 0
	ycMax = 0
	z1 = 0
	z2 = 0
	n1 = Link[j].node1
	n2 = Link[j].node2
	z1 = Link[j].offset1
	z2 = Link[j].offset2
	if Node[n1].type_ == OUTFALL {
		if (float64(0.0)) >= (float64(z1 - Node[n1].newDepth)) {
			z1 = (float64(0.0))
		} else {
			z1 = (float64(z1 - Node[n1].newDepth))
		}
	}
	if Node[n2].type_ == OUTFALL {
		if (float64(0.0)) >= (float64(z2 - Node[n2].newDepth)) {
			z2 = (float64(0.0))
		} else {
			z2 = (float64(z2 - Node[n2].newDepth))
		}
	}
	flowClass = SUBCRITICAL
	*fasnh = float64(1.0)
	if (y1 > float64(0.0001)) && (y2 > float64(0.0001)) {
		if q < float64(0.0) {
			if z1 > float64(0.0) {
				*yN = link_getYnorm(j, M.Fabs(q))
				*yC = link_getYcrit(j, M.Fabs(q))
				if (*yN) <= (*yC) {
					ycMin = (*yN)
				} else {
					ycMin = (*yC)
				}
				if y1 < ycMin {
					flowClass = UP_CRITICAL
				}
			}
		} else {
			if z2 > float64(0.0) {
				*yN = link_getYnorm(j, M.Fabs(q))
				*yC = link_getYcrit(j, M.Fabs(q))
				if (*yN) <= (*yC) {
					ycMin = (*yN)
				} else {
					ycMin = (*yC)
				}
				if (*yN) >= (*yC) {
					ycMax = (*yN)
				} else {
					ycMax = (*yC)
				}
				if y2 < ycMin {
					flowClass = DN_CRITICAL
				} else if y2 < ycMax {
					if float64(ycMax-ycMin) < float64(0.0001) {
						*fasnh = float64(0.0)
					} else {
						*fasnh = float64((float64(ycMax - y2)) / (float64(ycMax - ycMin)))
					}
				}
			}
		}
	} else if (y1 <= float64(0.0001)) && (y2 <= float64(0.0001)) {
		flowClass = DRY
	} else if y2 > float64(0.0001) {
		if h2 < float64(Node[n1].invertElev+Link[j].offset1) {
			flowClass = UP_DRY
		} else if z1 > float64(0.0) {
			*yN = link_getYnorm(j, M.Fabs(q))
			*yC = link_getYcrit(j, M.Fabs(q))
			flowClass = UP_CRITICAL
		}
	} else {
		if h1 < float64(Node[n2].invertElev+Link[j].offset2) {
			flowClass = DN_DRY
		} else if z2 > float64(0.0) {
			*yN = link_getYnorm(j, M.Fabs(q))
			*yC = link_getYcrit(j, M.Fabs(q))
			flowClass = DN_CRITICAL
		}
	}
	return flowClass
}

// dwflow.c:417
func findSurfArea(j int32, q float64, length float64, h1 *float64, h2 *float64, y1 *float64, y2 *float64) {
	var (
		n1            int32
		n2            int32
		flowDepth1    float64
		flowDepth2    float64
		flowDepthMid  float64
		width1        float64
		width2        float64
		widthMid      float64
		surfArea1     float64
		surfArea2     float64
		criticalDepth float64
		normalDepth   float64
		fullDepth     float64
		fasnh         float64
		xsect_2       *TXsect
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = n1, n2, flowDepth1, flowDepth2, flowDepthMid, width1, width2, widthMid, surfArea1, surfArea2, criticalDepth, normalDepth
	_, _, _ = fullDepth, fasnh, xsect_2
	n1 = 0
	n2 = 0
	flowDepth1 = 0
	flowDepth2 = 0
	flowDepthMid = 0
	width1 = 0
	width2 = 0
	widthMid = 0
	surfArea1 = float64(0.0)
	surfArea2 = float64(0.0)
	criticalDepth = 0
	normalDepth = 0
	fullDepth = 0
	fasnh = float64(1.0)
	xsect_2 = &Link[j].xsect
	n1 = Link[j].node1
	n2 = Link[j].node2
	flowDepth1 = (*y1)
	flowDepth2 = (*y2)
	normalDepth = float64((float64(flowDepth1 + flowDepth2)) / float64(2.0))
	criticalDepth = normalDepth
	fullDepth = xsect_2.yFull
	if (flowDepth1 >= fullDepth) && (flowDepth2 >= fullDepth) {
		Link[j].flowClass = SUBCRITICAL
	} else {
		Link[j].flowClass = getFlowClass(j, q, (*h1), (*h2), (*y1), (*y2), &criticalDepth, &normalDepth, &fasnh)
	}
	switch Link[j].flowClass {
	case SUBCRITICAL:
		flowDepthMid = float64(float64(0.5) * (float64(flowDepth1 + flowDepth2)))
		if flowDepthMid < float64(0.0001) {
			flowDepthMid = float64(0.0001)
		}
		width1 = dwflow_getWidth(xsect_2, flowDepth1)
		width2 = dwflow_getWidth(xsect_2, flowDepth2)
		widthMid = dwflow_getWidth(xsect_2, flowDepthMid)
		surfArea1 = float64(float64((float64(width1+widthMid))*length) / float64(4.0))
		surfArea2 = float64(float64(float64((float64(widthMid+width2))*length)/float64(4.0)) * fasnh)
		break
	case UP_CRITICAL:
		flowDepth1 = criticalDepth
		if normalDepth < criticalDepth {
			flowDepth1 = normalDepth
		}
		if (flowDepth1) >= (float64(0.0001)) {
			flowDepth1 = (flowDepth1)
		} else {
			flowDepth1 = (float64(0.0001))
		}
		*h1 = float64(float64(Node[n1].invertElev+Link[j].offset1) + flowDepth1)
		flowDepthMid = float64(float64(0.5) * (float64(flowDepth1 + flowDepth2)))
		if flowDepthMid < float64(0.0001) {
			flowDepthMid = float64(0.0001)
		}
		width2 = dwflow_getWidth(xsect_2, flowDepth2)
		widthMid = dwflow_getWidth(xsect_2, flowDepthMid)
		surfArea2 = float64(float64((float64(widthMid+width2))*length) * float64(0.5))
		break
	case DN_CRITICAL:
		flowDepth2 = criticalDepth
		if normalDepth < criticalDepth {
			flowDepth2 = normalDepth
		}
		if (flowDepth2) >= (float64(0.0001)) {
			flowDepth2 = (flowDepth2)
		} else {
			flowDepth2 = (float64(0.0001))
		}
		*h2 = float64(float64(Node[n2].invertElev+Link[j].offset2) + flowDepth2)
		width1 = dwflow_getWidth(xsect_2, flowDepth1)
		flowDepthMid = float64(float64(0.5) * (float64(flowDepth1 + flowDepth2)))
		if flowDepthMid < float64(0.0001) {
			flowDepthMid = float64(0.0001)
		}
		widthMid = dwflow_getWidth(xsect_2, flowDepthMid)
		surfArea1 = float64(float64((float64(width1+widthMid))*length) * float64(0.5))
		break
	case UP_DRY:
		flowDepth1 = float64(0.0001)
		flowDepthMid = float64(float64(0.5) * (float64(flowDepth1 + flowDepth2)))
		if flowDepthMid < float64(0.0001) {
			flowDepthMid = float64(0.0001)
		}
		width1 = dwflow_getWidth(xsect_2, flowDepth1)
		width2 = dwflow_getWidth(xsect_2, flowDepth2)
		widthMid = dwflow_getWidth(xsect_2, flowDepthMid)
		surfArea2 = float64(float64((float64(widthMid+width2))*length) / float64(4.0))
		if Link[j].offset1 <= float64(0.0) {
			surfArea1 = float64(float64((float64(width1+widthMid))*length) / float64(4.0))
		}
		break
	case DN_DRY:
		flowDepth2 = float64(0.0001)
		flowDepthMid = float64(float64(0.5) * (float64(flowDepth1 + flowDepth2)))
		if flowDepthMid < float64(0.0001) {
			flowDepthMid = float64(0.0001)
		}
		width1 = dwflow_getWidth(xsect_2, flowDepth1)
		width2 = dwflow_getWidth(xsect_2, flowDepth2)
		widthMid = dwflow_getWidth(xsect_2, flowDepthMid)
		surfArea1 = float64(float64((float64(widthMid+width1))*length) / float64(4.0))
		if Link[j].offset2 <= float64(0.0) {
			surfArea2 = float64(float64((float64(width2+widthMid))*length) / float64(4.0))
		}
		break
	case DRY:
		surfArea1 = float64(float64(float64(0.0001)*length) / float64(2.0))
		surfArea2 = surfArea1
		break
	}
	Link[j].surfArea1 = surfArea1
	Link[j].surfArea2 = surfArea2
	*y1 = flowDepth1
	*y2 = flowDepth2
}

// dwflow.c:554
func findLocalLosses(j int32, a1 float64, a2 float64, aMid float64, q float64) float64 {
	var (
		losses float64
	)
	_ = losses
	losses = float64(0.0)
	q = M.Fabs(q)
	if a1 > float64(0.0001) {
		losses = float64(losses + (float64(Link[j].cLossInlet * (float64(q / a1)))))
	}
	if a2 > float64(0.0001) {
		losses = float64(losses + (float64(Link[j].cLossOutlet * (float64(q / a2)))))
	}
	if aMid > float64(0.0001) {
		losses = float64(losses + (float64(Link[j].cLossAvg * (float64(q / aMid)))))
	}
	return losses
}

// dwflow.c:575
func getSlotWidth(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		yNorm float64
	)
	_ = yNorm
	yNorm = float64(y_2 / xsect_2.yFull)
	if ((SurchargeMethod != SLOT) || (xsect_isOpen(xsect_2.type_) != 0)) || (yNorm < CrownCutoff) {
		return float64(0.0)
	}
	if yNorm > float64(1.78) {
		return float64(float64(0.01) * xsect_2.wMax)
	}
	return float64(float64(xsect_2.wMax*float64(0.5423)) * M.Exp((-(M.Pow(yNorm, float64(2.4))))))
}

// dwflow.c:592
func dwflow_getWidth(xsect_2 *TXsect, y_2 float64) float64 {
	var (
		wSlot float64
	)
	_ = wSlot
	wSlot = getSlotWidth(xsect_2, y_2)
	if wSlot > float64(0.0) {
		return wSlot
	}
	if (float64(y_2/xsect_2.yFull) >= CrownCutoff) && (!(xsect_isOpen(xsect_2.type_) != 0)) {
		y_2 = float64(CrownCutoff * xsect_2.yFull)
	}
	return xsect_getWofY(xsect_2, y_2)
}

// dwflow.c:609
func dwflow_getArea(xsect_2 *TXsect, y_2 float64, wSlot float64) float64 {
	if y_2 >= xsect_2.yFull {
		return float64(xsect_2.aFull + float64((float64(y_2-xsect_2.yFull))*wSlot))
	}
	return xsect_getAofY(xsect_2, y_2)
}

// dwflow.c:623
func getHydRad(xsect_2 *TXsect, y_2 float64) float64 {
	if y_2 >= xsect_2.yFull {
		return xsect_2.rFull
	}
	return xsect_getRofY(xsect_2, y_2)
}

// dwflow.c:637
func checkNormalFlow(j int32, q float64, y1 float64, y2 float64, a1 float64, r1 float64) float64 {
	var (
		check      int32
		k          int32
		n1         int32
		n2         int32
		hasOutfall int32
		qNorm      float64
		f1         float64
	)
	_, _, _, _, _, _, _ = check, k, n1, n2, hasOutfall, qNorm, f1
	check = 0
	k = Link[j].subIndex
	n1 = Link[j].node1
	n2 = Link[j].node2
	hasOutfall = R.B2I(((Node[n1].type_ == OUTFALL) || (Node[n2].type_ == OUTFALL)))
	qNorm = 0
	f1 = 0
	if ((NormalFlowLtd == SLOPE) || (NormalFlowLtd == BOTH)) || (hasOutfall != 0) {
		if y1 < y2 {
			check = 1
		}
	}
	if ((!(check != 0)) && ((NormalFlowLtd == FROUDE) || (NormalFlowLtd == BOTH))) && (!(hasOutfall != 0)) {
		if (y1 > float64(0.0001)) && (y2 > float64(0.0001)) {
			f1 = link_getFroude(j, float64(q/a1), y1)
			if f1 >= float64(1.0) {
				check = 1
			}
		}
	}
	if check != 0 {
		qNorm = float64(float64(Conduit[k].beta*a1) * M.Pow(r1, float64(float64(2.0)/float64(3.0))))
		if qNorm < q {
			Link[j].normalFlow = 1
			return qNorm
		}
	}
	return q
}
