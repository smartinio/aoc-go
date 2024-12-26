package main

import (
	_ "embed"
	"main/perf"
	"strings"
)

type Pos struct{ x, y int }
type Dir struct{ x, y int }
type Vertex struct {
	pos   Pos
	steps int
}

//go:embed input.txt
var input string
var H = strings.Count(input, "\n")
var W = strings.Index(input, "\n")

func solution() (int, int) {
	part1, part2 := 0, 0
	start := findStart()
	graph := dfs(start, []Vertex{}, 0)

	for t1, a := range graph[:len(graph)-1] {
		for _, b := range graph[t1+2:] {
			delta := dist(a.pos, b.pos)
			saved := b.steps - a.steps - delta

			if saved < 100 {
				continue
			}

			if delta == 2 {
				part1++
			}

			if delta <= 20 {
				part2++
			}
		}
	}

	return part1, part2
}

func dfs(pos Pos, path []Vertex, count int) []Vertex {
	graph := append(path, Vertex{pos, count})

	for _, d := range []Dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		next := Pos{pos.x + d.x, pos.y + d.y}
		visited := count > 0 && next == path[count-1].pos

		if !visited && charAt(next) != '#' {
			return dfs(next, graph, count+1)
		}
	}

	return graph
}

func findStart() Pos {
	for y := range H {
		for x := range W {
			if charAt(Pos{x, y}) == 'S' {
				return Pos{x, y}
			}
		}
	}

	return Pos{-1, -1}
}

func dist(a, b Pos) int {
	dx := max(a.x, b.x) - min(a.x, b.x)
	dy := max(a.y, b.y) - min(a.y, b.y)

	return dx + dy
}

func charAt(pos Pos) rune {
	x, y := pos.x, pos.y

	if x >= W || y >= H || x < 0 || y < 0 {
		return -1
	}

	return rune(input[y*(W+1)+x])
}

func main() {
	perf.Bench(1, solution)
}
