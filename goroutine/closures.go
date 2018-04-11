package main

import (
	"fmt"
	"time"
)

func outer(name string) {
	// variable in outer function
	text := "Modified " + name

	// foo is a inner function and has access to text variable, is a closure
	// closures have access to variables even after exiting this block
	foo := func() {
		fmt.Println(text)
	}
	text = "AB"
	// calling the closure
	go foo()
}

func main() {
	outer("hello")
	time.Sleep(2 * time.Second)
}
