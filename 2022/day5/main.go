package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"container/list"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	crates := initCrates(scanner)
	move(scanner, crates)

	var p1 strings.Builder

	for i := 1; i <= len(crates); i++ {
		p1.WriteString(string(crates[i].Front().Value.(uint8)))
	}

	fmt.Println(p1.String())
}

func initCrates(s *bufio.Scanner) map[int]*list.List {
	crates := make(map[int]*list.List)

	for s.Scan() {
		line := s.Text()

		if line[1] == '1' {
			break
		}

		crateIndex := 1
		for i := 1; i < len(line); i += 4 {
			if line[i] != ' ' {
				if crates[crateIndex] == nil {
					crates[crateIndex] = list.New()
				}

				crates[crateIndex].PushBack(line[i])
			}

			crateIndex++
		}
	}

	s.Scan() //get rid of the empty line

	return crates
}

func move(s *bufio.Scanner, crates map[int]*list.List) {
	for s.Scan() {
		line := s.Text()
		re := regexp.MustCompile(`\d[\d,]*`)

		instructions := re.FindAllString(line, -1)
		moves, _ := strconv.Atoi(instructions[0])
		from, _ := strconv.Atoi(instructions[1])
		to, _ := strconv.Atoi(instructions[2])

		for i := 0; i < moves; i++ {
			e := crates[from].Remove(crates[from].Front())
			crates[to].PushFront(e)
		}
	}
}
