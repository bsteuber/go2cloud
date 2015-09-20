package client

import (
	"fmt"
	"net/http"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func UploadSgf(server string, content string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Error sending sgf: %v", e)
		}
	}()
	resp, err := http.Post(server,
		"application/x-go-sgf",
		strings.NewReader(content))
	check(err)
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("Server returned wrong status code: %v",
			resp.StatusCode))
	}
	return nil
}
