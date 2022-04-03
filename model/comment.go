package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

//一个评论的结构体
//影评的结构体

type Comment struct {
	gorm.Model
	Postid 		int
	Score       sql.NullInt64
	Flag 		bool	`form:"flag" json:"flag" binding:"required"`
	Txt         string	`gorm:"type:varchar(120)"`//文本
	Username    string	`gorm:"type:varchar(30)"`//用户名
	CommentTime string	`gorm:"type:varchar(50)"`//评论时间
	MovieName string	`gorm:"type:varchar(120)"`
}
