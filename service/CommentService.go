package service

import (
	"errors"
	"gin_start/config"
	"gin_start/repository"
)

//实现评论的创建功能，已认证的用户可以对文章发表评论。
//实现评论的读取功能，支持获取某篇文章的所有评论列表。

func CreateComment(comment *repository.Comment) error {
	if comment.PostID == 0 {
		return errors.New("文章id不能为空")
	}
	config.Db.Debug().Create(&comment)
	return nil
}

func FindComment(postId int) (repository.Comment, error) {
	var comment repository.Comment
	config.Db.Debug().Find(&comment)
	return comment, nil
}
