package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error { //چون از وُید فانکشن یعنی فانکشن بدون ریترن استفاده کرده
	// در واقع انگار فقط فیلدهای استراکتِ «مای ارور» رو مقداردهی کرده
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
