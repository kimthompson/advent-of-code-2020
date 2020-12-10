package main

import (
  "bufio"
  "log"
  "fmt"
  //"strings"
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

  for _, line := range lines {
    fmt.Println(chalk.Magenta, line)
  }
}
