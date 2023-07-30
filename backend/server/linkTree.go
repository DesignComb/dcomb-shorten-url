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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "error could not retrieve tree from db " + err.Error()})
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
		processedLinks = append(processedLinks, processedLink)
	}
	res.Success(c, gin.H{
		"title":       tree.Title,
		"cover":       coverImageUrl,
		"description": tree.Description,
		"linkTree":    processedLinks,
	})
}

type RequestData struct {
	Title       string       `json:"title"`
	ImageID     uint64       `json:"imageId"`
	Description string       `json:"description"`
	LinkTree    []model.Link `json:"linkTree"`
}

func createLinkTree(c *gin.Context) {
	var requestData RequestData
	var tree model.Tree
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error parsing JSON " + err.Error()})
		return
	}

	// logged in user
	loginUserId, _ := c.Get("user_id")
	if loginUserId, ok := loginUserId.(string); ok {
		tree.UserId, _ = strconv.ParseUint(loginUserId, 10, 64)
	}
	tree.Title = requestData.Title
	tree.Description = requestData.Description
	tree.ImageId = requestData.ImageID

	// todo：db transaction
	// todo：foreign key constraint
	tree, err = model.CreateTree(tree)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not create tree in db " + err.Error()})
		return
	}

	links, err := model.CreateLinks(requestData.LinkTree)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not create links in db " + err.Error()})
		return
	}

	res.Success(c, gin.H{
		"id":          tree.ID,
		"title":       tree.Title,
		"description": tree.Description,
		"imageId":     tree.ImageId,
		"linkTree":    links,
	})
}

func updateLinkTree(c *gin.Context) {
	var requestData RequestData
	var updateTree model.Tree

	// todo: 有傳的欄位才更新，目前會把沒傳的清空
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not parse json " + err.Error()})
		return
	}

	// logged in user
	loginUserId, _ := c.Get("user_id")
	if loginUserId, ok := loginUserId.(string); ok {
		updateTree.UserId, _ = strconv.ParseUint(loginUserId, 10, 64)
	}

	treeId, err := strconv.ParseUint(c.Param("treeId"), 10, 64)
	tree, err := model.FindTree(treeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not find tree in DB " + err.Error()})
		return
	}
	updateTree.ID = tree.ID
	updateTree.Title = requestData.Title
	updateTree.Description = requestData.Description
	updateTree.ImageId = requestData.ImageID
	if updateTree.UserId != tree.UserId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "no permission"})
		return
	}

	err = model.UpdateTree(updateTree)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not update link tree in DB " + err.Error()})
		return
	}

	err = model.DeleteTreeLinks(treeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not delete links in db " + err.Error()})
		return
	}

	links, err := model.CreateLinks(requestData.LinkTree)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not create links in db " + err.Error()})
		return
	}

	res.Success(c, gin.H{
		"id":          tree.ID,
		"title":       tree.Title,
		"description": tree.Description,
		"imageId":     tree.ImageId,
		"linkTree":    links,
	})
}

func deleteLinkTree(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("treeId"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not parse treeId from url " + err.Error()})
		return
	}

	treeId, err := strconv.ParseUint(c.Param("treeId"), 10, 64)
	tree, err := model.FindTree(treeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not find tree in DB. " + err.Error()})
		return
	}

	// logged in user
	loginUserId, _ := c.Get("user_id")
	loginUserId, _  = strconv.ParseUint(fmt.Sprintf("%v", loginUserId), 10, 64)
	if loginUserId != tree.UserId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "no permission"})
		return
	}

	err = model.DeleteTree(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not delete from db " + err.Error()})
		return
	}

	res.Success(c, nil)
}