package main

import (
	_ "embed"
	"main/perf"
	"main/utils"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var INVALID = Result{-1, -1}

type Result struct {
	count, entanglement int
}

func solution() (part1 int, part2 int) {
	packages := parseInput()

	part1 = split(packages, 3).entanglement
	part2 = split(packages, 4).entanglement

	return part1, part2
}

func split(packages []int, groups int) Result {
	for i := range packages {
		best := INVALID

		for g1 := range utils.Choose(packages, i+1) {
			if weight(packages)-weight(g1) == (groups-1)*weight(g1) {
				best = optimal(best, Result{len(g1), entanglement(g1)})
			}
		}

		if best != INVALID {
			return best
		}
	}

	return INVALID
}

func optimal(a, b Result) Result {
	switch {
	case b == INVALID:
		return a
	case a == INVALID:
		return b
	case a.count < b.count:
		return a
	case b.count < a.count:
		return b
	case a.entanglement < b.entanglement:
		return a
	default:
		return b
	}
}

func weight(group []int) (w int) {
	for _, weight := range group {
		w += weight
	}
	return w
}

func entanglement(group []int) (e int) {
	e = 1
	for _, weight := range group {
		e *= weight
	}
	return e
}

func parseInput() []int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	packages := make([]int, len(lines))

	for i, line := range lines {
		n, _ := strconv.Atoi(line)
		packages[i] = int(n)
	}

	return packages
}

func main() {
	perf.Bench(1, solution)
}
