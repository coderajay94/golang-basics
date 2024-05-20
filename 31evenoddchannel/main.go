package main

import "fmt"

func main() {

	fmt.Println("welcome to print even odd using channels")

	even := make(chan int)
	odd := make(chan int)

	go printNumbers(odd, even)

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

}

func printNumbers(odd chan int, even chan int) {

	for {
		select {
		case o := <-odd:
			fmt.Println(o)
		case e := <-even:
			fmt.Println(e)
		}
	}

}
