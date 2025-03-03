package main

import "fmt"

type UserStatus string

const (
	Active   UserStatus = "active"
	Inactive UserStatus = "inactive"
	Left     UserStatus = "left"
)

type User struct {
	ID       string
	Username string
	Status   UserStatus
}

func main() {
	user := User{
		ID:       "123",
		Username: "john",
		Status:   Left,
	}
	fmt.Println(user)
}
