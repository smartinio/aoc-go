package main

import (
	_ "embed"
	"main/perf"
	"slices"
	"strings"
)

//go:embed input.txt
var input string
var diagram = []byte(input)
var w = strings.Index(input, "\n")
var h = strings.Count(input, "\n")
var dirs = []int{-1, 0, 1}

func solution() (int, int) {
	part1, part2 := 0, 0

	for y := range h {
		for x := range w {
			if removable(x, y) {
				part1++
			}
		}
	}

	for {
		removed := 0

		for y := range h {
			for x := range w {
				if removable(x, y) {
					remove(x, y)
					removed++
				}
			}
		}

		if removed == 0 {
			break
		}

		part2 += removed
	}

	return part1, part2
}

func removable(x int, y int) bool {
	return charAt(x, y) == '@' && countAdj(x, y) < 4
}

func countAdj(x int, y int) int {
	sum := 0

	for _, dy := range dirs {
		for _, dx := range dirs {
			nx, ny := x+dx, y+dy
			if !(dx == 0 && dy == 0) && charAt(nx, ny) == '@' {
				sum++
			}
		}
	}

	return sum
}

func charAt(x int, y int) byte {
	if x >= w || y >= h || x < 0 || y < 0 {
		return 0
	}

	return diagram[y*(w+1)+x]
}

func remove(x int, y int) {
	if x >= w || y >= h || x < 0 || y < 0 {
		return
	}

	index := y*(w+1) + x
	diagram = slices.Replace(diagram, index, index+1, 'x')
}

func main() {
	perf.Bench(1, solution)
}
