package main

import (
	_ "embed"
	"main/perf"
	"main/utils"
	"regexp"
)

//go:embed input.txt
var input string

func solution() (part1 int, part2 int) {
	re := regexp.MustCompile(`cpy (\d+) (?:c|b)`)
	factors := utils.FindAllIntGroups(re, input)[:2]
	salt := factors[0] * factors[1]

	n := 0
	for i := 0; true; i++ {
		digit := i % 2
		n = (digit << i) | n
		part1 = n - salt

		if part1 > 0 {
			break
		}
	}

	return part1, 1337
}

func main() {
	perf.Bench(1, solution)
}
