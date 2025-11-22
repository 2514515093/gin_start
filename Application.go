package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run("0.0.0.0:8082") // 默认监听 0.0.0.0:8080
}

func CreateDb() (*gorm.DB, error) {
	dsn := "root:root123@tcp(211.159.169.85:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//创建表
	//db.AutoMigrate(&Student{})
	return db, err
}
