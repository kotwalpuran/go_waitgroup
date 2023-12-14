package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	urls := []string{
		"https://www.amazon.com",
		"https://google.com",
		"https://fb.com",
		"https://twitter.com",
		"https://gitlab.com",
		"https://github.com",
		"https://twilio.com",
		"https://facebook.com",
		"https://generativeai.net",
	}
	for _, v := range urls {
		go doHttp(v)
		wg.Add(1)
	}
	wg.Wait()
}

func doHttp(u string) {
	defer wg.Done()
	// set 3 seconds timeout
	httpClient := http.Client{
		Timeout: 3 * time.Second,
	}
	r, e := httpClient.Get(u)
	if e != nil {
		fmt.Printf("Something went wrong, error: %v\n", e)
	} else {
		fmt.Printf("url: %s, response status Code: %d\n", u, r.StatusCode)
	}
}
