package model

type User struct {
	Id       int			`form:"id" json:"id" binding:"required"`
	Username string			`form:"username" json:"username" binding:"required"`
	Password string
	Time string			`form:"time" json:"time" binding:"required"`
	Txt  string				`form:"txt" json:"txt" binding:"required"`
}
