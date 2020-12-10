package main

import (
  "bufio"
  "log"
  "fmt"
  "strings"
  "os"

  "github.com/ttacon/chalk"
)

const BYR = "byr"
const IYR = "iyr"
const EYR = "eyr"
const HGT = "hgt"
const HCL = "hcl"
const ECL = "ecl"
const PID = "pid"
const CID = "cid"

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

  passports := ParseData(lines)

  var validPassports []Passport 
  for _, pspt := range passports {
    if pspt.Valid {
      validPassports = append(validPassports, pspt)
    }
  }

  fmt.Println(chalk.Green, len(validPassports), "out of", len(passports), "are valid.")
}

func ParseData(lines []string) []Passport {
  var bundles []string

  // get each set of attributes onto one line
  start := 0
  for i, line := range lines {
    if line == "" {
      ans := strings.Join(lines[start:i], " ")
      start = i + 1
      bundles = append(bundles, ans)
    }
  }

  var pspts []Passport

  // examine the strings to fill in my Passport struct (if able)
  for _, bundle := range bundles {
    var pspt Passport

    pspt.BYR = LookForAttr(bundle, BYR)
    pspt.IYR = LookForAttr(bundle, IYR)
    pspt.EYR = LookForAttr(bundle, EYR)
    pspt.HGT = LookForAttr(bundle, HGT)
    pspt.HCL = LookForAttr(bundle, HCL)
    pspt.ECL = LookForAttr(bundle, ECL)
    pspt.PID = LookForAttr(bundle, PID)
    pspt.CID = LookForAttr(bundle, CID)
    pspt.Valid = pspt.BYR != "" &&
      pspt.IYR != "" &&
      pspt.EYR != "" &&
      pspt.HGT != "" &&
      pspt.HCL != "" &&
      pspt.ECL != "" &&
      pspt.PID != ""

    fmt.Println(chalk.Blue, pspt.Valid)

    pspts = append(pspts, pspt)
  }

  return pspts
}

func LookForAttr(psptString string, attr string) string {
  loc := strings.Index(psptString, attr)

  if loc >= 0 {
    temp := psptString[loc:]
    value := strings.Split(strings.Split(temp, ":")[1], " ")[0]
    return value 
  }

  return ""
}

type Passport struct {
  BYR string
  IYR string
  EYR string
  HGT string
  HCL string
  ECL string
  PID string
  CID string
  Valid bool
}
