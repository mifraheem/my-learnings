package main

import "fmt"

func main() {
	us1 := User{"Ifraheem", "admin123", "ifraheem@gmail.com"}
	us2 := User{"John", "john123", "john@gmail.com"}
	fmt.Println(us1, us2)
	fmt.Println(us1.CheckPassword())
	newD := us1.ChangeEmail("ifraheem@google.com")
	fmt.Println(newD)

}

type User struct {
	Username string
	Password string
	Email    string
}

func (u User) CheckPassword() string {
	return u.Password
}

func (u User) ChangeEmail(mail string) string {
	u.Email = mail
	return u.Email
}
