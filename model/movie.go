package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

//电影结构体
type Movie struct{
	gorm.Model
	Movieid		sql.NullInt64  //id
	Name         string  `gorm:"type:varchar(120)"`         //名字
	Director     string  `gorm:"type:varchar(120)"`         //导演
	Screenwriter string  `gorm:"type:varchar(120)"` //编剧
	Starring     string  `gorm:"type:varchar(120)"`         //主演明星
	Style        string `gorm:"type:varchar(120)"`              //风格
	Area         string  `gorm:"type:varchar(120)"`                 //地区
	Language     string  `gorm:"type:varchar(120)"`         //语言
	ReleaseData  string  `gorm:"type:varchar(120)"`   //发行日期
	Length       int                 //时长
	IMDb         string  `gorm:"type:varchar(120)"`                 //IMDb码
	Score        float32                //分数
}
