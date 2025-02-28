package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("### Welcome to Golang Todo Application...! ###\n")
	fmt.Println("Server is running on http://localhost:8080")
	http.HandleFunc("/", helloGo)
	http.HandleFunc("/show", showStudents)

	http.ListenAndServe(":8080", nil)

}

func helloGo(writer http.ResponseWriter, request *http.Request) {
	var greeting string = "Hello, am using Golang!"
	fmt.Fprintf(writer, greeting)
}
func showStudents(writer http.ResponseWriter, request *http.Request) {
	var students = []string{"John", "Doe", "Jane", "Doe"}
	for _, std := range students {
		fmt.Fprintln(writer, std)
	}
}
