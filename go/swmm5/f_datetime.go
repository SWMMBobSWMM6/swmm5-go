// Generated from EPA SWMM 5.2.4 (datetime.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
	M "swmm5go/glibm"
)

var MonthTxt = [13]R.CPtr{L695, L696, L697, L698, L699, L700, L701, L702, L703, L704, L705, L706}
var DaysPerMonth = [2][12]int32{[12]int32{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}, [12]int32{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}}
var DateDelta int32 = 693594
var SecsPerDay float64 = float64(86400.0)
var DateFormat int32

// datetime.c:51
func divMod(n_2 int32, d int32, result *int32, remainder *int32) {
	if d == 0 {
		*result = 0
		*remainder = 0
	} else {
		*result = (n_2 / d)
		*remainder = (n_2 - (d * (*result)))
	}
}

// datetime.c:74
func isLeapYear(year int32) int32 {
	if ((year % 4) == 0) && (((year % 100) != 0) || ((year % 400) == 0)) {
		return 1
	} else {
		return 0
	}
	return 0
}

// datetime.c:89
func datetime_findMonth(month R.CPtr) int32 {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 0; i < 12; i++ {
		if (((func() int32 {
			if ((month.G(0)) >= 97) && ((month.G(0)) <= 122) {
				return ((month.G(0)) & (^32))
			}
			return (month.G(0))
		}()) == MonthTxt[i].G(0)) && ((func() int32 {
			if ((month.G(1)) >= 97) && ((month.G(1)) <= 122) {
				return ((month.G(1)) & (^32))
			}
			return (month.G(1))
		}()) == MonthTxt[i].G(1))) && ((func() int32 {
			if ((month.G(2)) >= 97) && ((month.G(2)) <= 122) {
				return ((month.G(2)) & (^32))
			}
			return (month.G(2))
		}()) == MonthTxt[i].G(2)) {
			return (i + 1)
		}
	}
	return 0
}

// datetime.c:108
func datetime_encodeDate(year int32, month int32, day int32) float64 {
	var (
		i int32
		j int32
	)
	_, _ = i, j
	i = 0
	j = 0
	i = isLeapYear(year)
	if (((((year >= 1) && (year <= 9999)) && (month >= 1)) && (month <= 12)) && (day >= 1)) && (day <= DaysPerMonth[i][(month-1)]) {
		for j = 0; j < (month - 1); j++ {
			day += DaysPerMonth[i][j]
		}
		i = (year - 1)
		i = ((((((i * 365) + (i / 4)) - (i / 100)) + (i / 400)) + day) - DateDelta)
		return float64(i)
	} else {
		return float64((-DateDelta))
	}
	return 0
}

// datetime.c:136
func datetime_encodeTime(hour int32, minute int32, second int32) float64 {
	var (
		s int32
	)
	_ = s
	s = 0
	if ((hour >= 0) && (minute >= 0)) && (second >= 0) {
		s = (((hour * 3600) + (minute * 60)) + second)
		return float64(float64(s) / SecsPerDay)
	} else {
		return float64(0.0)
	}
	return 0
}

// datetime.c:158
func datetime_decodeDate(date float64, year *int32, month *int32, day *int32) {
	var (
		D1   int32
		D4   int32
		D100 int32
		D400 int32
		y_2  int32
		m    int32
		d    int32
		i    int32
		k    int32
		t    int32
	)
	_, _, _, _, _, _, _, _, _, _ = D1, D4, D100, D400, y_2, m, d, i, k, t
	D1 = 0
	D4 = 0
	D100 = 0
	D400 = 0
	y_2 = 0
	m = 0
	d = 0
	i = 0
	k = 0
	t = 0
	D1 = 365
	D4 = ((D1 * 4) + 1)
	D100 = ((D4 * 25) - 1)
	D400 = ((D100 * 4) + 1)
	t = (R.D2I((M.Floor(date))) + DateDelta)
	if t <= 0 {
		*year = 0
		*month = 1
		*day = 1
	} else {
		t--
		y_2 = 1
		for t >= D400 {
			t -= D400
			y_2 += 400
		}
		divMod(t, D100, &i, &d)
		if i == 4 {
			i--
			d += D100
		}
		y_2 += (i * 100)
		divMod(d, D4, &i, &d)
		y_2 += (i * 4)
		divMod(d, D1, &i, &d)
		if i == 4 {
			i--
			d += D1
		}
		y_2 += i
		k = isLeapYear(y_2)
		m = 1
		for {
			i = DaysPerMonth[k][(m - 1)]
			if d < i {
				break
			}
			d -= i
			m++
		}
		*year = y_2
		*month = m
		*day = (d + 1)
	}
}

// datetime.c:224
func datetime_decodeTime(time float64, h *int32, m *int32, s *int32) {
	var (
		secs    int32
		mins    int32
		fracDay float64
	)
	_, _, _ = secs, mins, fracDay
	secs = 0
	mins = 0
	fracDay = float64((float64(time - M.Floor(time))) * SecsPerDay)
	secs = R.D2I((M.Floor(float64(fracDay + float64(0.5)))))
	if secs >= 86400 {
		secs = 86399
	}
	divMod(secs, 60, &mins, s)
	divMod(mins, 60, h, m)
	if (*h) > 23 {
		*h = 0
	}
}

// datetime.c:245
func datetime_dateToStr(date float64, s R.CPtr) {
	var (
		y_2 int32
		m   int32
		d   int32
	)
	_, _, _ = y_2, m, d
	y_2 = 0
	m = 0
	d = 0
	datetime_decodeDate(date, &y_2, &m, &d)
	switch DateFormat {
	case 0:
		R.Snprintf(s, 12, L36, y_2, MonthTxt[(m-1)], d)
		break
	case 1:
		R.Snprintf(s, 12, L37, m, d, y_2)
		break
	default:
		R.Snprintf(s, 12, L38, d, MonthTxt[(m-1)], y_2)
	}
}

// datetime.c:270
func datetime_timeToStr(time float64, s R.CPtr) {
	var (
		hr    int32
		min_2 int32
		sec   int32
	)
	_, _, _ = hr, min_2, sec
	hr = 0
	min_2 = 0
	sec = 0
	datetime_decodeTime(time, &hr, &min_2, &sec)
	R.Snprintf(s, 9, L39, hr, min_2, sec)
}

// datetime.c:284
func datetime_strToDate(s R.CPtr, d *float64) int32 {
	var (
		yr    int32
		mon   int32
		day   int32
		n_2   int32
		month R.CPtr
		sep1  int32
		sep2  int32
	)
	_, _, _, _, _, _, _ = yr, mon, day, n_2, month, sep1, sep2
	yr = 0
	mon = 0
	day = 0
	n_2 = 0
	month = R.Buf(4)
	sep1 = 0
	sep2 = 0
	*d = float64((-DateDelta))
	if (!(R.Strchr(s, 45)).IsNull()) || (!(R.Strchr(s, 47)).IsNull()) {
		switch DateFormat {
		case 0:
			n_2 = R.Sscanf(s, L40, &yr, &sep1, &mon, &sep2, &day)
			if n_2 < 3 {
				mon = 0
				n_2 = R.Sscanf(s, L41, &yr, &sep1, month, &sep2, &day)
				if n_2 < 3 {
					return 0
				}
			}
			break
		case 2:
			n_2 = R.Sscanf(s, L40, &day, &sep1, &mon, &sep2, &yr)
			if n_2 < 3 {
				mon = 0
				n_2 = R.Sscanf(s, L41, &day, &sep1, month, &sep2, &yr)
				if n_2 < 3 {
					return 0
				}
			}
			break
		default:
			n_2 = R.Sscanf(s, L40, &mon, &sep1, &day, &sep2, &yr)
			if n_2 < 3 {
				mon = 0
				n_2 = R.Sscanf(s, L42, month, &sep1, &day, &sep2, &yr)
				if n_2 < 3 {
					return 0
				}
			}
		}
		if mon == 0 {
			mon = datetime_findMonth(month)
		}
		*d = datetime_encodeDate(yr, mon, day)
	}
	if (*d) == float64((-DateDelta)) {
		return 0
	} else {
		return 1
	}
	return 0
}

// datetime.c:338
func datetime_strToTime(s R.CPtr, t *float64) int32 {
	var (
		n_2    int32
		hr     int32
		min_2  int32
		sec    int32
		endptr R.CPtr
	)
	_, _, _, _, _ = n_2, hr, min_2, sec, endptr
	n_2 = 0
	hr = 0
	min_2 = 0
	sec = 0
	endptr = (R.CPtr{})
	*t = R.Strtod(s, &endptr)
	if endptr.G(0) == 0 {
		*t = float64((*t) / (float64(24.0)))
		return 1
	}
	*t = float64(0.0)
	n_2 = R.Sscanf(s, L43, &hr, &min_2, &sec)
	if n_2 == 0 {
		return 0
	}
	*t = datetime_encodeTime(hr, min_2, sec)
	if ((hr >= 0) && (min_2 >= 0)) && (sec >= 0) {
		return 1
	} else {
		return 0
	}
	return 0
}

// datetime.c:369
func datetime_setDateFormat(fmt_2 int32) {
	if (fmt_2 >= 0) && (fmt_2 <= 1) {
		DateFormat = fmt_2
	}
}

// datetime.c:381
func datetime_addSeconds(date1 float64, seconds float64) float64 {
	var (
		d float64
		h int32
		m int32
		s int32
	)
	_, _, _, _ = d, h, m, s
	d = M.Floor(date1)
	h = 0
	m = 0
	s = 0
	datetime_decodeTime(date1, &h, &m, &s)
	return float64(d + float64((float64(float64(float64(float64(float64(3600.0)*float64(h))+float64(float64(60.0)*float64(m)))+float64(s))+seconds))/SecsPerDay))
}

// datetime.c:397
func datetime_addDays(date1 float64, date2 float64) float64 {
	var (
		d1 float64
		d2 float64
		h1 int32
		m1 int32
		s1 int32
		h2 int32
		m2 int32
		s2 int32
	)
	_, _, _, _, _, _, _, _ = d1, d2, h1, m1, s1, h2, m2, s2
	d1 = M.Floor(date1)
	d2 = M.Floor(date2)
	h1 = 0
	m1 = 0
	s1 = 0
	h2 = 0
	m2 = 0
	s2 = 0
	datetime_decodeTime(date1, &h1, &m1, &s1)
	datetime_decodeTime(date2, &h2, &m2, &s2)
	return float64(float64(d1+d2) + datetime_encodeTime((h1+h2), (m1+m2), (s1+s2)))
}

// datetime.c:416
func datetime_timeDiff(date1 float64, date2 float64) int64 {
	var (
		d1   float64
		d2   float64
		h    int32
		m    int32
		s    int32
		s1   int64
		s2   int64
		secs int64
	)
	_, _, _, _, _, _, _, _ = d1, d2, h, m, s, s1, s2, secs
	d1 = M.Floor(date1)
	d2 = M.Floor(date2)
	h = 0
	m = 0
	s = 0
	s1 = 0
	s2 = 0
	secs = 0
	datetime_decodeTime(date1, &h, &m, &s)
	s1 = int64((((3600 * h) + (60 * m)) + s))
	datetime_decodeTime(date2, &h, &m, &s)
	s2 = int64((((3600 * h) + (60 * m)) + s))
	secs = int64(R.D2I((M.Floor(float64(float64((float64(d1-d2))*SecsPerDay) + float64(0.5))))))
	secs += (s1 - s2)
	return secs
}

// datetime.c:439
func datetime_monthOfYear(date float64) int32 {
	var (
		year  int32
		month int32
		day   int32
	)
	_, _, _ = year, month, day
	year = 0
	month = 0
	day = 0
	datetime_decodeDate(date, &year, &month, &day)
	return month
}

// datetime.c:453
func datetime_dayOfYear(date float64) int32 {
	var (
		year        int32
		month       int32
		day         int32
		startOfYear float64
	)
	_, _, _, _ = year, month, day, startOfYear
	year = 0
	month = 0
	day = 0
	startOfYear = 0
	datetime_decodeDate(date, &year, &month, &day)
	startOfYear = datetime_encodeDate(year, 1, 1)
	return (R.D2I((M.Floor(float64(date - startOfYear)))) + 1)
}

// datetime.c:469
func datetime_dayOfWeek(date float64) int32 {
	var (
		t int32
	)
	_ = t
	t = (R.D2I((M.Floor(date))) + DateDelta)
	return ((t % 7) + 1)
}

// datetime.c:482
func datetime_hourOfDay(date float64) int32 {
	var (
		hour  int32
		min_2 int32
		sec   int32
	)
	_, _, _ = hour, min_2, sec
	hour = 0
	min_2 = 0
	sec = 0
	datetime_decodeTime(date, &hour, &min_2, &sec)
	return hour
}

// datetime.c:496
func datetime_daysPerMonth(year int32, month int32) int32 {
	if (month < 1) || (month > 12) {
		return 0
	}
	return (DaysPerMonth[isLeapYear(year)])[(month - 1)]
}

// datetime.c:510
func datetime_getTimeStamp(fmt_2 int32, aDate float64, stampSize int32, timeStamp R.CPtr) {
	var (
		dateStr       R.CPtr
		timeStr       R.CPtr
		oldDateFormat int32
	)
	_, _, _ = dateStr, timeStr, oldDateFormat
	dateStr = R.Buf(12)
	timeStr = R.Buf(9)
	oldDateFormat = DateFormat
	if stampSize < 21 {
		return
	}
	datetime_setDateFormat(fmt_2)
	datetime_dateToStr(aDate, dateStr)
	DateFormat = oldDateFormat
	datetime_timeToStr(aDate, timeStr)
	R.Snprintf(timeStamp, uint64(stampSize), L44, dateStr, timeStr)
}
