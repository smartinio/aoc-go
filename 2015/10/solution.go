package main

import (
	_ "embed"
	"fmt"
	"main/perf"
	"strings"
)

//go:embed input.txt
var input string

func solution() (part1 int, part2 int) {
	sequence := strings.TrimSpace(input)

	for range 40 {
		sequence = transform(sequence)
	}

	part1 = len(sequence)

	for range 10 {
		sequence = transform(sequence)
	}

	part2 = len(sequence)

	return part1, part2
}

func transform(str string) string {
	result := strings.Builder{}
	count, prev := 0, string(str[0])

	for _, char := range str + " " {
		if prev != string(char) {
			result.WriteString(fmt.Sprintf("%d", count))
			result.WriteString(prev)
			prev = string(char)
			count = 0
		}
		count++
	}

	return result.String()
}

func main() {
	perf.Bench(1, solution)
}
