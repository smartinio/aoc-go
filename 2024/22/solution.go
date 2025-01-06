package main

import (
	_ "embed"
	"main/perf"
	"main/utils"
	"strconv"
	"strings"
)

type Seq [4]int

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0

	lines := strings.Split(strings.TrimSpace(input), "\n")

	// part 1
	{
		secrets := utils.ConcurrentFunc(lines, func(line string) int {
			secret, _ := strconv.Atoi(line)
			for range 2000 {
				secret = evolve(secret)
			}
			return secret
		})

		for _, secret := range secrets {
			part1 += secret
		}
	}

	// part 2
	{
		results := utils.ConcurrentFunc(lines, func(line string) map[Seq]int {
			secret, _ := strconv.Atoi(line)
			deltas, price := process(secret)
			prices := make(map[Seq]int, len(deltas)-3)

			for j := range len(deltas) - 3 {
				seq := Seq(deltas[j : j+4])
				if _, ok := prices[seq]; !ok {
					prices[seq] = price[j+3]
				}
			}

			return prices
		})

		total := make(map[Seq]int, 400000)

		for _, prices := range results {
			for seq, price := range prices {
				total[seq] += price
			}
		}

		for _, price := range total {
			part2 = max(part2, price)
		}
	}

	return part1, part2
}

func process(secret int) ([]int, []int) {
	deltas, prices := make([]int, 2000), make([]int, 2000)
	prev := secret % 10

	for i := range 2000 {
		secret = evolve(secret)
		price := secret % 10
		prices[i] = price
		deltas[i] = price - prev
		prev = price
	}

	return deltas, prices
}

func evolve(secret int) int {
	secret ^= secret * 64
	secret %= 16777216
	secret ^= secret / 32
	secret %= 16777216
	secret ^= secret * 2048
	return secret % 16777216
}

func main() {
	perf.Bench(1, solution)
}
