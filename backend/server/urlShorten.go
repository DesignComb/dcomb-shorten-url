package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/model"
	"main/pkg/res"
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

	res.Success(c, gin.H{
		"origin":      url.Origin,
		"title":       url.Title,
		"description": url.Description,
		"image_id":    url.ImageId,
	})
}

func getAllUrlShorten(c *gin.Context) {
	urls, err := model.GetAllUrlShorten()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error getting all goly links " + err.Error()})
		return
	}
	res.Success(c, urls)
}

func getUserUrlShorten(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error could not parse id " + err.Error()})
		return
	}

	value, exists := c.Get("user_id")
	if exists {
		str, ok := value.(string)
		if ok {
			loginUserId, _ := strconv.ParseUint(str, 10, 64)
			url, err := model.GetUserUrlShorten(id, loginUserId)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error could not retrieve url from db " + err.Error()})
				return
			}
			res.Success(c, url)
		}
	} else {
		res.Unauthorized(c, nil, "not logged in")
		c.Abort()
	}
}

func getUrlShortenFromOrigin(c *gin.Context) {
	origin := c.Query("origin")
	if !(len(origin) > 0) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "origin is empty."})
		return
	}
	url := model.GetNonUserUrlShortenFromOrigin(origin)
	res.Success(c, url)
}

func getUserUrlShortenFromOrigin(c *gin.Context) {
	origin := c.Query("origin")
	if !(len(origin) > 0) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "origin is required."})
		return
	}

	userId := c.Param("userId")
	var loginUserId uint64
	if len(userId) > 0 {
		value, exists := c.Get("user_id")
		if value != userId {
			res.Unauthorized(c, nil, "no permission")
			c.Abort()
			return
		}
		if exists {
			str, ok := value.(string)
			if ok {
				loginUserId, _ = strconv.ParseUint(str, 10, 64)
			}
		}
	}
	userUrlShorten := model.GetUserUrlShortenFromOrigin(origin, loginUserId)
	nonUserUrlShorten := model.GetNonUserUrlShortenFromOrigin(origin)

	res.Success(c, gin.H{
		"userUrl":    userUrlShorten,
		"nonUserUrl": nonUserUrlShorten,
	})
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

	// logged in user
	userId := c.Param("userId")
	if len(userId) > 0 {
		loginUserId, _ := c.Get("user_id")
		if loginUserId != userId {
			res.Unauthorized(c, nil, "no permission")
			c.Abort()
			return
		}
		urlShorten.UserId, _ = strconv.ParseUint(userId, 10, 64)
		// todo：imageId 不存在或不屬於此user
	} else {
		// 非會員無法加title等
		urlShorten.Title = ""
		urlShorten.Description = ""
		urlShorten.ImageId = 0
	}

	urlShorten, err = model.CreateUrlShorten(urlShorten)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not create urlShorten in db " + err.Error()})
		return
	}

	res.Success(c, urlShorten)
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

	res.Success(c, urlShorten)
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

	res.Success(c, nil)
}

func searchUserUrlShorten(c *gin.Context) {
	keyword := c.Query("keyword")
	if len(keyword) < 3 {
		res.BadRequest(c, nil, "keyword is required")
		c.Abort()
		return
	}

	userId := c.Param("userId")
	var loginUserId uint64
	if len(userId) > 0 {
		value, exists := c.Get("user_id")
		if value != userId {
			res.Unauthorized(c, nil, "no permission")
			c.Abort()
			return
		}
		if exists {
			str, ok := value.(string)
			if ok {
				loginUserId, _ = strconv.ParseUint(str, 10, 64)
			}
		}
	}

	userUrlShorten := model.SearchUserUrlShorten(loginUserId, keyword)
	nonUserUrlShorten := model.SearchNonUserUrlShorten(keyword)

	res.Success(c, gin.H{
		"userUrl":    userUrlShorten,
		"nonUserUrl": nonUserUrlShorten,
	})
}

func searchNonUserUrlShorten(c *gin.Context) {
	keyword := c.Query("keyword")
	if len(keyword) < 3 {
		res.BadRequest(c, nil, "keyword is required")
		c.Abort()
		return
	}
	url := model.SearchNonUserUrlShorten(keyword)
	res.Success(c, url)
}
