#!/bin/bash

number=$1

mkdir -p "day${1}/a"
mkdir -p "day${1}/b"
cd "day${1}"
echo "https://adventofcode.com/2022/day/${number}" > README.md
touch a/a.go
touch b/b.go
touch input.txt
touch testinput.txt
