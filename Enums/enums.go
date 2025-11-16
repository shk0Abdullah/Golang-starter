package main

import "fmt"

// enumerated types
// we use custom types to craete the nums

// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Delivered
// 	Dispatched
// )
type OrderStatus string

const (
	Received   OrderStatus = "received"
	Delivered              = "delivered"
	Dispatched             = "dispatched"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status Changes: ", status)
}

func main() {
	changeOrderStatus(Received)
}
