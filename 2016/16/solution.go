package main

import (
	_ "embed"
	"main/perf"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func solution() (part1 string, part2 string) {
	data := strings.TrimSpace(input)

	// part 1
	{
		disk := dragon(data, 272)
		part1 = checksum(disk)
	}

	// part 2
	{
		disk := dragon(data, 35651584)
		part2 = checksum(disk)
	}

	return part1, part2
}

func dragon(a string, size int) string {
	if len(a) >= size {
		return a[:size]
	}

	sb := strings.Builder{}

	b := []byte(a)
	slices.Reverse(b)

	for i := range len(b) {
		if b[i] == '0' {
			b[i] = '1'
		} else {
			b[i] = '0'
		}
	}

	sb.WriteString(a)
	sb.WriteRune('0')
	sb.WriteString(string(b))

	return dragon(sb.String(), size)
}

func checksum(str string) string {
	if !even(len(str)) {
		return str
	}

	sb := strings.Builder{}

	for i := 0; i < len(str); i += 2 {
		if str[i] == str[i+1] {
			sb.WriteRune('1')
		} else {
			sb.WriteRune('0')
		}
	}

	return checksum(sb.String())
}

func even(n int) bool {
	return (n | 1) > n
}

func main() {
	perf.Bench(1, solution)
}
