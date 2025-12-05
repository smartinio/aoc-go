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
	lines := strings.Split(strings.TrimSpace(input), ",")

	for _, line := range lines {
		counted := make(map[int]bool)
		nums := strings.Split(line, "-")
		start, end := nums[0], nums[1]
		n1, _ := strconv.Atoi(start)
		n2, _ := strconv.Atoi(end)

		for times := 2; times <= max(len(start), len(end)); times++ {
			result := sumInvalid(start, times, n1, n2, counted)

			if times == 2 {
				part1 += result
			}

			part2 += result
		}
	}

	return part1, part2
}

func sumInvalid(start string, times, n1, n2 int, counted map[int]bool) int {
	sum := 0
	chunk := start[0 : len(start)/times]

	for {
		invalid, _ := strconv.Atoi(strings.Repeat(chunk, times))

		if invalid > n2 {
			break
		}

		if invalid >= n1 && !counted[invalid] {
			counted[invalid] = true
			sum += invalid
		}

		next, _ := strconv.Atoi(chunk)
		chunk = strconv.Itoa(next + 1)
	}

	return sum
}

func main() {
	perf.Bench(1, solution)
}
