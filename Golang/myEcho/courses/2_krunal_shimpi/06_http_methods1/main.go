package main

import (
	"fmt"
	"net/http"
)

//? type Handler interface {
//? ServeHTTP(ResponseWriter, *Request)
//? }

// type login int
// type welcome int

func myLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "using GET method")
	case "POST":
		fmt.Fprintln(w, "using POST method ")
	}
	fmt.Fprintln(w, "on login page")
}

func myWelcome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
            fmt.Fprintln(w, "using GET method")
        case "POST":
            fmt.Fprintln(w, "using POST method ")
        }
	fmt.Fprintln(w, "on welcome page")
}

func main() {
	http.HandleFunc("/login", myLogin)
	http.HandleFunc("/welcome", myWelcome)
	// http.Handle("/login", http.HandlerFunc(myLogin))
	// http.Handle("/welcome", http.HandlerFunc(myWelcome))
	// var i login
	// var j welcome
	// http.Handle("/login", i)
	// http.Handle("/welcome", j)
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
