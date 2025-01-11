package main

import (
	_ "embed"
	"main/perf"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Node struct {
	pos  int
	next *Node
	prev *Node
}

func solution() (part1 int, part2 int) {
	elfCount, _ := strconv.Atoi(strings.TrimSpace(input))

	// part 1
	{
		start := dll(elfCount)
		curr := start

		for curr.next != curr {
			next := curr.next
			curr.next = next.next
			curr = curr.next
		}

		part1 = curr.pos
	}

	// part 2
	{
		start := dll(elfCount)
		curr := start
		across := start

		for range elfCount / 2 {
			across = across.next
		}

		for size := elfCount; size > 1; size-- {
			across.next.prev = across.prev
			across.prev.next = across.next
			curr = curr.next

			if size%2 == 1 {
				across = across.next.next
			} else {
				across = across.next
			}
		}

		part2 = curr.pos
	}

	return part1, part2
}

func dll(n int) *Node {
	start := &Node{pos: 1}
	prev := start

	for range n - 1 {
		prev.next = &Node{pos: prev.pos + 1, prev: prev}
		prev = prev.next
	}

	prev.next = start
	start.prev = prev

	return start
}

func main() {
	perf.Bench(1, solution)
}
