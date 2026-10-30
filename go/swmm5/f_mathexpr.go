// Generated from EPA SWMM 5.2.4 (mathexpr.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var Err int32
var Bc int32
var PrevLex int32
var CurLex int32
var Len int32
var Pos int32
var S R.CPtr
var Token = R.Buf(255)
var Ivar int32
var Fvalue float64
var MathFunc = [20]R.CPtr{L1004, L1005, L1006, L1007, L1008, L1009, L1010, L1011, L739, L1012, L1013, L1014, L1015, L1016, L1017, L1018, L1019, L1020, L1021, (R.CPtr{})}
var mathexpr_getVariableIndex FP3

// mathexpr.c:109
func sametext(s1 R.CPtr, s2 R.CPtr) int32 {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 0; R.Toupper(s1.G(i)) == R.Toupper(s2.G(i)); i++ {
		if (!(s1.G((i + 1)) != 0)) && (!(s2.G((i + 1)) != 0)) {
			return (1)
		}
	}
	return (0)
}

// mathexpr.c:130
func isDigit(c int32) int32 {
	if (c >= 49) && (c <= 57) {
		return 1
	}
	if c == 48 {
		return 1
	}
	return 0
}

// mathexpr.c:139
func isLetter(c int32) int32 {
	if (c >= 97) && (c <= 122) {
		return 1
	}
	if (c >= 65) && (c <= 90) {
		return 1
	}
	if c == 95 {
		return 1
	}
	return 0
}

// mathexpr.c:149
func getToken() {
	var (
		c R.CPtr
	)
	_ = c
	c = R.BufFrom(L33, 2)
	Token.S(0, 0)
	for (Pos <= Len) && ((isLetter(S.G(Pos)) != 0) || (isDigit(S.G(Pos)) != 0)) {
		c.S(0, S.G(Pos))
		R.Strcat(Token, c)
		Pos++
	}
	Pos--
}

// mathexpr.c:165
func getMathFunc() int32 {
	var (
		i int32
	)
	_ = i
	i = 0
	for !MathFunc[i].IsNull() {
		if sametext(MathFunc[i], Token) != 0 {
			return (i + 10)
		}
		i++
	}
	return (0)
}

// mathexpr.c:178
func getVariable() int32 {
	if !(mathexpr_getVariableIndex != nil) {
		return 0
	}
	Ivar = mathexpr_getVariableIndex(Token)
	if Ivar >= 0 {
		return 8
	}
	return 0
}

// mathexpr.c:188
func getNumber() float64 {
	var (
		c            R.CPtr
		sNumber      R.CPtr
		errflag      int32
		decimalCount int32
	)
	_, _, _, _ = c, sNumber, errflag, decimalCount
	c = R.BufFrom(L33, 2)
	sNumber = R.Buf(255)
	errflag = 0
	decimalCount = 0
	sNumber.S(0, 0)
	for (Pos < Len) && (isDigit(S.G(Pos)) != 0) {
		c.S(0, S.G(Pos))
		R.Strcat(sNumber, c)
		Pos++
	}
	if Pos < Len {
		if S.G(Pos) == 46 {
			decimalCount++
			if decimalCount > 1 {
				Err = 1
			}
			R.Strcat(sNumber, L315)
			Pos++
			for (Pos < Len) && (isDigit(S.G(Pos)) != 0) {
				c.S(0, S.G(Pos))
				R.Strcat(sNumber, c)
				Pos++
			}
		}
		if (Pos < Len) && ((S.G(Pos) == 101) || (S.G(Pos) == 69)) {
			R.Strcat(sNumber, L316)
			Pos++
			if Pos >= Len {
				errflag = 1
			} else {
				if (S.G(Pos) == 45) || (S.G(Pos) == 43) {
					c.S(0, S.G(Pos))
					R.Strcat(sNumber, c)
					Pos++
				}
				if (Pos >= Len) || (!(isDigit(S.G(Pos)) != 0)) {
					errflag = 1
				} else {
					for (Pos < Len) && (isDigit(S.G(Pos)) != 0) {
						c.S(0, S.G(Pos))
						R.Strcat(sNumber, c)
						Pos++
					}
				}
			}
		}
	}
	Pos--
	if errflag != 0 {
		return float64(0)
	} else {
		return R.Atof(sNumber)
	}
	return 0
}

// mathexpr.c:252
func getOperand() int32 {
	var (
		code int32
	)
	_ = code
	code = 0
	switch S.G(Pos) {
	case 40:
		code = 1
		break
	case 41:
		code = 2
		break
	case 43:
		code = 3
		break
	case 45:
		code = 4
		if ((Pos < (Len - 1)) && (isDigit(S.G((Pos + 1))) != 0)) && ((CurLex <= 6) || (CurLex == 31)) {
			Pos++
			Fvalue = (-(getNumber()))
			code = 7
		}
		break
	case 42:
		code = 5
		break
	case 47:
		code = 6
		break
	case 94:
		code = 31
		break
	default:
		code = 0
	}
	return code
}

// mathexpr.c:280
func getLex() int32 {
	var (
		n_2 int32
	)
	_ = n_2
	n_2 = 0
	for (Pos < Len) && (S.G(Pos) == 32) {
		Pos++
	}
	if Pos >= Len {
		return 0
	}
	n_2 = getOperand()
	if n_2 == 0 {
		if isLetter(S.G(Pos)) != 0 {
			getToken()
			n_2 = getMathFunc()
			if n_2 == 0 {
				n_2 = getVariable()
			}
		} else if (S.G(Pos) == 46) || (isDigit(S.G(Pos)) != 0) {
			n_2 = 7
			Fvalue = getNumber()
		}
	}
	Pos++
	PrevLex = CurLex
	CurLex = n_2
	return n_2
}

// mathexpr.c:314
func newNode() *ExprTree {
	var (
		node *ExprTree
	)
	_ = node
	node = nil
	node = new(ExprTree)
	if !(node != nil) {
		Err = 2
	} else {
		node.opcode = 0
		node.ivar = (-1)
		node.fvalue = float64(0.0)
		node.left = nil
		node.right = nil
	}
	return node
}

// mathexpr.c:332
func getSingleOp(lex *int32) *ExprTree {
	var (
		opcode int32
		left   *ExprTree
		node   *ExprTree
	)
	_, _, _ = opcode, left, node
	opcode = 0
	left = nil
	node = nil
	if (*lex) == 1 {
		Bc++
		left = getTree()
	} else {
		if (((*lex) < 7) || ((*lex) == 9)) || ((*lex) > 30) {
			Err = 1
			return nil
		}
		opcode = (*lex)
		if ((*lex) == 7) || ((*lex) == 8) {
			left = newNode()
			left.opcode = opcode
			if (*lex) == 7 {
				left.fvalue = Fvalue
			}
			if (*lex) == 8 {
				left.ivar = Ivar
			}
		} else {
			*lex = getLex()
			if (*lex) != 1 {
				Err = 1
				return nil
			}
			Bc++
			left = newNode()
			left.left = getTree()
			left.opcode = opcode
		}
	}
	*lex = getLex()
	if (*lex) == 31 {
		node = newNode()
		node.left = left
		node.opcode = (*lex)
		*lex = getLex()
		node.right = getSingleOp(lex)
		left = node
	}
	return left
}

// mathexpr.c:397
func getOp(lex *int32) *ExprTree {
	var (
		opcode int32
		left   *ExprTree
		right  *ExprTree
		node   *ExprTree
		neg    int32
	)
	_, _, _, _, _ = opcode, left, right, node, neg
	opcode = 0
	left = nil
	right = nil
	node = nil
	neg = 0
	*lex = getLex()
	if (PrevLex == 0) || (PrevLex == 1) {
		if (*lex) == 4 {
			neg = 1
			*lex = getLex()
		} else if (*lex) == 3 {
			*lex = getLex()
		}
	}
	left = getSingleOp(lex)
	for ((*lex) == 5) || ((*lex) == 6) {
		opcode = (*lex)
		*lex = getLex()
		right = getSingleOp(lex)
		node = newNode()
		if Err != 0 {
			return nil
		}
		node.left = left
		node.right = right
		node.opcode = opcode
		left = node
	}
	if neg != 0 {
		node = newNode()
		if Err != 0 {
			return nil
		}
		node.left = left
		node.right = nil
		node.opcode = 9
		left = node
	}
	return left
}

// mathexpr.c:442
func getTree() *ExprTree {
	var (
		lex    int32
		opcode int32
		left   *ExprTree
		right  *ExprTree
		node   *ExprTree
	)
	_, _, _, _, _ = lex, opcode, left, right, node
	lex = 0
	opcode = 0
	left = nil
	right = nil
	node = nil
	left = getOp(&lex)
	for {
		if (lex == 0) || (lex == 2) {
			if lex == 2 {
				Bc--
			}
			break
		}
		if (lex != 3) && (lex != 4) {
			Err = 1
			break
		}
		opcode = lex
		right = getOp(&lex)
		node = newNode()
		if Err != 0 {
			break
		}
		node.left = left
		node.right = right
		node.opcode = opcode
		left = node
	}
	return left
}

// mathexpr.c:479
func traverseTree(tree *ExprTree, expr **MathExpr) {
	var (
		node *MathExpr
	)
	_ = node
	node = nil
	if tree == nil {
		return
	}
	traverseTree(tree.left, expr)
	traverseTree(tree.right, expr)
	node = new(MathExpr)
	if node != nil {
		node.fvalue = tree.fvalue
		node.opcode = tree.opcode
		node.ivar = tree.ivar
		node.next = nil
		node.prev = (*expr)
	}
	if (*expr) != nil {
		(*expr).next = node
	}
	*expr = node
}

// mathexpr.c:501
func deleteTree(tree *ExprTree) {
	if tree != nil {
		if tree.left != nil {
			deleteTree(tree.left)
		}
		if tree.right != nil {
			deleteTree(tree.right)
		}
		_ = tree
	}
}

// mathexpr.c:516
func mathexpr_eval(expr *MathExpr, getVariableValue_2 FP4) float64 {
	var (
		ExprStack  [1024]float64
		node       *MathExpr
		r1         float64
		r2         float64
		stackindex int32
	)
	_, _, _, _, _ = ExprStack, node, r1, r2, stackindex
	ExprStack = ([1024]float64{})
	node = expr
	r1 = 0
	r2 = 0
	stackindex = 0
	ExprStack[0] = float64(0.0)
	for (node != nil) && (stackindex >= 0) {
		switch node.opcode {
		case 3:
			r1 = ExprStack[stackindex]
			stackindex--
			if stackindex < 0 {
				break
			}
			r2 = ExprStack[stackindex]
			ExprStack[stackindex] = float64(r2 + r1)
			break
		case 4:
			r1 = ExprStack[stackindex]
			stackindex--
			if stackindex < 0 {
				break
			}
			r2 = ExprStack[stackindex]
			ExprStack[stackindex] = float64(r2 - r1)
			break
		case 5:
			r1 = ExprStack[stackindex]
			stackindex--
			if stackindex < 0 {
				break
			}
			r2 = ExprStack[stackindex]
			ExprStack[stackindex] = float64(r2 * r1)
			break
		case 6:
			r1 = ExprStack[stackindex]
			stackindex--
			if stackindex < 0 {
				break
			}
			r2 = ExprStack[stackindex]
			ExprStack[stackindex] = float64(r2 / r1)
			break
		case 7:
			stackindex++
			if stackindex >= 1024 {
				break
			}
			ExprStack[stackindex] = node.fvalue
			break
		case 8:
			if getVariableValue_2 != nil {
				r1 = getVariableValue_2(node.ivar)
			} else {
				r1 = float64(0.0)
			}
			stackindex++
			if stackindex >= 1024 {
				break
			}
			ExprStack[stackindex] = r1
			break
		case 9:
			ExprStack[stackindex] = (-ExprStack[stackindex])
			break
		case 10:
			r1 = ExprStack[stackindex]
			r2 = M.Cos(r1)
			ExprStack[stackindex] = r2
			break
		case 11:
			r1 = ExprStack[stackindex]
			r2 = M.Sin(r1)
			ExprStack[stackindex] = r2
			break
		case 12:
			r1 = ExprStack[stackindex]
			r2 = M.Tan(r1)
			ExprStack[stackindex] = r2
			break
		case 13:
			r1 = ExprStack[stackindex]
			if r1 == float64(0.0) {
				r2 = float64(0.0)
			} else {
				r2 = float64(float64(1.0) / M.Tan(r1))
			}
			ExprStack[stackindex] = r2
			break
		case 14:
			r1 = ExprStack[stackindex]
			r2 = M.Fabs(r1)
			ExprStack[stackindex] = r2
			break
		case 15:
			r1 = ExprStack[stackindex]
			if r1 < float64(0.0) {
				r2 = (-(float64(1.0)))
			} else if r1 > float64(0.0) {
				r2 = float64(1.0)
			} else {
				r2 = float64(0.0)
			}
			ExprStack[stackindex] = r2
			break
		case 16:
			r1 = ExprStack[stackindex]
			if r1 < float64(0.0) {
				r2 = float64(0.0)
			} else {
				r2 = M.Sqrt(r1)
			}
			ExprStack[stackindex] = r2
			break
		case 17:
			r1 = ExprStack[stackindex]
			if r1 <= float64(0) {
				r2 = float64(0.0)
			} else {
				r2 = M.Log(r1)
			}
			ExprStack[stackindex] = r2
			break
		case 18:
			r1 = ExprStack[stackindex]
			r2 = M.Exp(r1)
			ExprStack[stackindex] = r2
			break
		case 19:
			r1 = ExprStack[stackindex]
			r2 = M.Asin(r1)
			ExprStack[stackindex] = r2
			break
		case 20:
			r1 = ExprStack[stackindex]
			r2 = M.Acos(r1)
			ExprStack[stackindex] = r2
			break
		case 21:
			r1 = ExprStack[stackindex]
			r2 = M.Atan(r1)
			ExprStack[stackindex] = r2
			break
		case 22:
			r1 = ExprStack[stackindex]
			r2 = float64(float64(1.57079632679489661923) - M.Atan(r1))
			ExprStack[stackindex] = r2
			break
		case 23:
			r1 = ExprStack[stackindex]
			r2 = float64((float64(M.Exp(r1) - M.Exp((-r1)))) / float64(2.0))
			ExprStack[stackindex] = r2
			break
		case 24:
			r1 = ExprStack[stackindex]
			r2 = float64((float64(M.Exp(r1) + M.Exp((-r1)))) / float64(2.0))
			ExprStack[stackindex] = r2
			break
		case 25:
			r1 = ExprStack[stackindex]
			r2 = float64((float64(M.Exp(r1) - M.Exp((-r1)))) / (float64(M.Exp(r1) + M.Exp((-r1)))))
			ExprStack[stackindex] = r2
			break
		case 26:
			r1 = ExprStack[stackindex]
			r2 = float64((float64(M.Exp(r1) + M.Exp((-r1)))) / (float64(M.Exp(r1) - M.Exp((-r1)))))
			ExprStack[stackindex] = r2
			break
		case 27:
			r1 = ExprStack[stackindex]
			if r1 == float64(0.0) {
				r2 = float64(0.0)
			} else {
				r2 = M.Log10(r1)
			}
			ExprStack[stackindex] = r2
			break
		case 28:
			r1 = ExprStack[stackindex]
			if r1 <= float64(0.0) {
				r2 = float64(0.0)
			} else {
				r2 = float64(1.0)
			}
			ExprStack[stackindex] = r2
			break
		case 31:
			r1 = ExprStack[stackindex]
			stackindex--
			if stackindex < 0 {
				break
			}
			r2 = ExprStack[stackindex]
			if r2 <= float64(0.0) {
				r2 = float64(0.0)
			} else {
				r2 = M.Pow(r2, r1)
			}
			ExprStack[stackindex] = r2
			break
		}
		node = node.next
	}
	if stackindex >= 0 {
		r1 = ExprStack[stackindex]
	} else {
		r1 = float64(0.0)
	}
	if r1 != r1 {
		r1 = float64(0.0)
	}
	return r1
}

// mathexpr.c:735
func mathexpr_delete(expr *MathExpr) {
	if expr != nil {
		mathexpr_delete(expr.next)
	}
	_ = expr
}

// mathexpr.c:743
func mathexpr_create(formula R.CPtr, getVar FP3) *MathExpr {
	var (
		tree   *ExprTree
		expr   *MathExpr
		result *MathExpr
	)
	_, _, _ = tree, expr, result
	tree = nil
	expr = nil
	result = nil
	mathexpr_getVariableIndex = getVar
	Err = 0
	PrevLex = 0
	CurLex = 0
	S = formula
	Len = int32(R.Strlen(S))
	Pos = 0
	Bc = 0
	tree = getTree()
	if (Bc == 0) && (Err == 0) {
		traverseTree(tree, &expr)
		for expr != nil {
			result = expr
			expr = expr.prev
		}
	}
	deleteTree(tree)
	return result
}
