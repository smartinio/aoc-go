package main

import (
	_ "embed"
	"main/perf"
	"math"
	"regexp"
	"slices"
	"strconv"
)

//go:embed input.txt
var input string

type Boss struct {
	hp, dmg, armor int
}

type Item struct {
	name             string
	cost, dmg, armor int
}

type Context struct {
	boss     Boss
	base     float64
	finished func(int) bool
	best     func(float64, float64) float64
}

var weapons = []Item{
	{"Dagger", 8, 4, 0},
	{"Shortsword", 10, 5, 0},
	{"Warhammer", 25, 6, 0},
	{"Longsword", 40, 7, 0},
	{"Greataxe", 74, 8, 0},
}

var armors = []Item{
	{"Leather", 13, 0, 1},
	{"Chainmail", 31, 0, 2},
	{"Splintmail", 53, 0, 3},
	{"Bandedmail", 75, 0, 4},
	{"Platemail", 102, 0, 5},
}

var rings = []Item{
	{"Damage +1", 25, 1, 0},
	{"Damage +2", 50, 2, 0},
	{"Damage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
}

func solution() (part1 float64, part2 float64) {
	boss := parseInput()

	// part 1
	{
		ctx := Context{
			boss:     boss,
			base:     math.Inf(1),
			finished: func(hp int) bool { return hp > 0 },
			best:     func(a, b float64) float64 { return min(a, b) },
		}

		part1 = dfs([]Item{}, ctx)
	}

	// part 2
	{
		ctx := Context{
			boss:     boss,
			base:     0,
			finished: func(hp int) bool { return hp <= 0 },
			best:     func(a, b float64) float64 { return max(a, b) },
		}

		part2 = dfs([]Item{}, ctx)
	}

	return part1, part2
}

func dfs(items []Item, ctx Context) float64 {
	cost := ctx.base

	// required - buy weapon
	if len(items) == 0 {
		for _, w := range weapons {
			opt := dfs(append(items, w), ctx)
			cost = ctx.best(cost, opt)
		}
		return cost // cannot proceed to other options without weapon
	}

	// option 1 - fight
	opt1 := fight(items, ctx)
	cost = ctx.best(cost, opt1)

	// option 2 - buy armor
	if !slices.ContainsFunc(items, func(item Item) bool { return slices.Contains(armors, item) }) {
		for _, a := range armors {
			opt2 := dfs(append(items, a), ctx)
			cost = ctx.best(cost, opt2)
		}
	}

	// option 3 - buy rings
	if countRings(items) < 2 {
		for _, r := range rings {
			if !slices.Contains(items, r) {
				opt3 := dfs(append(items, r), ctx)
				cost = ctx.best(cost, opt3)
			}
		}
	}

	return cost
}

func fight(items []Item, ctx Context) float64 {
	dmg, armor, spent := stats(items)
	playerTurn, hp, b := true, 100, ctx.boss

	for b.hp > 0 && hp > 0 {
		if playerTurn {
			b.hp -= max(1, dmg-b.armor)
		} else {
			hp -= max(1, b.dmg-armor)
		}
		playerTurn = !playerTurn
	}

	if ctx.finished(hp) {
		return float64(spent)
	}

	return ctx.base
}

func stats(items []Item) (dmg int, armor int, spent int) {
	for _, item := range items {
		dmg += item.dmg
		armor += item.armor
		spent += item.cost
	}
	return dmg, armor, spent
}

func countRings(items []Item) (count int) {
	for _, item := range items {
		if slices.Contains(rings, item) {
			count++
		}
	}
	return count
}

func parseInput() (boss Boss) {
	re := regexp.MustCompile(`\d+`)
	stats := re.FindAllString(input, -1)

	boss.hp, _ = strconv.Atoi(stats[0])
	boss.dmg, _ = strconv.Atoi(stats[1])
	boss.armor, _ = strconv.Atoi(stats[2])

	return boss
}

func main() {
	perf.Bench(1, solution)
}
