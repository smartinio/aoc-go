package main

import (
	"fmt"
	"testing"
)

func BenchmarkSolution(b *testing.B) {
	part1, part2 := solution()
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}
