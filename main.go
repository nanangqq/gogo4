package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("req error")

func main() {
	var results = make(map[string]string)
	urls := []string{
		"https://google.com",
		"https://www.airbnb.com",
		"https://www.reddit.com",
		"https://soundcloud.com",
		"https://facebook.com",
		"https://instagram.com",
		"https://nomadcoders.co",
		"https://amazon.com",
	}

	for _, url := range urls {
		// fmt.Println(url)
		result := "ok"
		err := hitURL(url)
		if err != nil {
			result = "fail"
		}
		results[url] = result
	}
	// results["hi"] = "hello"
	// fmt.Println(results)

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("checking:", url)
	res, err := http.Get(url)
	fmt.Println(res.StatusCode)
	if err != nil || res.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
