package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var tinyurlAPI = "https://tinyurl.com/api-create.php?url="

func main() {
	if len(os.Args) < 2 {
		panic("url not provide")
	}
	targetURL := os.Args[1]
	res, err := http.Get(tinyurlAPI + targetURL)
	if err != nil {
		panic(err)
	}
	url, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(url))
}
