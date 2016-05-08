// Print duplicates lines with thier coutn
package main

import (
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	for _, line := range os.Args[1:] {
		count[line]++
	}
	for n, line2 := range count {
		fmt.Println("This is line %s, this is count %d", line2, n)
	}
}
