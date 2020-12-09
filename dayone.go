package main

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
)

func main () {
  file, err := ioutil.ReadFile("./data/dayone.json")

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
}

func FindSumToYearWithTwo(d []int ) int {
  for a := 0; a < len(d); a++ {
    fmt.Printf("trying %d with:\n", d[a])

    for b := 0; b < len(d); b++ {
      sum := d[a] + d[b]
      isYear := sum == 2020
      fmt.Printf("\t%d, which equals %d\n", d[b], sum)
      if isYear {
        product := d[a] * d[b]
        fmt.Printf("\nIt was %d and %d! ", d[a], d[b])
        fmt.Println("The product of these two is", product)
        return product
      }
    }
  }

  return 0
}
