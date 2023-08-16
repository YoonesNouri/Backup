package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()   // increment is of type 'func() int' that holds the closure returned by the wrapper
	fmt.Println(increment()) // That's why it can be called
	fmt.Println(increment())
	fmt.Printf("%#v", increment)
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
