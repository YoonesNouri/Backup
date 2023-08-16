package main

import (
	"fmt"
	"net/http"
)

type login int
type welcome int

func (l login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "on login page")
}

func (wl welcome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "on welcome page")
}

//? type Handler interface {
//? ServeHTTP(ResponseWriter, *Request)
//? }

func main() {
	// http.HandleFunc("/login", myLogin)
	// http.HandleFunc("/welcome", myWelcome)
	// http.Handle("/login", http.HandlerFunc(myLogin))
	// http.Handle("/welcome", http.HandlerFunc(myWelcome))
	var i login
	var j welcome
	http.Handle("/login", i)
	http.Handle("/welcome", j)
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
