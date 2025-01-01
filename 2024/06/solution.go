package main

import (
	_ "embed"
	"main/perf"
	"strings"
	"sync"
)

type Dir struct{ x, y int }
type Pos struct{ x, y int }
type PosDir struct{ x, y, dir int }
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string, value int) {
	c.mu.Lock()
	c.v[key] += value
	c.mu.Unlock()
}

//go:embed input.txt
var input string
var W = strings.Index(input, "\n")
var H = strings.Count(input, "\n")
var dirs []Dir = []Dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
var dirlen = len(dirs)

func solution() (int, int) {
	visited := make(map[Pos]int)
	start := findStart()
	curr := start
	dir := 0

outer:
	for {
		_, didVisit := visited[curr]

		if !didVisit {
			visited[curr] = dir
		}

		for {
			next := getNext(curr, dir)
			nextChar := charAt(next)

			if nextChar == 0 {
				break outer
			}

			if nextChar != '#' {
				curr = next
				break
			}

			dir = (dir + 1) % dirlen
		}
	}

	part2 := SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup

	for k, dir := range visited {
		if k == start {
			continue
		}

		wg.Add(1)
		go func(k Pos, dir int) {
			defer wg.Done()

			visitedPosDirs := make(map[PosDir]bool)
			block := Pos{k.x, k.y}
			d := dirs[dir]
			curr := Pos{block.x - d.x, block.y - d.y}

		outer:
			for {
				dk := PosDir{curr.x, curr.y, dir}

				if visitedPosDirs[dk] {
					part2.Inc("loops", 1)
					break
				}

				visitedPosDirs[dk] = true

				for {
					next := getNext(curr, dir)
					nextChar := charAt(next)

					if nextChar == 0 {
						break outer
					}

					if next != block && nextChar != '#' {
						curr = next
						break
					}

					dir = (dir + 1) % dirlen
				}
			}
		}(k, dir)
	}

	wg.Wait()

	return len(visited), part2.v["loops"]
}

func getNext(pos Pos, dir int) Pos {
	d := dirs[dir]

	return Pos{pos.x + d.x, pos.y + d.y}
}

func charAt(pos Pos) rune {
	x, y := pos.x, pos.y

	if x >= W || y >= H || x < 0 || y < 0 {
		return 0
	}

	return rune(input[y*(W+1)+x])
}

func findStart() Pos {
	for x := range W {
		for y := range H {
			if charAt(Pos{x, y}) == '^' {
				return Pos{x, y}
			}
		}
	}
	return Pos{}
}

func main() {
	perf.Bench(10, solution)
}
