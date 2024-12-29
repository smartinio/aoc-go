package main

import (
	_ "embed"
	"main/perf"
	"strings"
)

//go:embed input.txt
var input string

func solution() (string, string) {
	part1 := increment(strings.TrimSpace(input))

	for !valid(part1) {
		part1 = increment(part1)
	}

	part2 := increment(part1)

	for !valid(part2) {
		part2 = increment(part2)
	}

	return part1, part2
}

func increment(pass string) string {
	last := len(pass) - 1
	next := inc(pass[last])

	if len(pass) == 1 {
		return next
	}

	remaining := pass[0:last]

	if next == "a" {
		return increment(remaining) + next
	}

	return remaining + next
}

func inc(char byte) string {
	if char == 'z' {
		return "a"
	}
	return string(char + 1)
}

func valid(pass string) bool {
	if strings.ContainsAny(pass, "iol") {
		return false
	}

	straight := false
	pairs := map[[2]byte]bool{}

	for i := range pass[:len(pass)-2] {
		if pass[i] == pass[i+1]-1 && pass[i+1] == pass[i+2]-1 {
			straight = true
		}
	}

	for i := range pass[:len(pass)-1] {
		if pass[i] == pass[i+1] {
			pairs[[2]byte{pass[i], pass[i+1]}] = true
		}
	}

	return straight && len(pairs) >= 2
}

func main() {
	perf.Bench(1, solution)
}
