package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("fetch %v\n", err)
			return
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		//b, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Printf("fetch: reading %s,%v\n", url, err)
			return
		}
		//fmt.Printf("%s", b)
	}
}
