package main

import (
	_ "embed"
	"main/perf"
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Room struct {
	name string
	id   int
}

func solution() (int, int) {
	part1, part2 := 0, 0

	lines := strings.Split(strings.TrimSpace(input), "\n")
	re := regexp.MustCompile(`((?:[a-z]+-?)+)-(\d+)\[([a-z]+)\]`)
	rooms := []Room{}

	// part 1
	for _, line := range lines {
		count := map[rune]int{}
		matches := re.FindAllStringSubmatch(line, -1)
		m := matches[0][1:]

		for _, char := range m[0] {
			if char != '-' {
				count[char]++
			}
		}

		keys := slices.Collect(maps.Keys(count))

		slices.SortFunc(keys, func(a, b rune) int {
			cmp := count[b] - count[a]
			if cmp == 0 {
				return int(a) - int(b)
			}
			return cmp
		})

		if string(keys[:5]) == m[2] {
			id := atoi(m[1])
			part1 += id
			rooms = append(rooms, Room{m[0], id})
		}
	}

	// part 2
	for _, room := range rooms {
		if strings.Contains(rotate(room), "north") {
			part2 = room.id
			break
		}
	}

	return part1, part2
}

func rotate(room Room) string {
	sb := strings.Builder{}
	offset := int('a')

	for _, char := range room.name {
		if char == '-' {
			sb.WriteRune(' ')
		} else {
			rotated := offset + (int(char)-offset+room.id)%26
			sb.WriteRune(rune(rotated))
		}
	}

	return sb.String()
}

func atoi(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func main() {
	perf.Bench(1, solution)
}
