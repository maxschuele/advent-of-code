package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(p1(input), p2(input))
}

func p1(input string) int {
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
	for _, i := range strings.Split(input, "\n") {
		f := strings.Fields(i)

		total += vals[f[1]] + res[f[0]][f[1]]
	}

	return total
}

func p2(input string) int {
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
	for _, i := range strings.Split(input, "\n") {
		f := strings.Fields(i)

		total += vals[f[1]] + res[f[0]][f[1]]
	}

	return total
}
