#!/bin/bash

number=$1

mkdir "day${1}/{a,b}"
cd "day${1}"
echo "https://adventofcode.com/2022/day/${number}" > README.md
touch a/a.go
touch b/b.go
touch input.txt
touch testinput.txt
