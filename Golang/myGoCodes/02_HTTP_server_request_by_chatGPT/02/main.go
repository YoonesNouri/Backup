package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// URL to make the GET request to
	url := "https://jsonplaceholder.typicode.com/posts/1"

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

	// Unmarshal the JSON response into a Post struct
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Access the individual fields of the Post struct
	fmt.Println("User ID:", post.UserID)
	fmt.Println("Post ID:", post.ID)
	fmt.Println("Title:", post.Title)
	fmt.Println("Body:", post.Body)
}
