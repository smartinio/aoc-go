package main

import (
	_ "embed"
	"main/perf"
	"main/utils"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Seq []string
type Graph [][]rune
type Pos struct{ x, y int }
type Dir struct{ x, y int }
type Vertex struct {
	pos  Pos
	path string
}
type Key struct {
	from, to rune
	depth    int
}

const NIL = -1

//go:embed input.txt
var input string

var numpad = Graph{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{NIL, '0', 'A'},
}

var keypad = Graph{
	{NIL, '^', 'A'},
	{'<', 'v', '>'},
}

func solution() (int, int) {
	part1, part2 := 0, 0
	codes := strings.Split(strings.TrimSpace(input), "\n")
	re := regexp.MustCompile(`\d+`)

	for _, code := range codes {
		numeric, _ := strconv.Atoi(re.FindString(code))
		part1 += numeric * minSeqLen(code, 2)
		part2 += numeric * minSeqLen(code, 25)
	}

	return part1, part2
}

func minSeqLen(code string, numRobots int) int {
	curr := 'A'
	options := [][]string{}

	for _, target := range code {
		sequences := bfs(curr, target, numpad)
		curr = target
		options = append(options, sequences)
	}

	sequences := join(utils.CartesianProduct(options))
	cache := map[Key]int{}
	minLen := math.Inf(1)

	for _, seq := range sequences {
		seq = "A" + seq
		length := 0
		for i := range seq[:len(seq)-1] {
			a, b := seq[i], seq[i+1]
			length += dfs(rune(a), rune(b), numRobots, cache)
		}
		minLen = min(minLen, float64(length))
	}

	return int(minLen)
}

func dfs(from, to rune, depth int, cache map[Key]int) int {
	key := Key{from, to, depth}

	if depth == 1 {
		return len(bfs(from, to, keypad)[0])
	}

	if cached, ok := cache[key]; ok {
		return cached
	}

	sequences := bfs(from, to, keypad)
	minimum := math.Inf(1)

	for _, seq := range sequences {
		seq = "A" + seq
		length := 0
		for i := range seq[:len(seq)-1] {
			a, b := seq[i], seq[i+1]
			length += dfs(rune(a), rune(b), depth-1, cache)
		}
		minimum = min(minimum, float64(length))
	}

	cache[key] = int(minimum)

	return int(minimum)
}

func bfs(source rune, target rune, graph Graph) []string {
	dirs := []Dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	q := []Vertex{{getPos(source, graph), ""}}
	paths := []string{}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if len(paths) > 0 && len(curr.path) > len(paths[0]) {
			continue
		}

		if charAt(curr.pos, graph) == target {
			curr.path += "A"
			paths = append(paths, curr.path)
			continue
		}

		for _, d := range dirs {
			next := Pos{curr.pos.x + d.x, curr.pos.y + d.y}
			key := string(dirKey(d))

			if charAt(next, graph) != NIL && uniqueCount(curr.path) <= 2 {
				q = append(q, Vertex{next, curr.path + key})
			}
		}
	}

	return paths
}

func uniqueCount(str string) int {
	seen := map[rune]bool{}
	for _, char := range str {
		seen[char] = true
	}
	return len(seen)
}

func join(seqs [][]string) []string {
	out := []string{}
	for _, seq := range seqs {
		out = append(out, strings.Join(seq, ""))
	}
	return out
}

func dirKey(dir Dir) rune {
	return map[Dir]rune{
		{0, -1}: '^',
		{1, 0}:  '>',
		{0, 1}:  'v',
		{-1, 0}: '<',
	}[dir]
}

func getPos(char rune, graph Graph) Pos {
	for y, row := range graph {
		for x, col := range row {
			if col == char {
				return Pos{x, y}
			}
		}
	}
	return Pos{-1, -1}
}

func charAt(pos Pos, graph Graph) rune {
	x, y := pos.x, pos.y

	if x < 0 || y < 0 || x >= len(graph[0]) || y >= len(graph) {
		return NIL
	}

	return graph[y][x]
}

func main() {
	perf.Bench(10, solution)
}
