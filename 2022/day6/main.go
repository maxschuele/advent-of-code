package main

import (
	"fmt"
	"io"
	"log"
	"math/bits"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	input, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(uniqueChars(input, 4), uniqueChars(input, 14))
}

func uniqueChars(input []byte, want int) int {
	var b, x uint32
	for i := want - 1; i < len(input); i++ {
		for _, c := range input[i-(want-1) : i+1] {
			x = 1 << (c - 97)
			if x&b > 0 {
				break
			}
			b = b | x
		}

		if bits.OnesCount32(b) == want {
			return i + 1
		}
		b = 0
	}

	return -1
}
