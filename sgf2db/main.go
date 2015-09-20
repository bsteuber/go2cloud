package main

import (
	"github.com/bsteuber/go2cloud/sgf2db/server"
	"github.com/bsteuber/go2cloud/tools/env"
)

var port = env.Get("SGF2DB_PORT")

func main() {
	server.Run(port)
}
