package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	left, right := parseInput()

	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		l, r := left[i], right[i]
		sum += max(l, r) - min(l, r)
	}

	fmt.Println("total distance:", sum)
}

func part2() {
	left, right := parseInput()
	count := make(map[int]int)

	for i := 0; i < len(right); i++ {
		count[right[i]] += 1
	}

	score := 0

	for i := 0; i < len(left); i++ {
		l := left[i]
		score += l * count[l]
	}

	fmt.Println("similarity score:", score)
}

func parseInput() ([]int, []int) {
	file, _ := os.Open("2024-01.txt")

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
