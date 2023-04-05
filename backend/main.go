package main

import (
	"github.com/gin-gonic/gin"
	"main/config"
	"main/model"
	"main/server"
)

func main() {
	config.Init()
	gin.SetMode(config.Val.Mode)
	model.Setup()
	server.SetupAndListen()
}