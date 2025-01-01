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

type Sue = map[string]int

func solution() (int, int) {
	part1, part2 := 0, 0

	lines := strings.Split(strings.TrimSpace(input), "\n")

	re := regexp.MustCompile(`([a-z]+): (\d+)`)

	target := Sue{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	sues := []Sue{}

	for _, line := range lines {
		sue := Sue{}

		for _, k := range re.FindAllStringSubmatch(line, -1) {
			sue[k[1]], _ = strconv.Atoi(k[2])
		}

		sues = append(sues, sue)
	}

	// part 1
	{
	outer1:
		for i, sue := range sues {
			for k, v := range target {
				if val, ok := sue[k]; ok && val != v {
					continue outer1
				}
			}

			part1 = i + 1
			break
		}
	}

	// part 2
	{
	outer2:
		for i, sue := range sues {
			for k, v := range target {
				val, remembers := sue[k]

				if !remembers {
					continue
				}

				switch {
				case k == "cats" || k == "trees":
					if val <= v {
						continue outer2
					}
				case k == "pomeranians" || k == "goldfish":
					if val >= v {
						continue outer2
					}
				default:
					if val != v {
						continue outer2
					}
				}
			}

			part2 = i + 1
			break
		}
	}

	return part1, part2
}

func main() {
	perf.Bench(1, solution)
}
