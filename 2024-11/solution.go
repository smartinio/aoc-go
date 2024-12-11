package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Job struct {
	stone, runs int
}

func solution() (int, int) {
	part1, part2 := 0, 0
	nums := strings.Split(strings.TrimSpace(input), " ")
	cache := make(map[Job]int)

	for _, num := range nums {
		stone, _ := strconv.Atoi(num)
		part1 += 1 + spawns(Job{stone, 25 + 1}, cache)
		part2 += 1 + spawns(Job{stone, 75 + 1}, cache)
	}

	return part1, part2
}

func spawns(job Job, cache map[Job]int) int {
	if cj, ok := cache[job]; ok {
		return cj
	}

	stone, n := job.stone, job.runs
	total := 0

	for i := range n {
		if stone == 0 {
			stone = 1
		} else if digits(stone)%2 == 0 {
			a, b := split(stone)
			stone = a
			runs := n - i - 1
			if runs > 0 {
				next := Job{b, runs}
				total += 1 + spawns(next, cache)
			}
		} else {
			stone *= 2024
		}
	}

	cache[job] = total

	return total
}

func split(i int) (int, int) {
	half := digits(i) / 2
	b10 := int(math.Pow10(half))
	p1 := i / b10

	return p1, i - p1*b10
}

func digits(i int) int {
	if i == 0 {
		return 1
	}

	count := 0

	for i != 0 {
		i /= 10
		count++
	}

	return count
}

func main() {
	part1, part2 := 0, 0
	sum := 0
	n := 100 // increase samples if benching perf

	for range n {
		start := time.Now()
		part1, part2 = solution()
		sum += int(time.Since(start).Milliseconds())
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
	fmt.Println("avg:", sum/n, "ms")
}
