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
var prefix = strings.HasPrefix

func solution() (part1 string, part2 string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	numRe := regexp.MustCompile(`(\d+)`)
	letRe := regexp.MustCompile(`letter ([a-z])`)

	// part 1
	{
		l := []byte("abcdefgh")

		for _, line := range lines {
			switch {
			case prefix(line, "swap position"):
				n := utils.FindAllIntGroups(numRe, line)
				l[n[0]], l[n[1]] = l[n[1]], l[n[0]]

			case prefix(line, "swap letter"):
				letters := utils.FindAllStringGroups(letRe, line)
				a := slices.Index(l, letters[0][0])
				b := slices.Index(l, letters[1][0])
				l[a], l[b] = l[b], l[a]

			case prefix(line, "rotate left"):
				n := utils.FindAllIntGroups(numRe, line)[0]
				l = append(l[n:], l[:n]...)

			case prefix(line, "rotate right"):
				n := utils.FindAllIntGroups(numRe, line)[0]
				l = append(l[len(l)-n:], l[:len(l)-n]...)

			case prefix(line, "rotate based"):
				letter := utils.FindAllStringGroups(letRe, line)[0][0]
				idx := slices.Index(l, letter)
				steps := idx + 1 + utils.Btoi(idx >= 4)
				i := wrap(len(l)-steps, len(l))
				l = append(l[i:], l[:i]...)

			case prefix(line, "reverse"):
				n := utils.FindAllIntGroups(numRe, line)
				slices.Reverse(l[n[0] : n[1]+1])

			case prefix(line, "move"):
				n := utils.FindAllIntGroups(numRe, line)
				char := l[n[0]]
				l = append(l[:n[0]], l[n[0]+1:]...)
				tail := string(l[n[1]:])
				l = append(l[:n[1]], char)
				l = append(l, tail...)
			}
		}

		part1 = string(l)
	}

	// part 2
	{
		l := []byte("fbgdceah")
		slices.Reverse(lines)

		for _, line := range lines {
			switch {
			case prefix(line, "swap position"): // no change from p1: self-reversible
				n := utils.FindAllIntGroups(numRe, line)
				l[n[0]], l[n[1]] = l[n[1]], l[n[0]]

			case prefix(line, "swap letter"): // no change from p1: self-reversible
				letters := utils.FindAllStringGroups(letRe, line)
				a := slices.Index(l, letters[0][0])
				b := slices.Index(l, letters[1][0])
				l[a], l[b] = l[b], l[a]

			case prefix(line, "rotate right"): // swapped with rotate left
				n := utils.FindAllIntGroups(numRe, line)[0]
				l = append(l[n:], l[:n]...)

			case prefix(line, "rotate left"): // swapped with rotate right
				n := utils.FindAllIntGroups(numRe, line)[0]
				l = append(l[len(l)-n:], l[:len(l)-n]...)

			case prefix(line, "rotate based"):
				// mappings found by printing before/after index in part1
				// when rotating 'x.......', '.x......' etc based on pos of x
				rb := [8]int{7, 0, 4, 1, 5, 2, 6, 3}
				letter := utils.FindAllStringGroups(letRe, line)[0][0]
				idx := slices.Index(l, letter)
				n := wrap(idx-rb[idx], len(l))
				l = append(l[n:], l[:n]...)

			case prefix(line, "reverse"): // no change from p1: self-reversible
				n := utils.FindAllIntGroups(numRe, line)
				slices.Reverse(l[n[0] : n[1]+1])

			case prefix(line, "move"):
				n := utils.FindAllIntGroups(numRe, line)
				slices.Reverse(n) // added
				char := l[n[0]]
				l = append(l[:n[0]], l[n[0]+1:]...)
				tail := string(l[n[1]:])
				l = append(l[:n[1]], char)
				l = append(l, tail...)
			}
		}

		part2 = string(l)
	}

	return part1, part2
}

func wrap(i, max int) int {
	return ((i % max) + max) % max
}

func main() {
	perf.Bench(1, solution)
}
