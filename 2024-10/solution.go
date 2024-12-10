package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"
)

type Dir struct{ x, y int }
type Pos struct{ x, y int }

//go:embed input.txt
var input string
var grid [][]int
var W = strings.Index(input, "\n")
var H = strings.Count(input, "\n")
var DIRS = []Dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func solution() (int, int) {
	part1 := 0
	part2 := 0

	for y := range H {
		row := []int{}
		for x := range W {
			n := numberAt(Pos{x, y})
			row = append(row, n)
		}
		grid = append(grid, row)
	}

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
	x, y := pos.x, pos.y

	if x < 0 || y < 0 || x >= W || y >= H || slices.Contains(path, pos) || reached[pos] {
		return 0
	}

	n := grid[y][x]

	if n != expected {
		return 0
	}

	if n == 9 {
		reached[pos] = unique
		return 1
	}

	path = append([]Pos{pos}, path...)
	sum := 0

	for _, d := range DIRS {
		sum += score(Pos{x + d.x, y + d.y}, path, reached, expected+1, unique)
	}

	return sum
}

func numberAt(pos Pos) int {
	x, y := pos.x, pos.y

	if x >= W || y >= H || x < 0 || y < 0 {
		return 0
	}

	return int(input[int(y)*int(W+1)+int(x)]) - 48 // 48 == '0' rune
}

func main() {
	part1, part2 := 0, 0
	sum := 0
	n := 20

	for range n {
		start := time.Now()
		part1, part2 = solution()
		sum += int(time.Since(start).Milliseconds())
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
	fmt.Println("avg:", sum/n, "ms")
}
