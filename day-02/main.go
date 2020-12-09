package main

import (
  "fmt"
  "log"
  "strings"
  "strconv"
  "io/ioutil"

  "github.com/ttacon/chalk"
)

func main () {
  file, err := ioutil.ReadFile("data.txt")

  if err != nil {
    fmt.Println("File reading error", err)
    return
  }

  fileContents := string(file)
  arrayOfLines := strings.Split(fileContents, "\n")

  fmt.Println(chalk.White, "formatting data . . .")

  var cleaned []Password

  for a := 0; a < len(arrayOfLines); a++ {
    if arrayOfLines[a] != "" {
      attrs := strings.Split(arrayOfLines[a], " ")
      bounds := strings.Split(attrs[0], "-")
      min, err := strconv.Atoi(bounds[0])
      max, err := strconv.Atoi(bounds[1])

      if (err != nil) {
        fmt.Println(chalk.Red, "int parsing error", err)
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

  fmt.Println(chalk.Magenta, "data ready!")

  CountValidPasswords(cleaned)
}

func CountValidPasswords(passwords []Password) int {
  fmt.Println(chalk.White, "processing . . .")

  countValid := 0
  for _, pwd := range passwords {
    charCount := strings.Count(pwd.Password, pwd.Character)
    fmt.Println(chalk.White, "Password requires between", pwd.Min, "and", pwd.Max, "of the character", pwd.Character, "| found", charCount)
    if charCount >= pwd.Min && charCount <= pwd.Max {
      fmt.Println(chalk.Blue, "is good")
      countValid++
    }
  }

  if countValid <= 0 {
    log.Fatal(chalk.Red, "No valid passwords found")
  }

  fmt.Println(chalk.Green, countValid, "valid passwords found")
  return countValid
}

type Password struct {
  Min int 
  Max int
  Character string
  Password string
}
