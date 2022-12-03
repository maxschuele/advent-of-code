package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines), part2(lines))
}

func part1(lines []string) int {
	res := map[string]map[string]int{
		"A": {"X": 3, "Y": 6, "Z": 0},
		"B": {"X": 0, "Y": 3, "Z": 6},
		"C": {"X": 6, "Y": 0, "Z": 3},
	}

	vals := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	total := 0
	for _, line := range lines {
		fields := strings.Fields(line)

		total += vals[fields[1]] + res[fields[0]][fields[1]]
	}

	return total
}

func part2(lines []string) int {
	res := map[string]map[string]int{
		"A": {"X": 3, "Y": 1, "Z": 2},
		"B": {"X": 1, "Y": 2, "Z": 3},
		"C": {"X": 2, "Y": 3, "Z": 1},
	}

	vals := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	total := 0
	for _, line := range lines {
		fields := strings.Fields(line)

		total += vals[fields[1]] + res[fields[0]][fields[1]]
	}

	return total
}
