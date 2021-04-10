package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url string
	status string
}

var errRequestFailed = errors.New("req error")

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
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
		// result := "ok"
		// err := hitURL(url)
		// if err != nil {
		// 	result = "fail"
		// }
		// results[url] = result

		go hitURL(url, c)
	}
	// results["hi"] = "hello"
	// fmt.Println(results)

	for i := 0; i < len(urls); i++ {
		result := <-c  
		results[result.url] = result.status
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string, channel chan<- requestResult) {
	fmt.Println("checking:", url)
	res, err := http.Get(url)
	fmt.Println(res.StatusCode)
	status := "ok"
	if err != nil || res.StatusCode >= 400 {
		status = "fail"
	} 
	channel <- requestResult{url: url, status: status}
}
