// Generated from EPA SWMM 5.2.4 (inputrpt.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

// inputrpt.c:27
func inputrpt_writeInput() {
	var (
		m        int32
		i        int32
		k        int32
		lidCount int32
	)
	_, _, _, _ = m, i, k, lidCount
	m = 0
	i = 0
	k = 0
	lidCount = 0
	if ErrorCode != 0 {
		return
	}
	report_writeLine((L0))
	report_writeLine((L206))
	report_writeLine((L207))
	report_writeLine((L206))
	R.Fprintf(Frpt.file, L208, Nobjects[GAGE])
	R.Fprintf(Frpt.file, L209, Nobjects[SUBCATCH])
	R.Fprintf(Frpt.file, L210, Nobjects[NODE])
	R.Fprintf(Frpt.file, L211, Nobjects[LINK])
	R.Fprintf(Frpt.file, L212, Nobjects[POLLUT])
	R.Fprintf(Frpt.file, L213, Nobjects[LANDUSE])
	if Nobjects[POLLUT] > 0 {
		report_writeLine((L0))
		report_writeLine((L0))
		report_writeLine((L214))
		report_writeLine((L215))
		report_writeLine((L214))
		R.Fprintf(Frpt.file, L216)
		R.Fprintf(Frpt.file, L217)
		R.Fprintf(Frpt.file, L218)
		for i = 0; i < Nobjects[POLLUT]; i++ {
			R.Fprintf(Frpt.file, L219, Pollut[i].ID, QualUnitsWords[Pollut[i].units], Pollut[i].pptConcen, Pollut[i].gwConcen, float64(Pollut[i].kDecay*float64(86400)))
			if Pollut[i].coPollut >= 0 {
				R.Fprintf(Frpt.file, L220, Pollut[Pollut[i].coPollut].ID, Pollut[i].coFraction)
			}
		}
	}
	if Nobjects[LANDUSE] > 0 {
		report_writeLine((L0))
		report_writeLine((L0))
		report_writeLine((L221))
		report_writeLine((L222))
		report_writeLine((L221))
		R.Fprintf(Frpt.file, L223)
		R.Fprintf(Frpt.file, L224)
		R.Fprintf(Frpt.file, L225)
		for i = 0; i < Nobjects[LANDUSE]; i++ {
			R.Fprintf(Frpt.file, L226, Landuse[i].ID, Landuse[i].sweepInterval, Landuse[i].sweepRemoval, Landuse[i].sweepDays0)
		}
	}
	if Nobjects[GAGE] > 0 {
		report_writeLine((L0))
		report_writeLine((L0))
		report_writeLine((L227))
		report_writeLine((L228))
		report_writeLine((L227))
		R.Fprintf(Frpt.file, L229)
		R.Fprintf(Frpt.file, L230)
		R.Fprintf(Frpt.file, L231)
		for i = 0; i < Nobjects[GAGE]; i++ {
			if Gage[i].tSeries >= 0 {
				R.Fprintf(Frpt.file, L232, Gage[i].ID, Tseries[Gage[i].tSeries].ID)
				R.Fprintf(Frpt.file, L233, RainTypeWords[Gage[i].rainType], ((Gage[i].rainInterval) / 60))
			} else {
				R.Fprintf(Frpt.file, L234, Gage[i].ID, R.CP(Gage[i].fname[:]))
			}
		}
	}
	if Nobjects[SUBCATCH] > 0 {
		report_writeLine((L0))
		report_writeLine((L0))
		report_writeLine((L235))
		report_writeLine((L236))
		report_writeLine((L235))
		R.Fprintf(Frpt.file, L237)
		R.Fprintf(Frpt.file, L238)
		for i = 0; i < Nobjects[SUBCATCH]; i++ {
			R.Fprintf(Frpt.file, L239, Subcatch[i].ID, float64(Subcatch[i].area*UCF(LANDAREA)), float64(Subcatch[i].width*UCF(LENGTH)), float64(Subcatch[i].fracImperv*float64(100.0)), float64(Subcatch[i].slope*float64(100.0)), Gage[Subcatch[i].gage].ID)
			if Subcatch[i].outNode >= 0 {
				R.Fprintf(Frpt.file, L240, Node[Subcatch[i].outNode].ID)
			} else if Subcatch[i].outSubcatch >= 0 {
				R.Fprintf(Frpt.file, L240, Subcatch[Subcatch[i].outSubcatch].ID)
			}
			if Subcatch[i].lidArea != 0 {
				lidCount++
			}
		}
	}
	if lidCount > 0 {
		lid_writeSummary()
	}
	if Nobjects[NODE] > 0 {
		report_writeLine((L0))
		report_writeLine((L0))
		report_writeLine((L241))
		report_writeLine((L242))
		report_writeLine((L241))
		R.Fprintf(Frpt.file, L243)
		R.Fprintf(Frpt.file, L244)
		R.Fprintf(Frpt.file, L245)
		for i = 0; i < Nobjects[NODE]; i++ {
			R.Fprintf(Frpt.file, L246, Node[i].ID, NodeTypeWords[(Node[i].type_-JUNCTION)], float64(Node[i].invertElev*UCF(LENGTH)), float64(Node[i].fullDepth*UCF(LENGTH)), float64(float64(Node[i].pondedArea*UCF(LENGTH))*UCF(LENGTH)))
			if ((Node[i].extInflow != nil) || (Node[i].dwfInflow != nil)) || (Node[i].rdiiInflow != nil) {
				R.Fprintf(Frpt.file, L247)
			}
		}
	}
	if Nobjects[LINK] > 0 {
		report_writeLine((L0))
		report_writeLine((L0))
		report_writeLine((L241))
		report_writeLine((L248))
		report_writeLine((L241))
		R.Fprintf(Frpt.file, L249)
		R.Fprintf(Frpt.file, L250)
		for i = 0; i < Nobjects[LINK]; i++ {
			if Link[i].direction == 1 {
				R.Fprintf(Frpt.file, L251, Link[i].ID, Node[Link[i].node1].ID, Node[Link[i].node2].ID)
			} else {
				R.Fprintf(Frpt.file, L251, Link[i].ID, Node[Link[i].node2].ID, Node[Link[i].node1].ID)
			}
			if Link[i].type_ == PUMP {
				k = Link[i].subIndex
				R.Fprintf(Frpt.file, L252, PumpTypeWords[Pump[k].type_])
			} else {
				R.Fprintf(Frpt.file, L253, LinkTypeWords[(Link[i].type_-CONDUIT)])
			}
			if Link[i].type_ == CONDUIT {
				k = Link[i].subIndex
				R.Fprintf(Frpt.file, L254, float64(Conduit[k].length*UCF(LENGTH)), float64(float64(Conduit[k].slope*float64(100.0))*float64(Link[i].direction)), Conduit[k].roughness)
			}
		}
		report_writeLine((L0))
		report_writeLine((L0))
		report_writeLine((L255))
		report_writeLine((L256))
		report_writeLine((L255))
		R.Fprintf(Frpt.file, L257)
		R.Fprintf(Frpt.file, L258)
		R.Fprintf(Frpt.file, L259)
		for i = 0; i < Nobjects[LINK]; i++ {
			if Link[i].type_ == CONDUIT {
				k = Link[i].subIndex
				R.Fprintf(Frpt.file, L260, Link[i].ID)
				if Link[i].xsect.type_ == CUSTOM {
					R.Fprintf(Frpt.file, L261, Curve[Link[i].xsect.transect].ID)
				} else if Link[i].xsect.type_ == IRREGULAR {
					R.Fprintf(Frpt.file, L261, Transect[Link[i].xsect.transect].ID)
				} else if Link[i].xsect.type_ == STREET_XSECT {
					R.Fprintf(Frpt.file, L261, Street[Link[i].xsect.transect].ID)
				} else {
					R.Fprintf(Frpt.file, L261, XsectTypeWords[Link[i].xsect.type_])
				}
				R.Fprintf(Frpt.file, L262, float64(Link[i].xsect.yFull*UCF(LENGTH)), float64(float64(Link[i].xsect.aFull*UCF(LENGTH))*UCF(LENGTH)), float64(Link[i].xsect.rFull*UCF(LENGTH)), float64(Link[i].xsect.wMax*UCF(LENGTH)), Conduit[k].barrels, float64(Link[i].qFull*UCF(FLOW)))
			}
		}
	}
	if Nobjects[SHAPE] > 0 {
		report_writeLine((L0))
		report_writeLine((L0))
		report_writeLine((L206))
		report_writeLine((L263))
		report_writeLine((L206))
		for i = 0; i < Nobjects[SHAPE]; i++ {
			k = Shape[i].curve
			R.Fprintf(Frpt.file, L264, Curve[k].ID)
			R.Fprintf(Frpt.file, L265)
			for m = 1; m < 51; m++ {
				if (m % 5) == 1 {
					R.Fprintf(Frpt.file, L266)
				}
				R.Fprintf(Frpt.file, L267, Shape[i].areaTbl[m])
			}
			R.Fprintf(Frpt.file, L268)
			for m = 1; m < 51; m++ {
				if (m % 5) == 1 {
					R.Fprintf(Frpt.file, L266)
				}
				R.Fprintf(Frpt.file, L267, Shape[i].hradTbl[m])
			}
			R.Fprintf(Frpt.file, L269)
			for m = 1; m < 51; m++ {
				if (m % 5) == 1 {
					R.Fprintf(Frpt.file, L266)
				}
				R.Fprintf(Frpt.file, L267, Shape[i].widthTbl[m])
			}
		}
	}
	if Nobjects[TRANSECT] > 0 {
		report_writeLine((L0))
		report_writeLine((L0))
		report_writeLine((L227))
		report_writeLine((L270))
		report_writeLine((L227))
		for i = 0; i < Nobjects[TRANSECT]; i++ {
			R.Fprintf(Frpt.file, L271, Transect[i].ID)
			R.Fprintf(Frpt.file, L265)
			for m = 1; m < 51; m++ {
				if (m % 5) == 1 {
					R.Fprintf(Frpt.file, L266)
				}
				R.Fprintf(Frpt.file, L267, Transect[i].areaTbl[m])
			}
			R.Fprintf(Frpt.file, L268)
			for m = 1; m < 51; m++ {
				if (m % 5) == 1 {
					R.Fprintf(Frpt.file, L266)
				}
				R.Fprintf(Frpt.file, L267, Transect[i].hradTbl[m])
			}
			R.Fprintf(Frpt.file, L269)
			for m = 1; m < 51; m++ {
				if (m % 5) == 1 {
					R.Fprintf(Frpt.file, L266)
				}
				R.Fprintf(Frpt.file, L267, Transect[i].widthTbl[m])
			}
		}
	}
	if Nobjects[STREET] > 0 {
		report_writeLine((L0))
		report_writeLine((L0))
		report_writeLine((L272))
		report_writeLine((L273))
		report_writeLine((L272))
		for i = 0; i < Nobjects[STREET]; i++ {
			R.Fprintf(Frpt.file, L274, Street[i].ID)
			R.Fprintf(Frpt.file, L265)
			for m = 1; m < Street[i].transect.nTbl; m++ {
				if (m % 5) == 1 {
					R.Fprintf(Frpt.file, L266)
				}
				R.Fprintf(Frpt.file, L267, Street[i].transect.areaTbl[m])
			}
			R.Fprintf(Frpt.file, L268)
			for m = 1; m < Street[i].transect.nTbl; m++ {
				if (m % 5) == 1 {
					R.Fprintf(Frpt.file, L266)
				}
				R.Fprintf(Frpt.file, L267, Street[i].transect.hradTbl[m])
			}
			R.Fprintf(Frpt.file, L269)
			for m = 1; m < Street[i].transect.nTbl; m++ {
				if (m % 5) == 1 {
					R.Fprintf(Frpt.file, L266)
				}
				R.Fprintf(Frpt.file, L267, Street[i].transect.widthTbl[m])
			}
		}
	}
	report_writeLine((L0))
}
