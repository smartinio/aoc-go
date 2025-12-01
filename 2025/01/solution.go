package main

import (
	"bufio"
	_ "embed"
	"main/perf"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0

	pos := 50
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		clicks, _ := strconv.Atoi(line[1:])
		laps := int(clicks / 100)
		steps := clicks % 100
		start := pos

		switch dir {
		case 'L':
			pos -= steps
		case 'R':
			pos += steps
		}

		next := wrap(pos, 100)

		// ends on a 0
		if next == 0 {
			part1++
			part2++
		}

		// passes by a 0
		if start != 0 && (pos < 0 || pos > 100) {
			part2++
		}

		// passes by a 0 multiple times
		part2 += laps

		pos = next
	}

	return part1, part2
}

func wrap(i, max int) int {
	return ((i % max) + max) % max
}

func main() {
	perf.Bench(1, solution)
}
