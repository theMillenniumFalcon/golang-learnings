package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

func main() {
	// go greeter("Hello")
	// greeter("World")

	websitelist := []string{
		"http://google.com",
		"http://fb.com",
		"http://youtube.com",
		"http://twitter.com",
		"http://go.dev",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endPoint string) {
	defer wg.Done()

	res, err := http.Get(endPoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endPoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endPoint)
	}
}
