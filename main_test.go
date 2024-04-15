package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"testing"
)

var numOkRequest int

func Test1(t *testing.T) {
	totalRequest := 10
	wg := sync.WaitGroup{}
	wg.Add(totalRequest)
	for range totalRequest {
		go func() {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			_, err := http.Get("http://localhost/")

			if err == nil {
				numOkRequest++
			}

		}()
	}
	wg.Wait()

	if numOkRequest == totalRequest {
		fmt.Println("Ok")
	} else {
		t.Fatalf("error making http request: " + strconv.Itoa(numOkRequest))
	}
}
