package main

import (
	_ "embed"
	"main/perf"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Range = struct{ start, end int }

func solution() (int, int) {
	part1, part2 := 0, 0
	ranges, ids := parseInput()

	for i, a := range ranges {
		for j := i + 1; j < len(ranges); j++ {
			b := ranges[j]
			normalize(a, b)
		}

		for _, id := range ids {
			if id >= a.start && id <= a.end {
				part1++
			}
		}

		if a.end != 0 {
			part2 += 1 + a.end - a.start
		}
	}

	return part1, part2
}

func normalize(a *Range, b *Range) {
	if contains(a, b) {
		b.start, b.end = 0, 0
	} else if contains(b, a) {
		a.start, a.end = 0, 0
	} else if overlaps(a, b) {
		a.end = b.start - 1
	} else if overlaps(b, a) {
		b.end = a.start - 1
	}
}

func contains(a *Range, b *Range) bool {
	return b.start >= a.start && b.end <= a.end
}

func overlaps(a *Range, b *Range) bool {
	return b.start > a.start && b.start <= a.end && b.end > a.end
}

func parseInput() ([]*Range, []int) {
	blocks := strings.Split(input, "\n\n")
	rangeStrings := strings.Split(blocks[0], "\n")
	idStrings := strings.Split(strings.TrimSpace(blocks[1]), "\n")
	ranges := make([]*Range, len(rangeStrings))
	ids := make([]int, len(idStrings))

	for i, rng := range rangeStrings {
		r := strings.Split(rng, "-")
		start, _ := strconv.Atoi(r[0])
		end, _ := strconv.Atoi(r[1])
		ranges[i] = &Range{start, end}
	}

	for i, str := range idStrings {
		id, _ := strconv.Atoi(str)
		ids[i] = id
	}

	return ranges, ids
}

func main() {
	perf.Bench(1, solution)
}
