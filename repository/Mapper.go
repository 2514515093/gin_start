package repository

import "gorm.io/gorm"

//设计数据库表结构，至少包含以下几个表：
//users 表：存储用户信息，包括 id 、 username 、 password 、 email 等字段。
//posts 表：存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
//comments 表：存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。
//使用 GORM 定义对应的 Go 模型结构体。

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Posts    []Post `gorm:"foreignKey:UserID"`
	salt     string
}

type Post struct {
	gorm.Model
	Title    string    `gorm:"not null"`
	Content  string    `gorm:"type:text;not null"`
	UserID   uint      `gorm:"not null"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	gorm.Model
	Content      string `gorm:"type:text;not null"`
	PostID       uint   `gorm:"not null"`
	SendUserId   uint
	ReciveUserId uint
}
