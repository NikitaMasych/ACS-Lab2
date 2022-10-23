package main

import (
	"server/config"
	"server/listener"
)

func main() {
	listener.StartListen(config.ServerType, config.ServerHost+":"+config.ServerPort)
}
