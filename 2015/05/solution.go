package main

import (
	_ "embed"
	"main/perf"
	"maps"
	"regexp"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0

	lines := strings.Split(strings.TrimSpace(input), "\n")

	// part 1
	{
		vowelsRe := regexp.MustCompile(`[aeiou]`)
		bannedRe := regexp.MustCompile(`ab|cd|pq|xy`)

		for _, line := range lines {
			vowels := vowelsRe.FindAllString(line, -1)
			if len(vowels) < 3 {
				continue
			}

			prev := '0'
			if !strings.ContainsFunc(line, func(e rune) bool {
				r := e == prev
				prev = e
				return r
			}) {
				continue
			}

			banned := bannedRe.FindAllString(line, -1)
			if len(banned) > 0 {
				continue
			}

			part1++
		}
	}

	// part 2
	for _, line := range lines {
		indexesByPair := map[string][]int{}

		for i := range line[:len(line)-1] {
			pair := line[i : i+2]
			indexesByPair[pair] = append(indexesByPair[pair], i)
		}

		pairIndexes := slices.Collect(maps.Values(indexesByPair))

		if !slices.ContainsFunc(pairIndexes, func(indexes []int) bool {
			return len(indexes) > 2 || len(indexes) == 2 && indexes[1]-indexes[0] > 1
		}) {
			continue
		}

		triples := []string{}
		for i := range line[:len(line)-2] {
			triples = append(triples, line[i:i+3])
		}

		if !slices.ContainsFunc(triples, func(t string) bool {
			return t[0] == t[2]
		}) {
			continue
		}

		part2++
	}

	return part1, part2
}

func main() {
	perf.Bench(1, solution)
}
