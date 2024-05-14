package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const geturl = "http://localhost:8000/get"

func main() {

	fmt.Println("calling apis via get/post methods...")

	//CallGetService(geturl)
	CallPostService()
}

func CallPostService() {

	geturl := "http://localhost:8000/post"

	//use string concatenation or string formatting to insert the value of geturl into the JSON.

	payload := strings.NewReader(`{
		"country":"India",
		"state":"Punjab",
		"city":"Abohar",
		"pincode":152116,
		"phone":"9872079423",
		"url":"` + geturl + `"

	}`)

	resp, err := http.Post(geturl, "application/json", payload)
	if err != nil {
		fmt.Println("error while posting request: ")
		panic(err)
	}

	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response from post request is: ", string(response))
}

func CallGetService(geturl string) {

	response, err := http.Get(geturl)

	if err != nil {
		fmt.Println("error getting response from get api")
		panic(err)
	}

	respbyte, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error reading response from api response")
		panic(err)
	}

	//way 1 to get response
	output := string(respbyte)

	fmt.Println(output)

	//way 2 to get response

	var stringsbuilder strings.Builder
	stringsbuilder.WriteString(output)
	fmt.Println("printing response:", stringsbuilder.String())

}
