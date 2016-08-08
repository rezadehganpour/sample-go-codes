package main

import (
	"./switch-package"
	"fmt"
)

type Person struct {
	Name    string
	Message string
}

func main() {
	person := Person{"Reza", "Hello"}
	fmt.Println(switching.PrintWithSwitch(person))
}
