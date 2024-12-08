package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Context struct {
	target int
	parts  []int
	concat bool
}

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

		if combines(parts[0], 1, Context{testval, parts, false}) {
			part1 += testval
			part2 += testval
		} else if combines(parts[0], 1, Context{testval, parts, true}) {
			part2 += testval
		}
	}

	return part1, part2
}

func combines(a, i int, ctx Context) bool {
	if i == len(ctx.parts) {
		return a == ctx.target
	}

	if a > ctx.target {
		return false
	}

	b, j := ctx.parts[i], i+1

	return combines(a+b, j, ctx) || combines(a*b, j, ctx) || ctx.concat && combines(concat(a, b), j, ctx)
}

func concat(a int, b int) int {
	return a*int(math.Pow10(digits(b))) + b
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
	n := 10 // increase samples if benching perf

	for range n {
		start := time.Now()
		part1, part2 = solution()
		sum += int(time.Since(start).Milliseconds())
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
	fmt.Println("avg:", sum/n, "ms")
}
