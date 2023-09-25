package main

import (
	"fmt"
)

func main() {
	server := NewApiServer(":3000")
	server.runServer()
	fmt.Println("Hello")
}

//
