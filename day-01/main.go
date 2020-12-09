package main

import (
  "fmt"
  "log"
  "io/ioutil"
  "encoding/json"
)

func main () {
  file, err := ioutil.ReadFile("data.json")

  if err != nil {
    fmt.Println("File reading error", err)
    return
  }

  jsonString := string(file)
  fmt.Println("Contents of file:", jsonString)

  var data []int
  json.Unmarshal([]byte(jsonString), &data)
  fmt.Println("to array", data)

  FindSumToYearWithTwo(data)
  FindSumToYearWithThree(data)
}

func FindSumToYearWithTwo(d []int) (int) {
  for a := 0; a < len(d); a++ {
    fmt.Printf("trying %d with:\n", d[a])
    for b := 0; b < len(d); b++ {
      fmt.Printf("\t\t%d...\n", d[b])
      if a != b {
        sum := d[a] + d[b]
        isYear := sum == 2020
        if isYear {
          product := d[a] * d[b]
          fmt.Printf("\nIt was %d and %d! ", d[a], d[b])
          fmt.Printf("The product of these two is %d.\n", product)
          return product
        }
      }
    }
  }

  log.Panic("No matching sum found")
  return 0
}

func FindSumToYearWithThree(d []int) (int) {
  for a := 0; a < len(d); a++ {
    fmt.Printf("trying %d with:\n", d[a])
    for b := 0; b < len(d); b++ {
      fmt.Printf("\ttrying %d with:\n", d[b])
      for c := 0; c < len(d); c++ {
        if a != b || b != c || c != a {
          fmt.Printf("\t\t%d...\n", d[c])
          sum := d[a] + d[b] + d[c]
          isYear := sum == 2020
          if isYear {
            product := d[a] * d[b] * d[c]
            fmt.Printf("\nIt was %d, %d, and %d! ", d[a], d[b], d[c])
            fmt.Printf("The product of these three is %d.\n", product)
            return product
          }
        }
      }
    }
  }

  log.Panic("No matching sum found")
  return 0
}
