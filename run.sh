#!/bin/sh

pwd = `pwd`

cd $1/day$2/
go build main.go && time ./main
rm main
cd $pwd