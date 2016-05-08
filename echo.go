// This Go program will print the echo command that was passed to it
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, temp string
	fmt.Println(os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		s += temp + os.Args[i]
		temp = " "
	}
	fmt.Println(s)
}
