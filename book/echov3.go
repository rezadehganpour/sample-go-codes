// This is the best solution for the echo practice

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:], " ")
}
