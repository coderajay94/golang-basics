package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("welcome to KBC!")

	bank := make(map[int]float64)
	//fmt.Println(len(bank))

	for {

		rand.Seed(time.Now().UnixNano())
		questionID := rand.Intn(4) + 1

		_, err := bank[questionID]
		if !err {
			KBCQuestions(questionID)
			reader := bufio.NewReader(os.Stdin)
			answer, _ := reader.ReadString('\n')

			ans, _ := strconv.ParseFloat(strings.TrimSpace(answer), 64)
			bank[questionID] = float64(ans)
		}

		if len(bank) == 4 {
			total := 0
			for key, _ := range bank {
				if key == 1 && bank[key] == 1 {
					total = total + 1
				} else if key == 2 && bank[key] == 2 {

					total = total + 1
				} else if key == 3 && bank[key] == 2 {

					total = total + 1
				} else if key == 4 && bank[key] == 1 {

					total = total + 1
				}
			}
			if total == 4 {
				fmt.Println("You Won!!!!!!!!!!! All Answers are correct")
			} else {
				fmt.Println("-----------------Result---------------------")
				fmt.Println("1. right answer is : option 1. Murgi Pehle aayi")
				fmt.Println("2. right answer is : option 2. Akal Badi")
				fmt.Println("3. right answer is : option 2. Godhe pe gadha")
				fmt.Println("4. right answer is : option 1. 9 Lakha haar")
				fmt.Println("Your total score is:", total)
				fmt.Println("--------------------------------------------")
			}
			read := bufio.NewReader(os.Stdin)
			ans, _ := read.ReadString('\n')
			fmt.Println("-------------------------------------", ans)
			goto Exit
		}
		//RollDice()
	}
Exit: // More logic here that needs to be executed
	fmt.Println("Thanks for playing game with us!")

	// to set the source of random numbers, to read from current time in milliseconds

}

func KBCQuestions(no int) {

	switch no {
	case 1:
		{
			fmt.Printf("1. Murgi Pehla aaya ya 2. anda? \n")
			//fallthrough - use this if you want to test remaining scenarios as well
		}
	case 2:
		{
			fmt.Printf("1. Bhains Badi ya 2. Akal? \n")
		}
	case 3:
		{
			fmt.Printf("1. gadhe pe ghoda ya 2. ghode pe gadha? \n")
		}
	case 4:
		{
			fmt.Printf("1. 9 Lakha haar ya 2. 9 lakha car \n")
		}
		// default:
		// 	{
		// 		fmt.Println("default case")
		// 	}
		// case 5:
		// 	{
		// 		fmt.Printf("")
		// 	}
		// case 6:
		// 	{
		// 		fmt.Printf("Got %v! you can open now and move by %v  ", dice, dice)
		// 	}

	}
}

func RollDice() {
	rand.Seed(time.Now().UnixNano())

	dice := rand.Intn(6) + 1

	switch dice {
	case 1:
		{
			fmt.Printf("Got %v! you can open now  ", dice)
		}
	case 2:
		{
			fmt.Printf("Got %v! you can move by  %v", dice, dice)
		}
	case 3:
		{
			fmt.Printf("Got %v! you can move by  %v", dice, dice)

		}
	case 4:
		{
			fmt.Printf("Got %v! you can move by  %v", dice, dice)

		}
	case 5:
		{
			fmt.Printf("Got %v! you can move by  %v", dice, dice)
		}
	case 6:
		{
			fmt.Printf("Got %v! you can open now and move by %v  ", dice, dice)
		}

	}
}
