package main

import "fmt"
import "os"
import "strconv"
import "strings"
import "bufio"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func Min(x int, y int) int {
  if x <  y {
    return x
  }
  return y
}

func Max(x int, y int) int {
  if x > y {
    return x
  }
  return y
}


func main() {
  f, err := os.Open("../input.txt")
  check(err)
  var fileScanner = bufio.NewScanner(f)
  fileScanner.Split(bufio.ScanLines)

  var totalPairs = 0 

  for fileScanner.Scan() {
    var pairs = strings.Split(fileScanner.Text(), ",")
    var one = strings.Split(pairs[0], "-")
    var two = strings.Split(pairs[1], "-")

    start_one , _ := strconv.Atoi(one[0])
    end_one, _ := strconv.Atoi(one[1])
    start_two, _ := strconv.Atoi(two[0])
    end_two, _ := strconv.Atoi(two[1])

    var calc = Max(start_one, start_two) - Min(end_one, end_two)
    if calc <= 0 {
      totalPairs += 1
    }
    
  }

  f.Close()
  fmt.Println(totalPairs)
}

