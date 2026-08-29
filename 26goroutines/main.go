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
		wg.Add(1)
		go getStatusCode(website)
	}

	wg.Wait()
	fmt.Println("\nSignals:", signals)
}

var wg sync.WaitGroup
var mut sync.Mutex
var signals = []string{}

func greeter(s string) {

	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", endpoint, err)
		return
	}
	defer res.Body.Close()

	mut.Lock()
	signals = append(signals, fmt.Sprintf("%s => %d", endpoint, res.StatusCode))
	mut.Unlock()

	fmt.Printf("Status code for %s: %d\n", endpoint, res.StatusCode)

}
