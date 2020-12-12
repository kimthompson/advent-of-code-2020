package main

import (
  "bufio"
  "log"
  "regexp"
  "strings"
  "strconv"
  "fmt"
  "os"

  "github.com/ttacon/chalk"
)

func main () {
  file, err := os.Open("data.txt")

  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  var lines []string
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  for _, ruleString := range lines {
    fmt.Println(chalk.Blue, ruleString)
    words := strings.Split(ruleString, " ")
    bagName := strings.Join(words[:2], " ")
    fmt.Println(chalk.Red, "bagName:", bagName)
    for a, word := range words {
      numberRegex := regexp.MustCompile(`^\d{1,}$`)
      isNumber := numberRegex.MatchString(word)
      var rule Rule
      if isNumber {
        num, _ := strconv.Atoi(word)
        rule.amount = num
        rule.modifier = words[a+1]
        rule.color = words[a+2]

        fmt.Println(chalk.Magenta, rule)
      }
    }
  }
}

type Rule struct {
  amount int
  modifier string
  color string
}
