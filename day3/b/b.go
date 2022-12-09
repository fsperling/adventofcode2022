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

  var totalPrio = 0 

  for fileScanner.Scan() {
    var backpack1 = strings.Split(fileScanner.Text(), "")
    fileScanner.Scan()
    var backpack2 = strings.Split(fileScanner.Text(), "")
    fileScanner.Scan()
    var backpack3 = strings.Split(fileScanner.Text(), "")

    // put first two arrays in a map each
    first := make(map[string]int)
    for i := 0; i < len(backpack1); i++ {
      first[backpack1[i]] = 1
    }

    second := make(map[string]int)
    for i := 0; i < len(backpack2); i++ {
      second[backpack2[i]] = 1
    }

    // check if items from the 3rd are contained in 1 and 2
    for j := 0; j < len(backpack3); j++ {
      _, ok1 := first[backpack3[j]]
      _, ok2 := second[backpack3[j]]

      if ok1 && ok2 {
        fmt.Printf("need to add: %s\n", backpack3[j])
        fmt.Println([]rune(backpack3[j])[0])

        // calculate their prio value
        if strings.ToUpper(backpack3[j]) == backpack3[j] {
          var prio = []rune(backpack3[j])[0] - 38
          totalPrio += int(prio)
        } else {
          var prio = []rune(backpack3[j])[0] - 96
          totalPrio += int(prio)
        }
        break
      }
    }
    
  }

  f.Close()
  fmt.Println(totalPrio)
}

