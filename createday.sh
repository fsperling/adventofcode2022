#!/bin/bash

number=$1

mkdir -p "day${1}"
cd "day${1}"
echo "https://adventofcode.com/2022/day/${number}" > README.md
touch a.py
touch b.py
touch input.txt
touch testinput.txt
