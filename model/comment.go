package model

import "time"//一个评论的结构体
//影评的结构体

type Comment struct {
	Id          int
	Score       int
	Flag 		bool
	Likes   	int64
	Txt         string//文本
	Username    string//用户名
	CommentTime time.Time//评论时间
	MovieName string
}
