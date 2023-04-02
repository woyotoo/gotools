package main

import (
	"log"
)

func SayHello(name string) string {
	return "Hello, " + name + "."
}

func main() {
	log.Println("hello, go tools")

	SayHello("World")
	SayHello("Golang")
}
