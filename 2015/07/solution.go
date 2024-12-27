package main

import (
	_ "embed"
	"main/perf"
	"strconv"
	"strings"
)

type Source struct {
	ins []string
	op  string
}

//go:embed input.txt
var input string

func solution() (uint16, uint16) {
	sources := parseInput()

	part1 := dfs("a", sources, map[string]uint16{})

	part2 := dfs("a", sources, map[string]uint16{"b": part1})

	return part1, part2
}

func dfs(wire string, sources map[string]Source, cache map[string]uint16) uint16 {
	if cached, ok := cache[wire]; ok {
		return cached
	}

	source, isSource := sources[wire]

	if !isSource {
		val, _ := strconv.Atoi(wire)
		return uint16(val)
	}

	var result uint16

	if source.op == "PASSTHROUGH" {
		result = dfs(source.ins[0], sources, cache)
	} else if source.op == "NOT" {
		result = ^dfs(source.ins[0], sources, cache)
	} else {
		a := dfs(source.ins[0], sources, cache)
		b := dfs(source.ins[1], sources, cache)
		result = transform(source, a, b)
	}

	cache[wire] = result

	return result
}

func transform(source Source, a, b uint16) uint16 {
	switch source.op {
	case "AND":
		return a & b
	case "LSHIFT":
		return a << b
	case "OR":
		return a | b
	case "RSHIFT":
		return a >> b
	default:
		return uint16(0)
	}
}

func parseInput() map[string]Source {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sources := map[string]Source{}

	for _, line := range lines {
		p := strings.Split(line, " ")
		if len(p) == 3 {
			sources[p[2]] = Source{[]string{p[0]}, "PASSTHROUGH"}
		} else if p[0] == "NOT" {
			sources[p[3]] = Source{[]string{p[1]}, "NOT"}
		} else {
			sources[p[4]] = Source{[]string{p[0], p[2]}, p[1]}
		}
	}

	return sources
}

func main() {
	perf.Bench(20, solution)
}
