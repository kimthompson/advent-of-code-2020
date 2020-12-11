package main

import (
  "bufio"
  "log"
  "fmt"
  "regexp"
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
  
  groupsWithTotals := GetTotalsPerGroup(groups)
  total := 0

  for _, group := range groupsWithTotals {
    fmt.Println(chalk.Magenta, group)
    total = total + group.total
  }

  fmt.Println(chalk.Green, "Total:", total)
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

func GetGroupTotal(answers string) int {
  unique := map[string]bool{}

  for _, char := range answers {
    re := regexp.MustCompile(`^[a-z]$`)
    isValidChar := re.MatchString(string(char))
    if isValidChar {
      unique[string(char)] = true
    }
  }

  return len(unique)
}

func GetTotalsPerGroup(groups [][]string) []Group {
  var groupTotals []Group

  for _, group := range groups {
    answers := strings.Join(group, " ")

    var groupTotal Group
    groupTotal.answerString = answers
    groupTotal.total = GetGroupTotal(answers)
    groupTotals = append(groupTotals, groupTotal)
  }

  return groupTotals
}

type Group struct {
  answerString string
  total int
}
