package main

import (
	_ "embed"
	"main/perf"
	"strings"
)

type Dir struct{ x, y int }
type Pos struct{ x, y int }

//go:embed input.txt
var input string

var dirs = map[rune]Dir{
	'U': {0, -1},
	'R': {1, 0},
	'D': {0, 1},
	'L': {-1, 0},
}

const NIL = -1

func solution() (part1 string, part2 string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// part 1
	{
		keypad := [][]rune{
			{'1', '2', '3'},
			{'4', '5', '6'},
			{'7', '8', '9'},
		}

		pos := Pos{1, 1} // key '5'

		digits := strings.Builder{}

		for _, line := range lines {
			for _, dir := range line {
				d := dirs[dir]
				pos = move(pos, d, keypad)
			}

			digit := keypad[pos.y][pos.x]
			digits.WriteRune(digit)
		}

		part1 = digits.String()
	}

	// part 2
	{
		keypad := [][]rune{
			{NIL, NIL, '1', NIL, NIL},
			{NIL, '2', '3', '4', NIL},
			{'5', '6', '7', '8', '9'},
			{NIL, 'A', 'B', 'C', NIL},
			{NIL, NIL, 'D', NIL, NIL},
		}

		pos := Pos{0, 2} // key '5'

		digits := strings.Builder{}

		for _, line := range lines {
			for _, dir := range line {
				pos = move(pos, dirs[dir], keypad)
			}

			digit := keypad[pos.y][pos.x]
			digits.WriteRune(digit)
		}

		part2 = digits.String()
	}

	return part1, part2
}

func move(pos Pos, d Dir, kp [][]rune) (n Pos) {
	n.x = pos.x + d.x
	n.y = pos.y + d.y

	if n.x < 0 || n.y < 0 || n.x >= len(kp) || n.y >= len(kp) || kp[n.y][n.x] == NIL {
		return pos
	}

	return n
}

func main() {
	perf.Bench(1, solution)
}
