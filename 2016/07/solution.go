package main

import (
	_ "embed"
	"main/perf"
	"regexp"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func solution() (part1 int, part2 int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	hypernets := regexp.MustCompile(`\[[a-z]+\]`)
	supernets := regexp.MustCompile(`\][a-z]+|[a-z]+\[`)

	// part 1
	for _, line := range lines {
		hns := hypernets.FindAllString(line, -1)
		sns := supernets.FindAllString(line, -1)

		if slices.ContainsFunc(hns, containsAbba) {
			continue
		}

		if !slices.ContainsFunc(sns, containsAbba) {
			continue
		}

		part1++
	}

	// part 2
	for _, line := range lines {
		hns := hypernets.FindAllString(line, -1)
		sns := supernets.FindAllString(line, -1)
		abas := getAbas(sns)
		babs := getAbas(hns)

		if slices.ContainsFunc(abas, func(aba string) bool {
			bab := string([]byte{aba[1], aba[0], aba[1]})
			return slices.Contains(babs, bab)
		}) {
			part2++
			continue
		}
	}

	return part1, part2
}

func getAbas(hs []string) (abas []string) {
	for _, h := range hs {
		for i := 0; i < len(h)-2; i++ {
			if h[i] != h[i+1] && h[i] == h[i+2] {
				aba := h[i : i+3]
				if !slices.Contains(abas, aba) {
					abas = append(abas, aba)
				}
			}
		}
	}
	return abas
}

func containsAbba(h string) bool {
	for i := 0; i < len(h)-3; i++ {
		if h[i] != h[i+1] && h[i] == h[i+3] && h[i+1] == h[i+2] {
			return true
		}
	}
	return false
}

func main() {
	perf.Bench(1, solution)
}
