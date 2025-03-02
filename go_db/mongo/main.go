package main

import (
	"fmt"
	"log"
	"mongo/router"
	"net/http"
)

func main() {
	fmt.Println("Mongo db Connectivity in golang")
	r := router.Router()

	fmt.Println("Server is running at port 4000\n http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
