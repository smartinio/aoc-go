package main

import (
	_ "embed"
	"main/perf"
	"strings"
)

type Dir struct{ x, y int }

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0

	dirs := map[string]Dir{
		"^": {0, -1},
		">": {1, 0},
		"v": {0, 1},
		"<": {-1, 0},
	}

	moves := strings.Split(input, "")

	// part 1
	{
		visited := map[[2]int]bool{{0, 0}: true}
		x, y := 0, 0

		for _, move := range moves {
			d := dirs[move]
			x += d.x
			y += d.y
			key := [2]int{x, y}
			visited[key] = true
		}

		part1 = len(visited)
	}

	// part 2
	{
		visited := map[[2]int]bool{{0, 0}: true}
		x, y := 0, 0
		x2, y2 := 0, 0

		for i, move := range moves {
			d := dirs[move]
			var key [2]int

			if i%2 == 0 {
				x += d.x
				y += d.y
				key = [2]int{x, y}
			} else {
				x2 += d.x
				y2 += d.y
				key = [2]int{x2, y2}
			}

			visited[key] = true
		}

		part2 = len(visited)
	}

	return part1, part2
}

func main() {
	perf.Bench(1, solution)
}
