package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// way-1
// you can use constraints which has alredy created default values for u as T
func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var vals []T
	for _, value := range values {
		val := mapFunc(value)
		vals = append(vals, val)
	}
	return vals
}

// way-2
// passing function into the arguments
func MapValues2[T int | float64](values []T, mapFunc func(T) T) []T {
	var vals []T
	for _, value := range values {
		val := mapFunc(value)
		vals = append(vals, val)
	}
	return vals
}

func main() {
	fmt.Println("testing generics in functions")
	arr := []int{11, 21, 3, 14}
	fmt.Println(arr)

	//MapValues method only accepts integers, not float
	//we can use generics to solve this problem
	processed := MapValues(arr, func(num int) int {
		return 2 * num
	})

	fmt.Println(processed)

	arr2 := []float64{11.1, 21.2, 3.3, 14.4}
	fmt.Println(arr2)

	processed2 := MapValues2(arr2, func(num float64) float64 {
		return 2 * num
	})
	fmt.Println(processed2)

	arr3 := []float64{11.1, 21.2, 3.3, 14.4}
	fmt.Println(arr3)

	processed3 := MapValues(arr2, func(num float64) float64 {
		return 2 * num
	})
	fmt.Println(processed3)

}
