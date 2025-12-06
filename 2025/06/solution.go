package main

import (
	_ "embed"
	"main/perf"
	"main/utils"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0
	parts := strings.Split(strings.TrimSpace(input), "\n")
	lines := parts[0 : len(parts)-1]
	regex := regexp.MustCompile(`([^\s]+)`)
	operators := utils.FindAllStringGroups(regex, parts[len(parts)-1])

	// part 1
	{
		problems := make([][]int, len(operators))

		for i := range problems {
			problems[i] = make([]int, len(parts)-1)
		}

		for i, line := range lines {
			numbers := utils.FindAllIntGroups(regex, line)
			for j, n := range numbers {
				problems[j][i] = n
			}
		}

		for i, problem := range problems {
			part1 += compute(problem, operators[i])
		}
	}

	// part 2
	{
		pIndex := 0
		height := len(lines)
		width := len(slices.MaxFunc(lines, compareLength))
		problems := make([][]int, len(operators))

		for i := range problems {
			problems[i] = []int{}
		}

		for x := range width {
			sb := strings.Builder{}

			for y := range height {
				c := charAt(lines[y], x)
				if c != ' ' {
					sb.WriteByte(c)
				}
			}

			if sb.Len() == 0 {
				pIndex++
			} else {
				n, _ := strconv.Atoi(sb.String())
				problems[pIndex] = append(problems[pIndex], n)
			}
		}

		for i, problem := range problems {
			part2 += compute(problem, operators[i])
		}
	}

	return part1, part2
}

func compute(nums []int, operator string) int {
	switch operator {
	case "+":
		return add(nums)
	case "*":
		return mul(nums)
	default:
		panic("unknown operator")
	}
}

func mul(nums []int) int {
	prod := 1
	for _, n := range nums {
		prod *= n
	}
	return prod
}

func add(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func charAt(line string, x int) byte {
	if x >= len(line) {
		return ' '
	}
	return line[x]
}

func compareLength(a, b string) int {
	return len(a) - len(b)
}

func main() {
	perf.Bench(1, solution)
}
