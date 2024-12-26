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

func solution() (int, int) {
	part1, part2 := 0, 0

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		d := strings.Split(line, "x")

		l, _ := strconv.Atoi(d[0])
		w, _ := strconv.Atoi(d[1])
		h, _ := strconv.Atoi(d[2])

		a, b, c := 2*l*w, 2*w*h, 2*h*l

		dims := []int{l, w, h}

		slices.Sort(dims)

		small := dims[0] * dims[1]

		part1 += a + b + c + small

		part2 += 2*dims[0] + 2*dims[1] + w*l*h
	}

	return part1, part2
}

func main() {
	perf.Bench(1, solution)
}
