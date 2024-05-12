package main

import "fmt"

func main() {

	fmt.Println("hello")

	defer fmt.Println("world")
	CallMe()
	defer fmt.Println("1. welcome")
	defer fmt.Println("2. wannacam")
	defer fmt.Println("3. namaste")

	//defer will put at last
	//another defer comes and will execute at last, but before the first defer and so on

}

//functions defer will execute seprately
func CallMe() {
	defer fmt.Println("a. idli")
	defer fmt.Println("b. sambhar")
	defer fmt.Println("c. uthappam")
}

//put into stack, so last defer will exexute first
// welcome,
// wannacam,
// namaste
