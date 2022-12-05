#!/bin/bash

number=$1

mkdir "day${1}"
cd "day${1}"
echo "https://adventofcode.com/2022/day/${number}" > README.md
touch a.go
touch b.go
touch input.txt
touch testinput.txt
