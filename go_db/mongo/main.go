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

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is running at port 4000\nlocalhost:4000")
}
