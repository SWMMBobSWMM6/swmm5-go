// main.go - command-line runner for the Go port: same arguments and console messages as EPA's
// run/main.c.   runswmm <input file> <report file> [output file]
// Written by Claude (Anthropic); see AI-USE.md.
package main

import (
	"os"
	"time"

	R "swmm5go/crt"
)

func main() {
	version := swmm_getVersion()
	vMajor := version / 10000
	vMinor := (version - 10000*vMajor) / 1000
	vRelease := version - 10000*vMajor - 1000*vMinor
	argv := os.Args[1:]
	switch {
	case len(argv) == 0:
		R.Out("\nNot Enough Arguments (See Help --help)\n\n")
	case len(argv) == 1:
		a := argv[0]
		if a == "--help" || a == "-h" {
			R.Out("\n\nSTORMWATER MANAGEMENT MODEL (SWMM) HELP\n\nCOMMANDS:\n\t--help (-h)       SWMM Help\n" +
				"\t--version (-v)    Build Version\n\nRUNNING A SIMULATION:\n" +
				"\t runswmm <input file> <report file> <optional output file>\n\n")
		} else if a == "--version" || a == "-v" {
			R.Out(R.Sprintf("\n%d.%d.%d\n\n", vMajor, vMinor, vRelease))
		} else {
			R.Out("\nUnknown Argument (See Help --help)\n\n")
		}
	default:
		start := time.Now()
		R.Out(R.Sprintf("\n... EPA SWMM %d.%d (Build %d.%d.%d) [Go port]\n", vMajor, vMinor, vMajor, vMinor, vRelease))
		out := ""
		if len(argv) > 2 {
			out = argv[2]
		}
		swmm_run(R.Lit(argv[0]), R.Lit(argv[1]), R.Lit(out))
		msg := R.Sprintf("\n\n... EPA SWMM completed in %.2f seconds.", time.Since(start).Seconds())
		errbuf := R.Buf(128)
		if swmm_getError(errbuf, 127) > 0 {
			msg += " There are errors.\n"
		} else if swmm_getWarnings() > 0 {
			msg += " There are warnings.\n"
		} else {
			msg += "\n"
		}
		R.Out(msg)
	}
	R.Fflush(R.Stdout)
}
