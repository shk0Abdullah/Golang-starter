package main

import (
	"fmt"
)

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Payment Done using Razorpay", amount)

}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)

}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Getting paid using Paypal:", amount)
}

func main() {
	PaymentGw := paypal{}
	Newpayment := payment{
		gateway: PaymentGw,
	}
	Newpayment.makePayment(100)
}
