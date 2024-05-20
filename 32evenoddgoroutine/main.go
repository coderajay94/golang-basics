package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("printing numbers using 2 goroutines")

	wg := &sync.WaitGroup{}
	//create 2 channels
	odd := make(chan int)
	even := make(chan int)

	wg.Add(2)

	go putOdd(odd, wg)
	go putEven(even, wg)

	for i := 1; i <= 5; i++ {
		fmt.Println(<-odd)
		fmt.Println(<-even)
	}

	wg.Wait()
}

func putOdd(odd chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		odd <- i
	}
}

func putEven(even chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		even <- i
	}
}
