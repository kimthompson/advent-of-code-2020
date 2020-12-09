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

  //CountValidPasswordsMinCharacters(cleaned)
  CountValidPasswordsCharacterLocation(cleaned)
}

func CountValidPasswordsMinCharacters(passwords []Password) int {
  fmt.Println(chalk.White, "processing pt. 1 . . .")

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
    log.Panic(chalk.Red, "no valid passwords found")
  }

  fmt.Println(chalk.Green, countValid, "valid passwords found")
  return countValid
}

func CountValidPasswordsCharacterLocation(passwords []Password) int {
  fmt.Println(chalk.White, "processing pt. 2 . . .")

  countValid := 0

  for _, pwd := range passwords {
    fmt.Println(chalk.White, "Password", chalk.Magenta, pwd.Password, chalk.White, "requires", pwd.Character, "in either the", pwd.Min, "position or the", pwd.Max, "position and has a length of", len(pwd.Password))

    if pwd.Min - 1 >= 0 && pwd.Max - 1 < len(pwd.Password) {
      first := string(pwd.Password[pwd.Min - 1])
      second := string(pwd.Password[pwd.Max - 1])

      firstMatch := first == pwd.Character
      secondMatch := second == pwd.Character

      if ((firstMatch || secondMatch) && !(firstMatch && secondMatch)) {
        fmt.Println(chalk.Blue, "is good")
        countValid++
      }
      
    } else {
      log.Panic(chalk.Red, "index doesn't exist")
    }
    
  }

  if countValid <= 0 {
    log.Panic(chalk.Red, "no valid passwords found")
  }

  fmt.Println(chalk.Green, countValid, "valid passwords found")
  return countValid
}

// "Min" and "Max" don't really describe the second problem's data relationship, but it works fine. I want to use the same struct for both problems
type Password struct {
  Min int 
  Max int
  Character string
  Password string
}
