package main

import "fmt"

type payment struct {
	gateway paymenter
}

type paymenter interface {
	pay(amount float32)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake payment", amount)
}

func main() {
	// razorpayGw := razorpay{}
	fakeGw := fakepayment{}
	newPayment := payment{
		gateway: fakeGw,
	}
	newPayment.gateway.pay(100.0)

}
