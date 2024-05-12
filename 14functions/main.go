package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")

	greeter("welcome")

	sum := adder(12, 12)
	fmt.Println("sum of 2 numbers is:", sum)

	total := greatAdder(12, 13, 15, 17, 19)
	fmt.Println("sum of great adder is:", total)

	//	nums := []int{12, 13, 15, 17}
	//note:  check why we can't pass nums in below greatAdder function
	total2 := greatAdder(12, 143, 2333, 1223)
	fmt.Println("sum of great adder2 is:", total2)

	func() {
		fmt.Println("calling anonymous function")
	}()
	//() will automatically calls the anonymous function

	line := func() string {
		return "hello world"
	}()

	fmt.Printf("line is %v \n", line)

}

//below ...int is same as []int - type is slice
func greatAdder(nums ...int) int {
	fmt.Printf("type of nums is %T \n", nums)
	total := 0
	for _, value := range nums {
		total = total + value
	}
	return total
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func greeter(greet string) {
	fmt.Println("hello! ", greet)
}
