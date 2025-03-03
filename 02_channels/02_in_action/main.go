package main

import (
	"fmt"
	"time"
)

func main() {
	// create channel
	// ch := make(chan string) // create unbuffered channel (no size in params)
	ch := make(chan string, 1) // create buffered channel (size in params)
	go greeter(ch)             // completes immediately if buffered

	// sleep for 2 seconds
	time.Sleep(2 * time.Second)
	fmt.Println("Main ready")

	// receive greet
	greet := <-ch

	// sleep for 2 seconds
	time.Sleep(2 * time.Second)
	fmt.Println("Greet received")
	fmt.Println(greet)

}

// send only channel
func greeter(ch chan<- string) {
	fmt.Println("Greeter is ready & waiting for sending greet...")
	// send greet
	ch <- "Hello"
	fmt.Println("Greet completed")
}
