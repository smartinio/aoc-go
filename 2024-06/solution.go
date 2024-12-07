package main

import (
	_ "embed"
	"fmt"
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
	visited := make(map[Pos]int)
	sx, sy := findStart()
	x, y := sx, sy
	dir := 0
	c := SafeCounter{v: make(map[string]int)}

	for x >= 0 && y >= 0 && x < w && y < h {
		key := Pos{x, y}
		_, didVisit := visited[key]

		if !didVisit {
			visited[key] = dir
		}

		for {
			nx, ny := getNext(x, y, dir)
			next := charAt(nx, ny)

			if next != '#' {
				x, y = nx, ny
				break
			}

			dir = (dir + 1) % dirlen
		}
	}

	var wg sync.WaitGroup

	for k, dir := range visited {
		if k == (Pos{sx, sy}) {
			continue
		}

		wg.Add(1)

		go func(k Pos, dir int) {
			defer wg.Done()

			visitedDir := make(map[PosDir]bool)

			bx := k.x
			by := k.y

			d := dirs[dir]
			x, y := bx-d.x, by-d.y

			sum := 0

			for x >= 0 && y >= 0 && x < w && y < h {
				dk := PosDir{x, y, dir}

				if visitedDir[dk] {
					sum++
					break
				}

				visitedDir[dk] = true

				for {
					nx, ny := getNext(x, y, dir)
					next := charAt(nx, ny)

					if next != '#' && !(nx == bx && ny == by) {
						x, y = nx, ny
						break
					}

					dir = (dir + 1) % dirlen
				}
			}

			c.Inc("loops", sum)
		}(k, dir)
	}

	wg.Wait()

	fmt.Println("part1:", len(visited))
	fmt.Println("part2:", c.v["loops"])
}

func getNext(x int, y int, dir int) (int, int) {
	d := dirs[dir]
	nx, ny := x+d.x, y+d.y
	return nx, ny
}

func charAt(x int, y int) rune {
	if x >= w || y >= h || x < 0 || y < 0 {
		return 0
	}

	return rune(input[int(y)*int(w+1)+int(x)])
}

func findStart() (int, int) {
	for x := range w {
		for y := range h {
			if charAt(x, y) == '^' {
				return x, y
			}
		}
	}

	return 0, 0
}
