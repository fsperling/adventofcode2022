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
    
    // Draws
    if match[0] == "A" && match[1] == "X" ||
      match[0] == "B" && match[1] == "Y" ||
      match[0] == "C" && match[1] == "Z" {
      score += 3
    }
 
    // Wins
    if match[0] == "A" && match[1] == "Y" ||
      match[0] == "B" && match[1] == "Z" ||
      match[0] == "C" && match[1] == "X" {
      score += 6
    }

    switch match[1] {
    case "X":
      score += 1
    case "Y":
      score += 2
    case "Z":
      score += 3
    }
    totalScore = totalScore + score
  }

  f.Close()
  fmt.Println(totalScore)
}

