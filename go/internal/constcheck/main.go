// constcheck - why the Go back end types every float literal and wraps every float operation:
// Go evaluates constant expressions exactly and rounds once, where C rounds after every operation.
// cd go && go run ./internal/constcheck     (compare: C's 0.1*3.0 is 0x3fd3333333333334)
// Written by Claude (Anthropic); see AI-USE.md.
package main

import (
	"fmt"
	"math"
)

func main() {
	a := 0.1 * 3.0                   // untyped constants: the exact 0.3, rounded once
	b := float64(float64(0.1) * 3.0) // typed, as the back end writes it: the double nearest 0.1, times 3, rounded
	x := 0.1
	c := x * 3.0 // at run time, as C computes it
	fmt.Printf("constant 0.1 * 3.0                 = %.21f  bits %016x\n", a, math.Float64bits(a))
	fmt.Printf("typed float64(float64(0.1) * 3.0)  = %.21f  bits %016x\n", b, math.Float64bits(b))
	fmt.Printf("run time x * 3.0 (x = 0.1)         = %.21f  bits %016x\n", c, math.Float64bits(c))
	z := 0.0
	fmt.Printf("constant -0.0 has its sign bit: %v; run-time -z (z = 0.0): %v\n", math.Signbit(-0.0), math.Signbit(-z))
}
