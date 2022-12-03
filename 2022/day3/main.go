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
	fmt.Println(p1(input), p2(input))
}

func p1(input string) int {
	total := 0

	for _, i := range strings.Split(input, "\n") {
		r1 := []rune(i[:len(i)/2])
		r2 := []rune(i[len(r1):])

		in := util.Intersect(r1, r2)

		total += letterValue(in[0])
	}

	return total
}

func p2(input string) int {
	total := 0
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines)-2; i += 3 {
		r1 := []rune(lines[i])
		r2 := []rune(lines[i+1])
		r3 := []rune(lines[i+2])

		in := util.Intersect(util.Intersect(r1, r2), r3)

		total += letterValue(in[0])
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
