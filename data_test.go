package main

import (
	"testing"
)

func Test_readData_fileFound(t *testing.T) {
	input := "./test_resources/test_data.json"

	_, err := readData(input)

	if err != nil {
		t.Errorf("File that exists should not cause an error")
	}
}

func Test_readData_fileNotFound(t *testing.T) {
	input := "./test_resources/not_a_real_file.json"

	_, err := readData(input)

	if err == nil {
		t.Errorf("File that doesn't exist should cause an error")
	}
}

func Test_readData_invalidFile(t *testing.T) {
	input := "./test_resources/bad_data.json"

	_, err := readData(input)

	if err == nil {
		t.Errorf("Invalid file should cause an error")
	}
}

func Test_readData_urlPresent(t *testing.T) {
	input := "./test_resources/test_data.json"

	data, _ := readData(input)

	if data["Twitter"].url != "https://mobile.twitter.com/{}" {
		t.Errorf("Failed to read Twitter url from test_data")
	}

}

func Test_filterByErrorType_message(t *testing.T) {
	filepath := "./test_resources/test_data.json"
	errorType := "message"

	data, _ := readData(filepath)
	filterByErrorType(errorType, data)

	for site := range data {
		if data[site].errorType != errorType {
			t.Errorf("Filtered on errorType for %s but found entry with %s", errorType, data[site].errorType)
		}
	}
}
