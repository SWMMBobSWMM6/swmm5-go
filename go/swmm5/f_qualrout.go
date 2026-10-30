// Generated from EPA SWMM 5.2.4 (qualrout.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var ZeroVolume float64 = float64(0.0353147)
var ZeroDepth float64 = float64(0.003281)

// qualrout.c:63
func qualrout_init() {
	var (
		i     int32
		p     int32
		isWet int32
		c     float64
	)
	_, _, _, _ = i, p, isWet, c
	i = 0
	p = 0
	isWet = 0
	c = 0
	for i = 0; i < Nobjects[NODE]; i++ {
		isWet = R.B2I((Node[i].newDepth > ZeroDepth))
		for p = 0; p < Nobjects[POLLUT]; p++ {
			c = float64(0.0)
			if isWet != 0 {
				c = Pollut[p].initConcen
			}
			Node[i].oldQual[p] = c
			Node[i].newQual[p] = c
		}
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		isWet = R.B2I((Link[i].newDepth > ZeroDepth))
		for p = 0; p < Nobjects[POLLUT]; p++ {
			c = float64(0.0)
			if isWet != 0 {
				c = Pollut[p].initConcen
			}
			Link[i].oldQual[p] = c
			Link[i].newQual[p] = c
		}
	}
}

// qualrout.c:100
func qualrout_execute(tStep float64) {
	var (
		i    int32
		j    int32
		qIn  float64
		vAvg float64
	)
	_, _, _, _ = i, j, qIn, vAvg
	i = 0
	j = 0
	qIn = 0
	vAvg = 0
	for i = 0; i < Nobjects[LINK]; i++ {
		findLinkMassFlow(i, tStep)
	}
	for j = 0; j < Nobjects[NODE]; j++ {
		Node[j].qualInflow = Node[j].inflow
		qIn = Node[j].qualInflow
		vAvg = float64((float64(Node[j].oldVolume + Node[j].newVolume)) / float64(2.0))
		if Node[j].treatment != nil {
			if qIn < float64(1e-10) {
				qIn = float64(0.0)
			}
			treatmnt_setInflow(qIn, Node[j].newQual)
		}
		if (Node[j].type_ == STORAGE) || (Node[j].oldVolume > ZeroVolume) {
			findStorageQual(j, tStep)
		} else {
			findNodeQual(j)
		}
		if Node[j].treatment != nil {
			treatmnt_treat(j, qIn, vAvg, tStep)
		}
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		findLinkQual(i, tStep)
	}
}

// qualrout.c:146
func getMixedQual(c float64, v1 float64, wIn float64, qIn float64, tStep float64) float64 {
	var (
		vIn  float64
		cIn  float64
		cMax float64
	)
	_, _, _ = vIn, cIn, cMax
	vIn = 0
	cIn = 0
	cMax = 0
	if qIn <= float64(1e-10) {
		return c
	}
	vIn = float64(qIn * tStep)
	cIn = float64(float64(wIn*tStep) / vIn)
	if (c) >= (cIn) {
		cMax = (c)
	} else {
		cMax = (cIn)
	}
	c = float64((float64(float64(c*v1) + float64(wIn*tStep))) / (float64(v1 + vIn)))
	if (c) <= (cMax) {
		c = (c)
	} else {
		c = (cMax)
	}
	if (c) >= (float64(0.0)) {
		c = (c)
	} else {
		c = (float64(0.0))
	}
	return c
}

// qualrout.c:179
func findLinkMassFlow(i int32, tStep float64) {
	var (
		j     int32
		p     int32
		qLink float64
		w     float64
	)
	_, _, _, _ = j, p, qLink, w
	j = 0
	p = 0
	qLink = 0
	w = 0
	qLink = Link[i].newFlow
	j = Link[i].node2
	if qLink < float64(0.0) {
		j = Link[i].node1
	}
	qLink = M.Fabs(qLink)
	if RouteModel != DW {
		qLink = float64(qLink - (inlet_capturedFlow(i)))
	}
	for p = 0; p < Nobjects[POLLUT]; p++ {
		w = float64(qLink * Link[i].oldQual[p])
		Node[j].newQual[p] = float64(Node[j].newQual[p] + w)
		Link[i].totalLoad[p] = float64(Link[i].totalLoad[p] + (float64(w * tStep)))
	}
	Node[j].qualInflow = float64(Node[j].qualInflow + qLink)
}

// qualrout.c:221
func findNodeQual(j int32) {
	var (
		p     int32
		qNode float64
	)
	_, _ = p, qNode
	p = 0
	qNode = 0
	qNode = Node[j].qualInflow
	if qNode > float64(1e-10) {
		for p = 0; p < Nobjects[POLLUT]; p++ {
			Node[j].newQual[p] = float64(Node[j].newQual[p] / qNode)
		}
	} else {
		for p = 0; p < Nobjects[POLLUT]; p++ {
			if Node[j].newDepth > ZeroDepth {
				Node[j].newQual[p] = Node[j].oldQual[p]
			} else {
				Node[j].newQual[p] = float64(0.0)
			}
		}
	}
}

// qualrout.c:253
func findLinkQual(i int32, tStep float64) {
	var (
		j       int32
		k       int32
		p       int32
		wIn     float64
		qIn     float64
		qSeep   float64
		v1      float64
		v2      float64
		c1      float64
		c2      float64
		vEvap   float64
		vLosses float64
		fEvap   float64
		barrels float64
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = j, k, p, wIn, qIn, qSeep, v1, v2, c1, c2, vEvap, vLosses
	_, _ = fEvap, barrels
	j = 0
	k = 0
	p = 0
	wIn = 0
	qIn = 0
	qSeep = 0
	v1 = 0
	v2 = 0
	c1 = 0
	c2 = 0
	vEvap = 0
	vLosses = 0
	fEvap = 0
	barrels = 0
	j = Link[i].node1
	if Link[i].newFlow < float64(0.0) {
		j = Link[i].node2
	}
	if (Link[i].type_ != CONDUIT) || (Link[i].xsect.type_ == DUMMY) {
		for p = 0; p < Nobjects[POLLUT]; p++ {
			Link[i].newQual[p] = Node[j].newQual[p]
		}
		return
	}
	k = Link[i].subIndex
	barrels = float64(Conduit[k].barrels)
	qIn = float64(M.Fabs(Conduit[k].q1) * barrels)
	qSeep = float64(Conduit[k].seepLossRate * barrels)
	vEvap = float64(float64(Conduit[k].evapLossRate*barrels) * tStep)
	v1 = Link[i].oldVolume
	v2 = Link[i].newVolume
	vLosses = float64(float64(qSeep*tStep) + vEvap)
	fEvap = float64(1.0)
	if (vEvap > float64(0.0)) && (v1 > ZeroVolume) {
		fEvap = float64(fEvap + (float64(vEvap / v1)))
	}
	if RouteModel == SF {
		findSFLinkQual(i, qSeep, fEvap, tStep)
		return
	}
	if RouteModel == DW {
		qIn = float64(qIn + float64((float64(float64(v2+vLosses)-v1))/tStep))
		if (qIn) >= (float64(0.0)) {
			qIn = (qIn)
		} else {
			qIn = (float64(0.0))
		}
	}
	for p = 0; p < Nobjects[POLLUT]; p++ {
		c1 = Link[i].oldQual[p]
		massbal_addSeepageLoss(p, float64(qSeep*c1))
		c1 = float64(c1 * fEvap)
		c2 = getReactedQual(p, c1, v1, tStep)
		wIn = float64(Node[j].newQual[p] * qIn)
		c2 = getMixedQual(c2, v1, wIn, qIn, tStep)
		if (v2 < ZeroVolume) || (Link[i].newDepth <= ZeroDepth) {
			massbal_addToFinalStorage(p, float64(c2*v2))
			c2 = float64(0.0)
		}
		Link[i].newQual[p] = c2
	}
}

// qualrout.c:357
func findSFLinkQual(i int32, qSeep float64, fEvap float64, tStep float64) {
	var (
		j        int32
		p        int32
		c1       float64
		c2       float64
		lossRate float64
	)
	_, _, _, _, _ = j, p, c1, c2, lossRate
	j = Link[i].node1
	p = 0
	c1 = 0
	c2 = 0
	lossRate = 0
	for p = 0; p < Nobjects[POLLUT]; p++ {
		c1 = Node[j].newQual[p]
		massbal_addSeepageLoss(p, float64(qSeep*c1))
		c1 = float64(c1 * fEvap)
		c2 = c1
		if Pollut[p].kDecay > float64(0.0) {
			c2 = float64(c1 * M.Exp(float64((-Pollut[p].kDecay)*tStep)))
			if (float64(0.0)) >= (c2) {
				c2 = (float64(0.0))
			} else {
				c2 = (c2)
			}
			lossRate = float64((float64(c1 - c2)) * Link[i].newFlow)
			massbal_addReactedMass(p, lossRate)
		}
		Link[i].newQual[p] = c2
	}
}

// qualrout.c:398
func findStorageQual(j int32, tStep float64) {
	var (
		p      int32
		k      int32
		qIn    float64
		wIn    float64
		v1     float64
		c1     float64
		c2     float64
		qExfil float64
		vEvap  float64
		fEvap  float64
	)
	_, _, _, _, _, _, _, _, _, _ = p, k, qIn, wIn, v1, c1, c2, qExfil, vEvap, fEvap
	p = 0
	k = 0
	qIn = 0
	wIn = 0
	v1 = 0
	c1 = 0
	c2 = 0
	qExfil = float64(0.0)
	vEvap = float64(0.0)
	fEvap = float64(1.0)
	qIn = Node[j].qualInflow
	v1 = Node[j].oldVolume
	if Node[j].type_ == STORAGE {
		updateHRT(j, Node[j].oldVolume, qIn, tStep)
		k = Node[j].subIndex
		qExfil = float64(Storage[k].exfilLoss / tStep)
		vEvap = Storage[k].evapLoss
		if (vEvap > float64(0.0)) && (v1 > ZeroVolume) {
			fEvap = float64(fEvap + (float64(vEvap / v1)))
		}
	}
	for p = 0; p < Nobjects[POLLUT]; p++ {
		c1 = Node[j].oldQual[p]
		massbal_addSeepageLoss(p, float64(qExfil*c1))
		c1 = float64(c1 * fEvap)
		if (Node[j].treatment == nil) || (Node[j].treatment[p].equation == nil) {
			c1 = getReactedQual(p, c1, v1, tStep)
		}
		wIn = Node[j].newQual[p]
		c2 = getMixedQual(c1, v1, wIn, qIn, tStep)
		if ((Node[j].newVolume <= ZeroVolume) || (Node[j].newDepth <= ZeroDepth)) && (qIn <= float64(1e-10)) {
			massbal_addToFinalStorage(p, float64(c2*Node[j].newVolume))
			c2 = float64(0.0)
		}
		Node[j].newQual[p] = c2
	}
}

// qualrout.c:478
func updateHRT(j int32, v float64, q float64, tStep float64) {
	var (
		k   int32
		hrt float64
	)
	_, _ = k, hrt
	k = Node[j].subIndex
	hrt = Storage[k].hrt
	if v < float64(1e-10) {
		hrt = float64(0.0)
	} else {
		hrt = float64(float64((float64(hrt+tStep))*v) / (float64(v + float64(q*tStep))))
	}
	if (hrt) >= (float64(0.0)) {
		Storage[k].hrt = (hrt)
	} else {
		Storage[k].hrt = (float64(0.0))
	}
}

// qualrout.c:498
func getReactedQual(p int32, c float64, v1 float64, tStep float64) float64 {
	var (
		c2       float64
		lossRate float64
		kDecay   float64
	)
	_, _, _ = c2, lossRate, kDecay
	c2 = 0
	lossRate = 0
	kDecay = Pollut[p].kDecay
	if kDecay == float64(0.0) {
		return c
	}
	c2 = float64(c * (float64(float64(1.0) - float64(kDecay*tStep))))
	if (float64(0.0)) >= (c2) {
		c2 = (float64(0.0))
	} else {
		c2 = (c2)
	}
	lossRate = float64(float64((float64(c-c2))*v1) / tStep)
	massbal_addReactedMass(p, lossRate)
	return c2
}
