#!/bin/sh

path="$1/day$2"

mkdir $path
touch $path/input.txt

if [[ ! -e $path/main.go ]]; then
touch $path/main.go
echo "package main

import (
	_ \"embed\"
	\"fmt\"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, \"\n\")
	fmt.Println(part1(input), part2(input))
}

func part1(lines []string) int {
	total := 0

	for _, line := range lines {

	}

	return total
}

func part2(lines []string) int {
	total := 0

	for _, line := range lines {
		
	}

	return total
}" >> $path/main.go
fi