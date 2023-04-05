package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/config"
)

func SetupAndListen() {
	// init
	router := gin.Default()

	// cors
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://172.19.0.1", "http://54.249.0.5", "http://54.249.0.5", "https://dco.tw", "http://localhost:8001"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Content-Length", "Accept-Language", "Accept-Encoding", "Connection", "Access-Control-Allow-Origin"},
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
	}))

	// short url
	router.GET("/api/r/:redirect", redirect)
	//router.Get("/urlShorten", getAllUrlShorten)
	router.GET("/api/urlShorten/:id", getUrlShorten)
	router.POST("/api/urlShorten", createUrlShorten)
	//router.Patch("/urlShorten", updateUrlShorten)
	//router.Delete("/urlShorten/:id", deleteUrlShorten)

	// google login
	//router.POST("/api/google/login", login)

	// log
	router.Use(gin.Logger())
	data, _ := json.MarshalIndent(router.Routes(), "", "  ")
	fmt.Println(string(data))

	router.Run(":" + config.Val.Port)
	log.WithFields(log.Fields{
		"Port": config.Val.Port,
	}).Info("serve started")
}
