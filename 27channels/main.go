package main

import "fmt"

func main() {

	channel := make(chan string)
	channel2 := make(chan string)
	fmt.Println("testing the channels in golang")

	go func() {
		channel <- "channel 1 source 1"
	}()
	// go func() {
	// 	channel <- "channel 1 source 2"
	// }()
	// go func() {
	// 	channel2 <- "source channel2 1"
	// }()
	go func() {
		channel2 <- "source channel2 2"
	}()

	//fmt.Println(<-channel, <-channel)
	select {
	case msgchannel1 := <-channel:
		fmt.Println(msgchannel1)
	case msgchannel2 := <-channel2:
		fmt.Println(msgchannel2)
	}
}

//im above examples main is also a goroutine
//so anoanymous functions are writting into teh channel
//and main goroutine is reading from the channel
//so print statement turns it into blocking call as its waiting for the response for both values
//
//main goroutine wants to finish the execution
//but we have to write the join model for our go-routines so that main can wait and work as fork-join model
//here print statement from main goroutine is doing joining the goroutine to main goroutine
