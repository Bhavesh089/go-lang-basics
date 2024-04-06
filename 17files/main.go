package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "this ned to the file"
	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	fmt.Println("Length is: ", length)

	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {

	databyte, err := ioutil.ReadFile(filename)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	fmt.Println("Text data inside the file", string(databyte))
}
func checkNilError(err error) {

	if err != nil {
		panic(err)
	}
}
