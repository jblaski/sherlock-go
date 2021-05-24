package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
  main.go is concerned with accepting and validating user input.
  it will then pass the arguments to the relevant functions.
*/
func main() {
	// get input args
	argsError := validateInput(os.Args[1:])
	if argsError != nil {
		fmt.Println(argsError)
		os.Exit(1)
	}
	username := os.Args[1]
	fmt.Println("Supplied args: " + strings.Join(os.Args[1:], ", "))

	// unmarshall
	data, err := readData("./resources/data.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Loaded data for " + fmt.Sprint(len(data)) + " sites")

	// filter down to currently supported types
	filterByErrorType("status_code", data)
	fmt.Println("Filtered down to " + fmt.Sprint(len(data)) + " sites")

	// make calls
	fmt.Println(makeCalls(data, "status_code", username))
}

func validateInput(args []string) error {
	if len(args) != 1 {
		return errors.New("Error: expected 1 argument, got " + fmt.Sprint(len(args)))
	}
	return nil
}
