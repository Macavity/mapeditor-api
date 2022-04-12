package main

import (
	"Macavity/mapeditor-server/server"
)

var serverInstance = server.Server{}

func main() {
	serverInstance.Run()
}
