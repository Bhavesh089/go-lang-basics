package main

import "fmt"

func main() {
	fmt.Println("Welcome to the if else statements")

	score := 29

	var result string

	if score > 20 {
		result = "You are the winner of the match"
	} else if score < 20 {
		result = "Match is still on going"
	} else {
		result = "scores are tied"
	}

	fmt.Println(result)
}
