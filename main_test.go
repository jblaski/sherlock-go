package main

import (
  "testing"

)

func Test_validateInput_3args(t *testing.T) {
  input := []string{"arg1", "arg2", "arg3"}

  actual := validateInput(input)
  if actual == nil {
    t.Errorf("Input of 3 arguments should return an error")
  }
}

func Test_validateInput_0args(t *testing.T) {
  input := []string{}

  actual := validateInput(input)

  if actual == nil {
    t.Errorf("Input of 0 argument should return an error")
  }
}

func Test_validateInput_1arg(t *testing.T) {
  input := []string{"arg1"}

  actual := validateInput(input)

  if actual != nil {
    t.Errorf("Input of 1 arguments should be valid")
  }
}
