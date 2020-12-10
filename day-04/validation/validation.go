package validation

import (
  "regexp"
  "strconv"
)

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
