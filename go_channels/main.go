package main

import (
	"fmt"
	"time"
)

func sum(result chan int, n int, n2 int) {
	time.Sleep(time.Second * 5)
	sumData := n + n2
	result <- sumData
}

func main() {
	result := make(chan int, 2)
	go sum(result, 10, 20)

	go sum(result, 30, 40)

	res := <-result
	res2 := <-result
	fmt.Println(res)
	fmt.Println(res2)
}
