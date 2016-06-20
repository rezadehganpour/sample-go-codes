package main

import "fmt"

type Printer func(s string)

func main() {
	print_message(CustomPrinter("Please"))
}

func print_message(do Printer) {
	m, c := create_message()
	do(m)
	do(c)
}

func create_message() (message string, alternative string) {
	message = "Hello"
	alternative = "Bye"
	return
}

func CustomPrinter(s string) Printer {
	return func(custom string) {
		fmt.Println(custom, s)
	}
}
