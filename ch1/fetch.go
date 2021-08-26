// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	urlPrefix = "http://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, urlPrefix) {
			url = strings.Join([]string{urlPrefix, url}, "")
		}
		resp, getErr := http.Get(url)
		if getErr != nil {
			fmt.Fprintf(os.Stderr, "fetch: %V/n", getErr)
			os.Exit(1)
		}
		body, bodyErr := ioutil.ReadAll(resp.Body)
		status := resp.Status
		if bodyErr != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s:%V/n", url, bodyErr)
			os.Exit(1)
		}
		fmt.Printf("Status: %s\nContent:\n%s", status, body)
	}
}
