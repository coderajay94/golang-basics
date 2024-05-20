package main

import (
	"fmt"
	"time"
)

// 2 ways you can do the join in goroutines
// first is you can use waitgroups and second is channels
func main() {

	fmt.Println("testing fork join model")
	now := time.Now()

	channel1 := make(chan struct{})
	//without goroutines it takes around elapsed time: 923.7032ms
	//with channels it takes elapsed time: 500.0014ms

	//lets add go and make it goroutine
	go greeter1(channel1) //-> forked from here
	go greeter2(channel1)
	go greeter3(channel1)
	go greeter4(channel1)

	//we will have to join here to the main
	//read from channel
	<-channel1
	<-channel1
	<-channel1
	<-channel1
	fmt.Println("elapsed time:", time.Since(now))

}

func greeter1(channel1 chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("executed greeter1")
	channel1 <- struct{}{}
}
func greeter2(channel2 chan struct{}) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("executed greeter2")
	channel2 <- struct{}{}
}
func greeter3(channel3 chan struct{}) {
	fmt.Println("executed greeter3")
	channel3 <- struct{}{}
}
func greeter4(channel4 chan struct{}) {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("executed greeter4")
	channel4 <- struct{}{}
}
