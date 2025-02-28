package main

import "fmt"

func main() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	fmt.Println("this is duffer program")
	myDefer()
}

func myDefer() {
	test := 0
	for test < 10 {
		defer fmt.Println(test)
		test++
	}
}
