package main

import (
	_ "embed"
	"main/perf"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Deltas map[string]map[string]int

func solution() (int, int) {
	part1, part2 := 0, 0

	deltas := parseInput()

	part1 = dfs([]string{}, deltas)

	deltas["Me"] = map[string]int{}

	part2 = dfs([]string{}, deltas)

	return part1, part2
}

func dfs(seating []string, deltas Deltas) int {
	if len(seating) == len(deltas) {
		return sum(seating, deltas)
	}

	maximum := 0

	for person := range deltas {
		if !slices.Contains(seating, person) {
			total := dfs(append(seating, person), deltas)
			maximum = max(maximum, total)
		}
	}

	return maximum
}

func sum(seating []string, deltas Deltas) int {
	total, mod := 0, len(seating)

	for i, person := range seating {
		left, right := seating[wrap(i-1, mod)], seating[wrap(i+1, mod)]
		total += deltas[person][left] + deltas[person][right]
	}

	return total
}

func wrap(i, max int) int {
	return ((i % max) + max) % max
}

func parseInput() Deltas {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	deltas := Deltas{}

	for _, line := range lines {
		l := strings.Split(strings.TrimRight(line, "."), " ")
		person, sign, h, neighbor := l[0], l[2], l[3], l[10]
		happiness, _ := strconv.Atoi(h)

		if sign == "lose" {
			happiness *= -1
		}

		if deltas[person] == nil {
			deltas[person] = map[string]int{}
		}

		deltas[person][neighbor] = happiness
	}

	return deltas
}

func main() {
	perf.Bench(1, solution)
}
