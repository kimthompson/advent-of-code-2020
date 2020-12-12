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
  groupsWithExclusiveTotals := GetExclusiveTotalsPerGroup(groups)

  total := 0
  for _, group := range groupsWithTotals {
    total = total + group.total
  }

  exclusiveTotal := 0
  for _, group := range groupsWithExclusiveTotals {
    exclusiveTotal = exclusiveTotal + group.total
  }

  fmt.Println(chalk.Green, "Part 1 Total:", total)
  fmt.Println(chalk.Green, "Part 2 Total:", exclusiveTotal)
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
func GetTotalsPerGroup(groups [][]string) []Group {
  var groupTotals []Group

  for _, group := range groups {
    answers := strings.Join(group, " ")

    var groupTotal Group
    groupTotal.answers = group
    groupTotal.answerString = answers
    groupTotal.total = GetGroupTotal(answers)
    groupTotals = append(groupTotals, groupTotal)
  }

  return groupTotals
}

func GetExclusiveTotalsPerGroup(groups [][]string) []Group {
  var groupTotals []Group

  for _, group := range groups {
    answers := strings.Join(group, " ")

    var groupTotal Group
    groupTotal.answers = group
    groupTotal.answerString = answers
    groupTotal.total = GetExclusiveGroupTotal(group)
    groupTotals = append(groupTotals, groupTotal)
  }

  return groupTotals
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

func GetExclusiveGroupTotal(answers []string) int {
  answersString := strings.Join(answers, " ")
  unique := map[string]string{}

  for _, char := range answersString {
    re := regexp.MustCompile(`^[a-z]$`)
    isValidChar := re.MatchString(string(char))
    if isValidChar {
      unique[string(char)] = string(char)
    }
  }

  groupSize := len(answers)
  total := 0
  for _, char := range unique {
    answerCount := strings.Count(answersString, string(char))

    if answerCount == groupSize {
      total = total + 1
    }
  }

  return total
}

type Group struct {
  answers []string
  answerString string
  total int
}
