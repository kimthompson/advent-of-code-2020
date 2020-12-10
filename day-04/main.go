package main

import (
  "bufio"
  "log"
  "fmt"
  "regexp"
  "strings"
  "strconv"
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

  passports := ParseDataForPassports(lines)

  PrettyPrintPassports(passports)

  validPassports := ValidatePassportsSimple(passports)
  newValidPassports := ValidatePassportsComplex(passports)

  fmt.Println(chalk.Green, len(validPassports), "out of", len(passports), "are valid at first glance.")
  fmt.Println(chalk.Green, len(newValidPassports), "out of", len(passports), "are truly valid.")
}

func ParseDataForPassports(lines []string) []Passport {
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

func ValidatePassportsSimple(passports []Passport) []Passport {
  var result []Passport
  for _, pspt := range passports {
    if (
      pspt.BYR != "" &&
      pspt.IYR != "" &&
      pspt.EYR != "" &&
      pspt.HGT != "" &&
      pspt.HCL != "" &&
      pspt.ECL != "" &&
      pspt.PID != "") {
      result = append(result, pspt)
    }
  }
  return result
}

func ValidatePassportsComplex(passports []Passport) []Passport {
  var result []Passport

  for _, pspt := range passports {
    if (
      ValidateBYR(pspt.BYR) &&
      ValidateIYR(pspt.IYR) &&
      ValidateEYR(pspt.EYR) &&
      ValidateHGT(pspt.HGT) &&
      pspt.HCL != "" &&
      pspt.ECL != "" &&
      pspt.PID != "") {
      result = append(result, pspt)
    }
  }
  return result
}

func ValidateBYR(attr string) bool {
  re := regexp.MustCompile(`\d\d\d\d`)
  isFourDigits := re.MatchString(attr)
  year := 0

  if isFourDigits {
    parsedYear, _ := strconv.Atoi(attr)
    year = parsedYear
  }

  return isFourDigits && year >= 1920 && year <= 2002
}

func ValidateIYR(attr string) bool {
  re := regexp.MustCompile(`\d\d\d\d`)
  isFourDigits := re.MatchString(attr)
  year := 0

  if isFourDigits {
    parsedYear, _ := strconv.Atoi(attr)
    year = parsedYear
  }

  return isFourDigits && year >= 2010 && year <= 2020
}

func ValidateEYR(attr string) bool {
  re := regexp.MustCompile(`\d\d\d\d`)
  isFourDigits := re.MatchString(attr)
  year := 0

  if isFourDigits {
    parsedYear, _ := strconv.Atoi(attr)
    year = parsedYear
  }

  return isFourDigits && year >= 2020 && year <= 2030
}

func ValidateHGT(attr string) bool {
  unitRegex := regexp.MustCompile(`(cm|in)`)
  unitsFound := unitRegex.FindAllString(attr, -1)
  unit := ""

  if len(unitsFound) == 1 {
    unit = unitsFound[0]
  }

  heightRegex := regexp.MustCompile(`\d*`)
  heightString := heightRegex.FindString(attr)
  height, _ := strconv.Atoi(heightString)

  isValidHeight := false

  if unit == "cm" {
    isValidHeight = height >= 150 && height <= 193
  } else if unit == "in" {
    isValidHeight = height >= 59 && height <= 76
  }

  return isValidHeight
}

func PrettyPrintPassports(passports []Passport) {
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

type Passport struct {
  BYR string
  IYR string
  EYR string
  HGT string
  HCL string
  ECL string
  PID string
  CID string
}
