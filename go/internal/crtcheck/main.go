// crtcheck - compare the Go port's printf, strtod and sscanf with the system glibc's results written by
// tools/crt/test/mkfmt.c.   cd go && go run ./internal/crtcheck FMT.txt
// Written by Claude (Anthropic); see AI-USE.md.
package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	R "swmm5go/crt"
)

func main() {
	fh, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(fh)
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	total := map[string]int{}
	bad := map[string]int{}
	shown := 0
	out := R.Buf(1024)
	unhex := func(s string) []byte {
		b, err := hex.DecodeString(s)
		if err != nil {
			panic(err)
		}
		return b
	}
	cstr := func(b []byte) R.CPtr { return R.Lit(string(b)) }
	for sc.Scan() {
		f := strings.Split(sc.Text(), "\t")
		kind, fm, arg, want := f[0], f[1], f[2], f[3]
		var got string
		switch kind {
		case "f":
			b, _ := strconv.ParseUint(arg, 16, 64)
			x := math.Float64frombits(b)
			if fm == "%*.*f" {
				R.Snprintf(out, 1024, R.Lit(fm), int32(9), int32(3), x)
			} else {
				R.Snprintf(out, 1024, R.Lit(fm), x)
			}
			got = hex.EncodeToString([]byte(R.Str(out)))
		case "i":
			b, _ := strconv.ParseUint(arg, 16, 32)
			R.Snprintf(out, 1024, R.Lit(fm), int32(uint32(b)))
			got = hex.EncodeToString([]byte(R.Str(out)))
		case "l":
			b, _ := strconv.ParseUint(arg, 16, 64)
			R.Snprintf(out, 1024, R.Lit(fm), int64(b))
			got = hex.EncodeToString([]byte(R.Str(out)))
		case "s":
			R.Snprintf(out, 1024, R.Lit(fm), cstr(unhex(arg)))
			got = hex.EncodeToString([]byte(R.Str(out)))
		case "d":
			s := cstr(unhex(arg))
			var end R.CPtr
			v := R.Strtod(s, &end)
			got = fmt.Sprintf("%016x:%d", math.Float64bits(v), end.O-s.O)
		case "n":
			s := cstr(unhex(arg))
			i1, d1, b := int32(-77), -77.0, R.Lit("?")
			b = R.BufFrom(b, 32)
			n := R.Sscanf(s, R.Lit("%d %lf %s"), &i1, &d1, b)
			got = fmt.Sprintf("%d:%d:%016x:%s", n, i1, math.Float64bits(d1), hex.EncodeToString([]byte(R.Str(b))))
		}
		key := kind + " " + fm
		total[key]++
		if got != want {
			bad[key]++
			if shown < 25 {
				shown++
				fmt.Printf("  %s %s arg=%s  glibc=%q  go=%q\n", kind, fm, arg, dec(kind, want), dec(kind, got))
			}
		}
	}
	nb, nt := 0, 0
	for k, n := range total {
		nt += n
		nb += bad[k]
		if bad[k] > 0 {
			fmt.Printf("%-12s %d of %d differ\n", k, bad[k], n)
		}
	}
	kinds := map[string][2]int{}
	for k, n := range total {
		g := map[byte]string{'f': "printf", 'i': "printf", 'l': "printf", 's': "printf", 'd': "strtod", 'n': "sscanf"}[k[0]]
		v := kinds[g]
		kinds[g] = [2]int{v[0] + n, v[1] + bad[k]}
	}
	for _, g := range []string{"printf", "strtod", "sscanf"} {
		fmt.Printf("%s: %d cases, %d differ\n", g, kinds[g][0], kinds[g][1])
	}
	fmt.Printf("%d of %d cases differ (%d formats)\n", nb, nt, len(total))
}

func dec(kind, s string) string {
	if kind == "d" || kind == "n" {
		return s
	}
	b, _ := hex.DecodeString(s)
	return string(b)
}
