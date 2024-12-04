package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := parseInput()
	part1, part2 := 0, 0

	for y, line := range lines {
		for x, char := range line {
			if char == 'X' {
				part1 += countMASFrom(x, y, lines)
			}

			if char == 'A' && isBetweenMultipleMS(x, y, lines) {
				part2 += 1
			}
		}
	}

	fmt.Println("xmas count:", part1)
	fmt.Println("x-mas count:", part2)
}

func countMASFrom(x int, y int, lines []string) int {
	count := 0
	dirs := []int{0, 1, -1}
	mas := "MAS"

	dirSpellsMAS := func(dx int, dy int) bool {
		for i, char := range mas {
			k := i + 1
			dxx, dyy := x+k*dx, y+k*dy

			if dyy < 0 || dxx < 0 || dyy > len(lines)-1 || dxx > len(lines[y])-1 || lines[dyy][dxx] != byte(char) {
				return false
			}
		}

		return true
	}

	for _, dx := range dirs {
		for _, dy := range dirs {
			if (dx != 0 || dy != 0) && dirSpellsMAS(dx, dy) {
				count += 1
			}
		}
	}

	return count
}

func isBetweenMultipleMS(x int, y int, lines []string) bool {
	if x == 0 || x == len(lines[y])-1 || y == 0 || y == len(lines)-1 {
		return false
	}

	dirs := []int{1, -1}
	match := false

	for _, dx := range dirs {
		for _, dy := range dirs {
			if lines[y+dy][x+dx] == 'M' && lines[y-dy][x-dx] == 'S' {
				if match {
					return true
				}
				match = true
			}
		}
	}

	return false
}

func parseInput() []string {
	file, _ := os.Open("2024-04/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
