package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	resp, err := http.Post("http://192.168.178.61:8080/upload",
		"application/x-go-sgf",
		strings.NewReader("Foorg"))
	check(err)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(string(b))
}
