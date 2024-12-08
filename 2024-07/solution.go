package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func solution() (int, int) {
	part1 := 0
	part2 := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	re := regexp.MustCompile(`[^\d]+`)

	for scanner.Scan() {
		line := scanner.Text()
		stringNums := re.Split(line, -1)
		nums := []int{}

		for _, str := range stringNums {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}

		testval := nums[0]
		parts := nums[1:]

		if combines(parts, testval, false) {
			part1 += testval
			part2 += testval
		} else if combines(parts, testval, true) {
			part2 += testval
		}
	}

	return part1, part2
}

func combines(ns []int, target int, concatEnabled bool) bool {
	if len(ns) == 0 || ns[0] > target {
		return false
	}

	if len(ns) == 1 {
		return ns[0] == target
	}

	add := func() bool {
		next := append([]int{ns[0] + ns[1]}, ns[2:]...)
		return combines(next, target, concatEnabled)
	}

	multiply := func() bool {
		next := append([]int{ns[0] * ns[1]}, ns[2:]...)
		return combines(next, target, concatEnabled)
	}

	concat := func() bool {
		c, _ := strconv.Atoi(fmt.Sprintf("%d%d", ns[0], ns[1]))
		next := append([]int{c}, ns[2:]...)
		return combines(next, target, concatEnabled)
	}

	return add() || multiply() || concatEnabled && concat()
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
