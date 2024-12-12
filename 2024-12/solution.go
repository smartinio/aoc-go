package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"
)

type Pos struct{ x, y int }
type Region struct{ area, perimeter, sides int }
type Side struct{ dir, sx, sy, ex, ey int }

//go:embed input.txt
var input string
var H = strings.Count(input, "\n")
var W = strings.Index(input, "\n")
var NOT_FOUND = Pos{-1, -1}
var DIRS = []Pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func solution() (int, int) {
	part1, part2 := 0, 0
	unvisited := make(map[Pos]bool)

	for x := range W {
		for y := range H {
			unvisited[Pos{x, y}] = true
		}
	}

	for {
		pos := getPos(unvisited)

		if pos == NOT_FOUND {
			break
		}

		region := mapRegion(pos, unvisited, make(map[Side]bool))

		part1 += region.area * region.perimeter
		part2 += region.area * region.sides
	}

	return part1, part2
}

func getPos(unvisited map[Pos]bool) Pos {
	for k := range unvisited {
		return k
	}
	return NOT_FOUND
}

func mapRegion(pos Pos, unvisited map[Pos]bool, sides map[Side]bool) Region {
	delete(unvisited, pos)

	neighbors := getNeighbors(pos)
	sideDirs := getSideDirs(pos)
	perimeter := len(sideDirs)
	area := 1

	for _, sideDir := range sideDirs {
		di := slices.Index(DIRS, sideDir)

		rightward := DIRS[(di+1)%len(DIRS)]
		leftward := DIRS[(di+3)%len(DIRS)]

		start := walk(leftward, pos, sideDir)
		end := walk(rightward, pos, sideDir)

		sides[Side{di, start.x, start.y, end.x, end.y}] = true
	}

	for _, neighbor := range neighbors {
		if _, ok := unvisited[neighbor]; ok {
			next := mapRegion(neighbor, unvisited, sides)
			area += next.area
			perimeter += next.perimeter
		}
	}

	return Region{area, perimeter, len(sides)}
}

func walk(d Pos, pos Pos, sideDir Pos) Pos {
	movedPos := pos
	movedSide := Pos{pos.x + sideDir.x, pos.y + sideDir.y}

	for {
		next := Pos{movedPos.x + d.x, movedPos.y + d.y}
		nextSide := Pos{movedSide.x + d.x, movedSide.y + d.y}

		if charAt(next) == charAt(pos) && charAt(nextSide) != charAt(pos) {
			movedPos = next
			movedSide = nextSide
		} else {
			break
		}
	}

	return movedPos
}

func getSideDirs(pos Pos) []Pos {
	char := charAt(pos)
	sideDirs := []Pos{}

	for _, dir := range DIRS {
		if charAt(Pos{pos.x + dir.x, pos.y + dir.y}) != char {
			sideDirs = append(sideDirs, dir)
		}
	}

	return sideDirs
}

func getNeighbors(pos Pos) []Pos {
	char := charAt(pos)
	neighbors := []Pos{}

	for _, d := range DIRS {
		neighbor := Pos{pos.x + d.x, pos.y + d.y}
		nc := charAt(neighbor)

		if nc != -1 && nc == char {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func charAt(pos Pos) rune {
	x, y := pos.x, pos.y

	if x >= W || y >= H || x < 0 || y < 0 {
		return -1
	}

	return rune(input[y*(W+1)+x])
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
