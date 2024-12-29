package main

import (
	_ "embed"
	"main/perf"
	"main/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Pair [2]string

func solution() (float64, float64) {
	weights, locations := parseInput()
	part1, part2 := math.Inf(1), 0.0

	for _, start := range locations {
		part1 = min(part1, dfs1([]string{start}, locations, 0, weights))
		part2 = max(part2, dfs2([]string{start}, locations, 0, weights))
	}

	return part1, part2
}

func dfs1(path, locations []string, dist float64, weights map[Pair]float64) float64 {
	if len(path) == len(locations) {
		return dist
	}

	curr, minimum := path[0], math.Inf(1)

	for _, next := range locations {
		if !slices.Contains(path, next) {
			w := weights[Pair{curr, next}]
			nextPath := append([]string{next}, path...)
			minimum = min(minimum, dfs1(nextPath, locations, dist+w, weights))
		}
	}

	return minimum
}

func dfs2(path, locations []string, dist float64, weights map[Pair]float64) float64 {
	if len(path) == len(locations) {
		return dist
	}

	curr, maximum := path[0], 0.0

	for _, next := range locations {
		if !slices.Contains(path, next) {
			w := weights[Pair{curr, next}]
			nextPath := append([]string{next}, path...)
			maximum = max(maximum, dfs2(nextPath, locations, dist+w, weights))
		}
	}

	return maximum
}

func parseInput() (map[Pair]float64, []string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	weights := map[Pair]float64{}
	locations := utils.NewSet[string]()

	for _, line := range lines {
		s := strings.Split(line, " ")
		k1 := Pair{s[0], s[2]}
		k2 := Pair{s[2], s[0]}
		d, _ := strconv.ParseFloat(s[4], 64)
		weights[k1], weights[k2] = d, d
		locations.Add(s[0], s[2])
	}

	return weights, slices.Collect(locations.Values())
}

func main() {
	perf.Bench(1, solution)
}
