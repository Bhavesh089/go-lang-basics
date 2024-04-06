package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://api.agify.io?name=meelad"

func main() {
	fmt.Println("WElcome tho handling urls in golang")
	fmt.Println("my url: ", myurl)

	//parsing url

	result, _ := url.Parse(myurl)

	fmt.Println("result", result.Scheme)
	fmt.Println("result", result.Host)
	fmt.Println("result", result.Path)
	fmt.Println("result", result.Port())
	fmt.Println("result", result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of query params are: %T\n", qparams)

	for _, value := range qparams {

		fmt.Println("Values", value)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "api.agify.io",
		RawQuery: "name=meelad",
	}

	fmt.Println("This is url: ", partsOfUrl.String())
}
