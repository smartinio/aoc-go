package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Pos struct{ x, y int }
type Dir struct{ x, y int }
type Vertex struct {
	pos   Pos
	steps int
}

const W, H = 71, 71

//go:embed input.txt
var input string

func solution() (int, Pos) {
	corrupt, bytes := parseInput()

	part1 := bfs(corrupt, 1024)

	idx, _ := slices.BinarySearchFunc(bytes, -1, func(a Pos, target int) int {
		return target - bfs(corrupt, corrupt[a])
	})

	part2 := bytes[idx]

	return part1, part2
}

func bfs(corrupted map[Pos]int, time int) int {
	start, end := Pos{0, 0}, Pos{W - 1, H - 1}
	dirs := []Dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	q := []Vertex{{start, 0}}
	visited := map[Pos]bool{start: true}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.pos == end {
			return curr.steps
		}

		for _, d := range dirs {
			nextPos := Pos{curr.pos.x + d.x, curr.pos.y + d.y}

			if !blocked(nextPos, corrupted, time) && !visited[nextPos] {
				q = append(q, Vertex{nextPos, curr.steps + 1})
				visited[nextPos] = true
			}
		}
	}

	return -1
}

func blocked(pos Pos, corrupt map[Pos]int, time int) bool {
	x, y := pos.x, pos.y

	if x >= W || y >= H || x < 0 || y < 0 {
		return true
	}

	corruptedAt, corrupted := corrupt[pos]

	return corrupted && corruptedAt <= time
}

func parseInput() (map[Pos]int, []Pos) {
	corrupt, bytes := map[Pos]int{}, []Pos{}
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for i, line := range lines {
		time := i + 1
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		corrupt[Pos{x, y}] = time
		bytes = append(bytes, Pos{x, y})
	}

	return corrupt, bytes
}

func main() {
	part1, part2 := 0, Pos{}
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
