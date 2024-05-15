package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("welcome to mod sections...")
	router := mux.NewRouter()
	router.HandleFunc("/home", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Greeter() {
	fmt.Println("welcome from Greeter method...")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to the mod classes</h1>"))
}
