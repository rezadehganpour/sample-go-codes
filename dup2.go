// Read input from user

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		count[scanner.Text()]++
	}
	for line, n := range count {
		fmt.Println("This is the count: %d, this is the line %s", n, line)
	}
}
