package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/maxschuele/advent-of-code/util"
)

//go:embed input.txt
var input string

func main() {
	var (
		elves   []int
		counter int
	)

	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			elves = append(elves, counter)

			counter = 0
			continue
		}

		counter += util.Atoi(v)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	max3 := elves[0] + elves[1] + elves[2]

	fmt.Println(elves[0], max3)
}
