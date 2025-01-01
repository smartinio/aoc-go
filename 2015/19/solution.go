package main

import (
	_ "embed"
	"main/perf"
	"regexp"
	"strings"
)

type Reps map[string][]string

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0

	replacements, molecule := parseInput()

	// part 1
	{
		molecules := map[string]bool{}

		for old, news := range replacements {
			re := regexp.MustCompile(old)
			for _, indexRange := range re.FindAllStringIndex(molecule, -1) {
				pos := indexRange[0]
				for _, new := range news {
					replaced := molecule[0:pos] + strings.Replace(molecule[pos:], old, new, 1)
					molecules[replaced] = true
				}
			}
		}

		part1 = len(molecules)
	}

	// part 2
	part2 = dfs(molecule, "e", 0, replacements, map[string]bool{})

	return part1, part2
}

func dfs(mol, target string, steps int, replacements Reps, visited map[string]bool) int {
	if mol == target {
		return steps
	}

	if visited[mol] {
		return -1
	}

	visited[mol] = true

	for old, news := range replacements {
		for _, new := range news {
			replaced := strings.Replace(mol, new, old, 1)
			totalSteps := dfs(replaced, target, steps+1, replacements, visited)

			if totalSteps != -1 {
				return totalSteps
			}
		}
	}

	return -1
}

func parseInput() (Reps, string) {
	blocks := strings.Split(strings.TrimSpace(input), "\n\n")

	replacements := Reps{}
	start := blocks[1]

	for _, rep := range strings.Split(blocks[0], "\n") {
		r := strings.Split(rep, " => ")
		replacements[r[0]] = append(replacements[r[0]], r[1])
	}

	return replacements, start
}

func main() {
	perf.Bench(1, solution)
}
