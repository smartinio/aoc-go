package main

import (
	_ "embed"
	"main/perf"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func solution() (part1 int, part2 int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// part 1
	{
		regs := map[string]int{}
		compute(lines, regs)
		part1 = regs["a"]
	}

	// part 2
	{
		regs := map[string]int{"c": 1}
		compute(lines, regs)
		part2 = regs["a"]
	}

	return part1, part2
}

func compute(lines []string, regs map[string]int) {
	ptr := 0

	for ptr >= 0 && ptr < len(lines) {
		a := strings.Split(lines[ptr], " ")
		x := getX(a[1], regs)

		switch a[0] {
		case "cpy":
			regs[a[2]] = x
		case "inc":
			regs[a[1]]++
		case "dec":
			regs[a[1]]--
		case "jnz":
			if x != 0 {
				ptr += atoi(a[2])
				continue
			}
		}

		ptr++
	}
}

func getX(s string, regs map[string]int) int {
	if s[0] >= 'a' && s[0] <= 'z' {
		return regs[s]
	}
	return atoi(s)
}

func atoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}

func main() {
	perf.Bench(1, solution)
}
