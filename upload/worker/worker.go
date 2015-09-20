package worker

import (
	"fmt"
	"github.com/bsteuber/go2cloud/sgf2db/client"
	"io/ioutil"
	"os"
	"path/filepath"
)

var allSgfs []string

type Config struct {
	SGFDir string
	Server string
}

func (c Config) walkSgfs() {
	walkFn := func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".sgf" {
			allSgfs = append(allSgfs, path)
		}
		return nil
	}
	filepath.Walk(c.SGFDir, walkFn)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (c Config) upload(filename string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Error sending sgf: %v", e)
		}
	}()

	data, err := ioutil.ReadFile(filename)
	sgfContent := string(data)
	fmt.Printf("sgf:\n%v", sgfContent)
	check(err)
	err = client.UploadSgf(c.Server, sgfContent)
	check(err)
	return nil
}

func Start(c Config) {
	c.walkSgfs()
	for i, path := range allSgfs {
		if i > 10 {
			break
		}
		c.upload(path)
	}

}
