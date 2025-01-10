package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"main/perf"
	"maps"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func solution() (part1 int, part2 int) {
	salt := strings.TrimSpace(input)

	part1 = solve(salt, 1)
	part2 = solve(salt, 2017)

	return part1, part2
}

func solve(salt string, times int) int {
	triples := map[byte][]int{}
	keys := map[int]bool{}

	for i := 0; true; i++ {
		text := fmt.Sprintf("%s%d", salt, i)
		stream := hash(text, times)
		newTrips := findRepeats(stream, 3, 1)

		if len(newTrips) == 0 {
			continue
		}

		for _, char := range findRepeats(stream, 5, -1) {
			t := triples[char]

			for _, key := range filterGTE(t, i-1000) {
				keys[key] = true
			}

			triples[char] = t[len(t):]
		}

		if len(keys) >= 64 {
			all := slices.Collect(maps.Keys(keys))
			slices.Sort(all)
			return all[63]
		}

		char := newTrips[0]
		triples[char] = append(triples[char], i)
	}

	return -1
}

func hash(text string, times int) string {
	for range times {
		text = getMd5(text)
	}

	return text
}

func filterGTE(slice []int, minimum int) []int {
	gti := slices.IndexFunc(slice, func(v int) bool { return v >= minimum })

	if gti == -1 {
		return []int{}
	}

	return slice[gti:]
}

func findRepeats(s string, n, max int) []byte {
	if n > len(s) {
		return []byte{}
	}

	for i := 0; i < n-1; i++ {
		if s[i] != s[i+1] {
			return findRepeats(s[1:], n, max-1)
		}
	}

	repeats := []byte{s[0]}

	if max == 1 {
		return repeats
	}

	return append(repeats, findRepeats(s[n:], n, max-1)...)
}

func getMd5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	perf.Bench(1, solution)
}
