package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("dice number is :", diceNumber)
	case 2:
		fmt.Println("dice number is :", diceNumber)
	case 3:
		fmt.Println("dice number is :", diceNumber)
	case 4:
		fmt.Println("dice number is :", diceNumber)

	case 5:
		fmt.Println("dice number is :", diceNumber)

	case 6:
		fmt.Println("dice number is :", diceNumber)

	default:
		fmt.Println("What was that????")
	}
}
