package main

import (
	_ "embed"
	"main/perf"
	"main/utils"
	"regexp"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

const maxInt = int(^uint32(0)) // 9 for example.txt

type Range struct{ start, end int }

func solution() (part1 int, part2 int) {
	ranges := parseInput()

	// part 1
	for _, r := range ranges {
		if r.start > part1 {
			break
		}

		part1 = max(part1, r.end+1)
	}

	// part 2
	end := 0

	for _, r := range ranges {
		part2 += max(0, r.start-(end+1))
		end = max(end, r.end)
	}

	part2 += maxInt - end

	return part1, part2
}

func parseInput() []Range {
	re := regexp.MustCompile(`(\d+)`)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ranges := make([]Range, len(lines))

	for i, line := range lines {
		r := utils.FindAllIntGroups(re, line)
		ranges[i] = Range{r[0], r[1]}
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return a.start - b.start
	})

	return ranges
}

func main() {
	perf.Bench(1, solution)
}
