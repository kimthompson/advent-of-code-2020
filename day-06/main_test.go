package main

import (
  "testing"
)

func TestGetGroupTotal(t *testing.T) {
  input := "lfhtzrxcj fbtzlhrcj lzdaftrjphco w"
  want := 15
  result := GetGroupTotal(input)
  if want != result {
      t.Fatalf(`ValidateBYR("%v") = %d, wanted %d`, input, result, want)
  }

  input = "rsdywfzikhnlqc kifcrndqlswzhy hdfkiylsnqcrzw"
  want = 14
  result = GetGroupTotal(input)
  if want != result {
      t.Fatalf(`ValidateBYR("%v") = %d, wanted %d`, input, result, want)
  }

  input = "jfkdls jdklwudkalb dlsw a abc"
  want = 11
  result = GetGroupTotal(input)
  if want != result {
      t.Fatalf(`ValidateBYR("%v") = %d, wanted %d`, input, result, want)
  }
}
