package model

//一个评论的结构体
//影评的结构体

type Comment struct {
	Id          int
	Postid 		int
	Score       int   `form:"score" json:"score" binding:"required"`
	Flag 		bool	`form:"flag" json:"flag" binding:"required"`
	Txt         string	`form:"txt" json:"txt" binding:"required"`//文本
	Username    string	`form:"username" json:"username" binding:"username"`//用户名
	CommentTime string	`form:"commenttime" json:"commenttime" binding:"required"`//评论时间
	MovieName string	`form:"moviename" json:"moviename" binding:"required"`
}
