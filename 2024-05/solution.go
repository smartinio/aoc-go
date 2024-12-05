package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rules, updates := parseInput()
	part1, part2 := 0, 0

	compareRules := func(a string, b string) int {
		rule := fmt.Sprintf("%s|%s", b, a)

		if rules[rule] {
			return 1
		}

		return -1
	}

	for _, update := range updates {
		mid := len(update) / 2

		if slices.IsSortedFunc(update, compareRules) {
			val, _ := strconv.Atoi(update[mid])
			part1 += val
			continue
		}

		slices.SortFunc(update, compareRules)
		val, _ := strconv.Atoi(update[mid])
		part2 += val
	}

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func parseInput() (map[string]bool, [][]string) {
	file, _ := os.Open("2024-05/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	rules := make(map[string]bool)
	updates := [][]string{}

	for scanner.Scan() {
		word := scanner.Text()

		if word[2] == '|' {
			rules[word] = true
		} else {
			update := strings.Split(word, ",")
			updates = append(updates, update)
		}
	}

	return rules, updates
}
