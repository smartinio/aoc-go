package main

import (
	_ "embed"
	"main/perf"
	"maps"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func solution() (part1 string, part2 string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	count := map[int]map[rune]int{}
	wl := len(lines[0])

	for i := range wl {
		count[i] = map[rune]int{}
	}

	for _, line := range lines {
		for i, char := range line {
			count[i][char]++
		}
	}

	p1r := make([]rune, wl)
	p2r := make([]rune, wl)

	for i := range wl {
		counts := count[i]
		chars := slices.Collect(maps.Keys(counts))
		compare := func(a, b rune) int {
			return counts[a] - counts[b]
		}
		p1r[i] = slices.MaxFunc(chars, compare)
		p2r[i] = slices.MinFunc(chars, compare)
	}

	part1 = string(p1r)
	part2 = string(p2r)

	return part1, part2
}

func main() {
	perf.Bench(1, solution)
}
