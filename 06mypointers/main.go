package main

import "fmt"

func main() {

	var name string = "pointers"
	fmt.Println("class ", name)

	//value is <nil> for below as its declared but not initalized yet.
	//her *int means this is a pointer to an integer value
	var ptr *int
	fmt.Println("int value pointer is", ptr)

	var num int = 12

	//how to store a pointer to existing value
	var pt *int = &num
	fmt.Println(pt)

	fmt.Println("printing value from pointer ", *pt)

	*pt = *pt + 3

	fmt.Println("updated ptr value:", *pt)

}
