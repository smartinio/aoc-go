package main

import (
	_ "embed"
	"main/perf"
	"slices"
	"strconv"
	"strings"
)

type Pos struct{ x, y int }
type Tiles map[Pos]bool
type PosLookup map[int][]Pos
type Dangers struct{ x, y PosLookup }

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	reds := make([]Pos, len(lines))

	for i, line := range lines {
		d := strings.Split(line, ",")
		x, _ := strconv.Atoi(d[0])
		y, _ := strconv.Atoi(d[1])

		reds[i] = Pos{x, y}
	}

	// part 1
	for i := range reds {
		for j := i + 1; j < len(reds); j++ {
			a, b := reds[i], reds[j]
			part1 = max(part1, area(a, b))
		}
	}

	// part 2
	dangers := perimeter(reds)

	for i := range reds {
		for j := i + 1; j < len(reds); j++ {
			a, b := reds[i], reds[j]

			if !crosscheck(a, b, dangers) {
				continue
			}

			part2 = max(part2, area(a, b))
		}
	}

	return part1, part2
}

func crosscheck(a, b Pos, dangers Dangers) bool {
	c, d := Pos{a.x, b.y}, Pos{b.x, a.y}

	for _, p1 := range []Pos{a, b} {
		for _, p2 := range []Pos{c, d} {
			if !clear(p1, p2, dangers) {
				return false
			}
		}
	}

	return true
}

func perimeter(reds []Pos) Dangers {
	inner, outer := make(Tiles), make(Tiles)
	dangers := Dangers{make(PosLookup), make(PosLookup)}

	for i := range reds {
		a, b := reds[i], reds[(i+1)%len(reds)]
		diffx, diffy := b.x-a.x, b.y-a.y
		dx, dy := normalize(diffx), normalize(diffy)
		steps := abs(diffx) + abs(diffy)

		for s := range steps {
			p := Pos{a.x + dx*s, a.y + dy*s}
			op := Pos{p.x + dy, p.y}
			inner[p] = true

			if dy == 0 {
				op = Pos{p.x, p.y - dx}
			}

			outer[op] = !inner[op]
		}
	}

	for k, danger := range outer {
		if danger {
			dangers.x[k.x] = append(dangers.x[k.x], k)
			dangers.y[k.y] = append(dangers.y[k.y], k)
		}
	}

	return dangers
}

func clear(a, b Pos, dangers Dangers) bool {
	if a.x == b.x {
		start, end := min(a.y, b.y), max(a.y, b.y)
		outers := dangers.x[a.x]

		return !slices.ContainsFunc(outers, func(out Pos) bool {
			return out.y >= start && out.y <= end
		})
	}

	start, end := min(a.x, b.x), max(a.x, b.x)
	outers := dangers.y[a.y]

	return !slices.ContainsFunc(outers, func(out Pos) bool {
		return out.x >= start && out.x <= end
	})
}

func area(a, b Pos) int {
	dx := 1 + abs(a.x-b.x)
	dy := 1 + abs(a.y-b.y)

	return dx * dy
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func normalize(a int) int {
	return min(1, max(-1, a))
}

func main() {
	perf.Bench(1, solution)
}
