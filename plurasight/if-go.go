package main

import "fmt"

func main() {
	var value bool = false
	PrintSomethingGood(value)
}

func PrintSomethingGood(value bool) {
	if value {
		fmt.Println("I am sexy and I know it!")
	} else {
		fmt.Println("Love me like you do.")
	}
}
