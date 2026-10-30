// Generated from EPA SWMM 5.2.4 (hotstart.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

var fileVersion int32

// hotstart.c:72
func hotstart_open() int32 {
	if !(openHotstartFile1() != 0) {
		return 0
	}
	if !(openHotstartFile2() != 0) {
		return 0
	}
	return 1
}

// hotstart.c:82
func hotstart_close() {
	if Fhotstart2.file != nil {
		saveRunoff()
		saveRouting()
		R.Fclose(Fhotstart2.file)
	}
}

// hotstart.c:94
func openHotstartFile1() int32 {
	var (
		nSubcatch  int32
		nNodes     int32
		nLinks     int32
		nPollut    int32
		nLandUses  int32
		flowUnits  int32
		fStamp     R.CPtr
		fileStamp  R.CPtr
		fStampx    R.CPtr
		fileStamp2 R.CPtr
		fileStamp3 R.CPtr
		fileStamp4 R.CPtr
	)
	_, _, _, _, _, _, _, _, _, _, _, _ = nSubcatch, nNodes, nLinks, nPollut, nLandUses, flowUnits, fStamp, fileStamp, fStampx, fileStamp2, fileStamp3, fileStamp4
	nSubcatch = 0
	nNodes = 0
	nLinks = 0
	nPollut = 0
	nLandUses = 0
	flowUnits = 0
	fStamp = R.BufFrom(L160, 15)
	fileStamp = R.BufFrom(L160, 15)
	fStampx = R.BufFrom(L161, 16)
	fileStamp2 = R.BufFrom(L162, 16)
	fileStamp3 = R.BufFrom(L163, 16)
	fileStamp4 = R.BufFrom(L164, 16)
	if Fhotstart1.mode != USE_FILE {
		return 1
	}
	if (R.Set(&Fhotstart1.file, R.Fopen(R.CP(Fhotstart1.name[:]), L165))) == nil {
		report_writeErrorMsg(ERR_HOTSTART_FILE_OPEN, R.CP(Fhotstart1.name[:]))
		return 0
	}
	R.FreadC(fStampx, R.Strlen(fileStamp2), Fhotstart1.file)
	if R.Strcmp(fStampx, fileStamp4) == 0 {
		fileVersion = 4
	} else if R.Strcmp(fStampx, fileStamp3) == 0 {
		fileVersion = 3
	} else if R.Strcmp(fStampx, fileStamp2) == 0 {
		fileVersion = 2
	} else {
		R.Rewind(Fhotstart1.file)
		R.FreadC(fStamp, R.Strlen(fileStamp), Fhotstart1.file)
		if R.Strcmp(fStamp, fileStamp) != 0 {
			report_writeErrorMsg(ERR_HOTSTART_FILE_FORMAT, L0)
			return 0
		}
		fileVersion = 1
	}
	nSubcatch = (-1)
	nNodes = (-1)
	nLinks = (-1)
	nPollut = (-1)
	nLandUses = (-1)
	flowUnits = (-1)
	if fileVersion >= 2 {
		R.Fread1(&nSubcatch, 4, Fhotstart1.file)
	} else {
		nSubcatch = Nobjects[SUBCATCH]
	}
	if fileVersion >= 3 {
		R.Fread1(&nLandUses, 4, Fhotstart1.file)
	} else {
		nLandUses = Nobjects[LANDUSE]
	}
	R.Fread1(&nNodes, 4, Fhotstart1.file)
	R.Fread1(&nLinks, 4, Fhotstart1.file)
	R.Fread1(&nPollut, 4, Fhotstart1.file)
	R.Fread1(&flowUnits, 4, Fhotstart1.file)
	if (((((nSubcatch != Nobjects[SUBCATCH]) || (nLandUses != Nobjects[LANDUSE])) || (nNodes != Nobjects[NODE])) || (nLinks != Nobjects[LINK])) || (nPollut != Nobjects[POLLUT])) || (flowUnits != FlowUnits) {
		report_writeErrorMsg(ERR_HOTSTART_FILE_FORMAT, L0)
		return 0
	}
	if fileVersion >= 3 {
		readRunoff()
	}
	readRouting()
	R.Fclose(Fhotstart1.file)
	if ErrorCode != 0 {
		return 0
	} else {
		return 1
	}
	return 0
}

// hotstart.c:180
func openHotstartFile2() int32 {
	var (
		nSubcatch int32
		nLandUses int32
		nNodes    int32
		nLinks    int32
		nPollut   int32
		flowUnits int32
		fileStamp R.CPtr
	)
	_, _, _, _, _, _, _ = nSubcatch, nLandUses, nNodes, nLinks, nPollut, flowUnits, fileStamp
	nSubcatch = 0
	nLandUses = 0
	nNodes = 0
	nLinks = 0
	nPollut = 0
	flowUnits = 0
	fileStamp = R.BufFrom(L164, 16)
	if Fhotstart2.mode != SAVE_FILE {
		return 1
	}
	if (R.Set(&Fhotstart2.file, R.Fopen(R.CP(Fhotstart2.name[:]), L166))) == nil {
		report_writeErrorMsg(ERR_HOTSTART_FILE_OPEN, R.CP(Fhotstart2.name[:]))
		return 0
	}
	nSubcatch = Nobjects[SUBCATCH]
	nLandUses = Nobjects[LANDUSE]
	nNodes = Nobjects[NODE]
	nLinks = Nobjects[LINK]
	nPollut = Nobjects[POLLUT]
	flowUnits = FlowUnits
	R.FwriteC(fileStamp, R.Strlen(fileStamp), Fhotstart2.file)
	R.Fwrite1(nSubcatch, 4, Fhotstart2.file)
	R.Fwrite1(nLandUses, 4, Fhotstart2.file)
	R.Fwrite1(nNodes, 4, Fhotstart2.file)
	R.Fwrite1(nLinks, 4, Fhotstart2.file)
	R.Fwrite1(nPollut, 4, Fhotstart2.file)
	R.Fwrite1(flowUnits, 4, Fhotstart2.file)
	return 1
}

// hotstart.c:222
func saveRouting() {
	var (
		i int32
		j int32
		x [3]float32
	)
	_, _, _ = i, j, x
	i = 0
	j = 0
	x = ([3]float32{})
	for i = 0; i < Nobjects[NODE]; i++ {
		x[0] = float32(float64(Node[i].newDepth))
		x[1] = float32(float64(Node[i].newLatFlow))
		R.Fwrite(x[:], 2, Fhotstart2.file)
		if Node[i].type_ == STORAGE {
			j = Node[i].subIndex
			x[0] = float32(float64(Storage[j].hrt))
			R.Fwrite(x[0:], 1, Fhotstart2.file)
		}
		for j = 0; j < Nobjects[POLLUT]; j++ {
			x[0] = float32(float64(Node[i].newQual[j]))
			R.Fwrite(x[0:], 1, Fhotstart2.file)
		}
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		x[0] = float32(float64(Link[i].newFlow))
		x[1] = float32(float64(Link[i].newDepth))
		x[2] = float32(float64(Link[i].setting))
		R.Fwrite(x[:], 3, Fhotstart2.file)
		for j = 0; j < Nobjects[POLLUT]; j++ {
			x[0] = float32(float64(Link[i].newQual[j]))
			R.Fwrite(x[0:], 1, Fhotstart2.file)
		}
	}
}

// hotstart.c:267
func readRouting() {
	var (
		i   int32
		j   int32
		x   float32
		xgw [4]float64
		f   *R.CFile
	)
	_, _, _, _, _ = i, j, x, xgw, f
	i = 0
	j = 0
	x = 0
	xgw = ([4]float64{})
	f = Fhotstart1.file
	if fileVersion == 2 {
		xgw[2] = float64(0.0)
		xgw[3] = (-(float64(10000000000.0)))
		for i = 0; i < Nobjects[SUBCATCH]; i++ {
			if !(readFloat(&x, f) != 0) {
				return
			}
			xgw[0] = float64(x)
			if !(readFloat(&x, f) != 0) {
				return
			}
			xgw[1] = float64(x)
			if Subcatch[i].groundwater != nil {
				gwater_setState(i, xgw[:])
			}
		}
	}
	for i = 0; i < Nobjects[NODE]; i++ {
		if !(readFloat(&x, f) != 0) {
			return
		}
		Node[i].newDepth = float64(x)
		if !(readFloat(&x, f) != 0) {
			return
		}
		Node[i].newLatFlow = float64(x)
		if (fileVersion >= 4) && (Node[i].type_ == STORAGE) {
			if !(readFloat(&x, f) != 0) {
				return
			}
			j = Node[i].subIndex
			Storage[j].hrt = float64(x)
		}
		for j = 0; j < Nobjects[POLLUT]; j++ {
			if !(readFloat(&x, f) != 0) {
				return
			}
			Node[i].newQual[j] = float64(x)
		}
		if fileVersion <= 2 {
			for j = 0; j < Nobjects[POLLUT]; j++ {
				if !(readFloat(&x, f) != 0) {
					return
				}
			}
		}
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		if !(readFloat(&x, f) != 0) {
			return
		}
		Link[i].newFlow = float64(x)
		if !(readFloat(&x, f) != 0) {
			return
		}
		Link[i].newDepth = float64(x)
		if !(readFloat(&x, f) != 0) {
			return
		}
		Link[i].setting = float64(x)
		Link[i].targetSetting = float64(x)
		link_setTargetSetting(i)
		link_setSetting(i, float64(0.0))
		for j = 0; j < Nobjects[POLLUT]; j++ {
			if !(readFloat(&x, f) != 0) {
				return
			}
			Link[i].newQual[j] = float64(x)
		}
	}
}

// hotstart.c:356
func saveRunoff() {
	var (
		i int32
		j int32
		k int32
		x [6]float64
		f *R.CFile
	)
	_, _, _, _, _ = i, j, k, x, f
	i = 0
	j = 0
	k = 0
	x = ([6]float64{})
	f = Fhotstart2.file
	for i = 0; i < Nobjects[SUBCATCH]; i++ {
		for j = 0; j < 3; j++ {
			x[j] = Subcatch[i].subArea[j].depth
		}
		x[3] = Subcatch[i].newRunoff
		R.Fwrite(x[:], 4, f)
		for j = 0; j < 6; j++ {
			x[j] = float64(0.0)
		}
		infil_getState(i, x[:])
		R.Fwrite(x[:], 6, f)
		if Subcatch[i].groundwater != nil {
			gwater_getState(i, x[:])
			R.Fwrite(x[:], 4, f)
		}
		if Subcatch[i].snowpack != nil {
			for j = 0; j < 3; j++ {
				snow_getState(i, j, x[:])
				R.Fwrite(x[:], 5, f)
			}
		}
		if Nobjects[POLLUT] > 0 {
			for j = 0; j < Nobjects[POLLUT]; j++ {
				x[0] = Subcatch[i].newQual[j]
				R.Fwrite(x[:], 1, f)
			}
			for j = 0; j < Nobjects[POLLUT]; j++ {
				x[0] = Subcatch[i].pondedQual[j]
				R.Fwrite(x[:], 1, f)
			}
			for k = 0; k < Nobjects[LANDUSE]; k++ {
				for j = 0; j < Nobjects[POLLUT]; j++ {
					x[0] = Subcatch[i].landFactor[k].buildup[j]
					R.Fwrite(x[:], uint64(Nobjects[POLLUT]), f)
				}
				x[0] = Subcatch[i].landFactor[k].lastSwept
				R.Fwrite(x[:], 1, f)
			}
		}
	}
}

// hotstart.c:430
func readRunoff() {
	var (
		i int32
		j int32
		k int32
		x [6]float64
		f *R.CFile
	)
	_, _, _, _, _ = i, j, k, x, f
	i = 0
	j = 0
	k = 0
	x = ([6]float64{})
	f = Fhotstart1.file
	for i = 0; i < Nobjects[SUBCATCH]; i++ {
		for j = 0; j < 3; j++ {
			if !(readDouble(&Subcatch[i].subArea[j].depth, f) != 0) {
				return
			}
		}
		if !(readDouble(&Subcatch[i].newRunoff, f) != 0) {
			return
		}
		for j = 0; j < 6; j++ {
			if !(readDouble(&x[j], f) != 0) {
				return
			}
		}
		infil_setState(i, x[:])
		if Subcatch[i].groundwater != nil {
			for j = 0; j < 4; j++ {
				if !(readDouble(&x[j], f) != 0) {
					return
				}
			}
			gwater_setState(i, x[:])
		}
		if Subcatch[i].snowpack != nil {
			for j = 0; j < 3; j++ {
				for k = 0; k < 5; k++ {
					if !(readDouble(&x[k], f) != 0) {
						return
					}
				}
				snow_setState(i, j, x[:])
			}
		}
		if Nobjects[POLLUT] > 0 {
			for j = 0; j < Nobjects[POLLUT]; j++ {
				if !(readDouble(&Subcatch[i].newQual[j], f) != 0) {
					return
				}
			}
			for j = 0; j < Nobjects[POLLUT]; j++ {
				if !(readDouble(&Subcatch[i].pondedQual[j], f) != 0) {
					return
				}
			}
			for k = 0; k < Nobjects[LANDUSE]; k++ {
				for j = 0; j < Nobjects[POLLUT]; j++ {
					if !(readDouble(&Subcatch[i].landFactor[k].buildup[j], f) != 0) {
						return
					}
				}
				if !(readDouble(&Subcatch[i].landFactor[k].lastSwept, f) != 0) {
					return
				}
			}
		}
	}
}

// hotstart.c:499
func readFloat(x *float32, f *R.CFile) int32 {
	R.Fread1(x, 4, f)
	if (*(x)) != (*(x)) {
		report_writeErrorMsg(ERR_HOTSTART_FILE_READ, L0)
		*(x) = float32(float64(float64(0.0)))
		return 0
	}
	return 1
}

// hotstart.c:521
func readDouble(x *float64, f *R.CFile) int32 {
	if R.Feof(f) != 0 {
		*(x) = float64(0.0)
		report_writeErrorMsg(ERR_HOTSTART_FILE_READ, L0)
		return 0
	}
	R.Fread1(x, 8, f)
	if (*(x)) != (*(x)) {
		*(x) = float64(0.0)
		return 0
	}
	return 1
}
