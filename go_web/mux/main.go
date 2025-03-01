package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Name    string
	Age     int
	Address string
}

var users []User
var allUsers = []User{
	{"John", 25, "New York"},
	{"Jane", 30, "London"},
	{"Bob", 35, "Paris"},
}

func main() {
	fmt.Println("Learning about Mux in Go")
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/users", listUsersHandler).Methods("GET")
	r.HandleFunc("/users/{age}", userByAge)

	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("Hello, World!")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang Mux series </h1>"))
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing users")

	data, err := json.MarshalIndent(allUsers, "", "\t")
	if err != nil {
		panic(err)
	}
	w.Write(data)
}

func userByAge(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	agestr := vars["age"]
	fmt.Println(agestr)
	age, err := strconv.Atoi(agestr)
	errorHandler(err)
	target_users := []User{}
	for _, user := range allUsers {
		fmt.Println("Looping through users")
		if user.Age == age {
			fmt.Println("getting target user")
			target_users = append(target_users, user)
			break
		}
	}

	data, err := json.MarshalIndent(target_users, "", "\t")
	errorHandler(err)
	w.Write(data)
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
