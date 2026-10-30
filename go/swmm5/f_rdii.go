// Generated from EPA SWMM 5.2.4 (rdii.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

var ZERO_RDII float64 = float64(0.0001)
var FileStamp = R.BufFrom(L352, 11)
var UHGroup []TUHGroup
var RdiiStep int32
var NumRdiiNodes int32
var RdiiNodeIndex []int32
var RdiiNodeFlow []float32
var RdiiFlowUnits int32
var RdiiStartDate float64
var RdiiEndDate float64
var TotalRainVol float64
var TotalRdiiVol float64
var RdiiFileType int32

// rdii.c:152
func rdii_readRdiiInflow(tok []R.CPtr, ntoks int32) int32 {
	var (
		j      int32
		k      int32
		a_2    float64
		inflow *TRdiiInflow
	)
	_, _, _, _ = j, k, a_2, inflow
	j = 0
	k = 0
	a_2 = 0
	inflow = nil
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(NODE, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	k = project_findObject(UNITHYD, tok[1])
	if k < 0 {
		return error_setInpError(ERR_NAME, tok[1])
	}
	if (!(getDouble(tok[2], &a_2) != 0)) || (a_2 < float64(0.0)) {
		return error_setInpError(ERR_NUMBER, tok[2])
	}
	inflow = Node[j].rdiiInflow
	if inflow == nil {
		inflow = new(TRdiiInflow)
		if !(inflow != nil) {
			return error_setInpError(ERR_MEMORY, L0)
		}
	}
	inflow.unitHyd = k
	inflow.area = float64(a_2 / UCF(LANDAREA))
	Node[j].rdiiInflow = inflow
	return 0
}

// rdii.c:198
func rdii_initUnitHyd(j int32) {
	var (
		i int32
		m int32
	)
	_, _ = i, m
	i = 0
	m = 0
	for m = 0; m < 12; m++ {
		for i = 0; i < 3; i++ {
			UnitHyd[j].iaMax[m][i] = float64(0.0)
			UnitHyd[j].iaRecov[m][i] = float64(0.0)
			UnitHyd[j].iaInit[m][i] = float64(0.0)
			UnitHyd[j].r[m][i] = float64(0.0)
			UnitHyd[j].tPeak[m][i] = 0
			UnitHyd[j].tBase[m][i] = 0
		}
	}
}

// rdii.c:224
func rdii_readUnitHydParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		i int32
		j int32
		k int32
		m int32
		g int32
		x [6]float64
	)
	_, _, _, _, _, _ = i, j, k, m, g, x
	i = 0
	j = 0
	k = 0
	m = 0
	g = 0
	x = ([6]float64{})
	j = project_findObject(UNITHYD, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	if UnitHyd[j].ID.IsNull() {
		UnitHyd[j].ID = project_findID(UNITHYD, tok[0])
	}
	if ntoks == 2 {
		g = project_findObject(GAGE, tok[1])
		if g < 0 {
			return error_setInpError(ERR_NAME, tok[1])
		}
		UnitHyd[j].rainGage = g
		Gage[g].isUsed = 1
		return 0
	} else if ntoks < 6 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	m = datetime_findMonth(tok[1])
	if m == 0 {
		if !(match(tok[1], L351) != 0) {
			return error_setInpError(ERR_KEYWORD, tok[1])
		}
	}
	k = findmatch(tok[2], UHTypeWords[:])
	if k < 0 {
		return readOldUHFormat(j, m, tok, ntoks)
	}
	for i = 0; i < 3; i++ {
		if !(getDouble(tok[(i+3)], &x[i]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[(i+3)])
		}
	}
	for i = 3; i < 6; i++ {
		x[i] = float64(0.0)
		if ntoks > (i + 3) {
			if !(getDouble(tok[(i+3)], &x[i]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[(i+2)])
			}
		}
	}
	setUnitHydParams(j, k, m, x[:])
	return 0
}

// rdii.c:293
func readOldUHFormat(j int32, m int32, tok []R.CPtr, ntoks int32) int32 {
	var (
		i int32
		k int32
		p [9]float64
		x [6]float64
	)
	_, _, _, _ = i, k, p, x
	i = 0
	k = 0
	p = ([9]float64{})
	x = ([6]float64{})
	if ntoks < 11 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	for i = 0; i < 9; i++ {
		if !(getDouble(tok[(i+2)], &p[i]) != 0) {
			return error_setInpError(ERR_NUMBER, tok[(i+2)])
		}
	}
	for i = 0; i < 3; i++ {
		x[(i + 3)] = float64(0.0)
		if ntoks > (i + 11) {
			if !(getDouble(tok[(i+11)], &x[(i+3)]) != 0) {
				return error_setInpError(ERR_NUMBER, tok[(i+11)])
			}
		}
	}
	for k = 0; k < 3; k++ {
		for i = 0; i < 3; i++ {
			x[i] = p[((3 * k) + i)]
			setUnitHydParams(j, k, m, x[:])
		}
	}
	return 0
}

// rdii.c:342
func setUnitHydParams(j int32, i int32, m int32, x []float64) {
	var (
		m1    int32
		m2    int32
		t     float64
		k     float64
		tBase float64
	)
	_, _, _, _, _ = m1, m2, t, k, tBase
	m1 = 0
	m2 = 0
	t = 0
	k = 0
	tBase = 0
	if m == 0 {
		m1 = 0
		m2 = 11
	} else {
		m1 = (m - 1)
		m2 = m1
	}
	for m = m1; m <= m2; m++ {
		UnitHyd[j].r[m][i] = x[0]
		t = x[1]
		k = x[2]
		tBase = float64(t * (float64(float64(1.0) + k)))
		UnitHyd[j].tPeak[m][i] = R.D2L((float64(t * float64(3600.0))))
		UnitHyd[j].tBase[m][i] = R.D2L((float64(tBase * float64(3600.0))))
		UnitHyd[j].iaMax[m][i] = x[3]
		UnitHyd[j].iaRecov[m][i] = x[4]
		UnitHyd[j].iaInit[m][i] = x[5]
	}
}

// rdii.c:389
func rdii_deleteRdiiInflow(j int32) {
	if Node[j].rdiiInflow != nil {
		_ = Node[j].rdiiInflow
		Node[j].rdiiInflow = nil
	}
}

// rdii.c:408
func rdii_openRdii() {
	var (
		fStamp R.CPtr
	)
	_ = fStamp
	fStamp = R.BufFrom(L352, 11)
	RdiiNodeIndex = nil
	RdiiNodeFlow = nil
	NumRdiiNodes = 0
	RdiiStartDate = float64((-693594))
	if IgnoreRDII != 0 {
		return
	}
	if Frdii.mode != USE_FILE {
		createRdiiFile()
	}
	if (Frdii.mode == NO_FILE) || (ErrorCode != 0) {
		return
	}
	Frdii.file = R.Fopen(R.CP(Frdii.name[:]), L353)
	if Frdii.file == nil {
		if Frdii.mode == SCRATCH_FILE {
			report_writeErrorMsg(ERR_RDII_FILE_SCRATCH, L0)
		} else {
			report_writeErrorMsg(ERR_RDII_FILE_OPEN, R.CP(Frdii.name[:]))
		}
		return
	}
	R.FreadC(fStamp, R.Strlen(FileStamp), Frdii.file)
	if R.Strcmp(fStamp, FileStamp) == 0 {
		RdiiFileType = BINARY
		ErrorCode = readRdiiFileHeader()
	} else {
		R.Fclose(Frdii.file)
		RdiiFileType = TEXT
		openRdiiTextFile()
	}
	if ErrorCode != 0 {
		report_writeErrorMsg(ErrorCode, R.CP(Frdii.name[:]))
	} else {
		readRdiiFlows()
	}
}

// rdii.c:470
func openRdiiTextFile() {
	Frdii.file = R.Fopen(R.CP(Frdii.name[:]), L13)
	if Frdii.file == nil {
		if Frdii.mode == SCRATCH_FILE {
			report_writeErrorMsg(ERR_RDII_FILE_SCRATCH, L0)
		} else {
			report_writeErrorMsg(ERR_RDII_FILE_OPEN, R.CP(Frdii.name[:]))
		}
		return
	}
	ErrorCode = readRdiiTextFileHeader()
	if ErrorCode != 0 {
		report_writeErrorMsg(ErrorCode, R.CP(Frdii.name[:]))
	}
}

// rdii.c:497
func rdii_closeRdii() {
	if Frdii.file != nil {
		R.Fclose(Frdii.file)
	}
	if Frdii.mode == SCRATCH_FILE {
		R.Remove(R.CP(Frdii.name[:]))
	}
	{
		if RdiiNodeIndex != nil {
			_ = RdiiNodeIndex
			RdiiNodeIndex = nil
		}
	}
	{
		if RdiiNodeFlow != nil {
			_ = RdiiNodeFlow
			RdiiNodeFlow = nil
		}
	}
}

// rdii.c:512
func rdii_getNumRdiiFlows(aDate float64) int32 {
	if NumRdiiNodes == 0 {
		return 0
	}
	if !(Frdii.file != nil) {
		return 0
	}
	for !(R.Feof(Frdii.file) != 0) {
		if RdiiStartDate == float64((-693594)) {
			return 0
		}
		if aDate < RdiiStartDate {
			return 0
		}
		if aDate < RdiiEndDate {
			return NumRdiiNodes
		} else {
			readRdiiFlows()
		}
	}
	return 0
}

// rdii.c:542
func rdii_getRdiiFlow(i int32, j *int32, q *float64) {
	if (i >= 0) && (i < NumRdiiNodes) {
		*j = RdiiNodeIndex[i]
		*q = float64(RdiiNodeFlow[i])
	}
}

// rdii.c:560
func readRdiiFileHeader() int32 {
	var (
		i int32
		j int32
	)
	_, _ = i, j
	i = 0
	j = 0
	R.Fread1(&RdiiStep, 4, Frdii.file)
	if RdiiStep <= 0 {
		return ERR_RDII_FILE_FORMAT
	}
	R.Fread1(&NumRdiiNodes, 4, Frdii.file)
	if NumRdiiNodes <= 0 {
		return ERR_RDII_FILE_FORMAT
	}
	RdiiNodeIndex = make([]int32, int(uint64(NumRdiiNodes)))
	if !(RdiiNodeIndex != nil) {
		return ERR_MEMORY
	}
	RdiiNodeFlow = make([]float32, int(uint64(NumRdiiNodes)))
	if !(RdiiNodeFlow != nil) {
		return ERR_MEMORY
	}
	if R.Feof(Frdii.file) != 0 {
		return ERR_RDII_FILE_FORMAT
	}
	R.Fread(RdiiNodeIndex, uint64(NumRdiiNodes), Frdii.file)
	for i = 0; i < NumRdiiNodes; i++ {
		j = RdiiNodeIndex[i]
		if Node[j].rdiiInflow == nil {
			return ERR_RDII_FILE_FORMAT
		}
	}
	if R.Feof(Frdii.file) != 0 {
		return ERR_RDII_FILE_FORMAT
	}
	return 0
}

// rdii.c:595
func readRdiiTextFileHeader() int32 {
	var (
		i    int32
		j    int32
		line R.CPtr
		s1   R.CPtr
		s2   R.CPtr
	)
	_, _, _, _, _ = i, j, line, s1, s2
	i = 0
	j = 0
	line = R.Buf(1025)
	s1 = R.BufFrom(L0, 1025)
	s2 = R.Buf(1025)
	R.Fgets(line, 1024, Frdii.file)
	if (!(R.Sscanf(line, L169, s1) != 0)) || (R.Strcmp(s1, L181) != 0) {
		return ERR_RDII_FILE_FORMAT
	}
	R.Fgets(line, 1024, Frdii.file)
	RdiiStep = 0
	R.Fgets(line, 1024, Frdii.file)
	if (!(R.Sscanf(line, L182, &RdiiStep) != 0)) || (RdiiStep <= 0) {
		return ERR_RDII_FILE_FORMAT
	}
	R.Fgets(line, 1024, Frdii.file)
	R.Fgets(line, 1024, Frdii.file)
	if R.Sscanf(line, L44, s1, s2) < 2 {
		return ERR_RDII_FILE_FORMAT
	}
	RdiiFlowUnits = findmatch(s2, FlowUnitWords[:])
	if RdiiFlowUnits < 0 {
		return ERR_RDII_FILE_FORMAT
	}
	R.Fgets(line, 1024, Frdii.file)
	if (R.Sscanf(line, L182, &NumRdiiNodes) < 1) || (NumRdiiNodes <= 0) {
		return ERR_RDII_FILE_FORMAT
	}
	RdiiNodeIndex = make([]int32, int(uint64(NumRdiiNodes)))
	if !(RdiiNodeIndex != nil) {
		return ERR_MEMORY
	}
	RdiiNodeFlow = make([]float32, int(uint64(NumRdiiNodes)))
	if !(RdiiNodeFlow != nil) {
		return ERR_MEMORY
	}
	for i = 0; i < NumRdiiNodes; i++ {
		if R.Feof(Frdii.file) != 0 {
			return ERR_RDII_FILE_FORMAT
		}
		R.Fgets(line, 1024, Frdii.file)
		if !(R.Sscanf(line, L169, s1) != 0) {
			return ERR_RDII_FILE_FORMAT
		}
		j = project_findObject(NODE, s1)
		if j < 0 {
			return ERR_RDII_FILE_FORMAT
		}
		RdiiNodeIndex[i] = j
	}
	if R.Feof(Frdii.file) != 0 {
		return ERR_RDII_FILE_FORMAT
	}
	R.Fgets(line, 1024, Frdii.file)
	return 0
}

// rdii.c:663
func readRdiiFlows() {
	if RdiiFileType == TEXT {
		readRdiiTextFlows()
	} else {
		RdiiStartDate = float64((-693594))
		RdiiEndDate = float64((-693594))
		if R.Feof(Frdii.file) != 0 {
			return
		}
		R.Fread1(&RdiiStartDate, 8, Frdii.file)
		if RdiiStartDate == float64((-693594)) {
			return
		}
		if R.Fread(RdiiNodeFlow, uint64(NumRdiiNodes), Frdii.file) < uint64(NumRdiiNodes) {
			RdiiStartDate = float64((-693594))
		} else {
			RdiiEndDate = datetime_addSeconds(RdiiStartDate, float64(RdiiStep))
		}
	}
}

// rdii.c:686
func readRdiiTextFlows() {
	var (
		i     int32
		n_2   int32
		yr    int32
		mon   int32
		day   int32
		hr    int32
		min_2 int32
		sec   int32
		x     float64
		line  R.CPtr
		s     R.CPtr
	)
	_, _, _, _, _, _, _, _, _, _, _ = i, n_2, yr, mon, day, hr, min_2, sec, x, line, s
	i = 0
	n_2 = 0
	yr = 0
	mon = 0
	day = 0
	hr = 0
	min_2 = 0
	sec = 0
	x = 0
	line = R.Buf(1025)
	s = R.Buf(1025)
	RdiiStartDate = float64((-693594))
	for i = 0; i < NumRdiiNodes; i++ {
		if R.Feof(Frdii.file) != 0 {
			return
		}
		R.Fgets(line, 1024, Frdii.file)
		n_2 = R.Sscanf(line, L354, s, &yr, &mon, &day, &hr, &min_2, &sec, &x)
		if n_2 < 8 {
			return
		}
		RdiiNodeFlow[i] = float32(float64((float64(x / Qcf[RdiiFlowUnits]))))
	}
	RdiiStartDate = float64(datetime_encodeDate(yr, mon, day) + datetime_encodeTime(hr, min_2, sec))
	RdiiEndDate = datetime_addSeconds(RdiiStartDate, float64(RdiiStep))
}

// rdii.c:720
func createRdiiFile() {
	var (
		hasRdii     int32
		elapsedTime float64
		duration    float64
		currentDate float64
	)
	_, _, _, _ = hasRdii, elapsedTime, duration, currentDate
	hasRdii = 0
	elapsedTime = 0
	duration = 0
	currentDate = 0
	RdiiStep = WetStep
	NumRdiiNodes = getNumRdiiNodes()
	if NumRdiiNodes == 0 {
		Frdii.mode = NO_FILE
		return
	} else if Frdii.mode == NO_FILE {
		Frdii.mode = SCRATCH_FILE
	}
	validateRdii()
	initGageData()
	if ErrorCode != 0 {
		return
	}
	openRdiiProcessor()
	if !(ErrorCode != 0) {
		initUnitHydData()
		duration = float64(TotalDuration / float64(1000.0))
		elapsedTime = float64(0.0)
		for (elapsedTime <= duration) && (!(ErrorCode != 0)) {
			currentDate = float64(StartDateTime + float64(elapsedTime/float64(86400)))
			getRainfall(currentDate)
			getUnitHydRdii(currentDate)
			hasRdii = getNodeRdii()
			if hasRdii != 0 {
				saveRdiiFlows(currentDate)
			}
			elapsedTime = float64(elapsedTime + (float64(RdiiStep)))
		}
	}
	closeRdiiProcessor()
}

// rdii.c:793
func getNumRdiiNodes() int32 {
	var (
		j   int32
		n_2 int32
	)
	_, _ = j, n_2
	j = 0
	n_2 = 0
	n_2 = 0
	for j = 0; j < Nobjects[NODE]; j++ {
		if Node[j].rdiiInflow != nil {
			n_2++
		}
	}
	return n_2
}

// rdii.c:813
func validateRdii() {
	var (
		i    int32
		j    int32
		k    int32
		m    int32
		rsum float64
	)
	_, _, _, _, _ = i, j, k, m, rsum
	i = 0
	j = 0
	k = 0
	m = 0
	rsum = 0
	for j = 0; j < Nobjects[UNITHYD]; j++ {
		for m = 0; m < 12; m++ {
			rsum = float64(0.0)
			for k = 0; k < 3; k++ {
				if UnitHyd[j].tBase[m][k] == 0 {
					continue
				}
				if float64(UnitHyd[j].tPeak[m][k]) < float64(0.0) {
					report_writeErrorMsg(ERR_UNITHYD_TIMES, UnitHyd[j].ID)
				}
				if UnitHyd[j].r[m][k] < float64(0.0) {
					report_writeErrorMsg(ERR_UNITHYD_RATIOS, UnitHyd[j].ID)
				} else {
					rsum = float64(rsum + UnitHyd[j].r[m][k])
				}
			}
			if rsum > float64(1.01) {
				report_writeErrorMsg(ERR_UNITHYD_RATIOS, UnitHyd[j].ID)
			}
		}
	}
	for i = 0; i < Nobjects[NODE]; i++ {
		if Node[i].rdiiInflow != nil {
			if Node[i].rdiiInflow.area < float64(0.0) {
				report_writeErrorMsg(ERR_RDII_AREA, Node[i].ID)
			}
		}
	}
}

// rdii.c:877
func openRdiiProcessor() {
	var (
		j   int32
		n_2 int32
	)
	_, _ = j, n_2
	j = 0
	n_2 = 0
	UHGroup = nil
	RdiiNodeIndex = nil
	RdiiNodeFlow = nil
	TotalRainVol = float64(0.0)
	TotalRdiiVol = float64(0.0)
	if !(allocRdiiMemory() != 0) {
		report_writeErrorMsg(ERR_MEMORY, L0)
		return
	}
	if !(openNewRdiiFile() != 0) {
		report_writeErrorMsg(ERR_RDII_FILE_SCRATCH, L0)
		return
	}
	n_2 = 0
	for j = 0; j < Nobjects[NODE]; j++ {
		if (Node[j].rdiiInflow != nil) && (RdiiNodeIndex != nil) {
			RdiiNodeIndex[n_2] = j
			n_2++
		}
	}
}

// rdii.c:922
func allocRdiiMemory() int32 {
	var (
		i   int32
		k   int32
		n_2 int32
	)
	_, _, _ = i, k, n_2
	i = 0
	k = 0
	n_2 = 0
	UHGroup = make([]TUHGroup, int(uint64(Nobjects[UNITHYD])))
	if !(UHGroup != nil) {
		return 0
	}
	for i = 0; i < Nobjects[UNITHYD]; i++ {
		UHGroup[i].rainInterval = getRainInterval(i)
		for k = 0; k < 3; k++ {
			UHGroup[i].uh[k].pastRain = nil
			UHGroup[i].uh[k].pastMonth = (R.CPtr{})
			UHGroup[i].uh[k].maxPeriods = getMaxPeriods(i, k)
			n_2 = UHGroup[i].uh[k].maxPeriods
			if n_2 > 0 {
				UHGroup[i].uh[k].pastRain = make([]float64, int(uint64(n_2)))
				if !(UHGroup[i].uh[k].pastRain != nil) {
					return 0
				}
				UHGroup[i].uh[k].pastMonth = R.Buf(int(uint64(n_2)))
				if !(!UHGroup[i].uh[k].pastMonth.IsNull()) {
					return 0
				}
			}
		}
	}
	RdiiNodeIndex = make([]int32, int(uint64(NumRdiiNodes)))
	if !(RdiiNodeIndex != nil) {
		return 0
	}
	RdiiNodeFlow = make([]float32, int(uint64(NumRdiiNodes)))
	if !(RdiiNodeFlow != nil) {
		return 0
	}
	return 1
}

// rdii.c:970
func getRainInterval(i int32) int32 {
	var (
		ri    int32
		tLimb int32
		k     int32
		m     int32
	)
	_, _, _, _ = ri, tLimb, k, m
	ri = 0
	tLimb = 0
	k = 0
	m = 0
	ri = WetStep
	for m = 0; m < 12; m++ {
		for k = 0; k < 3; k++ {
			if UnitHyd[i].tPeak[m][k] > 0 {
				tLimb = int32(UnitHyd[i].tPeak[m][k])
				if (ri) <= (tLimb) {
					ri = (ri)
				} else {
					ri = (tLimb)
				}
				tLimb = int32((UnitHyd[i].tBase[m][k] - int64(tLimb)))
				if tLimb > 0 {
					if (ri) <= (tLimb) {
						ri = (ri)
					} else {
						ri = (tLimb)
					}
				}
			}
		}
	}
	return ri
}

// rdii.c:1005
func getMaxPeriods(i int32, k int32) int32 {
	var (
		m            int32
		n_2          int32
		nMax         int32
		rainInterval int32
	)
	_, _, _, _ = m, n_2, nMax, rainInterval
	m = 0
	n_2 = 0
	nMax = 0
	rainInterval = 0
	rainInterval = UHGroup[i].rainInterval
	nMax = 0
	for m = 0; m < 12; m++ {
		n_2 = int32(((UnitHyd[i].tBase[m][k] / int64(rainInterval)) + 1))
		if (n_2) >= (nMax) {
			nMax = (n_2)
		} else {
			nMax = (nMax)
		}
	}
	return nMax
}

// rdii.c:1034
func initGageData() {
	var (
		i int32
		g int32
	)
	_, _ = i, g
	i = 0
	g = 0
	for g = 0; g < Nobjects[GAGE]; g++ {
		if Gage[g].tSeries >= 0 {
			table_tseriesInit(&Tseries[Gage[g].tSeries])
		}
		gage_initState(g)
	}
	for i = 0; i < Nobjects[UNITHYD]; i++ {
		g = UnitHyd[i].rainGage
		if g >= 0 {
			if Gage[g].coGage >= 0 {
				UnitHyd[i].rainGage = Gage[g].coGage
			}
		}
	}
}

// rdii.c:1072
func initUnitHydData() {
	var (
		i     int32
		j     int32
		k     int32
		n_2   int32
		month int32
	)
	_, _, _, _, _ = i, j, k, n_2, month
	i = 0
	j = 0
	k = 0
	n_2 = 0
	month = 0
	month = (datetime_monthOfYear(StartDateTime) - 1)
	for i = 0; i < Nobjects[UNITHYD]; i++ {
		for k = 0; k < 3; k++ {
			UHGroup[i].uh[k].drySeconds = int64(((UHGroup[i].uh[k].maxPeriods * UHGroup[i].rainInterval) + 1))
			UHGroup[i].uh[k].period = (UHGroup[i].uh[k].maxPeriods + 1)
			UHGroup[i].uh[k].hasPastRain = 0
			UHGroup[i].uh[k].iaUsed = UnitHyd[i].iaInit[month][k]
		}
		UHGroup[i].gageDate = StartDateTime
		UHGroup[i].area = float64(0.0)
		UHGroup[i].rdii = float64(0.0)
	}
	for i = 0; i < Nobjects[UNITHYD]; i++ {
		UHGroup[i].isUsed = 0
	}
	for n_2 = 0; n_2 < NumRdiiNodes; n_2++ {
		j = RdiiNodeIndex[n_2]
		i = Node[j].rdiiInflow.unitHyd
		UHGroup[i].isUsed = 1
		UHGroup[i].lastDate = StartDateTime
		UHGroup[i].area = float64(UHGroup[i].area + Node[j].rdiiInflow.area)
	}
}

// rdii.c:1128
func openNewRdiiFile() int32 {
	var (
		j int32
	)
	_ = j
	j = 0
	if Frdii.mode == SCRATCH_FILE {
		getTempFileName(R.CP(Frdii.name[:]))
	}
	Frdii.file = R.Fopen(R.CP(Frdii.name[:]), L166)
	if Frdii.file == nil {
		return 0
	}
	R.FwriteC(FileStamp, R.Strlen(FileStamp), Frdii.file)
	R.Fwrite1(RdiiStep, 4, Frdii.file)
	R.Fwrite1(NumRdiiNodes, 4, Frdii.file)
	for j = 0; j < Nobjects[NODE]; j++ {
		if Node[j].rdiiInflow != nil {
			R.Fwrite1(j, 4, Frdii.file)
		}
	}
	return 1
}

// rdii.c:1163
func getRainfall(currentDate float64) {
	var (
		j            int32
		k            int32
		g            int32
		i            int32
		month        int32
		rainInterval int32
		rainDepth    float64
		excessDepth  float64
		gageDate     float64
	)
	_, _, _, _, _, _, _, _, _ = j, k, g, i, month, rainInterval, rainDepth, excessDepth, gageDate
	j = 0
	k = 0
	g = 0
	i = 0
	month = 0
	rainInterval = 0
	rainDepth = 0
	excessDepth = 0
	gageDate = 0
	month = (datetime_monthOfYear(currentDate) - 1)
	for g = 0; g < Nobjects[GAGE]; g++ {
		Gage[g].isCurrent = 0
	}
	for j = 0; j < Nobjects[UNITHYD]; j++ {
		g = UnitHyd[j].rainGage
		rainInterval = UHGroup[j].rainInterval
		for UHGroup[j].gageDate < currentDate {
			gageDate = UHGroup[j].gageDate
			Adjust.rainFactor = Adjust.rain[(datetime_monthOfYear(gageDate) - 1)]
			if !(Gage[g].isCurrent != 0) {
				gage_setState(g, gageDate)
				Gage[g].isCurrent = 1
			}
			rainDepth = float64(float64(Gage[g].rainfall*float64(rainInterval)) / float64(3600.0))
			TotalRainVol = float64(TotalRainVol + (float64(float64(rainDepth/UCF(RAINDEPTH)) * UHGroup[j].area)))
			for k = 0; k < 3; k++ {
				excessDepth = applyIA(j, k, gageDate, float64(rainInterval), rainDepth)
				updateDryPeriod(j, k, excessDepth, rainInterval)
				i = UHGroup[j].uh[k].period
				if i >= UHGroup[j].uh[k].maxPeriods {
					i = 0
				}
				UHGroup[j].uh[k].pastRain[i] = excessDepth
				UHGroup[j].uh[k].pastMonth.S(i, int32(int8(month)))
				UHGroup[j].uh[k].period = (i + 1)
			}
			UHGroup[j].gageDate = datetime_addSeconds(gageDate, float64(rainInterval))
		}
	}
}

// rdii.c:1231
func applyIA(j int32, k int32, aDate float64, dt float64, rainDepth float64) float64 {
	var (
		m            int32
		ia           float64
		netRainDepth float64
	)
	_, _, _ = m, ia, netRainDepth
	m = 0
	ia = 0
	netRainDepth = 0
	m = (datetime_monthOfYear(aDate) - 1)
	ia = float64(UnitHyd[j].iaMax[m][k] - UHGroup[j].uh[k].iaUsed)
	if (ia) >= (float64(0.0)) {
		ia = (ia)
	} else {
		ia = (float64(0.0))
	}
	if rainDepth > float64(0.0) {
		netRainDepth = float64(rainDepth - ia)
		if (netRainDepth) >= (float64(0.0)) {
			netRainDepth = (netRainDepth)
		} else {
			netRainDepth = (float64(0.0))
		}
		ia = float64(rainDepth - netRainDepth)
		UHGroup[j].uh[k].iaUsed = float64(UHGroup[j].uh[k].iaUsed + ia)
	} else {
		UHGroup[j].uh[k].iaUsed = float64(UHGroup[j].uh[k].iaUsed - (float64(float64(dt/float64(86400.0)) * UnitHyd[j].iaRecov[m][k])))
		if (UHGroup[j].uh[k].iaUsed) >= (float64(0.0)) {
			UHGroup[j].uh[k].iaUsed = (UHGroup[j].uh[k].iaUsed)
		} else {
			UHGroup[j].uh[k].iaUsed = (float64(0.0))
		}
		netRainDepth = float64(0.0)
	}
	return netRainDepth
}

// rdii.c:1276
func updateDryPeriod(j int32, k int32, rainDepth float64, rainInterval int32) {
	var (
		i int32
	)
	_ = i
	i = 0
	if rainDepth > float64(0.0) {
		if UHGroup[j].uh[k].drySeconds >= int64((rainInterval * UHGroup[j].uh[k].maxPeriods)) {
			for i = 0; i < UHGroup[j].uh[k].maxPeriods; i++ {
				UHGroup[j].uh[k].pastRain[i] = float64(0.0)
			}
			UHGroup[j].uh[k].period = 0
		}
		UHGroup[j].uh[k].drySeconds = 0
		UHGroup[j].uh[k].hasPastRain = 1
	} else {
		UHGroup[j].uh[k].drySeconds += (int64(rainInterval))
		if UHGroup[j].uh[k].drySeconds >= int64((rainInterval * UHGroup[j].uh[k].maxPeriods)) {
			UHGroup[j].uh[k].hasPastRain = 0
		} else {
			UHGroup[j].uh[k].hasPastRain = 1
		}
	}
}

// rdii.c:1321
func getUnitHydRdii(currentDate float64) {
	var (
		j            int32
		k            int32
		rainInterval int32
	)
	_, _, _ = j, k, rainInterval
	j = 0
	k = 0
	rainInterval = 0
	for j = 0; j < Nobjects[UNITHYD]; j++ {
		if !(UHGroup[j].isUsed != 0) {
			continue
		}
		if currentDate < UHGroup[j].lastDate {
			continue
		}
		UHGroup[j].lastDate = UHGroup[j].gageDate
		rainInterval = UHGroup[j].rainInterval
		UHGroup[j].rdii = float64(0.0)
		for k = 0; k < 3; k++ {
			if UHGroup[j].uh[k].hasPastRain != 0 {
				UHGroup[j].rdii = float64(UHGroup[j].rdii + (getUnitHydConvol(j, k, rainInterval)))
			}
		}
	}
}

// rdii.c:1358
func getUnitHydConvol(j int32, k int32, rainInterval int32) float64 {
	var (
		i    int32
		m    int32
		p    int32
		pMax int32
		t    float64
		u    float64
		v    float64
		rdii float64
		uh   *TUHData
	)
	_, _, _, _, _, _, _, _, _ = i, m, p, pMax, t, u, v, rdii, uh
	i = 0
	m = 0
	p = 0
	pMax = 0
	t = 0
	u = 0
	v = 0
	rdii = 0
	uh = nil
	rdii = float64(0.0)
	uh = &UHGroup[j].uh[k]
	i = (uh.period - 1)
	if i < 0 {
		i = (uh.maxPeriods - 1)
	}
	pMax = uh.maxPeriods
	p = 1
	for p < pMax {
		v = uh.pastRain[i]
		m = uh.pastMonth.G(i)
		if v > float64(0.0) {
			t = float64((float64(float64((p)) - float64(0.5))) * float64(rainInterval))
			u = float64(getUnitHydOrd(j, m, k, t) * UnitHyd[j].r[m][k])
			rdii = float64(rdii + (float64(u * v)))
		}
		p = (p + 1)
		i = (i - 1)
		if i < 0 {
			i = (uh.maxPeriods - 1)
		}
	}
	return rdii
}

// rdii.c:1411
func getUnitHydOrd(h int32, m int32, k int32, t float64) float64 {
	var (
		qPeak float64
		f     float64
		t1    float64
		t2    float64
		tBase float64
	)
	_, _, _, _, _ = qPeak, f, t1, t2, tBase
	qPeak = 0
	f = 0
	t1 = 0
	t2 = 0
	tBase = 0
	tBase = float64(UnitHyd[h].tBase[m][k])
	if t >= tBase {
		return float64(0.0)
	}
	qPeak = float64(float64(float64(2.0)/tBase) * float64(3600.0))
	t1 = float64(UnitHyd[h].tPeak[m][k])
	t2 = float64(tBase - t1)
	if t <= t1 {
		f = float64(t / t1)
	} else {
		f = float64(float64(1.0) - float64((float64(t-t1))/t2))
	}
	return float64((R.If[float64](((f) >= (float64(0.0))), (f), (float64(0.0)))) * qPeak)
}

// rdii.c:1446
func getNodeRdii() int32 {
	var (
		hasRdii int32
		i       int32
		j       int32
		n_2     int32
		rdii    float64
	)
	_, _, _, _, _ = hasRdii, i, j, n_2, rdii
	hasRdii = 0
	i = 0
	j = 0
	n_2 = 0
	rdii = 0
	for n_2 = 0; n_2 < NumRdiiNodes; n_2++ {
		j = RdiiNodeIndex[n_2]
		i = Node[j].rdiiInflow.unitHyd
		rdii = float64(float64(UHGroup[i].rdii*Node[j].rdiiInflow.area) / UCF(RAINFALL))
		if rdii < ZERO_RDII {
			rdii = float64(0.0)
		} else {
			hasRdii = 1
		}
		RdiiNodeFlow[n_2] = float32(float64(rdii))
		if rdii > float64(0.0) {
			TotalRdiiVol = float64(TotalRdiiVol + (float64(rdii * float64(RdiiStep))))
		}
	}
	return hasRdii
}

// rdii.c:1483
func saveRdiiFlows(currentDate float64) {
	R.Fwrite1(currentDate, 8, Frdii.file)
	R.Fwrite(RdiiNodeFlow, uint64(NumRdiiNodes), Frdii.file)
}

// rdii.c:1496
func closeRdiiProcessor() {
	if !(ErrorCode != 0) {
		report_writeRdiiStats(TotalRainVol, TotalRdiiVol)
	}
	freeRdiiMemory()
	if Frdii.file != nil {
		R.Fclose(Frdii.file)
	}
}

// rdii.c:1516
func freeRdiiMemory() {
	var (
		i int32
		k int32
	)
	_, _ = i, k
	i = 0
	k = 0
	if UHGroup != nil {
		for i = 0; i < Nobjects[UNITHYD]; i++ {
			for k = 0; k < 3; k++ {
				{
					if UHGroup[i].uh[k].pastRain != nil {
						_ = UHGroup[i].uh[k].pastRain
						UHGroup[i].uh[k].pastRain = nil
					}
				}
				{
					if !UHGroup[i].uh[k].pastMonth.IsNull() {
						_ = UHGroup[i].uh[k].pastMonth
						UHGroup[i].uh[k].pastMonth = (R.CPtr{})
					}
				}
			}
		}
		{
			if UHGroup != nil {
				_ = UHGroup
				UHGroup = nil
			}
		}
	}
	{
		if RdiiNodeIndex != nil {
			_ = RdiiNodeIndex
			RdiiNodeIndex = nil
		}
	}
	{
		if RdiiNodeFlow != nil {
			_ = RdiiNodeFlow
			RdiiNodeFlow = nil
		}
	}
}
