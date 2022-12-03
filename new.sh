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
	fmt.Println(p1(input), p2(input))
}

func p1(input string) int {
	total := 0

	return total
}

func p2(input string) int {
	total := 0

	return total
}" >> $path/main.go
fi