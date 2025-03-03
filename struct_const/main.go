package main

import (
	"fmt"
	"time"
)

type Order struct {
	id          string
	items       []string
	amount      float64
	destination string
}

func newOrder(items []string, amount float64, destination string) *Order {
	// generate new id
	myOrder := Order{
		id:          time.Nanosecond.String(),
		items:       items,
		amount:      amount,
		destination: destination,
	}
	return &myOrder
}

func main() {

	o1 := newOrder([]string{"item1", "item2"}, 10.99, "New York")
	fmt.Println(o1)

}
