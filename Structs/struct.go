package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        int
	amount    float32
	status    string
	createdAt time.Time
}

func (o *Order) ChangeStatus(status string) {

	o.status = status

}

func main() {

	// var order Order
	order := Order{
		id:        32,
		amount:    60.1,
		status:    "completed",
		createdAt: time.Now(),
	}
	order.ChangeStatus("pending")
	fmt.Println("My order Struct: ", order)

}
