package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(file)

	var p1, p2 int

	for reader.Scan() {
		f := strings.FieldsFunc(reader.Text(), func(r rune) bool { return r == ',' || r == '-' })

		var a [4]int

		for i := 0; i < 4; i++ {
			a[i], err = strconv.Atoi(f[i])
			if err != nil {
				log.Fatal(err)
			}
		}

		if a[0] >= a[2] && a[1] <= a[3] || a[0] <= a[2] && a[1] >= a[3] {
			p1++
		}

		if a[0] <= a[3] && a[1] >= a[2] {
			p2++
		}

	}

	fmt.Println(p1, p2)
}
