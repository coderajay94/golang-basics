package main

import "fmt"

func main() {
	fmt.Println("welcome to map classes")

	langs := make(map[string]string, 4)

	langs["py"] = "python"
	langs["js"] = "javascript"
	langs["rb"] = "ruby"

	for key, value := range langs {
		fmt.Printf("full Form of the %v is: %v \n", key, value)
	}

	_, isAvailable := langs["py"]

	if isAvailable {
		fmt.Println("py is available")
	}

}
