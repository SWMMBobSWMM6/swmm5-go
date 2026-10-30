// Generated from EPA SWMM 5.2.4 (swmm5.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var Finp TFile
var Fout TFile
var Frpt TFile
var Fclimate TFile
var Frain TFile
var Frunoff TFile
var Frdii TFile
var Fhotstart1 TFile
var Fhotstart2 TFile
var Finflows TFile
var Foutflows TFile
var Nperiods int64
var TotalStepCount int64
var ReportStepCount int64
var NonConvergeCount int64
var Msg = R.Buf(1025)
var ErrorMsg = R.Buf(1025)
var Title = R.Bufs(3, 1025)
var TempDir = R.Buf(260)
var InpDir = R.Buf(260)
var RptFlags TRptFlags
var Nobjects [18]int32
var Nnodes [5]int32
var Nlinks [5]int32
var UnitSystem int32
var FlowUnits int32
var InfilModel int32
var RouteModel int32
var ForceMainEqn int32
var LinkOffsets int32
var SurchargeMethod int32
var AllowPonding int32
var InertDamping int32
var NormalFlowLtd int32
var SlopeWeighting int32
var Compatibility int32
var SkipSteadyState int32
var IgnoreRainfall int32
var IgnoreRDII int32
var IgnoreSnowmelt int32
var IgnoreGwater int32
var IgnoreRouting int32
var IgnoreQuality int32
var ErrorCode int32
var Warnings int32
var WetStep int32
var DryStep int32
var ReportStep int32
var RuleStep int32
var SweepStart int32
var SweepEnd int32
var MaxTrials int32
var NumThreads int32
var NumEvents int32
var RouteStep float64
var MinRouteStep float64
var LengtheningStep float64
var StartDryDays float64
var CourantFactor float64
var MinSurfArea float64
var MinSlope float64
var RunoffError float64
var GwaterError float64
var FlowError float64
var QualError float64
var HeadTol float64
var SysFlowTol float64
var LatFlowTol float64
var CrownCutoff float64
var StartDate float64
var StartTime float64
var StartDateTime float64
var EndDate float64
var EndTime float64
var EndDateTime float64
var ReportStartDate float64
var ReportStartTime float64
var ReportStart float64
var ReportTime float64
var OldRunoffTime float64
var NewRunoffTime float64
var OldRoutingTime float64
var NewRoutingTime float64
var TotalDuration float64
var ElapsedTime float64
var Temp TTemp
var Evap TEvap
var Wind TWind
var Snow TSnow
var Adjust TAdjust
var Snowmelt []TSnowmelt
var Gage []TGage
var Subcatch []TSubcatch
var Aquifer []TAquifer
var UnitHyd []TUnitHyd
var Node []TNode
var Outfall []TOutfall
var Divider []TDivider
var Storage []TStorage
var Link []TLink
var Conduit []TConduit
var Pump []TPump
var Orifice []TOrifice
var Weir []TWeir
var Outlet []TOutlet
var Pollut []TPollut
var Landuse []TLanduse
var Pattern []TPattern
var Curve []TTable
var Tseries []TTable
var Transect []TTransect
var Street []TStreet
var Shape []TShape
var Event []TEvent
var Ucf = [10][2]float64{[2]float64{float64(43200.0), float64(1097280.0)}, [2]float64{float64(12.0), float64(304.8)}, [2]float64{float64(1036800.0), float64(26334720.0)}, [2]float64{float64(1.0), float64(0.3048)}, [2]float64{float64(2.2956e-5), float64(0.92903e-5)}, [2]float64{float64(1.0), float64(0.02832)}, [2]float64{float64(1.0), float64(1.608)}, [2]float64{float64(1.0), float64(1.8)}, [2]float64{float64(2.203e-6), float64(1.0e-6)}, [2]float64{float64(43560.0), float64(3048.0)}}
var Qcf = [6]float64{float64(1.0), float64(448.831), float64(0.64632), float64(0.02832), float64(28.317), float64(2.4466)}
var IsOpenFlag int32
var IsStartedFlag int32
var SaveResultsFlag int32
var ExceptionCount int32
var DoRunoff int32
var DoRouting int32
var RoutingDuration float64

// swmm5.c:186
func swmm_run(f1 R.CPtr, f2 R.CPtr, f3 R.CPtr) int32 {
	var (
		newHour     int64
		oldHour     int64
		theDay      int64
		theHour     int64
		elapsedTime float64
	)
	_, _, _, _, _ = newHour, oldHour, theDay, theHour, elapsedTime
	newHour = 0
	oldHour = 0
	theDay = 0
	theHour = 0
	elapsedTime = float64(0.0)
	IsOpenFlag = 0
	IsStartedFlag = 0
	SaveResultsFlag = 1
	ErrorCode = 0
	writecon(L642)
	swmm_open(f1, f2, f3)
	if !(ErrorCode != 0) {
		swmm_start(1)
		if !(ErrorCode != 0) {
			writecon(L643)
			for __dw1 := true; __dw1 || ((elapsedTime > float64(0.0)) && (!(ErrorCode != 0))); __dw1 = false {
				swmm_step(&elapsedTime)
				newHour = R.D2L((float64(elapsedTime * float64(24.0))))
				if newHour > oldHour {
					theDay = R.D2L(elapsedTime)
					theHour = R.D2L((float64((float64(elapsedTime - M.Floor(elapsedTime))) * float64(24.0))))
					writecon(L644)
					R.Snprintf(Msg, 1024, L645, theDay, theHour)
					writecon(Msg)
					oldHour = newHour
				}
			}
			writecon(L646)
			writecon(L647)
		}
		swmm_end()
	}
	if (!(ErrorCode != 0)) && (Fout.mode == SCRATCH_FILE) {
		writecon(L648)
		swmm_report()
	}
	swmm_close()
	return ErrorCode
}

// swmm5.c:256
func swmm_open(f1 R.CPtr, f2 R.CPtr, f3 R.CPtr) int32 {
	{
		datetime_setDateFormat(1)
		ErrorCode = 0
		ErrorMsg.S(0, 0)
		Warnings = 0
		IsOpenFlag = 0
		IsStartedFlag = 0
		ExceptionCount = 0
		R.Strcpy(InpDir, L0)
		project_open(f1, f2, f3)
		getAbsolutePath(f1, InpDir, 260)
		if ErrorCode != 0 {
			return ErrorCode
		}
		IsOpenFlag = 1
		report_writeLogo()
		project_readInput()
		if ErrorCode != 0 {
			return ErrorCode
		}
		report_writeTitle()
		project_validate()
	}
	return ErrorCode
}

// swmm5.c:314
func swmm_start(saveResults_2 int32) int32 {
	if ErrorCode != 0 {
		return ErrorCode
	}
	if !(IsOpenFlag != 0) {
		return (R.Set(&ErrorCode, ERR_API_NOT_OPEN))
	}
	if IsStartedFlag != 0 {
		return (R.Set(&ErrorCode, ERR_API_NOT_ENDED))
	}
	if !(RptFlags.disabled != 0) {
		if RptFlags.input != 0 {
			inputrpt_writeInput()
		}
		report_writeOptions()
	}
	SaveResultsFlag = saveResults_2
	ExceptionCount = 0
	{
		ElapsedTime = float64(0.0)
		RoutingDuration = TotalDuration
		NewRunoffTime = float64(0.0)
		NewRoutingTime = float64(0.0)
		ReportTime = float64(float64(1000) * float64(ReportStep))
		TotalStepCount = 0
		ReportStepCount = 0
		NonConvergeCount = 0
		IsStartedFlag = 1
		RunoffError = float64(0.0)
		GwaterError = float64(0.0)
		FlowError = float64(0.0)
		QualError = float64(0.0)
		if !(IgnoreRainfall != 0) {
			rain_open()
		}
		if ErrorCode != 0 {
			return ErrorCode
		}
		project_init()
		if Nobjects[SUBCATCH] > 0 {
			DoRunoff = 1
		} else {
			DoRunoff = 0
		}
		if (Nobjects[NODE] > 0) && (!(IgnoreRouting != 0)) {
			DoRouting = 1
		} else {
			DoRouting = 0
		}
		output_open()
		if DoRunoff != 0 {
			runoff_open()
		}
		if !(hotstart_open() != 0) {
			return ErrorCode
		}
		if DoRouting != 0 {
			routing_open()
		}
		massbal_open()
		stats_open()
		if (!(RptFlags.disabled != 0)) && (RptFlags.controls != 0) {
			report_writeControlActionsHeading()
		}
	}
	return ErrorCode
}

// swmm5.c:410
func swmm_step(elapsedTime *float64) int32 {
	*elapsedTime = float64(0.0)
	if ErrorCode != 0 {
		return ErrorCode
	}
	if !(IsOpenFlag != 0) {
		return (R.Set(&ErrorCode, ERR_API_NOT_OPEN))
	}
	if !(IsStartedFlag != 0) {
		return (R.Set(&ErrorCode, ERR_API_NOT_STARTED))
	}
	{
		if NewRoutingTime < RoutingDuration {
			execRouting()
		}
		if SaveResultsFlag != 0 {
			saveResults()
		}
		if NewRoutingTime < RoutingDuration {
			ElapsedTime = float64(NewRoutingTime / float64(86400000.0))
		} else {
			ElapsedTime = float64(0.0)
		}
		*elapsedTime = ElapsedTime
	}
	return ErrorCode
}

// swmm5.c:466
func swmm_stride(strideStep int32, elapsedTime *float64) int32 {
	var (
		realRouteStep float64
	)
	_ = realRouteStep
	realRouteStep = RouteStep
	*elapsedTime = float64(0.0)
	if ErrorCode != 0 {
		return ErrorCode
	}
	if !(IsOpenFlag != 0) {
		return (R.Set(&ErrorCode, ERR_API_NOT_OPEN))
	}
	if !(IsStartedFlag != 0) {
		return (R.Set(&ErrorCode, ERR_API_NOT_STARTED))
	}
	RoutingDuration = float64(NewRoutingTime + float64(float64(1000.0)*float64(strideStep)))
	if (TotalDuration) <= (RoutingDuration) {
		RoutingDuration = (TotalDuration)
	} else {
		RoutingDuration = (RoutingDuration)
	}
	if float64(strideStep) < RouteStep {
		RouteStep = float64(strideStep)
	}
	for __dw1 := true; __dw1 || (((*elapsedTime) > float64(0.0)) && (!(ErrorCode != 0))); __dw1 = false {
		swmm_step(elapsedTime)
	}
	RouteStep = realRouteStep
	RoutingDuration = TotalDuration
	if NewRoutingTime < TotalDuration {
		ElapsedTime = float64(NewRoutingTime / float64(86400000.0))
	} else {
		ElapsedTime = float64(0.0)
	}
	*elapsedTime = ElapsedTime
	return ErrorCode
}

// swmm5.c:514
func execRouting() {
	var (
		nextRoutingTime float64
		routingStep     float64
	)
	_, _ = nextRoutingTime, routingStep
	nextRoutingTime = 0
	routingStep = 0
	{
		TotalStepCount++
		if !(DoRouting != 0) {
			if (WetStep) <= (ReportStep) {
				routingStep = float64((WetStep))
			} else {
				routingStep = float64((ReportStep))
			}
		} else {
			routingStep = routing_getRoutingStep(RouteModel, RouteStep)
		}
		if routingStep <= float64(0.0) {
			ErrorCode = ERR_TIMESTEP
			return
		}
		nextRoutingTime = float64(NewRoutingTime + float64(float64(1000.0)*routingStep))
		if nextRoutingTime > RoutingDuration {
			routingStep = float64((float64(RoutingDuration - NewRoutingTime)) / float64(1000.0))
			if (routingStep) >= (float64(float64(1.0) / float64(1000.0))) {
				routingStep = (routingStep)
			} else {
				routingStep = (float64(float64(1.0) / float64(1000.0)))
			}
			nextRoutingTime = RoutingDuration
		}
		if DoRunoff != 0 {
			for NewRunoffTime < nextRoutingTime {
				runoff_execute()
				if ErrorCode != 0 {
					return
				}
			}
		} else {
			climate_setState(getDateTime(NewRoutingTime))
		}
		if DoRouting != 0 {
			routing_execute(RouteModel, routingStep)
		} else {
			NewRoutingTime = nextRoutingTime
		}
	}
}

// swmm5.c:579
func saveResults() {
	if NewRoutingTime >= ReportTime {
		if RptFlags.averages != 0 {
			if NewRoutingTime == ReportTime {
				output_updateAvgResults()
			}
			output_saveResults(ReportTime)
			if NewRoutingTime > ReportTime {
				output_updateAvgResults()
			}
		} else {
			output_saveResults(ReportTime)
		}
		ReportTime = float64(ReportTime + float64(float64(1000)*float64(ReportStep)))
	} else if RptFlags.averages != 0 {
		output_updateAvgResults()
	}
}

// swmm5.c:618
func swmm_end() int32 {
	if !(IsOpenFlag != 0) {
		return (R.Set(&ErrorCode, ERR_API_NOT_OPEN))
	}
	if IsStartedFlag != 0 {
		if Fout.file != nil {
			output_end()
		}
		if (!(ErrorCode != 0)) && (RptFlags.disabled == 0) {
			massbal_report()
			stats_report()
		}
		stats_close()
		massbal_close()
		if !(IgnoreRainfall != 0) {
			rain_close()
		}
		if DoRunoff != 0 {
			runoff_close()
		}
		if DoRouting != 0 {
			routing_close(RouteModel)
		}
		hotstart_close()
		IsStartedFlag = 0
	}
	return ErrorCode
}

// swmm5.c:655
func swmm_report() int32 {
	if !(ErrorCode != 0) {
		report_writeReport()
	}
	return ErrorCode
}

// swmm5.c:669
func swmm_writeLine(line R.CPtr) {
	if IsOpenFlag != 0 {
		report_writeLine(line)
	}
}

// swmm5.c:682
func swmm_close() int32 {
	if Fout.file != nil {
		output_close()
	}
	if IsOpenFlag != 0 {
		project_close()
	}
	report_writeSysTime()
	if Finp.file != nil {
		R.Fclose(Finp.file)
	}
	if Frpt.file != nil {
		R.Fclose(Frpt.file)
	}
	if Fout.file != nil {
		R.Fclose(Fout.file)
		if Fout.mode == SCRATCH_FILE {
			R.Remove(R.CP(Fout.name[:]))
		}
	}
	IsOpenFlag = 0
	IsStartedFlag = 0
	return 0
}

// swmm5.c:708
func swmm_getMassBalErr(runoffErr *float32, flowErr *float32, qualErr *float32) int32 {
	*runoffErr = float32(float64(float64(0.0)))
	*flowErr = float32(float64(float64(0.0)))
	*qualErr = float32(float64(float64(0.0)))
	if (IsOpenFlag != 0) && (!(IsStartedFlag != 0)) {
		*runoffErr = float32(float64(RunoffError))
		*flowErr = float32(float64(FlowError))
		*qualErr = float32(float64(QualError))
	}
	return 0
}

// swmm5.c:734
func swmm_getVersion() int32 {
	return 52004
}

// swmm5.c:749
func swmm_getWarnings() int32 {
	return Warnings
}

// swmm5.c:760
func swmm_getError(errMsg R.CPtr, msgLen int32) int32 {
	if (ErrorCode > 0) && (R.Strlen(ErrorMsg) == 0) {
		error_getMsg(ErrorCode, ErrorMsg)
	}
	sstrncpy(errMsg, ErrorMsg, uint64(msgLen))
	if (msgLen > 0) && (errMsg.G(0) == 10) {
		errMsg.S(0, 32)
	}
	return ErrorCode
}

// swmm5.c:780
func swmm_getCount(objType int32) int32 {
	if !(IsOpenFlag != 0) {
		return 0
	}
	if (objType < swmm_GAGE) || (objType > swmm_LINK) {
		return 0
	}
	return Nobjects[objType]
}

// swmm5.c:795
func swmm_getName(objType int32, index int32, name R.CPtr, size int32) {
	var (
		idName R.CPtr
	)
	_ = idName
	idName = (R.CPtr{})
	name.S(0, 0)
	if !(IsOpenFlag != 0) {
		return
	}
	if (objType < swmm_GAGE) || (objType > swmm_LINK) {
		return
	}
	if (index < 0) || (index >= Nobjects[objType]) {
		return
	}
	switch objType {
	case GAGE:
		idName = Gage[index].ID
		break
	case SUBCATCH:
		idName = Subcatch[index].ID
		break
	case NODE:
		idName = Node[index].ID
		break
	case LINK:
		idName = Link[index].ID
		break
	}
	if !idName.IsNull() {
		sstrncpy(name, idName, uint64(size))
	}
}

// swmm5.c:826
func swmm_getIndex(objType int32, name R.CPtr) int32 {
	if !(IsOpenFlag != 0) {
		return (-1)
	}
	if (objType < swmm_GAGE) || (objType > swmm_LINK) {
		return (-1)
	}
	return project_findObject(objType, name)
}

// swmm5.c:842
func swmm_getValue(property int32, index int32) float64 {
	if !(IsOpenFlag != 0) {
		return float64(0)
	}
	if property < 100 {
		return getSystemValue(property)
	}
	if property < 200 {
		return getGageValue(property, index)
	}
	if property < 300 {
		return getSubcatchValue(property, index)
	}
	if property < 400 {
		return getNodeValue(property, index)
	}
	if property < 500 {
		return getLinkValue(property, index)
	}
	return float64(0)
}

// swmm5.c:867
func swmm_setValue(property int32, index int32, value float64) {
	if !(IsOpenFlag != 0) {
		return
	}
	switch property {
	case swmm_GAGE_RAINFALL:
		if (index < 0) || (index >= Nobjects[GAGE]) {
			return
		}
		if value >= float64(0.0) {
			Gage[index].apiRainfall = value
		}
		return
	case swmm_SUBCATCH_RPTFLAG:
		if ((!(IsStartedFlag != 0)) && (index >= 0)) && (index < Nobjects[SUBCATCH]) {
			Subcatch[index].rptFlag = R.B2I((value > float64(0.0)))
		}
		return
	case swmm_NODE_LATFLOW:
		setNodeLatFlow(index, value)
		return
	case swmm_NODE_HEAD:
		setOutfallStage(index, value)
		return
	case swmm_NODE_RPTFLAG:
		if ((!(IsStartedFlag != 0)) && (index >= 0)) && (index < Nobjects[NODE]) {
			Node[index].rptFlag = R.B2I((value > float64(0.0)))
		}
		return
	case swmm_LINK_SETTING:
		setLinkSetting(index, value)
		return
	case swmm_LINK_RPTFLAG:
		if ((!(IsStartedFlag != 0)) && (index >= 0)) && (index < Nobjects[LINK]) {
			Link[index].rptFlag = R.B2I((value > float64(0.0)))
		}
		return
	case swmm_ROUTESTEP:
		setRoutingStep(value)
		return
	case swmm_REPORTSTEP:
		if (!(IsStartedFlag != 0)) && (value > float64(0)) {
			ReportStep = R.D2I(value)
		}
		return
	case swmm_NOREPORT:
		if !(IsStartedFlag != 0) {
			RptFlags.disabled = R.B2I((value > float64(0.0)))
		}
		return
	}
}

// swmm5.c:922
func swmm_getSavedValue(property int32, index int32, period int32) float64 {
	if !(IsOpenFlag != 0) {
		return float64(0)
	}
	if IsStartedFlag != 0 {
		return float64(0)
	}
	if (period < 1) || (int64(period) > Nperiods) {
		return float64(0)
	}
	if property == swmm_CURRENTDATE {
		return getSavedDate(period)
	}
	if (property >= 200) && (property < 300) {
		return getSavedSubcatchValue(property, index, period)
	}
	if property < 400 {
		return getSavedNodeValue(property, index, period)
	}
	if property < 500 {
		return getSavedLinkValue(property, index, period)
	}
	return float64(0)
}

// swmm5.c:949
func swmm_decodeDate(date float64, year *int32, month *int32, day *int32, hour *int32, minute *int32, second *int32, dayOfWeek *int32) {
	datetime_decodeDate(date, year, month, day)
	datetime_decodeTime(date, hour, minute, second)
	*dayOfWeek = datetime_dayOfWeek(date)
}

// swmm5.c:966
func getGageValue(property int32, index int32) float64 {
	if (index < 0) || (index >= Nobjects[GAGE]) {
		return float64(0)
	}
	if property == swmm_GAGE_RAINFALL {
		return Gage[index].reportRainfall
	}
	return float64(0)
}

// swmm5.c:982
func getSubcatchValue(property int32, index int32) float64 {
	var (
		subcatch *TSubcatch
	)
	_ = subcatch
	subcatch = nil
	if (index < 0) || (index >= Nobjects[SUBCATCH]) {
		return float64(0)
	}
	subcatch = &Subcatch[index]
	switch property {
	case swmm_SUBCATCH_AREA:
		return float64(subcatch.area * UCF(LANDAREA))
	case swmm_SUBCATCH_RAINGAGE:
		return float64(subcatch.gage)
	case swmm_SUBCATCH_RAINFALL:
		if subcatch.gage >= 0 {
			return Gage[subcatch.gage].reportRainfall
		} else {
			return float64(0.0)
		}
	case swmm_SUBCATCH_EVAP:
		return float64(subcatch.evapLoss * UCF(EVAPRATE))
	case swmm_SUBCATCH_INFIL:
		return float64(subcatch.infilLoss * UCF(RAINFALL))
	case swmm_SUBCATCH_RUNOFF:
		return float64(subcatch.newRunoff * UCF(FLOW))
	case swmm_SUBCATCH_RPTFLAG:
		return float64(R.B2I((subcatch.rptFlag > 0)))
	default:
		return float64(0)
	}
	return 0
}

// swmm5.c:1019
func getNodeValue(property int32, index int32) float64 {
	var (
		node *TNode
	)
	_ = node
	node = nil
	if (index < 0) || (index >= Nobjects[NODE]) {
		return float64(0)
	}
	node = &Node[index]
	switch property {
	case swmm_NODE_TYPE:
		return float64(node.type_)
	case swmm_NODE_ELEV:
		return float64(node.invertElev * UCF(LENGTH))
	case swmm_NODE_MAXDEPTH:
		return float64(node.fullDepth * UCF(LENGTH))
	case swmm_NODE_DEPTH:
		return float64(node.newDepth * UCF(LENGTH))
	case swmm_NODE_HEAD:
		return float64((float64(node.newDepth + node.invertElev)) * UCF(LENGTH))
	case swmm_NODE_VOLUME:
		return float64(node.newVolume * UCF(VOLUME))
	case swmm_NODE_LATFLOW:
		return float64(node.newLatFlow * UCF(FLOW))
	case swmm_NODE_INFLOW:
		return float64(node.inflow * UCF(FLOW))
	case swmm_NODE_OVERFLOW:
		return float64(node.overflow * UCF(FLOW))
	case swmm_NODE_RPTFLAG:
		return float64(R.B2I((node.rptFlag > 0)))
	default:
		return float64(0)
	}
	return 0
}

// swmm5.c:1059
func getLinkValue(property int32, index int32) float64 {
	var (
		link *TLink
	)
	_ = link
	link = nil
	if (index < 0) || (index >= Nobjects[LINK]) {
		return float64(0)
	}
	link = &Link[index]
	switch property {
	case swmm_LINK_TYPE:
		return float64(link.type_)
	case swmm_LINK_NODE1:
		return float64(link.node1)
	case swmm_LINK_NODE2:
		return float64(link.node2)
	case swmm_LINK_LENGTH:
		if link.type_ == CONDUIT {
			return float64(Conduit[link.subIndex].length * UCF(LENGTH))
		} else {
			return float64(0)
		}
	case swmm_LINK_SLOPE:
		if link.type_ == CONDUIT {
			return Conduit[link.subIndex].slope
		} else {
			return float64(0)
		}
		break
	case swmm_LINK_FULLDEPTH:
		return float64(link.xsect.yFull * UCF(LENGTH))
	case swmm_LINK_FULLFLOW:
		return float64(link.qFull * UCF(FLOW))
	case swmm_LINK_FLOW:
		return float64(float64(link.newFlow*UCF(FLOW)) * float64(link.direction))
	case swmm_LINK_VELOCITY:
		return float64(link_getVelocity(index, M.Fabs(link.newFlow), link.newDepth) * UCF(LENGTH))
	case swmm_LINK_DEPTH:
		return float64(link.newDepth * UCF(LENGTH))
	case swmm_LINK_TOPWIDTH:
		if link.type_ == CONDUIT {
			return float64(xsect_getWofY(&link.xsect, link.newDepth) * UCF(LENGTH))
		} else {
			return float64(0)
		}
	case swmm_LINK_SETTING:
		return link.setting
	case swmm_LINK_TIMEOPEN:
		if link.setting > float64(0.0) {
			return float64((float64(getDateTime(NewRoutingTime) - link.timeLastSet)) * float64(24.0))
		} else {
			return float64(0)
		}
	case swmm_LINK_TIMECLOSED:
		if link.setting == float64(0.0) {
			return float64((float64(getDateTime(NewRoutingTime) - link.timeLastSet)) * float64(24.0))
		} else {
			return float64(0)
		}
	case swmm_LINK_RPTFLAG:
		return float64(R.B2I((link.rptFlag > 0)))
	default:
		return float64(0)
	}
	return 0
}

// swmm5.c:1126
func getSystemValue(property int32) float64 {
	switch property {
	case swmm_STARTDATE:
		return StartDateTime
	case swmm_CURRENTDATE:
		return float64(StartDateTime + ElapsedTime)
	case swmm_ELAPSEDTIME:
		return ElapsedTime
	case swmm_ROUTESTEP:
		return RouteStep
	case swmm_MAXROUTESTEP:
		return getMaxRouteStep()
	case swmm_REPORTSTEP:
		return float64(ReportStep)
	case swmm_TOTALSTEPS:
		return float64(Nperiods)
	case swmm_NOREPORT:
		return float64(RptFlags.disabled)
	case swmm_FLOWUNITS:
		return float64(FlowUnits)
	default:
		return float64(0)
	}
	return 0
}

// swmm5.c:1159
func setNodeLatFlow(index int32, value float64) {
	if (index < 0) || (index >= Nobjects[NODE]) {
		return
	}
	Node[index].apiExtInflow = float64(value / UCF(FLOW))
}

// swmm5.c:1173
func setOutfallStage(index int32, value float64) {
	var (
		node *TNode
	)
	_ = node
	node = nil
	if (index < 0) || (index >= Nobjects[NODE]) {
		return
	}
	node = &Node[index]
	if node.type_ != OUTFALL {
		return
	}
	Outfall[node.subIndex].fixedStage = float64(value / UCF(LENGTH))
	Outfall[node.subIndex].type_ = FIXED_OUTFALL
}

// swmm5.c:1192
func setLinkSetting(index int32, value float64) {
	var (
		link *TLink
	)
	_ = link
	link = nil
	if (index < 0) || (index >= Nobjects[LINK]) {
		return
	}
	link = &Link[index]
	if (value < float64(0.0)) || (link.type_ == CONDUIT) {
		return
	}
	if (link.type_ != PUMP) && (value > float64(1.0)) {
		value = float64(1.0)
	}
	if link.targetSetting == value {
		return
	}
	link.targetSetting = value
	if float64(link.targetSetting*link.setting) == float64(0.0) {
		link.timeLastSet = float64(StartDateTime + ElapsedTime)
	}
	link_setSetting(index, float64(0.0))
}

// swmm5.c:1217
func getSavedDate(period int32) float64 {
	var (
		days float64
	)
	_ = days
	days = 0
	output_readDateTime(int64(period), &days)
	return days
}

// swmm5.c:1230
func getSavedSubcatchValue(property int32, index int32, period int32) float64 {
	var (
		outIndex int32
	)
	_ = outIndex
	outIndex = (Subcatch[index].rptFlag - 1)
	if outIndex < 0 {
		return float64(0)
	}
	output_readSubcatchResults(int64(period), outIndex)
	switch property {
	case swmm_SUBCATCH_RAINFALL:
		return float64(SubcatchResults[SUBCATCH_RAINFALL])
	case swmm_SUBCATCH_EVAP:
		return float64(SubcatchResults[SUBCATCH_EVAP])
	case swmm_SUBCATCH_INFIL:
		return float64(SubcatchResults[SUBCATCH_INFIL])
	case swmm_SUBCATCH_RUNOFF:
		return float64(SubcatchResults[SUBCATCH_RUNOFF])
	default:
		return float64(0)
	}
	return 0
}

// swmm5.c:1265
func getSavedNodeValue(property int32, index int32, period int32) float64 {
	var (
		outIndex int32
	)
	_ = outIndex
	outIndex = (Node[index].rptFlag - 1)
	if outIndex < 0 {
		return float64(0)
	}
	output_readNodeResults(int64(period), outIndex)
	switch property {
	case swmm_NODE_DEPTH:
		return float64(NodeResults[NODE_DEPTH])
	case swmm_NODE_HEAD:
		return float64(NodeResults[NODE_HEAD])
	case swmm_NODE_VOLUME:
		return float64(NodeResults[NODE_VOLUME])
	case swmm_NODE_LATFLOW:
		return float64(NodeResults[NODE_LATFLOW])
	case swmm_NODE_INFLOW:
		return float64(NodeResults[NODE_INFLOW])
	case swmm_NODE_OVERFLOW:
		return float64(NodeResults[NODE_OVERFLOW])
	default:
		return float64(0)
	}
	return 0
}

// swmm5.c:1304
func getSavedLinkValue(property int32, index int32, period int32) float64 {
	var (
		y_2      float64
		w        float64
		outIndex int32
	)
	_, _, _ = y_2, w, outIndex
	y_2 = 0
	w = 0
	outIndex = (Link[index].rptFlag - 1)
	if outIndex < 0 {
		return float64(0)
	}
	output_readLinkResults(int64(period), outIndex)
	switch property {
	case swmm_LINK_FLOW:
		return float64(LinkResults[LINK_FLOW])
	case swmm_LINK_DEPTH:
		return float64(LinkResults[LINK_DEPTH])
	case swmm_LINK_VELOCITY:
		return float64(LinkResults[LINK_VELOCITY])
	case swmm_LINK_TOPWIDTH:
		y_2 = float64(float64(LinkResults[LINK_DEPTH]) / UCF(LENGTH))
		w = xsect_getWofY(&Link[index].xsect, y_2)
		return float64(w * UCF(LENGTH))
	case swmm_LINK_SETTING:
		return float64(LinkResults[LINK_CAPACITY])
	default:
		return float64(0)
	}
	return 0
}

// swmm5.c:1345
func getMaxRouteStep() float64 {
	var (
		tmpCourantFactor float64
		result           float64
	)
	_, _ = tmpCourantFactor, result
	tmpCourantFactor = CourantFactor
	result = RouteStep
	if (!(IsStartedFlag != 0)) || (RouteModel != DW) {
		return result
	}
	CourantFactor = float64(1.0)
	result = routing_getRoutingStep(RouteModel, MinRouteStep)
	CourantFactor = tmpCourantFactor
	return result
}

// swmm5.c:1360
func setRoutingStep(value float64) {
	if value <= float64(0.0) {
		return
	}
	if value <= MinRouteStep {
		value = MinRouteStep
	}
	CourantFactor = float64(0.0)
	RouteStep = value
}

// swmm5.c:1378
func UCF(u int32) float64 {
	if u < FLOW {
		return Ucf[u][UnitSystem]
	} else {
		return Qcf[FlowUnits]
	}
	return 0
}

// swmm5.c:1453
func strcomp(s1 R.CPtr, s2 R.CPtr) int32 {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 0; (func() int32 {
		if ((s1.G(i)) >= 97) && ((s1.G(i)) <= 122) {
			return ((s1.G(i)) & (^32))
		}
		return (s1.G(i))
	}()) == (func() int32 {
		if ((s2.G(i)) >= 97) && ((s2.G(i)) <= 122) {
			return ((s2.G(i)) & (^32))
		}
		return (s2.G(i))
	}()); i++ {
		if (!(s1.G((i + 1)) != 0)) && (!(s2.G((i + 1)) != 0)) {
			return (1)
		}
	}
	return (0)
}

// swmm5.c:1518
func getElapsedTime(aDate float64, days *int32, hrs *int32, mins *int32) {
	var (
		x    float64
		secs int32
	)
	_, _ = x, secs
	x = 0
	secs = 0
	x = float64(aDate - ReportStart)
	if x <= float64(0.0) {
		*days = 0
		*hrs = 0
		*mins = 0
	} else {
		*days = R.D2I(x)
		datetime_decodeTime(x, hrs, mins, &secs)
	}
}

// swmm5.c:1543
func getDateTime(elapsedMsec float64) float64 {
	return datetime_addSeconds(StartDateTime, float64((float64(elapsedMsec+float64(1)))/float64(1000.0)))
}

// swmm5.c:1556
func isRelativePath(fname R.CPtr) int32 {
	if !(R.Strchr(fname, 58)).IsNull() {
		return 0
	}
	if fname.G(0) == 92 {
		return 0
	}
	if fname.G(0) == 47 {
		return 0
	}
	return 1
}

// swmm5.c:1617
func addAbsolutePath(fname R.CPtr) R.CPtr {
	var (
		n_2    uint64
		buffer R.CPtr
	)
	_, _ = n_2, buffer
	n_2 = 0
	buffer = R.Buf(259)
	if isRelativePath(fname) != 0 {
		n_2 = uint64(R.Snprintf(buffer, 259, L649, InpDir, fname))
		if n_2 > 0 {
			sstrncpy(fname, buffer, 259)
		}
	}
	return fname
}

// swmm5.c:1638
func writecon(s R.CPtr) {
	R.Fprintf(R.Stdout, L169, s)
	R.Fflush(R.Stdout)
}
