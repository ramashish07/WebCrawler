package main

import (
	"fmt"
	"log"
	"main/router"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
}