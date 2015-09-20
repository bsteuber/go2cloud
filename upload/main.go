package main

import (
	"github.com/bsteuber/go2cloud/tools/env"
	"github.com/bsteuber/go2cloud/upload/worker"
)

func main() {
	config := worker.Config{
		Server: env.Get("SGF2DB_SERVER"),
		SGFDir: env.Get("SGF_DIR")}
	worker.Start(config)
}
