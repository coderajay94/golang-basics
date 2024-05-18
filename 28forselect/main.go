package main

import "fmt"

func main() {

	chars := []string{"a", "b", "c"}
	//if you will use unbuffered channel it will be asynchrounous and blocked
	// till the time it will not be starting read, writer goroutine will be in waiting state
	//so buffered channel will not be in waiting state once it submits the data
	// and can be closed not into the unbuffered channel
	bufferedChannel := make(chan string, 3)

	//fmt.Println("checking for select condition in concurrency patterns")

	for _, character := range chars {
		select {
		case bufferedChannel <- character:
		}
	}

	close(bufferedChannel)

	for result := range bufferedChannel {
		fmt.Println(result)
	}
}
