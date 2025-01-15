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

	// part 1
	{
		lines := strings.Split(strings.TrimSpace(input), "\n")
		regs := map[string]int{"a": 7}
		compute(lines, regs)
		part1 = regs["a"]
	}

	// part 2
	{
		input = strings.Replace(input, "inc a", "mul c d a", 1)
		input = strings.Replace(input, "dec c", "cpy 0 c", 1)
		input = strings.Replace(input, "dec d", "cpy 0 d", 1)

		lines := strings.Split(strings.TrimSpace(input), "\n")
		regs := map[string]int{"a": 12}
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
		case "tgl":
			p := ptr + x

			if p < 0 || p >= len(lines) {
				break
			}

			v := lines[p]
			inst := ""

			switch v[0:3] {
			case "inc":
				inst = "dec" + v[3:]
			case "tgl", "dec":
				inst = "inc" + v[3:]
			case "jnz":
				inst = "cpy" + v[3:]
			case "cpy":
				inst = "jnz" + v[3:]
			}

			lines[p] = inst
		case "mul":
			regs[a[3]] += x * getX(a[2], regs)
		case "cpy":
			regs[a[2]] = x
		case "inc":
			regs[a[1]]++
		case "dec":
			regs[a[1]]--
		case "jnz":
			if x != 0 {
				ptr += getX(a[2], regs)
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
