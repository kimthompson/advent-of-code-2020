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

  dataCopy := data

  for a := 0; a < len(data); a++ {
    fmt.Printf("a at %d is %d\n", a, data[a])

    for b := 0; b < len(dataCopy); b++ {
      fmt.Printf("\tb at %d is %d\n", b, dataCopy[b])
    }
  }
}
