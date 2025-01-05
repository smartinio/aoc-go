package main

import (
	_ "embed"
	"main/perf"
	"regexp"
	"strconv"
	"strings"
)

type Screen = [][]bool

//go:embed input.txt
var input string

const W, H = 50, 6

func solution() (part1 int, part2 string) {
	screen := make(Screen, H)
	for i := range screen {
		screen[i] = make([]bool, W)
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	re := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		a := strings.Split(line, " ")
		n := re.FindAllString(line, -1)
		arg1, _ := strconv.Atoi(n[0])
		arg2, _ := strconv.Atoi(n[1])

		switch {
		case a[0] == "rect":
			screen = rect(screen, arg1, arg2)
		case a[0] == "rotate":
			if a[1] == "row" {
				screen = rotateRow(screen, arg1, arg2)
			} else {
				screen = rotateCol(screen, arg1, arg2)
			}
		}
	}

	for i := range H {
		for j := range W {
			if screen[i][j] {
				part1++
			}
		}
	}

	part2 = render(screen)

	return part1, part2
}

func rect(screen Screen, x, y int) Screen {
	for i := range y {
		for j := range x {
			screen[i][j] = true
		}
	}

	return screen
}

func rotateRow(screen Screen, y, d int) Screen {
	new := make([]bool, len(screen[y]))

	for x, pixel := range screen[y] {
		nx := wrap(x+d, len(new))
		new[nx] = pixel
	}

	screen[y] = new

	return screen
}

func rotateCol(screen Screen, x, d int) Screen {
	new := make([]bool, len(screen))

	for y := 0; y < len(screen); y++ {
		ny := wrap(y+d, len(new))
		new[ny] = screen[y][x]
	}

	for y, pixel := range new {
		screen[y][x] = pixel
	}

	return screen
}

func render(screen Screen) string {
	sb := strings.Builder{}
	sb.WriteString("\n")

	for y := range H {
		for x := range W {
			if screen[y][x] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func wrap(i, max int) int {
	return ((i % max) + max) % max
}

func main() {
	perf.Bench(1, solution)
}
