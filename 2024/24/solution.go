package main

import (
	_ "embed"
	"fmt"
	"main/perf"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Gate struct {
	a, b, op string
}

func solution() (int, string) {
	part1, part2 := 0, ""

	inputs, gates, outputs := parseInput()

	for i, wire := range outputs {
		digit := dfs(wire, gates, inputs)
		part1 = (digit << i) | part1
	}

	// set last arg to true to generate graphViz output
	graphViz(inputs, gates, false)

	// from manually looking at the graph via https://dreampuf.github.io/GraphvizOnline
	part2 = "jgb,rkf,rrs,rvc,vcg,z09,z20,z24"

	return part1, part2
}

func dfs(wire string, gates map[string]Gate, inputs map[string]int) int {
	if value, ok := inputs[wire]; ok {
		return value
	}

	gate := gates[wire]
	a, b := dfs(gate.a, gates, inputs), dfs(gate.b, gates, inputs)

	switch gate.op {
	case "AND":
		return a & b
	case "XOR":
		return a ^ b
	case "OR":
		return a | b
	}

	return -1
}

func graphViz(ins map[string]int, outs map[string]Gate, print bool) {
	inputs, wires, gates := "// inputs\n", "// wires\n", "// gates\n"

	shapes := map[string]string{"XOR": "invhouse", "OR": "box", "AND": "egg"}
	colors := map[string]string{"XOR": "red", "OR": "yellow", "AND": "green"}

	for k, v := range ins {
		inputs += fmt.Sprintf("%s [label=\"%s=%d\", shape=none];\n", k, k, v)
	}

	i := 1
	for out, v := range outs {
		wires += fmt.Sprintf("%s [label=\"%s\", shape=none];\n", v.a, v.a)
		wires += fmt.Sprintf("%s [label=\"%s\", shape=none];\n", v.b, v.b)
		gates += fmt.Sprintf("gate%d [label=\"%s\", style=\"filled\", ", i, v.op)
		gates += fmt.Sprintf("shape=\"%s\", fillcolor=\"%s\"];\n", shapes[v.op], colors[v.op])
		gates += fmt.Sprintf("%s -> gate%d;\n", v.a, i)
		gates += fmt.Sprintf("%s -> gate%d;\n", v.b, i)
		gates += fmt.Sprintf("gate%d -> %s;\n\n", i, out)
		i++
	}

	if print {
		fmt.Println(inputs)
		fmt.Println()
		fmt.Println(wires)
		fmt.Println()
		fmt.Println(gates)
	}
}

func parseInput() (map[string]int, map[string]Gate, []string) {
	re := regexp.MustCompile(`[A-Za-z0-9]+`)
	s := strings.Split(strings.TrimSpace(input), "\n\n")
	is := re.FindAllString(s[0], -1)
	gs := re.FindAllString(s[1], -1)

	inputs := map[string]int{}
	gates := map[string]Gate{}
	outputs := []string{}

	for i := 0; i < len(gs)-3; i += 4 {
		a, op, b, out := gs[i], gs[i+1], gs[i+2], gs[i+3]
		gates[out] = Gate{a, b, op}

		if out[0] == 'z' {
			outputs = append(outputs, out)
		}
	}

	for i := 0; i < len(is)-1; i += 2 {
		key, value := is[i], is[i+1]
		inputs[key], _ = strconv.Atoi(value)
	}

	slices.Sort(outputs)

	return inputs, gates, outputs
}

func main() {
	perf.Bench(1, solution)
}
