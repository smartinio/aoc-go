package perf

import (
	"fmt"
	"slices"
	"time"
)

type Result[P1 any, P2 any] struct {
	Part1 P1
	Part2 P2
	P99   time.Duration
	P50   time.Duration
	Avg   time.Duration
}

func Bench[P1 any, P2 any](n int, callback func() (P1, P2)) Result[P1, P2] {
	times := make([]time.Duration, n)

	var p1 P1
	var p2 P2

	for i := range n {
		start := time.Now()
		p1, p2 = callback()
		times[i] = time.Since(start)
	}

	slices.Sort(times)

	sum := 0
	for _, t := range times {
		sum += int(t)
	}

	avg := time.Duration(sum / n)
	p50 := times[int(0.50*float64(n))]
	p99 := times[int(0.99*float64(n))]

	result := Result[P1, P2]{p1, p2, p99, p50, avg}

	fmt.Println("part1:", result.Part1)
	fmt.Println("part2:", result.Part2)
	fmt.Println("---")
	fmt.Println("p50:", result.P50)
	fmt.Println("avg:", result.Avg)
	fmt.Println("p99:", result.P99)
	fmt.Println("---")

	return result
}
