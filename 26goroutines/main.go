package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	greeter("Hello")
	greeter("World")

	websiteList := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
	}

	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1)
	}

	wg.Wait()
}

var wg sync.WaitGroup //pointers

func greeter(s string) {

	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done() //very good use of defer to ensure that Done is called even if there's an error
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", endpoint, err)
	}

	fmt.Printf("Status code for %s: %d\n", endpoint, res.StatusCode)

}
