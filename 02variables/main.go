package main

import "fmt"

func main() {
	var username string = "bhavesh"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallFloat float32 = 255.37489273894723984
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//default values and some aliases

	var intergerValue int
	fmt.Println(intergerValue)
	fmt.Printf("variable is of type: %T \n", intergerValue)

	var stringValue string = "hi"
	fmt.Println(stringValue)
	fmt.Printf("variable is of type: %T \n", stringValue)

	//implicit type
	var randomValue = "hi"
	fmt.Println(randomValue)
	// randomValue = 3
	// fmt.Println(randomValue) bydefault randomvalue is string
}
