package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/model"
	"main/utils"
	"net/http"
	"strconv"
)

func redirect(c *gin.Context) {
	shortenUrl := c.Param("redirect")
	url, err := model.FindByUrl(shortenUrl)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not find url in DB " + err.Error()})
		return
	}
	// grab any stats you want...
	url.Clicked += 1
	err = model.UpdateUrlShorten(url)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{"origin": url.Origin})
}

func getAllUrlShorten(c *gin.Context) {
	urls, err := model.GetAllUrlShorten()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error getting all goly links " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data:": urls})
}

func getUrlShorten(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error could not parse id " + err.Error()})
		return
	}

	url, err := model.GetUrlShorten(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error could not retrieve url from db " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data:": url})
}

func createUrlShorten(c *gin.Context) {
	acceptHeader := c.Request.Header.Get("Accept")
	if acceptHeader != "" && acceptHeader != "application/json" {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"error": "Not Acceptable",
		})
		return
	}

	var urlShorten model.UrlShorten
	err := c.ShouldBindJSON(&urlShorten)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error parsing JSON " + err.Error()})
		return
	}

	if len(urlShorten.Origin) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "origin column is required"})
		return
	}

	if urlShorten.Random {
		urlShorten.Short = utils.RandomURL(8)
	} else {
		if len(urlShorten.Short) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "short column is required"})
			return
		}
	}

	urlShorten, err = model.CreateUrlShorten(urlShorten)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not create urlShorten in db " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data:": urlShorten})
}

func updateUrlShorten(c *gin.Context) {
	acceptHeader := c.Request.Header.Get("Accept")
	if acceptHeader != "" && acceptHeader != "application/json" {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"error": "Not Acceptable",
		})
		return
	}

	var urlShorten model.UrlShorten

	err := c.ShouldBindJSON(&urlShorten)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not parse json " + err.Error()})
		return
	}

	err = model.UpdateUrlShorten(urlShorten)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not update urlShorten link in DB " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data:": urlShorten})
}

func deleteUrlShorten(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not parse id from url " + err.Error()})
		return
	}

	err = model.DeleteUrlShorten(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not delete from db " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "urlShorten deleted."})
}

func SetupAndListen() {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://172.19.0.1", "http://54.249.0.5", "http://54.249.0.5", "https://dco.tw", "http://localhost:8001"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Content-Length", "Accept-Language", "Accept-Encoding", "Connection", "Access-Control-Allow-Origin"},
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
	}))
	
	router.GET("/api/r/:redirect", redirect)

	//router.Get("/urlShorten", getAllUrlShorten)
	router.GET("/api/urlShorten/:id", getUrlShorten)
	router.POST("/api/urlShorten", createUrlShorten)
	//router.Patch("/urlShorten", updateUrlShorten)
	//router.Delete("/urlShorten/:id", deleteUrlShorten)

	router.Use(gin.Logger())
	data, _ := json.MarshalIndent(router.Routes(), "", "  ")
	fmt.Println(string(data))

	router.Run(":5000")
}
