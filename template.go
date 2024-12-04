package main

import (
	"bufio"
	"os"
)

func main() {
	input := parseInput()
}

func parseInput() {
	file, _ := os.Open("$DAY/example.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
	}
}
