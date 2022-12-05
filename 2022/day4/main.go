package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var p1, p2 int

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		f := strings.FieldsFunc(scanner.Text(), func(r rune) bool { return r == ',' || r == '-' })

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
