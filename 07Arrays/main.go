package main

import "fmt"

func main() {

	fmt.Println("welcome to arrays")

	var furitsList [4]string

	furitsList[0] = "Apple"
	furitsList[3] = "pineApple"
	furitsList[2] = "grapes"

	fmt.Println("this is fruit list", furitsList)
	fmt.Println("This is the length of the list", len(furitsList))

	var vegList = [3]string{"tomato", "potato", "califlower"}

	fmt.Println("This is the veglist", len(vegList))
}
