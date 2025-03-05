package main

// import (
// 	"fmt"
// 	"time"
// )

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("Sending email to", email)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func main() {
// 	println("let's go in deep of channels")
// 	emailChan := make(chan string, 8)
// 	done := make(chan bool)

// 	go emailSender(emailChan, done)

// 	for i := 0; i < 10; i++ {
// 		emailChan <- fmt.Sprintf("%d@gmail.com", i)
// 	}
// 	fmt.Println("done sending email")
// 	close(emailChan)

// 	<-done

// }
