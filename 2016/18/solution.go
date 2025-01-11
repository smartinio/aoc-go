package main

import (
	_ "embed"
	"main/perf"
	"strings"
)

//go:embed input.txt
var input string

func solution() (part1 int, part2 int) {
	first := strings.TrimSpace(input)

	// part 1
	{
		prev := first
		for range 40 {
			part1 += strings.Count(prev, ".")
			prev = nextRow(prev)
		}
	}

	// part 2
	{
		prev := first
		for range 400000 {
			part2 += strings.Count(prev, ".")
			prev = nextRow(prev)
		}
	}

	return part1, part2
}

func nextRow(prev string) string {
	sb := strings.Builder{}
	pad := "." + prev + "."

	for i := 1; i < len(pad)-1; i++ {
		lcr := pad[i-1 : i+2]

		switch lcr {
		case "^^.", ".^^", "^..", "..^":
			sb.WriteRune('^')
		default:
			sb.WriteRune('.')
		}
	}

	return sb.String()
}

func main() {
	perf.Bench(1, solution)
}
