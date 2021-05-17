package main

import (
  "fmt"
  "os"
  "strings"
  "errors"
)

/*
  main.go is concerned with accepting and validating user input.
  it will then pass the arguments to the relevant functions.
*/
func main() {
  argsError := validateInput(os.Args[1:])
  if argsError != nil {
    fmt.Println(argsError)
    os.Exit(1)
  }
  fmt.Println("Supplied args: " + strings.Join(os.Args[1:], ", "))
}

func validateInput(args []string) error {
  if len(args) != 1 {
    return errors.New("Error: expected 1 argument, got " + fmt.Sprint(len(args)))
  }
  return nil
}
