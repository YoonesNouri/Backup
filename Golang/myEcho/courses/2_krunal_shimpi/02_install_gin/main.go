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
	fmt.Fprintf(w, `
<html>
	<head>
    	<title>My Web Page</title>
	</head>

	<body>
    	<h1>Hello, World!</h1>
   	 	<p>This is a simple HTML template.</p>
	</body>
</html>
`) //! must use backticks for HTML
}

func main() {
	var k myWebServerType
	http.ListenAndServe("localhost:8080", k)
}
