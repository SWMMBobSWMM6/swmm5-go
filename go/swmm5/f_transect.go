// Generated from EPA SWMM 5.2.4 (transect.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var Ntransects int32
var Nstations int32
var Station [1501]float64
var Elev [1501]float64
var Nleft float64
var Nright float64
var Nchannel float64
var Xleftbank float64
var Xrightbank float64
var Xfactor float64
var Yfactor float64
var Lfactor float64

// transect.c:70
func transect_create(n_2 int32) int32 {
	Ntransects = n_2
	if n_2 == 0 {
		return 0
	}
	Transect = make([]TTransect, int(uint64(Ntransects)))
	if Transect == nil {
		return ERR_MEMORY
	}
	Nchannel = float64(0.0)
	Nleft = float64(0.0)
	Nright = float64(0.0)
	Nstations = 0
	return 0
}

// transect.c:90
func transect_delete() {
	if Ntransects == 0 {
		return
	}
	{
		if Transect != nil {
			_ = Transect
			Transect = nil
		}
	}
	Ntransects = 0
}

// transect.c:104
func transect_readParams(count *int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		i       int32
		k       int32
		index   int32
		errcode int32
		x       [10]float64
		id      R.CPtr
	)
	_, _, _, _, _, _ = i, k, index, errcode, x, id
	i = 0
	k = 0
	index = (*count)
	errcode = 0
	x = ([10]float64{})
	id = (R.CPtr{})
	k = findmatch(tok[0], TransectKeyWords[:])
	if k < 0 {
		return error_setInpError(ERR_KEYWORD, tok[0])
	}
	switch k {
	case 0:
		transect_validate((index - 1))
		if ntoks < 4 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		for i = 1; i <= 3; i++ {
			if !(getDouble(tok[i], &x[i]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[i])
			}
		}
		return setManning(x[:])
	case 1:
		if ntoks < 10 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		id = project_findID(TRANSECT, tok[1])
		if id.IsNull() {
			return error_setInpError(ERR_NAME, tok[1])
		}
		for i = 2; i < 10; i++ {
			if !(getDouble(tok[i], &x[i]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[i])
			}
		}
		*count = (index + 1)
		return setParams(index, id, x[:])
	case 2:
		if ((ntoks - 1) % 2) > 0 {
			return error_setInpError(ERR_ITEMS, L0)
		}
		i = 1
		for i < ntoks {
			if !(getDouble(tok[i], &x[1]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[i])
			}
			if !(getDouble(tok[(i+1)], &x[2]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[(i+1)])
			}
			errcode = addStation(x[1], x[2])
			if errcode != 0 {
				return errcode
			}
			i += 2
		}
		return 0
	}
	return 0
}

// transect.c:194
func transect_validate(j int32) {
	var (
		i           int32
		ymin        float64
		ymax        float64
		oldNchannel float64
	)
	_, _, _, _ = i, ymin, ymax, oldNchannel
	i = 0
	ymin = 0
	ymax = 0
	oldNchannel = Nchannel
	if (j < 0) || (j >= Ntransects) {
		return
	}
	if Nstations < 2 {
		report_writeErrorMsg(ERR_TRANSECT_TOO_FEW, Transect[j].ID)
		return
	}
	if Nstations >= 1500 {
		report_writeErrorMsg(ERR_TRANSECT_TOO_MANY, Transect[j].ID)
		return
	}
	if Nchannel <= float64(0.0) {
		report_writeErrorMsg(ERR_TRANSECT_MANNING, Transect[j].ID)
		return
	}
	if Xleftbank > Xrightbank {
		report_writeErrorMsg(ERR_TRANSECT_OVERBANK, Transect[j].ID)
		return
	}
	Nchannel = float64(Nchannel * M.Sqrt(Lfactor))
	Transect[j].lengthFactor = Lfactor
	ymax = Elev[1]
	ymin = Elev[1]
	for i = 2; i <= Nstations; i++ {
		if (Elev[i]) >= (ymax) {
			ymax = (Elev[i])
		} else {
			ymax = (ymax)
		}
		if (Elev[i]) <= (ymin) {
			ymin = (Elev[i])
		} else {
			ymin = (ymin)
		}
	}
	if ymin >= ymax {
		report_writeErrorMsg(ERR_TRANSECT_NO_DEPTH, Transect[j].ID)
		return
	}
	Station[0] = Station[1]
	Elev[0] = ymax
	Nstations++
	Station[Nstations] = Station[(Nstations - 1)]
	Elev[Nstations] = Elev[0]
	Transect[j].nTbl = 51
	createTables(&Transect[j], ymin, ymax)
	Transect[j].roughness = oldNchannel
}

// transect.c:264
func createTables(transect *TTransect, ymin float64, ymax float64) {
	var (
		i     int32
		nLast int32
		dy    float64
		y_2   float64
	)
	_, _, _, _ = i, nLast, dy, y_2
	i = 0
	nLast = 0
	dy = 0
	y_2 = 0
	transect.yFull = float64(ymax - ymin)
	transect.wMax = float64(0.0)
	transect.areaTbl[0] = float64(0.0)
	transect.hradTbl[0] = float64(0.0)
	transect.widthTbl[0] = float64(0.0)
	dy = float64((float64(ymax - ymin)) / (float64(float64((transect.nTbl)) - float64(1))))
	y_2 = ymin
	for i = 1; i < transect.nTbl; i++ {
		y_2 = float64(y_2 + dy)
		transect.areaTbl[i] = float64(0.0)
		transect.hradTbl[i] = float64(0.0)
		transect.widthTbl[i] = float64(0.0)
		getGeometry(transect, i, y_2)
	}
	setMaxSectionFactor(transect)
	nLast = (transect.nTbl - 1)
	transect.aFull = transect.areaTbl[nLast]
	transect.rFull = transect.hradTbl[nLast]
	transect.wMax = transect.widthTbl[nLast]
	for i = 1; i <= nLast; i++ {
		transect.areaTbl[i] = float64(transect.areaTbl[i] / transect.aFull)
		transect.hradTbl[i] = float64(transect.hradTbl[i] / transect.rFull)
		transect.widthTbl[i] = float64(transect.widthTbl[i] / transect.wMax)
	}
	transect.widthTbl[0] = transect.widthTbl[1]
}

// transect.c:312
func setManning(n_2 []float64) int32 {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 1; i <= 3; i++ {
		if n_2[i] < float64(0.0) {
			return ERR_NUMBER
		}
	}
	if n_2[1] > float64(0.0) {
		Nleft = n_2[1]
	}
	if n_2[2] > float64(0.0) {
		Nright = n_2[2]
	}
	if n_2[3] > float64(0.0) {
		Nchannel = n_2[3]
	}
	if Nleft == float64(0.0) {
		Nleft = Nchannel
	}
	if Nright == float64(0.0) {
		Nright = Nchannel
	}
	return 0
}

// transect.c:334
func setParams(j int32, id R.CPtr, x []float64) int32 {
	if (j < 0) || (j >= Ntransects) {
		return ERR_NUMBER
	}
	Transect[j].ID = id
	Xleftbank = float64(x[3] / UCF(LENGTH))
	Xrightbank = float64(x[4] / UCF(LENGTH))
	Lfactor = x[7]
	if Lfactor == float64(0.0) {
		Lfactor = float64(1.0)
	}
	Xfactor = x[8]
	if Xfactor == float64(0.0) {
		Xfactor = float64(1.0)
	}
	Xleftbank = float64(Xleftbank * Xfactor)
	Xrightbank = float64(Xrightbank * Xfactor)
	Yfactor = float64(x[9] / UCF(LENGTH))
	Nstations = 0
	return 0
}

// transect.c:360
func addStation(y_2 float64, x float64) int32 {
	if Nstations < 0 {
		return ERR_TRANSECT_UNKNOWN
	}
	Nstations++
	if Nstations >= 1500 {
		return 0
	}
	Station[Nstations] = float64(float64(x*Xfactor) / UCF(LENGTH))
	Elev[Nstations] = float64((float64(y_2 + Yfactor)) / UCF(LENGTH))
	if (Nstations > 1) && (Station[Nstations] < Station[(Nstations-1)]) {
		return ERR_TRANSECT_SEQUENCE
	}
	return 0
}

// transect.c:388
func getGeometry(transect *TTransect, i int32, y_2 float64) {
	var (
		k        int32
		ylo      float64
		yhi      float64
		w        float64
		wp       float64
		wpSum    float64
		a_2      float64
		aSum     float64
		q        float64
		qSum     float64
		findFlow int32
	)
	_, _, _, _, _, _, _, _, _, _, _ = k, ylo, yhi, w, wp, wpSum, a_2, aSum, q, qSum, findFlow
	k = 0
	ylo = 0
	yhi = 0
	w = 0
	wp = 0
	wpSum = 0
	a_2 = 0
	aSum = 0
	q = 0
	qSum = 0
	findFlow = 0
	wpSum = float64(0.0)
	aSum = float64(0.0)
	qSum = float64(0.0)
	for k = 1; k <= Nstations; k++ {
		if Elev[(k-1)] >= Elev[k] {
			yhi = Elev[(k - 1)]
			ylo = Elev[k]
		} else {
			yhi = Elev[k]
			ylo = Elev[(k - 1)]
		}
		if ylo >= y_2 {
			continue
		}
		getSliceGeom(k, y_2, ylo, yhi, &w, &a_2, &wp)
		wpSum = float64(wpSum + wp)
		aSum = float64(aSum + a_2)
		transect.areaTbl[i] = float64(transect.areaTbl[i] + a_2)
		transect.widthTbl[i] = float64(transect.widthTbl[i] + w)
		if Elev[k] >= y_2 {
			findFlow = 1
		} else {
			findFlow = 0
		}
		q = getFlow(k, aSum, wpSum, findFlow)
		if q > float64(0.0) {
			qSum = float64(qSum + q)
			aSum = float64(0.0)
			wpSum = float64(0.0)
		}
	}
	aSum = transect.areaTbl[i]
	if aSum == float64(0.0) {
		transect.hradTbl[i] = transect.hradTbl[(i - 1)]
	} else {
		transect.hradTbl[i] = M.Pow(float64(float64(float64(qSum*Nchannel)/float64(1.49))/aSum), float64(1.5))
	}
}

// transect.c:468
func getSliceGeom(k int32, y_2 float64, ylo float64, yhi float64, w *float64, a_2 *float64, wp *float64) {
	var (
		width float64
		ratio float64
	)
	_, _ = width, ratio
	width = 0
	ratio = 0
	width = M.Fabs(float64(Station[k] - Station[(k-1)]))
	*w = width
	*wp = M.Sqrt(float64(float64(width*width) + float64((float64(yhi-ylo))*(float64(yhi-ylo)))))
	*a_2 = float64(0.0)
	if y_2 > yhi {
		*a_2 = float64(float64(width*(float64((float64(y_2-yhi))+(float64(y_2-ylo))))) / float64(2.0))
	} else if yhi > ylo {
		ratio = float64((float64(y_2 - ylo)) / (float64(yhi - ylo)))
		*a_2 = float64(float64(float64(float64(width*(float64(yhi-ylo)))/float64(2.0))*ratio) * ratio)
		*w = float64((*w) * ratio)
		*wp = float64((*wp) * ratio)
	}
}

// transect.c:519
func getFlow(k int32, a_2 float64, wp float64, findFlow int32) float64 {
	var (
		n_2 float64
	)
	_ = n_2
	n_2 = 0
	if findFlow == 0 {
		if k == (Nstations - 1) {
			findFlow = 1
		} else if Station[k] == Xleftbank {
			if (Nleft != Nchannel) && (Station[k] != Station[(k-1)]) {
				findFlow = 1
			}
		} else if Station[k] == Xrightbank {
			if (Nright != Nchannel) && (Station[k] != Station[(k+1)]) {
				findFlow = 1
			}
		}
	}
	if findFlow != 0 {
		n_2 = Nchannel
		if Station[(k-1)] < Xleftbank {
			n_2 = Nleft
		}
		if Station[k] > Xrightbank {
			n_2 = Nright
		}
		return float64(float64(float64(float64(1.486)/n_2)*a_2) * M.Pow(float64(a_2/wp), float64(float64(2.0)/float64(3.0))))
	}
	return float64(0.0)
}

// transect.c:570
func setMaxSectionFactor(transect *TTransect) {
	var (
		i  int32
		sf float64
	)
	_, _ = i, sf
	i = 0
	sf = 0
	transect.aMax = float64(0.0)
	transect.sMax = float64(0.0)
	for i = 1; i < transect.nTbl; i++ {
		sf = float64(transect.areaTbl[i] * M.Pow(transect.hradTbl[i], float64(float64(2.0)/float64(3.0))))
		if sf > transect.sMax {
			transect.sMax = sf
			transect.aMax = transect.areaTbl[i]
		}
	}
}

// transect.c:596
func transect_createStreetTransect(street *TStreet) {
	var (
		ymin float64
		ymax float64
		y1   float64
		y3   float64
		y4   float64
		w1   float64
		w2   float64
		w3   float64
		w4   float64
	)
	_, _, _, _, _, _, _, _, _ = ymin, ymax, y1, y3, y4, w1, w2, w3, w4
	ymin = 0
	ymax = 0
	y1 = 0
	y3 = 0
	y4 = 0
	w1 = 0
	w2 = 0
	w3 = 0
	w4 = 0
	ymin = float64(0.0)
	w1 = street.backWidth
	w2 = street.gutterWidth
	w3 = street.width
	w4 = float64(w3 - w2)
	y3 = float64(street.gutterDepression + float64(street.slope*w2))
	y1 = float64(street.curbHeight + street.gutterDepression)
	ymax = float64(float64(street.backSlope*street.backWidth) + y1)
	y4 = float64(y3 + float64(street.slope*w4))
	if (ymax) >= (y4) {
		ymax = (ymax)
	} else {
		ymax = (y4)
	}
	Station[0] = float64(0.0)
	Elev[0] = ymax
	Station[1] = w1
	Elev[1] = y1
	Station[2] = w1
	Elev[2] = float64(0.0)
	Station[3] = float64(w1 + w2)
	Elev[3] = y3
	Station[4] = float64(w1 + w3)
	Elev[4] = y4
	if street.sides == 1 {
		Station[5] = Station[4]
		Elev[5] = ymax
		Nstations = 5
		street.transect.nTbl = 51
	} else {
		Station[5] = float64(Station[4] + w4)
		Elev[5] = y3
		Station[6] = float64(Station[5] + w2)
		Elev[6] = float64(0.0)
		Station[7] = Station[6]
		Elev[7] = y1
		Station[8] = float64(Station[7] + w1)
		Elev[8] = ymax
		Nstations = 8
		street.transect.nTbl = 51
	}
	Nchannel = street.roughness
	if street.backWidth == float64(0.0) {
		Nleft = Nchannel
		Nright = Nchannel
		Xleftbank = Station[0]
		Xrightbank = Station[Nstations]
	} else {
		Nleft = street.backRoughness
		Nright = Nleft
		Xleftbank = Station[1]
		if street.sides == 2 {
			Xrightbank = Station[(Nstations - 1)]
		} else {
			Xrightbank = Station[Nstations]
		}
	}
	createTables(&street.transect, ymin, ymax)
	street.transect.roughness = street.roughness
}
