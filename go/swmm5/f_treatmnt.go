// Generated from EPA SWMM 5.2.4 (treatmnt.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

var PVMAX int32 = 5
var ErrCode int32
var J int32
var Dt float64
var Q float64
var V float64
var R_ []float64
var Cin []float64

// treatmnt.c:66
func treatmnt_open() int32 {
	R_ = nil
	Cin = nil
	if Nobjects[POLLUT] > 0 {
		R_ = make([]float64, int(uint64(Nobjects[POLLUT])))
		Cin = make([]float64, int(uint64(Nobjects[POLLUT])))
		if (R_ == nil) || (Cin == nil) {
			report_writeErrorMsg(ERR_MEMORY, L0)
			return 0
		}
	}
	return 1
}

// treatmnt.c:90
func treatmnt_close() {
	{
		if R_ != nil {
			_ = R_
			R_ = nil
		}
	}
	{
		if Cin != nil {
			_ = Cin
			Cin = nil
		}
	}
}

// treatmnt.c:103
func treatmnt_readExpression(tok []R.CPtr, ntoks int32) int32 {
	var (
		s        R.CPtr
		expr     R.CPtr
		i        int32
		j        int32
		k        int32
		p        int32
		equation *MathExpr
	)
	_, _, _, _, _, _, _ = s, expr, i, j, k, p, equation
	s = R.Buf(1025)
	expr = (R.CPtr{})
	i = 0
	j = 0
	k = 0
	p = 0
	equation = nil
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(NODE, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	p = project_findObject(POLLUT, tok[1])
	if p < 0 {
		return error_setInpError(ERR_NAME, tok[1])
	}
	sstrncpy(s, tok[2], 1024)
	for i = 3; i < ntoks; i++ {
		sstrcat(s, L33, 1024)
		sstrcat(s, tok[i], 1024)
	}
	if (func() int32 {
		if ((s.G(0)) >= 97) && ((s.G(0)) <= 122) {
			return ((s.G(0)) & (^32))
		}
		return (s.G(0))
	}()) == 82 {
		k = 0
	} else if (func() int32 {
		if ((s.G(0)) >= 97) && ((s.G(0)) <= 122) {
			return ((s.G(0)) & (^32))
		}
		return (s.G(0))
	}()) == 67 {
		k = 1
	} else {
		return error_setInpError(ERR_KEYWORD, tok[2])
	}
	expr = R.Strchr(s, 61)
	if expr.IsNull() {
		return error_setInpError(ERR_KEYWORD, L0)
	} else {
		expr = expr.Add(1)
	}
	if Node[j].treatment == nil {
		if !(createTreatment(j) != 0) {
			return error_setInpError(ERR_MEMORY, L0)
		}
	}
	equation = mathexpr_create(expr, treatmnt_getVariableIndex)
	if equation == nil {
		return error_setInpError(ERR_MATH_EXPR, L0)
	}
	if Node[j].treatment != nil {
		Node[j].treatment[p].treatType = k
		Node[j].treatment[p].equation = equation
	}
	return 0
}

// treatmnt.c:165
func treatmnt_delete(j int32) {
	var (
		p int32
	)
	_ = p
	p = 0
	if Node[j].treatment != nil {
		for p = 0; p < Nobjects[POLLUT]; p++ {
			mathexpr_delete(Node[j].treatment[p].equation)
		}
		_ = Node[j].treatment
	}
	Node[j].treatment = nil
}

// treatmnt.c:184
func treatmnt_setInflow(qIn float64, wIn []float64) {
	var (
		p int32
	)
	_ = p
	p = 0
	if qIn > float64(0.0) {
		for p = 0; p < Nobjects[POLLUT]; p++ {
			Cin[p] = float64(wIn[p] / qIn)
		}
	} else {
		for p = 0; p < Nobjects[POLLUT]; p++ {
			Cin[p] = float64(0.0)
		}
	}
}

// treatmnt.c:202
func treatmnt_treat(j int32, q float64, v float64, tStep float64) {
	var (
		p         int32
		cOut      float64
		massLost  float64
		treatment *TTreatment
	)
	_, _, _, _ = p, cOut, massLost, treatment
	p = 0
	cOut = 0
	massLost = 0
	treatment = nil
	if Node[j].treatment == nil {
		return
	}
	ErrCode = 0
	J = j
	Dt = tStep
	Q = q
	V = v
	for p = 0; p < Nobjects[POLLUT]; p++ {
		R_[p] = (-(float64(1.0)))
	}
	for p = 0; p < Nobjects[POLLUT]; p++ {
		treatment = &Node[j].treatment[p]
		if treatment.equation == nil {
			R_[p] = float64(0.0)
		} else if (treatment.treatType == REMOVAL) && (q <= float64(1e-10)) {
			R_[p] = float64(0.0)
		} else {
			getRemoval(p)
		}
	}
	if ErrCode == ERR_CYCLIC_TREATMENT {
		report_writeErrorMsg(ERR_CYCLIC_TREATMENT, Node[J].ID)
	} else {
		for p = 0; p < Nobjects[POLLUT]; p++ {
			if R_[p] == float64(0.0) {
				continue
			}
			treatment = &Node[j].treatment[p]
			if treatment.treatType == REMOVAL {
				if Cin[p] == float64(0.0) {
					cOut = Node[j].newQual[p]
				} else {
					cOut = float64((float64(float64(1.0) - R_[p])) * Cin[p])
				}
				if (cOut) <= (Node[j].newQual[p]) {
					cOut = (cOut)
				} else {
					cOut = (Node[j].newQual[p])
				}
			} else {
				cOut = float64((float64(float64(1.0) - R_[p])) * Node[j].newQual[p])
			}
			massLost = float64((float64(float64(float64(float64(Cin[p]*q)*tStep)+float64(Node[j].oldQual[p]*Node[j].oldVolume)) - float64(cOut*(float64(float64(q*tStep)+Node[j].oldVolume))))) / tStep)
			if (float64(0.0)) >= (massLost) {
				massLost = (float64(0.0))
			} else {
				massLost = (massLost)
			}
			massbal_addReactedMass(p, massLost)
			Node[j].newQual[p] = cOut
		}
	}
}

// treatmnt.c:288
func createTreatment(j int32) int32 {
	var (
		p int32
	)
	_ = p
	p = 0
	Node[j].treatment = make([]TTreatment, int(uint64(Nobjects[POLLUT])))
	if Node[j].treatment == nil {
		return 0
	}
	for p = 0; p < Nobjects[POLLUT]; p++ {
		Node[j].treatment[p].equation = nil
	}
	return 1
}

// treatmnt.c:311
func treatmnt_getVariableIndex(s R.CPtr) int32 {
	var (
		k int32
		m int32
	)
	_, _ = k, m
	k = 0
	m = PVMAX
	k = findmatch(s, ProcessVarWords[:])
	if k >= 0 {
		return k
	}
	k = project_findObject(POLLUT, s)
	if k >= 0 {
		return (k + m)
	}
	if ((func() int32 {
		if ((s.G(0)) >= 97) && ((s.G(0)) <= 122) {
			return ((s.G(0)) & (^32))
		}
		return (s.G(0))
	}()) == 82) && (s.G(1) == 95) {
		k = project_findObject(POLLUT, s.Add(2))
		if k >= 0 {
			return ((Nobjects[POLLUT] + k) + m)
		}
	}
	return (-1)
}

// treatmnt.c:340
func treatmnt_getVariableValue(varCode int32) float64 {
	var (
		p         int32
		a1        float64
		a2        float64
		y_2       float64
		treatment *TTreatment
	)
	_, _, _, _, _ = p, a1, a2, y_2, treatment
	p = 0
	a1 = 0
	a2 = 0
	y_2 = 0
	treatment = nil
	if varCode < PVMAX {
		switch varCode {
		case pvHRT:
			if Node[J].type_ == STORAGE {
				return float64(Storage[Node[J].subIndex].hrt / float64(3600.0))
			} else {
				return float64(0.0)
			}
		case pvDT:
			return Dt
		case pvFLOW:
			return float64(Q * UCF(FLOW))
		case pvDEPTH:
			y_2 = float64((float64(Node[J].oldDepth + Node[J].newDepth)) / float64(2.0))
			return float64(y_2 * UCF(LENGTH))
		case pvAREA:
			a1 = node_getSurfArea(J, Node[J].oldDepth)
			a2 = node_getSurfArea(J, Node[J].newDepth)
			return float64(float64(float64((float64(a1+a2))/float64(2.0))*UCF(LENGTH)) * UCF(LENGTH))
		default:
			return float64(0.0)
		}
	} else if varCode < (PVMAX + Nobjects[POLLUT]) {
		p = (varCode - PVMAX)
		treatment = &Node[J].treatment[p]
		if treatment.treatType == REMOVAL {
			return Cin[p]
		}
		return Node[J].newQual[p]
	} else {
		p = ((varCode - PVMAX) - Nobjects[POLLUT])
		if p >= Nobjects[POLLUT] {
			return float64(0.0)
		}
		return getRemoval(p)
	}
	return 0
}

// treatmnt.c:404
func getRemoval(p int32) float64 {
	var (
		c0        float64
		r         float64
		treatment *TTreatment
	)
	_, _, _ = c0, r, treatment
	c0 = Node[J].newQual[p]
	r = 0
	treatment = nil
	if (R_[p] > float64(1.0)) || (ErrCode != 0) {
		ErrCode = 1
		return float64(0.0)
	}
	if (R_[p] >= float64(0.0)) && (R_[p] <= float64(1.0)) {
		return R_[p]
	}
	R_[p] = float64(10.0)
	if c0 == float64(0.0) {
		R_[p] = float64(0.0)
		return float64(0.0)
	}
	treatment = &Node[J].treatment[p]
	r = mathexpr_eval(treatment.equation, treatmnt_getVariableValue)
	if (float64(0.0)) >= (r) {
		r = (float64(0.0))
	} else {
		r = (r)
	}
	if treatment.treatType == REMOVAL {
		if (float64(1.0)) <= (r) {
			r = (float64(1.0))
		} else {
			r = (r)
		}
		R_[p] = r
	} else {
		if (c0) <= (r) {
			r = (c0)
		} else {
			r = (r)
		}
		R_[p] = float64(float64(1.0) - float64(r/c0))
	}
	return R_[p]
}
