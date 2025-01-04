package main

import (
	_ "embed"
	"main/perf"
	"main/utils"
	"math"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

type Context struct {
	bossDmg int
	part    uint8
	cache   map[Key]float64
}

type Mana struct {
	spent, pool float64
}

type Effects struct {
	shield, poison, recharge int
}

type Key [8]int

func solution() (part1 float64, part2 float64) {
	bossHp, bossDmg := parseInput()

	// part 1
	{
		ctx := Context{
			bossDmg: bossDmg,
			part:    1,
			cache:   map[Key]float64{},
		}

		part1 = dfs(50, bossHp, true, Mana{0, 500}, Effects{}, ctx)
	}

	// part 2
	{
		ctx := Context{
			bossDmg: bossDmg,
			part:    2,
			cache:   map[Key]float64{},
		}

		part2 = dfs(50, bossHp, true, Mana{pool: 500}, Effects{}, ctx)
	}

	return part1, part2
}

func dfs(hp, bossHp int, myTurn bool, m Mana, e Effects, ctx Context) float64 {
	if ctx.part == 2 {
		hp--
	}

	key := Key{hp, bossHp, utils.Btoi(myTurn), int(m.spent), int(m.pool), e.shield, e.poison, e.recharge}

	if cached, ok := ctx.cache[key]; ok {
		return cached
	}

	if bossHp <= 0 {
		return m.spent
	}

	if hp <= 0 {
		return math.Inf(1)
	}

	// apply effects
	armor := 0
	if e.shield > 0 {
		armor = 7
		e.shield--
	}

	if e.poison > 0 {
		bossHp -= 3
		e.poison--
	}

	if e.recharge > 0 {
		m.pool += 101
		e.recharge--
	}

	// boss turn
	if !myTurn {
		dmg := max(1, ctx.bossDmg-armor)
		return dfs(hp-dmg, bossHp, true, m, e, ctx)
	}

	// my turn, pick best option
	cost := math.Inf(1)

	if mana, effects, ok := cast("magic missile", e, m); ok {
		cost = min(cost, dfs(hp, bossHp-4, false, mana, effects, ctx))
	}

	if mana, effects, ok := cast("drain", e, m); ok {
		cost = min(cost, dfs(hp+2, bossHp-2, false, mana, effects, ctx))
	}

	if mana, effects, ok := cast("shield", e, m); ok && e.shield == 0 {
		cost = min(cost, dfs(hp, bossHp, false, mana, effects, ctx))
	}

	if mana, effects, ok := cast("poison", e, m); ok && e.poison == 0 {
		cost = min(cost, dfs(hp, bossHp, false, mana, effects, ctx))
	}

	if mana, effects, ok := cast("recharge", e, m); ok && e.recharge == 0 {
		cost = min(cost, dfs(hp, bossHp, false, mana, effects, ctx))
	}

	ctx.cache[key] = cost

	return cost
}

func cast(spell string, e Effects, m Mana) (Mana, Effects, bool) {
	if mana, ok := spend(53, m); ok && spell == "magic missile" {
		return mana, e, true
	}

	if mana, ok := spend(73, m); ok && spell == "drain" {
		return mana, e, true
	}

	if mana, ok := spend(113, m); ok && spell == "shield" {
		e.shield = 6
		return mana, e, true
	}

	if mana, ok := spend(173, m); ok && spell == "poison" {
		e.poison = 6
		return mana, e, true
	}

	if mana, ok := spend(229, m); ok && spell == "recharge" {
		e.recharge = 5
		return mana, e, true
	}

	return m, e, false
}

func spend(amount float64, m Mana) (Mana, bool) {
	if m.pool >= amount {
		m.spent += amount
		m.pool -= amount
		return m, true
	}
	return m, false
}

func parseInput() (bossHp, bossDmg int) {
	re := regexp.MustCompile(`\d+`)
	stats := re.FindAllString(input, -1)

	bossHp, _ = strconv.Atoi(stats[0])
	bossDmg, _ = strconv.Atoi(stats[1])

	return bossHp, bossDmg
}

func main() {
	perf.Bench(1, solution)
}
