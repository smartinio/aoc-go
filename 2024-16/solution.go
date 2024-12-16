package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/x1m3/priorityQueue"
)

type Pos struct{ x, y int }
type Dir struct{ x, y int }
type Vertex struct{ x, y, dir int }
type Move struct {
	from  Vertex
	to    Vertex
	score int
}

//go:embed input.txt
var input string
var W = strings.Index(input, "\n")
var H = strings.Count(input, "\n")
var dirs = []Dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
var EAST = 1 // index in ^
var scores = map[Vertex]int{}

func solution() (int, int) {
	scores = map[Vertex]int{}
	start := Vertex{}

	for y := range H {
		for x := range W {
			if charAt(Pos{x, y}) == 'S' {
				start = Vertex{x, y, EAST}
			}
		}
	}

	part1, part2 := dijkstra(start)

	return part1, part2
}

func (i Vertex) HigherPriorityThan(other priorityQueue.Interface) bool {
	return scores[i] < scores[other.(Vertex)]
}

func dijkstra(source Vertex) (int, int) {
	pq := priorityQueue.New()
	prevs := map[Vertex][]Vertex{}

	enqueue := func(vertex Vertex, score int) {
		scores[vertex] = score
		pq.Push(vertex)
	}

	check := func(m Move) {
		curr, next, score := m.from, m.to, m.score
		newScore := scores[curr] + score

		if _, ok := scores[next]; !ok || scores[next] > newScore {
			enqueue(next, newScore)
			prevs[next] = []Vertex{curr}
		} else if scores[next] == newScore {
			prevs[next] = append(prevs[next], curr)
		}
	}

	enqueue(source, 0)

	for {
		r := pq.Pop()

		if r == nil {
			break
		}

		curr := r.(Vertex)
		pos := Pos{curr.x, curr.y}

		if charAt(pos) == 'E' {
			return scores[curr], countSeats(curr, prevs)
		}

		// forward
		{
			d := dirs[curr.dir]
			next := Vertex{curr.x + d.x, curr.y + d.y, curr.dir}
			if charAt(Pos{next.x, next.y}) != '#' {
				check(Move{from: curr, to: next, score: 1})
			}
		}

		// rotate 90 deg
		for _, rotation := range []int{1, -1} {
			dir := wrap(curr.dir+rotation, len(dirs))
			next := Vertex{curr.x, curr.y, dir}
			check(Move{from: curr, to: next, score: 1000})
		}
	}

	return -1, 0
}

func countSeats(target Vertex, prevs map[Vertex][]Vertex) int {
	seats := map[Pos]bool{}
	q := []Vertex{target}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		seats[Pos{curr.x, curr.y}] = true
		q = append(q, prevs[curr]...)
	}

	return len(seats)
}

func charAt(pos Pos) rune {
	x, y := pos.x, pos.y

	if x >= W || y >= H || x < 0 || y < 0 {
		return -1
	}

	return rune(input[y*(W+1)+x])
}

func wrap(i, max int) int {
	return ((i % max) + max) % max
}

func main() {
	part1, part2 := 0, 0
	sum := 0
	n := 25 // increase samples if benching perf

	for range n {
		start := time.Now()
		part1, part2 = solution()
		sum += int(time.Since(start).Milliseconds())
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
	fmt.Println("avg:", sum/n, "ms")
}
