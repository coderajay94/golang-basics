package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte | []rune | string | int
}

// generics in structs
type User[T CustomData] struct {
	ID   int
	Name string
	Data T
}

//generics in maps

type customMap[T comparable, V int | string] map[T]V

func main() {
	//generics in maps
	m := make(map[int]string)
	m[3] = "ajay kumar"

	fmt.Println(m)
	//----------------------------------------------------------------

	//generics in structs
	fmt.Println("Starting")

	usr1 := User[string]{
		ID: 1, Name: "ajay kumar", Data: "abc",
	}

	fmt.Println(usr1)

	usr2 := User[int]{
		ID: 1, Name: "ajay kumar", Data: 1234,
	}
	fmt.Println(usr2)
}
