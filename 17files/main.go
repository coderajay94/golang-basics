package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("welcome to writing files...")

	//use os package to create file in current working directory
	file, error := os.Create("./note.txt")
	if error != nil {
		panic(error)
	}

	len, error := io.WriteString(file, "welcome to our course, we will teach you everything about golang \nthank you!")
	checkNilError(error)

	if len > 0 {
		fmt.Println("file updated successfully")
	}

	defer file.Close()

	readFile("./note.txt")
}

func readFile(fileName string) {

	byteArray, error := ioutil.ReadFile(fileName)
	checkNilError(error)

	fmt.Println("content from file: ", string(byteArray))

}

func checkNilError(error error) {
	if error != nil {
		panic(error)
	}
}
