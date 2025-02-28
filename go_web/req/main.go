package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	api := "https://jsonplaceholder.typicode.com/todos"
	response, err := http.Get(api)
	errorHandler(err)
	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	errorHandler(err)
	fmt.Println(string(data))
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
