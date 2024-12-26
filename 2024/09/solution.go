package main

import (
	_ "embed"
	"main/perf"
	"slices"
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
		blocks := int(input[i]) - '0'
		id := FREE

		if i%2 == 0 {
			id = i / 2
		} else if blocks > 0 {
			spans = append(spans, Span{offset, blocks})
		} else {
			continue
		}

		disk = append(disk, slices.Repeat([]int{id}, blocks)...)
		offset += blocks
	}

	// part 1
	p1disk := append([]int{}, disk...)
	maxLen := len(p1disk) - 1
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
	p2disk := append([]int{}, disk...)
	maxSize := int(^uint(0) >> 1)
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
	perf.Bench(20, solution)
}
