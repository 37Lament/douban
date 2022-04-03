package dao

import (
	"douban/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dB *gorm.DB
//连接服务器

func InitDB() {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/douban?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
	dB = db
	defer db.Close()
	dB.AutoMigrate(&model.User{})
	dB.AutoMigrate(&model.Movie{})
	dB.AutoMigrate(&model.People{})
	dB.AutoMigrate(&model.Comment{})
}
