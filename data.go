package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// For loading data from file

// type SiteData struct {
// 	errorType          string `json:"errorType"`
// 	url                string `json:"url"`
// 	urlMain            string `json:"urlMain"`
// 	username_claimed   string `json:"username_claimed"`
// 	username_unclaimed string `json:"username_unclaimed"`
// }

func readData() {
	// Open our jsonFile
	jsonFile, err := os.Open("data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(string(byteValue))
}
