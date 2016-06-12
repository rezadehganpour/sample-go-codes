//Read from file
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	countLines(files)
}

func countLines(files []string) {
	fmt.Println("countLines")
	file := ""
	for _, file = range files {
		count := make(map[string]int)
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "This is the error: %v", err)
		} else {
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				count[scanner.Text()]++
			}
		}
		f.Close()
		for line, n := range count {
			if n > 1 {
				fmt.Printf("This is a line: %s, This is a number: %d, This is the file: %s\n", line, n, file)
			}
		}
	}
}
