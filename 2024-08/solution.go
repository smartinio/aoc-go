package main

import (
	_ "embed"
	"main/perf"
	"strings"
)

//go:embed input.txt
var input string
var H = strings.Count(input, "\n")
var W = strings.Index(input, "\n")

type Loc struct{ x, y int }

func solution() (int, int) {
	towers := make(map[rune][]Loc)
	p1nodes := make(map[Loc]bool)
	p2nodes := make(map[Loc]bool)

	for x := range W {
		for y := range H {
			loc := Loc{x, y}
			key := charAt(loc)

			if key != '.' {
				towers[key] = append(towers[key], loc)
			}
		}
	}

	for _, locs := range towers {
		for i, a := range locs {
			for _, b := range locs[i+1:] {
				dx, dy := a.x-b.x, a.y-b.y
				dirs := []int{1, -1}
				curr := a

				for _, dir := range dirs {
					// part 1
					antis := []Loc{
						{a.x + dx*dir, a.y + dy*dir},
						{b.x + dx*dir, b.y + dy*dir},
					}

					for _, anti := range antis {
						if anti != a && anti != b && isOnMap(anti) {
							p1nodes[anti] = true
						}
					}

					// part 2
					for {
						curr.x += dx * dir
						curr.y += dy * dir

						if !isOnMap(curr) {
							break
						}

						p2nodes[curr] = true
					}
				}
			}
		}
	}

	return len(p1nodes), len(p2nodes)
}

func isOnMap(loc Loc) bool {
	x, y := loc.x, loc.y

	return x >= 0 && y >= 0 && x < W && y < H
}

func charAt(pos Loc) rune {
	x, y := pos.x, pos.y

	if x >= W || y >= H || x < 0 || y < 0 {
		return -1
	}

	return rune(input[y*(W+1)+x])
}

func main() {
	perf.Bench(100, solution)
}
