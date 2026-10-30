// Generated from EPA SWMM 5.2.4 (project.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var Htable [18][]*HTentry
var MemPoolAllocated int32

// project.c:70
func omp_get_max_threads() int32 {
	return 1
}

// project.c:113
func project_open(f1 R.CPtr, f2 R.CPtr, f3 R.CPtr) {
	initPointers()
	setDefaults()
	openFiles(f1, f2, f3)
}

// project.c:129
func project_readInput() {
	createHashTables()
	input_countObjects()
	createObjects()
	input_readData()
	if ErrorCode != 0 {
		return
	}
	StartDateTime = float64(StartDate + StartTime)
	EndDateTime = float64(EndDate + EndTime)
	ReportStart = float64(ReportStartDate + ReportStartTime)
	if (ReportStart) >= (StartDateTime) {
		ReportStart = (ReportStart)
	} else {
		ReportStart = (StartDateTime)
	}
	if EndDateTime <= StartDateTime {
		report_writeErrorMsg(ERR_START_DATE, L0)
	} else if EndDateTime <= ReportStart {
		report_writeErrorMsg(ERR_REPORT_DATE, L0)
	} else {
		TotalDuration = M.Floor(float64((float64(EndDateTime - StartDateTime)) * float64(86400)))
		if float64(ReportStep) > TotalDuration {
			ReportStep = R.D2I((TotalDuration))
		}
		if float64(ReportStep) < RouteStep {
			report_writeErrorMsg(ERR_REPORT_STEP, L0)
		}
		TotalDuration = float64(TotalDuration * (float64(1000.0)))
	}
}

// project.c:186
func project_validate() {
	var (
		i   int32
		j   int32
		err int32
	)
	_, _, _ = i, j, err
	i = 0
	j = 0
	err = 0
	for i = 0; i < Nobjects[CURVE]; i++ {
		err = table_validate(&Curve[i])
		if err != 0 {
			report_writeErrorMsg(ERR_CURVE_SEQUENCE, Curve[i].ID)
		}
	}
	for i = 0; i < Nobjects[TSERIES]; i++ {
		err = table_validate(&Tseries[i])
		if err != 0 {
			report_writeTseriesErrorMsg(err, &Tseries[i])
		}
	}
	climate_validate()
	lid_validate()
	if Nobjects[SNOWMELT] == 0 {
		IgnoreSnowmelt = 1
	}
	if Nobjects[AQUIFER] == 0 {
		IgnoreGwater = 1
	}
	for i = 0; i < Nobjects[AQUIFER]; i++ {
		gwater_validateAquifer(i)
	}
	for i = 0; i < Nobjects[SUBCATCH]; i++ {
		subcatch_validate(i)
	}
	for i = 0; i < Nobjects[GAGE]; i++ {
		gage_validate(i)
	}
	for i = 0; i < Nobjects[SNOWMELT]; i++ {
		snow_validateSnowmelt(i)
	}
	j = 0
	for i = 0; i < Nobjects[CURVE]; i++ {
		if Curve[i].curveType == SHAPE_CURVE {
			Curve[i].refersTo = j
			Shape[j].curve = i
			if !(shape_validate(&Shape[j], &Curve[i]) != 0) {
				report_writeErrorMsg(ERR_CURVE_SEQUENCE, Curve[i].ID)
			}
			j++
		}
	}
	for i = 0; i < Nobjects[NODE]; i++ {
		Node[i].oldDepth = Node[i].fullDepth
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		link_validate(i)
	}
	for i = 0; i < Nobjects[NODE]; i++ {
		node_validate(i)
	}
	if DryStep < WetStep {
		report_writeWarningMsg(L319, L0)
		DryStep = WetStep
	}
	if RouteStep > float64(WetStep) {
		report_writeWarningMsg(L320, L0)
		RouteStep = float64(WetStep)
	}
	if RptFlags.subcatchments == ALL {
		for i = 0; i < Nobjects[SUBCATCH]; i++ {
			Subcatch[i].rptFlag = 1
		}
	}
	if RptFlags.nodes == ALL {
		for i = 0; i < Nobjects[NODE]; i++ {
			Node[i].rptFlag = 1
		}
	}
	if RptFlags.links == ALL {
		for i = 0; i < Nobjects[LINK]; i++ {
			Link[i].rptFlag = 1
		}
	}
	if RouteModel == DW {
		dynwave_validate()
	}
	inlet_validate()
	if NumThreads == 0 {
		NumThreads = omp_get_max_threads()
	} else {
		if (NumThreads) <= (omp_get_max_threads()) {
			NumThreads = (NumThreads)
		} else {
			NumThreads = (omp_get_max_threads())
		}
	}
	if Nobjects[LINK] < (4 * NumThreads) {
		NumThreads = 1
	}
}

// project.c:274
func project_close() {
	deleteObjects()
	deleteHashTables()
}

// project.c:287
func project_init() int32 {
	var (
		j int32
		k int32
	)
	_, _ = j, k
	j = 0
	k = 0
	climate_initState()
	lid_initState()
	for j = 0; j < Nobjects[TSERIES]; j++ {
		table_tseriesInit(&Tseries[j])
	}
	for j = 0; j < Nobjects[GAGE]; j++ {
		gage_initState(j)
	}
	k = 1
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		subcatch_initState(j)
		if Subcatch[j].rptFlag > 0 {
			Subcatch[j].rptFlag = k
			k++
		}
	}
	k = 1
	for j = 0; j < Nobjects[NODE]; j++ {
		node_initState(j)
		if Node[j].rptFlag > 0 {
			Node[j].rptFlag = k
			k++
		}
	}
	k = 1
	for j = 0; j < Nobjects[LINK]; j++ {
		link_initState(j)
		if Link[j].rptFlag > 0 {
			Link[j].rptFlag = k
			k++
		}
	}
	return ErrorCode
}

// project.c:334
func project_addObject(type_2 int32, id R.CPtr, n_2 int32) int32 {
	var (
		result int32
		len_2  int64
		newID  R.CPtr
	)
	_, _, _ = result, len_2, newID
	result = 0
	len_2 = 0
	newID = (R.CPtr{})
	if project_findObject(type_2, id) >= 0 {
		return 0
	}
	len_2 = int64(R.Strlen(id))
	newID = Alloc(int64((uint64((len_2 + 1)) * 1)))
	sstrncpy(newID, id, uint64(len_2))
	result = HTinsert(Htable[type_2], newID, n_2)
	if result == 0 {
		result = (-1)
	}
	return result
}

// project.c:364
func project_findObject(type_2 int32, id R.CPtr) int32 {
	return HTfind(Htable[type_2], id)
}

// project.c:377
func project_findID(type_2 int32, id R.CPtr) R.CPtr {
	return HTfindKey(Htable[type_2], id)
}

// project.c:445
func project_readOption(s1 R.CPtr, s2 R.CPtr) int32 {
	var (
		k       int32
		m       int32
		h       int32
		s       int32
		tStep   float64
		strDate R.CPtr
		aTime   float64
		aDate   float64
	)
	_, _, _, _, _, _, _, _ = k, m, h, s, tStep, strDate, aTime, aDate
	k = 0
	m = 0
	h = 0
	s = 0
	tStep = 0
	strDate = R.Buf(25)
	aTime = 0
	aDate = 0
	k = findmatch(s1, OptionWords[:])
	if k < 0 {
		return error_setInpError(ERR_KEYWORD, s1)
	}
	switch k {
	case FLOW_UNITS:
		m = findmatch(s2, FlowUnitWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, s2)
		}
		FlowUnits = m
		if FlowUnits <= MGD {
			UnitSystem = US
		} else {
			UnitSystem = SI
		}
		break
	case INFIL_MODEL:
		m = findmatch(s2, InfilModelWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, s2)
		}
		InfilModel = m
		break
	case ROUTE_MODEL:
		m = findmatch(s2, RouteModelWords[:])
		if m < 0 {
			m = findmatch(s2, OldRouteModelWords[:])
		}
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, s2)
		}
		if m == NO_ROUTING {
			IgnoreRouting = 1
		} else {
			RouteModel = m
		}
		if RouteModel == EKW {
			RouteModel = KW
		}
		break
	case START_DATE:
		if !(datetime_strToDate(s2, &StartDate) != 0) {
			return error_setInpError(ERR_DATETIME, s2)
		}
		break
	case START_TIME:
		if !(datetime_strToTime(s2, &StartTime) != 0) {
			return error_setInpError(ERR_DATETIME, s2)
		}
		break
	case END_DATE:
		if !(datetime_strToDate(s2, &EndDate) != 0) {
			return error_setInpError(ERR_DATETIME, s2)
		}
		break
	case END_TIME:
		if !(datetime_strToTime(s2, &EndTime) != 0) {
			return error_setInpError(ERR_DATETIME, s2)
		}
		break
	case REPORT_START_DATE:
		if !(datetime_strToDate(s2, &ReportStartDate) != 0) {
			return error_setInpError(ERR_DATETIME, s2)
		}
		break
	case REPORT_START_TIME:
		if !(datetime_strToTime(s2, &ReportStartTime) != 0) {
			return error_setInpError(ERR_DATETIME, s2)
		}
		break
	case SWEEP_START, SWEEP_END:
		sstrncpy(strDate, s2, 24)
		sstrcat(strDate, L35, 25)
		if !(datetime_strToDate(strDate, &aDate) != 0) {
			return error_setInpError(ERR_DATETIME, s2)
		}
		m = datetime_dayOfYear(aDate)
		if k == SWEEP_START {
			SweepStart = m
		} else {
			SweepEnd = m
		}
		break
	case START_DRY_DAYS:
		StartDryDays = R.Atof(s2)
		if StartDryDays < float64(0.0) {
			return error_setInpError(ERR_NUMBER, s2)
		}
		break
	case WET_STEP, DRY_STEP, REPORT_STEP, RULE_STEP:
		if !(datetime_strToTime(s2, &aTime) != 0) {
			return error_setInpError(ERR_DATETIME, s2)
		}
		datetime_decodeTime(aTime, &h, &m, &s)
		h += (24 * R.D2I(aTime))
		s = ((s + (60 * m)) + (3600 * h))
		if k == RULE_STEP {
			if s < 0 {
				return error_setInpError(ERR_NUMBER, s2)
			}
		} else if s <= 0 {
			return error_setInpError(ERR_NUMBER, s2)
		}
		switch k {
		case WET_STEP:
			WetStep = s
			break
		case DRY_STEP:
			DryStep = s
			break
		case REPORT_STEP:
			ReportStep = s
			break
		case RULE_STEP:
			RuleStep = s
			break
		}
		break
	case INERT_DAMPING:
		m = findmatch(s2, InertDampingWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, s2)
		} else {
			InertDamping = m
		}
		break
	case ALLOW_PONDING, SLOPE_WEIGHTING, SKIP_STEADY_STATE, IGNORE_RAINFALL, IGNORE_SNOWMELT, IGNORE_GWATER, IGNORE_ROUTING, IGNORE_QUALITY, IGNORE_RDII:
		m = findmatch(s2, NoYesWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, s2)
		}
		switch k {
		case ALLOW_PONDING:
			AllowPonding = m
			break
		case SLOPE_WEIGHTING:
			SlopeWeighting = m
			break
		case SKIP_STEADY_STATE:
			SkipSteadyState = m
			break
		case IGNORE_RAINFALL:
			IgnoreRainfall = m
			break
		case IGNORE_SNOWMELT:
			IgnoreSnowmelt = m
			break
		case IGNORE_GWATER:
			IgnoreGwater = m
			break
		case IGNORE_ROUTING:
			IgnoreRouting = m
			break
		case IGNORE_QUALITY:
			IgnoreQuality = m
			break
		case IGNORE_RDII:
			IgnoreRDII = m
			break
		}
		break
	case NORMAL_FLOW_LTD:
		m = findmatch(s2, NormalFlowWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, s2)
		}
		NormalFlowLtd = m
		break
	case FORCE_MAIN_EQN:
		m = findmatch(s2, ForceMainEqnWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, s2)
		}
		ForceMainEqn = m
		break
	case LINK_OFFSETS:
		m = findmatch(s2, LinkOffsetWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, s2)
		}
		LinkOffsets = m
		break
	case COMPATIBILITY:
		if strcomp(s2, L321) != 0 {
			Compatibility = SWMM3
		} else if strcomp(s2, L322) != 0 {
			Compatibility = SWMM4
		} else if strcomp(s2, L323) != 0 {
			Compatibility = SWMM5
		} else {
			return error_setInpError(ERR_KEYWORD, s2)
		}
		break
	case ROUTE_STEP, LENGTHENING_STEP:
		if !(getDouble(s2, &tStep) != 0) {
			if !(datetime_strToTime(s2, &aTime) != 0) {
				return error_setInpError(ERR_NUMBER, s2)
			} else {
				datetime_decodeTime(aTime, &h, &m, &s)
				h += (24 * R.D2I(aTime))
				s = ((s + (60 * m)) + (3600 * h))
				tStep = float64(s)
			}
		}
		if k == ROUTE_STEP {
			if tStep <= float64(0.0) {
				return error_setInpError(ERR_NUMBER, s2)
			}
			RouteStep = tStep
		} else {
			if (float64(0.0)) >= (tStep) {
				LengtheningStep = (float64(0.0))
			} else {
				LengtheningStep = (tStep)
			}
		}
		break
	case MIN_ROUTE_STEP:
		if (!(getDouble(s2, &MinRouteStep) != 0)) || (MinRouteStep < float64(0.0)) {
			return error_setInpError(ERR_NUMBER, s2)
		}
		break
	case NUM_THREADS:
		m = R.Atoi(s2)
		if m < 0 {
			return error_setInpError(ERR_NUMBER, s2)
		}
		NumThreads = m
		break
	case VARIABLE_STEP:
		if !(getDouble(s2, &CourantFactor) != 0) {
			return error_setInpError(ERR_NUMBER, s2)
		}
		if (CourantFactor < float64(0.0)) || (CourantFactor > float64(2.0)) {
			return error_setInpError(ERR_NUMBER, s2)
		}
		break
	case MIN_SURFAREA:
		if !(getDouble(s2, &MinSurfArea) != 0) {
			return error_setInpError(ERR_NUMBER, s2)
		}
		if MinSurfArea < float64(0.0) {
			return error_setInpError(ERR_NUMBER, s2)
		}
		break
	case MIN_SLOPE:
		if !(getDouble(s2, &MinSlope) != 0) {
			return error_setInpError(ERR_NUMBER, s2)
		}
		if (MinSlope < float64(0.0)) || (MinSlope >= float64(100)) {
			return error_setInpError(ERR_NUMBER, s2)
		}
		MinSlope = float64(MinSlope / (float64(100.0)))
		break
	case MAX_TRIALS:
		m = R.Atoi(s2)
		if m < 0 {
			return error_setInpError(ERR_NUMBER, s2)
		}
		MaxTrials = m
		break
	case HEAD_TOL:
		if !(getDouble(s2, &HeadTol) != 0) {
			return error_setInpError(ERR_NUMBER, s2)
		}
		break
	case SYS_FLOW_TOL:
		if !(getDouble(s2, &SysFlowTol) != 0) {
			return error_setInpError(ERR_NUMBER, s2)
		}
		SysFlowTol = float64(SysFlowTol / (float64(100.0)))
		break
	case LAT_FLOW_TOL:
		if !(getDouble(s2, &LatFlowTol) != 0) {
			return error_setInpError(ERR_NUMBER, s2)
		}
		LatFlowTol = float64(LatFlowTol / (float64(100.0)))
		break
	case SURCHARGE_METHOD:
		m = findmatch(s2, SurchargeWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, s2)
		}
		SurchargeMethod = m
		break
	case TEMPDIR:
		sstrncpy(TempDir, s2, 259)
		break
	}
	return 0
}

// project.c:773
func initPointers() {
	Gage = nil
	Subcatch = nil
	Node = nil
	Outfall = nil
	Divider = nil
	Storage = nil
	Link = nil
	Conduit = nil
	Pump = nil
	Orifice = nil
	Weir = nil
	Outlet = nil
	Pollut = nil
	Landuse = nil
	Pattern = nil
	Curve = nil
	Tseries = nil
	Transect = nil
	Shape = nil
	Aquifer = nil
	UnitHyd = nil
	Snowmelt = nil
	Event = nil
	MemPoolAllocated = 0
}

// project.c:808
func setDefaults() {
	var (
		i int32
		j int32
	)
	_, _ = i, j
	i = 0
	j = 0
	for i = 0; i < 3; i++ {
		sstrncpy(Title[i], L0, 0)
	}
	sstrncpy(TempDir, L0, 0)
	Frain.mode = SCRATCH_FILE
	Fclimate.mode = NO_FILE
	Frunoff.mode = NO_FILE
	Frdii.mode = NO_FILE
	Fhotstart1.mode = NO_FILE
	Fhotstart2.mode = NO_FILE
	Finflows.mode = NO_FILE
	Foutflows.mode = NO_FILE
	Frain.file = (nil)
	Fclimate.file = (nil)
	Frunoff.file = (nil)
	Frdii.file = (nil)
	Fhotstart1.file = (nil)
	Fhotstart2.file = (nil)
	Finflows.file = (nil)
	Foutflows.file = (nil)
	Fout.file = (nil)
	Fout.mode = NO_FILE
	UnitSystem = US
	FlowUnits = CFS
	InfilModel = HORTON
	RouteModel = DW
	SurchargeMethod = EXTRAN
	CrownCutoff = float64(0.96)
	AllowPonding = 0
	InertDamping = PARTIAL_DAMPING
	NormalFlowLtd = BOTH
	ForceMainEqn = H_W
	LinkOffsets = DEPTH_OFFSET
	LengtheningStep = float64(0)
	CourantFactor = float64(0.75)
	MinSurfArea = float64(0.0)
	MinSlope = float64(0.0)
	SkipSteadyState = 0
	IgnoreRainfall = 0
	IgnoreRDII = 0
	IgnoreSnowmelt = 0
	IgnoreGwater = 0
	IgnoreRouting = 0
	IgnoreQuality = 0
	WetStep = 300
	DryStep = 3600
	RuleStep = 0
	RouteStep = float64(20)
	MinRouteStep = float64(0.5)
	ReportStep = 900
	StartDryDays = float64(0.0)
	MaxTrials = 0
	HeadTol = float64(0.0)
	SysFlowTol = float64(0.05)
	LatFlowTol = float64(0.05)
	NumThreads = 1
	NumEvents = 0
	SlopeWeighting = 1
	Compatibility = SWMM4
	StartDate = datetime_encodeDate(2004, 1, 1)
	StartTime = datetime_encodeTime(0, 0, 0)
	StartDateTime = float64(StartDate + StartTime)
	EndDate = StartDate
	EndTime = float64(0.0)
	ReportStartDate = float64((-693594))
	ReportStartTime = float64((-693594))
	SweepStart = 1
	SweepEnd = 365
	RptFlags.disabled = 0
	RptFlags.input = 0
	RptFlags.continuity = 1
	RptFlags.flowStats = 1
	RptFlags.controls = 0
	RptFlags.subcatchments = 0
	RptFlags.nodes = 0
	RptFlags.links = 0
	RptFlags.averages = 0
	Temp.dataSource = NO_TEMP
	Temp.tSeries = (-1)
	Temp.ta = float64(70.0)
	Temp.elev = float64(0.0)
	Temp.anglat = float64(40.0)
	Temp.dtlong = float64(0.0)
	Temp.tmax = (-(float64(10000000000.0)))
	Wind.type_ = MONTHLY_WIND
	for i = 0; i < 12; i++ {
		Wind.aws[i] = float64(0.0)
	}
	Snow.snotmp = float64(34.0)
	Snow.tipm = float64(0.5)
	Snow.rnm = float64(0.6)
	for i = 0; i < 2; i++ {
		for j = 0; j < 10; j++ {
			Snow.adc[i][j] = float64(1.0)
		}
	}
	Evap.type_ = CONSTANT_EVAP
	for i = 0; i < 12; i++ {
		Evap.monthlyEvap[i] = float64(0.0)
		Evap.panCoeff[i] = float64(1.0)
	}
	Evap.recoveryPattern = (-1)
	Evap.recoveryFactor = float64(1.0)
	Evap.tSeries = (-1)
	Evap.dryOnly = 0
	for i = 0; i < 12; i++ {
		Adjust.temp[i] = float64(0.0)
		Adjust.evap[i] = float64(0.0)
		Adjust.rain[i] = float64(1.0)
		Adjust.hydcon[i] = float64(1.0)
	}
	Adjust.rainFactor = float64(1.0)
	Adjust.hydconFactor = float64(1.0)
}

// project.c:954
func openFiles(f1 R.CPtr, f2 R.CPtr, f3 R.CPtr) {
	Finp.file = (nil)
	Frpt.file = (nil)
	Fout.file = (nil)
	sstrncpy(R.CP(Finp.name[:]), f1, 259)
	sstrncpy(R.CP(Frpt.name[:]), f2, 259)
	sstrncpy(R.CP(Fout.name[:]), f3, 259)
	if ((strcomp(f1, f2) != 0) || (strcomp(f1, f3) != 0)) || (strcomp(f2, f3) != 0) {
		writecon(L324)
		ErrorCode = ERR_FILE_NAME
		return
	}
	if (R.Set(&Finp.file, R.Fopen(f1, L13))) == nil {
		writecon(L325)
		writecon(f1)
		ErrorCode = ERR_INP_FILE
		return
	}
	if (R.Set(&Frpt.file, R.Fopen(f2, L171))) == nil {
		writecon(L326)
		ErrorCode = ERR_RPT_FILE
		return
	}
}

// project.c:999
func createObjects() {
	var (
		j int32
		k int32
	)
	_, _ = j, k
	j = 0
	k = 0
	if ErrorCode != 0 {
		return
	}
	Gage = make([]TGage, int(uint64(Nobjects[GAGE])))
	Subcatch = make([]TSubcatch, int(uint64(Nobjects[SUBCATCH])))
	Node = make([]TNode, int(uint64(Nobjects[NODE])))
	Outfall = make([]TOutfall, int(uint64(Nnodes[OUTFALL])))
	Divider = make([]TDivider, int(uint64(Nnodes[DIVIDER])))
	Storage = make([]TStorage, int(uint64(Nnodes[STORAGE])))
	Link = make([]TLink, int(uint64(Nobjects[LINK])))
	Conduit = make([]TConduit, int(uint64(Nlinks[CONDUIT])))
	Pump = make([]TPump, int(uint64(Nlinks[PUMP])))
	Orifice = make([]TOrifice, int(uint64(Nlinks[ORIFICE])))
	Weir = make([]TWeir, int(uint64(Nlinks[WEIR])))
	Outlet = make([]TOutlet, int(uint64(Nlinks[OUTLET])))
	Pollut = make([]TPollut, int(uint64(Nobjects[POLLUT])))
	Landuse = make([]TLanduse, int(uint64(Nobjects[LANDUSE])))
	Pattern = make([]TPattern, int(uint64(Nobjects[TIMEPATTERN])))
	Curve = make([]TTable, int(uint64(Nobjects[CURVE])))
	Tseries = make([]TTable, int(uint64(Nobjects[TSERIES])))
	Aquifer = make([]TAquifer, int(uint64(Nobjects[AQUIFER])))
	UnitHyd = make([]TUnitHyd, int(uint64(Nobjects[UNITHYD])))
	Snowmelt = make([]TSnowmelt, int(uint64(Nobjects[SNOWMELT])))
	Shape = make([]TShape, int(uint64(Nobjects[SHAPE])))
	Event = make([]TEvent, int((uint64(NumEvents) + 1)))
	Event[NumEvents].start = float64(10000000000.0)
	Event[NumEvents].end = float64(float64(10000000000.0) + float64(1.0))
	lid_create(Nobjects[LID], Nobjects[SUBCATCH])
	ErrorCode = controls_create(Nobjects[CONTROL])
	if ErrorCode != 0 {
		return
	}
	ErrorCode = transect_create(Nobjects[TRANSECT])
	if ErrorCode != 0 {
		return
	}
	ErrorCode = street_create(Nobjects[STREET])
	if ErrorCode != 0 {
		return
	}
	ErrorCode = inlet_create(Nobjects[INLET])
	if ErrorCode != 0 {
		return
	}
	infil_create(Nobjects[SUBCATCH])
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		Subcatch[j].initBuildup = make([]float64, int(uint64(Nobjects[POLLUT])))
		Subcatch[j].oldQual = make([]float64, int(uint64(Nobjects[POLLUT])))
		Subcatch[j].newQual = make([]float64, int(uint64(Nobjects[POLLUT])))
		Subcatch[j].pondedQual = make([]float64, int(uint64(Nobjects[POLLUT])))
		Subcatch[j].totalLoad = make([]float64, int(uint64(Nobjects[POLLUT])))
	}
	for j = 0; j < Nobjects[NODE]; j++ {
		Node[j].oldQual = make([]float64, int(uint64(Nobjects[POLLUT])))
		Node[j].newQual = make([]float64, int(uint64(Nobjects[POLLUT])))
		Node[j].extInflow = nil
		Node[j].dwfInflow = nil
		Node[j].rdiiInflow = nil
		Node[j].treatment = nil
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		Link[j].inlet = nil
		Link[j].oldQual = make([]float64, int(uint64(Nobjects[POLLUT])))
		Link[j].newQual = make([]float64, int(uint64(Nobjects[POLLUT])))
		Link[j].totalLoad = make([]float64, int(uint64(Nobjects[POLLUT])))
	}
	for j = 0; j < Nobjects[LANDUSE]; j++ {
		Landuse[j].buildupFunc = make([]TBuildup, int(uint64(Nobjects[POLLUT])))
		Landuse[j].washoffFunc = make([]TWashoff, int(uint64(Nobjects[POLLUT])))
	}
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		Subcatch[j].landFactor = make([]TLandFactor, int(uint64(Nobjects[LANDUSE])))
		for k = 0; k < Nobjects[LANDUSE]; k++ {
			Subcatch[j].landFactor[k].buildup = make([]float64, int(uint64(Nobjects[POLLUT])))
		}
	}
	for j = 0; j < Nobjects[LANDUSE]; j++ {
		for k = 0; k < Nobjects[POLLUT]; k++ {
			Landuse[j].buildupFunc[k].funcType = NO_BUILDUP
			Landuse[j].buildupFunc[k].normalizer = PER_AREA
			Landuse[j].washoffFunc[k].funcType = NO_WASHOFF
		}
	}
	for j = 0; j < Nobjects[GAGE]; j++ {
		Gage[j].tSeries = (-1)
		sstrncpy(R.CP(Gage[j].fname[:]), L0, 0)
	}
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		Subcatch[j].outSubcatch = (-1)
		Subcatch[j].outNode = (-1)
		Subcatch[j].infil = (-1)
		Subcatch[j].groundwater = nil
		Subcatch[j].gwLatFlowExpr = nil
		Subcatch[j].gwDeepFlowExpr = nil
		Subcatch[j].snowpack = nil
		Subcatch[j].lidArea = float64(0.0)
		for k = 0; k < Nobjects[POLLUT]; k++ {
			Subcatch[j].initBuildup[k] = float64(0.0)
		}
	}
	for j = 0; j < Nobjects[UNITHYD]; j++ {
		rdii_initUnitHyd(j)
	}
	for j = 0; j < Nobjects[SNOWMELT]; j++ {
		snow_initSnowmelt(j)
	}
	for j = 0; j < Nnodes[STORAGE]; j++ {
		Storage[j].exfil = nil
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		Link[j].xsect.type_ = (-1)
		Link[j].cLossInlet = float64(0.0)
		Link[j].cLossOutlet = float64(0.0)
		Link[j].cLossAvg = float64(0.0)
		Link[j].hasFlapGate = 0
	}
	for j = 0; j < Nlinks[PUMP]; j++ {
		Pump[j].pumpCurve = (-1)
	}
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		Subcatch[j].rptFlag = 0
	}
	for j = 0; j < Nobjects[NODE]; j++ {
		Node[j].rptFlag = 0
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		Link[j].rptFlag = 0
	}
	for j = 0; j < Nobjects[CURVE]; j++ {
		table_init(&Curve[j])
	}
	for j = 0; j < Nobjects[TSERIES]; j++ {
		table_init(&Tseries[j])
	}
	for j = 0; j < Nobjects[TIMEPATTERN]; j++ {
		inflow_initDwfPattern(j)
	}
}

// project.c:1176
func deleteObjects() {
	var (
		j int32
		k int32
	)
	_, _ = j, k
	j = 0
	k = 0
	if Subcatch != nil {
		for j = 0; j < Nobjects[SUBCATCH]; j++ {
			for k = 0; k < Nobjects[LANDUSE]; k++ {
				{
					if Subcatch[j].landFactor[k].buildup != nil {
						_ = Subcatch[j].landFactor[k].buildup
						Subcatch[j].landFactor[k].buildup = nil
					}
				}
			}
			{
				if Subcatch[j].landFactor != nil {
					_ = Subcatch[j].landFactor
					Subcatch[j].landFactor = nil
				}
			}
			{
				if Subcatch[j].groundwater != nil {
					_ = Subcatch[j].groundwater
					Subcatch[j].groundwater = nil
				}
			}
			gwater_deleteFlowExpression(j)
			{
				if Subcatch[j].snowpack != nil {
					_ = Subcatch[j].snowpack
					Subcatch[j].snowpack = nil
				}
			}
		}
	}
	if Landuse != nil {
		for j = 0; j < Nobjects[LANDUSE]; j++ {
			{
				if Landuse[j].buildupFunc != nil {
					_ = Landuse[j].buildupFunc
					Landuse[j].buildupFunc = nil
				}
			}
			{
				if Landuse[j].washoffFunc != nil {
					_ = Landuse[j].washoffFunc
					Landuse[j].washoffFunc = nil
				}
			}
		}
	}
	if Subcatch != nil {
		for j = 0; j < Nobjects[SUBCATCH]; j++ {
			{
				if Subcatch[j].initBuildup != nil {
					_ = Subcatch[j].initBuildup
					Subcatch[j].initBuildup = nil
				}
			}
			{
				if Subcatch[j].oldQual != nil {
					_ = Subcatch[j].oldQual
					Subcatch[j].oldQual = nil
				}
			}
			{
				if Subcatch[j].newQual != nil {
					_ = Subcatch[j].newQual
					Subcatch[j].newQual = nil
				}
			}
			{
				if Subcatch[j].pondedQual != nil {
					_ = Subcatch[j].pondedQual
					Subcatch[j].pondedQual = nil
				}
			}
			{
				if Subcatch[j].totalLoad != nil {
					_ = Subcatch[j].totalLoad
					Subcatch[j].totalLoad = nil
				}
			}
		}
	}
	if Node != nil {
		for j = 0; j < Nobjects[NODE]; j++ {
			{
				if Node[j].oldQual != nil {
					_ = Node[j].oldQual
					Node[j].oldQual = nil
				}
			}
			{
				if Node[j].newQual != nil {
					_ = Node[j].newQual
					Node[j].newQual = nil
				}
			}
		}
	}
	if Link != nil {
		for j = 0; j < Nobjects[LINK]; j++ {
			{
				if Link[j].oldQual != nil {
					_ = Link[j].oldQual
					Link[j].oldQual = nil
				}
			}
			{
				if Link[j].newQual != nil {
					_ = Link[j].newQual
					Link[j].newQual = nil
				}
			}
			{
				if Link[j].totalLoad != nil {
					_ = Link[j].totalLoad
					Link[j].totalLoad = nil
				}
			}
		}
	}
	infil_delete()
	if Node != nil {
		for j = 0; j < Nnodes[STORAGE]; j++ {
			if Storage[j].exfil != nil {
				{
					if Storage[j].exfil.btmExfil != nil {
						_ = Storage[j].exfil.btmExfil
						Storage[j].exfil.btmExfil = nil
					}
				}
				{
					if Storage[j].exfil.bankExfil != nil {
						_ = Storage[j].exfil.bankExfil
						Storage[j].exfil.bankExfil = nil
					}
				}
				{
					if Storage[j].exfil != nil {
						_ = Storage[j].exfil
						Storage[j].exfil = nil
					}
				}
			}
		}
	}
	if Node != nil {
		for j = 0; j < Nnodes[OUTFALL]; j++ {
			if Outfall[j].wRouted != nil {
				_ = Outfall[j].wRouted
				Outfall[j].wRouted = nil
			}
		}
	}
	if Node != nil {
		for j = 0; j < Nobjects[NODE]; j++ {
			inflow_deleteExtInflows(j)
			inflow_deleteDwfInflows(j)
			rdii_deleteRdiiInflow(j)
			treatmnt_delete(j)
		}
	}
	if Tseries != nil {
		for j = 0; j < Nobjects[TSERIES]; j++ {
			table_deleteEntries(&Tseries[j])
		}
	}
	if Curve != nil {
		for j = 0; j < Nobjects[CURVE]; j++ {
			table_deleteEntries(&Curve[j])
		}
	}
	transect_delete()
	street_delete()
	inlet_delete()
	controls_delete()
	lid_delete()
	{
		if Gage != nil {
			_ = Gage
			Gage = nil
		}
	}
	{
		if Subcatch != nil {
			_ = Subcatch
			Subcatch = nil
		}
	}
	{
		if Node != nil {
			_ = Node
			Node = nil
		}
	}
	{
		if Outfall != nil {
			_ = Outfall
			Outfall = nil
		}
	}
	{
		if Divider != nil {
			_ = Divider
			Divider = nil
		}
	}
	{
		if Storage != nil {
			_ = Storage
			Storage = nil
		}
	}
	{
		if Link != nil {
			_ = Link
			Link = nil
		}
	}
	{
		if Conduit != nil {
			_ = Conduit
			Conduit = nil
		}
	}
	{
		if Pump != nil {
			_ = Pump
			Pump = nil
		}
	}
	{
		if Orifice != nil {
			_ = Orifice
			Orifice = nil
		}
	}
	{
		if Weir != nil {
			_ = Weir
			Weir = nil
		}
	}
	{
		if Outlet != nil {
			_ = Outlet
			Outlet = nil
		}
	}
	{
		if Pollut != nil {
			_ = Pollut
			Pollut = nil
		}
	}
	{
		if Landuse != nil {
			_ = Landuse
			Landuse = nil
		}
	}
	{
		if Pattern != nil {
			_ = Pattern
			Pattern = nil
		}
	}
	{
		if Curve != nil {
			_ = Curve
			Curve = nil
		}
	}
	{
		if Tseries != nil {
			_ = Tseries
			Tseries = nil
		}
	}
	{
		if Aquifer != nil {
			_ = Aquifer
			Aquifer = nil
		}
	}
	{
		if UnitHyd != nil {
			_ = UnitHyd
			UnitHyd = nil
		}
	}
	{
		if Snowmelt != nil {
			_ = Snowmelt
			Snowmelt = nil
		}
	}
	{
		if Shape != nil {
			_ = Shape
			Shape = nil
		}
	}
	{
		if Event != nil {
			_ = Event
			Event = nil
		}
	}
}

// project.c:1304
func createHashTables() {
	var (
		j int32
	)
	_ = j
	j = 0
	MemPoolAllocated = 0
	for j = 0; j < MAX_OBJ_TYPES; j++ {
		Htable[j] = HTcreate()
		if Htable[j] == nil {
			report_writeErrorMsg(ERR_MEMORY, L0)
		}
	}
	if AllocInit() == nil {
		report_writeErrorMsg(ERR_MEMORY, L0)
	} else {
		MemPoolAllocated = 1
	}
}

// project.c:1325
func deleteHashTables() {
	var (
		j int32
	)
	_ = j
	j = 0
	for j = 0; j < MAX_OBJ_TYPES; j++ {
		if Htable[j] != nil {
			HTfree(Htable[j])
		}
	}
	if MemPoolAllocated != 0 {
		AllocFreePool()
	}
}
