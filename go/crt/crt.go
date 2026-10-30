// Package crt is the C runtime layer of the Go port of EPA SWMM 5.2.4 (imported by the generated code
// as R). It reproduces the glibc and x86-64 behaviour SWMM's output depends on: NUL-terminated byte
// strings (CPtr), x86's float-to-integer conversion, printf formatting with glibc's rounding,
// strtod/sscanf parsing, and stdio files with C's positioning and end-of-file rules. The math functions
// are in package glibm (bit-exact ports of glibc 2.39). Ported from the Java port's R.java.
// Written by Claude (Anthropic); see AI-USE.md.
package crt

import (
	"math"
	"unsafe"
)

// Number is every C arithmetic type the port uses.
type Number interface {
	~int32 | ~uint32 | ~int64 | ~uint64 | ~float32 | ~float64
}

// Integer is every Go integer type an index or count can arrive as.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

var (
	QNaN    = math.Float64frombits(0x7ff8000000000000) // C's NAN (math.NaN() is 0x7ff8000000000001)
	Inf     = math.Inf(1)
	NegZero = math.Copysign(0, -1) // Go constants have no negative zero
)

// ------------------------------------------------------------ C char*: a byte slice plus an offset

type CPtr struct {
	B []byte
	O int32
}

func (p CPtr) G(i int32) int32          { return int32(int8(p.B[p.O+i])) } // signed char, as x86 C
func (p CPtr) S(i int32, v int32) int32 { p.B[p.O+i] = byte(v); return v }
func (p CPtr) Add(n int32) CPtr         { return CPtr{p.B, p.O + n} }
func (p CPtr) IsNull() bool             { return p.B == nil }

// CP views a byte array (a struct's char field) as a C string.
func CP(b []byte) CPtr { return CPtr{b, 0} }

func Peq(a, b CPtr) bool {
	if a.B == nil || b.B == nil {
		return a.B == nil && b.B == nil
	}
	return unsafe.SliceData(a.B) == unsafe.SliceData(b.B) && a.O == b.O
}

func Lit(s string) CPtr {
	b := make([]byte, len(s)+1)
	copy(b, s)
	return CPtr{b, 0}
}

func Buf[I Integer](n I) CPtr {
	m := int(n)
	if m < 1 {
		m = 1
	}
	return CPtr{make([]byte, m), 0}
}

func BufFrom(src CPtr, n int) CPtr {
	p := Buf(n)
	m := int(strlen(src))
	if m > n {
		m = n
	}
	copy(p.B, src.B[src.O:int(src.O)+m])
	return p
}

func BufOf(n int, vals ...int32) CPtr {
	p := Buf(n)
	for i, v := range vals {
		p.B[i] = byte(v)
	}
	return p
}

func Bufs(n, m int) []CPtr {
	a := make([]CPtr, n)
	for i := range a {
		a[i] = Buf(m)
	}
	return a
}

// Str is the Go string of a C string.
func Str(p CPtr) string { return string(p.B[p.O : p.O+strlen(p)]) }

func ReallocC[I Integer](p CPtr, n I) CPtr {
	q := Buf(n)
	if p.B != nil {
		copy(q.B, p.B[p.O:])
	}
	return q
}

func Realloc[T any, I Integer](a []T, n I) []T {
	r := make([]T, int(n))
	copy(r, a)
	return r
}

func Free(x any) {}

func MemsetC[I Integer](p CPtr, v int32, n I) {
	b := p.B[p.O : int(p.O)+int(n)]
	for i := range b {
		b[i] = byte(v)
	}
}

// Memset zeroes the first n bytes' worth of elements (SWMM only ever memsets numbers to 0).
func Memset[T Number, I Integer](a []T, v int32, n I) {
	var z T
	k := int(uint64(n) / uint64(unsafe.Sizeof(z)))
	if k > len(a) {
		k = len(a)
	}
	for i := 0; i < k; i++ {
		a[i] = 0
	}
}

// ------------------------------------------------------------ arithmetic with C semantics

// D2I is (int)x on x86-64: NaN or out of range gives INT_MIN.
func D2I(x float64) int32 {
	if !(x > -2147483649.0 && x < 2147483648.0) {
		return math.MinInt32
	}
	return int32(x)
}

func D2L(x float64) int64 {
	if !(x > -9223372036854775809.0 && x < 9223372036854775808.0) {
		return math.MinInt64
	}
	return int64(x)
}

func B2I(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// If is ?: for operands that are safe to evaluate eagerly.
func If[T any](c bool, a, b T) T {
	if c {
		return a
	}
	return b
}

// Set is C's assignment used as a value.
func Set[T any](p *T, v T) T { *p = v; return v }

func PreInc[T Number](p *T) T  { *p++; return *p }
func PreDec[T Number](p *T) T  { *p--; return *p }
func PostInc[T Number](p *T) T { v := *p; *p++; return v }
func PostDec[T Number](p *T) T { v := *p; *p--; return v }

func CPre(p *CPtr, d int32) CPtr  { *p = p.Add(d); return *p }
func CPost(p *CPtr, d int32) CPtr { v := *p; *p = p.Add(d); return v }

// One is &x seen as a one-element array (C passes &x where the callee indexes x[0]).
func One[T any](p *T) []T {
	if p == nil {
		return nil
	}
	return unsafe.Slice(p, 1)
}

// First is an array seen as a pointer to its first element.
func First[T any](a []T) *T {
	if len(a) == 0 {
		return nil
	}
	return &a[0]
}

// SameAddr is C pointer equality for two arrays.
func SameAddr[T any](a, b []T) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return unsafe.SliceData(a) == unsafe.SliceData(b)
}

// ------------------------------------------------------------ strings

func strlen(p CPtr) int32 {
	b := p.B
	e := p.O
	for b[e] != 0 {
		e++
	}
	return e - p.O
}

func Strlen(p CPtr) uint64 { return uint64(strlen(p)) }

func Strcpy(d, s CPtr) CPtr {
	n := strlen(s) + 1
	copy(d.B[d.O:d.O+n], s.B[s.O:s.O+n])
	return d
}

func Strcat(d, s CPtr) CPtr {
	Strcpy(d.Add(strlen(d)), s)
	return d
}

func Strcmp(a, b CPtr) int32 {
	i, j := a.O, b.O
	for {
		c, d := int32(a.B[i]), int32(b.B[j])
		if c != d {
			return c - d
		}
		if c == 0 {
			return 0
		}
		i++
		j++
	}
}

func Strchr(s CPtr, c int32) CPtr {
	c &= 255
	for i := s.O; ; i++ {
		if int32(s.B[i]) == c {
			return CPtr{s.B, i}
		}
		if s.B[i] == 0 {
			return CPtr{}
		}
	}
}

func Strrchr(s CPtr, c int32) CPtr {
	c &= 255
	e := s.O + strlen(s)
	if c == 0 {
		return CPtr{s.B, e}
	}
	for i := e - 1; i >= s.O; i-- {
		if int32(s.B[i]) == c {
			return CPtr{s.B, i}
		}
	}
	return CPtr{}
}

func Strstr(h, n CPtr) CPtr {
	hl, nl := strlen(h), strlen(n)
outer:
	for i := int32(0); i+nl <= hl; i++ {
		for k := int32(0); k < nl; k++ {
			if h.B[h.O+i+k] != n.B[n.O+k] {
				continue outer
			}
		}
		return CPtr{h.B, h.O + i}
	}
	return CPtr{}
}

func Strcspn(s, rej CPtr) uint64 {
	n, m := strlen(s), strlen(rej)
	for i := int32(0); i < n; i++ {
		for k := int32(0); k < m; k++ {
			if s.B[s.O+i] == rej.B[rej.O+k] {
				return uint64(i)
			}
		}
	}
	return uint64(n)
}

var tokNext CPtr

func inSet(c byte, set CPtr) bool {
	for k := set.O; set.B[k] != 0; k++ {
		if set.B[k] == c {
			return true
		}
	}
	return false
}

func Strtok(s, delim CPtr) CPtr {
	if !s.IsNull() {
		tokNext = s
	}
	p := tokNext
	if p.IsNull() {
		return CPtr{}
	}
	b := p.B
	i := p.O
	for b[i] != 0 && inSet(b[i], delim) {
		i++
	}
	if b[i] == 0 {
		tokNext = CPtr{}
		return CPtr{}
	}
	start := i
	for b[i] != 0 && !inSet(b[i], delim) {
		i++
	}
	if b[i] != 0 {
		b[i] = 0
		tokNext = CPtr{b, i + 1}
	} else {
		tokNext = CPtr{}
	}
	return CPtr{b, start}
}

func Toupper(c int32) int32 {
	if c >= 97 && c <= 122 {
		return c - 32
	}
	return c
}
