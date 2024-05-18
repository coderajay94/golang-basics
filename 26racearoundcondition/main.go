package main

import (
	"fmt"
	"sync"
)

var list = []int{}

var wg = &sync.WaitGroup{}
var mutex = &sync.Mutex{}

func main() {

	fmt.Println("testing race condition....")

	wg.Add(1)
	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("adding 1 into the list...")
		mutex.Lock()
		list = append(list, 1)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("adding 2 into the list...")
		mutex.Lock()
		list = append(list, 2)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("adding 3 into the list...")
		mutex.Lock()
		list = append(list, 3)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	wg.Wait()
	func(mutex *sync.Mutex) {
		mutex.Lock()
		fmt.Println(list)
		mutex.Unlock()
	}(mutex)
}
