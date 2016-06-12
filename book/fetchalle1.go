package main

import (
	"bufio"
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
		file, err := os.Create("fetchallurl.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Something went wrong. %v", err)
			os.Exit(1)
		}
		writer := bufio.NewWriter(file)
		os.Close("fetchallurl.txt")
		fmt.Fprintln(writer, "%s", <-ch)
		writer.Flush()
	}
	file, err := os.Create("fetchallurl.txt")
	writer := bufio.NewWriter(file)
	os.Close("fetchallurl.txt")
	elapse := time.Since(start).Seconds()
	fmt.Fprintln(writer, "%0.2fs", elapse)
	writer.Flush()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error: This is the url: %s, This is the error: %v", url, err)
		return
	}
	bytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error: Couldnt read the response. This is the error: %v", err)
		return
	}
	ch <- fmt.Sprintf("This is the bytes: %d, This is the elapse time: %0.2fs, and this is the url: %s", bytes, time.Since(start).Seconds(), url)
}
