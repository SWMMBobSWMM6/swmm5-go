// Generated from EPA SWMM 5.2.4 (inlet.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var InletDesigns []TInletDesign
var InletDesignCount int32
var UsesInlets int32
var InletTypeWords = [8]R.CPtr{L718, L719, L0, L720, L721, L722, L723, (R.CPtr{})}
var GrateTypeWords = [9]R.CPtr{L724, L725, L726, L727, L728, L729, L730, L731, (R.CPtr{})}
var ThroatAngleWords = [4]R.CPtr{L732, L733, L734, (R.CPtr{})}
var PlacementTypeWords = [4]R.CPtr{L735, L736, L737}
var SplashCoeffs = [7][4]float64{[4]float64{float64(2.22), float64(4.03), float64(0.65), float64(0.06)}, [4]float64{float64(0.74), float64(2.44), float64(0.27), float64(0.02)}, [4]float64{float64(1.76), float64(3.12), float64(0.45), float64(0.03)}, [4]float64{float64(0.30), float64(4.85), float64(1.31), float64(0.15)}, [4]float64{float64(0.99), float64(2.64), float64(0.36), float64(0.03)}, [4]float64{float64(0.51), float64(2.34), float64(0.2), float64(0.01)}, [4]float64{float64(0.28), float64(2.28), float64(0.18), float64(0.01)}}
var GrateOpeningRatios = [8]float64{float64(0.90), float64(0.80), float64(0.60), float64(0.35), float64(0.17), float64(0.34), float64(0.80), float64(1.00)}
var Sx float64
var SL float64
var Sw float64
var a float64
var W float64
var T float64
var n float64
var Nsides int32
var Tcrown float64
var Beta float64
var Qfactor float64
var xsect *TXsect
var InletFlow []float64
var FirstInlet *TInlet

// inlet.c:257
func inlet_create(numInlets int32) int32 {
	var (
		i int32
	)
	_ = i
	i = 0
	InletDesigns = nil
	InletFlow = nil
	InletDesignCount = 0
	UsesInlets = 0
	FirstInlet = nil
	InletDesigns = make([]TInletDesign, int(uint64(numInlets)))
	if InletDesigns == nil {
		return ERR_MEMORY
	}
	InletDesignCount = numInlets
	InletFlow = make([]float64, int(uint64(Nobjects[NODE])))
	if InletFlow == nil {
		return ERR_MEMORY
	}
	for i = 0; i < InletDesignCount; i++ {
		InletDesigns[i].customCurve = (-1)
		InletDesigns[i].curbInlet.length = float64(0.0)
		InletDesigns[i].grateInlet.length = float64(0.0)
		InletDesigns[i].slottedInlet.length = float64(0.0)
		InletDesigns[i].type_ = CUSTOM_INLET
	}
	return 0
}

// inlet.c:291
func inlet_delete() {
	var (
		inlet     *TInlet
		nextInlet *TInlet
	)
	_, _ = inlet, nextInlet
	inlet = FirstInlet
	nextInlet = nil
	for inlet != nil {
		nextInlet = inlet.nextInlet
		_ = inlet
		inlet = nextInlet
	}
	FirstInlet = nil
	{
		if InletFlow != nil {
			_ = InletFlow
			InletFlow = nil
		}
	}
	{
		if InletDesigns != nil {
			_ = InletDesigns
			InletDesigns = nil
		}
	}
}

// inlet.c:313
func inlet_readDesignParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		i int32
	)
	_ = i
	i = 0
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	i = project_findObject(INLET, tok[0])
	if i < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	InletDesigns[i].ID = project_findID(INLET, tok[0])
	InletDesigns[i].type_ = findmatch(tok[1], InletTypeWords[:])
	switch InletDesigns[i].type_ {
	case GRATE_INLET, DROP_GRATE_INLET:
		return readGrateInletParams(i, tok, ntoks)
	case CURB_INLET, DROP_CURB_INLET:
		return readCurbInletParams(i, tok, ntoks)
	case SLOTTED_INLET:
		return readSlottedInletParams(i, tok, ntoks)
	case CUSTOM_INLET:
		return readCustomInletParams(i, tok, ntoks)
	default:
		return error_setInpError(ERR_KEYWORD, tok[1])
	}
	return 0
}

// inlet.c:362
func inlet_readUsageParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		linkIndex   int32
		designIndex int32
		nodeIndex   int32
		numInlets   int32
		placement   int32
		flowLimit   float64
		pctClogged  float64
		aLocal      float64
		wLocal      float64
		inlet       *TInlet
	)
	_, _, _, _, _, _, _, _, _, _ = linkIndex, designIndex, nodeIndex, numInlets, placement, flowLimit, pctClogged, aLocal, wLocal, inlet
	linkIndex = 0
	designIndex = 0
	nodeIndex = 0
	numInlets = 1
	placement = AUTOMATIC
	flowLimit = float64(0.0)
	pctClogged = float64(0.0)
	aLocal = float64(0.0)
	wLocal = float64(0.0)
	inlet = nil
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	linkIndex = project_findObject(LINK, tok[0])
	if linkIndex < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	designIndex = project_findObject(INLET, tok[1])
	if designIndex < 0 {
		return error_setInpError(ERR_NAME, tok[1])
	}
	nodeIndex = project_findObject(NODE, tok[2])
	if nodeIndex < 0 {
		return error_setInpError(ERR_NAME, tok[2])
	}
	if ntoks > 3 {
		if (!(getInt(tok[3], &numInlets) != 0)) || (numInlets < 1) {
			return error_setInpError(ERR_NUMBER, tok[3])
		}
	}
	if ntoks > 4 {
		if ((!(getDouble(tok[4], &pctClogged) != 0)) || (pctClogged < float64(0.0))) || (pctClogged > float64(99.0)) {
			return error_setInpError(ERR_NUMBER, tok[4])
		}
	}
	if ntoks > 5 {
		if (!(getDouble(tok[5], &flowLimit) != 0)) || (flowLimit < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[5])
		}
	}
	if ntoks > 6 {
		if (!(getDouble(tok[6], &aLocal) != 0)) || (aLocal < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[6])
		}
	}
	if ntoks > 7 {
		if (!(getDouble(tok[7], &wLocal) != 0)) || (wLocal < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[7])
		}
	}
	if ntoks > 8 {
		placement = findmatch(tok[8], PlacementTypeWords[:])
		if placement < 0 {
			return error_setInpError(ERR_KEYWORD, tok[8])
		}
	}
	inlet = Link[linkIndex].inlet
	if inlet == nil {
		inlet = new(TInlet)
		if !(inlet != nil) {
			return error_setInpError(ERR_MEMORY, L0)
		}
		Link[linkIndex].inlet = inlet
		inlet.nextInlet = FirstInlet
		FirstInlet = inlet
	}
	inlet.linkIndex = linkIndex
	inlet.designIndex = designIndex
	inlet.nodeIndex = nodeIndex
	inlet.numInlets = numInlets
	inlet.placement = placement
	inlet.clogFactor = float64(float64(1.0) - (float64(pctClogged / float64(100.0))))
	inlet.flowLimit = float64(flowLimit / UCF(FLOW))
	inlet.localDepress = float64(aLocal / UCF(LENGTH))
	inlet.localWidth = float64(wLocal / UCF(LENGTH))
	inlet.flowFactor = float64(0.0)
	inlet.backflowRatio = float64(0.0)
	initInletStats(inlet)
	UsesInlets = 1
	return 0
}

// inlet.c:464
func inlet_validate() {
	var (
		i          int32
		j          int32
		inletType  int32
		inletValid int32
		inlet      *TInlet
		prevInlet  *TInlet
	)
	_, _, _, _, _, _ = i, j, inletType, inletValid, inlet, prevInlet
	i = 0
	j = 0
	inletType = 0
	inletValid = 0
	inlet = nil
	prevInlet = nil
	if !(UsesInlets != 0) {
		return
	}
	prevInlet = FirstInlet
	inlet = FirstInlet
	for inlet != nil {
		inletValid = 0
		i = inlet.linkIndex
		xsect = &Link[i].xsect
		inletType = InletDesigns[inlet.designIndex].type_
		if inletType == CUSTOM_INLET {
			j = InletDesigns[inlet.designIndex].customCurve
			if j >= 0 {
				if (Curve[j].curveType == DIVERSION_CURVE) || (Curve[j].curveType == RATING_CURVE) {
					inletValid = 1
				}
			}
		} else if ((xsect.type_ == TRAPEZOIDAL) || (xsect.type_ == RECT_OPEN)) && ((inletType == DROP_GRATE_INLET) || (inletType == DROP_CURB_INLET)) {
			inletValid = 1
		} else if ((xsect.type_ == STREET_XSECT) && (inletType != DROP_GRATE_INLET)) && (inletType != DROP_CURB_INLET) {
			inletValid = 1
		}
		if inletValid != 0 {
			Node[Link[i].node2].inlet = BYPASS
			Node[inlet.nodeIndex].inlet = CAPTURE
			inlet.backflow = float64(0.0)
			getConduitGeometry(inlet)
			inlet.flowFactor = float64(float64((float64(float64(0.56)/n))*M.Pow(SL, float64(0.5))) * M.Pow(Sx, float64(1.67)))
			prevInlet = inlet
			inlet = inlet.nextInlet
		} else {
			report_writeWarningMsg(L187, Link[i].ID)
			if inlet == FirstInlet {
				FirstInlet = inlet.nextInlet
				prevInlet = FirstInlet
				_ = inlet
				inlet = FirstInlet
			} else {
				prevInlet.nextInlet = inlet.nextInlet
				_ = inlet
				inlet = prevInlet.nextInlet
			}
			Link[i].inlet = nil
		}
	}
	getBackflowRatios()
}

// inlet.c:555
func inlet_findCapturedFlows(tStep float64) {
	var (
		i         int32
		j         int32
		m         int32
		placement int32
		q         float64
		inlet     *TInlet
	)
	_, _, _, _, _, _ = i, j, m, placement, q, inlet
	i = 0
	j = 0
	m = 0
	placement = 0
	q = 0
	inlet = nil
	if !(UsesInlets != 0) {
		return
	}
	R.Memset(InletFlow, 0, (uint64(Nobjects[NODE]) * 8))
	if RouteModel != DW {
		for j = 0; j < Nobjects[NODE]; j++ {
			if (float64(0.0)) >= (Node[j].newLatFlow) {
				Node[j].inflow = (float64(0.0))
			} else {
				Node[j].inflow = (Node[j].newLatFlow)
			}
		}
		for i = 0; i < Nobjects[LINK]; i++ {
			Node[Link[i].node2].inflow = float64(Node[Link[i].node2].inflow + (func() float64 {
				if (float64(0.0)) >= (Link[i].newFlow) {
					return (float64(0.0))
				}
				return (Link[i].newFlow)
			}()))
		}
	}
	for inlet = FirstInlet; inlet != nil; inlet = inlet.nextInlet {
		i = inlet.linkIndex
		j = Link[i].node2
		m = inlet.nodeIndex
		placement = getInletPlacement(inlet, j)
		if InletDesigns[inlet.designIndex].type_ == CUSTOM_INLET {
			q = M.Fabs(Link[i].newFlow)
			inlet.flowCapture = getCustomCapturedFlow(inlet, q, Node[j].newDepth)
		} else if placement == ON_GRADE {
			q = M.Fabs(Link[i].newFlow)
			inlet.flowCapture = getOnGradeCapturedFlow(inlet, q, Node[j].newDepth)
		} else {
			q = Node[j].inflow
			inlet.flowCapture = getOnSagCapturedFlow(inlet, q, Node[j].newDepth)
		}
		if M.Fabs(inlet.flowCapture) < float64(0.0001) {
			inlet.flowCapture = float64(0.0)
		}
		InletFlow[j] = float64(InletFlow[j] + inlet.flowCapture)
		inlet.backflow = float64(Node[m].overflow * inlet.backflowRatio)
		if M.Fabs(inlet.backflow) < float64(0.0001) {
			inlet.backflow = float64(0.0)
		}
	}
	for inlet = FirstInlet; inlet != nil; inlet = inlet.nextInlet {
		i = inlet.linkIndex
		j = Link[i].node2
		m = inlet.nodeIndex
		if (RouteModel != DW) && (getInletPlacement(inlet, j) == ON_SAG) {
			q = float64(Node[j].newVolume / tStep)
			q = float64(q + (func() float64 {
				if (Node[j].inflow) >= (float64(0.0)) {
					return (Node[j].inflow)
				}
				return (float64(0.0))
			}()))
			if InletFlow[j] > q {
				inlet.flowCapture = float64(inlet.flowCapture * (float64(q / InletFlow[j])))
			}
		}
		Node[j].newLatFlow = float64(Node[j].newLatFlow - (float64(inlet.flowCapture - inlet.backflow)))
		Node[m].newLatFlow = float64(Node[m].newLatFlow + inlet.flowCapture)
		if getDateTime(NewRoutingTime) > ReportStart {
			updateInletStats(inlet, M.Fabs(Link[i].newFlow))
		}
	}
}

// inlet.c:655
func inlet_adjustQualInflows() {
	var (
		i     int32
		j     int32
		m     int32
		p     int32
		qNet  float64
		inlet *TInlet
	)
	_, _, _, _, _, _ = i, j, m, p, qNet, inlet
	i = 0
	j = 0
	m = 0
	p = 0
	qNet = 0
	inlet = nil
	if !(UsesInlets != 0) {
		return
	}
	if (IgnoreQuality != 0) || (Nobjects[POLLUT] == 0) {
		return
	}
	for inlet = FirstInlet; inlet != nil; inlet = inlet.nextInlet {
		i = inlet.linkIndex
		j = Link[i].node2
		m = inlet.nodeIndex
		qNet = float64(inlet.flowCapture - inlet.backflow)
		if qNet > float64(0.0) {
			Node[m].qualInflow = float64(Node[m].qualInflow + qNet)
			for p = 0; p < Nobjects[POLLUT]; p++ {
				Node[m].newQual[p] = float64(Node[m].newQual[p] + (float64(qNet * Node[j].oldQual[p])))
			}
		} else {
			qNet = (-qNet)
			Node[j].qualInflow = float64(Node[j].qualInflow + qNet)
			for p = 0; p < Nobjects[POLLUT]; p++ {
				Node[j].newQual[p] = float64(Node[j].newQual[p] + (float64(qNet * Node[m].oldQual[p])))
			}
		}
	}
}

// inlet.c:706
func inlet_adjustQualOutflows() {
	var (
		j     int32
		p     int32
		q     float64
		w     float64
		inlet *TInlet
	)
	_, _, _, _, _ = j, p, q, w, inlet
	j = 0
	p = 0
	q = 0
	w = 0
	inlet = nil
	for j = 0; j < Nobjects[NODE]; j++ {
		if Node[j].inlet == CAPTURE {
			q = Node[j].overflow
			if q > float64(0.0) {
				StepFlowTotals.flooding = float64(StepFlowTotals.flooding - q)
				if !(IgnoreQuality != 0) {
					for p = 0; p < Nobjects[POLLUT]; p++ {
						w = float64(q * Node[j].newQual[p])
						StepQualTotals[p].flooding = float64(StepQualTotals[p].flooding - w)
					}
				}
			}
		}
	}
	if (!(IgnoreQuality != 0)) && (Nobjects[POLLUT] > 0) {
		for inlet = FirstInlet; inlet != nil; inlet = inlet.nextInlet {
			j = Link[inlet.linkIndex].node2
			q = float64(inlet.flowCapture - inlet.backflow)
			if (q > float64(0.0)) && (Node[j].newLatFlow < float64(0.0)) {
				for p = 0; p < Nobjects[POLLUT]; p++ {
					w = float64(q * Node[j].newQual[p])
					StepQualTotals[p].outflow = float64(StepQualTotals[p].outflow - w)
				}
			}
		}
	}
}

// inlet.c:774
func inlet_writeStatsReport() {
	var (
		j      int32
		header int32
	)
	_, _ = j, header
	j = 0
	header = 0
	if Nobjects[STREET] == 0 {
		return
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		if Link[j].xsect.type_ == STREET_XSECT {
			if !(header != 0) {
				writeStreetStatsHeader()
				header = 1
			}
			writeStreetStats(j)
		}
	}
	report_writeLine(L0)
}

// inlet.c:801
func inlet_capturedFlow(i int32) float64 {
	if Link[i].inlet != nil {
		return Link[i].inlet.flowCapture
	}
	return float64(0.0)
}

// inlet.c:814
func readGrateInletParams(i int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		grateType int32
		width     float64
		length    float64
		areaRatio float64
		vSplash   float64
	)
	_, _, _, _, _ = grateType, width, length, areaRatio, vSplash
	grateType = 0
	width = 0
	length = 0
	areaRatio = float64(0.0)
	vSplash = float64(0.0)
	if ntoks < 5 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	if (!(getDouble(tok[2], &length) != 0)) || (length <= float64(0.0)) {
		return error_setInpError(ERR_NUMBER, tok[2])
	}
	if (!(getDouble(tok[3], &width) != 0)) || (width <= float64(0.0)) {
		return error_setInpError(ERR_NUMBER, tok[3])
	}
	grateType = findmatch(tok[4], GrateTypeWords[:])
	if grateType < 0 {
		return error_setInpError(ERR_KEYWORD, tok[4])
	}
	if grateType == GENERIC {
		if ntoks < 6 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		if ((!(getDouble(tok[5], &areaRatio) != 0)) || (areaRatio <= float64(0.0))) || (areaRatio > float64(1.0)) {
			return error_setInpError(ERR_NUMBER, tok[5])
		}
		if ntoks > 6 {
			if (!(getDouble(tok[6], &vSplash) != 0)) || (vSplash < float64(0.0)) {
				return error_setInpError(ERR_NUMBER, tok[6])
			}
		}
	}
	InletDesigns[i].grateInlet.length = float64(length / UCF(LENGTH))
	InletDesigns[i].grateInlet.width = float64(width / UCF(LENGTH))
	InletDesigns[i].grateInlet.type_ = grateType
	InletDesigns[i].grateInlet.fracOpenArea = areaRatio
	InletDesigns[i].grateInlet.splashVeloc = float64(vSplash / UCF(LENGTH))
	if (InletDesigns[i].type_ == GRATE_INLET) && (InletDesigns[i].curbInlet.length > float64(0.0)) {
		InletDesigns[i].type_ = COMBO_INLET
	}
	return 0
}

// inlet.c:868
func readCurbInletParams(i int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		throatAngle int32
		height      float64
		length      float64
	)
	_, _, _ = throatAngle, height, length
	throatAngle = 0
	height = 0
	length = 0
	if ntoks < 4 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	if (!(getDouble(tok[2], &length) != 0)) || (length <= float64(0.0)) {
		return error_setInpError(ERR_NUMBER, tok[2])
	}
	if (!(getDouble(tok[3], &height) != 0)) || (height <= float64(0.0)) {
		return error_setInpError(ERR_NUMBER, tok[3])
	}
	throatAngle = VERTICAL_THROAT
	if (InletDesigns[i].type_ == CURB_INLET) && (ntoks > 4) {
		throatAngle = findmatch(tok[4], ThroatAngleWords[:])
		if throatAngle < 0 {
			return error_setInpError(ERR_KEYWORD, tok[4])
		}
	}
	InletDesigns[i].curbInlet.length = float64(length / UCF(LENGTH))
	InletDesigns[i].curbInlet.height = float64(height / UCF(LENGTH))
	InletDesigns[i].curbInlet.throatAngle = throatAngle
	if (InletDesigns[i].type_ == CURB_INLET) && (InletDesigns[i].grateInlet.length > float64(0.0)) {
		InletDesigns[i].type_ = COMBO_INLET
	}
	return 0
}

// inlet.c:911
func readSlottedInletParams(i int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		width  float64
		length float64
	)
	_, _ = width, length
	width = 0
	length = 0
	if ntoks < 4 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	if (!(getDouble(tok[2], &length) != 0)) || (length <= float64(0.0)) {
		return error_setInpError(ERR_NUMBER, tok[2])
	}
	if (!(getDouble(tok[3], &width) != 0)) || (width <= float64(0.0)) {
		return error_setInpError(ERR_NUMBER, tok[3])
	}
	InletDesigns[i].slottedInlet.length = float64(length / UCF(LENGTH))
	InletDesigns[i].slottedInlet.width = float64(width / UCF(LENGTH))
	return 0
}

// inlet.c:939
func readCustomInletParams(i int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		c int32
	)
	_ = c
	c = 0
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	} else {
		c = project_findObject(CURVE, tok[2])
		if c < 0 {
			return error_setInpError(ERR_NAME, tok[2])
		}
	}
	InletDesigns[i].customCurve = c
	return 0
}

// inlet.c:962
func initInletStats(inlet *TInlet) {
	if inlet != nil {
		inlet.flowCapture = float64(0.0)
		inlet.backflow = float64(0.0)
		inlet.stats.flowPeriods = 0
		inlet.stats.capturePeriods = 0
		inlet.stats.backflowPeriods = 0
		inlet.stats.peakFlow = float64(0.0)
		inlet.stats.peakFlowCapture = float64(0)
		inlet.stats.avgFlowCapture = float64(0)
		inlet.stats.bypassFreq = float64(0)
	}
}

// inlet.c:985
func updateInletStats(inlet *TInlet, q float64) {
	var (
		qCapture  float64
		qBackflow float64
		qNet      float64
		qBypass   float64
		fCapture  float64
	)
	_, _, _, _, _ = qCapture, qBackflow, qNet, qBypass, fCapture
	qCapture = inlet.flowCapture
	qBackflow = inlet.backflow
	qNet = float64(qCapture - qBackflow)
	qBypass = float64(q - qNet)
	fCapture = float64(0.0)
	if (q < float64(0.001)) && (qBackflow <= float64(0.0)) {
		return
	}
	inlet.stats.flowPeriods++
	if qNet > float64(0.0) {
		inlet.stats.capturePeriods++
		fCapture = float64(qNet / q)
		if (fCapture) <= (float64(1.0)) {
			fCapture = (fCapture)
		} else {
			fCapture = (float64(1.0))
		}
		inlet.stats.avgFlowCapture = float64(inlet.stats.avgFlowCapture + fCapture)
		if qBypass > float64(0.001) {
			inlet.stats.bypassFreq++
		}
	} else {
		inlet.stats.backflowPeriods++
	}
	if q > inlet.stats.peakFlow {
		inlet.stats.peakFlow = q
		inlet.stats.peakFlowCapture = float64(fCapture * float64(100.0))
	}
}

// inlet.c:1026
func writeStreetStatsHeader() {
	report_writeLine(L0)
	report_writeLine(L188)
	report_writeLine(L189)
	report_writeLine(L188)
	report_writeLine(L0)
	R.Fprintf(Frpt.file, L190)
	if UnitSystem == US {
		R.Fprintf(Frpt.file, L191, FlowUnitWords[FlowUnits], FlowUnitWords[FlowUnits], FlowUnitWords[FlowUnits])
	} else {
		R.Fprintf(Frpt.file, L192, FlowUnitWords[FlowUnits], FlowUnitWords[FlowUnits], FlowUnitWords[FlowUnits])
	}
	R.Fprintf(Frpt.file, L193)
}

// inlet.c:1055
func writeStreetStats(link int32) {
	var (
		k         int32
		t         int32
		placement int32
		maxSpread float64
		maxDepth  float64
		maxFlow   float64
		fp        float64
		cp        float64
		afc       float64
		bpf       float64
		inlet     *TInlet
	)
	_, _, _, _, _, _, _, _, _, _, _ = k, t, placement, maxSpread, maxDepth, maxFlow, fp, cp, afc, bpf, inlet
	k = 0
	t = 0
	placement = 0
	maxSpread = 0
	maxDepth = 0
	maxFlow = 0
	fp = 0
	cp = 0
	afc = float64(0.0)
	bpf = float64(0.0)
	inlet = nil
	k = Link[link].subIndex
	t = Link[link].xsect.transect
	inlet = Link[link].inlet
	maxFlow = LinkStats[link].maxFlow
	maxDepth = LinkStats[link].maxDepth
	maxSpread = float64(xsect_getWofY(&Link[link].xsect, maxDepth) / float64(Street[t].sides))
	if (maxSpread) <= (Street[t].width) {
		maxSpread = (maxSpread)
	} else {
		maxSpread = (Street[t].width)
	}
	R.Fprintf(Frpt.file, L194, Link[link].ID)
	R.Fprintf(Frpt.file, L195, float64(maxFlow*UCF(FLOW)))
	R.Fprintf(Frpt.file, L195, float64(maxSpread*UCF(LENGTH)))
	R.Fprintf(Frpt.file, L195, float64(maxDepth*UCF(LENGTH)))
	if inlet != nil {
		R.Fprintf(Frpt.file, L196, InletDesigns[inlet.designIndex].ID)
		placement = getInletPlacement(inlet, Link[inlet.linkIndex].node2)
		if placement == ON_GRADE {
			R.Fprintf(Frpt.file, L197)
		} else {
			R.Fprintf(Frpt.file, L198)
		}
		R.Fprintf(Frpt.file, L199, inlet.numInlets)
		fp = float64(float64(inlet.stats.flowPeriods) / float64(100.0))
		if fp > float64(0.0) {
			cp = float64(float64(inlet.stats.capturePeriods) / float64(100.0))
			R.Fprintf(Frpt.file, L200, inlet.stats.peakFlowCapture)
			if cp > float64(0.0) {
				afc = float64(inlet.stats.avgFlowCapture / cp)
				bpf = float64(inlet.stats.bypassFreq / cp)
			}
			R.Fprintf(Frpt.file, L200, afc)
			R.Fprintf(Frpt.file, L200, bpf)
			R.Fprintf(Frpt.file, L200, float64(float64(inlet.stats.backflowPeriods)/fp))
			R.Fprintf(Frpt.file, L200, float64(float64(float64(float64((float64(maxFlow/float64(Street[t].sides)))*UCF(FLOW))*float64(0.01))*inlet.stats.peakFlowCapture)/float64(inlet.numInlets)))
			R.Fprintf(Frpt.file, L200, float64(float64(float64(maxFlow*UCF(FLOW))*float64(0.01))*(float64(float64(100.0)-inlet.stats.peakFlowCapture))))
		}
	}
}

// inlet.c:1129
func getInletPlacement(inlet *TInlet, j int32) int32 {
	if inlet.placement == AUTOMATIC {
		if Node[j].degree > 0 {
			return ON_GRADE
		} else {
			return ON_SAG
		}
	} else {
		return inlet.placement
	}
	return 0
}

// inlet.c:1147
func getConduitGeometry(inlet *TInlet) {
	var (
		linkIndex int32
		t         int32
		k         int32
	)
	_, _, _ = linkIndex, t, k
	linkIndex = inlet.linkIndex
	t = 0
	k = Link[linkIndex].subIndex
	SL = Conduit[k].slope
	Beta = Conduit[k].beta
	xsect = &Link[linkIndex].xsect
	if xsect.type_ == STREET_XSECT {
		t = xsect.transect
		Sx = Street[t].slope
		a = Street[t].gutterDepression
		W = Street[t].gutterWidth
		n = Street[t].roughness
		Nsides = Street[t].sides
		Tcrown = Street[t].width
		Qfactor = inlet.flowFactor
		if (inlet != nil) && (float64(inlet.localDepress*inlet.localWidth) > float64(0)) {
			a = float64(a + inlet.localDepress)
			W = inlet.localWidth
		}
		if float64(W*a) > float64(0.0) {
			Sw = float64(Sx + float64(a/W))
		} else {
			Sw = Sx
		}
	} else {
		a = float64(0.0)
		W = float64(0.0)
		n = Conduit[k].roughness
		Nsides = 1
		Sx = float64(0.01)
		Sw = Sx
	}
}

// inlet.c:1200
func getFlowSpread(Q_2 float64) float64 {
	var (
		iter int32
		f    float64
		f1   float64
		Sr   float64
		Ts1  float64
		Ts2  float64
		Tw   float64
		Qs   float64
		Eo   float64
	)
	_, _, _, _, _, _, _, _, _ = iter, f, f1, Sr, Ts1, Ts2, Tw, Qs, Eo
	iter = 0
	f = 0
	f1 = 0
	Sr = 0
	Ts1 = 0
	Ts2 = 0
	Tw = 0
	Qs = 0
	Eo = 0
	f = Qfactor
	if a == float64(0.0) {
		Ts1 = M.Pow(float64(Q_2/f), float64(0.375))
	} else {
		f1 = float64(f * M.Pow(float64((float64(a/W))/Sx), float64(1.67)))
		Tw = M.Pow(float64(Q_2/f1), float64(0.375))
		if Tw <= W {
			Ts1 = Tw
		} else {
			Sr = float64((float64(Sx + float64(a/W))) / Sx)
			iter = 1
			Ts1 = float64(M.Pow(float64(Q_2/f), float64(0.375)) - W)
			if Ts1 <= float64(0) {
				Ts1 = float64(Tw - W)
			}
			for iter < 11 {
				Eo = getEo(Sr, Ts1, W)
				Qs = float64((float64(float64(1.0) - Eo)) * Q_2)
				Ts2 = M.Pow(float64(Qs/f), float64(0.375))
				if M.Fabs(float64(Ts2-Ts1)) < float64(0.01) {
					break
				}
				Ts1 = Ts2
				iter++
			}
			Ts1 = float64(Ts2 + W)
		}
	}
	if (Ts1) <= (Tcrown) {
		return (Ts1)
	}
	return (Tcrown)
}

// inlet.c:1248
func getEo(Sr float64, Ts float64, w float64) float64 {
	var (
		x float64
	)
	_ = x
	x = 0
	x = float64(Sr / (float64(Ts / w)))
	x = float64(M.Pow((float64(float64(1.0)+x)), float64(2.67)) - float64(1.0))
	x = float64(float64(1.0) + float64(Sr/x))
	return float64(float64(1.0) / x)
}

// inlet.c:1267
func getOnGradeCapturedFlow(inlet *TInlet, q float64, d float64) float64 {
	var (
		i         int32
		linkIndex int32
		qApproach float64
		qc        float64
		qCaptured float64
		qBypassed float64
		qMax      float64
	)
	_, _, _, _, _, _, _ = i, linkIndex, qApproach, qc, qCaptured, qBypassed, qMax
	i = 0
	linkIndex = 0
	qApproach = 0
	qc = 0
	qCaptured = 0
	qBypassed = 0
	qMax = 0
	if inlet.numInlets == 0 {
		return float64(0.0)
	}
	linkIndex = inlet.linkIndex
	qApproach = q
	if qApproach < float64(0.001) {
		return float64(0.0)
	}
	getConduitGeometry(inlet)
	qApproach = float64(qApproach / (float64(Nsides)))
	qBypassed = qApproach
	qCaptured = float64(0.0)
	qMax = float64(10000000000.0)
	if inlet.flowLimit > float64(0.0) {
		qMax = inlet.flowLimit
	}
	for i = 1; i <= inlet.numInlets; i++ {
		qc = float64(getOnGradeInletCapture(inlet.designIndex, qBypassed, d) * inlet.clogFactor)
		if (qc) <= (qMax) {
			qc = (qc)
		} else {
			qc = (qMax)
		}
		if (qc) <= (qBypassed) {
			qc = (qc)
		} else {
			qc = (qBypassed)
		}
		qCaptured = float64(qCaptured + qc)
		qBypassed = float64(qBypassed - qc)
		if qBypassed < float64(0.001) {
			break
		}
	}
	return R.Set(&qCaptured, float64(qCaptured*(float64(Nsides))))
}

// inlet.c:1325
func getOnGradeInletCapture(i int32, Q_2 float64, d float64) float64 {
	var (
		Q1     float64
		Qc     float64
		Lsweep float64
		Lcurb  float64
		Lgrate float64
	)
	_, _, _, _, _ = Q1, Qc, Lsweep, Lcurb, Lgrate
	Q1 = Q_2
	Qc = float64(0.0)
	Lsweep = float64(0.0)
	Lcurb = float64(0.0)
	Lgrate = float64(0.0)
	if InletDesigns[i].type_ == DROP_CURB_INLET {
		Qc = getOnSagInletCapture(i, d)
		if (Qc) <= (Q_2) {
			return (Qc)
		}
		return (Q_2)
	}
	if InletDesigns[i].type_ == DROP_GRATE_INLET {
		Qc = getGrateInletCapture(i, Q_2)
		if (Qc) <= (Q_2) {
			return (Qc)
		}
		return (Q_2)
	}
	T = getFlowSpread(Q_2)
	if InletDesigns[i].type_ == SLOTTED_INLET {
		Qc = getCurbInletCapture(Q_2, InletDesigns[i].slottedInlet.length)
		if (Qc) <= (Q_2) {
			return (Qc)
		}
		return (Q_2)
	}
	Lcurb = InletDesigns[i].curbInlet.length
	Lgrate = InletDesigns[i].grateInlet.length
	if Lcurb > float64(0.0) {
		Lsweep = float64(Lcurb - Lgrate)
		if Lsweep > float64(0.0) {
			Qc = getCurbInletCapture(Q1, Lsweep)
			Q1 = float64(Q1 - Qc)
		}
	}
	if (Lgrate > float64(0.0)) && (Q1 > float64(0.0)) {
		if Q1 != Q_2 {
			T = getFlowSpread(Q1)
		}
		Qc = float64(Qc + (getGrateInletCapture(i, Q1)))
	}
	return Qc
}

// inlet.c:1387
func getGrateInletCapture(i int32, Q_2 float64) float64 {
	var (
		grateType int32
		Lg        float64
		Wg        float64
		A_2       float64
		Y_2       float64
		Eo        float64
		V_2       float64
		Vo        float64
		Qo        float64
		Rf        float64
		Rs        float64
	)
	_, _, _, _, _, _, _, _, _, _, _ = grateType, Lg, Wg, A_2, Y_2, Eo, V_2, Vo, Qo, Rf, Rs
	grateType = 0
	Lg = 0
	Wg = 0
	A_2 = 0
	Y_2 = 0
	Eo = 0
	V_2 = 0
	Vo = 0
	Qo = Q_2
	Rf = float64(1.0)
	Rs = float64(0.0)
	Lg = InletDesigns[i].grateInlet.length
	Wg = InletDesigns[i].grateInlet.width
	if (xsect.type_ == TRAPEZOIDAL) || (xsect.type_ == RECT_OPEN) {
		A_2 = xsect_getAofS(xsect, float64(Q_2/Beta))
		Y_2 = xsect_getYofA(xsect, A_2)
		T = xsect_getWofY(xsect, Y_2)
		Eo = float64(float64(float64(Beta*M.Pow(float64(Y_2*Wg), float64(1.67)))/M.Pow(float64(Wg+float64(float64(2)*Y_2)), float64(0.67))) / Q_2)
		if ((Wg > float64(float64(0.99)*xsect.yBot)) && (xsect.type_ == TRAPEZOIDAL)) && (xsect.sBot > float64(0.0)) {
			Wg = xsect.yBot
			Sx = float64(float64(1.0) / xsect.sBot)
		}
	} else if a == float64(0.0) {
		A_2 = float64(float64(float64(T*T)*Sx) / float64(2.0))
		Eo = getGutterFlowRatio(Wg)
		if T >= Tcrown {
			Qo = float64(Qfactor * M.Pow(Tcrown, float64(2.67)))
		}
	} else {
		if T <= W {
			A_2 = float64(float64(float64(T*T)*Sw) / float64(2.0))
		} else {
			A_2 = float64((float64(float64(float64(T*T)*Sx) + float64(a*W))) / float64(2.0))
		}
		Eo = getGutterFlowRatio(W)
		if Eo < float64(1.0) {
			if T >= Tcrown {
				Qo = float64(float64(Qfactor*M.Pow(Tcrown, float64(2.67))) / (float64(float64(1.0) - Eo)))
			}
			Eo = float64(Eo * getGutterAreaRatio(Wg, A_2))
		}
	}
	V_2 = float64(Qo / A_2)
	grateType = InletDesigns[i].grateInlet.type_
	if (grateType < 0) || (grateType == GENERIC) {
		Vo = InletDesigns[i].grateInlet.splashVeloc
	} else {
		Vo = getSplashOverVelocity(grateType, Lg)
	}
	if V_2 > Vo {
		Rf = float64(float64(1.0) - float64(float64(0.09)*(float64(V_2-Vo))))
	}
	if Eo < float64(1.0) {
		Rs = float64(float64(1.0) / (float64(float64(1.0) + (float64(float64(float64(float64(0.15)*M.Pow(V_2, float64(1.8)))/Sx) / M.Pow(Lg, float64(2.3)))))))
	}
	return float64(Q_2 * (float64(float64(Rf*Eo) + float64(Rs*(float64(float64(1.0)-Eo))))))
}

// inlet.c:1477
func getCurbInletCapture(Q_2 float64, L float64) float64 {
	var (
		Se float64
		Lt float64
		Sr float64
		Eo float64
		E  float64
	)
	_, _, _, _, _ = Se, Lt, Sr, Eo, E
	Se = Sx
	Lt = 0
	Sr = 0
	Eo = float64(0.0)
	E = float64(1.0)
	if a > float64(0.0) {
		Sr = float64(Sw / Sx)
		Eo = getEo(Sr, float64(T-W), W)
		Se = float64(Sx + float64((float64(a/W))*Eo))
	}
	Lt = float64(float64(float64(float64(0.6)*M.Pow(Q_2, float64(0.42)))*M.Pow(SL, float64(0.3))) * M.Pow(float64(float64(1.0)/(float64(n*Se))), float64(0.6)))
	if L < Lt {
		E = float64(float64(1.0) - (float64(L / Lt)))
		E = float64(float64(1) - M.Pow(E, float64(1.8)))
	}
	if (E) <= (float64(1.0)) {
		E = (E)
	} else {
		E = (float64(1.0))
	}
	if (E) >= (float64(0.0)) {
		E = (E)
	} else {
		E = (float64(0.0))
	}
	return float64(E * Q_2)
}

// inlet.c:1518
func getGutterFlowRatio(w float64) float64 {
	if T <= w {
		return float64(1.0)
	} else if a > float64(0.0) {
		return getEo(float64(Sw/Sx), float64(T-w), w)
	} else {
		return float64(float64(1.0) - M.Pow((float64(float64(1.0)-float64(w/T))), float64(2.67)))
	}
	return 0
}

// inlet.c:1535
func getGutterAreaRatio(Wg float64, A_2 float64) float64 {
	var (
		As float64
		Ag float64
	)
	_, _ = As, Ag
	As = 0
	Ag = 0
	if Wg >= W {
		return float64(1.0)
	}
	if T <= Wg {
		return float64(1.0)
	}
	if T <= W {
		return float64(Wg / T)
	}
	As = float64(float64(float64(0.5)*(float64((float64(T-W))*(float64(T-W))))) * Sx)
	Ag = float64(Wg * (float64(float64((float64(T*Sx))+a) - (float64(float64(Wg*Sw) / float64(2.0))))))
	return float64(Ag / (float64(A_2 - As)))
}

// inlet.c:1557
func getSplashOverVelocity(grateType int32, L float64) float64 {
	return float64(float64(float64(SplashCoeffs[grateType][0]+float64(SplashCoeffs[grateType][1]*L))-float64(float64(SplashCoeffs[grateType][2]*L)*L)) + float64(float64(float64(SplashCoeffs[grateType][3]*L)*L)*L))
}

// inlet.c:1574
func getOnSagCapturedFlow(inlet *TInlet, q float64, d float64) float64 {
	var (
		linkIndex   int32
		designIndex int32
		totalInlets int32
		qCaptured   float64
		qMax        float64
	)
	_, _, _, _, _ = linkIndex, designIndex, totalInlets, qCaptured, qMax
	linkIndex = 0
	designIndex = 0
	totalInlets = 0
	qCaptured = float64(0.0)
	qMax = float64(10000000000.0)
	if inlet.numInlets == 0 {
		return float64(0.0)
	}
	totalInlets = (Nsides * inlet.numInlets)
	linkIndex = inlet.linkIndex
	designIndex = inlet.designIndex
	getConduitGeometry(inlet)
	if inlet.flowLimit > float64(0.0) {
		qMax = inlet.flowLimit
	}
	qCaptured = getOnSagInletCapture(designIndex, M.Fabs(d))
	qCaptured = float64(qCaptured * inlet.clogFactor)
	if (qCaptured) <= (qMax) {
		qCaptured = (qCaptured)
	} else {
		qCaptured = (qMax)
	}
	qCaptured = float64(qCaptured * (float64(totalInlets)))
	return qCaptured
}

// inlet.c:1610
func getOnSagInletCapture(i int32, d float64) float64 {
	var (
		Lsweep float64
		Lcurb  float64
		Lgrate float64
		Qsw    float64
		Qso    float64
		Qgw    float64
		Qgo    float64
		Qcw    float64
		Qco    float64
	)
	_, _, _, _, _, _, _, _, _ = Lsweep, Lcurb, Lgrate, Qsw, Qso, Qgw, Qgo, Qcw, Qco
	Lsweep = float64(0.0)
	Lcurb = float64(0.0)
	Lgrate = float64(0.0)
	Qsw = float64(0.0)
	Qso = float64(0.0)
	Qgw = float64(0.0)
	Qgo = float64(0.0)
	Qcw = float64(0.0)
	Qco = float64(0.0)
	if InletDesigns[i].slottedInlet.length > float64(0.0) {
		return getOnSagSlottedFlow(i, d)
	}
	Lgrate = InletDesigns[i].grateInlet.length
	if Lgrate > float64(0.0) {
		findOnSagGrateFlows(i, d, &Qgw, &Qgo)
	}
	Lcurb = InletDesigns[i].curbInlet.length
	if Lcurb > float64(0.0) {
		Lsweep = float64(Lcurb - Lgrate)
		if Lsweep > float64(0.0) {
			findOnSagCurbFlows(i, d, Lsweep, &Qsw, &Qso)
		}
		if Qgo > float64(0.0) {
			findOnSagCurbFlows(i, d, Lgrate, &Qcw, &Qco)
		}
	}
	return float64(float64(float64(float64(Qgw+Qgo)+Qsw)+Qso) + Qco)
}

// inlet.c:1644
func findOnSagGrateFlows(i int32, d float64, Qw *float64, Qo *float64) {
	var (
		grateType int32
		Lg        float64
		Wg        float64
		P         float64
		Ao        float64
		di        float64
	)
	_, _, _, _, _, _ = grateType, Lg, Wg, P, Ao, di
	grateType = InletDesigns[i].grateInlet.type_
	Lg = InletDesigns[i].grateInlet.length
	Wg = InletDesigns[i].grateInlet.width
	P = 0
	Ao = 0
	di = 0
	if InletDesigns[i].type_ == DROP_GRATE_INLET {
		di = d
		P = float64(float64(2.0) * (float64(Lg + Wg)))
	} else {
		if d <= float64(Wg*Sw) {
			Wg = float64(d / Sw)
		}
		di = float64(d - float64((float64(Wg/float64(2.0)))*Sw))
		P = float64(Lg + float64(float64(2.0)*Wg))
	}
	if grateType == GENERIC {
		Ao = float64(float64(Lg*Wg) * InletDesigns[i].grateInlet.fracOpenArea)
	} else {
		Ao = float64(float64(Lg*Wg) * GrateOpeningRatios[grateType])
	}
	if d <= float64(float64(float64(1.79)*Ao)/P) {
		*Qw = float64(float64(float64(3.0)*P) * M.Pow(di, float64(1.5)))
	} else {
		*Qo = float64(float64(float64(0.67)*Ao) * M.Sqrt(float64(float64(float64(2.0)*float64(32.16))*di)))
	}
}

// inlet.c:1703
func findOnSagCurbFlows(i int32, d float64, L float64, Qw *float64, Qo *float64) {
	var (
		throatAngle int32
		h           float64
		Qweir       float64
		Qorif       float64
		P           float64
		dweir       float64
		dorif       float64
		r           float64
	)
	_, _, _, _, _, _, _, _ = throatAngle, h, Qweir, Qorif, P, dweir, dorif, r
	throatAngle = InletDesigns[i].curbInlet.throatAngle
	h = InletDesigns[i].curbInlet.height
	Qweir = 0
	Qorif = 0
	P = 0
	dweir = 0
	dorif = 0
	r = 0
	if L <= float64(0.0) {
		return
	}
	if InletDesigns[i].type_ == DROP_CURB_INLET {
		L = float64(L * float64(4.0))
	}
	dorif = float64(float64(1.4) * h)
	if d > dorif {
		*Qo = getCurbOrificeFlow(d, h, L, throatAngle)
		return
	}
	if (a == float64(0.0)) || (L > float64(12.0)) {
		dweir = h
		if d < dweir {
			*Qw = float64(float64(float64(3.0)*L) * M.Pow(d, float64(1.5)))
			return
		} else {
			Qweir = float64(float64(float64(3.0)*L) * M.Pow(dweir, float64(1.5)))
		}
	} else {
		P = float64(L + float64(float64(1.8)*W))
		dweir = float64(h + a)
		if d < dweir {
			*Qw = float64(float64(float64(2.3)*P) * M.Pow(d, float64(1.5)))
			return
		} else {
			Qweir = float64(float64(float64(2.3)*P) * M.Pow(dweir, float64(1.5)))
		}
	}
	Qorif = getCurbOrificeFlow(dorif, h, L, throatAngle)
	r = float64((float64(d - dweir)) / (float64(dorif - dweir)))
	*Qw = float64((float64(float64(1.0) - r)) * Qweir)
	*Qo = float64(r * Qorif)
}

// inlet.c:1764
func getCurbOrificeFlow(di float64, h float64, L float64, throatAngle int32) float64 {
	var (
		d float64
	)
	_ = d
	d = di
	if throatAngle == HORIZONTAL_THROAT {
		d = float64(di - float64(h/float64(2.0)))
	} else if throatAngle == INCLINED_THROAT {
		d = float64(di - float64((float64(h/float64(2.0)))*float64(0.7071)))
	}
	return float64(float64(float64(float64(0.67)*h)*L) * M.Sqrt(float64(float64(float64(2.0)*float64(32.16))*d)))
}

// inlet.c:1785
func getOnSagSlottedFlow(i int32, d float64) float64 {
	var (
		L float64
		w float64
	)
	_, _ = L, w
	L = InletDesigns[i].slottedInlet.length
	w = InletDesigns[i].slottedInlet.width
	if d <= float64(float64(2.587)*w) {
		return float64(float64(float64(2.48)*L) * M.Pow(d, float64(1.5)))
	} else {
		return float64(float64(float64(float64(0.8)*L)*w) * M.Sqrt(float64(float64(64.32)*d)))
	}
	return 0
}

// inlet.c:1805
func getBackflowRatios() {
	var (
		inlet      *TInlet
		area       float64
		f          float64
		n_2        int32
		inletNodes []TInletNode
	)
	_, _, _, _, _ = inlet, area, f, n_2, inletNodes
	inlet = nil
	area = 0
	f = 0
	n_2 = 0
	inletNodes = make([]TInletNode, int(uint64(Nobjects[NODE])))
	if inletNodes == nil {
		return
	}
	for inlet = FirstInlet; inlet != nil; inlet = inlet.nextInlet {
		n_2 = inlet.nodeIndex
		inletNodes[n_2].numInletLinks++
		area = getInletArea(inlet)
		if area > float64(0.0) {
			inletNodes[n_2].numStdInletLinks++
			inletNodes[n_2].totalInletArea = float64(inletNodes[n_2].totalInletArea + area)
		} else {
			inletNodes[n_2].numCustomInlets += inlet.numInlets
		}
	}
	for inlet = FirstInlet; inlet != nil; inlet = inlet.nextInlet {
		n_2 = inlet.nodeIndex
		f = float64(float64(inletNodes[n_2].numStdInletLinks) / float64(inletNodes[n_2].numInletLinks))
		area = getInletArea(inlet)
		if area == float64(0.0) {
			inlet.backflowRatio = float64(float64(float64(inlet.numInlets)/float64(inletNodes[n_2].numCustomInlets)) * (float64(float64(1.0) - f)))
		} else {
			inlet.backflowRatio = float64(float64(area/inletNodes[n_2].totalInletArea) * f)
		}
	}
	_ = inletNodes
}

// inlet.c:1870
func getInletArea(inlet *TInlet) float64 {
	var (
		area       float64
		curbLength float64
		i          int32
		grateType  int32
	)
	_, _, _, _ = area, curbLength, i, grateType
	area = float64(0.0)
	curbLength = 0
	i = inlet.designIndex
	grateType = InletDesigns[i].grateInlet.type_
	if InletDesigns[i].grateInlet.length > float64(0.0) {
		area = float64(InletDesigns[i].grateInlet.length * InletDesigns[i].grateInlet.width)
		if grateType == GENERIC {
			area = float64(area * InletDesigns[i].grateInlet.fracOpenArea)
		} else {
			area = float64(area * GrateOpeningRatios[grateType])
		}
	}
	curbLength = float64(InletDesigns[i].curbInlet.length - InletDesigns[i].grateInlet.length)
	if curbLength > float64(0.0) {
		area = float64(area + (float64(curbLength * InletDesigns[i].curbInlet.height)))
	}
	if InletDesigns[i].slottedInlet.length > float64(0.0) {
		area = float64(InletDesigns[i].slottedInlet.length * InletDesigns[i].slottedInlet.width)
	}
	return float64(float64(area*float64(inlet.numInlets)) * inlet.clogFactor)
}

// inlet.c:1902
func getCustomCapturedFlow(inlet *TInlet, q float64, d float64) float64 {
	var (
		i          int32
		j          int32
		sides      int32
		c          int32
		qApproach  float64
		qBypassed  float64
		qCaptured  float64
		qIncrement float64
		qMax       float64
	)
	_, _, _, _, _, _, _, _, _ = i, j, sides, c, qApproach, qBypassed, qCaptured, qIncrement, qMax
	i = inlet.designIndex
	j = 0
	sides = 1
	c = 0
	qApproach = 0
	qBypassed = 0
	qCaptured = 0
	qIncrement = 0
	qMax = float64(10000000000.0)
	if inlet.numInlets == 0 {
		return float64(0.0)
	}
	qMax = float64(10000000000.0)
	if inlet.flowLimit > float64(0.0) {
		qMax = inlet.flowLimit
	}
	xsect = &Link[inlet.linkIndex].xsect
	if xsect.type_ == STREET_XSECT {
		sides = Street[xsect.transect].sides
	}
	qApproach = float64(q / float64(sides))
	qBypassed = qApproach
	qCaptured = float64(0.0)
	c = InletDesigns[i].customCurve
	if c >= 0 {
		if Curve[c].curveType == DIVERSION_CURVE {
			for j = 1; j <= inlet.numInlets; j++ {
				qIncrement = float64(float64(inlet.clogFactor*table_lookupEx(&Curve[c], float64(qBypassed*UCF(FLOW)))) / UCF(FLOW))
				if (qIncrement) <= (qMax) {
					qIncrement = (qIncrement)
				} else {
					qIncrement = (qMax)
				}
				if (qIncrement) <= (qBypassed) {
					qIncrement = (qIncrement)
				} else {
					qIncrement = (qBypassed)
				}
				qCaptured = float64(qCaptured + qIncrement)
				qBypassed = float64(qBypassed - qIncrement)
				if qBypassed < float64(0.001) {
					break
				}
			}
		} else if Curve[c].curveType == RATING_CURVE {
			qCaptured = float64(float64(float64(float64(inlet.numInlets)*inlet.clogFactor)*table_lookupEx(&Curve[c], float64(d*UCF(LENGTH)))) / UCF(FLOW))
		}
		qCaptured = float64(qCaptured * (float64(sides)))
	}
	return qCaptured
}
