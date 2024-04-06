package main

import "fmt"

func main() {

	fmt.Println("Welcome to functions in golang")
	greeter()

	greeterTwo()

	result := adder(3, 5)

	resultPro, myMessage := proAdder(2, 4, 5, 2, 41, 1)
	fmt.Println("This is pro result", resultPro, myMessage)
	fmt.Println("Result is: ", result)
}

func adder(val1 int, val2 int) int {

	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "this is the string"
}

func greeterTwo() {
	fmt.Println("Another method")
}

func greeter() {

	fmt.Println("Namastey from golang")
}
