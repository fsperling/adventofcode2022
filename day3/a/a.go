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
    var match = strings.Split(fileScanner.Text(), "")
    fmt.Printf("%v\n", match)
    var length = len(match)

    // put first half of array in a map
    firstHalf := make(map[string]int)
    for i := 0; i < length/2; i++ {
      firstHalf[match[i]] = 1
    }

    // check if items from second half are contained in first half
    for j := length/2; j < length; j++ {
      _, ok := firstHalf[match[j]]
      if ok {
        fmt.Printf("need to add: %s\n", match[j])
        fmt.Println([]rune(match[j])[0])

        // calculate their prio value
        if strings.ToUpper(match[j]) == match[j] {
          var prio = []rune(match[j])[0] - 38
          fmt.Printf("prio for: is %d\n", prio)
          totalPrio += int(prio)
        } else {
          var prio = []rune(match[j])[0] - 96
          fmt.Printf("prio for: is %d\n", prio)
          totalPrio += int(prio)
        }
        break
      }
    }
    
  }

  f.Close()
  fmt.Println(totalPrio)
}

