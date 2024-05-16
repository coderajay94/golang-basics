package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coderajay94/mongoapi/router"
)

func main() {
	fmt.Println("welcome to mongo db tutorial")

	router := router.Router()

	fmt.Println("Started the server")
	log.Fatal(http.ListenAndServe(":4000", router))
	fmt.Println("Server Started on port 4000...")
}
