package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var sgfDir = os.Getenv("GO_SGF_DIR")
var server = os.Getenv("GO_SGF2DB_SERVER")

var allSgfs []string

func walkSgfs(path string, info os.FileInfo, err error) error {
	if filepath.Ext(path) == ".sgf" {
		allSgfs = append(allSgfs, path)
	}
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sendSgf(path string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Error sending sgf: %v", e)
		}
	}()
	f, err := os.Open(path)
	check(err)
	defer f.Close()

	r := bufio.NewReader(f)
	resp, err := http.Post(server,
		"application/x-go-sgf",
		r)
	check(err)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(io.LimitReader(resp.Body, 65536))
	check(err)
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("Server returned wrong status code: %v, Error: %v",
			resp.StatusCode, b))
	}
	return nil
}

func checkEnv() {
	if sgfDir == "" {
		panic("GO_SGF_DIR not set")
	}
	if server == "" {
		panic("GO_SGF2DB_SERVER not set")
	}
}

func main() {
	checkEnv()
	filepath.Walk(sgfDir, walkSgfs)
	for i, path := range allSgfs {
		if i > 10 {
			break
		}
		sendSgf(path)
	}
}
