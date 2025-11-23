package service

import (
	"errors"
	"gin_start/config"
	"gin_start/repository"

	"gorm.io/gorm"
)

//实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
//实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
//实现文章的更新功能，只有文章的作者才能更新自己的文章。
//实现文章的删除功能，只有文章的作者才能删除自己的文章。

func CreatePost(post *repository.Post) error {
	if post.UserID == 0 {
		return errors.New("用户id不能为空")
	}
	config.Db.Debug().Create(&post)
	return nil
}

func FindPostById(postId int) ([]repository.Post, error) {
	var posts []repository.Post
	if postId == 0 {
		config.Db.Debug().Preload("Comments").Find(&posts)
	} else {
		config.Db.Debug().Find(&posts)
	}
	return posts, nil
}

func UpdatePost(post *repository.Post, user *repository.User) (int, error) {
	postdetail, _ := FindPostById(int(post.ID))
	if postdetail[0].UserID != user.ID {
		return 0, errors.New("只能更新自己的文章")
	}
	result := config.Db.Debug().Model(&repository.Post{}).Where("id = ?", post.ID).Updates((map[string]interface{}{
		"title":   post.Title,
		"content": post.Content,
	}))
	return int(result.RowsAffected), nil
}

func DelPost(postId int, user *repository.User) (int, error) {
	postdetail, _ := FindPostById(postId)
	if postdetail[0].UserID != user.ID {
		return 0, errors.New("只能删除自己的文章")
	}
	result := config.Db.Debug().Delete(&repository.Post{Model: gorm.Model{ID: uint(postId)}})
	return int(result.RowsAffected), nil
}
