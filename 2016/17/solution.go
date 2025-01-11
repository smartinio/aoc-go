package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"main/perf"
	"math"
	"strings"
)

type Dir struct{ x, y int }
type Pos struct{ x, y int }
type Result struct {
	len  float64
	path string
}

//go:embed input.txt
var input string
var dirs = map[Dir]string{{0, -1}: "U", {1, 0}: "R", {0, 1}: "D", {-1, 0}: "L"}
var upDownLeftRight = []Dir{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
var OUTSIDE = Pos{-1, -1}
var LONGEST = Result{math.Inf(1), ""}
var SHORTEST = Result{0, ""}

func solution() (part1 string, part2 float64) {
	seed := strings.TrimSpace(input)

	part1 = dfs1(Pos{0, 0}, seed, "").path
	part2 = dfs2(Pos{0, 0}, seed, "").len

	return part1, part2
}

func dfs1(pos Pos, seed, path string) Result {
	if pos == (Pos{3, 3}) {
		return Result{float64(len(path)), path}
	}

	best := LONGEST

	for _, d := range adj(seed, path) {
		if next := getNext(pos, d); next != OUTSIDE {
			result := dfs1(next, seed, path+dirs[d])
			if result.len < best.len {
				best = result
			}
		}
	}

	return best
}

func dfs2(pos Pos, seed, path string) Result {
	if pos == (Pos{3, 3}) {
		return Result{float64(len(path)), path}
	}

	best := SHORTEST

	for _, d := range adj(seed, path) {
		if next := getNext(pos, d); next != OUTSIDE {
			result := dfs2(next, seed, path+dirs[d])
			if result.len > best.len {
				best = result
			}
		}
	}

	return best
}

func adj(seed, path string) (a []Dir) {
	hash := getMd5(fmt.Sprintf("%s%s", seed, path))[:4]

	for i, d := range upDownLeftRight {
		if hash[i] > 'a' && hash[i] < 'g' {
			a = append(a, d)
		}
	}

	return a
}

func getNext(pos Pos, d Dir) Pos {
	next := Pos{pos.x + d.x, pos.y + d.y}

	if next.x < 0 || next.y < 0 || next.x > 3 || next.y > 3 {
		return OUTSIDE
	}

	return next
}

func getMd5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	perf.Bench(1, solution)
}
