package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func main () {
  file, err := ioutil.ReadFile("data.txt")

  if err != nil {
    fmt.Println("File reading error", err)
    return
  }

  fileContents := string(file)
  arrayOfLines := strings.Split(fileContents, "\n")

  fmt.Println("formatting data . . .")

  var cleaned []Password

  for a := 0; a < len(arrayOfLines); a++ {
    if arrayOfLines[a] != "" {
      attrs := strings.Split(arrayOfLines[a], " ")
      bounds := strings.Split(attrs[0], "-")
      min, err := strconv.Atoi(bounds[0])
      max, err := strconv.Atoi(bounds[1])

      if (err != nil) {
        fmt.Println("int parsing error", err)
        min = 0
        max = 0
      }

      password := Password{
        Min: min,
        Max: max,
        Character: strings.TrimSuffix(attrs[1], ":"),
        Password: attrs[2],
      }
      cleaned = append(cleaned, password)
    }
  }

  fmt.Println("data ready!")

  CountValidPasswords(cleaned)
}

func CountValidPasswords(data []Password) int {
  fmt.Println("processing . . .")
  return 0
}

type Password struct {
  Min int 
  Max int
  Character string
  Password string
}
