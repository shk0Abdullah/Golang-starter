package main

import "fmt"
import "strings"


func main(){
	// dict := map[string]int{
	// 	"r" : 1,
	// 	"p" : 2,
	// 	"s" : 3, 
	// }
	var inp int 
	fmt.Println("---MY Rock Paper Scissor Game--- \n Press 0 For Exit \n 1. Player vs Player \n 2. Player vs Computer")
	fmt.Scanln(&inp)
	if inp == 1{
		playervsplayer()
	}
	// else if inp == 0 {
	// 	...
	// }
	// else{
	// 	...
	// }

}
func playervsplayer(){
			var p1 string
			var inp1, inp2 int  
			fmt.Println("Enter the Player1's input: ")
			fmt.Scanln(&p1)
			if strings.ToLower(p1) == "rock" || strings.ToLower(p1) == "r"{
					inp1 = 1
			}else if strings.ToLower(p1) == "paper" || strings.ToLower(p1) == "p"{
					inp1 = 2
			}else if strings.ToLower(p1) == "scissor" || strings.ToLower(p1) == "s"{
					inp1 = 3
			}else{
					fmt.Println("Error")
					// Exit the program
			}
			var p2 string 
			fmt.Println("Enter Player2's choice:")
			fmt.Scanln(&p2)
			if strings.ToLower(p2) == "rock" || strings.ToLower(p2) == "r"{
				inp2 = 1
			}else if strings.ToLower(p2) == "paper" || strings.ToLower(p2) == "p"{
				inp2 = 2
			}else if strings.ToLower(p2) == "scissor" || strings.ToLower(p2) == "s"{
				inp2 = 3
			}else{
				fmt.Println("Error")
				// Exit the program
			}
			evaluate(inp1, inp2)
		
		}

func evaluate (p1 int, p2 int ){

	if p1 == 1 && p2 == 2 {
		fmt.Println("Player2 Wins!")
	}else if p1 ==1 && p2 == 3 {
		fmt.Println("Player1 Wins!")
	}else if p1 ==1 && p2 == 1 {
		fmt.Println("Draw!")
	}else if p1 ==2 && p2 == 2 {
		fmt.Println("Draw!")
	}else if p1 ==3 && p2 == 3 {
		fmt.Println("Draw!")
	}else if p1 ==2 && p2 == 1 {
		fmt.Println("Player1 Wins!")
	}else if p1 ==2 && p2 == 3 {
		fmt.Println("Player2 Wins!")
	}else if p1 ==3 && p2 == 1 {
		fmt.Println("Player2 Wins!")
	}else if p1 ==3 && p2 == 2 {
		fmt.Println("Player1 Wins!")
	}else{
		fmt.Println("Fucking Go")
	}
}

