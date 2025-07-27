package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Printf("Hello, %s\n", name)
}

func main() {
	go sayHello("Minna")
	go sayHello("Anu")

	time.Sleep(1 * time.Second)
}