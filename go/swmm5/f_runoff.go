// Generated from EPA SWMM 5.2.4 (runoff.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

var IsRaining int32
var HasRunoff int32
var HasSnow int32
var runoff_Nsteps int32
var MaxSteps int32
var MaxStepsPos int64
var HasWetLids int32
var OutflowLoad []float64

// runoff.c:82
func runoff_open() int32 {
	IsRaining = 0
	HasRunoff = 0
	HasSnow = 0
	runoff_Nsteps = 0
	if !(odesolve_open(4) != 0) {
		report_writeErrorMsg(ERR_ODE_SOLVER, L0)
	}
	OutflowLoad = nil
	if Nobjects[POLLUT] > 0 {
		OutflowLoad = make([]float64, int(uint64(Nobjects[POLLUT])))
		if !(OutflowLoad != nil) {
			report_writeErrorMsg(ERR_MEMORY, L0)
		}
	}
	switch Frunoff.mode {
	case USE_FILE:
		if (R.Set(&Frunoff.file, R.Fopen(R.CP(Frunoff.name[:]), L165))) == nil {
			report_writeErrorMsg(ERR_RUNOFF_FILE_OPEN, R.CP(Frunoff.name[:]))
		} else {
			runoff_initFile()
		}
		break
	case SAVE_FILE:
		if (R.Set(&Frunoff.file, R.Fopen(R.CP(Frunoff.name[:]), L166))) == nil {
			report_writeErrorMsg(ERR_RUNOFF_FILE_OPEN, R.CP(Frunoff.name[:]))
		} else {
			runoff_initFile()
		}
		break
	}
	return ErrorCode
}

// runoff.c:124
func runoff_close() {
	odesolve_close()
	{
		if OutflowLoad != nil {
			_ = OutflowLoad
			OutflowLoad = nil
		}
	}
	if Frunoff.file != nil {
		if Frunoff.mode == SAVE_FILE {
			R.Fseek(Frunoff.file, MaxStepsPos, 0)
			R.Fwrite1(runoff_Nsteps, 4, Frunoff.file)
		}
		R.Fclose(Frunoff.file)
	}
	if Fclimate.file != nil {
		R.Fclose(Fclimate.file)
	}
}

// runoff.c:155
func runoff_execute() {
	var (
		j             int32
		day           int32
		runoffStep    float64
		oldRunoffStep float64
		runoff        float64
		currentDate   float64
		canSweep      int32
	)
	_, _, _, _, _, _, _ = j, day, runoffStep, oldRunoffStep, runoff, currentDate, canSweep
	j = 0
	day = 0
	runoffStep = 0
	oldRunoffStep = 0
	runoff = 0
	currentDate = 0
	canSweep = 0
	if ErrorCode != 0 {
		return
	}
	oldRunoffStep = float64((float64(NewRunoffTime - OldRunoffTime)) / float64(1000.0))
	currentDate = getDateTime(NewRunoffTime)
	climate_setState(currentDate)
	if Nobjects[SUBCATCH] == 0 {
		OldRunoffTime = NewRunoffTime
		NewRunoffTime = float64(NewRunoffTime + (float64(float64((DryStep)) * float64(1000.0))))
		if (NewRunoffTime) <= (TotalDuration) {
			NewRunoffTime = (NewRunoffTime)
		} else {
			NewRunoffTime = (TotalDuration)
		}
		return
	}
	IsRaining = 0
	for j = 0; j < Nobjects[GAGE]; j++ {
		gage_setState(j, currentDate)
		if Gage[j].rainfall > float64(0.0) {
			IsRaining = 1
		}
	}
	if Frunoff.mode == USE_FILE {
		runoff_readFromFile()
		return
	}
	day = datetime_dayOfYear(currentDate)
	canSweep = 0
	if SweepStart <= SweepEnd {
		if (day >= SweepStart) && (day <= SweepEnd) {
			canSweep = 1
		}
	} else if (day <= SweepEnd) || (day >= SweepStart) {
		canSweep = 1
	}
	runoffStep = runoff_getTimeStep(currentDate)
	if runoffStep <= float64(0.0) {
		ErrorCode = ERR_TIMESTEP
		return
	}
	OldRunoffTime = NewRunoffTime
	NewRunoffTime = float64(NewRunoffTime + (float64(float64(1000) * runoffStep)))
	if NewRunoffTime > TotalDuration {
		runoffStep = float64((float64(TotalDuration - OldRunoffTime)) / float64(1000.0))
		NewRunoffTime = TotalDuration
	}
	for j = 0; j < Nobjects[GAGE]; j++ {
		gage_updatePastRain(j, R.D2I(runoffStep))
	}
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		subcatch_setOldState(j)
	}
	if oldRunoffStep > float64(0.0) {
		runoff_getOutfallRunon(oldRunoffStep)
	}
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		if Subcatch[j].area == float64(0.0) {
			continue
		}
		subcatch_getRunon(j)
		if !(IgnoreSnowmelt != 0) {
			snow_plowSnow(j, runoffStep)
		}
	}
	HasSnow = 0
	HasRunoff = 0
	HasWetLids = 0
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		if Subcatch[j].area == float64(0.0) {
			continue
		}
		runoff = subcatch_getRunoff(j, runoffStep)
		if runoff > float64(0.0) {
			HasRunoff = 1
		}
		if Subcatch[j].newSnowDepth > float64(0.0) {
			HasSnow = 1
		}
		if IgnoreQuality != 0 {
			continue
		}
		if runoff < float64(2.31481e-08) {
			surfqual_getBuildup(j, runoffStep)
		}
		if (canSweep != 0) && (Subcatch[j].rainfall <= float64(2.31481e-08)) {
			surfqual_sweepBuildup(j, currentDate)
		}
		surfqual_getWashoff(j, runoff, runoffStep)
	}
	stats_updateMaxRunoff()
	runoff_Nsteps++
	if Frunoff.mode == SAVE_FILE {
		runoff_saveToFile(float32(float64(runoffStep)))
	}
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		Subcatch[j].runon = float64(0.0)
	}
}

// runoff.c:299
func runoff_getTimeStep(currentDate float64) float64 {
	var (
		j        int32
		timeStep int64
		maxStep  int64
	)
	_, _, _ = j, timeStep, maxStep
	j = 0
	timeStep = 0
	maxStep = int64(DryStep)
	timeStep = datetime_timeDiff(climate_getNextEvapDate(), currentDate)
	if (float64(timeStep) > float64(0.0)) && (timeStep < maxStep) {
		maxStep = timeStep
	}
	for j = 0; j < Nobjects[GAGE]; j++ {
		timeStep = datetime_timeDiff(gage_getNextRainDate(j, currentDate), currentDate)
		if (timeStep > 0) && (timeStep < maxStep) {
			maxStep = timeStep
		}
	}
	if (((IsRaining != 0) || (HasSnow != 0)) || (HasRunoff != 0)) || (HasWetLids != 0) {
		timeStep = int64(WetStep)
	} else {
		timeStep = int64(DryStep)
	}
	if timeStep > maxStep {
		timeStep = maxStep
	}
	return float64(timeStep)
}

// runoff.c:335
func runoff_initFile() {
	var (
		nSubcatch int32
		nPollut   int32
		flowUnits int32
		fileStamp R.CPtr
		fStamp    R.CPtr
	)
	_, _, _, _, _ = nSubcatch, nPollut, flowUnits, fileStamp, fStamp
	nSubcatch = 0
	nPollut = 0
	flowUnits = 0
	fileStamp = R.BufFrom(L533, 13)
	fStamp = R.BufFrom(L533, 13)
	MaxSteps = 0
	if Frunoff.mode == SAVE_FILE {
		nSubcatch = Nobjects[SUBCATCH]
		nPollut = Nobjects[POLLUT]
		flowUnits = FlowUnits
		R.FwriteC(fileStamp, R.Strlen(fileStamp), Frunoff.file)
		R.Fwrite1(nSubcatch, 4, Frunoff.file)
		R.Fwrite1(nPollut, 4, Frunoff.file)
		R.Fwrite1(flowUnits, 4, Frunoff.file)
		MaxStepsPos = R.Ftell(Frunoff.file)
		R.Fwrite1(MaxSteps, 4, Frunoff.file)
	}
	if Frunoff.mode == USE_FILE {
		R.FreadC(fStamp, R.Strlen(fileStamp), Frunoff.file)
		if R.Strcmp(fStamp, fileStamp) != 0 {
			report_writeErrorMsg(ERR_RUNOFF_FILE_FORMAT, L0)
			return
		}
		nSubcatch = (-1)
		nPollut = (-1)
		flowUnits = (-1)
		R.Fread1(&nSubcatch, 4, Frunoff.file)
		R.Fread1(&nPollut, 4, Frunoff.file)
		R.Fread1(&flowUnits, 4, Frunoff.file)
		R.Fread1(&MaxSteps, 4, Frunoff.file)
		if (((nSubcatch != Nobjects[SUBCATCH]) || (nPollut != Nobjects[POLLUT])) || (flowUnits != FlowUnits)) || (MaxSteps <= 0) {
			report_writeErrorMsg(ERR_RUNOFF_FILE_FORMAT, L0)
		}
	}
}

// runoff.c:391
func runoff_saveToFile(tStep float32) {
	var (
		j   int32
		n_2 int32
	)
	_, _ = j, n_2
	j = 0
	n_2 = ((9 + Nobjects[POLLUT]) - 1)
	R.Fwrite1(tStep, 4, Frunoff.file)
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		subcatch_getResults(j, float64(1.0), SubcatchResults)
		R.Fwrite(SubcatchResults, uint64(n_2), Frunoff.file)
	}
}

// runoff.c:412
func runoff_readFromFile() {
	var (
		i        int32
		j        int32
		nResults int32
		kount    int32
		tStep    float32
		gw       *TGroundwater
	)
	_, _, _, _, _, _ = i, j, nResults, kount, tStep, gw
	i = 0
	j = 0
	nResults = 0
	kount = 0
	tStep = 0
	gw = nil
	if runoff_Nsteps > MaxSteps {
		report_writeErrorMsg(ERR_RUNOFF_FILE_END, L0)
		return
	}
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		subcatch_setOldState(j)
	}
	kount = 0
	kount += (int32(R.Fread1(&tStep, 4, Frunoff.file)))
	nResults = ((9 + Nobjects[POLLUT]) - 1)
	for j = 0; j < Nobjects[SUBCATCH]; j++ {
		kount += (int32(R.Fread(SubcatchResults, uint64(nResults), Frunoff.file)))
		Subcatch[j].newSnowDepth = float64(float64(SubcatchResults[SUBCATCH_SNOWDEPTH]) / UCF(RAINDEPTH))
		Subcatch[j].evapLoss = float64(float64(SubcatchResults[SUBCATCH_EVAP]) / UCF(RAINFALL))
		Subcatch[j].infilLoss = float64(float64(SubcatchResults[SUBCATCH_INFIL]) / UCF(RAINFALL))
		Subcatch[j].newRunoff = float64(float64(SubcatchResults[SUBCATCH_RUNOFF]) / UCF(FLOW))
		gw = Subcatch[j].groundwater
		if gw != nil {
			gw.newFlow = float64(float64(SubcatchResults[SUBCATCH_GW_FLOW]) / UCF(FLOW))
			gw.lowerDepth = float64(Aquifer[gw.aquifer].bottomElev - (float64(float64(SubcatchResults[SUBCATCH_GW_ELEV]) / UCF(LENGTH))))
			gw.theta = float64(SubcatchResults[SUBCATCH_SOIL_MOIST])
		}
		for i = 0; i < Nobjects[POLLUT]; i++ {
			Subcatch[j].newQual[i] = float64(SubcatchResults[(SUBCATCH_WASHOFF + i)])
		}
	}
	if kount < (1 + (Nobjects[SUBCATCH] * nResults)) {
		report_writeErrorMsg(ERR_RUNOFF_FILE_READ, L0)
		return
	}
	OldRunoffTime = NewRunoffTime
	NewRunoffTime = float64(OldRunoffTime + float64(float64((tStep))*float64(1000.0)))
	if (NewRunoffTime) <= (TotalDuration) {
		NewRunoffTime = (NewRunoffTime)
	} else {
		NewRunoffTime = (TotalDuration)
	}
	runoff_Nsteps++
}

// runoff.c:491
func runoff_getOutfallRunon(tStep float64) {
	var (
		i int32
		k int32
		p int32
		w float64
	)
	_, _, _, _ = i, k, p, w
	i = 0
	k = 0
	p = 0
	w = 0
	for i = 0; i < Nnodes[OUTFALL]; i++ {
		k = Outfall[i].routeTo
		if k < 0 {
			continue
		}
		if Subcatch[k].area == float64(0.0) {
			continue
		}
		subcatch_addRunonFlow(k, float64(Outfall[i].vRouted/tStep))
		massbal_updateRunoffTotals(RUNOFF_RUNON, Outfall[i].vRouted)
		Outfall[i].vRouted = float64(0.0)
		for p = 0; p < Nobjects[POLLUT]; p++ {
			w = float64(Outfall[i].wRouted[p] * float64(28.317))
			massbal_updateLoadingTotals(DEPOSITION_LOAD, p, float64(w*Pollut[p].mcf))
			Subcatch[k].newQual[p] = float64(Subcatch[k].newQual[p] + (float64(w / tStep)))
			Outfall[i].wRouted[p] = float64(0.0)
		}
	}
}
