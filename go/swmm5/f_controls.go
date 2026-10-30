// Generated from EPA SWMM 5.2.4 (controls.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var ObjectWords = [10]R.CPtr{L656, L657, L658, L659, L660, L661, L662, L663, L32, (R.CPtr{})}
var AttribWords = [22]R.CPtr{L664, L665, L314, L666, L667, L183, L668, L669, L670, L671, L672, L673, L674, L675, L676, L677, L21, L678, L679, L680, L681, (R.CPtr{})}
var RelOpWords = [7]R.CPtr{L31, L682, L683, L684, L685, L686, (R.CPtr{})}
var StatusWords = [3]R.CPtr{L687, L688, (R.CPtr{})}
var ConduitWords = [3]R.CPtr{L689, L690, (R.CPtr{})}
var SettingTypeWords = [4]R.CPtr{L691, L692, L693, (R.CPtr{})}
var IntensityWord R.CPtr = L694
var Rules []TRule
var ActionList *TActionList
var InputState int32
var RuleCount int32
var ControlValue float64
var SetPoint float64
var CurrentDate float64
var CurrentTime float64
var VariableCount int32
var ExpressionCount int32
var CurrentVariable int32
var CurrentExpression int32
var NamedVariable []TNamedVariable
var Expression []TExpression

// controls.c:226
func controls_init() {
	Rules = nil
	NamedVariable = nil
	Expression = nil
	RuleCount = 0
	VariableCount = 0
	ExpressionCount = 0
}

// controls.c:243
func controls_addToCount(s R.CPtr) {
	if match(s, L29) != 0 {
		VariableCount++
	} else if match(s, L30) != 0 {
		ExpressionCount++
	}
}

// controls.c:257
func controls_create(n_2 int32) int32 {
	var (
		r int32
	)
	_ = r
	r = 0
	ActionList = nil
	InputState = r_PRIORITY
	RuleCount = n_2
	if RuleCount > 0 {
		Rules = make([]TRule, int(uint64(RuleCount)))
		if Rules == nil {
			return ERR_MEMORY
		}
		for r = 0; r < RuleCount; r++ {
			Rules[r].ID = (R.CPtr{})
			Rules[r].firstPremise = nil
			Rules[r].lastPremise = nil
			Rules[r].thenActions = nil
			Rules[r].elseActions = nil
			Rules[r].priority = float64(0.0)
		}
	}
	CurrentVariable = (-1)
	CurrentExpression = (-1)
	if VariableCount > 0 {
		NamedVariable = make([]TNamedVariable, int(uint64(VariableCount)))
		if NamedVariable == nil {
			return ERR_MEMORY
		}
	}
	if ExpressionCount > 0 {
		Expression = make([]TExpression, int(uint64(ExpressionCount)))
		if Expression == nil {
			return ERR_MEMORY
		}
	}
	return 0
}

// controls.c:302
func controls_delete() {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 0; i < ExpressionCount; i++ {
		mathexpr_delete(Expression[i].expression)
		Expression[i].expression = nil
	}
	{
		if Expression != nil {
			_ = Expression
			Expression = nil
		}
	}
	{
		if NamedVariable != nil {
			_ = NamedVariable
			NamedVariable = nil
		}
	}
	if RuleCount == 0 {
		return
	}
	deleteActionList()
	deleteRules()
}

// controls.c:326
func controls_addVariable(tok []R.CPtr, nToks int32) int32 {
	var (
		v1  TVariable
		k   int32
		err int32
	)
	_, _, _ = v1, k, err
	v1 = (TVariable{})
	k = 0
	err = 0
	CurrentVariable++
	if nToks < 5 {
		return ERR_ITEMS
	}
	if findExactMatch(tok[1], AttribWords[:]) >= 0 {
		return error_setInpError(ERR_KEYWORD, tok[1])
	}
	if !(match(tok[2], L31) != 0) {
		return error_setInpError(ERR_KEYWORD, tok[2])
	}
	if (!(match(tok[3], L32) != 0)) && (nToks < 6) {
		return ERR_ITEMS
	}
	k = 3
	err = getPremiseVariable(tok, nToks, &k, &v1)
	if err > 0 {
		return err
	}
	k = CurrentVariable
	NamedVariable[k].variable = v1
	sstrncpy(R.CP(NamedVariable[k].name[:]), tok[1], 32)
	return 0
}

// controls.c:357
func controls_addExpression(tok []R.CPtr, nToks int32) int32 {
	var (
		i    int32
		k    int32
		s    R.CPtr
		expr *MathExpr
	)
	_, _, _, _ = i, k, s, expr
	i = 0
	k = 0
	s = R.Buf(1025)
	expr = nil
	CurrentExpression++
	if nToks < 4 {
		return ERR_ITEMS
	}
	k = CurrentExpression
	Expression[k].expression = nil
	sstrncpy(R.CP(Expression[k].name[:]), tok[1], 32)
	sstrncpy(s, tok[3], 1024)
	for i = 4; i < nToks; i++ {
		sstrcat(s, L33, 1024)
		sstrcat(s, tok[i], 1024)
	}
	expr = mathexpr_create(s, getVariableIndex)
	if expr == nil {
		return error_setInpError(ERR_MATH_EXPR, L0)
	}
	Expression[k].expression = expr
	return 0
}

// controls.c:393
func getVariableIndex(varName R.CPtr) int32 {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 0; i < VariableCount; i++ {
		if match(varName, R.CP(NamedVariable[i].name[:])) != 0 {
			return i
		}
	}
	return (-1)
}

// controls.c:410
func getNamedVariableValue(varIndex int32) float64 {
	return getVariableValue(NamedVariable[varIndex].variable)
}

// controls.c:422
func getExpressionIndex(exprName R.CPtr) int32 {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 0; i < ExpressionCount; i++ {
		if match(exprName, R.CP(Expression[i].name[:])) != 0 {
			return i
		}
	}
	return (-1)
}

// controls.c:439
func controls_addRuleClause(r int32, keyword int32, tok []R.CPtr, nToks int32) int32 {
	switch keyword {
	case r_RULE:
		if Rules[r].ID.IsNull() {
			Rules[r].ID = project_findID(CONTROL, tok[1])
		}
		InputState = r_RULE
		if nToks > 2 {
			return ERR_RULE
		}
		return 0
	case r_IF:
		if InputState != r_RULE {
			return ERR_RULE
		}
		InputState = r_IF
		return addPremise(r, r_AND, tok, nToks)
	case r_AND:
		if InputState == r_IF {
			return addPremise(r, r_AND, tok, nToks)
		} else if (InputState == r_THEN) || (InputState == r_ELSE) {
			return addAction(r, tok, nToks)
		} else {
			return ERR_RULE
		}
	case r_OR:
		if InputState != r_IF {
			return ERR_RULE
		}
		return addPremise(r, r_OR, tok, nToks)
	case r_THEN:
		if InputState != r_IF {
			return ERR_RULE
		}
		InputState = r_THEN
		return addAction(r, tok, nToks)
	case r_ELSE:
		if InputState != r_THEN {
			return ERR_RULE
		}
		InputState = r_ELSE
		return addAction(r, tok, nToks)
	case r_PRIORITY:
		if (InputState != r_THEN) && (InputState != r_ELSE) {
			return ERR_RULE
		}
		InputState = r_PRIORITY
		if !(getDouble(tok[1], &Rules[r].priority) != 0) {
			return ERR_NUMBER
		}
		if nToks > 2 {
			return ERR_RULE
		}
		return 0
	}
	return 0
}

// controls.c:495
func controls_evaluate(currentTime float64, elapsedTime float64, tStep float64) int32 {
	var (
		r      int32
		result int32
		p      *TPremise
		a_2    *TAction
	)
	_, _, _, _ = r, result, p, a_2
	r = 0
	result = 0
	p = nil
	a_2 = nil
	CurrentDate = M.Floor(currentTime)
	CurrentTime = float64(currentTime - M.Floor(currentTime))
	ElapsedTime = elapsedTime
	if RuleCount == 0 {
		return 0
	}
	clearActionList()
	for r = 0; r < RuleCount; r++ {
		result = 1
		p = Rules[r].firstPremise
		for p != nil {
			if p.type_ == r_OR {
				if result == 0 {
					result = evaluatePremise(p, tStep)
				}
			} else {
				if result == 0 {
					break
				}
				result = evaluatePremise(p, tStep)
			}
			p = p.next
		}
		if result == 1 {
			a_2 = Rules[r].thenActions
		} else {
			a_2 = Rules[r].elseActions
		}
		for a_2 != nil {
			updateActionValue(a_2, currentTime, tStep)
			updateActionList(a_2)
			a_2 = a_2.next
		}
	}
	if ActionList != nil {
		return executeActionList(currentTime)
	} else {
		return 0
	}
	return 0
}

// controls.c:556
func addPremise(r int32, type_2 int32, tok []R.CPtr, nToks int32) int32 {
	var (
		relation  int32
		n_2       int32
		err       int32
		value     float64
		p         *TPremise
		v1        TVariable
		v2        TVariable
		obj       int32
		exprIndex int32
		varIndex  int32
	)
	_, _, _, _, _, _, _, _, _, _ = relation, n_2, err, value, p, v1, v2, obj, exprIndex, varIndex
	relation = 0
	n_2 = 0
	err = 0
	value = (-(float64(10000000000.0)))
	p = nil
	v1 = (TVariable{})
	v2 = (TVariable{})
	obj = 0
	exprIndex = 0
	varIndex = (-1)
	if nToks < 4 {
		return ERR_ITEMS
	}
	v1.attribute = (-1)
	v1.object = (-1)
	v1.index = (-1)
	n_2 = 1
	exprIndex = getExpressionIndex(tok[1])
	if exprIndex < 0 {
		varIndex = getVariableIndex(tok[n_2])
		if varIndex >= 0 {
			v1 = NamedVariable[varIndex].variable
		} else {
			err = getPremiseVariable(tok, nToks, &n_2, &v1)
			if err > 0 {
				return err
			}
		}
	}
	n_2++
	if n_2 >= nToks {
		return error_setInpError(ERR_ITEMS, L0)
	}
	relation = findExactMatch(tok[n_2], RelOpWords[:])
	if relation < 0 {
		return error_setInpError(ERR_KEYWORD, tok[n_2])
	}
	v2.attribute = (-1)
	v2.object = (-1)
	v2.index = (-1)
	n_2++
	if n_2 >= nToks {
		return error_setInpError(ERR_ITEMS, L0)
	}
	varIndex = getVariableIndex(tok[n_2])
	if varIndex >= 0 {
		v2 = NamedVariable[varIndex].variable
	} else {
		obj = findmatch(tok[n_2], ObjectWords[:])
		if obj >= 0 {
			err = getPremiseVariable(tok, nToks, &n_2, &v2)
			if err > 0 {
				return ERR_RULE
			}
			if (exprIndex < 0) && (v1.attribute != v2.attribute) {
				report_writeWarningMsg(L34, Rules[r].ID)
			}
		} else {
			err = getPremiseValue(tok[n_2], v1.attribute, &value)
			if err > 0 {
				return err
			}
		}
	}
	n_2++
	if (n_2 < nToks) && (findmatch(tok[n_2], RuleKeyWords[:]) >= 0) {
		return ERR_RULE
	}
	p = new(TPremise)
	if !(p != nil) {
		return ERR_MEMORY
	}
	p.type_ = type_2
	p.exprIndex = exprIndex
	p.lhsVar = v1
	p.rhsVar = v2
	p.relation = relation
	p.value = value
	p.next = nil
	if Rules[r].firstPremise == nil {
		Rules[r].firstPremise = p
	} else {
		Rules[r].lastPremise.next = p
	}
	Rules[r].lastPremise = p
	return 0
}

// controls.c:668
func getPremiseVariable(tok []R.CPtr, nToks int32, k *int32, v *TVariable) int32 {
	var (
		n_2    int32
		object int32
		index  int32
		obj    int32
		attrib int32
	)
	_, _, _, _, _ = n_2, object, index, obj, attrib
	n_2 = (*k)
	object = (-1)
	index = (-1)
	obj = 0
	attrib = 0
	obj = findmatch(tok[n_2], ObjectWords[:])
	if obj < 0 {
		return error_setInpError(ERR_KEYWORD, tok[n_2])
	}
	n_2++
	if n_2 >= nToks {
		return error_setInpError(ERR_ITEMS, L0)
	}
	switch obj {
	case r_GAGE:
		index = project_findObject(GAGE, tok[n_2])
		if index < 0 {
			return error_setInpError(ERR_NAME, tok[n_2])
		}
		object = r_GAGE
		break
	case r_NODE:
		index = project_findObject(NODE, tok[n_2])
		if index < 0 {
			return error_setInpError(ERR_NAME, tok[n_2])
		}
		object = r_NODE
		break
	case r_LINK, r_CONDUIT, r_PUMP, r_ORIFICE, r_WEIR, r_OUTLET:
		index = project_findObject(LINK, tok[n_2])
		if index < 0 {
			return error_setInpError(ERR_NAME, tok[n_2])
		}
		object = r_LINK
		break
	default:
		n_2--
	}
	n_2++
	if n_2 >= nToks {
		return error_setInpError(ERR_ITEMS, L0)
	}
	if object == r_GAGE {
		attrib = getGageAttrib(tok[n_2])
	} else {
		attrib = findmatch(tok[n_2], AttribWords[:])
	}
	if attrib < 0 {
		return error_setInpError(ERR_KEYWORD, tok[n_2])
	}
	if obj == r_GAGE {
	} else if obj == r_NODE {
		switch attrib {
		case r_DEPTH, r_MAXDEPTH, r_HEAD, r_VOLUME, r_INFLOW:
			break
		default:
			return error_setInpError(ERR_KEYWORD, tok[n_2])
		}
	} else if ((object == r_LINK) && (index >= 0)) && ((attrib == r_TIMEOPEN) || (attrib == r_TIMECLOSED)) {
	} else if (obj == r_LINK) || (obj == r_CONDUIT) {
		switch attrib {
		case r_STATUS, r_DEPTH, r_FULLFLOW, r_FULLDEPTH, r_FLOW, r_LENGTH, r_SLOPE, r_VELOCITY:
			break
		default:
			return error_setInpError(ERR_KEYWORD, tok[n_2])
		}
	} else if obj == r_PUMP {
		switch attrib {
		case r_FLOW, r_SETTING, r_STATUS:
			break
		default:
			return error_setInpError(ERR_KEYWORD, tok[n_2])
		}
	} else if ((obj == r_ORIFICE) || (obj == r_WEIR)) || (obj == r_OUTLET) {
		switch attrib {
		case r_FLOW, r_SETTING:
			break
		default:
			return error_setInpError(ERR_KEYWORD, tok[n_2])
		}
	} else {
		switch attrib {
		case r_TIME, r_DATE, r_CLOCKTIME, r_DAY, r_MONTH, r_DAYOFYEAR:
			break
		default:
			return error_setInpError(ERR_KEYWORD, tok[n_2])
		}
	}
	v.object = object
	v.index = index
	v.attribute = attrib
	*k = n_2
	return 0
}

// controls.c:797
func getGageAttrib(token R.CPtr) int32 {
	var (
		attrib int32
	)
	_ = attrib
	attrib = 0
	if match(token, IntensityWord) != 0 {
		return 0
	}
	attrib = R.Atoi(token)
	if (attrib < 1) || (attrib > 48) {
		return (-1)
	}
	return attrib
}

// controls.c:824
func getPremiseValue(token R.CPtr, attrib int32, value *float64) int32 {
	var (
		strDate R.CPtr
	)
	_ = strDate
	strDate = R.Buf(25)
	switch attrib {
	case r_STATUS:
		*value = float64(findmatch(token, StatusWords[:]))
		if (*value) < float64(0.0) {
			*value = float64(findmatch(token, ConduitWords[:]))
		}
		if (*value) < float64(0.0) {
			return error_setInpError(ERR_KEYWORD, token)
		}
		break
	case r_TIME, r_CLOCKTIME, r_TIMEOPEN, r_TIMECLOSED:
		if !(datetime_strToTime(token, value) != 0) {
			return error_setInpError(ERR_DATETIME, token)
		}
		break
	case r_DATE:
		if !(datetime_strToDate(token, value) != 0) {
			return error_setInpError(ERR_DATETIME, token)
		}
		break
	case r_DAY:
		if !(getDouble(token, value) != 0) {
			return error_setInpError(ERR_NUMBER, token)
		}
		if ((*value) < float64(1.0)) || ((*value) > float64(7.0)) {
			return error_setInpError(ERR_DATETIME, token)
		}
		break
	case r_MONTH:
		if !(getDouble(token, value) != 0) {
			return error_setInpError(ERR_NUMBER, token)
		}
		if ((*value) < float64(1.0)) || ((*value) > float64(12.0)) {
			return error_setInpError(ERR_DATETIME, token)
		}
		break
	case r_DAYOFYEAR:
		sstrncpy(strDate, token, 6)
		sstrcat(strDate, L35, 25)
		if datetime_strToDate(strDate, value) != 0 {
			*value = float64(datetime_dayOfYear((*value)))
		} else if ((!(getDouble(token, value) != 0)) || ((*value) < float64(1))) || ((*value) > float64(365)) {
			return error_setInpError(ERR_DATETIME, token)
		}
		break
	default:
		if !(getDouble(token, value) != 0) {
			return error_setInpError(ERR_NUMBER, token)
		}
	}
	return 0
}

// controls.c:889
func addAction(r int32, tok []R.CPtr, nToks int32) int32 {
	var (
		obj     int32
		link    int32
		attrib  int32
		curve   int32
		tseries int32
		n_2     int32
		err     int32
		values  [3]float64
		a_2     *TAction
	)
	_, _, _, _, _, _, _, _, _ = obj, link, attrib, curve, tseries, n_2, err, values, a_2
	obj = 0
	link = 0
	attrib = 0
	curve = (-1)
	tseries = (-1)
	n_2 = 0
	err = 0
	values = [3]float64{float64(1.0), float64(0.0), float64(0.0)}
	a_2 = nil
	if nToks < 6 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	obj = findmatch(tok[1], ObjectWords[:])
	if (((((obj != r_LINK) && (obj != r_CONDUIT)) && (obj != r_PUMP)) && (obj != r_ORIFICE)) && (obj != r_WEIR)) && (obj != r_OUTLET) {
		return error_setInpError(ERR_KEYWORD, tok[1])
	}
	link = project_findObject(LINK, tok[2])
	if link < 0 {
		return error_setInpError(ERR_NAME, tok[2])
	}
	switch obj {
	case r_CONDUIT:
		if Link[link].type_ != CONDUIT {
			return error_setInpError(ERR_NAME, tok[2])
		}
		break
	case r_PUMP:
		if Link[link].type_ != PUMP {
			return error_setInpError(ERR_NAME, tok[2])
		}
		break
	case r_ORIFICE:
		if Link[link].type_ != ORIFICE {
			return error_setInpError(ERR_NAME, tok[2])
		}
		break
	case r_WEIR:
		if Link[link].type_ != WEIR {
			return error_setInpError(ERR_NAME, tok[2])
		}
		break
	case r_OUTLET:
		if Link[link].type_ != OUTLET {
			return error_setInpError(ERR_NAME, tok[2])
		}
		break
	}
	attrib = findmatch(tok[3], AttribWords[:])
	if attrib < 0 {
		return error_setInpError(ERR_KEYWORD, tok[3])
	}
	if obj == r_CONDUIT {
		if attrib == r_STATUS {
			values[0] = float64(findmatch(tok[5], ConduitWords[:]))
			if values[0] < float64(0.0) {
				return error_setInpError(ERR_KEYWORD, tok[5])
			}
		} else {
			return error_setInpError(ERR_KEYWORD, tok[3])
		}
	} else if obj == r_PUMP {
		if attrib == r_STATUS {
			values[0] = float64(findmatch(tok[5], StatusWords[:]))
			if values[0] < float64(0.0) {
				return error_setInpError(ERR_KEYWORD, tok[5])
			}
		} else if attrib == r_SETTING {
			err = setActionSetting(tok, nToks, &curve, &tseries, &attrib, values[:])
			if err > 0 {
				return err
			}
		} else {
			return error_setInpError(ERR_KEYWORD, tok[3])
		}
	} else if ((obj == r_ORIFICE) || (obj == r_WEIR)) || (obj == r_OUTLET) {
		if attrib == r_SETTING {
			err = setActionSetting(tok, nToks, &curve, &tseries, &attrib, values[:])
			if err > 0 {
				return err
			}
			if (attrib == r_SETTING) && ((values[0] < float64(0.0)) || (values[0] > float64(1.0))) {
				return error_setInpError(ERR_NUMBER, tok[5])
			}
		} else {
			return error_setInpError(ERR_KEYWORD, tok[3])
		}
	} else {
		return error_setInpError(ERR_KEYWORD, tok[1])
	}
	n_2 = 6
	if (curve >= 0) || (tseries >= 0) {
		n_2 = 7
	}
	if attrib == r_PID {
		n_2 = 9
	}
	if (n_2 < nToks) && (findmatch(tok[n_2], RuleKeyWords[:]) >= 0) {
		return ERR_RULE
	}
	a_2 = new(TAction)
	if !(a_2 != nil) {
		return ERR_MEMORY
	}
	a_2.rule = r
	a_2.link = link
	a_2.attribute = attrib
	a_2.curve = curve
	a_2.tseries = tseries
	a_2.value = values[0]
	if attrib == r_PID {
		a_2.kp = values[0]
		a_2.ki = values[1]
		a_2.kd = values[2]
		a_2.e1 = float64(0.0)
		a_2.e2 = float64(0.0)
	}
	if InputState == r_THEN {
		a_2.next = Rules[r].thenActions
		Rules[r].thenActions = a_2
	} else {
		a_2.next = Rules[r].elseActions
		Rules[r].elseActions = a_2
	}
	return 0
}

// controls.c:1028
func setActionSetting(tok []R.CPtr, nToks int32, curve *int32, tseries *int32, attrib *int32, values []float64) int32 {
	var (
		k int32
		m int32
	)
	_, _ = k, m
	k = 0
	m = 0
	if nToks < 6 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	k = findmatch(tok[5], SettingTypeWords[:])
	if (k >= 0) && (nToks < 7) {
		return error_setInpError(ERR_ITEMS, L0)
	}
	switch k {
	case r_CURVE:
		m = project_findObject(CURVE, tok[6])
		if m < 0 {
			return error_setInpError(ERR_NAME, tok[6])
		}
		*curve = m
		break
	case r_TIMESERIES:
		m = project_findObject(TSERIES, tok[6])
		if m < 0 {
			return error_setInpError(ERR_NAME, tok[6])
		}
		*tseries = m
		Tseries[m].refersTo = CONTROL
		break
	case r_PID:
		if nToks < 9 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		for m = 6; m <= 8; m++ {
			if !(getDouble(tok[m], &values[(m-6)]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[m])
			}
		}
		*attrib = r_PID
		break
	default:
		if !(getDouble(tok[5], &values[0]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[5])
		}
	}
	return 0
}

// controls.c:1086
func updateActionValue(a_2 *TAction, currentTime float64, dt float64) {
	if a_2.curve >= 0 {
		a_2.value = table_lookup(&Curve[a_2.curve], ControlValue)
	} else if a_2.tseries >= 0 {
		a_2.value = table_tseriesLookup(&Tseries[a_2.tseries], currentTime, 1)
	} else if a_2.attribute == r_PID {
		a_2.value = getPIDSetting(a_2, dt)
	}
}

// controls.c:1111
func getPIDSetting(a_2 *TAction, dt float64) float64 {
	var (
		e0        float64
		setting   float64
		p         float64
		i         float64
		d         float64
		update    float64
		tolerance float64
	)
	_, _, _, _, _, _, _ = e0, setting, p, i, d, update, tolerance
	e0 = 0
	setting = 0
	p = 0
	i = 0
	d = 0
	update = 0
	tolerance = float64(0.0001)
	dt = float64(dt * (float64(1440.0)))
	e0 = float64(SetPoint - ControlValue)
	if M.Fabs(e0) > float64(1e-06) {
		if SetPoint != float64(0.0) {
			e0 = float64(e0 / SetPoint)
		} else {
			e0 = float64(e0 / ControlValue)
		}
	}
	if M.Fabs(float64(e0-a_2.e1)) < tolerance {
		a_2.e2 = float64(0.0)
		a_2.e1 = float64(0.0)
	}
	p = (float64(e0 - a_2.e1))
	if a_2.ki == float64(0.0) {
		i = float64(0.0)
	} else {
		i = float64(float64(e0*dt) / a_2.ki)
	}
	d = float64(float64(a_2.kd*(float64(float64(e0-float64(float64(2.0)*a_2.e1))+a_2.e2))) / dt)
	update = float64(a_2.kp * (float64(float64(p+i) + d)))
	if M.Fabs(update) < tolerance {
		update = float64(0.0)
	}
	setting = float64(Link[a_2.link].targetSetting + update)
	a_2.e2 = a_2.e1
	a_2.e1 = e0
	if setting < float64(0.0) {
		setting = float64(0.0)
	}
	if (Link[a_2.link].type_ != PUMP) && (setting > float64(1.0)) {
		setting = float64(1.0)
	}
	return setting
}

// controls.c:1168
func updateActionList(a_2 *TAction) {
	var (
		listItem *TActionList
		a1       *TAction
		priority float64
	)
	_, _, _ = listItem, a1, priority
	listItem = nil
	a1 = nil
	priority = Rules[a_2.rule].priority
	listItem = ActionList
	for listItem != nil {
		a1 = listItem.action
		if !(a1 != nil) {
			break
		}
		if a1.link == a_2.link {
			if priority > Rules[a1.rule].priority {
				listItem.action = a_2
			}
			return
		}
		listItem = listItem.next
	}
	if !(listItem != nil) {
		listItem = new(TActionList)
		listItem.next = ActionList
		ActionList = listItem
	}
	listItem.action = a_2
}

// controls.c:1206
func executeActionList(currentTime float64) int32 {
	var (
		listItem *TActionList
		nextItem *TActionList
		a1       *TAction
		count    int32
	)
	_, _, _, _ = listItem, nextItem, a1, count
	listItem = nil
	nextItem = nil
	a1 = nil
	count = 0
	listItem = ActionList
	for listItem != nil {
		a1 = listItem.action
		if !(a1 != nil) {
			break
		}
		if a1.link >= 0 {
			if Link[a1.link].targetSetting != a1.value {
				Link[a1.link].targetSetting = a1.value
				if (((RptFlags.controls != 0) && (a1.curve < 0)) && (a1.tseries < 0)) && (a1.attribute != r_PID) {
					report_writeControlAction(currentTime, Link[a1.link].ID, a1.value, Rules[a1.rule].ID)
				}
				count++
			}
		}
		nextItem = listItem.next
		listItem = nextItem
	}
	return count
}

// controls.c:1243
func evaluatePremise(p *TPremise, tStep float64) int32 {
	var (
		lhsValue float64
		rhsValue float64
		result   int32
	)
	_, _, _ = lhsValue, rhsValue, result
	lhsValue = 0
	rhsValue = 0
	result = 0
	if p.exprIndex >= 0 {
		lhsValue = mathexpr_eval(Expression[p.exprIndex].expression, getNamedVariableValue)
	} else {
		lhsValue = getVariableValue(p.lhsVar)
	}
	if p.value == (-(float64(10000000000.0))) {
		rhsValue = getVariableValue(p.rhsVar)
	} else {
		rhsValue = p.value
	}
	if (lhsValue == (-(float64(10000000000.0)))) || (rhsValue == (-(float64(10000000000.0)))) {
		return 0
	}
	switch p.lhsVar.attribute {
	case r_TIME, r_CLOCKTIME:
		return compareTimes(lhsValue, p.relation, rhsValue, float64(tStep/float64(2.0)))
	case r_TIMEOPEN, r_TIMECLOSED:
		result = compareTimes(lhsValue, p.relation, rhsValue, float64(tStep/float64(2.0)))
		ControlValue = float64(lhsValue * float64(24.0))
		return result
	default:
		return compareValues(lhsValue, p.relation, rhsValue)
	}
	return 0
}

// controls.c:1286
func getVariableValue(v TVariable) float64 {
	var (
		i int32
		j int32
	)
	_, _ = i, j
	i = (-1)
	j = (-1)
	if v.object == r_GAGE {
		return getRainValue(v)
	}
	if v.object == r_NODE {
		i = v.index
	}
	if v.object == r_LINK {
		j = v.index
	}
	switch v.attribute {
	case r_TIME:
		return ElapsedTime
	case r_DATE:
		return CurrentDate
	case r_CLOCKTIME:
		return CurrentTime
	case r_DAY:
		return float64(datetime_dayOfWeek(CurrentDate))
	case r_MONTH:
		return float64(datetime_monthOfYear(CurrentDate))
	case r_DAYOFYEAR:
		return float64(datetime_dayOfYear(CurrentDate))
	case r_STATUS:
		if (j < 0) || ((Link[j].type_ != CONDUIT) && (Link[j].type_ != PUMP)) {
			return (-(float64(10000000000.0)))
		} else {
			return Link[j].setting
		}
	case r_SETTING:
		if (j < 0) || (((Link[j].type_ != PUMP) && (Link[j].type_ != ORIFICE)) && (Link[j].type_ != WEIR)) {
			return (-(float64(10000000000.0)))
		} else {
			return Link[j].setting
		}
	case r_FLOW:
		if j < 0 {
			return (-(float64(10000000000.0)))
		} else {
			return float64(float64(float64(Link[j].direction)*Link[j].newFlow) * UCF(FLOW))
		}
	case r_FULLFLOW, r_FULLDEPTH, r_VELOCITY, r_LENGTH, r_SLOPE:
		if j < 0 {
			return (-(float64(10000000000.0)))
		} else if Link[j].type_ != CONDUIT {
			return (-(float64(10000000000.0)))
		}
		switch v.attribute {
		case r_FULLFLOW:
			return float64(Link[j].qFull * UCF(FLOW))
		case r_FULLDEPTH:
			return float64(Link[j].xsect.yFull * UCF(LENGTH))
		case r_VELOCITY:
			return float64(link_getVelocity(j, Link[j].newFlow, Link[j].newDepth) * UCF(LENGTH))
		case r_LENGTH:
			return float64(Conduit[Link[j].subIndex].length * UCF(LENGTH))
		case r_SLOPE:
			return Conduit[Link[j].subIndex].slope
		default:
			return (-(float64(10000000000.0)))
		}
		fallthrough
	case r_DEPTH:
		if j >= 0 {
			return float64(Link[j].newDepth * UCF(LENGTH))
		} else if i >= 0 {
			return float64(Node[i].newDepth * UCF(LENGTH))
		} else {
			return (-(float64(10000000000.0)))
		}
	case r_MAXDEPTH:
		if i >= 0 {
			return float64(Node[i].fullDepth * UCF(LENGTH))
		} else {
			return (-(float64(10000000000.0)))
		}
	case r_HEAD:
		if i < 0 {
			return (-(float64(10000000000.0)))
		}
		return float64((float64(Node[i].newDepth + Node[i].invertElev)) * UCF(LENGTH))
	case r_VOLUME:
		if i < 0 {
			return (-(float64(10000000000.0)))
		}
		return (float64(Node[i].newVolume * UCF(VOLUME)))
	case r_INFLOW:
		if i < 0 {
			return (-(float64(10000000000.0)))
		} else {
			return float64(Node[i].newLatFlow * UCF(FLOW))
		}
	case r_TIMEOPEN:
		if j < 0 {
			return (-(float64(10000000000.0)))
		}
		if Link[j].setting <= float64(0.0) {
			return (-(float64(10000000000.0)))
		}
		return float64(float64(CurrentDate+CurrentTime) - Link[j].timeLastSet)
	case r_TIMECLOSED:
		if j < 0 {
			return (-(float64(10000000000.0)))
		}
		if Link[j].setting > float64(0.0) {
			return (-(float64(10000000000.0)))
		}
		return float64(float64(CurrentDate+CurrentTime) - Link[j].timeLastSet)
	default:
		return (-(float64(10000000000.0)))
	}
	return 0
}

// controls.c:1388
func getRainValue(v TVariable) float64 {
	if v.index < 0 {
		return (-(float64(10000000000.0)))
	} else if Gage[v.index].isUsed == 0 {
		return float64(0.0)
	} else if v.attribute == 0 {
		return Gage[v.index].rainfall
	} else {
		return gage_getPastRain(v.index, v.attribute)
	}
	return 0
}

// controls.c:1405
func compareTimes(lhsValue float64, relation int32, rhsValue float64, halfStep float64) int32 {
	if relation == EQ {
		if (lhsValue >= float64(rhsValue-halfStep)) && (lhsValue < float64(rhsValue+halfStep)) {
			return 1
		}
		return 0
	} else if relation == NE {
		if (lhsValue < float64(rhsValue-halfStep)) || (lhsValue >= float64(rhsValue+halfStep)) {
			return 1
		}
		return 0
	} else {
		return compareValues(lhsValue, relation, rhsValue)
	}
	return 0
}

// controls.c:1432
func compareValues(lhsValue float64, relation int32, rhsValue float64) int32 {
	SetPoint = rhsValue
	ControlValue = lhsValue
	switch relation {
	case EQ:
		if lhsValue == rhsValue {
			return 1
		}
		break
	case NE:
		if lhsValue != rhsValue {
			return 1
		}
		break
	case LT:
		if lhsValue < rhsValue {
			return 1
		}
		break
	case LE:
		if lhsValue <= rhsValue {
			return 1
		}
		break
	case GT:
		if lhsValue > rhsValue {
			return 1
		}
		break
	case GE:
		if lhsValue >= rhsValue {
			return 1
		}
		break
	}
	return 0
}

// controls.c:1455
func clearActionList() {
	var (
		listItem *TActionList
	)
	_ = listItem
	listItem = nil
	listItem = ActionList
	for listItem != nil {
		listItem.action = nil
		listItem = listItem.next
	}
}

// controls.c:1473
func deleteActionList() {
	var (
		listItem *TActionList
		nextItem *TActionList
	)
	_, _ = listItem, nextItem
	listItem = nil
	nextItem = nil
	listItem = ActionList
	for listItem != nil {
		nextItem = listItem.next
		_ = listItem
		listItem = nextItem
	}
	ActionList = nil
}

// controls.c:1494
func deleteRules() {
	var (
		p     *TPremise
		pnext *TPremise
		a_2   *TAction
		anext *TAction
		r     int32
	)
	_, _, _, _, _ = p, pnext, a_2, anext, r
	p = nil
	pnext = nil
	a_2 = nil
	anext = nil
	r = 0
	for r = 0; r < RuleCount; r++ {
		p = Rules[r].firstPremise
		for p != nil {
			pnext = p.next
			_ = p
			p = pnext
		}
		a_2 = Rules[r].thenActions
		for a_2 != nil {
			anext = a_2.next
			_ = a_2
			a_2 = anext
		}
		a_2 = Rules[r].elseActions
		for a_2 != nil {
			anext = a_2.next
			_ = a_2
			a_2 = anext
		}
	}
	{
		if Rules != nil {
			_ = Rules
			Rules = nil
		}
	}
	RuleCount = 0
}

// controls.c:1536
func findExactMatch(s R.CPtr, keyword []R.CPtr) int32 {
	var (
		i int32
	)
	_ = i
	i = 0
	for !keyword[i].IsNull() {
		if strcomp(s, keyword[i]) != 0 {
			return (i)
		}
		i++
	}
	return (-1)
}
