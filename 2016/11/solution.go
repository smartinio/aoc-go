package main

import (
	_ "embed"
	"main/perf"
	"main/utils"
	"regexp"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type State [15]int

type Item struct {
	kind, element string
	floor         int
}

func solution() (part1 float64, part2 float64) {
	part1 = bfs(parseInput(1))
	input2 := parseInput(2)
	part2 = bfs(input2)

	return part1, part2
}

type Vertex struct {
	elevator int
	steps    float64
	items    []Item
}

func bfs(initial []Item) float64 {
	q := []Vertex{{1, 0, initial}}
	visited := map[State]bool{}

	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		if topFloor(v.items) {
			return v.steps
		}

		floorItems := getFloor(v.items, v.elevator)

		// choose 1-2 items
		for i := range 2 {
			for chosen := range utils.Choose(floorItems, i+1) {
				for _, dy := range []int{1, -1} {
					next := make([]Item, len(v.items))
					copy(next, v.items)
					nextFloor := max(1, min(4, v.elevator+dy))

					for _, c := range chosen {
						idx := slices.Index(next, c)
						next[idx].floor = nextFloor
					}

					nextState := makeState(nextFloor, next)

					if valid(next) && !visited[nextState] {
						q = append(q, Vertex{nextFloor, v.steps + 1, next})
						visited[nextState] = true
					}
				}
			}
		}
	}

	return -1
}

func heuristic(slice []Item) (h int) {
	for _, item := range slice {
		h += 4 - item.floor
	}
	return h
}

func valid(slice []Item) bool {
	for i := range 4 {
		floor := getFloor(slice, i+1)
		if !stable(floor) {
			return false
		}
	}
	return true
}

func stable(floor []Item) bool {
	chips := map[string]bool{}
	rtgs := map[string]bool{}

	for _, item := range floor {
		if item.kind == "chip" {
			chips[item.element] = true
		} else {
			rtgs[item.element] = true
		}
	}

	if len(chips) == 0 || len(rtgs) == 0 {
		return true
	}

	for element := range chips {
		if !rtgs[element] {
			return false
		}
	}

	return true
}

func getFloor(slice []Item, f int) (floor []Item) {
	for _, item := range slice {
		if item.floor == f {
			floor = append(floor, item)
		}
	}
	return
}

func topFloor(slice []Item) bool {
	for _, item := range slice {
		if item.floor != 4 {
			return false
		}
	}
	return true
}

func makeState(elevator int, items []Item) (state State) {
	state[0] = elevator
	for i, item := range items {
		state[i+1] = item.floor
	}
	return
}

func parseInput(part int) []Item {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	chipRe := regexp.MustCompile(`[a-z]+-compatible microchip`)
	genRe := regexp.MustCompile(`[a-z]+ generator`)
	items := []Item{}

	for i, line := range lines {
		if strings.Contains(line, "nothing") {
			continue
		}

		floor := i + 1

		if part == 2 && floor == 1 {
			line += "An elerium generator. An elerium-compatible microchip. A dilithium generator. A dilithium-compatible microchip."
		}

		for _, chip := range chipRe.FindAllString(line, -1) {
			element := chip[:strings.Index(chip, "-")]
			items = append(items, Item{"chip", element, floor})
		}

		for _, gen := range genRe.FindAllString(line, -1) {
			element := gen[:strings.Index(gen, " ")]
			items = append(items, Item{"rtg", element, floor})
		}
	}

	return items
}

func main() {
	perf.Bench(1, solution)
}
