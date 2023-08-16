package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	// URL to make the GET request to
	url := "https://jsonplaceholder.typicode.com/posts/1" // An example API endpoint

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}

	// Make sure the response body is closed after reading from it
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Convert the response body to a string and print it
	fmt.Println(string(body))
}
