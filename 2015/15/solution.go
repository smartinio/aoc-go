package main

import (
	_ "embed"
	"main/perf"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Cookie = map[Ingredient]int
type Ingredient struct {
	capacity, durability, flavor, texture, kcal int
}

func solution() (int, int) {
	part1, part2 := 0, 0

	re := regexp.MustCompile(`-?\d+`)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ingredients := []Ingredient{}

	for _, line := range lines {
		n := re.FindAllString(line, -1)
		ing := Ingredient{atoi(n[0]), atoi(n[1]), atoi(n[2]), atoi(n[3]), atoi(n[4])}
		ingredients = append(ingredients, ing)
	}

	part1 = dfs(0, []int{}, ingredients, -1)

	part2 = dfs(0, []int{}, ingredients, 500)

	return part1, part2
}

func dfs(n int, used []int, ingredients []Ingredient, targetKcal int) int {
	if n == len(ingredients) {
		if sum(used) < 100 {
			return 0
		}

		cookie := Cookie{}

		for i := range ingredients {
			cookie[ingredients[i]] = used[i]
		}

		if targetKcal != -1 && kcal(cookie) != targetKcal {
			return 0
		}

		return score(cookie)
	}

	highest := 0
	remaining := 100 - sum(used)

	for i := range remaining {
		score := dfs(n+1, append(used, i+1), ingredients, targetKcal)
		highest = max(highest, score)
	}

	return highest
}

func sum(used []int) (total int) {
	for _, val := range used {
		total += val
	}
	return total
}

func kcal(cookie Cookie) (total int) {
	for s, count := range cookie {
		total += count * s.kcal
	}
	return total
}

func score(cookie Cookie) int {
	t := Ingredient{}

	for ing, count := range cookie {
		t.capacity += count * ing.capacity
		t.durability += count * ing.durability
		t.flavor += count * ing.flavor
		t.texture += count * ing.texture
	}

	t.capacity = max(0, t.capacity)
	t.durability = max(0, t.durability)
	t.flavor = max(0, t.flavor)
	t.texture = max(0, t.texture)

	return t.capacity * t.durability * t.flavor * t.texture
}

func atoi(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func main() {
	perf.Bench(1, solution)
}
