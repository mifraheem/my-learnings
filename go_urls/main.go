package main

import (
	"fmt"
	"io"
	"net/http"
)

const myurl string = "https://jsonplaceholder.typicode.com/todos"

func main() {
	fmt.Println("Welcome to Handling URLs in Golang")
	// res, _ := url.Parse(myurl)
	// fmt.Println(res)
	// fmt.Println(res.Scheme)
	// fmt.Println(res.Host)
	// fmt.Println(res.Path)
	// fmt.Println(res.Port())

	response, err := http.Get(myurl)
	errorHandler(err)
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	// fmt.Println(response.Header)
	// fmt.Println(response.Cookies())
	fmt.Println(response.ContentLength)

	content, err := io.ReadAll(response.Body)

	errorHandler(err)
	fmt.Println(string(content))

}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
