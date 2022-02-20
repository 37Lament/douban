package model


type People struct {
	Id int
	Name string		`form:"name" json:"name" binding:"required"`
	Birthdate string `form:"birthdate" json:"birthdate" binding:"required"`
	Birthplace string  `form:"birthplace" json:"birthplace" binding:"required"`
	Job string			`form:"job" json:"job" binding:"required"`
	IMDb string			`form:"IMDb" json:"IMDb" binding:"required"`

}
