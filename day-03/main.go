package main

import (
  "errors"
  "fmt"
  "log"
  "strings"
  "io/ioutil"

  "github.com/ttacon/chalk"
)

func main () {
  file, err := ioutil.ReadFile("data.txt")

  if err != nil {
    log.Fatal(chalk.Red, "index doesn't exist")
    return
  }

  fileContents := string(file)
  rows := strings.Split(fileContents, "\n")

  fmt.Println(chalk.Magenta, "data ready!")

  TraverseForest(3, 1, rows[0:26])
}

func TraverseForest(right int, down int, rows []string) (int, error) {
  if (right < 1 || down < 1) {
    return 0, errors.New("need positive, non-zero inputs")
  }

  fmt.Println(chalk.White, "processing pt. 1 . . .")

  currentXPos := right
  currentYPos := down
  treesHit := 0

  for currentYPos < len(rows) {
    row := rows[currentYPos]
    if row != "" {
      fmt.Printf("Position: (%d, %d) %v\n", currentXPos, currentYPos, row)
      
      charAtPos := ""
      if currentXPos > len(row) {
        modPos := currentXPos % (len(row))
        charAtPos = string(row[modPos])
        fmt.Println("modPos", modPos, charAtPos)
      } else {
        charAtPos = string(row[currentXPos])
        fmt.Println("pos", currentXPos, charAtPos)
      }

      currentXPos += right
      currentYPos += down

      treesHit++
    }
  }

  fmt.Println(chalk.Green, treesHit, "trees hit!")
  return treesHit, nil
}
