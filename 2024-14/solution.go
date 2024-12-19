package main

import (
	"bufio"
	_ "embed"
	"main/perf"
	"regexp"
	"strconv"
	"strings"
)

type Robot struct {
	px, py, vx, vy int
}

const W, H = 101, 103

//go:embed input.txt
var input string

func solution() (int, int) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	re := regexp.MustCompile(`-?\d+`)
	robots := []Robot{}

	for scanner.Scan() {
		line := scanner.Text()
		m := re.FindAllString(line, -1)

		px, _ := strconv.Atoi(m[0])
		py, _ := strconv.Atoi(m[1])
		vx, _ := strconv.Atoi(m[2])
		vy, _ := strconv.Atoi(m[3])

		robots = append(robots, Robot{px, py, vx, vy})
	}

	startSafety := safetyFactor(1, robots)
	minSafety, minSeconds := startSafety, 0

	for seconds := 2; true; seconds++ {
		safety := safetyFactor(seconds, robots)

		if safety < minSafety {
			minSafety, minSeconds = safety, seconds
		} else if safety == startSafety {
			break
		}
	}

	part1 := safetyFactor(100, robots)
	part2 := minSeconds

	return part1, part2
}

func safetyFactor(seconds int, robots []Robot) int {
	q1, q2, q3, q4 := 0, 0, 0, 0

	for _, r := range robots {
		x := wrap(r.px+seconds*r.vx, W)
		y := wrap(r.py+seconds*r.vy, H)

		if x < W/2 && y < H/2 {
			q1++
		} else if x > W/2 && y < H/2 {
			q2++
		} else if x < W/2 && y > H/2 {
			q3++
		} else if x > W/2 && y > H/2 {
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func wrap(i, max int) int {
	return ((i % max) + max) % max
}

func main() {
	perf.Bench(10, solution)
}
