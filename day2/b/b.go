package main

import "fmt"
import "os"
import "strings"
import "bufio"

func check(e error) {
  if e != nil {
      panic(e)
  }
}

func main() {
  f, err := os.Open("../input.txt")
  check(err)
  var fileScanner = bufio.NewScanner(f)
  fileScanner.Split(bufio.ScanLines)

  var totalScore = 0 

  for fileScanner.Scan() {
    var score = 0
    var match = strings.Split(fileScanner.Text(), " ")
    
    // Rock
    if match[0] == "A" && match[1] == "Y" ||
      match[0] == "B" && match[1] == "X" ||
      match[0] == "C" && match[1] == "Z" {
      score += 1
    }
 
    // Paper
    if match[0] == "A" && match[1] == "Z" ||
      match[0] == "B" && match[1] == "Y" ||
      match[0] == "C" && match[1] == "X" {
      score += 2
    }

    // Scissor
    if match[0] == "A" && match[1] == "X" ||
      match[0] == "B" && match[1] == "Z" ||
      match[0] == "C" && match[1] == "Y" {
      score += 3
    }
    
    switch match[1] {
    case "X":
      score += 0
    case "Y":
      score += 3
    case "Z":
      score += 6
    }
    totalScore = totalScore + score
  }


  f.Close()
  fmt.Println(totalScore)
}

