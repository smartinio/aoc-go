package main

import (
	_ "embed"
	"main/perf"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Box struct {
	pos   [3]float64
	nodes []*Box
}

type Pair struct {
	a, b *Box
	dist float64
}

func solution() (int, int) {
	part1, part2 := 1, 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	upper := min(len(lines), 1000)
	boxes := make([]*Box, len(lines))
	pairs := []Pair{}

	for i, line := range lines {
		p := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(p[0], 64)
		y, _ := strconv.ParseFloat(p[1], 64)
		z, _ := strconv.ParseFloat(p[2], 64)
		boxes[i] = &Box{[3]float64{x, y, z}, []*Box{}}
	}

	for i, a := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			b := boxes[j]
			pairs = append(pairs, Pair{a, b, dist(a, b)})
		}
	}

	slices.SortFunc(pairs, func(a, b Pair) int {
		return int(a.dist - b.dist)
	})

	// part 1
	for _, pair := range pairs[0:upper] {
		visited := make(map[*Box]bool)

		if !connected(pair.a, pair.b, visited) {
			connect(pair)
		}
	}

	sizes := []int{}
	visited := make(map[*Box]bool)

	for _, pair := range pairs[0:upper] {
		if !visited[pair.a] {
			size := circsize(pair.a, visited)

			if size > 0 {
				sizes = append(sizes, size)
			}
		}
	}

	slices.Sort(sizes)

	for i := range 3 {
		part1 *= sizes[len(sizes)-1-i]
	}

	// part 2
	for _, pair := range pairs[upper:] {
		visited := make(map[*Box]bool)

		if !connected(pair.a, pair.b, visited) {
			connect(pair)
			size := circsize(pair.a, make(map[*Box]bool))

			if size == len(boxes) {
				part2 = int(pair.a.pos[0] * pair.b.pos[0])
				break
			}
		}
	}

	return part1, part2
}

func circsize(box *Box, visited map[*Box]bool) int {
	q := []*Box{box}
	prevLen := len(visited)

	for len(q) > 0 {
		b := q[0]
		visited[b] = true

		for _, node := range b.nodes {
			if !visited[node] {
				q = append(q, node)
			}
		}

		q = q[1:]
	}

	return len(visited) - prevLen
}

func connected(a, b *Box, visited map[*Box]bool) bool {
	for _, box := range a.nodes {
		if !visited[box] {
			visited[box] = true

			if box == b || connected(box, b, visited) {
				return true
			}
		}
	}

	return false
}

func connect(pair Pair) {
	pair.a.nodes = append(pair.a.nodes, pair.b)
	pair.b.nodes = append(pair.b.nodes, pair.a)
}

func dist(p *Box, q *Box) float64 {
	sum := 0.0

	for i := range 3 {
		sum += math.Pow(p.pos[i]-q.pos[i], 2)
	}

	return math.Sqrt(sum)
}

func main() {
	perf.Bench(1, solution)
}
