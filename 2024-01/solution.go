package main

import (
	"bufio"
	"main/perf"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution() (int, int) {
	part1, part2 := 0, 0
	left, right := parseInput()
	count := map[int]int{}

	sort.Ints(left)
	sort.Ints(right)

	for i := range left {
		l, r := left[i], right[i]
		part1 += max(l, r) - min(l, r)
		count[right[i]] += 1
	}

	for _, l := range left {
		part2 += l * count[l]
	}

	return part1, part2
}

func parseInput() ([]int, []int) {
	file, _ := os.Open("2024-01/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	left, right := []int{}, []int{}

	for scanner.Scan() {
		line := scanner.Text()
		ids := strings.Split(line, "   ")
		l, _ := strconv.Atoi(ids[0])
		r, _ := strconv.Atoi(ids[1])
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

func main() {
	perf.Bench(1, solution)
}
