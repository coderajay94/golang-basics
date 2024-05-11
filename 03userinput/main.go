package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello welcome to Rating system!")

	fmt.Println("How would you like to  rate our pizza")
	input := bufio.NewReader(os.Stdin)

	//read till user doesn't give next line or enter
	value, err := input.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading", err)
	}
	fmt.Println("Thanks for rating us, ", value)
}
