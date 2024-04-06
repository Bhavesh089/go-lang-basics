package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	println("how many ratings that you can give: ")

	input, _ := reader.ReadString('\n')

	println("Thanks for the rating: ", input)

}
