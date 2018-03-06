package main

import "fmt"

func test(msg chan string, done chan<- bool) {
	fmt.Println("*****HI ABHI *****\n")
	temp := <-msg
	msg <- "Hello from a goroutine!"
	fmt.Println(temp)
	done <- true
	//msg <- "Hello from a goroutine!" //send a string through the channel
}

func main() {
	msg := make(chan string, 2) //channel (bidirectional channel) of type string
	done := make(chan bool)

	msg <- "hi whats up"
	fmt.Println("*****I am main go routine*****")

	go test(msg, done)
	<-done
	fmt.Println("Main routine waiting for message")

	fmt.Println(<-msg) //receive the string from the channel
	//fmt.Println(<-msg)
	//fmt.Println(message)

	//ch := make(chan string, 2) //buffered channel of size 2
	//ch <- "Hello"
	//ch <- "World!"
	//fmt.Println(<-ch, <-ch)
}
