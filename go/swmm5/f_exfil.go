// Generated from EPA SWMM 5.2.4 (exfil.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

// exfil.c:34
func exfil_readStorageParams(k int32, tok []R.CPtr, ntoks int32, n_2 int32) int32 {
	var (
		i int32
		x [3]float64
	)
	_, _ = i, x
	i = 0
	x = ([3]float64{})
	if ntoks == (n_2 + 1) {
		if !(getDouble(tok[n_2], &x[1]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[n_2])
		}
		x[0] = float64(0.0)
		x[2] = float64(0.0)
	} else if ntoks < (n_2 + 3) {
		return error_setInpError(ERR_ITEMS, L0)
	} else {
		for i = 0; i < 3; i++ {
			if !(getDouble(tok[(n_2+i)], &x[i]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[(n_2+i)])
			}
		}
	}
	if x[1] == float64(0.0) {
		return 0
	}
	return createStorageExfil(k, x[:])
}

// exfil.c:74
func exfil_initState(k int32) {
	var (
		i      int32
		a_2    float64
		alast  float64
		d      float64
		aCurve *TTable
		exfil  *TExfil
	)
	_, _, _, _, _, _ = i, a_2, alast, d, aCurve, exfil
	i = 0
	a_2 = 0
	alast = 0
	d = 0
	aCurve = nil
	exfil = Storage[k].exfil
	if exfil != nil {
		grnampt_initState(exfil.btmExfil)
		grnampt_initState(exfil.bankExfil)
		switch Storage[k].shape {
		case TABULAR:
			i = Storage[k].aCurve
			exfil.btmArea = float64(0.0)
			exfil.bankMinDepth = float64(0.0)
			exfil.bankMaxDepth = float64(0.0)
			exfil.bankMaxArea = float64(0.0)
			if i >= 0 {
				aCurve = &Curve[i]
				Storage[k].exfil.btmArea = table_lookupEx(aCurve, float64(0.0))
				table_getFirstEntry(aCurve, &d, &a_2)
				alast = a_2
				for table_getNextEntry(aCurve, &d, &a_2) != 0 {
					if a_2 < alast {
						break
					} else if a_2 > alast {
						exfil.bankMaxArea = a_2
						exfil.bankMaxDepth = d
					} else if exfil.bankMaxArea == float64(0.0) {
						exfil.bankMinDepth = d
					} else {
						break
					}
					alast = a_2
				}
				exfil.btmArea = float64(exfil.btmArea / (float64(UCF(LENGTH) * UCF(LENGTH))))
				exfil.bankMaxArea = float64(exfil.bankMaxArea / (float64(UCF(LENGTH) * UCF(LENGTH))))
				exfil.bankMinDepth = float64(exfil.bankMinDepth / (UCF(LENGTH)))
				exfil.bankMaxDepth = float64(exfil.bankMaxDepth / (UCF(LENGTH)))
			}
			break
		case FUNCTIONAL:
			exfil.btmArea = Storage[k].a0
			if Storage[k].a2 == float64(0.0) {
				exfil.btmArea = float64(exfil.btmArea + Storage[k].a1)
			}
			exfil.bankMinDepth = float64(0.0)
			exfil.bankMaxDepth = float64(10000000000.0)
			exfil.bankMaxArea = float64(10000000000.0)
			break
		case CYLINDRICAL, CONICAL, PYRAMIDAL:
			exfil.btmArea = Storage[k].a0
			exfil.bankMinDepth = float64(0.0)
			exfil.bankMaxDepth = float64(10000000000.0)
			exfil.bankMaxArea = float64(10000000000.0)
			break
		}
	}
}

// exfil.c:158
func exfil_getLoss(exfil *TExfil, tStep float64, depth float64, area float64) float64 {
	var (
		exfilRate float64
	)
	_ = exfilRate
	exfilRate = float64(0.0)
	if exfil.btmExfil.IMDmax == float64(0.0) {
		exfilRate = float64(exfil.btmExfil.Ks * Adjust.hydconFactor)
	} else {
		exfilRate = grnampt_getInfil(exfil.btmExfil, tStep, float64(0.0), depth, MOD_GREEN_AMPT)
	}
	exfilRate = float64(exfilRate * exfil.btmArea)
	if depth > exfil.bankMinDepth {
		area = float64((func() float64 {
			if (area) <= (exfil.bankMaxArea) {
				return (area)
			}
			return (exfil.bankMaxArea)
		}()) - exfil.btmArea)
		if area > float64(0.0) {
			if exfil.btmExfil.IMDmax == float64(0.0) {
				exfilRate = float64(exfilRate + (float64(float64(area*exfil.btmExfil.Ks) * Adjust.hydconFactor)))
			} else {
				if depth > exfil.bankMaxDepth {
					depth = float64(float64(depth-exfil.bankMaxDepth) + float64((float64(exfil.bankMaxDepth-exfil.bankMinDepth))/float64(2.0)))
				} else {
					depth = float64((float64(depth - exfil.bankMinDepth)) / float64(2.0))
				}
				exfilRate = float64(exfilRate + (float64(area * grnampt_getInfil(exfil.bankExfil, tStep, float64(0.0), depth, MOD_GREEN_AMPT))))
			}
		}
	}
	return exfilRate
}

// exfil.c:218
func createStorageExfil(k int32, x []float64) int32 {
	var (
		exfil *TExfil
	)
	_ = exfil
	exfil = nil
	exfil = Storage[k].exfil
	if exfil == nil {
		exfil = new(TExfil)
		if exfil == nil {
			return error_setInpError(ERR_MEMORY, L0)
		}
		Storage[k].exfil = exfil
		exfil.btmExfil = nil
		exfil.bankExfil = nil
		exfil.btmExfil = new(TGrnAmpt)
		if exfil.btmExfil == nil {
			return error_setInpError(ERR_MEMORY, L0)
		}
		exfil.bankExfil = new(TGrnAmpt)
		if exfil.bankExfil == nil {
			return error_setInpError(ERR_MEMORY, L0)
		}
	}
	if !(grnampt_setParams(exfil.btmExfil, x) != 0) {
		return error_setInpError(ERR_NUMBER, L0)
	}
	grnampt_setParams(exfil.bankExfil, x)
	return 0
}
