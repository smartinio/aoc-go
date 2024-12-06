package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Dir struct{ x, y int }

//go:embed input.txt
var input string
var w int = strings.Index(input, "\n")
var h int = strings.Count(input, "\n")
var dirs []Dir = []Dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
var dirlen int = len(dirs)

func main() {
	visited := make(map[string]int)
	sx, sy := findStart()
	x, y := sx, sy
	dir := 0
	loops := 0

	for x >= 0 && y >= 0 && x < w && y < h {
		k := key(x, y)
		_, didVisit := visited[k]

		if !didVisit {
			visited[k] = dir
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

	for k, dir := range visited {
		if k == key(sx, sy) {
			continue
		}

		visitedDir := make(map[string]bool)
		kx := strings.Split(k, ":")
		bx, _ := strconv.Atoi(kx[0])
		by, _ := strconv.Atoi(kx[1])

		d := dirs[dir]
		x, y = bx-d.x, by-d.y

		for x >= 0 && y >= 0 && x < w && y < h {
			dk := dirKey(x, y, dir)

			if visitedDir[dk] {
				loops++
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
	}

	fmt.Println("part1:", len(visited))
	fmt.Println("part2:", loops)
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

func key(x int, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func dirKey(x int, y int, dir int) string {
	return fmt.Sprintf("%d:%d:%d", x, y, dir)
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
