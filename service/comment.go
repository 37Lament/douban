package service

import (
	"douban/dao"
	"douban/model"
)

func AddComment(comment model.Comment) error {
	return dao.InsertComment(comment)
}

func AddlevelComment(comment model.Comment) error {
	return dao.InsertlevelComment(comment)
}
