package main

import (
	"fmt"
	"net/http"
)

//? type Handler interface {
//? ServeHTTP(ResponseWriter, *Request)
//? }

func myLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<html>
	<head>
    	login
	</head>

	<body>
    	<h1>this is login endpoint</h1>
	</body>
</html>
`)
}

func myWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<html>
	<head>
    	welcome
	</head>

	<body>
    	<h1>this is welcome endpoint</h1>
	</body>
</html>
`)
}

func main() {
	http.HandleFunc("/login", myLogin)
	http.HandleFunc("/welcome", myWelcome)
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
