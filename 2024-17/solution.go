package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

//go:embed input.txt
var input string

type Opcode int
type Regs struct{ A, B, C int }
type Program struct {
	ptr   int
	initA int
	orig  string
	out   []string
	regs  *Regs
}

func (p *Program) Output(value int) {
	p.out = append(p.out, fmt.Sprintf("%d", value))
}

func (p *Program) Approaches() bool {
	if len(p.out) > 1+len(p.orig)/2 {
		return false
	}

	origParts := strings.Split(p.orig, ",")

	for i, part := range p.out {
		if origParts[i] != part {
			return false
		}
	}

	return true
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

// shitty brute force-ish solution
// cant be bothered improving because bitwise operations hurt my brain
func solution() (string, int) {
	part1, part2 := "", 0

	s := strings.Split(input, "\n\n")
	re := regexp.MustCompile(`\d+`)
	regsRaw, inputs := s[0], s[1]
	rs := re.FindAllString(regsRaw, -1)
	ns := re.FindAllString(inputs, -1)

	// part 1
	{
		prg := Program{}
		regs := Regs{}
		prg.regs = &regs
		regs.A, _ = strconv.Atoi(rs[0])
		regs.B, _ = strconv.Atoi(rs[1])
		regs.C, _ = strconv.Atoi(rs[2])
		prg.initA = regs.A

		for prg.ptr <= len(ns)-2 {
			code, _ := strconv.Atoi(ns[prg.ptr])
			opd, _ := strconv.Atoi(ns[prg.ptr+1])
			perform(code, opd, &regs, &prg)
		}

		part1 = strings.Join(prg.out, ",")
	}

	// part2
	{
		concurrency := 500000
		res := map[int]bool{}
		orig := strings.TrimSpace(strings.Split(inputs, " ")[1])
		wg := sync.WaitGroup{}

		j := 1
		for {
			for range concurrency {
				// found this number by printing the first cases that started with the target output
				// this one was early and matched lots of digits so using it as base for brute force
				startingPoint := "1011011010110111010111101"
				pad := strconv.FormatInt(int64(j), 2)
				of, _ := strconv.ParseInt(pad+startingPoint, 2, 64)
				offset := int(of)

				wg.Add(1)
				compute := func(initial int) {
					defer wg.Done()
					prg := Program{orig: orig}
					regs := Regs{}
					prg.regs = &regs
					regs.A = initial
					regs.B, _ = strconv.Atoi(rs[1])
					regs.C, _ = strconv.Atoi(rs[2])
					prg.initA = regs.A

					for prg.ptr <= len(ns)-2 {
						code, _ := strconv.Atoi(ns[prg.ptr])
						opd, _ := strconv.Atoi(ns[prg.ptr+1])
						perform(code, opd, &regs, &prg)
						if len(prg.out) > 0 && !prg.Approaches() {
							return
						}
					}

					joined := strings.Join(prg.out, ",")

					if joined == orig {
						res[initial] = true
					}
				}

				go compute(offset)
				j++
			}

			wg.Wait()

			if len(res) > 0 {
				break
			}
		}

		part2 = int(^uint(0) >> 1) // max int

		for initial := range res {
			part2 = min(part2, initial)
		}
	}

	return part1, part2
}

func perform(code int, opd int, regs *Regs, prg *Program) {
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
		prg.Output(combo() % 8)
	case BDV:
		regs.B = divide()
	case CDV:
		regs.C = divide()
	}

	prg.ptr += jumps
}

func main() {
	part1, part2 := "", 0
	sum := 0
	n := 1 // increase samples if benching perf

	for range n {
		start := time.Now()
		part1, part2 = solution()
		sum += int(time.Since(start).Milliseconds())
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
	fmt.Println("avg:", sum/n, "ms")
}
