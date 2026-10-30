// Generated from EPA SWMM 5.2.4 (node.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

// node.c:102
func node_readParams(j int32, type_2 int32, k int32, tok []R.CPtr, ntoks int32) int32 {
	switch type_2 {
	case JUNCTION:
		return junc_readParams(j, k, tok, ntoks)
	case OUTFALL:
		return outfall_readParams(j, k, tok, ntoks)
	case STORAGE:
		return storage_readParams(j, k, tok, ntoks)
	case DIVIDER:
		return divider_readParams(j, k, tok, ntoks)
	default:
		return 0
	}
	return 0
}

// node.c:125
func node_setParams(j int32, type_2 int32, k int32, x []float64) {
	Node[j].type_ = type_2
	Node[j].subIndex = k
	Node[j].invertElev = float64(x[0] / UCF(LENGTH))
	Node[j].crownElev = Node[j].invertElev
	Node[j].initDepth = float64(0.0)
	Node[j].newVolume = float64(0.0)
	Node[j].fullVolume = float64(0.0)
	Node[j].fullDepth = float64(0.0)
	Node[j].surDepth = float64(0.0)
	Node[j].pondedArea = float64(0.0)
	Node[j].degree = 0
	Node[j].inlet = NO_INLET
	switch type_2 {
	case JUNCTION:
		Node[j].fullDepth = float64(x[1] / UCF(LENGTH))
		Node[j].initDepth = float64(x[2] / UCF(LENGTH))
		Node[j].surDepth = float64(x[3] / UCF(LENGTH))
		Node[j].pondedArea = float64(x[4] / (float64(UCF(LENGTH) * UCF(LENGTH))))
		break
	case OUTFALL:
		Outfall[k].type_ = R.D2I(x[1])
		Outfall[k].fixedStage = float64(x[2] / UCF(LENGTH))
		Outfall[k].tideCurve = R.D2I(x[3])
		Outfall[k].stageSeries = R.D2I(x[4])
		Outfall[k].hasFlapGate = int32(int8(R.D2I(x[5])))
		Outfall[k].routeTo = R.D2I(x[6])
		Outfall[k].wRouted = nil
		if Outfall[k].routeTo >= 0 {
			Outfall[k].wRouted = make([]float64, int(uint64(Nobjects[POLLUT])))
		}
		break
	case STORAGE:
		Node[j].fullDepth = float64(x[1] / UCF(LENGTH))
		Node[j].initDepth = float64(x[2] / UCF(LENGTH))
		Storage[k].shape = R.D2I(x[3])
		Storage[k].a1 = x[4]
		Storage[k].a2 = x[5]
		Storage[k].a0 = x[6]
		Storage[k].aCurve = R.D2I(x[7])
		Node[j].surDepth = float64(x[8] / UCF(LENGTH))
		Storage[k].fEvap = x[9]
		break
	case DIVIDER:
		Divider[k].link = R.D2I(x[1])
		Divider[k].type_ = R.D2I(x[2])
		Divider[k].flowCurve = R.D2I(x[3])
		Divider[k].qMin = float64(x[4] / UCF(FLOW))
		Divider[k].dhMax = x[5]
		Divider[k].cWeir = x[6]
		Node[j].fullDepth = float64(x[7] / UCF(LENGTH))
		Node[j].initDepth = float64(x[8] / UCF(LENGTH))
		Node[j].surDepth = float64(x[9] / UCF(LENGTH))
		Node[j].pondedArea = float64(x[10] / (float64(UCF(LENGTH) * UCF(LENGTH))))
		break
	}
}

// node.c:200
func node_validate(j int32) {
	var (
		inflow *TDwfInflow
	)
	_ = inflow
	inflow = nil
	if (Node[j].fullDepth > Node[j].oldDepth) && (Node[j].oldDepth > float64(0.0)) {
		report_writeWarningMsg(L317, Node[j].ID)
	}
	if Node[j].initDepth > float64(Node[j].fullDepth+Node[j].surDepth) {
		report_writeErrorMsg(ERR_NODE_DEPTH, Node[j].ID)
	}
	if Node[j].type_ == STORAGE {
		if node_getVolume(j, Node[j].fullDepth) < float64(0.0) {
			report_writeErrorMsg(ERR_STORAGE_VOLUME, Node[j].ID)
		}
	}
	if Node[j].type_ == DIVIDER {
		divider_validate(j)
	}
	inflow = Node[j].dwfInflow
	for inflow != nil {
		inflow_initDwfInflow(inflow)
		inflow = inflow.next
	}
}

// node.c:237
func node_initState(j int32) {
	var (
		p int32
		k int32
	)
	_, _ = p, k
	p = 0
	k = 0
	Node[j].oldDepth = Node[j].initDepth
	Node[j].newDepth = Node[j].oldDepth
	Node[j].crownElev = Node[j].invertElev
	Node[j].fullVolume = node_getVolume(j, Node[j].fullDepth)
	Node[j].oldVolume = node_getVolume(j, Node[j].oldDepth)
	Node[j].newVolume = Node[j].oldVolume
	for p = 0; p < Nobjects[POLLUT]; p++ {
		Node[j].oldQual[p] = float64(0.0)
		Node[j].newQual[p] = float64(0.0)
	}
	Node[j].oldLatFlow = float64(0.0)
	Node[j].newLatFlow = float64(0.0)
	Node[j].apiExtInflow = float64(0.0)
	Node[j].losses = float64(0.0)
	if Node[j].type_ == STORAGE {
		k = Node[j].subIndex
		Storage[k].hrt = float64(0.0)
		if Storage[k].exfil != nil {
			exfil_initState(k)
		}
	}
	if Node[j].type_ == OUTFALL {
		k = Node[j].subIndex
		if Outfall[k].routeTo >= 0 {
			Outfall[k].vRouted = float64(0.0)
			for p = 0; p < Nobjects[POLLUT]; p++ {
				Outfall[k].wRouted[p] = float64(0.0)
			}
		}
	}
}

// node.c:293
func node_setOldHydState(j int32) {
	Node[j].oldDepth = Node[j].newDepth
	Node[j].oldVolume = Node[j].newVolume
	Node[j].oldFlowInflow = Node[j].inflow
	Node[j].oldNetInflow = float64(Node[j].inflow - Node[j].outflow)
}

// node.c:308
func node_setOldQualState(j int32) {
	var (
		p int32
	)
	_ = p
	p = 0
	for p = 0; p < Nobjects[POLLUT]; p++ {
		Node[j].oldQual[p] = Node[j].newQual[p]
		Node[j].newQual[p] = float64(0.0)
	}
}

// node.c:325
func node_initFlows(j int32, tStep float64) {
	Node[j].inflow = Node[j].newLatFlow
	Node[j].outflow = Node[j].losses
	if Node[j].newVolume > Node[j].fullVolume {
		Node[j].overflow = float64((float64(Node[j].newVolume - Node[j].fullVolume)) / tStep)
	} else {
		Node[j].overflow = float64(0.0)
	}
}

// node.c:345
func node_getDepth(j int32, v float64) float64 {
	switch Node[j].type_ {
	case STORAGE:
		return storage_getDepth(j, v)
	default:
		return float64(0.0)
	}
	return 0
}

// node.c:362
func node_getVolume(j int32, d float64) float64 {
	switch Node[j].type_ {
	case STORAGE:
		return storage_getVolume(j, d)
	default:
		if Node[j].fullDepth > float64(0.0) {
			return float64(Node[j].fullVolume * (float64(d / Node[j].fullDepth)))
		} else {
			return float64(0.0)
		}
	}
	return 0
}

// node.c:383
func node_getSurfArea(j int32, d float64) float64 {
	switch Node[j].type_ {
	case STORAGE:
		return storage_getSurfArea(j, d)
	default:
		return float64(0.0)
	}
	return 0
}

// node.c:400
func node_getOutflow(j int32, k int32) float64 {
	switch Node[j].type_ {
	case DIVIDER:
		return divider_getOutflow(j, k)
	case STORAGE:
		return storage_getOutflow(j, k)
	default:
		return float64(Node[j].inflow + Node[j].overflow)
	}
	return 0
}

// node.c:418
func node_getMaxOutflow(j int32, q float64, tStep float64) float64 {
	var (
		qMax float64
	)
	_ = qMax
	qMax = 0
	if Node[j].fullVolume > float64(0.0) {
		qMax = float64(Node[j].inflow + float64(Node[j].oldVolume/tStep))
		if q > qMax {
			q = qMax
		}
	}
	if (float64(0.0)) >= (q) {
		return (float64(0.0))
	}
	return (q)
}

// node.c:438
func node_getSystemOutflow(j int32, isFlooded *int32) float64 {
	var (
		outflow float64
	)
	_ = outflow
	outflow = float64(0.0)
	*isFlooded = 0
	if Node[j].type_ == OUTFALL {
		if Node[j].outflow == float64(0.0) {
			outflow = Node[j].inflow
		} else {
			if Node[j].inflow == float64(0.0) {
				outflow = (-Node[j].outflow)
				Node[j].inflow = M.Fabs(outflow)
			}
		}
		Node[j].overflow = float64(0.0)
		Node[j].newVolume = float64(0.0)
	} else if ((RouteModel != DW) && (Node[j].degree == 0)) && (Node[j].type_ != STORAGE) {
		if Node[j].outflow == float64(0.0) {
			outflow = Node[j].inflow
		}
		Node[j].overflow = float64(0.0)
		Node[j].newVolume = float64(0.0)
	} else {
		if Node[j].newVolume <= Node[j].fullVolume {
			outflow = Node[j].overflow
		}
		if outflow > float64(0.0) {
			*isFlooded = 1
		}
	}
	return outflow
}

// node.c:497
func node_getResults(j int32, f float64, x []float32) {
	var (
		p  int32
		z  float64
		f1 float64
	)
	_, _, _ = p, z, f1
	p = 0
	z = 0
	f1 = float64(float64(1.0) - f)
	z = float64((float64(float64(f1*Node[j].oldDepth) + float64(f*Node[j].newDepth))) * UCF(LENGTH))
	x[NODE_DEPTH] = float32(float64(z))
	z = float64(Node[j].invertElev * UCF(LENGTH))
	x[NODE_HEAD] = float32(x[NODE_DEPTH] + float32(float64(z)))
	z = float64((float64(float64(f1*Node[j].oldVolume) + float64(f*Node[j].newVolume))) * UCF(VOLUME))
	x[NODE_VOLUME] = float32(float64(z))
	z = float64((float64(float64(f1*Node[j].oldLatFlow) + float64(f*Node[j].newLatFlow))) * UCF(FLOW))
	x[NODE_LATFLOW] = float32(float64(z))
	z = float64((float64(float64(f1*Node[j].oldFlowInflow) + float64(f*Node[j].inflow))) * UCF(FLOW))
	x[NODE_INFLOW] = float32(float64(z))
	z = float64(Node[j].overflow * UCF(FLOW))
	x[NODE_OVERFLOW] = float32(float64(z))
	if !(IgnoreQuality != 0) {
		for p = 0; p < Nobjects[POLLUT]; p++ {
			z = float64(float64(f1*Node[j].oldQual[p]) + float64(f*Node[j].newQual[p]))
			x[(NODE_QUAL + p)] = float32(float64(z))
		}
	}
}

// node.c:532
func node_setOutletDepth(j int32, yNorm float64, yCrit float64, z float64) {
	switch Node[j].type_ {
	case STORAGE:
		return
	case OUTFALL:
		outfall_setOutletDepth(j, yNorm, yCrit, z)
		break
	default:
		if z > float64(0.0) {
			Node[j].newDepth = float64(0.0)
		} else {
			if (yNorm) <= (yCrit) {
				Node[j].newDepth = (yNorm)
			} else {
				Node[j].newDepth = (yCrit)
			}
		}
	}
}

// node.c:562
func node_getPondedArea(j int32, d float64) float64 {
	var (
		a_2 float64
	)
	_ = a_2
	a_2 = 0
	if (d <= Node[j].fullDepth) || (Node[j].pondedArea == float64(0.0)) {
		return node_getSurfArea(j, d)
	}
	d = float64(d - Node[j].fullDepth)
	a_2 = Node[j].pondedArea
	if a_2 <= float64(0.0) {
		a_2 = node_getSurfArea(j, Node[j].fullDepth)
	}
	return a_2
}

// node.c:589
func node_getLosses(j int32, tStep float64) float64 {
	if Node[j].type_ == STORAGE {
		return storage_getLosses(j, tStep)
	} else {
		return float64(0.0)
	}
	return 0
}

// node.c:606
func junc_readParams(j int32, k int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		i  int32
		x  [6]float64
		id R.CPtr
	)
	_, _, _ = i, x, id
	i = 0
	x = ([6]float64{})
	id = (R.CPtr{})
	if ntoks < 2 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(NODE, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	for i = 1; i <= 5; i++ {
		x[(i - 1)] = float64(0.0)
		if i < ntoks {
			if !(getDouble(tok[i], &x[(i-1)]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[i])
			}
		}
	}
	for i = 1; i <= 4; i++ {
		if x[i] < float64(0.0) {
			return error_setInpError(ERR_NUMBER, tok[(i+1)])
		}
	}
	Node[j].ID = id
	node_setParams(j, JUNCTION, k, x[:])
	return 0
}

// node.c:654
func storage_readParams(j int32, k int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		i   int32
		m   int32
		n_2 int32
		x   [10]float64
		y_2 [3]float64
		A_2 float64
		B   float64
		L   float64
		W_2 float64
		Z   float64
		id  R.CPtr
	)
	_, _, _, _, _, _, _, _, _, _, _ = i, m, n_2, x, y_2, A_2, B, L, W_2, Z, id
	i = 0
	m = 0
	n_2 = 0
	x = ([10]float64{})
	y_2 = ([3]float64{})
	A_2 = 0
	B = 0
	L = 0
	W_2 = 0
	Z = 0
	id = (R.CPtr{})
	if ntoks < 6 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(NODE, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	for i = 1; i <= 3; i++ {
		if !(getDouble(tok[i], &x[(i-1)]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[i])
		}
	}
	m = findmatch(tok[4], RelationWords[:])
	if m < 0 {
		return error_setInpError(ERR_KEYWORD, tok[4])
	}
	x[3] = float64(m)
	x[4] = float64(0.0)
	x[5] = float64(0.0)
	x[6] = float64(0.0)
	x[7] = (-(float64(1.0)))
	x[8] = float64(0.0)
	x[9] = float64(0.0)
	if m == TABULAR {
		i = project_findObject(CURVE, tok[5])
		if i < 0 {
			return error_setInpError(ERR_NAME, tok[5])
		}
		x[7] = float64(i)
		n_2 = 6
	} else {
		if ntoks < 8 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		for i = 5; i <= 7; i++ {
			if !(getDouble(tok[i], &y_2[(i-5)]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[i])
			}
		}
		n_2 = 8
	}
	switch m {
	case FUNCTIONAL:
		if y_2[2] < float64(0.0) {
			return error_setInpError(ERR_NUMBER, tok[7])
		}
		break
	case CYLINDRICAL, CONICAL, PARABOLOID, PYRAMIDAL:
		if y_2[0] <= float64(0.0) {
			return error_setInpError(ERR_NUMBER, tok[5])
		}
		if y_2[1] <= float64(0.0) {
			return error_setInpError(ERR_NUMBER, tok[6])
		}
		if y_2[2] < float64(0.0) {
			return error_setInpError(ERR_NUMBER, tok[7])
		}
		break
	}
	if (m == PARABOLOID) && (y_2[2] == float64(0.0)) {
		return error_setInpError(ERR_NUMBER, tok[7])
	}
	switch m {
	case FUNCTIONAL:
		x[4] = y_2[0]
		x[5] = y_2[1]
		x[6] = y_2[2]
		break
	case CYLINDRICAL:
		A_2 = float64(y_2[0] / float64(2.0))
		B = float64(y_2[1] / float64(2.0))
		x[4] = float64(0.0)
		x[5] = float64(0.0)
		x[6] = float64(float64(float64(3.141592654)*A_2) * B)
		break
	case CONICAL:
		A_2 = float64(y_2[0] / float64(2.0))
		B = float64(y_2[1] / float64(2.0))
		Z = y_2[2]
		x[4] = float64(float64(float64(float64(2.0)*float64(3.141592654))*B) * Z)
		x[5] = float64(float64(float64(float64(float64(3.141592654)*B)/A_2)*Z) * Z)
		x[6] = float64(float64(float64(3.141592654)*A_2) * B)
		break
	case PARABOLOID:
		A_2 = float64(y_2[0] / float64(2.0))
		B = float64(y_2[1] / float64(2.0))
		Z = y_2[2]
		x[4] = float64(float64(float64(float64(3.141592654)*A_2)*B) / Z)
		x[5] = float64(0.0)
		x[6] = float64(0.0)
		break
	case PYRAMIDAL:
		L = y_2[0]
		W_2 = y_2[1]
		Z = y_2[2]
		x[4] = float64(float64(float64(2.0)*(float64(L+W_2))) * Z)
		x[5] = float64(float64(float64(4.0)*Z) * Z)
		x[6] = float64(L * W_2)
		break
	}
	if ntoks > n_2 {
		if !(getDouble(tok[n_2], &x[8]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[n_2])
		}
		n_2++
	}
	if ntoks > n_2 {
		if !(getDouble(tok[n_2], &x[9]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[n_2])
		}
		n_2++
	}
	Node[j].ID = id
	node_setParams(j, STORAGE, k, x[:])
	if ntoks > n_2 {
		return exfil_readStorageParams(k, tok, ntoks, n_2)
	}
	return 0
}

// node.c:813
func storage_getDepth(j int32, v float64) float64 {
	var (
		k          int32
		i          int32
		shape      int32
		a0         float64
		a1         float64
		a2         float64
		d          float64
		e          float64
		storageVol TStorageVol
	)
	_, _, _, _, _, _, _, _, _ = k, i, shape, a0, a1, a2, d, e, storageVol
	k = Node[j].subIndex
	i = Storage[k].aCurve
	shape = Storage[k].shape
	a0 = Storage[k].a0
	a1 = Storage[k].a1
	a2 = Storage[k].a2
	d = 0
	e = 0
	storageVol = (TStorageVol{})
	if (Node[j].fullVolume > float64(0.0)) && (v >= Node[j].fullVolume) {
		return Node[j].fullDepth
	}
	if v == float64(0.0) {
		return float64(0.0)
	}
	v = float64(v * (UCF(VOLUME)))
	storageVol.k = k
	storageVol.v = v
	switch shape {
	case TABULAR:
		i = Storage[k].aCurve
		if i >= 0 {
			d = table_getStorageDepth(&Curve[i], v)
		} else {
			d = float64(0.0)
		}
		break
	case CYLINDRICAL:
		d = float64(v / a0)
		break
	case PARABOLOID:
		d = M.Sqrt(float64(float64(float64(2.0)*v) / a1))
		break
	case FUNCTIONAL:
		if a2 == float64(0.0) {
			d = float64(v / (float64(a0 + a1)))
		} else if a0 == float64(0.0) {
			e = float64(float64(1.0) / (float64(a2 + float64(1.0))))
			d = M.Pow(float64(v/(float64(a1*e))), e)
		} else if (a2 == float64(1.0)) && (a1 > float64(0.0)) {
			d = float64((float64(M.Sqrt(float64(float64(a0*a0)+float64(float64(float64(2.0)*a1)*v))) - a0)) / a1)
		} else {
			d = float64(v / (float64(a0 + a1)))
			findroot_Newton(float64(0.0), float64(Node[j].fullDepth*UCF(LENGTH)), &d, float64(0.001), storage_getVolDiff, &storageVol)
		}
		break
	case CONICAL, PYRAMIDAL:
		d = float64(v / a0)
		findroot_Newton(float64(0.0), float64(Node[j].fullDepth*UCF(LENGTH)), &d, float64(0.001), storage_getVolDiff, &storageVol)
		break
	default:
		d = float64(0.0)
	}
	d = float64(d / (UCF(LENGTH)))
	if d > Node[j].fullDepth {
		d = Node[j].fullDepth
	}
	return d
}

// node.c:906
func storage_getVolDiff(y_2 float64, f *float64, df *float64, p any) {
	var (
		k          int32
		storageVol *TStorageVol
	)
	_, _ = k, storageVol
	k = 0
	storageVol = nil
	storageVol = p.(*TStorageVol)
	k = storageVol.k
	*f = float64(storage_getVolume(k, y_2) - storageVol.v)
	*df = storage_getSurfArea(k, y_2)
}

// node.c:930
func storage_getVolume(j int32, d float64) float64 {
	var (
		k   int32
		i   int32
		n_2 float64
		v   float64
	)
	_, _, _, _ = k, i, n_2, v
	k = Node[j].subIndex
	i = 0
	n_2 = 0
	v = 0
	if d == float64(0.0) {
		return float64(0.0)
	}
	if (d >= Node[j].fullDepth) && (Node[j].fullVolume > float64(0.0)) {
		return Node[j].fullVolume
	}
	switch Storage[k].shape {
	case TABULAR:
		i = Storage[k].aCurve
		if i >= 0 {
			return float64(table_getStorageVolume(&Curve[i], float64(d*UCF(LENGTH))) / UCF(VOLUME))
		} else {
			return float64(0.0)
		}
	case FUNCTIONAL:
		d = float64(d * (UCF(LENGTH)))
		n_2 = float64(Storage[k].a2 + float64(1.0))
		v = float64((float64(Storage[k].a0 * d)) + float64(float64(Storage[k].a1/n_2)*M.Pow(d, n_2)))
		return float64(v / UCF(VOLUME))
	case CYLINDRICAL, CONICAL, PARABOLOID, PYRAMIDAL:
		d = float64(d * (UCF(LENGTH)))
		v = float64(d * (float64(Storage[k].a0 + float64(d*(float64(float64(Storage[k].a1/float64(2.0))+float64(float64(d*Storage[k].a2)/float64(3.0))))))))
		return float64(v / UCF(VOLUME))
	default:
		return float64(0.0)
	}
	return 0
}

// node.c:979
func storage_getSurfArea(j int32, d float64) float64 {
	var (
		area float64
		k    int32
		i    int32
	)
	_, _, _ = area, k, i
	area = float64(0.0)
	k = Node[j].subIndex
	i = 0
	switch Storage[k].shape {
	case TABULAR:
		i = Storage[k].aCurve
		if i >= 0 {
			area = table_lookupEx(&Curve[i], float64(d*UCF(LENGTH)))
		}
		break
	case FUNCTIONAL:
		area = float64(Storage[k].a0 + float64(Storage[k].a1*M.Pow(float64(d*UCF(LENGTH)), Storage[k].a2)))
		break
	case CYLINDRICAL, CONICAL, PARABOLOID, PYRAMIDAL:
		d = float64(d * (UCF(LENGTH)))
		area = float64(Storage[k].a0 + float64(d*(float64(Storage[k].a1+float64(d*Storage[k].a2)))))
		break
	default:
		return float64(0.0)
	}
	return float64(float64(area/UCF(LENGTH)) / UCF(LENGTH))
}

// node.c:1022
func storage_getOutflow(j int32, i int32) float64 {
	var (
		k   int32
		a_2 float64
		y_2 float64
	)
	_, _, _ = k, a_2, y_2
	k = 0
	a_2 = 0
	y_2 = 0
	if Link[i].type_ != CONDUIT {
		return float64(0.0)
	}
	y_2 = float64(Node[j].newDepth - Link[i].offset1)
	if y_2 <= float64(0.0) {
		return float64(0.0)
	}
	if y_2 >= Link[i].xsect.yFull {
		return Link[i].qFull
	}
	k = Link[i].subIndex
	a_2 = xsect_getAofY(&Link[i].xsect, y_2)
	return float64(Conduit[k].beta * xsect_getSofA(&Link[i].xsect, a_2))
}

// node.c:1052
func storage_getLosses(j int32, tStep float64) float64 {
	var (
		k         int32
		depth     float64
		area      float64
		evapRate  float64
		exfilRate float64
		totalLoss float64
		lossRatio float64
		exfil     *TExfil
	)
	_, _, _, _, _, _, _, _ = k, depth, area, evapRate, exfilRate, totalLoss, lossRatio, exfil
	k = 0
	depth = 0
	area = 0
	evapRate = float64(0.0)
	exfilRate = float64(0.0)
	totalLoss = float64(0.0)
	lossRatio = 0
	exfil = nil
	k = Node[j].subIndex
	evapRate = float64(Evap.rate * Storage[k].fEvap)
	exfil = Storage[k].exfil
	if (evapRate > float64(0.0)) || (exfil != nil) {
		depth = Node[j].newDepth
		area = storage_getSurfArea(j, depth)
		if Node[j].newVolume > float64(0.0001) {
			evapRate = float64(area * evapRate)
		}
		if exfil != nil {
			exfilRate = exfil_getLoss(exfil, tStep, depth, area)
		}
		totalLoss = float64((float64(evapRate + exfilRate)) * tStep)
		if totalLoss > Node[j].newVolume {
			lossRatio = float64(Node[j].newVolume / totalLoss)
			evapRate = float64(evapRate * lossRatio)
			exfilRate = float64(exfilRate * lossRatio)
		}
	}
	Storage[Node[j].subIndex].evapLoss = float64(evapRate * tStep)
	Storage[Node[j].subIndex].exfilLoss = float64(exfilRate * tStep)
	return float64(evapRate + exfilRate)
}

// node.c:1113
func divider_readParams(j int32, k int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		i   int32
		m   int32
		m1  int32
		m2  int32
		n_2 int32
		x   [11]float64
		id  R.CPtr
	)
	_, _, _, _, _, _, _ = i, m, m1, m2, n_2, x, id
	i = 0
	m = 0
	m1 = 0
	m2 = 0
	n_2 = 0
	x = ([11]float64{})
	id = (R.CPtr{})
	if ntoks < 4 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(NODE, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	if !(getDouble(tok[1], &x[0]) != 0) {
		return error_setInpError(ERR_NUMBER, tok[1])
	}
	for i = 1; i < 11; i++ {
		x[i] = float64(0.0)
	}
	if (R.Strlen(tok[2]) == 0) || (R.Strcmp(tok[2], L275) == 0) {
		x[1] = (-(float64(1.0)))
	} else {
		m1 = project_findObject(LINK, tok[2])
		if m1 < 0 {
			return error_setInpError(ERR_NAME, tok[2])
		}
		x[1] = float64(m1)
	}
	n_2 = 4
	m1 = findmatch(tok[3], DividerTypeWords[:])
	if m1 < 0 {
		return error_setInpError(ERR_KEYWORD, tok[3])
	}
	x[2] = float64(m1)
	x[3] = float64((-1))
	if m1 == TABULAR_DIVIDER {
		if ntoks < 5 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		m2 = project_findObject(CURVE, tok[4])
		if m2 < 0 {
			return error_setInpError(ERR_NAME, tok[4])
		}
		x[3] = float64(m2)
		n_2 = 5
	}
	if m1 == CUTOFF_DIVIDER {
		if ntoks < 5 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		if !(getDouble(tok[4], &x[4]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[4])
		}
		n_2 = 5
	}
	if m1 == WEIR_DIVIDER {
		if ntoks < 7 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		for i = 4; i < 7; i++ {
			if !(getDouble(tok[i], &x[i]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[i])
			}
		}
		n_2 = 7
	}
	if m1 == OVERFLOW_DIVIDER {
		n_2 = 4
	}
	m = 7
	for i = n_2; (i < ntoks) && (m < 11); i++ {
		if !(getDouble(tok[i], &x[m]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[i])
		}
		m++
	}
	Node[j].ID = id
	node_setParams(j, DIVIDER, k, x[:])
	return 0
}

// node.c:1216
func divider_validate(j int32) {
	var (
		i int32
		k int32
	)
	_, _ = i, k
	i = 0
	k = 0
	k = Node[j].subIndex
	i = Divider[k].link
	if (i < 0) || ((Link[i].node1 != j) && (Link[i].node2 != j)) {
		report_writeErrorMsg(ERR_DIVIDER_LINK, Node[j].ID)
	}
	if Divider[k].type_ == WEIR_DIVIDER {
		if (Divider[k].dhMax <= float64(0.0)) || (Divider[k].cWeir <= float64(0.0)) {
			report_writeErrorMsg(ERR_WEIR_DIVIDER, Node[j].ID)
		} else {
			Divider[k].qMax = float64(float64(Divider[k].cWeir*M.Pow(Divider[k].dhMax, float64(1.5))) / UCF(FLOW))
			if Divider[k].qMin > Divider[k].qMax {
				report_writeErrorMsg(ERR_WEIR_DIVIDER, Node[j].ID)
			}
		}
	}
}

// node.c:1251
func divider_getOutflow(j int32, k int32) float64 {
	var (
		i    int32
		m    int32
		qIn  float64
		qOut float64
		f    float64
	)
	_, _, _, _, _ = i, m, qIn, qOut, f
	i = 0
	m = 0
	qIn = 0
	qOut = 0
	f = 0
	qIn = float64(Node[j].inflow + Node[j].overflow)
	i = Node[j].subIndex
	switch Divider[i].type_ {
	case CUTOFF_DIVIDER:
		if qIn <= Divider[i].qMin {
			qOut = float64(0.0)
		} else {
			qOut = float64(qIn - Divider[i].qMin)
		}
		break
	case OVERFLOW_DIVIDER:
		if k != Divider[i].link {
			qOut = qIn
		} else {
			qOut = float64(qIn - Node[j].outflow)
		}
		if qOut < float64(1e-05) {
			qOut = float64(0.0)
		}
		return qOut
	case WEIR_DIVIDER:
		if qIn <= Divider[i].qMin {
			qOut = float64(0.0)
		} else {
			f = float64((float64(qIn - Divider[i].qMin)) / (float64(Divider[i].qMax - Divider[i].qMin)))
			if f > float64(1.0) {
				qOut = float64(Divider[i].qMax * M.Sqrt(f))
			} else {
				qOut = float64(float64(Divider[i].cWeir*M.Pow(float64(f*Divider[i].dhMax), float64(1.5))) / UCF(FLOW))
			}
		}
		break
	case TABULAR_DIVIDER:
		m = Divider[i].flowCurve
		if m >= 0 {
			qOut = float64(table_lookup(&Curve[m], float64(qIn*UCF(FLOW))) / UCF(FLOW))
		} else {
			qOut = float64(0.0)
		}
		break
	default:
		qOut = float64(0.0)
	}
	if qOut > qIn {
		qOut = qIn
	}
	if k != Divider[i].link {
		qOut = float64(qIn - qOut)
	}
	return qOut
}

// node.c:1333
func outfall_readParams(j int32, k int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		i   int32
		m   int32
		n_2 int32
		x   [7]float64
		id  R.CPtr
	)
	_, _, _, _, _ = i, m, n_2, x, id
	i = 0
	m = 0
	n_2 = 0
	x = ([7]float64{})
	id = (R.CPtr{})
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	id = project_findID(NODE, tok[0])
	if id.IsNull() {
		return error_setInpError(ERR_NAME, tok[0])
	}
	if !(getDouble(tok[1], &x[0]) != 0) {
		return error_setInpError(ERR_NUMBER, tok[1])
	}
	i = findmatch(tok[2], OutfallTypeWords[:])
	if i < 0 {
		return error_setInpError(ERR_KEYWORD, tok[2])
	}
	x[1] = float64(i)
	x[2] = float64(0.0)
	x[3] = (-(float64(1.0)))
	x[4] = (-(float64(1.0)))
	x[5] = float64(0.0)
	x[6] = (-(float64(1.0)))
	n_2 = 4
	if i >= FIXED_OUTFALL {
		if ntoks < 4 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		n_2 = 5
		switch i {
		case FIXED_OUTFALL:
			if !(getDouble(tok[3], &x[2]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[3])
			}
			break
		case TIDAL_OUTFALL:
			m = project_findObject(CURVE, tok[3])
			if m < 0 {
				return error_setInpError(ERR_NAME, tok[3])
			}
			x[3] = float64(m)
			break
		case TIMESERIES_OUTFALL:
			m = project_findObject(TSERIES, tok[3])
			if m < 0 {
				return error_setInpError(ERR_NAME, tok[3])
			}
			x[4] = float64(m)
			Tseries[m].refersTo = TIMESERIES_OUTFALL
		}
	}
	if ntoks == n_2 {
		m = findmatch(tok[(n_2-1)], NoYesWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, tok[(n_2-1)])
		}
		x[5] = float64(m)
	}
	if ntoks == (n_2 + 1) {
		m = project_findObject(SUBCATCH, tok[n_2])
		if m < 0 {
			return error_setInpError(ERR_NAME, tok[n_2])
		}
		x[6] = float64(m)
	}
	Node[j].ID = id
	node_setParams(j, OUTFALL, k, x[:])
	return 0
}

// node.c:1413
func outfall_setOutletDepth(j int32, yNorm float64, yCrit float64, z float64) {
	var (
		x           float64
		y_2         float64
		yNew        float64
		stage       float64
		k           int32
		i           int32
		currentDate float64
	)
	_, _, _, _, _, _, _ = x, y_2, yNew, stage, k, i, currentDate
	x = 0
	y_2 = 0
	yNew = 0
	stage = 0
	k = 0
	i = Node[j].subIndex
	currentDate = 0
	switch Outfall[i].type_ {
	case FREE_OUTFALL:
		if z > float64(0.0) {
			Node[j].newDepth = float64(0.0)
		} else {
			if (yNorm) <= (yCrit) {
				Node[j].newDepth = (yNorm)
			} else {
				Node[j].newDepth = (yCrit)
			}
		}
		return
	case NORMAL_OUTFALL:
		if z > float64(0.0) {
			Node[j].newDepth = float64(0.0)
		} else {
			Node[j].newDepth = yNorm
		}
		return
	case FIXED_OUTFALL:
		stage = Outfall[i].fixedStage
		break
	case TIDAL_OUTFALL:
		k = Outfall[i].tideCurve
		table_getFirstEntry(&Curve[k], &x, &y_2)
		currentDate = float64(NewRoutingTime / float64(86400000.0))
		x = float64(x + (float64((float64(currentDate - M.Floor(currentDate))) * float64(24.0))))
		stage = float64(table_lookup(&Curve[k], x) / UCF(LENGTH))
		break
	case TIMESERIES_OUTFALL:
		k = Outfall[i].stageSeries
		currentDate = float64(StartDateTime + float64(NewRoutingTime/float64(86400000.0)))
		stage = float64(table_tseriesLookup(&Tseries[k], currentDate, 1) / UCF(LENGTH))
		break
	default:
		stage = Node[j].invertElev
	}
	if (yCrit) <= (yNorm) {
		yCrit = (yCrit)
	} else {
		yCrit = (yNorm)
	}
	if float64(float64(yCrit+z)+Node[j].invertElev) < stage {
		yNew = float64(stage - Node[j].invertElev)
	} else if z > float64(0.0) {
		if stage < float64(Node[j].invertElev+z) {
			if (float64(0.0)) >= (float64(stage - Node[j].invertElev)) {
				yNew = (float64(0.0))
			} else {
				yNew = (float64(stage - Node[j].invertElev))
			}
		} else {
			yNew = float64(z + yCrit)
		}
	} else {
		yNew = yCrit
	}
	Node[j].newDepth = yNew
}
