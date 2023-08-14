package main

import "fmt"

func main() {
	server := NewAPISever(":3000")
	server.Run()
	fmt.Println("Hello")
}
