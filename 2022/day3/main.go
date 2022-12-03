package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"

	"github.com/maxschuele/advent-of-code/util"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines), part2(lines))
}

func part1(lines []string) int {
	total := 0

	for _, line := range lines {
		r1 := []rune(line[:len(line)/2])
		r2 := []rune(line[len(r1):])

		intersection := util.Intersect(r1, r2)

		total += letterValue(intersection[0])
	}

	return total
}

func part2(lines []string) int {
	total := 0

	for i := 0; i < len(lines)-2; i += 3 {
		r1 := []rune(lines[i])
		r2 := []rune(lines[i+1])
		r3 := []rune(lines[i+2])

		intersection := util.Intersect(util.Intersect(r1, r2), r3)

		total += letterValue(intersection[0])
	}

	return total
}

func letterValue(r rune) int {
	if unicode.IsLetter(r) {
		if unicode.IsLower(r) {
			return int(r) - 96
		}

		if unicode.IsUpper(r) {
			return int(r) - 38
		}
	}

	panic(r)
}
