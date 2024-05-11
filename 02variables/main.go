package main

import "fmt"

//outside main variable declaration
var JWTToken = "token data"

//constants declaration
const APIKey = "API key"

func main() {
	fmt.Println("variables")

	var num1 int = 23

	fmt.Printf("%T  and %v \n", num1, num1)

	var monthly float32 = 1200.12349
	var salary float64 = 120990.1234567890111213141516171811920

	fmt.Println("monthly and yearly salary", monthly, salary)

	fmt.Println(JWTToken)

	fmt.Println(APIKey)

	//only works inside the function
	//automaticlly detects the type
	data := "hello world"
	fmt.Printf("%v, %T \n", data, data)

}
