package main

import (
	_ "embed"
	"main/perf"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0

	lines := strings.Split(strings.TrimSpace(input), ",")

	for _, line := range lines {
		nums := strings.Split(line, "-")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])

	idScan:
		for i := n1; i <= n2; i++ {
			num := strconv.Itoa(i)
			length := len(num)

			if length%2 == 0 {
				halfA := num[0 : length/2]
				halfB := num[length/2:]

				if halfA == halfB {
					part1 += i
				}
			}

		segmentScan:
			for k := 1; k <= length/2; k++ {
				slice := num[0:k]

				for m := k; m <= length-k; m += k {
					other := num[m : m+k]

					if slice != other {
						continue segmentScan
					}
				}

				if length%k == 0 {
					part2 += i
					continue idScan
				}
			}
		}
	}

	return part1, part2
}

func main() {
	perf.Bench(1, solution)
}
