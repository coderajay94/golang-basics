package main

import "fmt"

func main() {

	days := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}

	fmt.Println("welcome to the loops")

	//using for loop to initialize, condition & increment
	for index := 0; index < len(days); index++ {
		fmt.Printf("%v day is %v \n", index+1, days[index])
	}

	//using range for index
	for i := range days {
		fmt.Printf("%v day is %v \n", i, days[i])
	}
	fmt.Println()
	//using range for index and value both
	for index, value := range days {
		fmt.Printf("%v day is %v \n", index+1, value)
	}
	fmt.Println()

	val := 11

	//equal to while loop
	for ind := 1; ind <= val; {

		if ind == 11 {
			//just break the loop execution
			break
		}

		if ind == 6 {
			ind++
			//will skip this iteration but continue the loop
			continue
		}
		fmt.Printf("value is %v \n", ind)
		if ind == 10 {
			goto exit
		}
		ind++

	}

	//created labels in golang - label: code
	//and call this lable by -  goto label
exit:
	fmt.Println("called label by- go to exit, so going to exit")

}
