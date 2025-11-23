package config

import (
	"fmt"
	"gin_start/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitMysql() {
	dsn := "root:root123@tcp(211.159.169.85:3306)/gin_start?charset=utf8mb4&parseTime=True&loc=Local"
	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Db.AutoMigrate(&repository.User{}, &repository.Post{}, &repository.Comment{})
	fmt.Println("mysql初始化成功")
}
