package main

import (
	_ "embed"
	"main/perf"
	"math"
	"slices"
	"strings"
)

type Pos struct{ x, y int }
type Dir struct{ x, y int }
type Key [8]rune
type Pair [2]Pos

//go:embed input.txt
var input string
var W = strings.Index(input, "\n")
var H = strings.Count(input, "\n")

func solution() (part1 float64, part2 float64) {
	start, points := parseInput()

	dists := map[Pair]float64{}

	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			dist := bfs(p1, p2)
			dists[Pair{p1, p2}] = dist
			dists[Pair{p2, p1}] = dist
		}
	}

	part1 = dfs(0, []Pos{start}, points, dists, map[Key]float64{}, 1)
	part2 = dfs(0, []Pos{start}, points, dists, map[Key]float64{}, 2)

	return part1, part2
}

func dfs(steps float64, path, points []Pos, dists map[Pair]float64, cache map[Key]float64, part int) float64 {
	curr := path[len(path)-1]

	if len(path) == len(points) {
		if part == 2 {
			return steps + bfs(curr, path[0])
		}

		return steps
	}

	key := makeKey(path)

	if cached, ok := cache[key]; ok {
		return steps + cached
	}

	low := math.Inf(1)

	for _, next := range points {
		if !slices.Contains(path, next) {
			dist := dists[Pair{curr, next}]
			low = min(low, dfs(dist, append(path, next), points, dists, cache, part))
		}
	}

	cache[key] = low

	return steps + low
}

func makeKey(visited []Pos) (key Key) {
	current := visited[len(visited)-1]
	key[0] = charAt(current)
	for i := 0; i < len(visited); i++ {
		key[i+1] = charAt(visited[i])
	}
	slices.Sort(key[1:])
	return
}

func bfs(start, end Pos) float64 {
	q := []Pos{start}
	steps := map[Pos]float64{start: 0}
	dirs := []Dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr == end {
			return steps[curr]
		}

		for _, d := range dirs {
			next := Pos{curr.x + d.x, curr.y + d.y}

			if _, visited := steps[next]; !visited && charAt(next) != '#' {
				q = append(q, next)
				steps[next] = steps[curr] + 1
			}
		}
	}

	return math.Inf(1)
}

func charAt(pos Pos) rune {
	x, y := pos.x, pos.y

	if x >= W || y >= H || x < 0 || y < 0 {
		return '#'
	}

	return rune(input[y*(W+1)+x])
}

func parseInput() (start Pos, points []Pos) {
	for y := range H {
		for x := range W {
			char := charAt(Pos{x, y})
			if char >= '0' && char <= '9' {
				points = append(points, Pos{x, y})
			}
			if char == '0' {
				start = Pos{x, y}
			}
		}
	}
	return
}

func main() {
	perf.Bench(1, solution)
}
