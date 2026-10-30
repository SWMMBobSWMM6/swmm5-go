// Generated from EPA SWMM 5.2.4 (shape.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	M "swmm5go/glibm"
)

var Atotal float64
var Ptotal float64

// shape.c:42
func shape_validate(shape *TShape, curve *TTable) int32 {
	if !(computeShapeTables(shape, curve) != 0) {
		return 0
	}
	if !(normalizeShapeTables(shape) != 0) {
		return 0
	}
	return 1
}

// shape.c:64
func computeShapeTables(shape *TShape, curve *TTable) int32 {
	var (
		i     int32
		n_2   int32
		dy    float64
		y_2   float64
		y1    float64
		y2    float64
		w     float64
		w1    float64
		w2    float64
		yLast float64
		wLast float64
		wMax  float64
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = i, n_2, dy, y_2, y1, y2, w, w1, w2, yLast, wLast, wMax
	i = 0
	n_2 = 0
	dy = 0
	y_2 = 0
	y1 = 0
	y2 = 0
	w = 0
	w1 = 0
	w2 = 0
	yLast = 0
	wLast = 0
	wMax = 0
	if !(table_getFirstEntry(curve, &y1, &w1) != 0) {
		return 0
	}
	if ((y1 < float64(0.0)) || (y1 >= float64(1.0))) || (w1 < float64(0.0)) {
		return 0
	}
	wMax = w1
	if y1 != float64(0.0) {
		y2 = y1
		w2 = w1
		y1 = float64(0.0)
		w1 = float64(0.0)
	} else {
		if !(table_getNextEntry(curve, &y2, &w2) != 0) {
			return 0
		}
		if (y2 < y1) || (w2 < float64(0.0)) {
			return 0
		}
		if y2 > float64(1.0) {
			y2 = float64(1.0)
		}
		if w2 > wMax {
			wMax = w2
		}
	}
	shape.nTbl = 51
	n_2 = (shape.nTbl - 1)
	dy = float64(float64(1.0) / float64((n_2)))
	shape.areaTbl[0] = float64(0.0)
	shape.hradTbl[0] = float64(0.0)
	shape.widthTbl[0] = w1
	Ptotal = w1
	Atotal = float64(0.0)
	y_2 = float64(0.0)
	w = w1
	for i = 1; i <= n_2; i++ {
		yLast = y_2
		wLast = w
		y_2 = float64(y_2 + dy)
		if M.Fabs(float64(y_2-float64(1.0))) < float64(1e-06) {
			y_2 = float64(1.0)
		}
		if y_2 > y2 {
			if !(getNextInterval(curve, y_2, yLast, wLast, &y1, &y2, &w1, &w2, &wMax) != 0) {
				return 0
			}
			yLast = y1
			wLast = w1
		}
		w = shape_getWidth(y_2, y1, y2, w1, w2)
		Atotal = float64(Atotal + (shape_getArea(y_2, w, yLast, wLast)))
		Ptotal = float64(Ptotal + (getPerim(y_2, w, yLast, wLast)))
		if y_2 == float64(1.0) {
			Ptotal = float64(Ptotal + w2)
		}
		shape.widthTbl[i] = w
		shape.areaTbl[i] = Atotal
		if Ptotal > float64(0.0) {
			shape.hradTbl[i] = float64(Atotal / Ptotal)
		} else {
			shape.hradTbl[i] = float64(0.0)
		}
	}
	shape.aFull = shape.areaTbl[n_2]
	shape.rFull = shape.hradTbl[n_2]
	shape.wMax = wMax
	getSmax(shape)
	return 1
}

// shape.c:189
func getSmax(shape *TShape) {
	var (
		i   int32
		n_2 int32
		sf  float64
	)
	_, _, _ = i, n_2, sf
	i = 0
	n_2 = (shape.nTbl - 1)
	sf = 0
	shape.sMax = float64(0.0)
	shape.aMax = float64(0.0)
	for i = 1; i <= n_2; i++ {
		sf = float64(shape.areaTbl[i] * M.Pow(shape.hradTbl[i], float64(float64(2.0)/float64(3.0))))
		if sf > shape.sMax {
			shape.sMax = sf
			shape.aMax = shape.areaTbl[i]
		}
	}
}

// shape.c:216
func normalizeShapeTables(shape *TShape) int32 {
	var (
		i     int32
		n_2   int32
		aFull float64
		rFull float64
		wMax  float64
	)
	_, _, _, _, _ = i, n_2, aFull, rFull, wMax
	i = 0
	n_2 = (shape.nTbl - 1)
	aFull = shape.aFull
	rFull = shape.rFull
	wMax = shape.wMax
	if ((aFull == float64(0.0)) || (rFull == float64(0.0))) || (wMax == float64(0.0)) {
		return 0
	}
	for i = 0; i <= n_2; i++ {
		shape.areaTbl[i] = float64(shape.areaTbl[i] / aFull)
		shape.hradTbl[i] = float64(shape.hradTbl[i] / rFull)
		shape.widthTbl[i] = float64(shape.widthTbl[i] / wMax)
	}
	return 1
}

// shape.c:246
func getNextInterval(curve *TTable, y_2 float64, yLast float64, wLast float64, y1 *float64, y2 *float64, w1 *float64, w2 *float64, wMax *float64) int32 {
	for y_2 > (*y2) {
		if (*y2) > yLast {
			Atotal = float64(Atotal + (shape_getArea((*y2), (*w2), yLast, wLast)))
			Ptotal = float64(Ptotal + (getPerim((*y2), (*w2), yLast, wLast)))
			yLast = (*y2)
			wLast = (*w2)
		}
		*y1 = (*y2)
		*w1 = (*w2)
		if !(table_getNextEntry(curve, y2, w2) != 0) {
			*y2 = float64(1.0)
			return 1
		}
		if (*w2) > (*wMax) {
			*wMax = (*w2)
		}
		if ((*y2) < (*y1)) || ((*w2) < float64(0.0)) {
			return 0
		}
		if (*y2) > float64(1.0) {
			*y2 = float64(1.0)
		}
	}
	return 1
}

// shape.c:307
func shape_getWidth(y_2 float64, y1 float64, y2 float64, w1 float64, w2 float64) float64 {
	if y2 == y1 {
		return w2
	}
	return float64(w1 + float64(float64((float64(y_2-y1))/(float64(y2-y1)))*(float64(w2-w1))))
}

// shape.c:328
func shape_getArea(y_2 float64, w float64, y1 float64, w1 float64) float64 {
	var (
		wMin float64
		wMax float64
	)
	_, _ = wMin, wMax
	wMin = 0
	wMax = 0
	if w > w1 {
		wMin = w1
		wMax = w
	} else {
		wMin = w
		wMax = w1
	}
	return float64((float64(wMin + float64((float64(wMax-wMin))/float64(2.0)))) * (float64(y_2 - y1)))
}

// shape.c:353
func getPerim(y_2 float64, w float64, y1 float64, w1 float64) float64 {
	var (
		dy float64
		dw float64
	)
	_, _ = dy, dw
	dy = float64(y_2 - y1)
	dw = float64(M.Fabs(float64(w-w1)) / float64(2.0))
	return float64(float64(2.0) * M.Sqrt(float64(float64(dy*dy)+float64(dw*dw))))
}
