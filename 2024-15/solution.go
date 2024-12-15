package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

type Dir struct{ x, y int }
type Pos struct{ x, y int }

//go:embed input.txt
var input string
var UP, RIGHT, DOWN, LEFT = Dir{0, -1}, Dir{1, 0}, Dir{0, 1}, Dir{-1, 0}

func solution() (int, int) {
	part1, part2 := 0, 0

	// part 1
	{
		replacements := []string{}

		grid, moves, robot := parseInput(replacements)

		moveRobot(grid, moves, robot)

		part1 = sumCoords("O", grid)
	}

	// part 2
	{
		replacements := []string{"O", "[]", "#", "##", ".", "..", "@", "@."}

		grid, moves, robot := parseInput(replacements)

		moveRobot(grid, moves, robot)

		part2 = sumCoords("[", grid)
	}

	return part1, part2
}

func moveRobot(grid [][]string, moves []Dir, robot Pos) {
	for _, move := range moves {
		movables := make(map[Pos]bool)
		movable := isMovable(robot, move, grid, movables)

		if !movable {
			continue
		}

		chars := make(map[Pos]string)

		for from := range movables {
			chars[from] = charAt(from, grid)
			grid[from.y][from.x] = "."
		}

		for from := range movables {
			to := nextPos(from, move)
			grid[to.y][to.x] = chars[from]
		}

		robot = nextPos(robot, move)
	}
}

func isMovable(pos Pos, dir Dir, grid [][]string, acc map[Pos]bool) bool {
	next := nextPos(pos, dir)
	char, nextChar := charAt(pos, grid), charAt(next, grid)
	left, right := nextPos(pos, LEFT), nextPos(pos, RIGHT)

	if nextChar == "#" {
		acc[pos] = false
	} else if nextChar == "." {
		acc[pos] = true
	} else {
		acc[pos] = isMovable(next, dir, grid, acc)
	}

	if dir == LEFT || dir == RIGHT {
		return acc[pos]
	}

	if _, visited := acc[left]; !visited && char == "]" {
		acc[pos] = acc[pos] && isMovable(left, dir, grid, acc)
	} else if _, visited := acc[right]; !visited && char == "[" {
		acc[pos] = acc[pos] && isMovable(right, dir, grid, acc)
	}

	return acc[pos]
}

func charAt(pos Pos, grid [][]string) string {
	return grid[pos.y][pos.x]
}

func nextPos(pos Pos, dir Dir) Pos {
	return Pos{pos.x + dir.x, pos.y + dir.y}
}

func sumCoords(box string, grid [][]string) int {
	sum := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == box {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func strToDir(s string) Dir {
	switch s {
	case "^":
		return UP
	case ">":
		return RIGHT
	case "v":
		return DOWN
	default:
		return LEFT
	}
}

func parseInput(replacements []string) ([][]string, []Dir, Pos) {
	d := strings.Split(input, "\n\n")
	m := strings.Split(strings.ReplaceAll(d[1], "\n", ""), "")
	str := strings.NewReplacer(replacements...).Replace(d[0])
	lines := strings.Split(str, "\n")

	grid := [][]string{}
	for _, text := range lines {
		grid = append(grid, strings.Split(text, ""))
	}

	moves := []Dir{}
	for _, moveStr := range m {
		moves = append(moves, strToDir(moveStr))
	}

	robot := Pos{}
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "@" {
				robot = Pos{x, y}
			}
		}
	}

	return grid, moves, robot
}

func main() {
	part1, part2 := 0, 0
	sum := 0
	n := 25 // increase samples if benching perf

	for range n {
		start := time.Now()
		part1, part2 = solution()
		sum += int(time.Since(start).Milliseconds())
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
	fmt.Println("avg:", sum/n, "ms")
}
