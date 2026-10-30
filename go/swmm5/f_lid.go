// Generated from EPA SWMM 5.2.4 (lid.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var LidLayerWords = [8]R.CPtr{L991, L992, L742, L993, L994, L995, L996, (R.CPtr{})}
var LidTypeWords = [9]R.CPtr{L997, L998, L953, L999, L1000, L1001, L1002, L1003, (R.CPtr{})}
var LidProcs []TLidProc
var LidCount int32
var LidGroups []*LidGroup
var GroupCount int32
var lid_EvapRate float64
var NativeInfil float64
var lid_MaxNativeInfil float64
var initLidRptFile_colCount int32 = 14
var initLidRptFile_head1 = [15]R.CPtr{L1022, L1023, L1024, L1024, L1025, L1026, L1027, L1028, L1025, L1029, L1025, L1026, L1027, L1030}
var initLidRptFile_head2 = [15]R.CPtr{L1022, L1031, L1032, L1033, L1034, L1035, L1035, L1036, L1037, L1038, L1039, L1039, L1040, L1041}
var initLidRptFile_units1 = [15]R.CPtr{L1042, L1043, L1044, L1044, L1044, L1044, L1044, L1044, L1044, L1044, L1045, L1045, L1046, L1047}
var initLidRptFile_units2 = [15]R.CPtr{L1042, L1043, L1048, L1048, L1048, L1048, L1048, L1048, L1048, L1048, L1049, L1049, L1046, L1050}
var initLidRptFile_line9 = R.BufFrom(L1051, 11)

// lid.c:237
func lid_create(lidCount int32, subcatchCount int32) {
	var (
		j int32
	)
	_ = j
	j = 0
	LidProcs = nil
	LidGroups = nil
	LidCount = lidCount
	GroupCount = subcatchCount
	if GroupCount > 0 {
		LidGroups = make([]*LidGroup, int(uint64(GroupCount)))
		if LidGroups == nil {
			ErrorCode = ERR_MEMORY
			return
		}
	}
	for j = 0; j < GroupCount; j++ {
		LidGroups[j] = nil
	}
	if LidCount == 0 {
		return
	}
	LidProcs = make([]TLidProc, int(uint64(LidCount)))
	if LidProcs == nil {
		ErrorCode = ERR_MEMORY
		return
	}
	for j = 0; j < LidCount; j++ {
		LidProcs[j].lidType = (-1)
		LidProcs[j].surface.thickness = float64(0.0)
		LidProcs[j].surface.voidFrac = float64(1.0)
		LidProcs[j].surface.roughness = float64(0.0)
		LidProcs[j].surface.surfSlope = float64(0.0)
		LidProcs[j].pavement.thickness = float64(0.0)
		LidProcs[j].soil.thickness = float64(0.0)
		LidProcs[j].storage.thickness = float64(0.0)
		LidProcs[j].storage.kSat = float64(0.0)
		LidProcs[j].drain.coeff = float64(0.0)
		LidProcs[j].drain.offset = float64(0.0)
		LidProcs[j].drainMat.thickness = float64(0.0)
		LidProcs[j].drainMat.roughness = float64(0.0)
		LidProcs[j].drainRmvl = nil
		LidProcs[j].drainRmvl = make([]float64, int(uint64(Nobjects[POLLUT])))
		if LidProcs[j].drainRmvl == nil {
			ErrorCode = ERR_MEMORY
			return
		}
	}
}

// lid.c:304
func lid_delete() {
	var (
		j int32
	)
	_ = j
	j = 0
	for j = 0; j < GroupCount; j++ {
		freeLidGroup(j)
	}
	{
		if LidGroups != nil {
			_ = LidGroups
			LidGroups = nil
		}
	}
	for j = 0; j < LidCount; j++ {
		if LidProcs[j].drainRmvl != nil {
			_ = LidProcs[j].drainRmvl
			LidProcs[j].drainRmvl = nil
		}
	}
	{
		if LidProcs != nil {
			_ = LidProcs
			LidProcs = nil
		}
	}
	GroupCount = 0
	LidCount = 0
}

// lid.c:322
func freeLidGroup(j int32) {
	var (
		lidGroup    *LidGroup
		lidList     *TLidList
		lidUnit     *TLidUnit
		nextLidUnit *TLidList
	)
	_, _, _, _ = lidGroup, lidList, lidUnit, nextLidUnit
	lidGroup = LidGroups[j]
	lidList = nil
	lidUnit = nil
	nextLidUnit = nil
	if lidGroup == nil {
		return
	}
	lidList = lidGroup.lidList
	for lidList != nil {
		lidUnit = lidList.lidUnit
		if lidUnit.rptFile != nil {
			if lidUnit.rptFile.file != nil {
				R.Fclose(lidUnit.rptFile.file)
			}
			_ = lidUnit.rptFile
		}
		nextLidUnit = lidList.nextLidUnit
		_ = lidUnit
		_ = lidList
		lidList = nextLidUnit
	}
	_ = lidGroup
	LidGroups[j] = nil
}

// lid.c:355
func lid_readProcParams(toks []R.CPtr, ntoks int32) int32 {
	var (
		j int32
		m int32
	)
	_, _ = j, m
	j = 0
	m = 0
	if ntoks < 2 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(LID, toks[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, toks[0])
	}
	if LidProcs[j].ID.IsNull() {
		LidProcs[j].ID = project_findID(LID, toks[0])
	}
	m = findmatch(toks[1], LidTypeWords[:])
	if m >= 0 {
		LidProcs[j].lidType = m
		return 0
	} else {
		m = findmatch(toks[1], LidLayerWords[:])
	}
	switch m {
	case SURF:
		return readSurfaceData(j, toks, ntoks)
	case SOIL:
		return readSoilData(j, toks, ntoks)
	case STOR:
		return readStorageData(j, toks, ntoks)
	case PAVE:
		return readPavementData(j, toks, ntoks)
	case 5:
		return readDrainData(j, toks, ntoks)
	case DRAINMAT:
		return readDrainMatData(j, toks, ntoks)
	case REMOVALS:
		return readRemovalsData(j, toks, ntoks)
	}
	return error_setInpError(ERR_KEYWORD, toks[1])
}

// lid.c:415
func lid_readGroupParams(toks []R.CPtr, ntoks int32) int32 {
	var (
		i             int32
		j             int32
		k             int32
		n_2           int32
		x             [6]float64
		fname         R.CPtr
		drainSubcatch int32
		drainNode     int32
	)
	_, _, _, _, _, _, _, _ = i, j, k, n_2, x, fname, drainSubcatch, drainNode
	i = 0
	j = 0
	k = 0
	n_2 = 0
	x = ([6]float64{})
	fname = (R.CPtr{})
	drainSubcatch = (-1)
	drainNode = (-1)
	if ntoks < 8 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(SUBCATCH, toks[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, toks[0])
	}
	k = project_findObject(LID, toks[1])
	if k < 0 {
		return error_setInpError(ERR_NAME, toks[1])
	}
	n_2 = R.Atoi(toks[2])
	if n_2 < 0 {
		return error_setInpError(ERR_NUMBER, toks[2])
	}
	if n_2 == 0 {
		return 0
	}
	for i = 3; i <= 7; i++ {
		if (!(getDouble(toks[i], &x[(i-3)]) != 0)) || (x[(i-3)] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, toks[i])
		}
	}
	for i = 2; i <= 3; i++ {
		if x[i] > float64(100.0) {
			return error_setInpError(ERR_NUMBER, toks[(i+3)])
		}
	}
	if (ntoks >= 9) && (R.Strcmp(toks[8], L275) != 0) {
		fname = toks[8]
	}
	if (ntoks >= 10) && (R.Strcmp(toks[9], L275) != 0) {
		drainSubcatch = project_findObject(SUBCATCH, toks[9])
		if drainSubcatch < 0 {
			drainNode = project_findObject(NODE, toks[9])
			if drainNode < 0 {
				return error_setInpError(ERR_NAME, toks[9])
			}
		}
	}
	x[5] = float64(0.0)
	if ntoks >= 11 {
		if ((!(getDouble(toks[10], &x[5]) != 0)) || (x[5] < float64(0.0))) || (x[5] > float64(100.0)) {
			return error_setInpError(ERR_NUMBER, toks[10])
		}
	}
	return addLidUnit(j, k, n_2, x[:], fname, drainSubcatch, drainNode)
}

// lid.c:499
func addLidUnit(j int32, k int32, n_2 int32, x []float64, fname R.CPtr, drainSubcatch int32, drainNode int32) int32 {
	var (
		lidUnit  *TLidUnit
		lidList  *TLidList
		lidGroup *LidGroup
	)
	_, _, _ = lidUnit, lidList, lidGroup
	lidUnit = nil
	lidList = nil
	lidGroup = nil
	lidGroup = LidGroups[j]
	if !(lidGroup != nil) {
		lidGroup = new(LidGroup)
		if !(lidGroup != nil) {
			return error_setInpError(ERR_MEMORY, L0)
		}
		lidGroup.lidList = nil
		LidGroups[j] = lidGroup
	}
	lidUnit = new(TLidUnit)
	if !(lidUnit != nil) {
		return error_setInpError(ERR_MEMORY, L0)
	}
	lidUnit.rptFile = nil
	lidList = new(TLidList)
	if !(lidList != nil) {
		_ = lidUnit
		return error_setInpError(ERR_MEMORY, L0)
	}
	lidList.lidUnit = lidUnit
	lidList.nextLidUnit = lidGroup.lidList
	lidGroup.lidList = lidList
	lidUnit.lidIndex = k
	lidUnit.number = n_2
	lidUnit.area = float64(x[0] / (float64((UCF(LENGTH)) * (UCF(LENGTH)))))
	lidUnit.fullWidth = float64(x[1] / UCF(LENGTH))
	lidUnit.initSat = float64(x[2] / float64(100.0))
	lidUnit.fromImperv = float64(x[3] / float64(100.0))
	lidUnit.toPerv = R.B2I((x[4] > float64(0.0)))
	lidUnit.fromPerv = float64(x[5] / float64(100.0))
	lidUnit.drainSubcatch = drainSubcatch
	lidUnit.drainNode = drainNode
	if !fname.IsNull() {
		if !(createLidRptFile(lidUnit, fname) != 0) {
			return error_setInpError(ERR_RPT_FILE, fname)
		}
	}
	return 0
}

// lid.c:567
func createLidRptFile(lidUnit *TLidUnit, fname R.CPtr) int32 {
	var (
		rptFile *TLidRptFile
	)
	_ = rptFile
	rptFile = nil
	rptFile = new(TLidRptFile)
	if rptFile == nil {
		return 0
	}
	lidUnit.rptFile = rptFile
	rptFile.file = R.Fopen(fname, L171)
	if rptFile.file == nil {
		return 0
	}
	return 1
}

// lid.c:581
func readSurfaceData(j int32, toks []R.CPtr, ntoks int32) int32 {
	var (
		i int32
		x [5]float64
	)
	_, _ = i, x
	i = 0
	x = ([5]float64{})
	if ntoks < 7 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	for i = 2; i < 7; i++ {
		if (!(getDouble(toks[i], &x[(i-2)]) != 0)) || (x[(i-2)] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, toks[i])
		}
	}
	if x[1] >= float64(1.0) {
		return error_setInpError(ERR_NUMBER, toks[3])
	}
	if x[0] == float64(0.0) {
		x[1] = float64(0.0)
	}
	LidProcs[j].surface.thickness = float64(x[0] / UCF(RAINDEPTH))
	LidProcs[j].surface.voidFrac = float64(float64(1.0) - x[1])
	LidProcs[j].surface.roughness = x[2]
	LidProcs[j].surface.surfSlope = float64(x[3] / float64(100.0))
	LidProcs[j].surface.sideSlope = x[4]
	return 0
}

// lid.c:616
func readPavementData(j int32, toks []R.CPtr, ntoks int32) int32 {
	var (
		i int32
		x [7]float64
	)
	_, _ = i, x
	i = 0
	x = ([7]float64{})
	if ntoks < 7 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	for i = 2; i < 7; i++ {
		if (!(getDouble(toks[i], &x[(i-2)]) != 0)) || (x[(i-2)] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, toks[i])
		}
	}
	x[5] = float64(0.0)
	if ntoks > 7 {
		if (!(getDouble(toks[7], &x[5]) != 0)) || (x[5] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, toks[7])
		}
	}
	x[6] = float64(0.0)
	if ntoks > 8 {
		if ((!(getDouble(toks[8], &x[6]) != 0)) || (x[6] < float64(0.0))) || (x[6] > float64(1.0)) {
			return error_setInpError(ERR_NUMBER, toks[8])
		}
	}
	x[1] = float64(x[1] / (float64(x[1] + float64(1.0))))
	LidProcs[j].pavement.thickness = float64(x[0] / UCF(RAINDEPTH))
	LidProcs[j].pavement.voidFrac = x[1]
	LidProcs[j].pavement.impervFrac = x[2]
	LidProcs[j].pavement.kSat = float64(x[3] / UCF(RAINFALL))
	LidProcs[j].pavement.clogFactor = x[4]
	LidProcs[j].pavement.regenDays = x[5]
	LidProcs[j].pavement.regenDegree = x[6]
	return 0
}

// lid.c:669
func readSoilData(j int32, toks []R.CPtr, ntoks int32) int32 {
	var (
		i int32
		x [7]float64
	)
	_, _ = i, x
	i = 0
	x = ([7]float64{})
	if ntoks < 9 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	for i = 2; i < 9; i++ {
		if (!(getDouble(toks[i], &x[(i-2)]) != 0)) || (x[(i-2)] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, toks[i])
		}
	}
	LidProcs[j].soil.thickness = float64(x[0] / UCF(RAINDEPTH))
	LidProcs[j].soil.porosity = x[1]
	LidProcs[j].soil.fieldCap = x[2]
	LidProcs[j].soil.wiltPoint = x[3]
	LidProcs[j].soil.kSat = float64(x[4] / UCF(RAINFALL))
	LidProcs[j].soil.kSlope = x[5]
	LidProcs[j].soil.suction = float64(x[6] / UCF(RAINDEPTH))
	return 0
}

// lid.c:703
func readStorageData(j int32, toks []R.CPtr, ntoks int32) int32 {
	var (
		i       int32
		covered int32
		x       [6]float64
	)
	_, _, _ = i, covered, x
	i = 0
	covered = 0
	x = ([6]float64{})
	if ntoks < 6 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	for i = 2; i < 6; i++ {
		if (!(getDouble(toks[i], &x[(i-2)]) != 0)) || (x[(i-2)] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, toks[i])
		}
	}
	if ntoks > 6 {
		if match(toks[6], L5) != 0 {
			covered = 1
		}
	}
	x[1] = float64(x[1] / (float64(x[1] + float64(1.0))))
	LidProcs[j].storage.thickness = float64(x[0] / UCF(RAINDEPTH))
	LidProcs[j].storage.voidFrac = x[1]
	LidProcs[j].storage.kSat = float64(x[2] / UCF(RAINFALL))
	LidProcs[j].storage.clogFactor = x[3]
	LidProcs[j].storage.covered = covered
	return 0
}

// lid.c:749
func readDrainData(j int32, toks []R.CPtr, ntoks int32) int32 {
	var (
		i int32
		x [6]float64
	)
	_, _ = i, x
	i = 0
	x = ([6]float64{})
	if ntoks < 6 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	for i = 0; i < 6; i++ {
		x[i] = float64(0.0)
	}
	for i = 2; i < 8; i++ {
		if (ntoks > i) && ((!(getDouble(toks[i], &x[(i-2)]) != 0)) || (x[(i-2)] < float64(0.0))) {
			return error_setInpError(ERR_NUMBER, toks[i])
		}
	}
	i = (-1)
	if ntoks >= 9 {
		i = project_findObject(CURVE, toks[8])
		if i < 0 {
			return error_setInpError(ERR_NAME, toks[8])
		}
	}
	LidProcs[j].drain.coeff = x[0]
	LidProcs[j].drain.expon = x[1]
	LidProcs[j].drain.offset = float64(x[2] / UCF(RAINDEPTH))
	LidProcs[j].drain.delay = float64(x[3] * float64(3600.0))
	LidProcs[j].drain.hOpen = float64(x[4] / UCF(RAINDEPTH))
	LidProcs[j].drain.hClose = float64(x[5] / UCF(RAINDEPTH))
	LidProcs[j].drain.qCurve = i
	return 0
}

// lid.c:794
func readDrainMatData(j int32, toks []R.CPtr, ntoks int32) int32 {
	var (
		i int32
		x [3]float64
	)
	_, _ = i, x
	i = 0
	x = ([3]float64{})
	if ntoks < 5 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	if LidProcs[j].lidType != GREEN_ROOF {
		return 0
	}
	for i = 2; i < 5; i++ {
		if (!(getDouble(toks[i], &x[(i-2)]) != 0)) || (x[(i-2)] < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, toks[i])
		}
	}
	LidProcs[j].drainMat.thickness = float64(x[0] / UCF(RAINDEPTH))
	LidProcs[j].drainMat.voidFrac = x[1]
	LidProcs[j].drainMat.roughness = x[2]
	return 0
}

// lid.c:828
func readRemovalsData(j int32, toks []R.CPtr, ntoks int32) int32 {
	var (
		i    int32
		p    int32
		rmvl float64
	)
	_, _, _ = i, p, rmvl
	i = 2
	p = 0
	rmvl = 0
	if ntoks < 4 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	for ntoks > i {
		p = project_findObject(POLLUT, toks[i])
		if p < 0 {
			return error_setInpError(ERR_NAME, toks[i])
		}
		i++
		if ntoks == i {
			return error_setInpError(ERR_ITEMS, L0)
		}
		if ((!(getDouble(toks[i], &rmvl) != 0)) || (rmvl < float64(0.0))) || (rmvl > float64(100.0)) {
			return error_setInpError(ERR_NUMBER, toks[i])
		}
		LidProcs[j].drainRmvl[p] = float64(rmvl / float64(100.0))
		i++
	}
	return 0
}

// lid.c:869
func lid_writeSummary() {
	var (
		j        int32
		k        int32
		pctArea  float64
		lidUnit  *TLidUnit
		lidList  *TLidList
		lidGroup *LidGroup
	)
	_, _, _, _, _, _ = j, k, pctArea, lidUnit, lidList, lidGroup
	j = 0
	k = 0
	pctArea = 0
	lidUnit = nil
	lidList = nil
	lidGroup = nil
	R.Fprintf(Frpt.file, L276)
	R.Fprintf(Frpt.file, L276)
	R.Fprintf(Frpt.file, L277)
	R.Fprintf(Frpt.file, L278)
	R.Fprintf(Frpt.file, L277)
	R.Fprintf(Frpt.file, L279)
	R.Fprintf(Frpt.file, L280)
	R.Fprintf(Frpt.file, L281)
	for j = 0; j < GroupCount; j++ {
		lidGroup = LidGroups[j]
		if lidGroup == nil {
			continue
		}
		lidList = lidGroup.lidList
		for lidList != nil {
			lidUnit = lidList.lidUnit
			k = lidUnit.lidIndex
			pctArea = float64(float64(float64(lidUnit.area*float64(lidUnit.number))/Subcatch[j].area) * float64(100.0))
			R.Fprintf(Frpt.file, L282, Subcatch[j].ID, LidProcs[k].ID)
			R.Fprintf(Frpt.file, L283, lidUnit.number, float64(lidUnit.area*(float64((UCF(LENGTH))*(UCF(LENGTH))))), float64(lidUnit.fullWidth*UCF(LENGTH)), pctArea, float64(lidUnit.fromImperv*float64(100.0)), float64(lidUnit.fromPerv*float64(100.0)))
			lidList = lidList.nextLidUnit
		}
	}
}

// lid.c:918
func lid_validate() {
	var (
		j int32
	)
	_ = j
	j = 0
	for j = 0; j < LidCount; j++ {
		validateLidProc(j)
	}
	for j = 0; j < GroupCount; j++ {
		validateLidGroup(j)
	}
}

// lid.c:932
func validateLidProc(j int32) {
	var (
		layerMissing int32
	)
	_ = layerMissing
	layerMissing = 0
	if LidProcs[j].lidType < 0 {
		report_writeErrorMsg(ERR_LID_TYPE, LidProcs[j].ID)
		return
	}
	switch LidProcs[j].lidType {
	case BIO_CELL, RAIN_GARDEN:
		if LidProcs[j].soil.thickness <= float64(0.0) {
			layerMissing = 1
		}
		break
	case GREEN_ROOF:
		if LidProcs[j].soil.thickness <= float64(0.0) {
			layerMissing = 1
		}
		if LidProcs[j].drainMat.thickness <= float64(0.0) {
			layerMissing = 1
		}
		break
	case POROUS_PAVEMENT:
		if LidProcs[j].pavement.thickness <= float64(0.0) {
			layerMissing = 1
		}
		break
	case INFIL_TRENCH:
		if LidProcs[j].storage.thickness <= float64(0.0) {
			layerMissing = 1
		}
		break
	}
	if layerMissing != 0 {
		report_writeErrorMsg(ERR_LID_LAYER, LidProcs[j].ID)
		return
	}
	if LidProcs[j].lidType == POROUS_PAVEMENT {
		if ((((LidProcs[j].pavement.thickness <= float64(0.0)) || (LidProcs[j].pavement.kSat <= float64(0.0))) || (LidProcs[j].pavement.voidFrac <= float64(0.0))) || (LidProcs[j].pavement.voidFrac > float64(1.0))) || (LidProcs[j].pavement.impervFrac > float64(1.0)) {
			sstrncpy(Msg, LidProcs[j].ID, 1024)
			sstrcat(Msg, L284, 1024)
			report_writeErrorMsg(ERR_LID_PARAMS, Msg)
		}
	}
	if LidProcs[j].soil.thickness > float64(0.0) {
		if ((((LidProcs[j].soil.porosity <= float64(0.0)) || (LidProcs[j].soil.fieldCap >= LidProcs[j].soil.porosity)) || (LidProcs[j].soil.wiltPoint >= LidProcs[j].soil.fieldCap)) || (LidProcs[j].soil.kSat <= float64(0.0))) || (LidProcs[j].soil.kSlope < float64(0.0)) {
			sstrncpy(Msg, LidProcs[j].ID, 1024)
			sstrcat(Msg, L285, 1024)
			report_writeErrorMsg(ERR_LID_PARAMS, Msg)
		}
	}
	if LidProcs[j].storage.thickness > float64(0.0) {
		if (LidProcs[j].storage.voidFrac <= float64(0.0)) || (LidProcs[j].storage.voidFrac > float64(1.0)) {
			sstrncpy(Msg, LidProcs[j].ID, 1024)
			sstrcat(Msg, L286, 1024)
			report_writeErrorMsg(ERR_LID_PARAMS, Msg)
		}
	} else {
		LidProcs[j].storage.voidFrac = float64(1.0)
		LidProcs[j].drain.offset = float64(0.0)
	}
	if (LidProcs[j].drain.hOpen > float64(0.0)) && (LidProcs[j].drain.hOpen <= LidProcs[j].drain.hClose) {
		sstrncpy(Msg, LidProcs[j].ID, 1024)
		sstrcat(Msg, L287, 1024)
		report_writeErrorMsg(ERR_LID_PARAMS, Msg)
	}
	if LidProcs[j].lidType == VEG_SWALE {
		if (float64(LidProcs[j].surface.roughness*LidProcs[j].surface.surfSlope) <= float64(0.0)) || (LidProcs[j].surface.thickness == float64(0.0)) {
			sstrncpy(Msg, LidProcs[j].ID, 1024)
			sstrcat(Msg, L288, 1024)
			report_writeErrorMsg(ERR_LID_PARAMS, Msg)
		} else {
			LidProcs[j].surface.alpha = float64(float64(float64(1.49)*M.Sqrt(LidProcs[j].surface.surfSlope)) / LidProcs[j].surface.roughness)
		}
	} else {
		if LidProcs[j].surface.roughness > float64(0.0) {
			LidProcs[j].surface.alpha = float64(float64(float64(1.49)/LidProcs[j].surface.roughness) * M.Sqrt(LidProcs[j].surface.surfSlope))
		} else {
			LidProcs[j].surface.alpha = float64(0.0)
		}
	}
	if LidProcs[j].drainMat.roughness > float64(0.0) {
		LidProcs[j].drainMat.alpha = float64(float64(float64(1.49)/LidProcs[j].drainMat.roughness) * M.Sqrt(LidProcs[j].surface.surfSlope))
	} else {
		LidProcs[j].drainMat.alpha = float64(0.0)
	}
	if LidProcs[j].pavement.thickness > float64(0.0) {
		LidProcs[j].pavement.clogFactor = float64(LidProcs[j].pavement.clogFactor * (float64(float64(LidProcs[j].pavement.thickness*LidProcs[j].pavement.voidFrac) * (float64(float64(1.0) - LidProcs[j].pavement.impervFrac)))))
	}
	if LidProcs[j].storage.thickness > float64(0.0) {
		LidProcs[j].storage.clogFactor = float64(LidProcs[j].storage.clogFactor * (float64(LidProcs[j].storage.thickness * LidProcs[j].storage.voidFrac)))
	} else {
		LidProcs[j].storage.clogFactor = float64(0.0)
	}
	LidProcs[j].surface.canOverflow = 1
	switch LidProcs[j].lidType {
	case ROOF_DISCON:
		LidProcs[j].surface.canOverflow = 0
		break
	case INFIL_TRENCH, POROUS_PAVEMENT, BIO_CELL, RAIN_GARDEN, GREEN_ROOF:
		if LidProcs[j].surface.alpha > float64(0.0) {
			LidProcs[j].surface.canOverflow = 0
		}
	}
	if LidProcs[j].lidType == RAIN_BARREL {
		LidProcs[j].storage.voidFrac = float64(1.0)
		LidProcs[j].storage.kSat = float64(0.0)
	}
	if LidProcs[j].lidType == GREEN_ROOF {
		LidProcs[j].storage.thickness = LidProcs[j].drainMat.thickness
		LidProcs[j].storage.voidFrac = LidProcs[j].drainMat.voidFrac
		LidProcs[j].storage.clogFactor = float64(0.0)
		LidProcs[j].storage.kSat = float64(0.0)
	}
}

// lid.c:1113
func validateLidGroup(j int32) {
	var (
		k            int32
		p            [3]float64
		totalArea    float64
		totalLidArea float64
		fromImperv   float64
		fromPerv     float64
		lidUnit      *TLidUnit
		lidList      *TLidList
		lidGroup     *LidGroup
	)
	_, _, _, _, _, _, _, _, _ = k, p, totalArea, totalLidArea, fromImperv, fromPerv, lidUnit, lidList, lidGroup
	k = 0
	p = ([3]float64{})
	totalArea = Subcatch[j].area
	totalLidArea = float64(0.0)
	fromImperv = float64(0.0)
	fromPerv = float64(0.0)
	lidUnit = nil
	lidList = nil
	lidGroup = nil
	lidGroup = LidGroups[j]
	if lidGroup == nil {
		return
	}
	lidList = lidGroup.lidList
	for lidList != nil {
		lidUnit = lidList.lidUnit
		k = lidUnit.lidIndex
		totalLidArea = float64(totalLidArea + (float64(lidUnit.area * float64(lidUnit.number))))
		fromImperv = float64(fromImperv + lidUnit.fromImperv)
		fromPerv = float64(fromPerv + lidUnit.fromPerv)
		lidUnit.soilInfil.Ks = float64(0.0)
		if LidProcs[k].soil.thickness > float64(0.0) {
			p[0] = float64(LidProcs[k].soil.suction * UCF(RAINDEPTH))
			p[1] = float64(LidProcs[k].soil.kSat * UCF(RAINFALL))
			p[2] = float64((float64(LidProcs[k].soil.porosity - LidProcs[k].soil.wiltPoint)) * (float64(float64(1.0) - lidUnit.initSat)))
			if grnampt_setParams(&lidUnit.soilInfil, p[:]) == 0 {
				sstrncpy(Msg, LidProcs[k].ID, 1024)
				sstrcat(Msg, L285, 1024)
				report_writeErrorMsg(ERR_LID_PARAMS, Msg)
			}
		}
		if LidProcs[k].lidType == VEG_SWALE {
			if (Subcatch[j].infilModel == GREEN_AMPT) || (Subcatch[j].infilModel == MOD_GREEN_AMPT) {
				grnampt_getParams(j, p[:])
				if grnampt_setParams(&lidUnit.soilInfil, p[:]) == 0 {
					sstrncpy(Msg, LidProcs[k].ID, 1024)
					sstrcat(Msg, L289, 1024)
					report_writeErrorMsg(ERR_LID_PARAMS, Msg)
				}
			}
			if lidUnit.fullWidth <= float64(0.0) {
				sstrncpy(Msg, LidProcs[k].ID, 1024)
				sstrcat(Msg, L290, 1024)
				report_writeErrorMsg(ERR_LID_PARAMS, Msg)
			}
		}
		if Subcatch[j].fracImperv >= float64(0.999) {
			lidUnit.toPerv = 0
		}
		if (lidUnit.drainNode == (-1)) && (lidUnit.drainSubcatch == (-1)) {
			lidUnit.drainNode = Subcatch[j].outNode
			lidUnit.drainSubcatch = Subcatch[j].outSubcatch
		}
		lidList = lidList.nextLidUnit
	}
	if totalLidArea > float64(float64(1.001)*totalArea) {
		report_writeErrorMsg(ERR_LID_AREAS, Subcatch[j].ID)
	}
	if (fromImperv > float64(1.001)) || (fromPerv > float64(1.001)) {
		report_writeErrorMsg(ERR_LID_CAPTURE_AREA, Subcatch[j].ID)
	}
	if totalLidArea > float64(float64(0.999)*totalArea) {
		totalLidArea = totalArea
	}
	Subcatch[j].lidArea = totalLidArea
}

// lid.c:1211
func lid_initState() {
	var (
		i           int32
		j           int32
		k           int32
		lidUnit     *TLidUnit
		lidList     *TLidList
		lidGroup    *LidGroup
		initVol     float64
		initDryTime float64
	)
	_, _, _, _, _, _, _, _ = i, j, k, lidUnit, lidList, lidGroup, initVol, initDryTime
	i = 0
	j = 0
	k = 0
	lidUnit = nil
	lidList = nil
	lidGroup = nil
	initVol = 0
	initDryTime = float64(StartDryDays * float64(86400))
	HasWetLids = 0
	for j = 0; j < GroupCount; j++ {
		lidGroup = LidGroups[j]
		if lidGroup == nil {
			continue
		}
		lidGroup.pervArea = float64(0.0)
		lidGroup.flowToPerv = float64(0.0)
		lidGroup.oldDrainFlow = float64(0.0)
		lidGroup.newDrainFlow = float64(0.0)
		lidList = lidGroup.lidList
		for lidList != nil {
			lidUnit = lidList.lidUnit
			k = lidUnit.lidIndex
			lidUnit.surfaceDepth = float64(0.0)
			lidUnit.storageDepth = float64(0.0)
			lidUnit.soilMoisture = float64(0.0)
			lidUnit.paveDepth = float64(0.0)
			lidUnit.dryTime = initDryTime
			lidUnit.volTreated = float64(0.0)
			lidUnit.nextRegenDay = LidProcs[k].pavement.regenDays
			initVol = float64(0.0)
			if LidProcs[k].soil.thickness > float64(0.0) {
				lidUnit.soilMoisture = float64(LidProcs[k].soil.wiltPoint + float64(lidUnit.initSat*(float64(LidProcs[k].soil.porosity-LidProcs[k].soil.wiltPoint))))
				initVol = float64(initVol + (float64(lidUnit.soilMoisture * LidProcs[k].soil.thickness)))
			}
			if LidProcs[k].storage.thickness > float64(0.0) {
				lidUnit.storageDepth = float64(lidUnit.initSat * LidProcs[k].storage.thickness)
				initVol = float64(initVol + (float64(lidUnit.storageDepth * LidProcs[k].storage.voidFrac)))
			}
			if lidUnit.initSat > float64(0.0) {
				HasWetLids = 1
			}
			lidproc_initWaterBalance(lidUnit, initVol)
			lidUnit.volTreated = float64(0.0)
			if lidUnit.rptFile != nil {
				initLidRptFile(Title[0], LidProcs[k].ID, Subcatch[j].ID, lidUnit)
			}
			lidUnit.oldDrainFlow = float64(0.0)
			lidUnit.newDrainFlow = float64(0.0)
			for i = 0; i < 4; i++ {
				lidUnit.oldFluxRates[i] = float64(0.0)
			}
			if lidUnit.soilInfil.Ks > float64(0.0) {
				grnampt_initState(&lidUnit.soilInfil)
			}
			if isLidPervious(lidUnit.lidIndex) != 0 {
				lidGroup.pervArea = float64(lidGroup.pervArea + (float64(lidUnit.area * float64(lidUnit.number))))
			}
			lidList = lidList.nextLidUnit
		}
	}
}

// lid.c:1302
func lid_setOldGroupState(j int32) {
	var (
		lidList *TLidList
	)
	_ = lidList
	lidList = nil
	if LidGroups[j] != nil {
		LidGroups[j].oldDrainFlow = LidGroups[j].newDrainFlow
		LidGroups[j].newDrainFlow = float64(0.0)
		lidList = LidGroups[j].lidList
		for lidList != nil {
			lidList.lidUnit.oldDrainFlow = lidList.lidUnit.newDrainFlow
			lidList.lidUnit.newDrainFlow = float64(0.0)
			lidList = lidList.nextLidUnit
		}
	}
}

// lid.c:1326
func isLidPervious(k int32) int32 {
	return R.B2I(((LidProcs[k].storage.thickness == float64(0.0)) || (LidProcs[k].storage.kSat > float64(0.0))))
}

// lid.c:1339
func getSurfaceDepth(j int32) float64 {
	var (
		k        int32
		depth    float64
		lidUnit  *TLidUnit
		lidList  *TLidList
		lidGroup *LidGroup
	)
	_, _, _, _, _ = k, depth, lidUnit, lidList, lidGroup
	k = 0
	depth = float64(0.0)
	lidUnit = nil
	lidList = nil
	lidGroup = nil
	lidGroup = LidGroups[j]
	if lidGroup == nil {
		return float64(0.0)
	}
	if Subcatch[j].lidArea == float64(0.0) {
		return float64(0.0)
	}
	lidList = lidGroup.lidList
	for lidList != nil {
		lidUnit = lidList.lidUnit
		k = lidUnit.lidIndex
		depth = float64(depth + (float64(float64(float64(lidUnit.surfaceDepth*LidProcs[k].surface.voidFrac)*lidUnit.area) * float64(lidUnit.number))))
		lidList = lidList.nextLidUnit
	}
	return float64(depth / Subcatch[j].lidArea)
}

// lid.c:1370
func lid_getPervArea(j int32) float64 {
	if LidGroups[j] != nil {
		return LidGroups[j].pervArea
	} else {
		return float64(0.0)
	}
	return 0
}

// lid.c:1383
func lid_getFlowToPerv(j int32) float64 {
	if LidGroups[j] != nil {
		return LidGroups[j].flowToPerv
	}
	return float64(0.0)
}

// lid.c:1397
func lid_getStoredVolume(j int32) float64 {
	var (
		total    float64
		lidUnit  *TLidUnit
		lidList  *TLidList
		lidGroup *LidGroup
	)
	_, _, _, _ = total, lidUnit, lidList, lidGroup
	total = float64(0.0)
	lidUnit = nil
	lidList = nil
	lidGroup = nil
	lidGroup = LidGroups[j]
	if (lidGroup == nil) || (Subcatch[j].lidArea == float64(0.0)) {
		return float64(0.0)
	}
	lidList = lidGroup.lidList
	for lidList != nil {
		lidUnit = lidList.lidUnit
		total = float64(total + (float64(float64(lidUnit.waterBalance.finalVol*lidUnit.area) * float64(lidUnit.number))))
		lidList = lidList.nextLidUnit
	}
	return total
}

// lid.c:1424
func lid_getDrainFlow(j int32, timePeriod int32) float64 {
	if LidGroups[j] != nil {
		if timePeriod == PREVIOUS {
			return LidGroups[j].oldDrainFlow
		} else {
			return LidGroups[j].newDrainFlow
		}
	}
	return float64(0.0)
}

// lid.c:1442
func lid_addDrainLoads(j int32, c []float64, tStep float64) {
	var (
		isRunoffLoad int32
		p            int32
		r            float64
		w            float64
		lidUnit      *TLidUnit
		lidList      *TLidList
		lidGroup     *LidGroup
	)
	_, _, _, _, _, _, _ = isRunoffLoad, p, r, w, lidUnit, lidList, lidGroup
	isRunoffLoad = 0
	p = 0
	r = 0
	w = 0
	lidUnit = nil
	lidList = nil
	lidGroup = nil
	lidGroup = LidGroups[j]
	if lidGroup != nil {
		lidList = lidGroup.lidList
		for lidList != nil {
			lidUnit = lidList.lidUnit
			isRunoffLoad = R.B2I(((lidUnit.drainNode >= 0) || (lidUnit.drainSubcatch == j)))
			if !(lidUnit.toPerv != 0) {
				for p = 0; p < Nobjects[POLLUT]; p++ {
					w = float64(float64(float64(float64(lidUnit.newDrainFlow*c[p])*tStep)*float64(28.317)) * Pollut[p].mcf)
					r = LidProcs[lidUnit.lidIndex].drainRmvl[p]
					massbal_updateLoadingTotals(BMP_REMOVAL_LOAD, p, float64(r*w))
					if isRunoffLoad != 0 {
						massbal_updateLoadingTotals(RUNOFF_LOAD, p, float64(w*(float64(float64(1.0)-r))))
					}
				}
			}
			lidList = lidList.nextLidUnit
		}
	}
}

// lid.c:1497
func lid_addDrainRunon(j int32) {
	var (
		i        int32
		k        int32
		p        int32
		q        float64
		w        float64
		lidUnit  *TLidUnit
		lidList  *TLidList
		lidGroup *LidGroup
	)
	_, _, _, _, _, _, _, _ = i, k, p, q, w, lidUnit, lidList, lidGroup
	i = 0
	k = 0
	p = 0
	q = 0
	w = 0
	lidUnit = nil
	lidList = nil
	lidGroup = nil
	lidGroup = LidGroups[j]
	if lidGroup != nil {
		lidList = lidGroup.lidList
		for lidList != nil {
			lidUnit = lidList.lidUnit
			i = lidUnit.lidIndex
			k = lidUnit.drainSubcatch
			if (k >= 0) && (k != j) {
				q = lidUnit.oldDrainFlow
				subcatch_addRunonFlow(k, q)
				for p = 0; p < Nobjects[POLLUT]; p++ {
					w = float64(float64(q*Subcatch[j].oldQual[p]) * float64(28.317))
					w = float64(w * (float64(float64(1.0) - LidProcs[i].drainRmvl[p])))
					Subcatch[k].newQual[p] = float64(Subcatch[k].newQual[p] + w)
				}
			}
			lidList = lidList.nextLidUnit
		}
	}
}

// lid.c:1549
func lid_addDrainInflow(j int32, f float64) {
	var (
		i        int32
		k        int32
		p        int32
		q        float64
		w        float64
		w1       float64
		w2       float64
		lidUnit  *TLidUnit
		lidList  *TLidList
		lidGroup *LidGroup
	)
	_, _, _, _, _, _, _, _, _, _ = i, k, p, q, w, w1, w2, lidUnit, lidList, lidGroup
	i = 0
	k = 0
	p = 0
	q = 0
	w = 0
	w1 = 0
	w2 = 0
	lidUnit = nil
	lidList = nil
	lidGroup = nil
	lidGroup = LidGroups[j]
	if lidGroup != nil {
		lidList = lidGroup.lidList
		for lidList != nil {
			lidUnit = lidList.lidUnit
			i = lidUnit.lidIndex
			k = lidUnit.drainNode
			if k >= 0 {
				q = float64(float64((float64(float64(1.0)-f))*lidUnit.oldDrainFlow) + float64(f*lidUnit.newDrainFlow))
				Node[k].newLatFlow = float64(Node[k].newLatFlow + q)
				massbal_addInflowFlow(WET_WEATHER_INFLOW, q)
				for p = 0; p < Nobjects[POLLUT]; p++ {
					w1 = float64(lidUnit.oldDrainFlow * Subcatch[j].oldQual[p])
					w2 = float64(lidUnit.newDrainFlow * Subcatch[j].newQual[p])
					w = float64(float64((float64(float64(1.0)-f))*w1) + float64(f*w2))
					w = float64(w * (float64(float64(1.0) - LidProcs[i].drainRmvl[p])))
					Node[k].newQual[p] = float64(Node[k].newQual[p] + w)
					massbal_addInflowQual(WET_WEATHER_INFLOW, p, w)
				}
			}
			lidList = lidList.nextLidUnit
		}
	}
}

// lid.c:1609
func lid_getRunoff(j int32, tStep float64) {
	var (
		theLidGroup *LidGroup
		lidList     *TLidList
		lidUnit     *TLidUnit
		lidArea     float64
		qImperv     float64
		qPerv       float64
		lidInflow   float64
		qRunoff     float64
		qDrain      float64
		qReturn     float64
	)
	_, _, _, _, _, _, _, _, _, _ = theLidGroup, lidList, lidUnit, lidArea, qImperv, qPerv, lidInflow, qRunoff, qDrain, qReturn
	theLidGroup = nil
	lidList = nil
	lidUnit = nil
	lidArea = 0
	qImperv = float64(0.0)
	qPerv = float64(0.0)
	lidInflow = float64(0.0)
	qRunoff = float64(0.0)
	qDrain = float64(0.0)
	qReturn = float64(0.0)
	theLidGroup = LidGroups[j]
	if !(theLidGroup != nil) {
		return
	}
	lidList = theLidGroup.lidList
	if !(lidList != nil) {
		return
	}
	lid_EvapRate = Evap.rate
	if (Evap.dryOnly != 0) && (Subcatch[j].rainfall > float64(0.0)) {
		lid_EvapRate = float64(0.0)
	}
	findNativeInfil(j, tStep)
	if Subcatch[j].area > Subcatch[j].lidArea {
		qImperv = getImpervAreaRunoff(j)
		qPerv = getPervAreaRunoff(j)
	}
	for lidList != nil {
		lidUnit = lidList.lidUnit
		lidArea = float64(lidUnit.area * float64(lidUnit.number))
		if lidArea > float64(0.0) {
			lidInflow = float64((float64(float64(qImperv*lidUnit.fromImperv) + float64(qPerv*lidUnit.fromPerv))) / lidArea)
			VlidIn = float64(VlidIn + (float64(float64(lidInflow*lidArea) * tStep)))
			lidInflow = float64(lidInflow + getRainInflow(j, lidUnit))
			if Subcatch[j].area == Subcatch[j].lidArea {
				lidInflow = float64(lidInflow + Subcatch[j].runon)
			}
			evalLidUnit(j, lidUnit, lidArea, lidInflow, tStep, &qRunoff, &qDrain, &qReturn)
		}
		lidList = lidList.nextLidUnit
	}
	theLidGroup.newDrainFlow = qDrain
	theLidGroup.flowToPerv = qReturn
	VlidOut = float64(qRunoff * tStep)
	VlidDrain = float64(qDrain * tStep)
	VlidReturn = float64(qReturn * tStep)
}

// lid.c:1697
func findNativeInfil(j int32, tStep float64) {
	var (
		nonLidArea float64
	)
	_ = nonLidArea
	nonLidArea = 0
	nonLidArea = float64(Subcatch[j].area - Subcatch[j].lidArea)
	if (nonLidArea > float64(0.0)) && (Subcatch[j].fracImperv < float64(1.0)) {
		NativeInfil = float64(float64(Vinfil/nonLidArea) / tStep)
	} else {
		NativeInfil = infil_getInfil(j, tStep, Subcatch[j].rainfall, Subcatch[j].runon, getSurfaceDepth(j))
	}
	if (!(IgnoreGwater != 0)) && (Subcatch[j].groundwater != nil) {
		lid_MaxNativeInfil = float64(Subcatch[j].groundwater.maxInfilVol / tStep)
	} else {
		lid_MaxNativeInfil = float64(10000000000.0)
	}
}

// lid.c:1734
func getRainInflow(j int32, lidUnit *TLidUnit) float64 {
	var (
		lidProc *TLidProc
	)
	_ = lidProc
	lidProc = &LidProcs[lidUnit.lidIndex]
	if (lidProc.lidType == RAIN_BARREL) && (lidProc.storage.covered == 1) {
		return float64(0.0)
	}
	return Subcatch[j].rainfall
}

// lid.c:1751
func getImpervAreaRunoff(j int32) float64 {
	var (
		i          int32
		q          float64
		nonLidArea float64
	)
	_, _, _ = i, q, nonLidArea
	i = 0
	q = float64(0.0)
	nonLidArea = 0
	for i = IMPERV0; i <= IMPERV1; i++ {
		q = float64(q + (float64(Subcatch[j].subArea[i].runoff * Subcatch[j].subArea[i].fArea)))
	}
	if (Subcatch[j].subArea[IMPERV0].routeTo == TO_PERV) && (Subcatch[j].fracImperv < float64(1.0)) {
		q = float64(q * Subcatch[j].subArea[IMPERV0].fOutlet)
	}
	nonLidArea = float64(Subcatch[j].area - Subcatch[j].lidArea)
	return float64(q * nonLidArea)
}

// lid.c:1781
func getPervAreaRunoff(j int32) float64 {
	var (
		q          float64
		nonLidArea float64
	)
	_, _ = q, nonLidArea
	q = float64(0.0)
	nonLidArea = 0
	q = float64(Subcatch[j].subArea[PERV].runoff * Subcatch[j].subArea[PERV].fArea)
	if (Subcatch[j].subArea[PERV].routeTo == TO_IMPERV) && (Subcatch[j].fracImperv > float64(0.0)) {
		q = float64(q * Subcatch[j].subArea[PERV].fOutlet)
	}
	nonLidArea = float64(Subcatch[j].area - Subcatch[j].lidArea)
	return float64(q * nonLidArea)
}

// lid.c:1807
func evalLidUnit(j int32, lidUnit *TLidUnit, lidArea float64, lidInflow float64, tStep float64, qRunoff *float64, qDrain *float64, qReturn *float64) {
	var (
		lidProc   *TLidProc
		lidRunoff float64
		lidEvap   float64
		lidInfil  float64
		lidDrain  float64
	)
	_, _, _, _, _ = lidProc, lidRunoff, lidEvap, lidInfil, lidDrain
	lidProc = nil
	lidRunoff = 0
	lidEvap = 0
	lidInfil = 0
	lidDrain = 0
	lidProc = &LidProcs[lidUnit.lidIndex]
	lidEvap = float64(0.0)
	lidInfil = float64(0.0)
	lidRunoff = float64(lidproc_getOutflow(lidUnit, lidProc, lidInflow, lid_EvapRate, NativeInfil, lid_MaxNativeInfil, tStep, &lidEvap, &lidInfil, &lidDrain) * lidArea)
	lidDrain = float64(lidDrain * lidArea)
	if (lidUnit.toPerv != 0) && (Subcatch[j].area > Subcatch[j].lidArea) {
		*qReturn = float64((*qReturn) + lidRunoff)
		lidRunoff = float64(0.0)
		if (lidUnit.drainNode == Subcatch[j].outNode) && (lidUnit.drainSubcatch == Subcatch[j].outSubcatch) {
			*qReturn = float64((*qReturn) + lidDrain)
			lidDrain = float64(0.0)
		}
	}
	if lidUnit.drainNode >= 0 {
		massbal_updateRunoffTotals(RUNOFF_DRAINS, float64(lidDrain*tStep))
	}
	lidUnit.newDrainFlow = lidDrain
	Vevap = float64(Vevap + (float64(float64(lidEvap*tStep) * lidArea)))
	VlidInfil = float64(VlidInfil + (float64(float64(lidInfil*tStep) * lidArea)))
	if isLidPervious(lidUnit.lidIndex) != 0 {
		Vpevap = float64(Vpevap + (float64(float64(lidEvap*tStep) * lidArea)))
	}
	if Subcatch[j].rainfall > float64(2.31481e-08) {
		lidUnit.dryTime = float64(0.0)
	} else {
		lidUnit.dryTime = float64(lidUnit.dryTime + tStep)
	}
	lidproc_saveResults(lidUnit, UCF(RAINFALL), UCF(RAINDEPTH))
	*qRunoff = float64((*qRunoff) + lidRunoff)
	*qDrain = float64((*qDrain) + lidDrain)
}

// lid.c:1890
func lid_writeWaterBalance() {
	var (
		j        int32
		k        int32
		ucf      float64
		inflow   float64
		outflow  float64
		err      float64
		lidUnit  *TLidUnit
		lidList  *TLidList
		lidGroup *LidGroup
	)
	_, _, _, _, _, _, _, _, _ = j, k, ucf, inflow, outflow, err, lidUnit, lidList, lidGroup
	j = 0
	k = 0
	ucf = UCF(RAINDEPTH)
	inflow = 0
	outflow = 0
	err = 0
	lidUnit = nil
	lidList = nil
	lidGroup = nil
	for j = 0; j < GroupCount; j++ {
		if LidGroups[j] != nil {
			k++
		}
	}
	if k == 0 {
		return
	}
	R.Fprintf(Frpt.file, L291)
	R.Fprintf(Frpt.file, L292)
	if UnitSystem == US {
		R.Fprintf(Frpt.file, L293)
	} else {
		R.Fprintf(Frpt.file, L294)
	}
	R.Fprintf(Frpt.file, L295)
	for j = 0; j < GroupCount; j++ {
		lidGroup = LidGroups[j]
		if (!(lidGroup != nil)) || (Subcatch[j].lidArea == float64(0.0)) {
			continue
		}
		lidList = lidGroup.lidList
		for lidList != nil {
			lidUnit = lidList.lidUnit
			k = lidUnit.lidIndex
			R.Fprintf(Frpt.file, L296, Subcatch[j].ID, LidProcs[k].ID)
			R.Fprintf(Frpt.file, L297, float64(lidUnit.waterBalance.inflow*ucf), float64(lidUnit.waterBalance.evap*ucf), float64(lidUnit.waterBalance.infil*ucf), float64(lidUnit.waterBalance.surfFlow*ucf), float64(lidUnit.waterBalance.drainFlow*ucf), float64(lidUnit.waterBalance.initVol*ucf), float64(lidUnit.waterBalance.finalVol*ucf))
			inflow = float64(lidUnit.waterBalance.initVol + lidUnit.waterBalance.inflow)
			outflow = float64(float64(float64(float64(lidUnit.waterBalance.finalVol+lidUnit.waterBalance.evap)+lidUnit.waterBalance.infil)+lidUnit.waterBalance.surfFlow) + lidUnit.waterBalance.drainFlow)
			if inflow > float64(0.0) {
				err = float64((float64(inflow - outflow)) / inflow)
			} else {
				err = float64(1.0)
			}
			R.Fprintf(Frpt.file, L298, float64(err*float64(100.0)))
			lidList = lidList.nextLidUnit
		}
	}
}

// lid.c:1972
func initLidRptFile(title R.CPtr, lidID R.CPtr, subcatchID R.CPtr, lidUnit *TLidUnit) {
	var (
		i int32
		f *R.CFile
	)
	_, _ = i, f
	i = 0
	f = lidUnit.rptFile.file
	if f == nil {
		return
	}
	R.Fprintf(f, L299)
	R.Fprintf(f, L300, title)
	R.Fprintf(f, L301, lidID, subcatchID)
	for i = 0; i < initLidRptFile_colCount; i++ {
		R.Fprintf(f, L169, initLidRptFile_head1[i])
	}
	for i = 0; i < initLidRptFile_colCount; i++ {
		R.Fprintf(f, L169, initLidRptFile_head2[i])
	}
	if UnitSystem == US {
		for i = 0; i < initLidRptFile_colCount; i++ {
			R.Fprintf(f, L169, initLidRptFile_units1[i])
		}
	} else {
		for i = 0; i < initLidRptFile_colCount; i++ {
			R.Fprintf(f, L169, initLidRptFile_units2[i])
		}
	}
	R.Fprintf(f, L302)
	for i = 1; i < initLidRptFile_colCount; i++ {
		R.Fprintf(f, L303, initLidRptFile_line9)
	}
	lidUnit.rptFile.wasDry = 1
	sstrncpy(R.CP(lidUnit.rptFile.results[:]), L0, 0)
}
