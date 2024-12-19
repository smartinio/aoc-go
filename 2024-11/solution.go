package main

import (
	_ "embed"
	"main/perf"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Stone struct {
	number, blinks int
}

func solution() (int, int) {
	part1, part2 := 0, 0
	chars := strings.Split(strings.TrimSpace(input), " ")
	cache := make(map[Stone]int)

	for _, str := range chars {
		number, _ := strconv.Atoi(str)
		part1 += 1 + splits(Stone{number, 25}, cache)
		part2 += 1 + splits(Stone{number, 75}, cache)
	}

	return part1, part2
}

func splits(stone Stone, cache map[Stone]int) int {
	if cached, ok := cache[stone]; ok {
		return cached
	}

	num, blinks := stone.number, stone.blinks
	total := 0

	for i := range blinks {
		if num == 0 {
			num = 1
		} else if digits(num)%2 == 0 {
			left, right := split(num)
			num = left
			total += 1 + splits(Stone{right, blinks - i - 1}, cache)
		} else {
			num *= 2024
		}
	}

	cache[stone] = total

	return total
}

func split(i int) (int, int) {
	half := digits(i) / 2
	tens := int(math.Pow10(half))
	left := i / tens
	right := i - left*tens

	return left, right
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
	perf.Bench(20, solution)
}
