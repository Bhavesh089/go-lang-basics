package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("welcome to the conversions")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please give the rating")
	input, _ := reader.ReadString('\n')

	conversion, err := strconv.ParseInt(strings.TrimSpace(input), 0, 64)

	if err != nil {
		fmt.Println("hell we go the err", err)
	} else {
		println("Thanks for the rating, ", conversion+1)

	}

}
