package main

import (
	"main/model"
	"main/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}