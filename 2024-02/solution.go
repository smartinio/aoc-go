package main

import (
	"bufio"
	"main/perf"
	"os"
	"strconv"
	"strings"
)

func solution() (int, int) {
	part1, part2 := 0, 0
	reports := parseInput()

	for _, report := range reports {
		if isSafe(report) {
			part1 += 1
		}
	}

	brute := func(report []int) {
		if isSafe(report) {
			part2 += 1
			return
		}

		for i := range report {
			dampenedReport := dampen(report, i)

			if isSafe(dampenedReport) {
				part2 += 1
				return
			}
		}
	}

	for _, report := range reports {
		brute(report)
	}

	return part1, part2
}

func isSafe(report []int) bool {
	increasing, decreasing := false, false
	prev := -1

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

func main() {
	perf.Bench(100, solution)
}
