// Generated from EPA SWMM 5.2.4 (gwater.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var GWTOL float64 = float64(0.0001)
var XTOL float64 = float64(0.001)
var GWVarWords = [12]R.CPtr{L707, L708, L709, L710, L711, L712, L713, L714, L715, L716, L717, (R.CPtr{})}
var Area float64
var gwater_Infil float64
var MaxEvap float64
var AvailEvap float64
var UpperEvap float64
var LowerEvap float64
var UpperPerc float64
var LowerLoss float64
var GWFlow float64
var MaxUpperPerc float64
var MaxGWFlowPos float64
var MaxGWFlowNeg float64
var FracPerv float64
var TotalDepth float64
var Theta float64
var HydCon float64
var Hgw float64
var Hstar float64
var Hsw float64
var gwater_Tstep float64
var A TAquifer
var GW *TGroundwater
var LatFlowExpr *MathExpr
var DeepFlowExpr *MathExpr

// gwater.c:119
func gwater_readAquiferParams(j int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		i  int32
		p  int32
		x  [12]float64
		id R.CPtr
	)
	_, _, _, _ = i, p, x, id
	i = 0
	p = 0
	x = ([12]float64{})
	id = (R.CPtr{})
	if ntoks < 13 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(AQUIFER, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	for i = 0; i < 11; i++ {
		x[i] = float64(0.0)
	}
	for i = 1; i < 13; i++ {
		if !(getDouble(tok[i], &x[(i-1)]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[i])
		}
	}
	p = (-1)
	if ntoks > 13 {
		p = project_findObject(TIMEPATTERN, tok[13])
		if p < 0 {
			return error_setInpError(ERR_NAME, tok[13])
		}
	}
	Aquifer[j].ID = id
	Aquifer[j].porosity = x[0]
	Aquifer[j].wiltingPoint = x[1]
	Aquifer[j].fieldCapacity = x[2]
	Aquifer[j].conductivity = float64(x[3] / UCF(RAINFALL))
	Aquifer[j].conductSlope = x[4]
	Aquifer[j].tensionSlope = float64(x[5] / UCF(LENGTH))
	Aquifer[j].upperEvapFrac = x[6]
	Aquifer[j].lowerEvapDepth = float64(x[7] / UCF(LENGTH))
	Aquifer[j].lowerLossCoeff = float64(x[8] / UCF(RAINFALL))
	Aquifer[j].bottomElev = float64(x[9] / UCF(LENGTH))
	Aquifer[j].waterTableElev = float64(x[10] / UCF(LENGTH))
	Aquifer[j].upperMoisture = x[11]
	Aquifer[j].upperEvapPat = p
	return 0
}

// gwater.c:179
func gwater_readGroundwaterParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		i   int32
		j   int32
		k   int32
		m   int32
		n_2 int32
		x   [11]float64
		gw  *TGroundwater
	)
	_, _, _, _, _, _, _ = i, j, k, m, n_2, x, gw
	i = 0
	j = 0
	k = 0
	m = 0
	n_2 = 0
	x = ([11]float64{})
	gw = nil
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(SUBCATCH, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	if ntoks < 11 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	k = project_findObject(AQUIFER, tok[1])
	if k < 0 {
		return error_setInpError(ERR_NAME, tok[1])
	}
	n_2 = project_findObject(NODE, tok[2])
	if n_2 < 0 {
		return error_setInpError(ERR_NAME, tok[2])
	}
	for i = 0; i < 7; i++ {
		if !(getDouble(tok[(i+3)], &x[i]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[(i+3)])
		}
	}
	for i = 7; i < 11; i++ {
		x[i] = (-(float64(10000000000.0)))
		m = (i + 3)
		if (ntoks > m) && (tok[m].G(0) != 42) {
			if !(getDouble(tok[m], &x[i]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[m])
			}
			if i < 10 {
				x[i] = float64(x[i] / (UCF(LENGTH)))
			}
		}
	}
	if !(Subcatch[j].groundwater != nil) {
		gw = new(TGroundwater)
		if !(gw != nil) {
			return error_setInpError(ERR_MEMORY, L0)
		}
		Subcatch[j].groundwater = gw
	} else {
		gw = Subcatch[j].groundwater
	}
	gw.aquifer = k
	gw.node = n_2
	gw.surfElev = float64(x[0] / UCF(LENGTH))
	gw.a1 = x[1]
	gw.b1 = x[2]
	gw.a2 = x[3]
	gw.b2 = x[4]
	gw.a3 = x[5]
	gw.fixedDepth = float64(x[6] / UCF(LENGTH))
	gw.nodeElev = x[7]
	gw.bottomElev = x[8]
	gw.waterTableElev = x[9]
	gw.upperMoisture = x[10]
	return 0
}

// gwater.c:258
func gwater_readFlowExpression(tok []R.CPtr, ntoks int32) int32 {
	var (
		i       int32
		j       int32
		k       int32
		exprStr R.CPtr
		expr    *MathExpr
	)
	_, _, _, _, _ = i, j, k, exprStr, expr
	i = 0
	j = 0
	k = 0
	exprStr = R.Buf(1025)
	expr = nil
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(SUBCATCH, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	k = 1
	if match(tok[1], L158) != 0 {
		k = 1
	} else if match(tok[1], L159) != 0 {
		k = 2
	} else {
		return error_setInpError(ERR_KEYWORD, tok[1])
	}
	sstrncpy(exprStr, tok[2], 1024)
	for i = 3; i < ntoks; i++ {
		sstrcat(exprStr, L33, uint64((1024 + 1)))
		sstrcat(exprStr, tok[i], uint64((1024 + 1)))
	}
	if k == 1 {
		mathexpr_delete(Subcatch[j].gwLatFlowExpr)
	} else {
		mathexpr_delete(Subcatch[j].gwDeepFlowExpr)
	}
	expr = mathexpr_create(exprStr, gwater_getVariableIndex)
	if expr == nil {
		return error_setInpError(ERR_MATH_EXPR, L0)
	}
	if k == 1 {
		Subcatch[j].gwLatFlowExpr = expr
	} else {
		Subcatch[j].gwDeepFlowExpr = expr
	}
	return 0
}

// gwater.c:315
func gwater_deleteFlowExpression(j int32) {
	mathexpr_delete(Subcatch[j].gwLatFlowExpr)
	mathexpr_delete(Subcatch[j].gwDeepFlowExpr)
}

// gwater.c:328
func gwater_validateAquifer(j int32) {
	var (
		p int32
	)
	_ = p
	p = 0
	if ((((((((((Aquifer[j].porosity <= float64(0.0)) || (Aquifer[j].fieldCapacity >= Aquifer[j].porosity)) || (Aquifer[j].wiltingPoint >= Aquifer[j].fieldCapacity)) || (Aquifer[j].conductivity <= float64(0.0))) || (Aquifer[j].conductSlope < float64(0.0))) || (Aquifer[j].tensionSlope < float64(0.0))) || (Aquifer[j].upperEvapFrac < float64(0.0))) || (Aquifer[j].lowerEvapDepth < float64(0.0))) || (Aquifer[j].waterTableElev < Aquifer[j].bottomElev)) || (Aquifer[j].upperMoisture > Aquifer[j].porosity)) || (Aquifer[j].upperMoisture < Aquifer[j].wiltingPoint) {
		report_writeErrorMsg(ERR_AQUIFER_PARAMS, Aquifer[j].ID)
	}
	p = Aquifer[j].upperEvapPat
	if (p >= 0) && (Pattern[p].type_ != MONTHLY_PATTERN) {
		report_writeErrorMsg(ERR_AQUIFER_PARAMS, Aquifer[j].ID)
	}
}

// gwater.c:359
func gwater_validate(j int32) {
	var (
		a_2 TAquifer
		gw  *TGroundwater
	)
	_, _ = a_2, gw
	a_2 = (TAquifer{})
	gw = nil
	gw = Subcatch[j].groundwater
	if gw != nil {
		a_2 = Aquifer[gw.aquifer]
		if gw.bottomElev == (-(float64(10000000000.0))) {
			gw.bottomElev = a_2.bottomElev
		}
		if gw.waterTableElev == (-(float64(10000000000.0))) {
			gw.waterTableElev = a_2.waterTableElev
		}
		if gw.upperMoisture == (-(float64(10000000000.0))) {
			gw.upperMoisture = a_2.upperMoisture
		}
		if gw.surfElev < gw.waterTableElev {
			report_writeErrorMsg(ERR_GROUND_ELEV, Subcatch[j].ID)
		}
	}
}

// gwater.c:382
func gwater_initState(j int32) {
	var (
		a_2 TAquifer
		gw  *TGroundwater
	)
	_, _ = a_2, gw
	a_2 = (TAquifer{})
	gw = nil
	gw = Subcatch[j].groundwater
	if gw != nil {
		a_2 = Aquifer[gw.aquifer]
		gw.theta = gw.upperMoisture
		if gw.theta >= a_2.porosity {
			gw.theta = float64(a_2.porosity - XTOL)
		}
		gw.lowerDepth = float64(gw.waterTableElev - gw.bottomElev)
		if gw.lowerDepth >= float64(gw.surfElev-gw.bottomElev) {
			gw.lowerDepth = float64(float64(gw.surfElev-gw.bottomElev) - XTOL)
		}
		gw.oldFlow = float64(0.0)
		gw.newFlow = float64(0.0)
		gw.evapLoss = float64(0.0)
		gw.maxInfilVol = float64(float64((float64(gw.surfElev-gw.waterTableElev))*(float64(a_2.porosity-gw.theta))) / subcatch_getFracPerv(j))
	}
}

// gwater.c:425
func gwater_getState(j int32, x []float64) {
	var (
		gw *TGroundwater
	)
	_ = gw
	gw = Subcatch[j].groundwater
	x[0] = gw.theta
	x[1] = float64(gw.bottomElev + gw.lowerDepth)
	x[2] = gw.newFlow
	x[3] = gw.maxInfilVol
}

// gwater.c:441
func gwater_setState(j int32, x []float64) {
	var (
		gw *TGroundwater
	)
	_ = gw
	gw = Subcatch[j].groundwater
	if gw == nil {
		return
	}
	gw.theta = x[0]
	gw.lowerDepth = float64(x[1] - gw.bottomElev)
	gw.oldFlow = x[2]
	if x[3] != (-(float64(10000000000.0))) {
		gw.maxInfilVol = x[3]
	}
}

// gwater.c:458
func gwater_getVolume(j int32) float64 {
	var (
		a_2        TAquifer
		gw         *TGroundwater
		upperDepth float64
	)
	_, _, _ = a_2, gw, upperDepth
	a_2 = (TAquifer{})
	gw = nil
	upperDepth = 0
	gw = Subcatch[j].groundwater
	if gw == nil {
		return float64(0.0)
	}
	a_2 = Aquifer[gw.aquifer]
	upperDepth = float64(float64(gw.surfElev-gw.bottomElev) - gw.lowerDepth)
	return float64((float64(upperDepth * gw.theta)) + (float64(gw.lowerDepth * a_2.porosity)))
}

// gwater.c:477
func gwater_getGroundwater(j int32, evap float64, infil float64, tStep float64) {
	var (
		n_2      int32
		x        [2]float64
		vUpper   float64
		nodeFlow float64
	)
	_, _, _, _ = n_2, x, vUpper, nodeFlow
	n_2 = 0
	x = ([2]float64{})
	vUpper = 0
	nodeFlow = 0
	GW = Subcatch[j].groundwater
	if GW == nil {
		return
	}
	LatFlowExpr = Subcatch[j].gwLatFlowExpr
	DeepFlowExpr = Subcatch[j].gwDeepFlowExpr
	A = Aquifer[GW.aquifer]
	FracPerv = subcatch_getFracPerv(j)
	if FracPerv <= float64(0.0) {
		return
	}
	Area = Subcatch[j].area
	infil = float64(float64(infil/Area) / tStep)
	gwater_Infil = infil
	gwater_Tstep = tStep
	evap = float64(float64(evap/Area) / tStep)
	MaxEvap = float64(Evap.rate * FracPerv)
	if (float64(MaxEvap - evap)) >= (float64(0.0)) {
		AvailEvap = (float64(MaxEvap - evap))
	} else {
		AvailEvap = (float64(0.0))
	}
	TotalDepth = float64(GW.surfElev - GW.bottomElev)
	if TotalDepth <= float64(0.0) {
		return
	}
	n_2 = GW.node
	if GW.nodeElev != (-(float64(10000000000.0))) {
		Hstar = float64(GW.nodeElev - GW.bottomElev)
	} else {
		Hstar = float64(Node[n_2].invertElev - GW.bottomElev)
	}
	if GW.fixedDepth > float64(0.0) {
		Hsw = float64(float64(GW.fixedDepth+Node[n_2].invertElev) - GW.bottomElev)
	} else {
		Hsw = float64(float64(Node[n_2].newDepth+Node[n_2].invertElev) - GW.bottomElev)
	}
	x[THETA] = GW.theta
	x[LOWERDEPTH] = GW.lowerDepth
	vUpper = float64((float64(TotalDepth - x[LOWERDEPTH])) * (float64(x[THETA] - A.fieldCapacity)))
	if (float64(0.0)) >= (vUpper) {
		vUpper = (float64(0.0))
	} else {
		vUpper = (vUpper)
	}
	MaxUpperPerc = float64(vUpper / tStep)
	MaxGWFlowPos = float64(float64(x[LOWERDEPTH]*A.porosity) / tStep)
	MaxGWFlowNeg = float64(float64((float64(TotalDepth-x[LOWERDEPTH]))*(float64(A.porosity-x[THETA]))) / tStep)
	nodeFlow = float64((float64(Node[n_2].inflow + float64(Node[n_2].newVolume/tStep))) / Area)
	MaxGWFlowNeg = (-(R.If[float64](((MaxGWFlowNeg) <= (nodeFlow)), (MaxGWFlowNeg), (nodeFlow))))
	odesolve_integrate(x[:], 2, float64(0), tStep, GWTOL, tStep, getDxDt)
	if (x[THETA]) >= (A.wiltingPoint) {
		x[THETA] = (x[THETA])
	} else {
		x[THETA] = (A.wiltingPoint)
	}
	if x[THETA] >= A.porosity {
		x[THETA] = float64(A.porosity - XTOL)
		x[LOWERDEPTH] = float64(TotalDepth - XTOL)
	}
	if (x[LOWERDEPTH]) >= (float64(0.0)) {
		x[LOWERDEPTH] = (x[LOWERDEPTH])
	} else {
		x[LOWERDEPTH] = (float64(0.0))
	}
	if x[LOWERDEPTH] >= TotalDepth {
		x[LOWERDEPTH] = float64(TotalDepth - XTOL)
	}
	GW.theta = x[THETA]
	GW.lowerDepth = x[LOWERDEPTH]
	getFluxes(GW.theta, GW.lowerDepth)
	GW.oldFlow = GW.newFlow
	GW.newFlow = GWFlow
	GW.evapLoss = float64(UpperEvap + LowerEvap)
	GW.maxInfilVol = float64(float64((float64(TotalDepth-x[LOWERDEPTH]))*(float64(A.porosity-x[THETA]))) / FracPerv)
	updateMassBal(Area, tStep)
	stats_updateGwaterStats(j, infil, GW.evapLoss, GWFlow, LowerLoss, GW.theta, float64(GW.lowerDepth+GW.bottomElev), tStep)
}

// gwater.c:605
func updateMassBal(area float64, tStep float64) {
	var (
		vInfil     float64
		vUpperEvap float64
		vLowerEvap float64
		vLowerPerc float64
		vGwater    float64
		ft2sec     float64
	)
	_, _, _, _, _, _ = vInfil, vUpperEvap, vLowerEvap, vLowerPerc, vGwater, ft2sec
	vInfil = 0
	vUpperEvap = 0
	vLowerEvap = 0
	vLowerPerc = 0
	vGwater = 0
	ft2sec = float64(area * tStep)
	vInfil = float64(gwater_Infil * ft2sec)
	vUpperEvap = float64(UpperEvap * ft2sec)
	vLowerEvap = float64(LowerEvap * ft2sec)
	vLowerPerc = float64(LowerLoss * ft2sec)
	vGwater = float64(float64(float64(0.5)*(float64(GW.oldFlow+GW.newFlow))) * ft2sec)
	massbal_updateGwaterTotals(vInfil, vUpperEvap, vLowerEvap, vLowerPerc, vGwater)
}

// gwater.c:631
func getFluxes(theta float64, lowerDepth float64) {
	var (
		upperDepth float64
	)
	_ = upperDepth
	upperDepth = 0
	if (lowerDepth) >= (float64(0.0)) {
		lowerDepth = (lowerDepth)
	} else {
		lowerDepth = (float64(0.0))
	}
	if (lowerDepth) <= (TotalDepth) {
		lowerDepth = (lowerDepth)
	} else {
		lowerDepth = (TotalDepth)
	}
	upperDepth = float64(TotalDepth - lowerDepth)
	Hgw = lowerDepth
	Theta = theta
	gwater_getEvapRates(theta, upperDepth)
	UpperPerc = getUpperPerc(theta, upperDepth)
	if (UpperPerc) <= (MaxUpperPerc) {
		UpperPerc = (UpperPerc)
	} else {
		UpperPerc = (MaxUpperPerc)
	}
	if DeepFlowExpr != nil {
		LowerLoss = float64(mathexpr_eval(DeepFlowExpr, gwater_getVariableValue) / UCF(RAINFALL))
	} else {
		LowerLoss = float64(float64(A.lowerLossCoeff*lowerDepth) / TotalDepth)
	}
	if (LowerLoss) <= (float64(lowerDepth / gwater_Tstep)) {
		LowerLoss = (LowerLoss)
	} else {
		LowerLoss = (float64(lowerDepth / gwater_Tstep))
	}
	GWFlow = getGWFlow(lowerDepth)
	if LatFlowExpr != nil {
		GWFlow = float64(GWFlow + (float64(mathexpr_eval(LatFlowExpr, gwater_getVariableValue) / UCF(GWFLOW))))
	}
	if GWFlow >= float64(0.0) {
		if (GWFlow) <= (MaxGWFlowPos) {
			GWFlow = (GWFlow)
		} else {
			GWFlow = (MaxGWFlowPos)
		}
	} else {
		if (GWFlow) >= (MaxGWFlowNeg) {
			GWFlow = (GWFlow)
		} else {
			GWFlow = (MaxGWFlowNeg)
		}
	}
}

// gwater.c:677
func getDxDt(t float64, x []float64, dxdt []float64) {
	var (
		qUpper float64
		qLower float64
		denom  float64
	)
	_, _, _ = qUpper, qLower, denom
	qUpper = 0
	qLower = 0
	denom = 0
	getFluxes(x[THETA], x[LOWERDEPTH])
	qUpper = float64(float64(gwater_Infil-UpperEvap) - UpperPerc)
	qLower = float64(float64(float64(UpperPerc-LowerLoss)-LowerEvap) - GWFlow)
	denom = float64(TotalDepth - x[LOWERDEPTH])
	if denom > float64(0.0) {
		dxdt[THETA] = float64(qUpper / denom)
	} else {
		dxdt[THETA] = float64(0.0)
	}
	denom = float64(A.porosity - x[THETA])
	if denom > float64(0.0) {
		dxdt[LOWERDEPTH] = float64(qLower / denom)
	} else {
		dxdt[LOWERDEPTH] = float64(0.0)
	}
}

// gwater.c:713
func gwater_getEvapRates(theta float64, upperDepth float64) {
	var (
		p         int32
		month     int32
		f         float64
		lowerFrac float64
		upperFrac float64
	)
	_, _, _, _, _ = p, month, f, lowerFrac, upperFrac
	p = 0
	month = 0
	f = 0
	lowerFrac = 0
	upperFrac = 0
	UpperEvap = float64(0.0)
	LowerEvap = float64(0.0)
	if gwater_Infil > float64(0.0) {
		return
	}
	upperFrac = A.upperEvapFrac
	f = float64(1.0)
	p = A.upperEvapPat
	if p >= 0 {
		month = datetime_monthOfYear(getDateTime(NewRunoffTime))
		f = Pattern[p].factor[(month - 1)]
	}
	upperFrac = float64(upperFrac * f)
	if theta > A.wiltingPoint {
		UpperEvap = float64(upperFrac * MaxEvap)
		if (UpperEvap) <= (AvailEvap) {
			UpperEvap = (UpperEvap)
		} else {
			UpperEvap = (AvailEvap)
		}
	}
	if A.lowerEvapDepth > float64(0.0) {
		lowerFrac = float64((float64(A.lowerEvapDepth - upperDepth)) / A.lowerEvapDepth)
		if (float64(0.0)) >= (lowerFrac) {
			lowerFrac = (float64(0.0))
		} else {
			lowerFrac = (lowerFrac)
		}
		if (lowerFrac) <= (float64(1.0)) {
			lowerFrac = (lowerFrac)
		} else {
			lowerFrac = (float64(1.0))
		}
		LowerEvap = float64(float64(lowerFrac*(float64(float64(1.0)-upperFrac))) * MaxEvap)
		if (LowerEvap) <= (float64(AvailEvap - UpperEvap)) {
			LowerEvap = (LowerEvap)
		} else {
			LowerEvap = (float64(AvailEvap - UpperEvap))
		}
	}
}

// gwater.c:769
func getUpperPerc(theta float64, upperDepth float64) float64 {
	var (
		delta  float64
		dhdz   float64
		hydcon float64
	)
	_, _, _ = delta, dhdz, hydcon
	delta = 0
	dhdz = 0
	hydcon = 0
	if (upperDepth <= float64(0.0)) || (theta <= A.fieldCapacity) {
		return float64(0.0)
	}
	delta = float64(theta - A.porosity)
	hydcon = float64(A.conductivity * M.Exp(float64(delta*A.conductSlope)))
	delta = float64(theta - A.fieldCapacity)
	dhdz = float64(float64(1.0) + float64(float64(float64(A.tensionSlope*float64(2.0))*delta)/upperDepth))
	HydCon = hydcon
	return float64(hydcon * dhdz)
}

// gwater.c:799
func getGWFlow(lowerDepth float64) float64 {
	var (
		q  float64
		t1 float64
		t2 float64
		t3 float64
	)
	_, _, _, _ = q, t1, t2, t3
	q = 0
	t1 = 0
	t2 = 0
	t3 = 0
	if lowerDepth <= Hstar {
		return float64(0.0)
	}
	if GW.b1 == float64(0.0) {
		t1 = GW.a1
	} else {
		t1 = float64(GW.a1 * M.Pow(float64((float64(lowerDepth-Hstar))*UCF(LENGTH)), GW.b1))
	}
	if GW.b2 == float64(0.0) {
		t2 = GW.a2
	} else if Hsw > Hstar {
		t2 = float64(GW.a2 * M.Pow(float64((float64(Hsw-Hstar))*UCF(LENGTH)), GW.b2))
	} else {
		t2 = float64(0.0)
	}
	t3 = float64(float64(float64(float64(GW.a3*lowerDepth)*Hsw)*UCF(LENGTH)) * UCF(LENGTH))
	q = float64((float64(float64(t1-t2) + t3)) / UCF(GWFLOW))
	if (q < float64(0.0)) && (GW.a3 != float64(0.0)) {
		q = float64(0.0)
	}
	return q
}

// gwater.c:834
func gwater_getVariableIndex(s R.CPtr) int32 {
	var (
		k int32
	)
	_ = k
	k = 0
	k = findmatch(s, GWVarWords[:])
	if k >= 0 {
		return k
	}
	return (-1)
}

// gwater.c:850
func gwater_getVariableValue(varIndex int32) float64 {
	switch varIndex {
	case gwvHGW:
		return float64(Hgw * UCF(LENGTH))
	case gwvHSW:
		return float64(Hsw * UCF(LENGTH))
	case gwvHCB:
		return float64(Hstar * UCF(LENGTH))
	case gwvHGS:
		return float64(TotalDepth * UCF(LENGTH))
	case gwvKS:
		return float64(A.conductivity * UCF(RAINFALL))
	case gwvK:
		return float64(HydCon * UCF(RAINFALL))
	case gwvTHETA:
		return Theta
	case gwvPHI:
		return A.porosity
	case gwvFI:
		return float64(gwater_Infil * UCF(RAINFALL))
	case gwvFU:
		return float64(UpperPerc * UCF(RAINFALL))
	case gwvA:
		return float64(Area * UCF(LANDAREA))
	default:
		return float64(0.0)
	}
	return 0
}
