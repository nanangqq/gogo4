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
		"http://google.com",
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
	for _, result := range results {
		fmt.Println(result)
	}
}

func hitURL(url string) error {
	fmt.Println("checking:", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
