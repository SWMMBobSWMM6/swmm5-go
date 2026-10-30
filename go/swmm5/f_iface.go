// Generated from EPA SWMM 5.2.4 (iface.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

var IfaceFlowUnits int32
var IfaceStep int32
var NumIfacePolluts int32
var IfacePolluts []int32
var NumIfaceNodes int32
var IfaceNodes []int32
var OldIfaceValues [][]float64
var NewIfaceValues [][]float64
var IfaceFrac float64
var OldIfaceDate float64
var NewIfaceDate float64

// iface.c:66
func iface_readFileParams(tok []R.CPtr, ntoks int32) int32 {
	var (
		k     int32
		j     int32
		fname R.CPtr
	)
	_, _, _ = k, j, fname
	k = 0
	j = 0
	fname = R.Buf(260)
	if ntoks < 2 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	k = int32(int8(findmatch(tok[0], FileModeWords[:])))
	if k < 0 {
		return error_setInpError(ERR_KEYWORD, tok[0])
	}
	j = findmatch(tok[1], FileTypeWords[:])
	if j < 0 {
		return error_setInpError(ERR_KEYWORD, tok[1])
	}
	if ntoks < 3 {
		return 0
	}
	sstrncpy(fname, tok[2], 259)
	switch j {
	case RAINFALL_FILE:
		Frain.mode = k
		sstrncpy(R.CP(Frain.name[:]), addAbsolutePath(fname), 259)
		break
	case RUNOFF_FILE:
		Frunoff.mode = k
		sstrncpy(R.CP(Frunoff.name[:]), addAbsolutePath(fname), 259)
		break
	case HOTSTART_FILE:
		if k == USE_FILE {
			Fhotstart1.mode = k
			sstrncpy(R.CP(Fhotstart1.name[:]), addAbsolutePath(fname), 259)
		} else if k == SAVE_FILE {
			Fhotstart2.mode = k
			sstrncpy(R.CP(Fhotstart2.name[:]), addAbsolutePath(fname), 259)
		}
		break
	case RDII_FILE:
		Frdii.mode = k
		sstrncpy(R.CP(Frdii.name[:]), fname, 259)
		break
	case INFLOWS_FILE:
		if k != USE_FILE {
			return error_setInpError(ERR_ITEMS, L0)
		}
		Finflows.mode = k
		sstrncpy(R.CP(Finflows.name[:]), addAbsolutePath(fname), 259)
		break
	case OUTFLOWS_FILE:
		if k != SAVE_FILE {
			return error_setInpError(ERR_ITEMS, L0)
		}
		Foutflows.mode = k
		sstrncpy(R.CP(Foutflows.name[:]), addAbsolutePath(fname), 259)
		break
	}
	return 0
}

// iface.c:138
func iface_openRoutingFiles() {
	NumIfacePolluts = 0
	IfacePolluts = nil
	NumIfaceNodes = 0
	IfaceNodes = nil
	OldIfaceValues = nil
	NewIfaceValues = nil
	if (Foutflows.mode != NO_FILE) && (Finflows.mode != NO_FILE) {
		if strcomp(R.CP(Foutflows.name[:]), R.CP(Finflows.name[:])) != 0 {
			report_writeErrorMsg(ERR_ROUTING_FILE_NAMES, L0)
			return
		}
	}
	if Foutflows.mode == SAVE_FILE {
		openFileForOutput()
	}
	if Finflows.mode == USE_FILE {
		openFileForInput()
	}
}

// iface.c:170
func iface_closeRoutingFiles() {
	{
		if IfacePolluts != nil {
			_ = IfacePolluts
			IfacePolluts = nil
		}
	}
	{
		if IfaceNodes != nil {
			_ = IfaceNodes
			IfaceNodes = nil
		}
	}
	if OldIfaceValues != nil {
		project_freeMatrix(OldIfaceValues)
	}
	if NewIfaceValues != nil {
		project_freeMatrix(NewIfaceValues)
	}
	if Finflows.file != nil {
		R.Fclose(Finflows.file)
	}
	if Foutflows.file != nil {
		R.Fclose(Foutflows.file)
	}
}

// iface.c:187
func iface_getNumIfaceNodes(currentDate float64) int32 {
	if OldIfaceDate > currentDate {
		return 0
	}
	for (NewIfaceDate < currentDate) && (NewIfaceDate != float64((-693594))) {
		setOldIfaceValues()
		readNewIfaceValues()
	}
	if NewIfaceDate == float64((-693594)) {
		return 0
	}
	IfaceFrac = float64((float64(currentDate - OldIfaceDate)) / (float64(NewIfaceDate - OldIfaceDate)))
	if (float64(0.0)) >= (IfaceFrac) {
		IfaceFrac = (float64(0.0))
	} else {
		IfaceFrac = (IfaceFrac)
	}
	if (IfaceFrac) <= (float64(1.0)) {
		IfaceFrac = (IfaceFrac)
	} else {
		IfaceFrac = (float64(1.0))
	}
	return NumIfaceNodes
}

// iface.c:219
func iface_getIfaceNode(index int32) int32 {
	if (index >= 0) && (index < NumIfaceNodes) {
		return IfaceNodes[index]
	} else {
		return (-1)
	}
	return 0
}

// iface.c:232
func iface_getIfaceFlow(index int32) float64 {
	var (
		q1 float64
		q2 float64
	)
	_, _ = q1, q2
	q1 = 0
	q2 = 0
	if (index >= 0) && (index < NumIfaceNodes) {
		q1 = OldIfaceValues[index][0]
		q2 = NewIfaceValues[index][0]
		return float64(float64((float64(float64(1.0)-IfaceFrac))*q1) + float64(IfaceFrac*q2))
	} else {
		return float64(0.0)
	}
	return 0
}

// iface.c:253
func iface_getIfaceQual(index int32, pollut int32) float64 {
	var (
		j  int32
		c1 float64
		c2 float64
	)
	_, _, _ = j, c1, c2
	j = 0
	c1 = 0
	c2 = 0
	if (index >= 0) && (index < NumIfaceNodes) {
		j = IfacePolluts[pollut]
		if j < 0 {
			return float64(0.0)
		}
		c1 = OldIfaceValues[index][(j + 1)]
		c2 = NewIfaceValues[index][(j + 1)]
		return float64(float64((float64(float64(1.0)-IfaceFrac))*c1) + float64(IfaceFrac*c2))
	} else {
		return float64(0.0)
	}
	return 0
}

// iface.c:281
func iface_saveOutletResults(reportDate float64, file *R.CFile) {
	var (
		i       int32
		p       int32
		yr      int32
		mon     int32
		day     int32
		hr      int32
		min_2   int32
		sec     int32
		theDate R.CPtr
	)
	_, _, _, _, _, _, _, _, _ = i, p, yr, mon, day, hr, min_2, sec, theDate
	i = 0
	p = 0
	yr = 0
	mon = 0
	day = 0
	hr = 0
	min_2 = 0
	sec = 0
	theDate = R.Buf(26)
	datetime_decodeDate(reportDate, &yr, &mon, &day)
	datetime_decodeTime(reportDate, &hr, &min_2, &sec)
	R.Snprintf(theDate, 26, L167, yr, mon, day, hr, min_2, sec)
	for i = 0; i < Nobjects[NODE]; i++ {
		if !(isOutletNode(i) != 0) {
			continue
		}
		R.Fprintf(file, L168, Node[i].ID)
		R.Fprintf(file, L169, theDate)
		R.Fprintf(file, L170, float64(Node[i].inflow*UCF(FLOW)))
		for p = 0; p < Nobjects[POLLUT]; p++ {
			R.Fprintf(file, L170, Node[i].newQual[p])
		}
	}
}

// iface.c:313
func openFileForOutput() {
	var (
		i   int32
		n_2 int32
	)
	_, _ = i, n_2
	i = 0
	n_2 = 0
	Foutflows.file = R.Fopen(R.CP(Foutflows.name[:]), L171)
	if Foutflows.file == nil {
		report_writeErrorMsg(ERR_ROUTING_FILE_OPEN, R.CP(Foutflows.name[:]))
		return
	}
	R.Fprintf(Foutflows.file, L172)
	R.Fprintf(Foutflows.file, L173, Title[0])
	R.Fprintf(Foutflows.file, L174, ReportStep)
	R.Fprintf(Foutflows.file, L175, (Nobjects[POLLUT] + 1))
	R.Fprintf(Foutflows.file, L176, FlowUnitWords[FlowUnits])
	for i = 0; i < Nobjects[POLLUT]; i++ {
		R.Fprintf(Foutflows.file, L177, Pollut[i].ID, QualUnitsWords[Pollut[i].units])
	}
	n_2 = 0
	for i = 0; i < Nobjects[NODE]; i++ {
		if isOutletNode(i) != 0 {
			n_2++
		}
	}
	R.Fprintf(Foutflows.file, L178, n_2)
	for i = 0; i < Nobjects[NODE]; i++ {
		if isOutletNode(i) != 0 {
			R.Fprintf(Foutflows.file, L173, Node[i].ID)
		}
	}
	R.Fprintf(Foutflows.file, L179)
	for i = 0; i < Nobjects[POLLUT]; i++ {
		R.Fprintf(Foutflows.file, L180, Pollut[i].ID)
	}
	if ReportStart == StartDateTime {
		iface_saveOutletResults(ReportStart, Foutflows.file)
	}
}

// iface.c:377
func openFileForInput() {
	var (
		err  int32
		line R.CPtr
		s    R.CPtr
	)
	_, _, _ = err, line, s
	err = 0
	line = R.Buf(1025)
	s = R.Buf(1025)
	Finflows.file = R.Fopen(R.CP(Finflows.name[:]), L13)
	if Finflows.file == nil {
		report_writeErrorMsg(ERR_ROUTING_FILE_OPEN, R.CP(Finflows.name[:]))
		return
	}
	R.Fgets(line, 1024, Finflows.file)
	if (!(R.Sscanf(line, L169, s) != 0)) || (!(strcomp(s, L181) != 0)) {
		report_writeErrorMsg(ERR_ROUTING_FILE_FORMAT, R.CP(Finflows.name[:]))
		return
	}
	R.Fgets(line, 1024, Finflows.file)
	IfaceStep = 0
	R.Fgets(line, 1024, Finflows.file)
	if (!(R.Sscanf(line, L182, &IfaceStep) != 0)) || (IfaceStep <= 0) {
		report_writeErrorMsg(ERR_ROUTING_FILE_FORMAT, R.CP(Finflows.name[:]))
		return
	}
	err = getIfaceFilePolluts()
	if err > 0 {
		report_writeErrorMsg(err, R.CP(Finflows.name[:]))
		return
	}
	err = getIfaceFileNodes()
	if err > 0 {
		report_writeErrorMsg(err, R.CP(Finflows.name[:]))
		return
	}
	OldIfaceValues = project_createMatrix(NumIfaceNodes, (1 + NumIfacePolluts))
	NewIfaceValues = project_createMatrix(NumIfaceNodes, (1 + NumIfacePolluts))
	if (OldIfaceValues == nil) || (NewIfaceValues == nil) {
		report_writeErrorMsg(ERR_MEMORY, L0)
		return
	}
	readNewIfaceValues()
	OldIfaceDate = NewIfaceDate
}

// iface.c:450
func getIfaceFilePolluts() int32 {
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
	s1 = R.Buf(1025)
	s2 = R.Buf(1025)
	R.Fgets(line, 1024, Finflows.file)
	NumIfacePolluts = (-1)
	if R.Sscanf(line, L182, &NumIfacePolluts) != 0 {
		NumIfacePolluts--
	}
	if NumIfacePolluts < 0 {
		return ERR_ROUTING_FILE_FORMAT
	}
	R.Fgets(line, 1024, Finflows.file)
	if R.Sscanf(line, L44, s1, s2) < 2 {
		return ERR_ROUTING_FILE_FORMAT
	}
	if !(strcomp(s1, L183) != 0) {
		return ERR_ROUTING_FILE_FORMAT
	}
	IfaceFlowUnits = findmatch(s2, FlowUnitWords[:])
	if IfaceFlowUnits < 0 {
		return ERR_ROUTING_FILE_FORMAT
	}
	if Nobjects[POLLUT] > 0 {
		IfacePolluts = make([]int32, int(uint64(Nobjects[POLLUT])))
		if !(IfacePolluts != nil) {
			return ERR_MEMORY
		}
		for i = 0; i < Nobjects[POLLUT]; i++ {
			IfacePolluts[i] = (-1)
		}
	}
	if (NumIfacePolluts > 0) && (Nobjects[POLLUT] > 0) {
		for i = 0; i < NumIfacePolluts; i++ {
			if R.Feof(Finflows.file) != 0 {
				return ERR_ROUTING_FILE_FORMAT
			}
			R.Fgets(line, 1024, Finflows.file)
			if R.Sscanf(line, L44, s1, s2) < 2 {
				return ERR_ROUTING_FILE_FORMAT
			}
			j = project_findObject(POLLUT, s1)
			if j < 0 {
				continue
			}
			if !(strcomp(s2, QualUnitsWords[Pollut[j].units]) != 0) {
				return ERR_ROUTING_FILE_NOMATCH
			}
			IfacePolluts[j] = i
		}
	}
	return 0
}

// iface.c:510
func getIfaceFileNodes() int32 {
	var (
		i    int32
		line R.CPtr
		s    R.CPtr
	)
	_, _, _ = i, line, s
	i = 0
	line = R.Buf(1025)
	s = R.Buf(1025)
	R.Fgets(line, 1024, Finflows.file)
	if (!(R.Sscanf(line, L182, &NumIfaceNodes) != 0)) || (NumIfaceNodes <= 0) {
		return ERR_ROUTING_FILE_FORMAT
	}
	IfaceNodes = make([]int32, int(uint64(NumIfaceNodes)))
	if !(IfaceNodes != nil) {
		return ERR_MEMORY
	}
	for i = 0; i < NumIfaceNodes; i++ {
		if R.Feof(Finflows.file) != 0 {
			return ERR_ROUTING_FILE_FORMAT
		}
		IfaceNodes[i] = 0
		R.Fgets(line, 1024, Finflows.file)
		if !(R.Sscanf(line, L169, s) != 0) {
			return ERR_ROUTING_FILE_FORMAT
		}
		IfaceNodes[i] = project_findObject(NODE, s)
	}
	if R.Feof(Finflows.file) != 0 {
		return ERR_ROUTING_FILE_FORMAT
	}
	R.Fgets(line, 1024, Finflows.file)
	return 0
}

// iface.c:550
func readNewIfaceValues() {
	var (
		i     int32
		j     int32
		s     R.CPtr
		yr    int32
		mon   int32
		day   int32
		hr    int32
		min_2 int32
		sec   int32
		line  R.CPtr
	)
	_, _, _, _, _, _, _, _, _, _ = i, j, s, yr, mon, day, hr, min_2, sec, line
	i = 0
	j = 0
	s = (R.CPtr{})
	yr = 0
	mon = 0
	day = 0
	hr = 0
	min_2 = 0
	sec = 0
	line = R.Buf(1025)
	NewIfaceDate = float64((-693594))
	for i = 0; i < NumIfaceNodes; i++ {
		if R.Feof(Finflows.file) != 0 {
			return
		}
		R.Fgets(line, 1024, Finflows.file)
		if (R.Strtok(line, L184)).IsNull() {
			return
		}
		s = R.Strtok((R.CPtr{}), L184)
		if s.IsNull() {
			return
		}
		yr = R.Atoi(s)
		s = R.Strtok((R.CPtr{}), L184)
		if s.IsNull() {
			return
		}
		mon = R.Atoi(s)
		s = R.Strtok((R.CPtr{}), L184)
		if s.IsNull() {
			return
		}
		day = R.Atoi(s)
		s = R.Strtok((R.CPtr{}), L184)
		if s.IsNull() {
			return
		}
		hr = R.Atoi(s)
		s = R.Strtok((R.CPtr{}), L184)
		if s.IsNull() {
			return
		}
		min_2 = R.Atoi(s)
		s = R.Strtok((R.CPtr{}), L184)
		if s.IsNull() {
			return
		}
		sec = R.Atoi(s)
		s = R.Strtok((R.CPtr{}), L184)
		if s.IsNull() {
			return
		}
		NewIfaceValues[i][0] = float64(R.Atof(s) / Qcf[IfaceFlowUnits])
		for j = 1; j <= NumIfacePolluts; j++ {
			s = R.Strtok((R.CPtr{}), L184)
			if s.IsNull() {
				return
			}
			NewIfaceValues[i][j] = R.Atof(s)
		}
	}
	NewIfaceDate = float64(datetime_encodeDate(yr, mon, day) + datetime_encodeTime(hr, min_2, sec))
}

// iface.c:613
func setOldIfaceValues() {
	var (
		i int32
		j int32
	)
	_, _ = i, j
	i = 0
	j = 0
	OldIfaceDate = NewIfaceDate
	for i = 0; i < NumIfaceNodes; i++ {
		for j = 0; j < (NumIfacePolluts + 1); j++ {
			OldIfaceValues[i][j] = NewIfaceValues[i][j]
		}
	}
}

// iface.c:633
func isOutletNode(i int32) int32 {
	if RouteModel == DW {
		return R.B2I((Node[i].type_ == OUTFALL))
	} else {
		return R.B2I((Node[i].degree == 0))
	}
	return 0
}
