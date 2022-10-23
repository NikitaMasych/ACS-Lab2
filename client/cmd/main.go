package main

import (
	"client/config"
	"client/sender"
)

func main() {
	sender.StartSending(config.ServerType, config.ServerHost+":"+config.ServerPort)
}
