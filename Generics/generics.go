package main

import "fmt"

// any == interface{}
// Better way int | string

// You can also use comparable keyword instead without writing all types like int | string | struct | float32 | bool instead of this use comparable keyword

type stack[T any] struct {
	element []T
}

func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	printSlice([]int{1, 23, 4})
	printSlice([]string{"1", "2Abdullah3", "6ali9"})

	// When using struct
	myStack := stack[string]{
		element: []string{"golang"},
	}
	fmt.Println(myStack)

}
