package main

import (
	_ "embed"
	"main/perf"
	"maps"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0

	inventory := parseInput()

	counts := map[int]int{}

	part1 = dfs(0, 0, []int{}, 150, inventory, counts)

	minCount := slices.Min(slices.Collect(maps.Keys(counts)))

	part2 = counts[minCount]

	return part1, part2
}

func dfs(acc int, count int, checked []int, target int, inventory []int, counts map[int]int) int {
	if acc > target {
		return 0
	}

	if acc == target {
		counts[count]++
		return 1
	}

	ways := 0

	for i, box := range inventory {
		if !slices.Contains(checked, i) {
			checked = append(checked, i)
			ways += dfs(acc+box, count+1, checked, target, inventory, counts)
		}
	}

	return ways
}

func parseInput() (inventory []int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		inventory = append(inventory, n)
	}

	return inventory
}

func main() {
	perf.Bench(1, solution)
}
