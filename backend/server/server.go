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
		AllowOrigins:  []string{"http://172.19.0.1", "http://54.249.0.5", "http://54.249.0.5", "https://dco.tw", "http://localhost:8001", "http://localhost:3000"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Content-Length", "Accept-Language", "Accept-Encoding", "Connection", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Set-Cookie"},
		AllowMethods:  []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Set-Cookie"},

		AllowCredentials: true,
	}))

	api := router.Group("/api")

	// google login
	api.GET("ouath/google/url", access)
	api.GET("ouath/google/login", login)

	//region short url
	api.GET("/r/:redirect", redirect)
	//api.Get("/urlShorten", getAllUrlShorten)
	api.GET("/urlShorten/origin", getUrlShortenFromOrigin)
	api.POST("/urlShorten", createUrlShorten)
	api.GET("/urlShorten/search", searchNonUserUrlShorten)
	api.GET("/metadata", getMetadata)
	//api.Patch("/urlShorten", updateUrlShorten)
	//api.Delete("/urlShorten/:id", deleteUrlShorten)

	// user short url
	userUrlShortenApi := api.Use(middleware.Auth())
	userUrlShortenApi.POST("/user/:userId/urlShorten", createUrlShorten)
	userUrlShortenApi.GET("/user/:userId/urlShorten/origin", getUserUrlShortenFromOrigin)
	userUrlShortenApi.GET("/urlShorten/:id", getUserUrlShorten)
	userUrlShortenApi.GET("/user/:userId/urlShorten/search", searchUserUrlShorten)

	//endregion

	//region link tree




	//endregion

	// google login
	//router.POST("/api/google/login", login)

	// user
	userApi := api.Use(middleware.Auth())
	userApi.GET("/user/info", GetUserInfo)
	userApi.GET("/user/urlShorten", GetUserUrlShorten)

	// user image
	userImageApi := api.Use(middleware.Auth())
	userImageApi.POST("/user/:userId/image", UploadImage)
	userImageApi.GET("/user/:userId/image/:imageId", GetImage)

	// log
	router.Use(gin.Logger())
	data, _ := json.MarshalIndent(router.Routes(), "", "  ")
	fmt.Println(string(data))

	router.Run(":" + config.Val.Port)
	log.WithFields(log.Fields{
		"Port": config.Val.Port,
	}).Info("serve started")
}
