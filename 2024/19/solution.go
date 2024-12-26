package main

import (
	_ "embed"
	"main/perf"
	"strings"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0

	s := strings.Split(strings.TrimSpace(input), "\n\n")
	colors := strings.Split(s[0], ", ")
	towels := strings.Split(s[1], "\n")
	colorSet := map[string]bool{}
	cache := map[string]int{}

	maxLen := 0
	for _, c := range colors {
		maxLen = max(maxLen, len(c))
		colorSet[c] = true
	}

	for _, towel := range towels {
		ways := bt(towel, maxLen, colorSet, cache)

		if ways != 0 {
			part1 += 1
			part2 += ways
		}
	}

	return part1, part2
}

func bt(towel string, maxLen int, colors map[string]bool, cache map[string]int) int {
	if ways, cached := cache[towel]; cached {
		return ways
	}

	if towel == "" {
		return 1
	}

	maxHeadLen := min(len(towel), maxLen)
	ways := 0

	for i := range maxHeadLen {
		head := towel[:i+1]

		if colors[head] {
			tail := towel[i+1:]
			ways += bt(tail, maxLen, colors, cache)
		}
	}

	cache[towel] = ways

	return ways
}

func main() {
	perf.Bench(100, solution)
}
