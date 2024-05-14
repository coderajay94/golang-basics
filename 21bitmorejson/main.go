package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name    string   `json:"courseName"`
	Price   float64  `json:"price"`
	Teacher string   `json:"teacher"`
	Topics  []string `json:"topics,omitempty"`
	Hours   int      `json:"hours"`
}

func main() {
	fmt.Println("welcome to bit more json tutorial...")

	//EncodeJSON()
	DecodeJSON()
}

func DecodeJSON() {

	stringDataFromWeb := []byte(`{
		"courseName": "java",  
		"price": 12500.1,      
		"teacher": "Raghu sir",
		"topics": [
			"basics",      
			"collections"
		],
		"hours": 2
}`)

	var course Course

	if !json.Valid(stringDataFromWeb) {
		fmt.Println("json is not valid")
	} else {
		fmt.Println("json is valid")
		error := json.Unmarshal(stringDataFromWeb, &course)
		if error != nil {
			panic(error)
		}

		fmt.Printf("ouput is %v\n", course)
	}

	var courseMap map[string]interface{}

	error := json.Unmarshal(stringDataFromWeb, &courseMap)
	if error != nil {
		panic(error)
	}

	for key, value := range courseMap {
		fmt.Printf("for key - %v value is - %v and type is - %T \n", key, value, value)
	}
}

func EncodeJSON() {

	courses := []Course{

		{"java", 12500.10, "Raghu sir", []string{"basics", "collections"}, 2},
		{"python", 10500.10, "Rajiv sir", []string{"basics"}, 4},
		{"golang", 11500.10, "Ajay sir", []string{"basics", "gokit"}, 3},
		{"cpp", 13500.10, "Akash sir", []string{}, 1},
	}

	//request, error := json.MarshalIndent(courses)
	request, error := json.MarshalIndent(courses, "", "\t")
	if error != nil {
		panic(error)
	}

	fmt.Printf("Json Request:%s \n ", request)

}
