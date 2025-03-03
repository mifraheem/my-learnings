package main

import (
	"fmt"
)

var generateId = func() func() int {
	var id int = 0
	return func() int {
		id += 1
		return id
	}
}()

type User struct {
	id   int
	Name string
	Age  int
	tst  string
}

func (u *User) assignId() {
	u.id = generateId()
}

func main() {
	fmt.Println("Structs method in golang")
	u1 := User{
		Name: "John",
		Age:  25,
	}
	u1.assignId()
	var u2 User
	u2.Name = "Ifraheem"
	u2.Age = 24
	u2.assignId()

	fmt.Println(u1)
	fmt.Println(u2)
}
