package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/edwingeng/deque/v2"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	crates, crates2 := parseCrates(scanner)
	moves := parseMoves(scanner)

	for _, m := range moves {
		for _, e := range crates[m[1]].DequeueMany(m[0]) {
			crates[m[2]].PushFront(e)
		}

		x := crates2[m[1]].DequeueMany(m[0])
		for i := len(x) - 1; i >= 0; i-- {
			crates2[m[2]].PushFront(x[i])
		}
	}

	var p1, p2 []byte
	for i := 1; i <= len(crates); i++ {
		f1, _ := crates[i].Front()
		f2, _ := crates2[i].Front()
		p1 = append(p1, f1)
		p2 = append(p2, f2)
	}

	fmt.Println(string(p1))
	fmt.Println(string(p2))

}

func parseCrates(s *bufio.Scanner) (map[int]*deque.Deque[byte], map[int]*deque.Deque[byte]) {
	crates := make(map[int]*deque.Deque[byte])
	crates2 := make(map[int]*deque.Deque[byte])

	for s.Scan() {
		line := s.Text()

		if line[1] == '1' {
			break
		}

		crateIndex := 1
		for i := 1; i < len(line); i += 4 {
			if line[i] != ' ' {
				if crates[crateIndex] == nil {
					crates[crateIndex] = deque.NewDeque[byte]()
					crates2[crateIndex] = deque.NewDeque[byte]()
				}

				crates[crateIndex].PushBack(line[i])
				crates2[crateIndex].PushBack(line[i])
			}

			crateIndex++
		}
	}

	s.Scan() //get rid of the empty line
	return crates, crates2
}

func parseMoves(s *bufio.Scanner) [][3]int {
	var moves [][3]int
	for s.Scan() {
		fields := strings.Fields(s.Text())
		n, _ := strconv.Atoi(fields[1])
		f, _ := strconv.Atoi(fields[3])
		t, _ := strconv.Atoi(fields[5])
		moves = append(moves, [3]int{n, f, t})
	}

	return moves
}
