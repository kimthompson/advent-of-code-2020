package main

import (
  "bufio"
  "log"
  "errors"
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

  fmt.Println(len(rows))

  TraverseForest(3, 1, rows)
}

func TraverseForest(right int, down int, rows []string) (int, error) {
  if (right < 1 || down < 1) {
    return 0, errors.New("need positive, non-zero inputs")
  }

  fmt.Println(chalk.White, "processing pt. 1 . . .")

  currentXPos := 1 
  currentYPos := 1
  treesHit := 0

  for currentYPos < len(rows) {
    row := rows[currentYPos]
    if row != "" {
      fmt.Printf("Position: (%d, %d) %v\n", currentXPos, currentYPos, row)
      
      modPos := currentXPos % (len(row))
      charAtPos := string(row[modPos])
      fmt.Println(charAtPos, "found at modified position", modPos)

      if charAtPos == "#" {
        treesHit++
      }

      currentXPos += right
      currentYPos += down
    }
  }

  fmt.Println(chalk.Green, treesHit, "trees hit!")
  return treesHit, nil
}
