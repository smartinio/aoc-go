package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"main/perf"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func solution() (part1 string, part2 string) {
	doorId := strings.TrimSpace(input)

	// part 1
	{
		bytes := []byte{}

		for i := 0; len(bytes) < 8; i++ {
			hash := getMd5(fmt.Sprintf("%s%d", doorId, i))

			if strings.HasPrefix(hash, "00000") {
				bytes = append(bytes, hash[5])
			}
		}

		part1 = string(bytes)
	}

	// part 2
	{
		bytes, found := make([]byte, 8), map[int]bool{}

		for i := 0; len(found) < 8; i++ {
			hash := getMd5(fmt.Sprintf("%s%d", doorId, i))

			if strings.HasPrefix(hash, "00000") {
				pos, err := strconv.Atoi(hash[5:6])

				if err == nil && pos >= 0 && pos <= 7 && !found[pos] {
					bytes[pos] = hash[6]
					found[pos] = true
				}
			}
		}

		part2 = string(bytes)
	}

	return part1, part2
}

func getMd5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	perf.Bench(1, solution)
}
