package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/maxschuele/advent-of-code/util"
)

//go:embed input.txt
var input string

func main() {
	var (
		p1, p2 int
		lines  = strings.Split(input, "\n")
	)

	for _, line := range lines {
		assignments := strings.Split(line, ",")

		a0 := util.ArrAtoi(strings.Split(assignments[0], "-"))
		a1 := util.ArrAtoi(strings.Split(assignments[1], "-"))

		range1 := util.IntRange(a0[0], a0[1])
		range2 := util.IntRange(a1[0], a1[1])

		intersection := util.Intersect(range1, range2)

		if cmp.Equal(range1, intersection) || cmp.Equal(range2, intersection) {
			p1++
		}

		if len(intersection) > 0 {
			p2++
		}
	}

	fmt.Println(p1, p2)
}
