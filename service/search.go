package service

import (
	"douban/dao"
	"douban/model"
)

func Searchsth(data string)([]model.Movie,error){
	return dao.Searching(data)
}