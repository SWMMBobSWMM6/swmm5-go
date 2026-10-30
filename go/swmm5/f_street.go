// Generated from EPA SWMM 5.2.4 (street.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

// street.c:29
func street_create(nStreets int32) int32 {
	Street = nil
	Nobjects[STREET] = 0
	Street = make([]TStreet, int(uint64(nStreets)))
	if Street == nil {
		return ERR_MEMORY
	}
	Nobjects[STREET] = nStreets
	return 0
}

// street.c:46
func street_delete() {
	{
		if Street != nil {
			_ = Street
			Street = nil
		}
	}
}

// street.c:58
func street_readParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		i      int32
		k      int32
		sides  int32
		x      [11]float64
		street *TStreet
	)
	_, _, _, _, _ = i, k, sides, x, street
	i = 0
	k = 0
	sides = 0
	x = ([11]float64{})
	street = nil
	if ntoks < 5 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	i = project_findObject(STREET, tok[0])
	if i < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	Street[i].ID = project_findID(STREET, tok[0])
	for k = 0; k <= 9; k++ {
		x[k] = float64(0.0)
	}
	for k = 1; k <= 4; k++ {
		if (!(getDouble(tok[k], &x[k]) != 0)) || (x[k] <= float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[k])
		}
	}
	if ntoks > 5 {
		if (!(getDouble(tok[5], &x[5]) != 0)) || (x[5] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[5])
		}
	}
	if ntoks > 6 {
		if (!(getDouble(tok[6], &x[6]) != 0)) || (x[6] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[6])
		}
	}
	sides = 2
	if ntoks > 7 {
		if ((!(getInt(tok[7], &sides) != 0)) || (sides < 1)) || (sides > 2) {
			return error_setInpError(ERR_NUMBER, tok[7])
		}
	}
	if ntoks > 8 {
		if (!(getDouble(tok[8], &x[8]) != 0)) || (x[8] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, tok[8])
		}
		if x[8] > float64(0.0) {
			if ntoks < 11 {
				return error_setInpError(ERR_ITEMS, L0)
			}
			for k = 9; k <= 10; k++ {
				if (!(getDouble(tok[k], &x[k]) != 0)) || (x[k] <= float64(0.0)) {
					return error_setInpError(ERR_NUMBER, tok[k])
				}
			}
		}
	}
	street = &Street[i]
	street.width = float64(x[1] / UCF(LENGTH))
	street.curbHeight = float64(x[2] / UCF(LENGTH))
	street.slope = float64(x[3] / float64(100.0))
	street.roughness = x[4]
	street.gutterDepression = float64(x[5] / UCF(LENGTH))
	street.gutterWidth = float64(x[6] / UCF(LENGTH))
	street.sides = sides
	street.backWidth = float64(x[8] / UCF(LENGTH))
	street.backSlope = float64(x[9] / float64(100.0))
	street.backRoughness = x[10]
	transect_createStreetTransect(street)
	return 0
}

// street.c:141
func street_getExtentFilled(link int32) float64 {
	var (
		k      int32
		t      int32
		filled float64
	)
	_, _, _ = k, t, filled
	k = 0
	t = 0
	filled = 0
	t = Link[link].xsect.transect
	if t < 0 {
		return float64(0.0)
	}
	if RouteModel == DW {
		if (Node[Link[link].node1].newDepth) >= (Node[Link[link].node2].newDepth) {
			filled = (Node[Link[link].node1].newDepth)
		} else {
			filled = (Node[Link[link].node2].newDepth)
		}
	} else {
		k = Link[link].subIndex
		if (Conduit[k].a1) >= (Conduit[k].a2) {
			filled = (Conduit[k].a1)
		} else {
			filled = (Conduit[k].a2)
		}
	}
	return filled
}
