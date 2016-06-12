package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "This url has a problem: %s, This is the error: %v", url, err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "This is the error: %v", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
		fmt.Printf("%s", strings.Split(resp.Status, " ")[0])
	}
}
