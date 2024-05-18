package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var websites = []string{
	"https://www.google.com",
	"https://www.fb.com",
	"https://www.github.com",
	"https://www.openai.com",
	"http://www.instagram.com",
}

var wg sync.WaitGroup //these are pointers
var mutex sync.Mutex

func main() {
	//fmt.Println("hello to goroutine")
	//go greeter("hello")
	//greeter("world")

	for _, website := range websites {
		go getStatusCode(website)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(website string) {
	defer wg.Done()
	resp, _ := http.Get(website)
	if resp.StatusCode != 200 {

		//as many other threads can also try to access signals can cause problems
		// so lock and update and others can wait till operation is completed
		mutex.Lock()
		signals = append(signals, website)
		mutex.Unlock()
		fmt.Printf("Status code: %v for website is %v \n", resp.StatusCode, website)
		//fmt.Printf("error occured while getting status code %v\n", website)
		//log.Fatal(err)
	} else {
		fmt.Printf("Status code: %v for website is %v \n", resp.StatusCode, website)
	}
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		fmt.Println(s)
// 		time.Sleep(1 * time.Second)
// 	}
// }
