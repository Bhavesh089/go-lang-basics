package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to the time study")

	presentTime := time.Now()

	fmt.Println("This is present time", presentTime)
	fmt.Println("Format time", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2024, time.March, 30, 18, 30, 0, 0, time.UTC)
	fmt.Println(createdDate)

}
