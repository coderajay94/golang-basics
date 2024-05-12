package main

import "fmt"

func main() {
	fmt.Println("welcome to array class")

	//declaration
	var fruits [4]string

	//intitalization
	fruits[0] = "apple"
	fruits[1] = "orange"
	fruits[3] = "papaya"

	fmt.Println(fruits)

	fmt.Println("fruits array length is:", len(fruits))

	//declaration and intialization together
	var subs = [3]string{"apple", "orange", "papaya"}
	fmt.Println("printing array is ", subs)
}
