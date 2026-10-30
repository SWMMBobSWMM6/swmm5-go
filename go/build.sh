#!/bin/sh
# build.sh - build the Go port of EPA SWMM 5.2.4:  bin/runswmm <input> <report> [output]
# Needs Go 1.22 or later. GOAMD64=v1 (the default) keeps the compiler from using FMA instructions
# anywhere except math.FMA, which the glibc math routines call on purpose (gate/check_go_fma.sh checks).
cd "$(dirname "$0")" && GOAMD64=v1 go build -trimpath -o bin/runswmm ./swmm5
