package main

import (
	_ "embed"
	"main/perf"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		unquoted, _ := strconv.Unquote(line)
		quoted := strconv.Quote(line)
		part1 += len(line) - len(unquoted)
		part2 += len(quoted) - len(line)
	}

	return part1, part2
}

func main() {
	perf.Bench(1, solution)
}
