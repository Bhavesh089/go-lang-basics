package main

import "fmt"

func main() {

	fmt.Println("maps in golang")

	lang := make(map[string]string)

	lang["python"] = "py"

	lang["java script"] = "js"

	lang["go lang"] = "gl"

	fmt.Println("List of all languages", lang)

	//delete
	delete(lang, "python")

	fmt.Println("This is after deleting python", lang)

	//looping the maps

	for key, value := range lang {
		fmt.Printf("for key, v value is %v %s", key, value)
	}
}
