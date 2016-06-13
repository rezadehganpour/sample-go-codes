package main

import "fmt"

func main() {
	var message string = "Hello World!"
	var message_pointer *string = &message

	fmt.Println("message: ", message)
	fmt.Println("message pointer: ", *message_pointer)
}
