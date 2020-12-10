package main

import (
  "bufio"
  "log"
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

  var rows []string
  for scanner.Scan() {
    rows = append(rows, scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  TraverseForest(3, 1, rows, false)
  TestSlopes(rows)
}

func TraverseForest(right int, down int, rows []string, quietMode bool) int {
  currentXPos := right
  currentYPos := down
  treesHit := 0

  for currentYPos < len(rows) {
    row := rows[currentYPos]
    if row != "" {
      if (!quietMode) {
        fmt.Printf("Position: (%d, %d) %v\n", currentXPos, currentYPos, row)
      }
      
      modPos := currentXPos % (len(row))
      charAtPos := string(row[modPos])
      if (!quietMode) {
        fmt.Println(charAtPos, "found at modified position", modPos)
      }

      if charAtPos == "#" {
        treesHit++
      }

      currentXPos += right
      currentYPos += down
    }
  }

  fmt.Println(chalk.Green, treesHit, "trees hit!")
  return treesHit
}

func TestSlopes(rows []string) int {
  first := TraverseForest(1, 1, rows, true)
  second := TraverseForest(3, 1, rows, true)
  third := TraverseForest(5, 1, rows, true)
  fourth := TraverseForest(7, 1, rows, true)
  fifth := TraverseForest(1, 2, rows, true)

  ans := first * second * third * fourth * fifth

  fmt.Println(chalk.Blue, ans)
  return ans
}
