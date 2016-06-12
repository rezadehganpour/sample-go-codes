package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("This is the url: %s, This is the error: %v", err)
		return
	}
	bytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Could not read the file: $v", err)
		return
	}
	elapse := time.Since(start).Seconds()
	ch <- fmt.Sprintf("This is bytes: %d, and thats how much time it spends: %.2fs with this url: %s", bytes, elapse, url)
}
