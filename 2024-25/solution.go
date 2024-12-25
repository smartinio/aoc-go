package main

import (
	_ "embed"
	"main/perf"
	"strings"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 1337
	keys, locks := [][]int{}, [][]int{}
	blocks := strings.Split(strings.TrimSpace(input), "\n\n")

	for _, block := range blocks {
		heights := count(block)

		if strings.HasPrefix(block, "#") {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}

	for _, lock := range locks {
		for _, key := range keys {
			if fits(lock, key) {
				part1++
			}
		}
	}

	return part1, part2
}

func count(block string) []int {
	result := make([]int, strings.Index(block, "\n"))
	lines := strings.Split(block, "\n")

	for _, line := range lines[1 : len(lines)-1] {
		for j, c := range line {
			if c == '#' {
				result[j]++
			}
		}
	}

	return result
}

func fits(lock, key []int) bool {
	for i := range len(key) {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}

func main() {
	perf.Bench(1, solution)
}
