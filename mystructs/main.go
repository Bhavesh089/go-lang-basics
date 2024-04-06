package main

import "fmt"

func main() {
	fmt.Println("Welcome to the struct class")
	bhavesh := User{"bhavesh", "bhavesh@gmail.com", "Yes", "Male", 23}

	fmt.Println("This is bhavesh class", bhavesh)
	fmt.Printf("This is bhavesh Gender %v and Email %v \n", bhavesh.Gender, bhavesh.Email)
	fmt.Printf("This is bhavesh  %+v\n", bhavesh)

}

type User struct {
	Name   string
	Email  string
	Status string
	Gender string
	Age    int
}
