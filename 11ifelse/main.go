package main

import "fmt"

func main() {

	RetryLogin(10)

}

func RetryLogin(count int) {
	if count <= 10 {
		fmt.Println("user is permanently blocked")
	} else if retry := 3; count <= retry {
		fmt.Println("user is blocked for 24 hours, try again later")
	} else {
		fmt.Println("user can try logging in again")
	}
}
