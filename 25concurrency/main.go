package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("hello to goroutine")
	go greeter("hello")
	greeter("world")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Second)
	}
}
