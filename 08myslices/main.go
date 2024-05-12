package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")

	var slice = []string{"ahhha", "world", "welcome"}

	fmt.Println(slice)

	slice = append(slice, "wannakum", "padhariye")

	fmt.Println(slice)

	//without any size
	fmt.Printf("Type of slice is: %T \n", slice)

	// slice.reverse()

	//read till ends
	fmt.Println("slice is sliced:", slice[2:])

	//last one is not inclusieve
	fmt.Println("slice is sliced:", slice[1:4])

	//use make function to create slice

	var vegs = make([]string, 3)
	//append only adds new element into the slice, it will leave empty the existing ones
	vegs = append(vegs, "potato", "banana", "tomato")

	fmt.Println("slice using make function", len(vegs))

	//test slices methods

	nums := make([]int, 5)

	nums[0] = 89
	nums[1] = 13
	nums[2] = 23
	nums[3] = 44
	nums[4] = 12

	sort.Ints(nums)
	var isSorted bool = sort.IntsAreSorted(nums)
	fmt.Println("sorted slice:", isSorted)

	//reverse method takes interface as in input, so need ot convert using sort.IntSlice(nums)
	reverse := sort.Reverse(sort.IntSlice(nums))
	fmt.Printf("reverse sorted slice: %v with type %T \n", reverse, reverse)
	fmt.Println("----------------------------------------------------------------")
	//how to remove element from a slice

	var courses = []string{"react", "java", "python", "golang", "ruby"}
	fmt.Println(courses)
	index := 2

	// it will not take it as a syntax, so use ... along with it
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
