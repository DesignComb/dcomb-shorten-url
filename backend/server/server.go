package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/config"
	"main/middleware"
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

	api := router.Group("/api")

	// google login
	api.GET("ouath/google/url", access)
	api.GET("ouath/google/login", login)

	// short url
	api.GET("/r/:redirect", redirect)
	//api.Get("/urlShorten", getAllUrlShorten)

	// user
	userApi := api.Use(middleware.Auth())
	userApi.GET("/user/info", GetUserInfo)

	urlShortenApi := api.Use(middleware.Auth())
	urlShortenApi.GET("/urlShorten/:id", getUrlShorten)
	urlShortenApi.POST("/urlShorten", createUrlShorten)
	//api.Patch("/urlShorten", updateUrlShorten)
	//api.Delete("/urlShorten/:id", deleteUrlShorten)


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
