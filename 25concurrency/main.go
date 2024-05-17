package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var websites = []string{
	"https://www.google.com",
	"https://www.fb.com",
	"https://www.github.com",
	"https://www.openai.com",
	"http://www.instagram.com",
}

var wg sync.WaitGroup //these are pointers

func main() {
	//fmt.Println("hello to goroutine")
	//go greeter("hello")
	//greeter("world")

	for _, website := range websites {
		go getStatusCode(website)
		wg.Add(1)
	}

	wg.Wait()
}

func getStatusCode(website string) {
	defer wg.Done()
	resp, err := http.Get(website)
	if err != nil {
		fmt.Printf("error occured while getting status code %v\n", website)
		log.Fatal(err)
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
