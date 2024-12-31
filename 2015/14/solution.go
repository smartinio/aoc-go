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

const time = 2503

type Deer struct {
	speed, stamina, rest int
}

func solution() (int, int) {
	deers := parseInput()
	part1, part2 := 0, 0

	// part 1
	for _, deer := range deers {
		part1 = max(part1, distance(time, deer))
	}

	// part 2
	{
		points := map[string]int{}

		for t := range time {
			best, end := 0, t+1

			for _, deer := range deers {
				best = max(best, distance(end, deer))
			}

			for name, deer := range deers {
				if distance(end, deer) >= best {
					points[name]++
					part2 = max(part2, points[name])
				}
			}
		}
	}

	return part1, part2
}

func distance(time int, d Deer) int {
	period := d.stamina + d.rest
	cycles := time / period
	dx := d.speed * d.stamina

	return cycles*dx + d.speed*min(d.stamina, time%period)
}

func atoi(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func parseInput() map[string]Deer {
	re := regexp.MustCompile(`\d+`)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	deers := map[string]Deer{}

	for _, line := range lines {
		ns := re.FindAllString(line, -1)
		name := strings.Split(line, " ")[0]
		speed, stamina, rest := atoi(ns[0]), atoi(ns[1]), atoi(ns[2])
		deers[name] = Deer{speed, stamina, rest}
	}

	return deers
}

func main() {
	perf.Bench(1, solution)
}
