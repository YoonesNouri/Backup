package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, World!")

	buffer := make([]byte, 5) //پنج بایت پنج بایت میخونه، به ورودی و خروجی فانکشن ها توجه کن
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error:", err)
			break
		}
		fmt.Printf("Read %d bytes: %q\n", n, buffer[:n])//when the format of buffer[:n] is %q or %s it writes strings
														//when the format of buffer[:n] is %v or %d it writes bytes 
	}
}
