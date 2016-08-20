package main

import (
	"github.com/sasd97/flatterer/server"
)

func main() {
	appServer := server.Make()
	appServer.Run()
}
