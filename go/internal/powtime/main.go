// powtime - nanoseconds per call of the Go port's glibc translations (glibm) and of Go's own math, on the
// same arguments as tools/libm/test/powtime.c.   cd go && go run ./internal/powtime [LABEL [nomath]]
// Written by Claude (Anthropic); see AI-USE.md.
package main

import (
	"fmt"
	"math"
	"os"
	"time"

	M "swmm5go/glibm"
)

const N = 2000000

var X, Y [N]float64
var s uint64 = 88172645463325252

func u() float64 { s ^= s << 13; s ^= s >> 7; s ^= s << 17; return float64(s>>11) * 0x1.0p-53 }

func run(name string, pow func(float64, float64) float64, sin, cos func(float64) float64) float64 {
	acc := 0.0
	for r := 0; r < 3; r++ {
		t0 := time.Now()
		a := 0.0
		for i := 0; i < N; i++ {
			a += pow(X[i], Y[i])
		}
		t1 := time.Now()
		b := 0.0
		for i := 0; i < N; i++ {
			b += sin(X[i])
		}
		t2 := time.Now()
		c := 0.0
		for i := 0; i < N; i++ {
			c += cos(X[i])
		}
		t3 := time.Now()
		acc += a + b + c
		if r == 2 {
			ns := func(d time.Duration) float64 { return float64(d.Nanoseconds()) / N }
			fmt.Printf("%-16s pow %5.1f ns   sin %5.1f ns   cos %5.1f ns\n", name, ns(t1.Sub(t0)), ns(t2.Sub(t1)), ns(t3.Sub(t2)))
		}
	}
	return acc
}

func main() {
	for i := 0; i < N; i++ {
		X[i] = 0.001 + 100.0*u()
		switch i % 3 {
		case 0:
			Y[i] = 0.6666666666666666
		case 1:
			Y[i] = 1.6666666666666667
		default:
			Y[i] = -4.0 + 8.0*u()
		}
	}
	label := "Go glibm"
	if len(os.Args) > 1 {
		label = os.Args[1]
	}
	a := run(label, M.Pow, M.Sin, M.Cos)
	b := 0.0
	if len(os.Args) < 3 || os.Args[2] != "nomath" {
		b = run("Go math", math.Pow, math.Sin, math.Cos)
	}
	if a+b == 0.12345 {
		fmt.Println()
	}
}
