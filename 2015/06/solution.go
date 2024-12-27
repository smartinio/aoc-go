package main

import (
	_ "embed"
	"main/perf"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0

	lines := strings.Split(strings.TrimSpace(input), "\n")
	re := regexp.MustCompile(`\d+`)

	// part 1
	{
		grid := make([][]bool, 1000)
		for i := range grid {
			grid[i] = make([]bool, 1000)
		}

		for _, line := range lines {
			c := re.FindAllString(line, -1)
			x1, y1, x2, y2 := atoi(c[0]), atoi(c[1]), atoi(c[2]), atoi(c[3])

			for y := y1; y <= y2; y++ {
				for x := x1; x <= x2; x++ {
					if strings.HasPrefix(line, "turn on") {
						grid[y][x] = true
					}
					if strings.HasPrefix(line, "turn off") {
						grid[y][x] = false
					}
					if strings.HasPrefix(line, "toggle") {
						grid[y][x] = !grid[y][x]
					}
				}
			}
		}

		for y := range grid {
			for _, on := range grid[y] {
				if on {
					part1++
				}
			}
		}
	}

	// part 2
	{
		grid := make([][]int, 1000)
		for i := range grid {
			grid[i] = make([]int, 1000)
		}

		for _, line := range lines {
			c := re.FindAllString(line, -1)
			x1, y1, x2, y2 := atoi(c[0]), atoi(c[1]), atoi(c[2]), atoi(c[3])

			for y := y1; y <= y2; y++ {
				for x := x1; x <= x2; x++ {
					if strings.HasPrefix(line, "turn on") {
						grid[y][x] += 1
					}
					if strings.HasPrefix(line, "turn off") {
						grid[y][x] = max(0, grid[y][x]-1)
					}
					if strings.HasPrefix(line, "toggle") {
						grid[y][x] += 2
					}
				}
			}
		}

		for y := range grid {
			for _, brightness := range grid[y] {
				part2 += brightness
			}
		}
	}

	return part1, part2
}

func atoi(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func main() {
	perf.Bench(1, solution)
}
