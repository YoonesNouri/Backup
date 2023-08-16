package main

import (
	"fmt"
	"net/http"
)

//* type Handler interface {
//* ServeHTTP(ResponseWriter, *Request)
//* }

type myWebServerType bool

func (m myWebServerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello site!")
	fmt.Fprintf(w, "request is: %+v ", r) // %+v: Formats structs with field names.
}

func main() {
	var k myWebServerType
	http.ListenAndServe("localhost:8080", k)
}
