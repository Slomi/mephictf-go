package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		if err := fetchAndPrint(url); err != nil {
			fmt.Fprintf(os.Stderr, "fetchAndPrint: %s\n", err.Error())
			os.Exit(1)
		}
	}
}

func fetchAndPrint(url string) error {
	rsp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("fetch: %s", err.Error())
	}
	fmt.Printf("%s: %s\n", url, rsp.Status)
	return nil
}
