package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	reports := parseInput()
	safeCount := 0

	for _, report := range reports {
		if isSafe(report) {
			safeCount += 1
		}
	}

	fmt.Println("part 1:", safeCount)
}

func part2() {
	reports := parseInput()
	safeCount := 0

outer:
	for _, report := range reports {
		if isSafe(report) {
			safeCount += 1
			continue
		}

		for i := range report {
			dampenedReport := dampen(report, i)

			if isSafe(dampenedReport) {
				safeCount += 1
				continue outer
			}
		}
	}

	fmt.Println("part 2:", safeCount)
}

func isSafe(report []int) bool {
	prev := -1
	increasing := false
	decreasing := false

	for _, level := range report {
		if prev != -1 {
			dx := absDelta(prev, level)
			increasing = increasing || level > prev
			decreasing = decreasing || level < prev

			if dx < 1 || dx > 3 || (increasing && decreasing) {
				return false
			}
		}

		prev = level
	}

	return true
}

func dampen(report []int, indexToRemove int) []int {
	dampenedReport := []int{}

	for i, level := range report {
		if i != indexToRemove {
			dampenedReport = append(dampenedReport, level)
		}
	}

	return dampenedReport
}

func absDelta(a int, b int) int {
	return max(a, b) - min(a, b)
}

func parseInput() [][]int {
	file, _ := os.Open("2024-02/input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	reports := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		ids := strings.Split(line, " ")

		report := []int{}

		for _, id := range ids {
			level, _ := strconv.Atoi(id)
			report = append(report, level)
		}

		reports = append(reports, report)
	}

	return reports
}
