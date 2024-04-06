package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to pointers")

	var point *int

	fmt.Println("This is the init point", point)

	number := 2

	var pointer = &number

	fmt.Println("This is the pointer of number", pointer)
	fmt.Println("This is the value of that pointer", *pointer)

	*pointer = *pointer + 2

	fmt.Println("new updater number: ", number)
}
