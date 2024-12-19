package main

import (
	_ "embed"
	"main/perf"
)

//go:embed example.txt
var input string

func solution() (int, int) {
	part1 := 0
	part2 := 0

	return part1, part2
}

func main() {
	perf.Bench(1, solution)
}
