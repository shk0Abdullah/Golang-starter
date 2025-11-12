package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Printf("In change Number func %d ", *num)
}

func main() {
	num := 1
	changeNum(&num)
	fmt.Printf("\n%d", num)
}
