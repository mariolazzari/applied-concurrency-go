package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	// never use time.Sleep in production code
	time.Sleep(1 * time.Second)
	goodbye()
}

func hello() {
	fmt.Println("Hello World")
}

func goodbye() {
	fmt.Println("Goodbye World")
}
