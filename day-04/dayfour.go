package main

import (
  "bufio"
  "log"
  "fmt"
  "strings"
  "os"

  "kimthompson.me/validation"
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

  passports := ParseDataForPassports(lines)

  PrettyPrintPassports(passports)

  validPassports := validation.ValidatePassportsSimple(passports)
  newValidPassports := validation.ValidatePassportsComplex(passports)

  fmt.Println(chalk.Green, len(validPassports), "out of", len(passports), "are valid at first glance.")
  fmt.Println(chalk.Green, len(newValidPassports), "out of", len(passports), "are truly valid.")
}

func ParseDataForPassports(lines []string) []validation.Passport {
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

  var pspts []validation.Passport

  // examine the strings to fill in my Passport struct (if able)
  for _, bundle := range bundles {
    var pspt validation.Passport

    pspt.BYR = LookForAttr(bundle, BYR)
    pspt.IYR = LookForAttr(bundle, IYR)
    pspt.EYR = LookForAttr(bundle, EYR)
    pspt.HGT = LookForAttr(bundle, HGT)
    pspt.HCL = LookForAttr(bundle, HCL)
    pspt.ECL = LookForAttr(bundle, ECL)
    pspt.PID = LookForAttr(bundle, PID)
    pspt.CID = LookForAttr(bundle, CID)

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

func PrettyPrintPassports(passports []validation.Passport) {
  for i, pspt := range passports {
    fmt.Println(chalk.Blue, "Passport", i)
    fmt.Println(chalk.Red, "==========")
    fmt.Println(chalk.Magenta, BYR, chalk.White, pspt.BYR)
    fmt.Println(chalk.Magenta, IYR, chalk.White, pspt.IYR)
    fmt.Println(chalk.Magenta, EYR, chalk.White, pspt.EYR)
    fmt.Println(chalk.Magenta, HGT, chalk.White, pspt.HGT)
    fmt.Println(chalk.Magenta, HCL, chalk.White, pspt.HCL)
    fmt.Println(chalk.Magenta, ECL, chalk.White, pspt.ECL)
    fmt.Println(chalk.Magenta, PID, chalk.White, pspt.PID)
    fmt.Println(chalk.Magenta, CID, chalk.White, pspt.CID)
    fmt.Println(chalk.Red, "==========\n")
  }
}
