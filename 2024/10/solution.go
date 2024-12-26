package main

import (
	_ "embed"
	"main/perf"
	"slices"
	"strings"
)

type Dir struct{ x, y int }
type Pos struct{ x, y int }

//go:embed input.txt
var input string
var W = strings.Index(input, "\n")
var H = strings.Count(input, "\n")
var DIRS = []Dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func solution() (int, int) {
	part1, part2 := 0, 0

	for y := range H {
		for x := range W {
			if numberAt(Pos{x, y}) == 0 {
				part1 += score(Pos{x, y}, []Pos{}, make(map[Pos]bool), 0, true)
				part2 += score(Pos{x, y}, []Pos{}, make(map[Pos]bool), 0, false)
			}
		}
	}

	return part1, part2
}

func score(pos Pos, path []Pos, reached map[Pos]bool, expected int, unique bool) int {
	n := numberAt(pos)

	if n != expected || slices.Contains(path, pos) || reached[pos] {
		return 0
	}

	if n == 9 {
		reached[pos] = unique
		return 1
	}

	path = append([]Pos{pos}, path...)
	sum := 0

	for _, d := range DIRS {
		next := Pos{pos.x + d.x, pos.y + d.y}
		sum += score(next, path, reached, expected+1, unique)
	}

	return sum
}

func numberAt(pos Pos) int {
	x, y := pos.x, pos.y

	if x >= W || y >= H || x < 0 || y < 0 {
		return -1
	}

	return int(input[y*(W+1)+x]) - '0'
}

func main() {
	perf.Bench(20, solution)
}
