package main

import (
	"gin_start/api"
	"gin_start/config"
	"gin_start/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitMysql()
	config.InitRedis()
	router := gin.Default()
	router.GET("/ping", handler(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/login", api.Login)
	router.POST("/register", api.Register)

	router.POST("/post/create", handler(), api.CreatePost)
	router.GET("/post/findById", api.FindPostById)
	router.PUT("/post/update", handler(), api.UpdatePost)
	router.DELETE("/post/delete/:postId", handler(), api.DeletePost)

	router.POST("/comment/create", handler(), api.CreateCm)
	router.GET("/comment/find", handler(), api.FindCm)

	router.Run("0.0.0.0:8082") // 默认监听 0.0.0.0:8080
}

func handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("sa-token")
		if token == "" {
			c.JSON(500, gin.H{"error": "请先进行登录"})
			c.Abort()
			return
		}
		user, err := utils.ParseTokenToUser(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		t, _ := utils.RdGet(strconv.FormatUint(uint64(user.ID), 10))
		if t == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "token已失效"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
