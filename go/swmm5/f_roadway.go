// Generated from EPA SWMM 5.2.4 (roadway.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	M "swmm5go/glibm"
)

var N_Cr_Low_Paved int32 = 4
var Cr_Low_Paved = [4][2]float64{[2]float64{float64(0.0), float64(2.85)}, [2]float64{float64(0.2), float64(2.95)}, [2]float64{float64(0.7), float64(3.03)}, [2]float64{float64(4.0), float64(3.05)}}
var N_Cr_Low_Gravel int32 = 8
var Cr_Low_Gravel = [8][2]float64{[2]float64{float64(0.0), float64(2.5)}, [2]float64{float64(0.5), float64(2.7)}, [2]float64{float64(1.0), float64(2.8)}, [2]float64{float64(1.5), float64(2.9)}, [2]float64{float64(2.0), float64(2.98)}, [2]float64{float64(2.5), float64(3.02)}, [2]float64{float64(3.0), float64(3.03)}, [2]float64{float64(4.0), float64(3.05)}}
var N_Cr_High_Paved int32 = 2
var Cr_High_Paved = [2][2]float64{[2]float64{float64(0.15), float64(3.05)}, [2]float64{float64(0.25), float64(3.10)}}
var N_Cr_High_Gravel int32 = 2
var Cr_High_Gravel = [2][2]float64{[2]float64{float64(0.15), float64(2.95)}, [2]float64{float64(0.30), float64(3.10)}}
var N_Kt_Paved int32 = 9
var Kt_Paved = [9][2]float64{[2]float64{float64(0.8), float64(1.0)}, [2]float64{float64(0.85), float64(0.98)}, [2]float64{float64(0.90), float64(0.92)}, [2]float64{float64(0.93), float64(0.85)}, [2]float64{float64(0.95), float64(0.80)}, [2]float64{float64(0.97), float64(0.70)}, [2]float64{float64(0.98), float64(0.60)}, [2]float64{float64(0.99), float64(0.50)}, [2]float64{float64(1.00), float64(0.40)}}
var N_Kt_Gravel int32 = 12
var Kt_Gravel = [12][2]float64{[2]float64{float64(0.75), float64(1.00)}, [2]float64{float64(0.80), float64(0.985)}, [2]float64{float64(0.83), float64(0.97)}, [2]float64{float64(0.86), float64(0.93)}, [2]float64{float64(0.89), float64(0.90)}, [2]float64{float64(0.90), float64(0.87)}, [2]float64{float64(0.92), float64(0.80)}, [2]float64{float64(0.94), float64(0.70)}, [2]float64{float64(0.96), float64(0.60)}, [2]float64{float64(0.98), float64(0.50)}, [2]float64{float64(0.99), float64(0.40)}, [2]float64{float64(1.00), float64(0.24)}}

// roadway.c:85
func roadway_getInflow(j int32, dir float64, hRoad float64, h1 float64, h2 float64) float64 {
	var (
		k             int32
		roadSurf      int32
		useVariableCd int32
		length        float64
		roadWidth     float64
		hWr           float64
		ht            float64
		cD            float64
		q             float64
		dqdh          float64
	)
	_, _, _, _, _, _, _, _, _, _ = k, roadSurf, useVariableCd, length, roadWidth, hWr, ht, cD, q, dqdh
	k = 0
	roadSurf = 0
	useVariableCd = 0
	length = 0
	roadWidth = 0
	hWr = 0
	ht = 0
	cD = 0
	q = float64(0.0)
	dqdh = float64(0.0)
	if Link[j].type_ != WEIR {
		return float64(0.0)
	}
	k = Link[j].subIndex
	roadWidth = Weir[k].roadWidth
	roadSurf = Weir[k].roadSurface
	cD = Weir[k].cDisch1
	if UnitSystem == SI {
		cD = float64(cD / float64(0.552))
	}
	useVariableCd = 0
	if (roadWidth > float64(0.0)) && (roadSurf >= 1) {
		useVariableCd = 1
	}
	hWr = float64(h1 - hRoad)
	ht = float64(h2 - hRoad)
	if hWr > float64(0.0001) {
		if useVariableCd != 0 {
			cD = getCd(hWr, ht, roadWidth, roadSurf)
		}
		length = Link[j].xsect.wMax
		q = float64(float64(cD*length) * M.Pow(hWr, float64(1.5)))
		dqdh = float64(float64(float64(1.5)*q) / hWr)
	}
	Link[j].dqdh = dqdh
	if (float64(h1 - hRoad)) >= (float64(0.0)) {
		Link[j].newDepth = (float64(h1 - hRoad))
	} else {
		Link[j].newDepth = (float64(0.0))
	}
	Link[j].flowClass = SUBCRITICAL
	if hRoad > h2 {
		if dir == float64(1.0) {
			Link[j].flowClass = DN_CRITICAL
		} else {
			Link[j].flowClass = UP_CRITICAL
		}
	}
	return float64(dir * q)
}

// roadway.c:146
func getCd(hWr float64, ht float64, roadWidth float64, roadSurf int32) float64 {
	var (
		kT  float64
		hL  float64
		htH float64
		cR  float64
	)
	_, _, _, _ = kT, hL, htH, cR
	kT = float64(1.0)
	hL = 0
	htH = 0
	cR = 0
	if hWr <= float64(0.0) {
		return float64(0.0)
	}
	hL = float64(hWr / roadWidth)
	if hL <= float64(0.15) {
		if roadSurf == PAVED {
			cR = getY(hWr, Cr_Low_Paved[:], N_Cr_Low_Paved)
		} else {
			cR = getY(hWr, Cr_Low_Gravel[:], N_Cr_Low_Gravel)
		}
	} else {
		if roadSurf == PAVED {
			cR = getY(hL, Cr_High_Paved[:], N_Cr_High_Paved)
		} else {
			cR = getY(hL, Cr_High_Gravel[:], N_Cr_High_Gravel)
		}
	}
	if ht > float64(0.0) {
		htH = float64(ht / hWr)
		if roadSurf == PAVED {
			kT = getY(htH, Kt_Paved[:], N_Kt_Paved)
		} else {
			kT = getY(htH, Kt_Gravel[:], N_Kt_Gravel)
		}
	}
	return float64(cR * kT)
}

// roadway.c:176
func getY(x float64, table [][2]float64, n_2 int32) float64 {
	var (
		i  int32
		x1 float64
		y1 float64
		dx float64
		dy float64
	)
	_, _, _, _, _ = i, x1, y1, dx, dy
	i = 0
	x1 = 0
	y1 = 0
	dx = 0
	dy = 0
	if x <= table[0][0] {
		return table[0][1]
	}
	if x >= (table[(n_2 - 1)])[0] {
		return (table[(n_2 - 1)])[1]
	}
	for i = 1; i < n_2; i++ {
		if x <= table[i][0] {
			x1 = (table[(i - 1)])[0]
			dx = float64(table[i][0] - x1)
			y1 = (table[(i - 1)])[1]
			dy = float64(table[i][1] - y1)
			return float64(y1 + float64(float64((float64(x-x1))*dy)/dx))
		}
	}
	return (table[(n_2 - 1)])[1]
}
