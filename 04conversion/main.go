package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Please rate your pizza")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error while converting your input,", err.Error())
		panic(err)
	} else {
		fmt.Println("Thanks for rating us, ", rating)
	}

}
