package main

import (
	"context"
	_ "embed"
	"main/perf"
	"strconv"
	"strings"
	"sync"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0
	num, _ := strconv.Atoi(strings.TrimSpace(input))

	part1 = solve(num, packageCount)
	part2 = solve(num, packageCount2)

	return part1, part2
}

func solve(num int, counter func(int) int) int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	houses := make(chan int)
	result := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(houses)

		i := 1
		for {
			select {
			case <-ctx.Done():
				return
			case houses <- 12 * i:
				i++
			}
		}
	}()

	for w := range 8 {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for houseNo := range houses {
				if counter(houseNo) >= num {
					result <- houseNo
					cancel()
					return
				}
			}
		}(w)
	}

	answer := <-result
	wg.Wait()

	return answer
}

func packageCount(houseNo int) (total int) {
	for i := range houseNo {
		elf := i + 1

		if houseNo%elf == 0 {
			total += elf
		}
	}

	return 10 * total
}

func packageCount2(houseNo int) (total int) {
	for i := houseNo/50 - 1; i < houseNo; i++ {
		elf := i + 1

		if elf*50 >= houseNo && houseNo%elf == 0 {
			total += elf
		}
	}

	return 11 * total
}

func main() {
	perf.Bench(1, solution)
}
