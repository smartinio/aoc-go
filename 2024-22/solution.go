package main

import (
	_ "embed"
	"main/perf"
	"strconv"
	"strings"
)

type Seq struct{ s0, s1, s2, s3 int }
type Sale struct{ buyer, s0, s1, s2, s3 int }

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0

	lines := strings.Split(strings.TrimSpace(input), "\n")

	// part 1
	for _, line := range lines {
		secret, _ := strconv.Atoi(line)
		for range 2000 {
			secret = evolve(secret)
		}
		part1 += secret
	}

	// part 2
	deltas := map[int][]int{}
	prices := map[int][]int{}

	for _, line := range lines {
		secret, _ := strconv.Atoi(line)
		deltas[secret], prices[secret] = process(secret)
	}

	seqPrices := map[Seq]int{}
	sold := map[Sale]bool{}

	for buyer, d := range deltas {
		for j := range d[:len(d)-3] {
			s := Seq{d[j], d[j+1], d[j+2], d[j+3]}
			sale := Sale{buyer, s.s0, s.s1, s.s2, s.s3}

			if !sold[sale] {
				seqPrices[s] += prices[buyer][j+3]
				sold[sale] = true
			}
		}
	}

	for _, price := range seqPrices {
		part2 = max(part2, price)
	}

	return part1, part2
}

func process(secret int) ([]int, []int) {
	deltas, prices := []int{}, []int{}
	prev := lastDigit(secret)

	for range 2000 {
		secret = evolve(secret)
		price := lastDigit(secret)
		prices = append(prices, price)
		deltas = append(deltas, price-prev)
		prev = price
	}

	return deltas, prices
}

func lastDigit(secret int) int {
	return secret % 10
}

func evolve(secret int) int {
	secret = mix(secret, secret*64)
	secret = prune(secret)
	secret = mix(secret, secret/32)
	secret = prune(secret)
	secret = mix(secret, secret*2048)
	return prune(secret)
}

func mix(secret int, value int) int {
	return secret ^ value
}

func prune(secret int) int {
	return secret % 16777216
}

func main() {
	perf.Bench(1, solution)
}
