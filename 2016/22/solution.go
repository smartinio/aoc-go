package main

import (
	_ "embed"
	"main/perf"
	"main/utils"
	"math"
	"regexp"
	"strings"
)

type Pos struct{ x, y int }

type Dir struct{ x, y int }

type Node struct{ x, y, size, used, avail int }

type Nodes map[Pos]*Node

type Item struct {
	node *Node
	prev *Item
}

//go:embed input.txt
var input string

func solution() (part1 int, part2 int) {
	nodes := parseInput()

	// part 1
	for _, from := range nodes {
		for _, to := range nodes {
			if viable(from, to) {
				part1++
			}
		}
	}

	// part 2
	{
		maxX := int(math.Sqrt(float64(len(nodes))) - 1)
		start, end := nodes[Pos{maxX, 0}], nodes[Pos{0, 0}]
		var empty *Node

		for _, v := range nodes {
			if v.used == 0 {
				empty = v
				break
			}
		}

		curr := start

		for curr != end {
			left := nodes[Pos{curr.x - 1, curr.y}]
			chain := bfs(empty, left, curr, nodes)

			for i := range len(chain) - 1 {
				to, from := chain[i], chain[i+1]
				move(from, to)
				part2++
			}

			move(curr, left)
			part2++

			empty = curr
			curr = left
		}
	}

	return part1, part2
}

func bfs(from, to, data *Node, nodes Nodes) []*Node {
	q := []*Item{{from, nil}}
	visited := map[*Node]bool{from: true}
	dirs := []Dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		curr := item.node

		if curr == to {
			return getPath(item, []*Node{})
		}

		for _, d := range dirs {
			next := nodes[Pos{curr.x + d.x, curr.y + d.y}]

			switch {
			case
				next == nil,
				next == data,
				!viable(next, &Node{avail: from.size}),
				visited[next]:
				continue
			default:
				q = append(q, &Item{next, item})
				visited[next] = true
			}
		}
	}

	return []*Node{}
}

func viable(from, to *Node) bool {
	return from != to && from.used > 0 && from.used <= to.avail
}

func move(from, to *Node) {
	to.used += from.used
	to.avail -= from.used
	from.avail = from.size
	from.used = 0
}

func getPath(item *Item, path []*Node) []*Node {
	if item.prev != nil {
		path = getPath(item.prev, path)
	}

	return append(path, item.node)
}

func parseInput() Nodes {
	re := regexp.MustCompile(`(\d+)`)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	nodes := make(Nodes, len(lines[2:]))

	for _, line := range lines[2:] {
		n := utils.FindAllIntGroups(re, line)
		pos := Pos{n[0], n[1]}
		nodes[pos] = &Node{n[0], n[1], n[2], n[3], n[4]}
	}

	return nodes
}

func main() {
	perf.Bench(1, solution)
}
