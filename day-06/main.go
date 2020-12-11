package main

import (
  "bufio"
  "log"
  "fmt"
  "strings"
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

  groups := GetGroups(lines)
  
  fmt.Println(chalk.Blue, GetTotalsPerGroup(groups))
}

func GetGroups(lines []string) [][]string {
  var groups [][]string

  start := 0
  for i, line := range lines {
    if line == "" {
      group := lines[start:i]
      start = i + 1
      groups = append(groups, group)
    }
  }

  return groups
}

func GetTotalsPerGroup(groups [][]string) map[string]int {
  groupTotals := map[string]int{}

  for _, group := range groups {
    answers := strings.Join(group, " ")
    
    unique := map[string]bool{}

    for _, char := range answers {
      if string(char) != " " {
        unique[string(char)] = true
      }
    }

    groupTotals[answers] = len(unique)
  }

  return groupTotals
}
