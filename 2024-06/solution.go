package main

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"
	"time"
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
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key] += value
	c.mu.Unlock()
}

//go:embed input.txt
var input string
var w int = strings.Index(input, "\n")
var h int = strings.Count(input, "\n")
var dirs []Dir = []Dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
var dirlen int = len(dirs)

func main() {
	part1, part2 := 0, 0
	sum := 0
	n := 1 // increase samples if benching perf

	for range n {
		start := time.Now()
		part1, part2 = solution()
		sum += int(time.Since(start).Milliseconds())
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
	fmt.Println("avg:", sum/n, "ms")
}

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

	if x >= w || y >= h || x < 0 || y < 0 {
		return 0
	}

	return rune(input[int(y)*int(w+1)+int(x)])
}

func findStart() Pos {
	for x := range w {
		for y := range h {
			if charAt(Pos{x, y}) == '^' {
				return Pos{x, y}
			}
		}
	}
	return Pos{}
}
