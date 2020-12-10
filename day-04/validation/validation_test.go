package validation

import (
  "testing"
)

func TestValidateBYR(t *testing.T) {
  input1 := "1992"
  want1 := true
  result1 := ValidateBYR(input1)
  if want1 != result1 {
      t.Fatalf(`ValidateBYR("%v") = %t, wanted %t`, input1, result1, want1)
  }

  input2 := "1910"
  want2 := false 
  result2 := ValidateBYR(input2)
  if want2 != result2 {
      t.Fatalf(`ValidateBYR("%v") = %t, wanted %t`, input2, result2, want2)
  }

  input3 := "2020"
  want3 := false 
  result3 := ValidateBYR(input3)
  if want3 != result3 {
      t.Fatalf(`ValidateBYR("%v") = %t, wanted %t`, input3, result3, want3)
  }

  input4 := ""
  want4 := false 
  result4 := ValidateBYR(input4)
  if want4 != result4 {
      t.Fatalf(`ValidateBYR("%v") = %t, wanted %t`, input4, result4, want4)
  }

  input5 := "eeeEEee"
  want5 := false 
  result5 := ValidateBYR(input5)
  if want5 != result5 {
      t.Fatalf(`ValidateBYR("%v") = %t, wanted %t`, input5, result5, want5)
  }
}

func TestValidateIYR(t *testing.T) {
  input1 := "2011"
  want1 := true
  result1 := ValidateIYR(input1)
  if want1 != result1 {
      t.Fatalf(`ValidateIYR("%v") = %t, wanted %t`, input1, result1, want1)
  }

  input2 := "1910"
  want2 := false 
  result2 := ValidateIYR(input2)
  if want2 != result2 {
      t.Fatalf(`ValidateIYR("%v") = %t, wanted %t`, input2, result2, want2)
  }

  input3 := "2024"
  want3 := false 
  result3 := ValidateIYR(input3)
  if want3 != result3 {
      t.Fatalf(`ValidateIYR("%v") = %t, wanted %t`, input3, result3, want3)
  }

  input4 := ""
  want4 := false 
  result4 := ValidateIYR(input4)
  if want4 != result4 {
      t.Fatalf(`ValidateIYR("%v") = %t, wanted %t`, input4, result4, want4)
  }

  input5 := "eeeEEee"
  want5 := false 
  result5 := ValidateIYR(input5)
  if want5 != result5 {
      t.Fatalf(`ValidateIYR("%v") = %t, wanted %t`, input5, result5, want5)
  }
}

func TestValidateEYR(t *testing.T) {
  input1 := "2021"
  want1 := true
  result1 := ValidateEYR(input1)
  if want1 != result1 {
      t.Fatalf(`ValidateEYR("%v") = %t, wanted %t`, input1, result1, want1)
  }

  input2 := "1910"
  want2 := false 
  result2 := ValidateEYR(input2)
  if want2 != result2 {
      t.Fatalf(`ValidateEYR("%v") = %t, wanted %t`, input2, result2, want2)
  }

  input3 := "2011"
  want3 := false 
  result3 := ValidateEYR(input3)
  if want3 != result3 {
      t.Fatalf(`ValidateEYR("%v") = %t, wanted %t`, input3, result3, want3)
  }

  input4 := ""
  want4 := false 
  result4 := ValidateEYR(input4)
  if want4 != result4 {
      t.Fatalf(`ValidateEYR("%v") = %t, wanted %t`, input4, result4, want4)
  }

  input5 := "eeeEEee"
  want5 := false 
  result5 := ValidateEYR(input5)
  if want5 != result5 {
      t.Fatalf(`ValidateEYR("%v") = %t, wanted %t`, input5, result5, want5)
  }
}

func TestValidateHGT(t *testing.T) {
  input1 := "154cm"
  want1 := true
  result1 := ValidateHGT(input1)
  if want1 != result1 {
      t.Fatalf(`ValidateHGT("%v") = %t, wanted %t`, input1, result1, want1)
  }

  input2 := "61in"
  want2 := true
  result2 := ValidateHGT(input2)
  if want2 != result2 {
      t.Fatalf(`ValidateHGT("%v") = %t, wanted %t`, input2, result2, want2)
  }

  input3 := "100cm"
  want3 := false 
  result3 := ValidateHGT(input3)
  if want3 != result3 {
      t.Fatalf(`ValidateHGT("%v") = %t, wanted %t`, input3, result3, want3)
  }

  input4 := "200in"
  want4 := false 
  result4 := ValidateHGT(input4)
  if want4 != result4 {
      t.Fatalf(`ValidateHGT("%v") = %t, wanted %t`, input4, result4, want4)
  }

  input5 := "eeeEEee"
  want5 := false 
  result5 := ValidateHGT(input5)
  if want5 != result5 {
      t.Fatalf(`ValidateHGT("%v") = %t, wanted %t`, input5, result5, want5)
  }

  input6 := "61"
  want6 := false 
  result6 := ValidateHGT(input6)
  if want6 != result6 {
      t.Fatalf(`ValidateHGT("%v") = %t, wanted %t`, input6, result6, want6)
  }
}

func TestValidateHCL(t *testing.T) {
  input1 := "#0200a0"
  want1 := true
  result1 := ValidateHCL(input1)
  if want1 != result1 {
      t.Fatalf(`ValidateHCL("%v") = %t, wanted %t`, input1, result1, want1)
  }

  input2 := "fff00d"
  want2 := false
  result2 := ValidateHCL(input2)
  if want2 != result2 {
      t.Fatalf(`ValidateHCL("%v") = %t, wanted %t`, input2, result2, want2)
  }

  input3 := "#xd1337"
  want3 := false 
  result3 := ValidateHCL(input3)
  if want3 != result3 {
      t.Fatalf(`ValidateHCL("%v") = %t, wanted %t`, input3, result3, want3)
  }

  input4 := "#fff"
  want4 := false 
  result4 := ValidateHCL(input4)
  if want4 != result4 {
      t.Fatalf(`ValidateHCL("%v") = %t, wanted %t`, input4, result4, want4)
  }

  input5 := "#eeEEee"
  want5 := false 
  result5 := ValidateHCL(input5)
  if want5 != result5 {
      t.Fatalf(`ValidateHCL("%v") = %t, wanted %t`, input5, result5, want5)
  }

  input6 := "#eeeeeee"
  want6 := false 
  result6 := ValidateHCL(input6)
  if want6 != result6 {
      t.Fatalf(`ValidateHCL("%v") = %t, wanted %t`, input6, result6, want6)
  }
}

func TestValidateECL(t *testing.T) {
  input1 := "amb"
  want1 := true
  result1 := ValidateECL(input1)
  if want1 != result1 {
      t.Fatalf(`ValidateECL("%v") = %t, wanted %t`, input1, result1, want1)
  }

  input2 := "blu"
  want2 := true 
  result2 := ValidateECL(input2)
  if want2 != result2 {
      t.Fatalf(`ValidateECL("%v") = %t, wanted %t`, input2, result2, want2)
  }

  input3 := "brn"
  want3 := true
  result3 := ValidateECL(input3)
  if want3 != result3 {
      t.Fatalf(`ValidateECL("%v") = %t, wanted %t`, input3, result3, want3)
  }

  input4 := "gry"
  want4 := true
  result4 := ValidateECL(input4)
  if want4 != result4 {
      t.Fatalf(`ValidateECL("%v") = %t, wanted %t`, input4, result4, want4)
  }

  input5 := "hzl"
  want5 := true
  result5 := ValidateECL(input5)
  if want5 != result5 {
      t.Fatalf(`ValidateECL("%v") = %t, wanted %t`, input5, result5, want5)
  }

  input6 := "oth"
  want6 := true
  result6 := ValidateECL(input6)
  if want6 != result6 {
      t.Fatalf(`ValidateECL("%v") = %t, wanted %t`, input6, result6, want6)
  }

  input7 := "blue"
  want7 := false
  result7 := ValidateECL(input7)
  if want7 != result7 {
      t.Fatalf(`ValidateECL("%v") = %t, wanted %t`, input7, result7, want7)
  }

  input8 := "#9998fe"
  want8 := false
  result8 := ValidateECL(input8)
  if want8 != result8 {
      t.Fatalf(`ValidateECL("%v") = %t, wanted %t`, input8, result8, want8)
  }

  input9 := "2003"
  want9 := false
  result9 := ValidateECL(input9)
  if want9 != result9 {
      t.Fatalf(`ValidateECL("%v") = %t, wanted %t`, input9, result9, want9)
  }
}

func TestValidatePID(t *testing.T) {
  input1 := "000312938"
  want1 := true
  result1 := ValidatePID(input1)
  if want1 != result1 {
      t.Fatalf(`ValidatePID("%v") = %t, wanted %t`, input1, result1, want1)
  }

  input2 := "138204893"
  want2 := true 
  result2 := ValidatePID(input2)
  if want2 != result2 {
      t.Fatalf(`ValidatePID("%v") = %t, wanted %t`, input2, result2, want2)
  }

  input3 := "brn"
  want3 := false
  result3 := ValidatePID(input3)
  if want3 != result3 {
      t.Fatalf(`ValidatePID("%v") = %t, wanted %t`, input3, result3, want3)
  }

  input4 := "81043920189573"
  want4 := false
  result4 := ValidatePID(input4)
  if want4 != result4 {
      t.Fatalf(`ValidatePID("%v") = %t, wanted %t`, input4, result4, want4)
  }

  input5 := "00003920189573"
  want5 := false
  result5 := ValidatePID(input5)
  if want5 != result5 {
      t.Fatalf(`ValidatePID("%v") = %t, wanted %t`, input5, result5, want5)
  }

  input6 := "abcdefghi"
  want6 := false
  result6 := ValidatePID(input6)
  if want6 != result6 {
      t.Fatalf(`ValidatePID("%v") = %t, wanted %t`, input6, result6, want6)
  }
}
