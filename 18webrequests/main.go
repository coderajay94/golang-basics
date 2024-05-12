package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// read apis for testing from
// https://reqres.in/
const url = "https://reqres.in/api/users/2"

func main() {

	fmt.Println("welcome to the http request in golang")

	response, error := http.Get(url)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	//fmt.Println("printing response object from http request", response)

	//read all response will give you byte array as an response
	resp, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}

	fmt.Println("response is:", (string(resp)))

	//use struct to hold the response object
	data := UserResponse{}
	json.Unmarshal(resp, &data)

	fmt.Println()
	fmt.Println("Email from http response:", data.Data.Email)
}

// created this struct from this site
// https://transform.tools/json-to-go
type UserResponse struct {
	Data struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`
	Support struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}
