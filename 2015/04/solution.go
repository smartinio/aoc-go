package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"main/perf"
	"strings"
)

//go:embed input.txt
var input string

func getMd5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func solution() (int, int) {
	part1, part2, i := 0, 0, 0
	seed := strings.TrimSpace(input)

	for {
		str := fmt.Sprintf("%s%d", seed, i)
		hash := getMd5(str)
		if strings.HasPrefix(hash, "00000") && part1 == 0 {
			part1 = i
		}
		if strings.HasPrefix(hash, "000000") {
			part2 = i
			break
		}
		i++
	}

	return part1, part2
}

func main() {
	perf.Bench(1, solution)
}
