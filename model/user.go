package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Id       sql.NullInt64
	Username string			`gorm:"type:varchar(30);unique_index"`
	Password string			`gorm:"type:varchar(120)"`
	Txt  string				`gorm:"type:varchar(120)"`
}
