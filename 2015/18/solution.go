package main

import (
	_ "embed"
	"main/perf"
	"slices"
	"strings"
)

//go:embed input.txt
var input string
var steps = 100
var size = strings.Index(input, "\n")
var ds = []int{-1, 0, 1}

type Grid [][]bool
type Pos struct{ x, y int }

func solution() (int, int) {
	part1, part2 := 0, 0

	// part 1
	{
		stuck := []Pos{}
		grid := animate(stuck)

		for _, row := range grid {
			part1 += count(row, true)
		}
	}

	// part 2
	{
		n := size - 1
		stuck := []Pos{{0, 0}, {0, n}, {n, 0}, {n, n}}
		grid := animate(stuck)

		for _, row := range grid {
			part2 += count(row, true)
		}
	}

	return part1, part2
}

func animate(stuck []Pos) Grid {
	grid := parseInput()
	size := len(grid)

	for _, pos := range stuck {
		grid[pos.y][pos.x] = true
	}

	for range steps {
		toggles := []Pos{}

		for y := range size {
			for x := range size {
				on := grid[y][x]
				adj := neighbors(Pos{x, y}, grid)
				adjOn := count(adj, true)
				if on && adjOn != 2 && adjOn != 3 {
					toggles = append(toggles, Pos{x, y})
				}
				if !on && adjOn == 3 {
					toggles = append(toggles, Pos{x, y})
				}
			}
		}

		for _, pos := range toggles {
			if !slices.Contains(stuck, pos) {
				grid[pos.y][pos.x] = !grid[pos.y][pos.x]
			}
		}
	}

	return grid
}

func count(slice []bool, val bool) (sum int) {
	for _, item := range slice {
		if item == val {
			sum++
		}
	}
	return sum
}

func neighbors(pos Pos, grid Grid) (adj []bool) {
	for _, dy := range ds {
		for _, dx := range ds {
			a := Pos{pos.x + dx, pos.y + dy}
			if a != pos && a.x >= 0 && a.y >= 0 && a.x < len(grid) && a.y < len(grid) {
				adj = append(adj, grid[a.y][a.x])
			}
		}
	}

	return adj
}

func parseInput() Grid {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make(Grid, len(lines))
	for i := range grid {
		grid[i] = make([]bool, len(lines))
	}

	for y, line := range lines {
		for x, char := range line {
			grid[y][x] = char == '#'
		}
	}

	return grid
}

func main() {
	perf.Bench(1, solution)
}
