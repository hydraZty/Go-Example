package main

import (
	"fmt"
	"net/http"
)

func main() {
	if resp, err := http.Get("http://example.com/"); err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println(resp)
	}
}
