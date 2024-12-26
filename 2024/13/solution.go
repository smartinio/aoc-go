package main

import (
	_ "embed"
	"main/perf"
	"regexp"
	"strconv"
	"strings"
)

type Pos struct{ x, y int }
type Machine struct{ ax, ay, bx, by int }

//go:embed input.txt
var input string

func solution() (int, int) {
	part1, part2 := 0, 0
	machines := strings.Split(input, "\n\n")
	re := regexp.MustCompile(`(\d+)`)

	for _, m := range machines {
		lines := strings.Split(m, "\n")

		as := re.FindAllString(lines[0], -1)
		bs := re.FindAllString(lines[1], -1)
		ts := re.FindAllString(lines[2], -1)

		ax, ay := atoi(as[0]), atoi(as[1])
		bx, by := atoi(bs[0]), atoi(bs[1])
		tx, ty := atoi(ts[0]), atoi(ts[1])

		machine := Machine{ax, ay, bx, by}

		{
			target := Pos{tx, ty}
			a, b := solve(target, machine)
			part1 += a*3 + b*1
		}

		{
			target := Pos{10000000000000 + tx, 10000000000000 + ty}
			a, b := solve(target, machine)
			part2 += a*3 + b*1
		}
	}

	return part1, part2
}

func solve(t Pos, m Machine) (int, int) {
	d := det(m)

	if d == 0 {
		return 0, 0
	}

	da := det(Machine{t.x, t.y, m.bx, m.by})
	db := det(Machine{m.ax, m.ay, t.x, t.y})

	a, b := da/d, db/d
	x, y := a*m.ax+b*m.bx, a*m.ay+b*m.by

	if x != t.x || y != t.y {
		return 0, 0
	}

	return a, b
}

func atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func det(m Machine) int {
	return m.ax*m.by - m.bx*m.ay
}

func main() {
	perf.Bench(100, solution)
}
