package main

import "fmt"

type Student struct {
	Name   string
	RollNo int
	Fee    float64
	Status bool
}

func main() {

	fmt.Println("hello world! to struct")

	ajay := Student{
		Name:   "ajay Kumar",
		RollNo: 123,
		Fee:    2000,
		Status: true,
	}

	fmt.Println(ajay)

	fmt.Printf("%+v\n", ajay)
}
