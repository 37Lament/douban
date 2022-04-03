package model

import "github.com/jinzhu/gorm"

type People struct {
	gorm.Model
	Name string		`gorm:"type:varchar(30)"`
	Birthdate string `gorm:"type:varchar(120)"`
	Birthplace string  `gorm:"type:varchar(120)"`
	Job string			`gorm:"type:varchar(120)"`
	IMDb string			`gorm:"type:varchar(120)"`

}
