package main

import (
	_ "embed"
	"main/perf"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, -1

	for i, char := range input {
		if char == '(' {
			part1++
		}
		if char == ')' {
			part1--
		}
		if part1 == -1 && part2 == -1 {
			part2 = i + 1
		}
	}

	return part1, part2
}

func main() {
	perf.Bench(1, solution)
}
