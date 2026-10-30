// check.go - compare the Go glibc ports (and math.FMA) bit for bit with system-libm vectors (mkvectors.c,
// mkedge.c). Run from go/ as: go run ../tools/libm/test/check.go VECTORS.bin   (it imports swmm5go/glibm)
// Written by Claude (Anthropic); see AI-USE.md.


package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"time"

	M "swmm5go/glibm"
)

func main() {
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	p := 0
	f64 := func() float64 { v := math.Float64frombits(binary.LittleEndian.Uint64(b[p:])); p += 8; return v }
	for p < len(b) {
		s := p
		for b[p] != 0 {
			p++
		}
		name := string(b[s:p])
		p++
		n := int(binary.LittleEndian.Uint32(b[p:]))
		p += 4
		na := 1
		if name == "pow" {
			na = 2
		} else if name == "fma" {
			na = 3
		}
		bad, nanSign, shown := 0, 0, 0
		t0 := time.Now()
		for i := 0; i < n; i++ {
			x := f64()
			y, z := 0.0, 0.0
			if na > 1 {
				y = f64()
			}
			if na > 2 {
				z = f64()
			}
			want := f64()
			var got float64
			switch name {
			case "exp":
				got = M.Exp(x)
			case "log":
				got = M.Log(x)
			case "log10":
				got = M.Log10(x)
			case "pow":
				got = M.Pow(x, y)
			case "sin":
				got = M.Sin(x)
			case "cos":
				got = M.Cos(x)
			case "tan":
				got = M.Tan(x)
			case "asin":
				got = M.Asin(x)
			case "acos":
				got = M.Acos(x)
			case "atan":
				got = M.Atan(x)
			default:
				got = math.FMA(x, y, z)
			}
			gb, wb := math.Float64bits(got), math.Float64bits(want)
			if math.IsNaN(got) && math.IsNaN(want) {
				if gb>>63 != wb>>63 {
					nanSign++
				}
				continue
			}
			if gb != wb {
				bad++
				if shown < 4 {
					shown++
					fmt.Printf("  %s(%v, %v, %v) = %v, libm %v\n", name, x, y, z, got, want)
				}
			}
		}
		fmt.Printf("%-6s %d of %d differ, %d NaN sign   (%.3f us/call)\n", name, bad, n, nanSign,
			float64(time.Since(t0).Nanoseconds())/1e3/float64(n))
	}
}
