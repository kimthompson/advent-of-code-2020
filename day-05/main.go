package main

import (
  "bufio"
  "log"
  "fmt"
  "os"

  "github.com/ttacon/chalk"
)

const ROW_MAX = 127
const COL_MAX = 7

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

  highestSeatId := 0

  for _, line := range lines {
    fmt.Println(chalk.Magenta, line)

    rowNum := FindRow(line[:7])
    colNum := FindCol(line[7:])
    seatId := (rowNum * 8) + colNum

    fmt.Println(chalk.Blue, "[", rowNum, ",", colNum, "]", seatId)

    if seatId > highestSeatId {
      highestSeatId = seatId
    }
  }

  fmt.Println(chalk.Green, "The highest seat ID is", highestSeatId)
}

func FindRow(code string) int {
  // I found it easier to wrap my head around physically representing each row and looping through it,
  // so I generated a slice with a number representing each row
  var planeRows []int
  for r := 0; r <= ROW_MAX; r++ {
    planeRows = append(planeRows, r)
  }

  for _, char := range code {
    if string(char) == "F" {
      planeRows = planeRows[:len(planeRows) / 2]
    } else if string(char) == "B" {
      planeRows = planeRows[len(planeRows) / 2:]
    }
  }
  return planeRows[0]
}

func FindCol(code string) int {
  var planeCols []int
  for c := 0; c <= COL_MAX; c++ {
    planeCols = append(planeCols, c)
  }

  for _, char := range code {
    if string(char) == "L" {
      planeCols = planeCols[:len(planeCols) / 2]
    } else if string(char) == "R" {
      planeCols = planeCols[len(planeCols) / 2:]
    }
  }
  return planeCols[0]
}
