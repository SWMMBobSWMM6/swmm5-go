// Generated from EPA SWMM 5.2.4 (table.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

// table.c:51
func table_interpolate(x float64, x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	var (
		dx float64
	)
	_ = dx
	dx = float64(x2 - x1)
	if M.Fabs(dx) < float64(1.0e-20) {
		return float64((float64(y1 + y2)) / float64(2.0))
	}
	return float64(y1 + float64(float64((float64(x-x1))*(float64(y2-y1)))/dx))
}

// table.c:67
func table_readCurve(tok []R.CPtr, ntoks int32) int32 {
	var (
		j   int32
		m   int32
		k   int32
		k1  int32
		x   float64
		y_2 float64
	)
	_, _, _, _, _, _ = j, m, k, k1, x, y_2
	j = 0
	m = 0
	k = 0
	k1 = 1
	x = 0
	y_2 = 0
	if ntoks < 2 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(CURVE, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	if Curve[j].ID.IsNull() {
		Curve[j].ID = project_findID(CURVE, tok[0])
		m = findmatch(tok[1], CurveTypeWords[:])
		if m < 0 {
			return error_setInpError(ERR_KEYWORD, tok[1])
		}
		Curve[j].curveType = m
		if ntoks == 2 {
			return 0
		}
		k1 = 2
	}
	for k = k1; k < ntoks; k = (k + 2) {
		if (k + 1) >= ntoks {
			return error_setInpError(ERR_ITEMS, L0)
		}
		if !(getDouble(tok[k], &x) != 0) {
			return error_setInpError(ERR_NUMBER, tok[k])
		}
		if !(getDouble(tok[(k+1)], &y_2) != 0) {
			return error_setInpError(ERR_NUMBER, tok[(k+1)])
		}
		table_addEntry(&Curve[j], x, y_2)
	}
	return 0
}

// table.c:113
func table_readTimeseries(tok []R.CPtr, ntoks int32) int32 {
	var (
		j     int32
		k     int32
		state int32
		x     float64
		y_2   float64
		d     float64
		t     float64
		fname R.CPtr
	)
	_, _, _, _, _, _, _, _ = j, k, state, x, y_2, d, t, fname
	j = 0
	k = 0
	state = 0
	x = 0
	y_2 = 0
	d = 0
	t = 0
	fname = R.Buf(260)
	if ntoks < 3 {
		return error_setInpError(ERR_ITEMS, L0)
	}
	j = project_findObject(TSERIES, tok[0])
	if j < 0 {
		return error_setInpError(ERR_NAME, tok[0])
	}
	if Tseries[j].ID.IsNull() {
		Tseries[j].ID = project_findID(TSERIES, tok[0])
	}
	if strcomp(tok[1], L1) != 0 {
		sstrncpy(fname, tok[2], 259)
		sstrncpy(R.CP(Tseries[j].file.name[:]), addAbsolutePath(fname), 259)
		Tseries[j].file.mode = USE_FILE
		return 0
	}
	x = float64(0.0)
	k = 1
	state = 1
	for k < ntoks {
		switch state {
		case 1:
			if datetime_strToDate(tok[k], &d) != 0 {
				Tseries[j].lastDate = d
				k++
			}
			state = 2
			break
		case 2:
			if k >= ntoks {
				return error_setInpError(ERR_ITEMS, L0)
			}
			if getDouble(tok[k], &t) != 0 {
				t = float64(t / (float64(24.0)))
			} else if !(datetime_strToTime(tok[k], &t) != 0) {
				return error_setInpError(ERR_NUMBER, tok[k])
			}
			x = float64(Tseries[j].lastDate + t)
			k++
			state = 3
			break
		case 3:
			if k >= ntoks {
				return error_setInpError(ERR_ITEMS, L0)
			}
			if !(getDouble(tok[k], &y_2) != 0) {
				return error_setInpError(ERR_NUMBER, tok[k])
			}
			table_addEntry(&Tseries[j], x, y_2)
			k++
			state = 1
			break
		}
	}
	return 0
}

// table.c:208
func table_addEntry(table *TTable, x float64, y_2 float64) int32 {
	var (
		entry *TTableEntry
	)
	_ = entry
	entry = nil
	entry = new(TTableEntry)
	if !(entry != nil) {
		return 0
	}
	entry.x = x
	entry.y = y_2
	entry.next = nil
	if table.firstEntry == nil {
		table.firstEntry = entry
	} else {
		table.lastEntry.next = entry
	}
	table.lastEntry = entry
	return 1
}

// table.c:231
func table_deleteEntries(table *TTable) {
	var (
		entry     *TTableEntry
		nextEntry *TTableEntry
	)
	_, _ = entry, nextEntry
	entry = nil
	nextEntry = nil
	entry = table.firstEntry
	for entry != nil {
		nextEntry = entry.next
		_ = entry
		entry = nextEntry
	}
	table.firstEntry = nil
	table.lastEntry = nil
	table.thisEntry = nil
	if table.file.file != nil {
		R.Fclose(table.file.file)
		table.file.file = (nil)
	}
}

// table.c:260
func table_init(table *TTable) {
	table.ID = (R.CPtr{})
	table.refersTo = (-1)
	table.firstEntry = nil
	table.lastEntry = nil
	table.thisEntry = table.firstEntry
	table.lastDate = float64(0.0)
	table.x1 = float64(0.0)
	table.x2 = float64(0.0)
	table.y1 = float64(0.0)
	table.y2 = float64(0.0)
	table.dxMin = float64(0.0)
	table.file.mode = NO_FILE
	table.file.file = (nil)
	table.curveType = (-1)
}

// table.c:285
func table_validate(table *TTable) int32 {
	var (
		result int32
		x1     float64
		x2     float64
		y1     float64
		y2     float64
		dx     float64
		dxMin  float64
	)
	_, _, _, _, _, _, _ = result, x1, x2, y1, y2, dx, dxMin
	result = 0
	x1 = 0
	x2 = 0
	y1 = 0
	y2 = 0
	dx = 0
	dxMin = float64(10000000000.0)
	if table.file.mode == USE_FILE {
		table.file.file = R.Fopen(R.CP(table.file.name[:]), L13)
		if table.file.file == nil {
			return ERR_TABLE_FILE_OPEN
		}
	}
	result = table_getFirstEntry(table, &x1, &y1)
	if (!(result != 0)) && (table.file.mode == USE_FILE) {
		return ERR_TABLE_FILE_READ
	}
	for table_getNextEntry(table, &x2, &y2) != 0 {
		dx = float64(x2 - x1)
		if dx <= float64(0.0) {
			table.x2 = x2
			return ERR_CURVE_SEQUENCE
		}
		if (dxMin) <= (dx) {
			dxMin = (dxMin)
		} else {
			dxMin = (dx)
		}
		x1 = x2
	}
	table.dxMin = dxMin
	if (table.file.mode == USE_FILE) && (!(R.Feof(table.file.file) != 0)) {
		return ERR_TABLE_FILE_READ
	}
	return 0
}

// table.c:331
func table_getFirstEntry(table *TTable, x *float64, y_2 *float64) int32 {
	var (
		entry *TTableEntry
	)
	_ = entry
	entry = nil
	*x = float64(0)
	*y_2 = float64(0.0)
	if table.file.mode == USE_FILE {
		if table.file.file == nil {
			return 0
		}
		R.Rewind(table.file.file)
		return table_getNextFileEntry(table, x, y_2)
	}
	entry = table.firstEntry
	if entry != nil {
		*x = entry.x
		*y_2 = entry.y
		table.thisEntry = entry
		return 1
	} else {
		return 0
	}
	return 0
}

// table.c:366
func table_getNextEntry(table *TTable, x *float64, y_2 *float64) int32 {
	var (
		entry *TTableEntry
	)
	_ = entry
	entry = nil
	if table.file.mode == USE_FILE {
		return table_getNextFileEntry(table, x, y_2)
	}
	entry = table.thisEntry.next
	if entry != nil {
		*x = entry.x
		*y_2 = entry.y
		table.thisEntry = entry
		return 1
	} else {
		return 0
	}
	return 0
}

// table.c:395
func table_lookup(table *TTable, x float64) float64 {
	var (
		x1    float64
		y1    float64
		x2    float64
		y2    float64
		entry *TTableEntry
	)
	_, _, _, _, _ = x1, y1, x2, y2, entry
	x1 = 0
	y1 = 0
	x2 = 0
	y2 = 0
	entry = nil
	entry = table.firstEntry
	if entry == nil {
		return float64(0.0)
	}
	x1 = entry.x
	y1 = entry.y
	if x <= x1 {
		return y1
	}
	for entry.next != nil {
		entry = entry.next
		x2 = entry.x
		y2 = entry.y
		if x <= x2 {
			return table_interpolate(x, x1, y1, x2, y2)
		}
		x1 = x2
		y1 = y2
	}
	return y1
}

// table.c:430
func table_getSlope(table *TTable, x float64) float64 {
	var (
		x1    float64
		y1    float64
		x2    float64
		y2    float64
		dx    float64
		entry *TTableEntry
	)
	_, _, _, _, _, _ = x1, y1, x2, y2, dx, entry
	x1 = 0
	y1 = 0
	x2 = 0
	y2 = 0
	dx = 0
	entry = nil
	entry = table.firstEntry
	if entry == nil {
		return float64(0.0)
	}
	x1 = entry.x
	y1 = entry.y
	x2 = x1
	y2 = y1
	for entry.next != nil {
		entry = entry.next
		x2 = entry.x
		y2 = entry.y
		if x <= x2 {
			break
		}
		x1 = x2
		y1 = y2
	}
	dx = float64(x2 - x1)
	if dx == float64(0.0) {
		return float64(0.0)
	}
	return float64((float64(y2 - y1)) / dx)
}

// table.c:464
func table_lookupEx(table *TTable, x float64) float64 {
	var (
		x1    float64
		y1    float64
		x2    float64
		y2    float64
		s     float64
		entry *TTableEntry
	)
	_, _, _, _, _, _ = x1, y1, x2, y2, s, entry
	x1 = 0
	y1 = 0
	x2 = 0
	y2 = 0
	s = float64(0.0)
	entry = nil
	entry = table.firstEntry
	if entry == nil {
		return float64(0.0)
	}
	x1 = entry.x
	y1 = entry.y
	if x <= x1 {
		if x1 > float64(0.0) {
			return float64(float64(x/x1) * y1)
		} else {
			return y1
		}
	}
	for entry.next != nil {
		entry = entry.next
		x2 = entry.x
		y2 = entry.y
		if x2 != x1 {
			s = float64((float64(y2 - y1)) / (float64(x2 - x1)))
		}
		if x <= x2 {
			return table_interpolate(x, x1, y1, x2, y2)
		}
		x1 = x2
		y1 = y2
	}
	if s < float64(0.0) {
		s = float64(0.0)
	}
	return float64(y1 + float64(s*(float64(x-x1))))
}

// table.c:503
func table_intervalLookup(table *TTable, x float64) float64 {
	var (
		entry *TTableEntry
	)
	_ = entry
	entry = nil
	entry = table.firstEntry
	if entry == nil {
		return float64(0.0)
	}
	if x < entry.x {
		return entry.y
	}
	for entry.next != nil {
		entry = entry.next
		if x < entry.x {
			return entry.y
		}
	}
	return entry.y
}

// table.c:527
func table_inverseLookup(table *TTable, y_2 float64) float64 {
	var (
		x1    float64
		y1    float64
		x2    float64
		y2    float64
		entry *TTableEntry
	)
	_, _, _, _, _ = x1, y1, x2, y2, entry
	x1 = 0
	y1 = 0
	x2 = 0
	y2 = 0
	entry = nil
	entry = table.firstEntry
	if entry == nil {
		return float64(0.0)
	}
	x1 = entry.x
	y1 = entry.y
	if y_2 <= y1 {
		return x1
	}
	for entry.next != nil {
		entry = entry.next
		x2 = entry.x
		y2 = entry.y
		if y_2 <= y2 {
			return table_interpolate(y_2, y1, x1, y2, x2)
		}
		x1 = x2
		y1 = y2
	}
	return x1
}

// table.c:562
func table_getMaxY(table *TTable, x float64) float64 {
	var (
		ymax  float64
		entry *TTableEntry
	)
	_, _ = ymax, entry
	ymax = 0
	entry = nil
	entry = table.firstEntry
	ymax = entry.y
	for (x > entry.x) && (entry.next != nil) {
		entry = entry.next
		if entry.y < ymax {
			return ymax
		}
		ymax = entry.y
	}
	return float64(0.0)
}

// table.c:587
func table_getStorageVolume(table *TTable, x float64) float64 {
	var (
		a_2   float64
		a1    float64
		x1    float64
		v     float64
		dx    float64
		dy    float64
		s     float64
		entry *TTableEntry
	)
	_, _, _, _, _, _, _, _ = a_2, a1, x1, v, dx, dy, s, entry
	a_2 = 0
	a1 = 0
	x1 = 0
	v = 0
	dx = float64(0.0)
	dy = float64(0.0)
	s = 0
	entry = nil
	v = float64(0.0)
	entry = table.firstEntry
	if entry == nil {
		return float64(0.0)
	}
	x1 = entry.x
	a1 = entry.y
	if x <= x1 {
		if x1 < float64(1.0e-6) {
			return float64(0.0)
		}
		return float64(float64(float64((float64(a1/x1))*x)*x) / float64(2.0))
	}
	for entry.next != nil {
		entry = entry.next
		if entry.x >= x {
			a_2 = table_interpolate(x, x1, a1, entry.x, entry.y)
			return float64(v + float64(float64((float64(a1+a_2))/float64(2.0))*(float64(x-x1))))
		} else {
			dx = float64(entry.x - x1)
			dy = float64(entry.y - a1)
			v = float64(v + float64(float64((float64(a1+entry.y))/float64(2.0))*dx))
			x1 = entry.x
			a1 = entry.y
		}
	}
	if dx > float64(1.0e-6) {
		s = float64(dy / dx)
		a_2 = float64(a1 + float64(s*(float64(x-x1))))
		if a_2 < float64(0.0) {
			v = float64(v - float64(float64(float64(a1*a1)/s)/float64(2.0)))
		} else {
			v = float64(v + float64(float64((float64(a1+a_2))/float64(2.0))*(float64(x-x1))))
		}
	}
	return v
}

// table.c:651
func table_getStorageDepth(table *TTable, v float64) float64 {
	var (
		a1    float64
		a2    float64
		d1    float64
		d2    float64
		dd    float64
		da    float64
		v1    float64
		v2    float64
		s     float64
		entry *TTableEntry
	)
	_, _, _, _, _, _, _, _, _, _ = a1, a2, d1, d2, dd, da, v1, v2, s, entry
	a1 = 0
	a2 = 0
	d1 = 0
	d2 = 0
	dd = float64(0.0)
	da = float64(0.0)
	v1 = 0
	v2 = 0
	s = 0
	entry = nil
	if v == float64(0.0) {
		return float64(0.0)
	}
	entry = table.firstEntry
	if entry == nil {
		return float64(0.0)
	}
	d1 = entry.x
	a1 = entry.y
	v1 = float64(float64(a1*d1) / float64(2.0))
	if v <= v1 {
		if a1 > float64(0.0) {
			return M.Sqrt(float64(float64(float64(float64(2.0)*v)*d1) / a1))
		} else {
			return float64(0.0)
		}
	}
	for entry.next != nil {
		entry = entry.next
		d2 = entry.x
		a2 = entry.y
		dd = float64(d2 - d1)
		da = float64(a2 - a1)
		v2 = float64(v1 + float64(float64((float64(a1+a2))/float64(2.0))*dd))
		if v <= v2 {
			if dd <= float64(0.0) {
				return d1
			}
			if da == float64(0.0) {
				if M.Fabs(float64(v2-v1)) < float64(1.0e-6) {
					return d1
				} else {
					return float64(d1 + float64(float64(dd*(float64(v-v1)))/(float64(v2-v1))))
				}
			}
			if da < float64(0.0) {
				d1 = d2
				a1 = a2
				v1 = v2
			}
			s = float64(da / dd)
			return float64(d1 + float64((float64(M.Sqrt(float64(float64(a1*a1)+float64(float64(float64(2.0)*s)*(float64(v-v1)))))-a1))/s))
		}
		d1 = d2
		a1 = a2
		v1 = v2
	}
	if (dd == float64(0.0)) || (da == float64(0.0)) {
		if a1 > float64(0.0) {
			dd = float64((float64(v - v1)) / a1)
		} else {
			dd = float64(0.0)
		}
	} else {
		s = float64(da / dd)
		dd = float64((float64(M.Sqrt(float64(float64(a1*a1)+float64(float64(float64(2.0)*s)*(float64(v-v1))))) - a1)) / s)
		if dd < float64(0.0) {
			dd = float64(0.0)
		}
	}
	return float64(d1 + dd)
}

// table.c:730
func table_tseriesInit(table *TTable) {
	table_getFirstEntry(table, &table.x1, &table.y1)
	table.x2 = table.x1
	table.y2 = table.y1
	table_getNextEntry(table, &table.x2, &table.y2)
}

// table.c:745
func table_tseriesLookup(table *TTable, x float64, extend int32) float64 {
	if ((table.x1 <= x) && (table.x2 >= x)) && (table.x1 != table.x2) {
		return table_interpolate(x, table.x1, table.y1, table.x2, table.y2)
	}
	if (table.file.mode == USE_FILE) && (R.Feof(table.file.file) != 0) {
		if extend == 1 {
			return table.y1
		} else {
			return float64(0)
		}
	}
	if (table.x1 == table.x2) || (x < table.x1) {
		table_getFirstEntry(table, &table.x1, &table.y1)
		if x < table.x1 {
			if extend == 1 {
				return table.y1
			} else {
				return float64(0)
			}
		}
	}
	table.x1 = table.x2
	table.y1 = table.y2
	for table_getNextEntry(table, &table.x2, &table.y2) != 0 {
		if x <= table.x2 {
			return table_interpolate(x, table.x1, table.y1, table.x2, table.y2)
		}
		table.x1 = table.x2
		table.y1 = table.y2
	}
	if extend == 1 {
		return table.y1
	} else {
		return float64(0.0)
	}
	return 0
}

// table.c:808
func table_getNextFileEntry(table *TTable, x *float64, y_2 *float64) int32 {
	var (
		line R.CPtr
		code int32
	)
	_, _ = line, code
	line = R.Buf(1025)
	code = 0
	if table.file.file == nil {
		return 0
	}
	for (!(R.Feof(table.file.file) != 0)) && (!(R.Fgets(line, 1024, table.file.file)).IsNull()) {
		code = table_parseFileLine(line, table, x, y_2)
		if code < 0 {
			continue
		}
		return code
	}
	return 0
}

// table.c:833
func table_parseFileLine(line R.CPtr, table *TTable, x *float64, y_2 *float64) int32 {
	var (
		n_2  int32
		s1   R.CPtr
		s2   R.CPtr
		s3   R.CPtr
		tStr R.CPtr
		yStr R.CPtr
		yy   float64
		d    float64
		t    float64
	)
	_, _, _, _, _, _, _, _, _ = n_2, s1, s2, s3, tStr, yStr, yy, d, t
	n_2 = 0
	s1 = R.Buf(50)
	s2 = R.Buf(50)
	s3 = R.Buf(50)
	tStr = (R.CPtr{})
	yStr = (R.CPtr{})
	yy = 0
	d = 0
	t = 0
	n_2 = R.Sscanf(line, L650, s1, s2, s3)
	tStr = R.Strtok(line, L184)
	if (tStr.IsNull()) || (tStr.G(0) == 59) {
		return (-1)
	}
	if n_2 == 2 {
		d = table.lastDate
		tStr = s1
		yStr = s2
	} else if n_2 == 3 {
		if !(datetime_strToDate(s1, &d) != 0) {
			return 0
		}
		table.lastDate = d
		tStr = s2
		yStr = s3
	} else {
		return 0
	}
	if getDouble(tStr, &t) != 0 {
		t = float64(t / (float64(24.0)))
	} else if !(datetime_strToTime(tStr, &t) != 0) {
		return 0
	}
	if !(getDouble(yStr, &yy) != 0) {
		return 0
	}
	*x = float64(d + t)
	*y_2 = yy
	return 1
}
