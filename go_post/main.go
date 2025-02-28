package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var myurl = "https://dummyjson.com/auth/login"

func main() {
	fmt.Println("Let's learn about post request in golang")
	// response, err := http.Get(myurl)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer response.Body.Close()
	// data, err := io.ReadAll(response.Body)
	// errorHandler(err)
	// fmt.Println(string(data))
	// performPostRequest()
	EncoreJson()
	performPostRequest()

}
func performPostRequest() {
	// Create a new request
	requestBody := strings.NewReader(`
	{
	"username":"emilys",
	"password":"emilyspass",
	"expiresInMins":20
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	errorHandler(err)
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	errorHandler(err)
	fmt.Println(string(data))

}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
	Time     int
}

func EncoreJson() {
	// var user user
	userData := []user{
		{"ifraheem", "123456", 20},
		{"emilys", "123456", 20},
		{"john", "123456", 20},
	}
	// package data into json data
	finalJson, err := json.MarshalIndent(userData, "", "\t")
	errorHandler(err)
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {

}
