package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/model"
	"main/pkg/res"
	"net/http"
	"strconv"
)

func getTree(c *gin.Context) {
	treeId, err := strconv.ParseUint(c.Param("treeId"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error could not parse treeId " + err.Error()})
		return
	}

	tree, err := model.FindTree(treeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error could not retrieve tree from db " + err.Error()})
		return
	}
	links, err := model.GetTreeAllLink(treeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error getting all links " + err.Error()})
		return
	}
	coverImageUrl := ""
	if tree.ImageId > 0 {
		image, _ := model.FindImage(tree.ImageId)
		coverImageUrl = s3ReadUrl + image.Uri
	}
	var processedLinks []gin.H

	for _, link := range links {
		linkImageUrl := ""
		if link.ImageId > 0 {
			image, _ := model.FindImage(link.ImageId)
			linkImageUrl = s3ReadUrl + image.Uri
		}
		processedLink := gin.H{
			"platformId":  link.PlatformId,
			"link":        link.Link,
			"title":       link.Title,
			"description": link.Description,
			"img":         linkImageUrl,
			"sort":        link.Sort,
			"isOnlyIcon":  link.IsOnlyIcon,
		}
		fmt.Println(processedLink)
		processedLinks = append(processedLinks, processedLink)
	}
	res.Success(c, gin.H{
		"title":       tree.Title,
		"cover":       coverImageUrl,
		"description": tree.Description,
		"linkTree":    processedLinks,
	})
}

func createLinkTree(c *gin.Context) {
	//var urlShorten model.UrlShorten
	//err := c.ShouldBindJSON(&urlShorten)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error parsing JSON " + err.Error()})
	//	return
	//}
	//
	//if len(urlShorten.Origin) == 0 {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "origin column is required"})
	//	return
	//}
	//
	//if urlShorten.Random {
	//	urlShorten.Short = utils.RandomURL(8)
	//} else {
	//	if len(urlShorten.Short) == 0 {
	//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "short column is required"})
	//		return
	//	}
	//}
	//
	//// logged in user
	//userId := c.Param("userId")
	//if len(userId) > 0 {
	//	loginUserId, _ := c.Get("user_id")
	//	if loginUserId != userId {
	//		res.Unauthorized(c, nil, "no permission")
	//		c.Abort()
	//		return
	//	}
	//	urlShorten.UserId, _ = strconv.ParseUint(userId, 10, 64)
	//	// todo：imageId 不存在或不屬於此user
	//} else {
	//	// 非會員無法加title等
	//	urlShorten.Title = ""
	//	urlShorten.Description = ""
	//	urlShorten.ImageId = 0
	//}
	//
	//urlShorten, err = model.CreateUrlShorten(urlShorten)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not create urlShorten in db " + err.Error()})
	//	return
	//}
	//
	//res.Success(c, urlShorten)
}
