package main

import "fmt"

func main() {
	fmt.Println("welcome to loops")

	days := []string{"sunday", "monday", "tuesday", "wednesday"}

	for index, day := range days {
		fmt.Printf("This index %v and %v day\n", index, day)
	}
	someValue := 1
	for someValue < 10 {
		fmt.Println("This is somevalue", someValue)
		someValue++
	}

	for d := 0; d < 10; d++ {

		if d == 2 {
			goto p
		}

		fmt.Println("this is d: ", d)
	}

p:

	fmt.Println("This is goto command")

}
