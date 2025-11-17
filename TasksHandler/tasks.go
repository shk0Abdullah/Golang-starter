package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	t := []string{}
	fmt.Println("Welcome to CLI Task Handler")
	for {
		fmt.Print("1.Add task\n2.Remove Task\n3.Show Tasks\n4.Exit\n")
		var tasks int
		fmt.Scanln(&tasks)
		if tasks == 1 {
			var task string
			fmt.Println("Write Your Tasks")
			fmt.Scanln(&task)
			t = append(t, task)
		} else if tasks == 2 {
			var remove int
			fmt.Println("Which one You wanted to Remove Write its index")
			fmt.Scanln(&remove)
			slices.Delete(t, remove, remove+1)

		} else if tasks == 3 {
			for i := range t {
				fmt.Println(t[i])
			}
		} else if tasks == 4 {
			fmt.Println("Clearing Memory Exiting...")
			os.Exit(0)
		} else {
			fmt.Println("Invalid Info")
		}
	}
}
