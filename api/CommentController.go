package api

import (
	"gin_start/repository"
	"gin_start/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCm(c *gin.Context) {
	var comment repository.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.CreateComment(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, "")

}

func FindCm(c *gin.Context) {
	commentId := c.Query("commentId")
	postIdInt, err := strconv.Atoi(commentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "postId 无效"})
		return
	}
	comment, err := service.FindComment(postIdInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comment})
}
