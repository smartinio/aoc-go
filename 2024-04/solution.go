package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string
var w int = strings.Index(input, "\n")
var h int = strings.Count(input, "\n")

func main() {
	part1, part2 := 0, 0

	for x := range w {
		for y := range h {
			char := charAt(x, y)

			if char == 'X' {
				part1 += countMASFrom(x, y)
			}

			if char == 'A' && isBetweenMultipleMS(x, y) {
				part2 += 1
			}
		}
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func charAt(x int, y int) rune {
	if x >= w || y >= h || x < 0 || y < 0 {
		return 0
	}

	return rune(input[int(y)*int(w+1)+int(x)])
}

func countMASFrom(x int, y int) int {
	count := 0
	dirs := []int{0, 1, -1}

	for _, dx := range dirs {
		for _, dy := range dirs {
			if !(dx == 0 && dy == 0) && dirSpells(x, y, dx, dy, "MAS") {
				count += 1
			}
		}
	}

	return count
}

func isBetweenMultipleMS(x int, y int) bool {
	dirs := []int{1, -1}
	match := false

	for _, dx := range dirs {
		for _, dy := range dirs {
			if dirSpells(x, y, dx, dy, "M") && dirSpells(x, y, -dx, -dy, "S") {
				if match {
					return true
				}
				match = true
			}
		}
	}

	return false
}

func dirSpells(x int, y int, dx int, dy int, word string) bool {
	for i, char := range word {
		if charAt(x+(i+1)*dx, y+(i+1)*dy) != char {
			return false
		}
	}

	return true
}
