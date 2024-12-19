package main

import (
	"main/perf"
	"os"
	"regexp"
	"strconv"
)

func solution() (int, int) {
	file, _ := os.ReadFile("2024-03/input.txt")

	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	subs := re.FindAllStringSubmatch(string(file), -1)

	part1, part2 := 0, 0
	enabled := true

	for _, sub := range subs {
		switch sub[1] {
		case "":
			enabled = sub[0] == "do()"
		default:
			a, _ := strconv.Atoi(sub[1])
			b, _ := strconv.Atoi(sub[2])

			part1 += a * b

			if enabled {
				part2 += a * b
			}
		}
	}

	return part1, part2
}

func main() {
	perf.Bench(100, solution)
}
