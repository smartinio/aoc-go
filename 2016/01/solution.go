package main

import (
	_ "embed"
	"main/perf"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Dir struct{ x, y int }
type Pos struct{ x, y int }

var dirs = []Dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func solution() (int, int) {
	part1, part2 := 0, -1
	directions := strings.Split(strings.TrimSpace(input), ", ")
	dir, pos := 0, Pos{}
	been := map[Pos]bool{}

	for _, line := range directions {
		switch line[0] {
		case 'L':
			dir = wrap(dir-1, len(dirs))
		case 'R':
			dir = wrap(dir+1, len(dirs))
		}

		d := dirs[dir]
		dist, _ := strconv.Atoi(line[1:])

		for range dist {
			pos.x += d.x
			pos.y += d.y

			if been[pos] && part2 == -1 {
				part2 = abs(pos.x) + abs(pos.y)
			}

			been[pos] = true
		}
	}

	part1 = abs(pos.x) + abs(pos.y)

	return part1, part2
}

func abs(n int) int {
	return max(n, 0) - min(n, 0)
}

func wrap(i, max int) int {
	return ((i % max) + max) % max
}

func main() {
	perf.Bench(1, solution)
}
