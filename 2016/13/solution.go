package main

import (
	_ "embed"
	"main/perf"
	"strconv"
	"strings"
)

type Pos struct{ x, y int }
type Dir struct{ x, y int }

//go:embed input.txt
var input string
var favNum, _ = strconv.Atoi(strings.TrimSpace(input))

func solution() (part1 int, part2 int) {
	part1 = bfs(Pos{31, 39}, 1)
	part2 = bfs(Pos{31, 39}, 2)

	return part1, part2
}

func bfs(target Pos, part int) int {
	start := Pos{1, 1}
	q := []Pos{start}
	steps := map[Pos]int{start: 0}
	dirs := []Dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		s := steps[curr]

		if part == 2 && s == 50 {
			return len(steps) + len(q)
		} else if curr == target {
			return s
		}

		for _, d := range dirs {
			next := Pos{curr.x + d.x, curr.y + d.y}

			if _, visited := steps[next]; !visited && !isWall(next) {
				q = append(q, next)
				steps[next] = s + 1
			}
		}
	}

	return -1
}

func isWall(pos Pos) bool {
	x, y := pos.x, pos.y

	if x < 0 || y < 0 {
		return true
	}

	val := x*x + 3*x + 2*x*y + y + y*y
	val += favNum
	bin := strconv.FormatInt(int64(val), 2)

	return strings.Count(bin, "1")%2 == 1
}

func main() {
	perf.Bench(1, solution)
}
