package main

import (
	_ "embed"
	"fmt"
	"main/perf"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Regs struct{ A, B, C int }
type Program struct {
	ptr   int
	initA int
	out   []string
	regs  *Regs
}

const (
	ADV = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

func solution() (string, int) {
	part1, part2 := "", 0

	s := strings.Split(input, "\n\n")
	re := regexp.MustCompile(`\d+`)
	rawRegs, rawInputs := s[0], s[1]
	inputs := re.FindAllString(rawInputs, -1)

	// part 1
	{
		regs := re.FindAllString(rawRegs, -1)
		a, _ := strconv.Atoi(regs[0])
		b, _ := strconv.Atoi(regs[1])
		c, _ := strconv.Atoi(regs[2])

		part1 = getOutput(a, b, c, inputs)
	}

	// part2
	{
		target := strings.TrimSpace(strings.Split(rawInputs, " ")[1])

		part2 = bt(0, target, inputs)
	}

	return part1, part2
}

func bt(a int, target string, inputs []string) int {
	for i := range 8 {
		A := (a << 3) | i
		output := getOutput(A, 0, 0, inputs)

		if output == target {
			return A
		}

		if strings.HasSuffix(target, output) {
			if next := bt(A, target, inputs); next != -1 {
				return next
			}
		}
	}

	return -1
}

func getOutput(regA int, regB int, regC int, inputs []string) string {
	prg := Program{regs: &Regs{regA, regB, regC}}
	prg.initA = regA

	for prg.ptr <= len(inputs)-2 {
		code, _ := strconv.Atoi(inputs[prg.ptr])
		opd, _ := strconv.Atoi(inputs[prg.ptr+1])
		execute(code, opd, &prg)
	}

	return strings.Join(prg.out, ",")
}

func execute(code int, opd int, prg *Program) {
	regs := prg.regs
	jumps := 2

	combo := func() int {
		switch opd {
		case 4:
			return regs.A
		case 5:
			return regs.B
		case 6:
			return regs.C
		default:
			return opd
		}
	}

	divide := func() int {
		operand := combo()
		num := float64(regs.A)
		den := math.Exp2(float64(operand))

		return int(math.Trunc(num / den))
	}

	switch code {
	case ADV:
		regs.A = divide()
	case BXL:
		regs.B ^= opd
	case BST:
		regs.B = combo() % 8
	case JNZ:
		if regs.A != 0 {
			prg.ptr = opd
			jumps = 0
		}
	case BXC:
		regs.B ^= regs.C
	case OUT:
		digit := fmt.Sprintf("%d", combo()%8)
		prg.out = append(prg.out, digit)
	case BDV:
		regs.B = divide()
	case CDV:
		regs.C = divide()
	}

	prg.ptr += jumps
}

func main() {
	perf.Bench(500, solution)
}
