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

	part1 = joltage(lines, 2)
	part2 = joltage(lines, 12)

	return part1, part2
}

func joltage(lines []string, length int) int {
	sum := 0

	for _, line := range lines {
		str := extract(line, length)
		value, _ := strconv.Atoi(str)
		sum += value
	}

	return sum
}

func extract(line string, length int) string {
	if length == 0 {
		return ""
	}

	for i := range 9 {
		num := strconv.Itoa(9 - i)
		end := len(line) - length + 1
		best := strings.Index(line[0:end], num)

		if best != -1 {
			return string(line[best]) + extract(line[best+1:], length-1)
		}
	}

	return ""
}

func main() {
	perf.Bench(1, solution)
}
