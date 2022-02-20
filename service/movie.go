package service

import (
	"douban/dao"
	"douban/model"
)

func CreateMovie(movie model.Movie) error {
	return dao.InsertMovie(movie)
}

func SelectMovie(movieid int)(dao.Page,error){
	return dao.QueryMovie(movieid)
}

