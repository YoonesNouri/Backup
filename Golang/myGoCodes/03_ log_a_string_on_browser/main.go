package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	message := "Hello, non-echo log!"
	fmt.Fprintln(w, message)
}

func main() {
	http.HandleFunc("/", handler)

	port := ":8000"
	fmt.Printf("Server started on http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
