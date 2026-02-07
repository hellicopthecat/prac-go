package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

type requestResult struct {
	url    string
	status string
}

func main() {
	var results = map[string]string{}
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}

	for _, url := range urls {
		go hitUrl(url, c)
	}
	for range len(urls) {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitUrl(url string, c chan<- requestResult) {
	fmt.Println("Checking URL")
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		c <- requestResult{url: url, status: "Failed"}
	} else {
		c <- requestResult{url: url, status: status}
	}
}
