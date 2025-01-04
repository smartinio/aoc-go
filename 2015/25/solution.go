package main

import (
	_ "embed"
	"main/perf"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func solution() (part1 int, part2 int) {
	row, col := parseInput()

	// find number of iterations needed to reach row, col
	n := 0

	// position cell x-wise
	for i := 1; i <= col; i++ {
		n += i
	}

	// position cell y-wise
	for i := 0; i < row-1; i++ {
		n += col + i
	}

	code := 20151125
	for range n - 1 {
		code = next(code)
	}

	return code, 1337
}

func next(code int) int {
	return (code * 252533) % 33554393
}

func parseInput() (row, col int) {
	re := regexp.MustCompile(`\d+`)
	m := re.FindAllString(input, 2)
	row, _ = strconv.Atoi(m[0])
	col, _ = strconv.Atoi(m[1])
	return row, col
}

func main() {
	perf.Bench(10, solution)
}
