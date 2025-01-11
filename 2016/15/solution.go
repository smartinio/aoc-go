package main

import (
	_ "embed"
	"main/perf"
	"main/utils"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

type Disc struct {
	positions, t0 int
}

func solution() (part1 int, part2 int) {
	discs := parseInput()

	// part 1
	for t := 0; true; t++ {
		if aligned(discs, t) {
			part1 = t
			break
		}
	}

	// part 2
	discs = append(discs, Disc{11, 0})
	for t := 0; true; t++ {
		if aligned(discs, t) {
			part2 = t
			break
		}
	}

	return part1, part2
}

func aligned(discs []Disc, t int) bool {
	for n, disc := range discs {
		position := (disc.t0 + t + n + 1) % disc.positions
		if position != 0 {
			return false
		}
	}
	return true
}

func parseInput() []Disc {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	discs := make([]Disc, len(lines))
	re := regexp.MustCompile(` (\d+)`)

	for i, line := range lines {
		d := utils.FindAllIntGroups(re, line)
		discs[i] = Disc{d[0], d[1]}
	}

	return discs
}

func main() {
	perf.Bench(1, solution)
}
