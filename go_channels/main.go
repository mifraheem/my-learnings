package main

import "fmt"

// func sum(result chan int, n int, n2 int) {
// 	time.Sleep(time.Second * 5)
// 	sumData := n + n2
// 	result <- sumData
// }

func main() {
	// channel
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "ping"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received from chan1:", chan1Val)
		case cha2Val := <-chan2:
			fmt.Println("Received from chan2:", cha2Val)
		}
	}

	// result := make(chan int, 2)
	// go sum(result, 10, 20)

	// go sum(result, 30, 40)

	// res := <-result
	// res2 := <-result
	// fmt.Println(res)
	// fmt.Println(res2)
}
