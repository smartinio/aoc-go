package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1 := 0
	part2 := 0

	return part1, part2
}

func main() {
	part1, part2 := 0, 0
	sum := 0
	n := 1 // increase samples if benching perf

	for range n {
		start := time.Now()
		part1, part2 = solution()
		sum += int(time.Since(start).Milliseconds())
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
	fmt.Println("avg:", sum/n, "ms")
}
