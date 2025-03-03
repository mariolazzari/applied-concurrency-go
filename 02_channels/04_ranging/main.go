package main

import (
	"fmt"
	"time"
)

var greetings = []string{"Ciao", "Hello", "Hola", "Bonjour", "Namaste"}

func main() {
	// create channel
	ch := make(chan string, 1)
	go greeter(ch)

	// sleep for 1 second
	time.Sleep(1 * time.Second)
	fmt.Println("Main ready")

	// range checks automatically for open channel
	for greeting := range ch {
		// greeting, ok := <-ch
		// check if channel open
		// if !ok {
		// 	return
		// }

		time.Sleep(500 * time.Millisecond)
		fmt.Println("Greet received", greeting)
	}

}

// send only channel
func greeter(ch chan<- string) {
	fmt.Println("Greeter ready")
	// send greet
	for _, g := range greetings {
		ch <- g
	}
	// close channel
	close(ch)
	fmt.Println("Greet completed")
}
