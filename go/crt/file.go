// file.go - C stdio over os.File: positioned reads and writes with a write buffer and a read buffer,
// C's end-of-file flag, fopen modes, binary fread/fwrite of single values and arrays (little-endian;
// a short read leaves the upper bytes of the old value, as C does), time and ctime.
// Written by Claude (Anthropic); see AI-USE.md.
package crt

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unsafe"
)

const bufsz = 1 << 16

type CFile struct {
	f      *os.File
	std    bool
	pos    int64
	eof    bool
	wbuf   []byte
	wpos   int64
	rbuf   []byte
	rstart int64
	rlen   int
}

var Stdout = &CFile{f: os.Stdout, std: true}

func (f *CFile) flush() {
	if len(f.wbuf) == 0 {
		return
	}
	var err error
	if f.std {
		_, err = f.f.Write(f.wbuf)
	} else {
		_, err = f.f.WriteAt(f.wbuf, f.wpos)
	}
	if err != nil {
		panic(err)
	}
	f.wbuf = f.wbuf[:0]
}

func (f *CFile) write(src []byte) {
	n := len(src)
	f.rlen = 0
	if f.wbuf == nil {
		f.wbuf = make([]byte, 0, bufsz)
	}
	if len(f.wbuf) > 0 && f.pos != f.wpos+int64(len(f.wbuf)) {
		f.flush()
	}
	if len(f.wbuf) == 0 {
		f.wpos = f.pos
	}
	if len(f.wbuf)+n > bufsz {
		f.flush()
		if n > bufsz {
			var err error
			if f.std {
				_, err = f.f.Write(src)
			} else {
				_, err = f.f.WriteAt(src, f.pos)
			}
			if err != nil {
				panic(err)
			}
			f.pos += int64(n)
			return
		}
		f.wpos = f.pos
	}
	f.wbuf = append(f.wbuf, src...)
	f.pos += int64(n)
}

func (f *CFile) read(dst []byte) int {
	f.flush()
	n := len(dst)
	got := 0
	for got < n {
		if f.rlen > 0 && f.pos >= f.rstart && f.pos < f.rstart+int64(f.rlen) {
			a := int(f.pos - f.rstart)
			m := copy(dst[got:], f.rbuf[a:f.rlen])
			got += m
			f.pos += int64(m)
			continue
		}
		if f.rbuf == nil {
			f.rbuf = make([]byte, bufsz)
		}
		f.rstart = f.pos
		r, _ := f.f.ReadAt(f.rbuf, f.pos)
		f.rlen = r
		if r == 0 {
			break
		}
	}
	return got
}

func (f *CFile) getc() int32 {
	if len(f.wbuf) == 0 && f.rlen > 0 && f.pos >= f.rstart && f.pos < f.rstart+int64(f.rlen) {
		c := f.rbuf[f.pos-f.rstart]
		f.pos++
		return int32(c)
	}
	var t [1]byte
	if f.read(t[:]) == 1 {
		return int32(t[0])
	}
	return -1
}

func (f *CFile) size() int64 {
	f.flush()
	st, err := f.f.Stat()
	if err != nil {
		panic(err)
	}
	return st.Size()
}

func Fopen(name, mode CPtr) *CFile {
	fn := Str(name)
	m := strings.NewReplacer("t", "", "b", "").Replace(Str(mode))
	if strings.HasPrefix(m, "r") {
		st, err := os.Stat(fn)
		if err != nil || !st.Mode().IsRegular() {
			return nil
		}
	}
	flag := os.O_RDWR | os.O_CREATE
	if m == "r" {
		flag = os.O_RDONLY
	}
	fh, err := os.OpenFile(fn, flag, 0644)
	if err != nil {
		return nil
	}
	f := &CFile{f: fh}
	if strings.HasPrefix(m, "w") {
		if fh.Truncate(0) != nil {
			fh.Close()
			return nil
		}
	}
	if strings.HasPrefix(m, "a") {
		f.pos = f.size()
	}
	return f
}

func Fclose(f *CFile) int32 {
	if f == nil {
		return -1
	}
	f.flush()
	if !f.std {
		if f.f.Close() != nil {
			return -1
		}
	}
	return 0
}

func Fflush(f *CFile) int32 {
	if f != nil {
		f.flush()
	}
	return 0
}

func Feof(f *CFile) int32 {
	if f.eof {
		return 1
	}
	return 0
}

func Rewind(f *CFile) {
	f.flush()
	f.pos = 0
	f.eof = false
}

func Fseek(f *CFile, off int64, whence int32) int32 {
	f.flush()
	p := off
	switch whence {
	case 1:
		p = f.pos + off
	case 2:
		p = f.size() + off
	}
	if p < 0 {
		return -1
	}
	f.pos = p
	f.eof = false
	return 0
}

func Ftell(f *CFile) int64 { return f.pos }

func Fgets(b CPtr, n int32, f *CFile) CPtr {
	m := int32(0)
	for m < n-1 {
		c := f.getc()
		if c < 0 {
			f.eof = true
			break
		}
		b.B[b.O+m] = byte(c)
		m++
		if c == 10 {
			break
		}
	}
	if m == 0 {
		return CPtr{}
	}
	b.B[b.O+m] = 0
	return b
}

func Remove(name CPtr) int32 {
	if os.Remove(Str(name)) != nil {
		return -1
	}
	return 0
}

func Mkstemp(template CPtr) int32 {
	n := strlen(template)
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for t := 0; t < 1000; t++ {
		for i := n - 6; i < n; i++ {
			template.S(i, int32(chars[rand.Intn(len(chars))]))
		}
		fh, err := os.OpenFile(Str(template), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
		if err == nil {
			fh.Close()
			return 3
		}
		if !os.IsExist(err) {
			return -1
		}
	}
	return -1
}

// RealpathOf is realpath(3) for the hand-written getAbsolutePath (an absolute path if it cannot resolve).
func RealpathOf(p string) string {
	if r, err := filepath.EvalSymlinks(p); err == nil {
		if a, err := filepath.Abs(r); err == nil {
			return a
		}
	}
	a, _ := filepath.Abs(p)
	return a
}

// ------------------------------------------------------------ binary I/O

func bitsOf[T Number](v T) uint64 {
	switch x := any(v).(type) {
	case float64:
		return math.Float64bits(x)
	case float32:
		return uint64(math.Float32bits(x))
	case int32:
		return uint64(uint32(x))
	case uint32:
		return uint64(x)
	case int64:
		return uint64(x)
	case uint64:
		return x
	}
	panic("bitsOf")
}

func fromBits[T Number](b uint64) T {
	var z T
	switch any(z).(type) {
	case float64:
		return any(math.Float64frombits(b)).(T)
	case float32:
		return any(math.Float32frombits(uint32(b))).(T)
	case int32:
		return any(int32(uint32(b))).(T)
	case uint32:
		return any(uint32(b)).(T)
	case int64:
		return any(int64(b)).(T)
	case uint64:
		return any(b).(T)
	}
	panic("fromBits")
}

// Fwrite1 writes the first size bytes (little-endian) of one value.
func Fwrite1[T Number, I Integer](v T, size I, f *CFile) uint64 {
	var io [8]byte
	binary.LittleEndian.PutUint64(io[:], bitsOf(v))
	f.write(io[:int(size)])
	return 1
}

func Fread1[T Number, I Integer](p *T, size I, f *CFile) uint64 {
	var io [8]byte
	binary.LittleEndian.PutUint64(io[:], bitsOf(*p))
	if f.read(io[:int(size)]) < int(size) {
		f.eof = true
		return 0
	}
	*p = fromBits[T](binary.LittleEndian.Uint64(io[:]))
	return 1
}

func Fwrite[T Number, I Integer](a []T, n I, f *CFile) uint64 {
	var z T
	sz := int(unsafe.Sizeof(z))
	buf := make([]byte, 8*int(n))
	for i := 0; i < int(n); i++ {
		binary.LittleEndian.PutUint64(buf[i*sz:], bitsOf(a[i]))
	}
	f.write(buf[:sz*int(n)])
	return uint64(n)
}

func Fread[T Number, I Integer](a []T, n I, f *CFile) uint64 {
	var z T
	sz := int(unsafe.Sizeof(z))
	var io [8]byte
	for i := 0; i < int(n); i++ {
		if f.read(io[:sz]) < sz {
			f.eof = true
			return uint64(i)
		}
		a[i] = fromBits[T](binary.LittleEndian.Uint64(io[:]))
	}
	return uint64(n)
}

func FwriteC[I Integer](p CPtr, n I, f *CFile) uint64 {
	f.write(p.B[p.O : int(p.O)+int(n)])
	return uint64(n)
}

func FreadC[I Integer](p CPtr, n I, f *CFile) uint64 {
	got := f.read(p.B[p.O : int(p.O)+int(n)])
	if got < int(n) {
		f.eof = true
	}
	return uint64(got)
}

// ------------------------------------------------------------ time

func Time() int64 { return time.Now().Unix() }

func Ctime(t int64) CPtr {
	d := time.Unix(t, 0).Local()
	return Lit(fmt.Sprintf("%s %s %2d %02d:%02d:%02d %d\n", d.Weekday().String()[:3], d.Month().String()[:3],
		d.Day(), d.Hour(), d.Minute(), d.Second(), d.Year()))
}

// Out writes a Go string to standard output at once (for the hand-written runner).
func Out(s string) {
	Stdout.write([]byte(s))
	Stdout.flush()
}
