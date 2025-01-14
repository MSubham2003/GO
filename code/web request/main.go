package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//Get request
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	defer response.Body.Close()//Close the response body after the function ends close the network connection to free up system resources
	if err != nil {
		fmt.Println("Error in GET request", err)
	}
	fmt.Printf("Type of response is %T\n", response)

	//Read response content
	content, err1 := io.ReadAll(response.Body)
	if err1 != nil {
		fmt.Println("Error while reading response", err1)
	}
	fmt.Printf("Type of content is %T\n", content)
	fmt.Println(string(content))
}
