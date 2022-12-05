package main

import "fmt"
import "os"
import "sort"
import "strconv"
import "bufio"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func trackNewElv(currentValue *int, allValues []int) {
  if *currentValue > allValues[0] {
   allValues[0] = *currentValue
   sort.Sort(sort.IntSlice(allValues))
  }
  *currentValue = 0
}


func main() {
  f, err := os.Open("input.txt")
  check(err)
  var fileScanner = bufio.NewScanner(f)
  fileScanner.Split(bufio.ScanLines)

  var mostCalories = []int {0, 0, 0}
  var currentElfCalories int = 0

  for fileScanner.Scan() {
    var calories = fileScanner.Text()
    
    if calories == "" {
      trackNewElv(&currentElfCalories, mostCalories)
     } else {
      intVar, _ := strconv.Atoi(calories)
      currentElfCalories = currentElfCalories + intVar
    }
  }

  trackNewElv(&currentElfCalories, mostCalories)

  f.Close()

  fmt.Printf("%v\n", mostCalories)
  fmt.Println(mostCalories[0] + mostCalories[1] +  mostCalories[2])
}

