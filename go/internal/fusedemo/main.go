// fusedemo - the two forms of x*y + z the Go back end chooses between: plain (a Go compiler may fuse it
// into one FMA) and float64(x*y) + z (the specification forbids fusing). gate/check_go_fma_arm64.sh compiles it.
// Written by Claude (Anthropic); see AI-USE.md.
package main

import (
	"fmt"
	"os"
	"strconv"
)

//go:noinline
func plain(x, y, z float64) float64 { return x*y + z }

//go:noinline
func wrapped(x, y, z float64) float64 { return float64(x*y) + z }

func main() {
	a, _ := strconv.ParseFloat(os.Args[1], 64)
	fmt.Println(plain(a, a, -1), wrapped(a, a, -1))
}
