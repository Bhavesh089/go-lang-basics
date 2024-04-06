package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://api.publicapis.org/entries"

func main() {
	fmt.Println("LCO web request")

	ressponse, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", ressponse)

	defer ressponse.Body.Close() //callers responsibilty to close the connections

	databytes, err := ioutil.ReadAll(ressponse.Body)

	if err != nil {
		fmt.Println(err)
	}
	content := string(databytes)

	fmt.Println("This is the content: ", content)
}
