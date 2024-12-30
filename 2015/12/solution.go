package main

import (
	_ "embed"
	"encoding/json"
	"main/perf"
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0

	// part 1
	{
		re := regexp.MustCompile(`-?\d+`)
		for _, line := range strings.Split(input, "\n") {
			for _, m := range re.FindAllString(line, -1) {
				num, _ := strconv.Atoi(m)
				part1 += num
			}
		}
	}

	// part 2
	{
		var object any
		json.Unmarshal([]byte(input), &object)
		part2 = dfs(object)
	}

	return part1, part2
}

func dfs(object any) int {
	sum := 0

	switch v := object.(type) {
	case map[string]any:
		mapValues := slices.Collect(maps.Values(v))
		if !slices.Contains(mapValues, "red") {
			for _, item := range v {
				sum += dfs(item)
			}
		}
	case []any:
		for _, item := range v {
			sum += dfs(item)
		}
	case float64:
		sum = int(v)
	}

	return sum
}

func main() {
	perf.Bench(1, solution)
}
