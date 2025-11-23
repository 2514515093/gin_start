package api

import (
	"gin_start/repository"
	"gin_start/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post repository.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": ""})
}

func FindPostById(c *gin.Context) {
	postId := c.Query("postId")
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		// 处理转换错误，比如返回 400
		//c.JSON(http.StatusBadRequest, gin.H{"error": "postId 无效"})
		//return
		//不做处理
	}
	result, err := service.FindPostById(postIdInt)
	if err != nil {
		// 处理转换错误，比如返回 400
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func UpdatePost(c *gin.Context) {
	var post repository.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, exists := c.Get("user")
	if !exists {
		// user 不存在
		return
	}
	user := u.(*repository.User)
	rowsAffected, err := service.UpdatePost(&post, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rowsAffected})
}

func DeletePost(c *gin.Context) {

	postId := c.Param("postId")
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		// 处理转换错误，比如返回 400
		c.JSON(http.StatusBadRequest, gin.H{"error": "postId 无效"})
		return
	}

	u, exists := c.Get("user")
	if !exists {
		// user 不存在
		return
	}
	user := u.(*repository.User)
	rowsAffected, err := service.DelPost(postIdInt, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rowsAffected})
}
