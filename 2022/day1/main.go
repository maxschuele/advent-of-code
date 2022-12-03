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
		sums    []int
		counter int
	)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			sums = append(sums, counter)

			counter = 0
			continue
		}

		counter += util.Atoi(line)
	}

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	fmt.Println(sums[0], sums[0]+sums[1]+sums[2])
}
