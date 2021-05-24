package main

import (
	"fmt"
	"net/http"
	"strings"
)

// return a list of strings representing sites where the user was found
func makeCalls(siteData map[string]SiteData, errorType string, username string) []string {
	var foundAt = []string{}
	for siteName := range siteData {
		usernameLink := constructLink(siteData[siteName], username)
		if checkLink(usernameLink) {
			fmt.Println(usernameLink)
			foundAt = append(foundAt, usernameLink)
		}
	}
	return foundAt
}

func constructLink(site SiteData, username string) string {
	return strings.ReplaceAll(site.url, "{}", username)
}

func checkLink(link string) bool {
	resp, _ := http.Get(link) //TODO check the error
	return resp.StatusCode != 404
}
