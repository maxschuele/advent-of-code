#!/bin/sh

mkdir p $1/day$2
touch $1/day$2/input.txt

if [[ ! -e $1/day$2/main.go ]]; then
touch $1/day$2/main.go
echo "package main

import (
	_ \"embed\"
	\"fmt\"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(p1(input), p2(input))
}

func p1(input string) int {
	total := 0

	return total
}

func p2(input string) int {
	total := 0

	return total
}" >> $1/day$2/main.go
fi