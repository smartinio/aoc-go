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
	part1 = splits(sx, sy, make(Seen))
	part2 = spawn(Key{sx, sy}, make(Cache))

	return part1, part2
}

func splits(x int, sy int, seen Seen) int {
	for y := sy; y < h; y++ {
		k := Key{x, y}

		if seen[k] {
			break
		}

		if charAt(x, y) == '^' {
			seen[k] = true
			return 1 + splits(x+1, y, seen) + splits(x-1, y, seen)
		}
	}

	return 0
}

func spawn(k Key, cache Cache) int {
	if cache[k] != 0 {
		return cache[k]
	}

	for y := k.y; y < h; y++ {
		if charAt(k.x, y) == '^' {
			a, b := Key{k.x + 1, y}, Key{k.x - 1, y}
			cache[a] = spawn(a, cache)
			cache[b] = spawn(b, cache)
			return cache[a] + cache[b]
		}
	}

	return 1
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
