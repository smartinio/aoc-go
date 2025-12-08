package main

import (
	_ "embed"
	"main/perf"
	"strings"
)

type Key = struct{ x, y int }
type Seen = map[Key]bool
type Cache = map[Key]int

//go:embed input.txt
var input string
var w = strings.Index(input, "\n")
var h = strings.Count(input, "\n") - 1

func solution() (int, int) {
	part1, part2 := 0, 0
	sy, sx := 0, strings.Index(input, "S")
	cache := make(Cache)
	part2 = beam(sx, sy, cache)
	part1 = len(cache) // only caching on split

	return part1, part2
}

func beam(x int, y int, cache Cache) int {
	if y >= h {
		return 1
	}

	k := Key{x, y}

	if cache[k] != 0 {
		return cache[k]
	}

	if charAt(x, y) == '^' {
		cache[k] = beam(x+1, y, cache) + beam(x-1, y, cache)
		return cache[k]
	}

	return beam(x, y+1, cache)
}

func charAt(x int, y int) byte {
	if x >= w || y >= h || x < 0 || y < 0 {
		return 0
	}

	return input[y*(w+1)+x]
}

func main() {
	perf.Bench(1, solution)
}
