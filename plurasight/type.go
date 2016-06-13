package main

import "fmt"

func main() {
	type Sample struct {
		name    string
		message string
	}
	s := Sample{"Reza", "Hello from USA"}
	fmt.Println(s.name, s.message)
}
