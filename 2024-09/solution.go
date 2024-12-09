package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"time"
)

//go:embed input.txt
var input string

const FREE = -1

type Span struct{ index, space int }

func solution() (int, int) {
	disk := []int{}
	spans := []Span{}
	offset := 0

	for i := range len(input) {
		digit := input[i]
		blocks, _ := strconv.Atoi(string(digit))

		if i%2 == 0 {
			id := i / 2

			for range blocks {
				disk = append(disk, id)
			}
		} else {
			for range blocks {
				disk = append(disk, FREE)
			}
			if blocks > 0 {
				spans = append(spans, Span{offset, blocks})
			}
		}

		offset += blocks
	}

	// part 1
	p1disk := make([]int, len(disk))
	maxLen := len(p1disk) - 1
	copy(p1disk, disk)
	{
		left := slices.Index(disk, FREE)

		for right := maxLen; right >= left; right-- {
			char := disk[right]

			if char == FREE {
				continue
			}

			p1disk[left] = char
			p1disk[right] = FREE

			for p1disk[left] != FREE && left < maxLen {
				left++
			}
		}
	}
	part1 := checksum(p1disk)

	// part 2
	maxSize := int(^uint(0) >> 1)
	p2disk := make([]int, len(disk))
	copy(p2disk, disk)
	{
		right := len(disk) - 1
		left := right - 1

		for left > 0 {
			for disk[right] == FREE && right > 0 {
				right--
			}

			left = right

			for disk[left] == disk[right] && left > 0 {
				left--
			}

			size := right - left

			if size >= maxSize {
				right = left
				continue
			}

			span := slices.IndexFunc(spans, func(span Span) bool {
				return span.index < left && span.space >= size
			})

			if span != -1 {
				offset := spans[span].index

				for i := range size {
					p2disk[offset+i] = disk[right-i]
					p2disk[right-i] = FREE
				}

				spans[span].index += size
				spans[span].space -= size
			} else {
				maxSize = min(maxSize, size)
			}

			right = left
		}
	}
	part2 := checksum(p2disk)

	return part1, part2
}

func checksum(disk []int) int {
	total := 0
	prev := disk[0]

	for i, id := range disk {
		if id != FREE {
			total += id * i
		}

		if id == prev {
			continue
		}

		prev = id
	}

	return total
}

func main() {
	part1, part2 := 0, 0
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
