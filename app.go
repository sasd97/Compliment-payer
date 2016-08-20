package main

import (
	"github.com/sasd97/flatterer-vk/server"
)

func main() {
	appServer := server.Make()
	appServer.Run()
}
