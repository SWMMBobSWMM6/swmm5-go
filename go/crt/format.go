// format.go - printf, sscanf, strtod, atof and atoi with glibc's results. strconv's 'f' and 'e'
// conversions are exact and round half to even on the binary value, as glibc does in its default
// rounding mode; %g is built from them the way C99 specifies. Written by Claude (Anthropic); see AI-USE.md.
package crt

import (
	"math"
	"strconv"
	"strings"
)

func isSpace(c int32) bool { return c == 32 || (c >= 9 && c <= 13) }

func isHex(c int32) bool {
	return (c >= 48 && c <= 57) || (c >= 65 && c <= 70) || (c >= 97 && c <= 102)
}

func low(c int32) int32 {
	if c >= 65 && c <= 90 {
		return c | 0x20
	}
	return c
}

var (
	defaultNaN = math.Float64frombits(0xfff8000000000000)
	posNaN     = math.Float64frombits(0x7ff8000000000000)
)

// scanDouble parses a C floating number at b[i:] (at most maxw chars). It returns the value, the end
// index (-1: no number) and the text parsed (empty for inf/nan).
func scanDouble(b []byte, i int, maxw int, skipSpace bool) (float64, int, string) {
	n := len(b)
	if skipSpace {
		for i < n && isSpace(int32(b[i])) {
			i++
		}
	}
	lim := n
	if maxw >= 0 && i+maxw < n {
		lim = i + maxw
	}
	g := func(k int) int32 {
		if k < lim && k < n {
			return int32(b[k])
		}
		return 0
	}
	neg := false
	if g(i) == 43 || g(i) == 45 {
		neg = g(i) == 45
		i++
	}
	if low(g(i)) == 105 && low(g(i+1)) == 110 && low(g(i+2)) == 102 {
		i += 3
		if low(g(i)) == 105 && low(g(i+1)) == 110 && low(g(i+2)) == 105 && low(g(i+3)) == 116 && low(g(i+4)) == 121 {
			i += 5
		}
		if neg {
			return math.Inf(-1), i, ""
		}
		return math.Inf(1), i, ""
	}
	if low(g(i)) == 110 && low(g(i+1)) == 97 && low(g(i+2)) == 110 {
		i += 3
		if g(i) == 40 {
			k := i + 1
			for {
				c := g(k)
				if (c >= 48 && c <= 57) || (c >= 65 && c <= 90) || (c >= 97 && c <= 122) || c == 95 {
					k++
					continue
				}
				break
			}
			if g(k) == 41 {
				// glibc reads the n-char-sequence with strtoull(seq, 0) and keeps it as the payload
				if pl, err := strconv.ParseUint(string(b[i+1:k]), 0, 64); err == nil || i+1 == k {
					if i+1 == k {
						pl = 0
					}
					v := math.Float64frombits(0x7ff8000000000000 | pl&(1<<51-1))
					if neg {
						v = -v
					}
					return v, k + 1, ""
				}
				i = k + 1
			}
		}
		if neg {
			return defaultNaN, i, ""
		}
		return posNaN, i, ""
	}
	if g(i) == 48 && low(g(i+1)) == 120 { // hexadecimal float (C99), as glibc's strtod reads it
		j, nh := i+2, 0
		for isHex(g(j)) {
			j++
			nh++
		}
		if g(j) == 46 {
			j++
			for isHex(g(j)) {
				j++
				nh++
			}
		}
		if nh > 0 {
			mant := string(b[i+2 : j])
			exp := "0"
			if low(g(j)) == 112 {
				k := j + 1
				if g(k) == 43 || g(k) == 45 {
					k++
				}
				if g(k) >= 48 && g(k) <= 57 {
					for g(k) >= 48 && g(k) <= 57 {
						k++
					}
					exp = string(b[j+1 : k])
					j = k
				}
			}
			if strings.HasPrefix(mant, ".") {
				mant = "0" + mant
			}
			text := "0x" + mant + "p" + exp
			if neg {
				text = "-" + text
			}
			v, err := strconv.ParseFloat(text, 64)
			if err != nil {
				if ne, ok := err.(*strconv.NumError); !ok || ne.Err != strconv.ErrRange {
					panic("strtod: " + text)
				}
			}
			return v, j, text
		}
	}
	ms, nd := i, 0
	for g(i) >= 48 && g(i) <= 57 {
		i++
		nd++
	}
	if g(i) == 46 {
		i++
		for g(i) >= 48 && g(i) <= 57 {
			i++
			nd++
		}
	}
	if nd == 0 {
		return 0, -1, ""
	}
	if low(g(i)) == 101 {
		k := i + 1
		if g(k) == 43 || g(k) == 45 {
			k++
		}
		if g(k) >= 48 && g(k) <= 57 {
			for g(k) >= 48 && g(k) <= 57 {
				k++
			}
			i = k
		}
	}
	text := string(b[ms:i])
	if neg {
		text = "-" + text
	}
	return parse(text, 64), i, text
}

func parse(text string, bits int) float64 {
	if strings.Contains(text, "x") {
		v, _ := strconv.ParseFloat(text, bits)
		return v
	}
	if strings.HasSuffix(text, ".") {
		text += "0"
	}
	if strings.HasPrefix(text, ".") || strings.HasPrefix(text, "-.") {
		text = strings.Replace(text, ".", "0.", 1)
	}
	v, err := strconv.ParseFloat(text, bits) // correctly rounded, as glibc's strtod
	if err != nil {
		if ne, ok := err.(*strconv.NumError); !ok || ne.Err != strconv.ErrRange {
			panic("strtod: " + text)
		}
	}
	return v
}

func Strtod(s CPtr, end *CPtr) float64 {
	v, e, _ := scanDouble(s.B, int(s.O), -1, true)
	if end != nil {
		if e < 0 {
			*end = s
		} else {
			*end = CPtr{s.B, int32(e)}
		}
	}
	if e < 0 {
		return 0
	}
	return v
}

func Atof(s CPtr) float64 { return Strtod(s, nil) }

// Atoi is strtol's long, then (int): it keeps the low 32 bits.
func Atoi(s CPtr) int32 {
	b := s.B
	i := s.O
	for b[i] != 0 && isSpace(int32(b[i])) {
		i++
	}
	neg := false
	if b[i] == 43 || b[i] == 45 {
		neg = b[i] == 45
		i++
	}
	var v int64
	over := false
	for b[i] >= 48 && b[i] <= 57 {
		d := int64(b[i] - 48)
		if v > (math.MaxInt64-d)/10 {
			over = true
		} else {
			v = v*10 + d
		}
		i++
	}
	if over {
		if neg {
			v = math.MinInt64
		} else {
			v = math.MaxInt64
		}
	} else if neg {
		v = -v
	}
	return int32(v)
}

// ------------------------------------------------------------ sscanf

func ch(b []byte, j int) int32 {
	if j < len(b) {
		return int32(b[j])
	}
	return 0
}

func Sscanf(s CPtr, fmtp CPtr, args ...any) int32 {
	b := s.B
	i := int(s.O)
	f := s2b(fmtp)
	k, ai, count := 0, 0, int32(0)
	anyInput := false
	n, nf := len(b), len(f)
	for k < nf {
		fc := int32(f[k])
		if isSpace(fc) {
			for k < nf && isSpace(int32(f[k])) {
				k++
			}
			for ch(b, i) != 0 && isSpace(ch(b, i)) {
				i++
			}
			continue
		}
		if fc != 37 {
			if ch(b, i) == 0 {
				if anyInput || count > 0 {
					return count
				}
				return -1
			}
			if ch(b, i) != fc {
				return count
			}
			i++
			k++
			continue
		}
		k++
		if f[k] == 37 {
			for ch(b, i) != 0 && isSpace(ch(b, i)) {
				i++
			}
			if ch(b, i) != 37 {
				return count
			}
			i++
			k++
			continue
		}
		suppress := false
		if f[k] == 42 {
			suppress = true
			k++
		}
		width := 0
		for k < nf && f[k] >= 48 && f[k] <= 57 {
			width = width*10 + int(f[k]-48)
			k++
		}
		for k < nf && (f[k] == 108 || f[k] == 104 || f[k] == 76) {
			k++
		}
		conv := int32(f[k])
		k++
		if conv != 99 && conv != 91 && conv != 110 {
			for ch(b, i) != 0 && isSpace(ch(b, i)) {
				i++
			}
		}
		if ch(b, i) == 0 && conv != 110 {
			if count == 0 && !anyInput {
				return -1
			}
			return count
		}
		switch conv {
		case 100, 105, 117:
			j := i
			lim := n
			if width != 0 {
				lim = i + width
			}
			neg := false
			if j < lim && (ch(b, j) == 43 || ch(b, j) == 45) {
				neg = ch(b, j) == 45
				j++
			}
			ds := j
			var v int64
			over := false
			for j < lim && ch(b, j) >= 48 && ch(b, j) <= 57 {
				d := int64(ch(b, j) - 48)
				if v > (math.MaxInt64-d)/10 {
					over = true
				} else {
					v = v*10 + d
				}
				j++
			}
			if j == ds {
				return count
			}
			i = j
			anyInput = true
			if over {
				if neg {
					v = math.MinInt64
				} else {
					v = math.MaxInt64
				}
			} else if neg {
				v = -v
			}
			if !suppress {
				switch a := args[ai].(type) {
				case *int32:
					*a = int32(v)
				case *int64:
					*a = v
				case *uint32:
					*a = uint32(v)
				default:
					panic("sscanf %d target")
				}
				ai++
				count++
			}
		case 102, 101, 103:
			w := -1
			if width != 0 {
				w = width
			}
			v, e, text := scanDouble(b, i, w, false)
			if e < 0 {
				return count
			}
			i = e
			anyInput = true
			if !suppress {
				switch a := args[ai].(type) {
				case *float64:
					*a = v
				case *float32:
					if text != "" {
						*a = float32(parse(text, 32))
					} else {
						*a = float32(v)
					}
				default:
					panic("sscanf %f target")
				}
				ai++
				count++
			}
		case 115:
			j := i
			w := 1 << 30
			if width != 0 {
				w = width
			}
			m := 0
			for ch(b, j) != 0 && !isSpace(ch(b, j)) && m < w {
				j++
				m++
			}
			if !suppress {
				d := args[ai].(CPtr)
				ai++
				copy(d.B[d.O:int(d.O)+m], b[i:i+m])
				d.B[int(d.O)+m] = 0
				count++
			}
			i = j
			anyInput = true
		case 99:
			cnt := 1
			if width != 0 {
				cnt = width
			}
			if ch(b, i) == 0 {
				if count == 0 && !anyInput {
					return -1
				}
				return count
			}
			var a any
			if !suppress {
				a = args[ai]
				ai++
			}
			t := 0
			for t < cnt && ch(b, i) != 0 {
				if !suppress {
					switch x := a.(type) {
					case CPtr:
						x.B[int(x.O)+t] = b[i]
					case *int32:
						*x = int32(int8(b[i]))
					}
				}
				i++
				t++
			}
			anyInput = true
			if !suppress {
				count++
			}
		default:
			panic("sscanf %" + string(rune(conv)))
		}
	}
	return count
}

// s2b is the bytes of a C string (not copied).
func s2b(p CPtr) []byte { return p.B[p.O : p.O+strlen(p)] }

// ------------------------------------------------------------ printf

func fixedDigits(a float64, prec int, alt bool) string {
	s := strconv.FormatFloat(a, 'f', prec, 64)
	if prec == 0 && alt {
		s += "."
	}
	return s
}

// expRaw is the digits of a rounded to prec+1 significant digits, and the decimal exponent.
func expRaw(a float64, prec int) (string, int) {
	if a == 0 {
		return strings.Repeat("0", prec+1), 0
	}
	s := strconv.FormatFloat(a, 'e', prec, 64) // d.ddde±XX
	ep := strings.IndexByte(s, 'e')
	x, _ := strconv.Atoi(s[ep+1:])
	return strings.Replace(s[:ep], ".", "", 1), x
}

func expDigits(a float64, prec int, alt, upper bool) string {
	d, x := expRaw(a, prec)
	mant := d[:1]
	if prec > 0 {
		mant += "." + d[1:]
	} else if alt {
		mant += "."
	}
	e := "e"
	if upper {
		e = "E"
	}
	sign := "+"
	if x < 0 {
		sign = "-"
		x = -x
	}
	xs := strconv.Itoa(x)
	if x < 10 {
		xs = "0" + xs
	}
	return mant + e + sign + xs
}

func formatFloat(a float64, conv byte, prec int, alt bool) string {
	lc := conv | 0x20
	if lc == 'f' {
		return fixedDigits(a, prec, alt)
	}
	if lc == 'e' {
		return expDigits(a, prec, alt, conv == 'E')
	}
	P := prec
	if P == 0 {
		P = 1
	}
	X := 0
	if a != 0 {
		_, X = expRaw(a, P-1)
	}
	var s string
	if P > X && X >= -4 {
		s = fixedDigits(a, P-1-X, alt)
	} else {
		s = expDigits(a, P-1, alt, conv == 'G')
	}
	if !alt {
		ep := strings.IndexAny(s, "eE")
		mant, ex := s, ""
		if ep >= 0 {
			mant, ex = s[:ep], s[ep:]
		}
		if strings.IndexByte(mant, '.') >= 0 {
			mant = strings.TrimRight(mant, "0")
			mant = strings.TrimSuffix(mant, ".")
		}
		s = mant + ex
	}
	return s
}

func toLong(a any) (int64, bool) { // value, and whether it is a 32-bit C int
	switch v := a.(type) {
	case int32:
		return int64(v), true
	case int:
		return int64(v), true
	case int64:
		return v, false
	case uint32:
		return int64(v), false
	case uint64:
		return int64(v), false
	case float64:
		return int64(v), false
	case float32:
		return int64(v), false
	}
	panic("printf integer arg")
}

func toDouble(a any) float64 {
	switch v := a.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int32:
		return float64(v)
	case int:
		return float64(v)
	case int64:
		return float64(v)
	}
	panic("printf float arg")
}

func format(fmtp CPtr, args []any) []byte {
	f := fmtp.B
	k, ai := int(fmtp.O), 0
	o := make([]byte, 0, 128)
	for f[k] != 0 {
		if f[k] != 37 {
			o = append(o, f[k])
			k++
			continue
		}
		k++
		if f[k] == 37 {
			o = append(o, '%')
			k++
			continue
		}
		left, plus, space, alt, zero := false, false, false, false, false
	flags:
		for ; ; k++ {
			switch f[k] {
			case '-':
				left = true
			case '+':
				plus = true
			case ' ':
				space = true
			case '#':
				alt = true
			case '0':
				zero = true
			default:
				break flags
			}
		}
		width, prec := 0, -1
		if f[k] == '*' {
			w, _ := toLong(args[ai])
			ai++
			width = int(w)
			if width < 0 {
				left = true
				width = -width
			}
			k++
		} else {
			for f[k] >= '0' && f[k] <= '9' {
				width = width*10 + int(f[k]-'0')
				k++
			}
		}
		if f[k] == '.' {
			k++
			prec = 0
			if f[k] == '*' {
				p, _ := toLong(args[ai])
				ai++
				prec = int(p)
				if prec < 0 {
					prec = -1
				}
				k++
			} else {
				for f[k] >= '0' && f[k] <= '9' {
					prec = prec*10 + int(f[k]-'0')
					k++
				}
			}
		}
		for f[k] == 'l' || f[k] == 'h' || f[k] == 'L' || f[k] == 'z' || f[k] == 'j' || f[k] == 't' || f[k] == 'q' {
			k++
		}
		conv := f[k]
		k++
		var body, sign string
		numeric := true
		switch conv {
		case 'd', 'i':
			v, _ := toLong(args[ai])
			ai++
			if v < 0 {
				sign = "-"
			} else if plus {
				sign = "+"
			} else if space {
				sign = " "
			}
			if v == math.MinInt64 {
				body = "9223372036854775808"
			} else if v < 0 {
				body = strconv.FormatInt(-v, 10)
			} else {
				body = strconv.FormatInt(v, 10)
			}
			if prec >= 0 {
				zero = false
				if prec == 0 && v == 0 {
					body = ""
				}
				for len(body) < prec {
					body = "0" + body
				}
			}
		case 'u', 'x', 'X', 'o':
			v, isInt := toLong(args[ai])
			ai++
			u := uint64(v)
			if isInt {
				u &= 0xFFFFFFFF
			}
			switch conv {
			case 'u':
				body = strconv.FormatUint(u, 10)
			case 'o':
				body = strconv.FormatUint(u, 8)
			case 'x':
				body = strconv.FormatUint(u, 16)
			default:
				body = strings.ToUpper(strconv.FormatUint(u, 16))
			}
			if prec >= 0 {
				zero = false
				if prec == 0 && u == 0 {
					body = ""
				}
				for len(body) < prec {
					body = "0" + body
				}
			}
		case 'c':
			numeric = false
			v, _ := toLong(args[ai])
			ai++
			body = string([]byte{byte(v)})
		case 's':
			numeric = false
			p := args[ai].(CPtr)
			ai++
			if p.IsNull() {
				body = "(null)"
			} else {
				body = string(s2b(p))
			}
			if prec >= 0 && len(body) > prec {
				body = body[:prec]
			}
		case 'f', 'F', 'e', 'E', 'g', 'G':
			v := toDouble(args[ai])
			ai++
			if prec < 0 {
				prec = 6
			}
			if math.Signbit(v) {
				sign = "-"
			} else if plus {
				sign = "+"
			} else if space {
				sign = " "
			}
			if math.IsNaN(v) || math.IsInf(v, 0) {
				body = "inf"
				if math.IsNaN(v) {
					body = "nan"
				}
				if conv == 'F' || conv == 'E' || conv == 'G' {
					body = strings.ToUpper(body)
				}
				zero = false
			} else {
				body = formatFloat(math.Abs(v), conv, prec, alt)
			}
		default:
			panic("printf %" + string(rune(conv)))
		}
		n := len(sign) + len(body)
		switch {
		case n >= width:
			o = append(append(o, sign...), body...)
		case left:
			o = append(append(o, sign...), body...)
			o = append(o, strings.Repeat(" ", width-n)...)
		case zero && numeric:
			o = append(o, sign...)
			o = append(o, strings.Repeat("0", width-n)...)
			o = append(o, body...)
		default:
			o = append(o, strings.Repeat(" ", width-n)...)
			o = append(append(o, sign...), body...)
		}
	}
	return o
}

func Fprintf(f *CFile, fmtp CPtr, args ...any) int32 {
	s := format(fmtp, args)
	if f == nil {
		return -1
	}
	f.write(s)
	return int32(len(s))
}

func Snprintf[I Integer](b CPtr, n I, fmtp CPtr, args ...any) int32 {
	s := format(fmtp, args)
	if int(n) > 0 {
		m := len(s)
		if m > int(n)-1 {
			m = int(n) - 1
		}
		copy(b.B[b.O:], s[:m])
		b.B[int(b.O)+m] = 0
	}
	return int32(len(s))
}

// Sprintf is printf into a Go string (for the hand-written runner).
func Sprintf(fmtp string, args ...any) string { return string(format(Lit(fmtp), args)) }
