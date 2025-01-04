package main

import (
	_ "embed"
	"main/perf"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Reg map[string]uint

func solution() (part1 uint, part2 uint) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// part 1
	{
		reg := map[string]uint{}
		compute(reg, lines)
		part1 = reg["b"]
	}

	// part 2
	{
		reg := map[string]uint{"a": 1}
		compute(reg, lines)
		part2 = reg["b"]
	}

	return part1, part2
}

func compute(reg Reg, lines []string) {
	ptr := 0

	for {
		if ptr < 0 || ptr >= len(lines) {
			break
		}

		line := strings.Replace(lines[ptr], ",", "", -1)
		a := strings.Split(line, " ")

		switch a[0] {
		case "hlf":
			reg[a[1]] /= 2
		case "tpl":
			reg[a[1]] *= 3
		case "inc":
			reg[a[1]]++
		case "jmp":
			offset, _ := strconv.Atoi(a[1])
			ptr += offset
			continue
		case "jie":
			if reg[a[1]]%2 == 0 {
				offset, _ := strconv.Atoi(a[2])
				ptr += offset
				continue
			}
		case "jio":
			if reg[a[1]] == 1 {
				offset, _ := strconv.Atoi(a[2])
				ptr += offset
				continue
			}
		}

		ptr++
	}
}

func main() {
	perf.Bench(1, solution)
}
