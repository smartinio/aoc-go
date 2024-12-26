package main

import (
	_ "embed"
	"main/perf"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func solution() (int, string) {
	part1, part2 := 0, ""

	conns := map[string][]string{}
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		c := strings.Split(line, "-")
		conns[c[0]] = append(conns[c[0]], c[1])
		conns[c[1]] = append(conns[c[1]], c[0])
	}

	// part 1
	{
		type Key struct{ a, b, c string }
		set := map[Key]bool{}

		for k, v := range conns {
			for i := range v {
				for j := range v[1:] {
					a, b := v[i], v[j]
					startsWithT := k == "t" || a[0] == 't' || b[0] == 't'

					if startsWithT && slices.Contains(conns[a], b) {
						n := []string{k, a, b}
						slices.Sort(n)
						key := Key{n[0], n[1], n[2]}
						set[key] = true
					}
				}
			}
		}

		part1 = len(set)
	}

	// part 2
	{
		maxScore, code := 0, ""

		for k, friends := range conns {
			mutuals := map[string]int{}

			for _, friend := range friends {
				for _, other := range conns[friend] {
					if slices.Contains(friends, other) {
						mutuals[friend]++
					}
				}
			}

			score := 1
			for _, v := range mutuals {
				score *= v
			}

			if score > maxScore {
				maxScore = score
				parts := []string{k}
				for k := range mutuals {
					parts = append(parts, k)
				}
				slices.Sort(parts)
				code = strings.Join(parts, ",")
			}
		}

		part2 = code
	}

	return part1, part2
}

func main() {
	perf.Bench(20, solution)
}
