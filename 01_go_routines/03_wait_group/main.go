package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // wait for 1 goroutine
	go hello(&wg)
	wg.Wait() // wait for all goroutines to finish (internal counter == 0)
	goodbye()
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the internal counter
	fmt.Println("Hello World")
}

func goodbye() {
	fmt.Println("Goodbye World")
}
