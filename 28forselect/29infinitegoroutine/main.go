package main

import (
	"fmt"
	"time"
)

func main() {

	//1. these kind of goroutines can create memory leaks and would be consuming memory and resources
	//running in background

	//create infinite task to keep running
	/*	fmt.Println("started main...")
		go func() {
			for {
				select {
				default:
					fmt.Println("doing work...")
				}
			}
		}()

		time.Sleep(5 * time.Second)
	*/

	//2. control these function calls from main function
	// by creating channel

	done := make(chan bool)

	go dowork(done)

	time.Sleep(3 * time.Second)

	close(done)
}

// read only channel by putting <-
func dowork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("doing work...")
		}
	}
}
