package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type SiteData struct {
	errorType         string
	url               string
	urlMain           string
	usernameClaimed   string
	usernameUnclaimed string
}

func readData(filepath string) (map[string]SiteData, error) {

	// open the file
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// unmarshall json
	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return nil, err
	}

	// key site name, value site data
	var siteData = make(map[string]SiteData)

	// convert to desired structure
	for key := range payload {
		data := payload[key].(map[string]interface{})
		siteData[key] = SiteData{
			errorType:         fmt.Sprintf("%s", data["errorType"]),
			url:               fmt.Sprintf("%s", data["url"]),
			urlMain:           fmt.Sprintf("%s", data["urlMain"]),
			usernameClaimed:   fmt.Sprintf("%s", data["username_claimed"]),
			usernameUnclaimed: fmt.Sprintf("%s", data["username_unclaimed"]),
		}
	}

	return siteData, nil
}
