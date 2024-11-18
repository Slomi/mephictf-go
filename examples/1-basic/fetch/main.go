package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]
	errs := make(chan error, len(urls))
	for _, url := range urls {
		f := func(url string) {
			errs <- fetchAndPrint(url)
		}
		go f(url)
	}
	for range urls {
		err := <-errs
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
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
