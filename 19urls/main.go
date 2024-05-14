package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to urls in golang")

	// const myurl = "https://reqres.in:8080/api/users/2?username=ajakumar&password=admin"

	// urlobj, err := url.Parse(myurl)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(urlobj.Scheme)
	// fmt.Println(urlobj.Host)
	// fmt.Println(urlobj.Path)
	// fmt.Println(urlobj.Port())
	// fmt.Println(urlobj.RawQuery)

	// //this will return url.Values which
	// qparams := urlobj.Query()

	// fmt.Printf("type of query params is %T, \n", qparams)

	// for params := range qparams {
	// 	fmt.Println(params)
	// }
	// fmt.Println(qparams["username"])
	// fmt.Println(qparams["password"])

	// newurl := &url.URL{
	// 	Scheme:   "https",
	// 	Host:     "google.com",
	// 	Path:     "/users/ajaykumar",
	// 	RawQuery: "username=admin?password=admin",
	// }

	// urlstring := newurl.String()
	// fmt.Println(urlstring)

	//num := 1
	//CallPrint(num)
	fmt.Print("fibonacci number at index 6 is:", Fibo(6))

}

func Fibo(num int) int {

	if num < 2 {
		return num
	}
	return Fibo(num-1) + Fibo(num-2)

}

func CallPrint(num int) {

	if num == 10 {
		fmt.Println("printing the number:", num)
		return
	}

	fmt.Println("printing the number:", num)
	CallPrint(num + 1)
}
