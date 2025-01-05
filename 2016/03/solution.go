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

func solution() (part1 int, part2 int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	re := regexp.MustCompile(`\d+`)

	// part 1
	for _, line := range lines {
		n := re.FindAllString(line, 3)
		a, b, c := atoi(n[0]), atoi(n[1]), atoi(n[2])

		if a+b > c && b+c > a && c+a > b {
			part1++
		}
	}

	// part 2
	for i := 0; i < len(lines)-2; i += 3 {
		n := re.FindAllString(lines[i], 3)
		a1, a2, a3 := atoi(n[0]), atoi(n[1]), atoi(n[2])
		n = re.FindAllString(lines[i+1], 3)
		b1, b2, b3 := atoi(n[0]), atoi(n[1]), atoi(n[2])
		n = re.FindAllString(lines[i+2], 3)
		c1, c2, c3 := atoi(n[0]), atoi(n[1]), atoi(n[2])

		if a1+b1 > c1 && b1+c1 > a1 && c1+a1 > b1 {
			part2++
		}

		if a2+b2 > c2 && b2+c2 > a2 && c2+a2 > b2 {
			part2++
		}

		if a3+b3 > c3 && b3+c3 > a3 && c3+a3 > b3 {
			part2++
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
