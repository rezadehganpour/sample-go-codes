package main

import "fmt"

func main() {
	m, s := generate_message()
	fmt.Println(m, s)
}

func generate_message() (message string, alt string) {
	return "Hello", "Poll"
}
