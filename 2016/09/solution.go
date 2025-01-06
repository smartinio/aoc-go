package main

import (
	_ "embed"
	"main/perf"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var markerRe = regexp.MustCompile(`\(\d+x\d+\)`)
var numberRe = regexp.MustCompile(`\d+`)

func solution() (part1 int, part2 int) {
	data := strings.TrimSpace(input)

	part1 = dfs1(data)
	part2 = dfs2(data)

	return part1, part2
}

func dfs1(data string) int {
	mi := markerRe.FindStringIndex(data)

	if len(mi) == 0 {
		return len(data)
	}

	marker := data[mi[0]:mi[1]]
	data = data[mi[1]:] // truncate marker
	length, repeat := parseMarker(marker)

	return mi[0] + repeat*length + dfs1(data[length:])
}

func dfs2(data string) int {
	mi := markerRe.FindStringIndex(data)

	if len(mi) == 0 {
		return len(data)
	}

	marker := data[mi[0]:mi[1]]
	data = data[mi[1]:] // truncate marker
	length, repeat := parseMarker(marker)

	return mi[0] + repeat*dfs2(data[:length]) + dfs2(data[length:])
}

func parseMarker(marker string) (length, repeat int) {
	m := numberRe.FindAllString(marker, 2)
	length, _ = strconv.Atoi(m[0])
	repeat, _ = strconv.Atoi(m[1])
	return
}

func main() {
	perf.Bench(1, solution)
}
