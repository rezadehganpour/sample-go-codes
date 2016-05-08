// A better version of echo.go

package main

import (
	"fmt"
	"os"
)

func main() {
	s, temp := "", ""
	fmt.Println(os.Args[0])
	for _, value := range os.Args[1:] {
		s += temp + value
		temp = " "
	}
	fmt.Println(s)
}
