package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file in args")
		os.Exit(1)
	}

	urlsByte, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	urls := strings.Split((string)(urlsByte), "\r\n")

	for _, url := range urls {
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("%s\terror\n", url)
			continue
		}
		fmt.Printf("%s\t%s\n", url, response.Status)
	}
}
