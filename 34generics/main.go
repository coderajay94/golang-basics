package main

import "fmt"

//way - 1
/*func Add(num1 int, num2 int) {
	fmt.Println("sum is:", num1+num2)
}*/

//way - 2
//using generics
//create new Type and mention it can use all types of data types
/*func Add[T int | float64 | string](a T, b T) {
	fmt.Println("sum is: ", a+b)
}*/

//way - 3
type Num interface {
	~int | float64 | string
}

func Add[T Num](a T, b T) {
	fmt.Println("sum is: ", a+b)
}

//way - 4
type UserId int

func main() {
	fmt.Println("testing generics...")

	Add(12, 12)
	//what if we change it to float64
	Add(12.5, 12.5)

	Add("ajay", "kumar")

	userId1 := UserId(1)
	userId2 := UserId(2)

	//below line will not work so change add tilde sign ~
	Add(userId1, userId2)
	//if you will use ~ sign, it will take type and see if it resolves to int
}
