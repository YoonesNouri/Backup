package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8) //هشت بایت هشت بایت میخونه
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) //when the format of b[:n] is %q or %s it writes strings
										  //when the format of b[:n] is %v or %d it writes bytes 
		if err == io.EOF {
			break
		}
	}
}
