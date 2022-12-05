package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"

	"github.com/maxschuele/advent-of-code/util"
)

func main() {
	var (
		store     [3]string
		p1, p2, i int
	)

	input, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		r1 := []rune(line[:len(line)/2])
		r2 := []rune(line[len(r1):])
		intersection := util.Intersect(r1, r2)
		p1 += letterValue(intersection[0])

		store[i] = line
		if i == 2 {
			r1 = []rune(store[0])
			r2 = []rune(store[1])
			r3 := []rune(store[2])
			intersection = util.Intersect(util.Intersect(r1, r2), r3)
			p2 += letterValue(intersection[0])
			i = 0
			continue
		}
		i++
	}

	fmt.Println(p1, p2)
}

func letterValue(r rune) int {
	if unicode.IsLower(r) {
		return int(r) - 96
	}

	if unicode.IsUpper(r) {
		return int(r) - 38
	}

	panic(r)
}
