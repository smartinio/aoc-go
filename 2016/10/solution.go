package main

import (
	_ "embed"
	"main/perf"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Node struct {
	kind, id string
}

var FOUND = Node{id: "found"}

func solution() (int, int) {
	part1, part2 := -1, 1
	sources, outputs := parseInput()
	cache := map[Node]int{}

	for i, output := range outputs {
		source := sources[output][0] // outputs only have 1 source
		chip := dfs(source, sources, 17, 61, cache)

		if found, ok := cache[FOUND]; ok && part1 == -1 {
			part1 = found
		}

		if i <= 2 {
			part2 *= chip
		}
	}

	return part1, part2
}

func dfs(target Node, sources map[string][]Node, lo, hi int, cache map[Node]int) int {
	if strings.HasPrefix(target.id, "value") {
		return ntoi(target.id)
	}

	if cached, ok := cache[target]; ok {
		return cached
	}

	a := dfs(sources[target.id][0], sources, lo, hi, cache)
	b := dfs(sources[target.id][1], sources, lo, hi, cache)

	if lo == min(a, b) && hi == max(a, b) {
		cache[FOUND] = ntoi(target.id)
	}

	result := -1
	switch target.kind {
	case "lo":
		result = min(a, b)
	case "hi":
		result = max(a, b)
	}

	cache[target] = result

	return result
}

func ntoi(target string) int {
	val, _ := strconv.Atoi(strings.Split(target, " ")[1])
	return val
}

func parseInput() (map[string][]Node, []string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sourceRe := regexp.MustCompile(`^[a-z]+ \d+`)
	targetRe := regexp.MustCompile(`[a-z]+ to ([a-z]+ \d+)`)
	sources := map[string][]Node{}
	outputs := []string{}

	for _, line := range lines {
		srcId := sourceRe.FindString(line)
		for _, t := range targetRe.FindAllStringSubmatch(line, -1) {
			src := Node{id: srcId, kind: t[0][:2]}
			sources[t[1]] = append(sources[t[1]], src)

			if strings.HasPrefix(t[1], "output") && !slices.Contains(outputs, t[1]) {
				outputs = append(outputs, t[1])
			}
		}
	}

	slices.SortFunc(outputs, func(a, b string) int {
		return ntoi(a) - ntoi(b)
	})

	return sources, outputs
}

func main() {
	perf.Bench(1, solution)
}
