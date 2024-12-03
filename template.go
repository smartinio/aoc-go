package main

import (
	"bufio"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {

}

func part2() {

}

func parseInput() {
	file, _ := os.Open("$DAY/example.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
	}
}
