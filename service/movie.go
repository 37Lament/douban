package service

import (
	"douban/dao"
	"douban/model"
)

func CreateMovie(movie model.Movie) error {
	return dao.InsertMovie(movie)
}

func SelectMovie(movieid int)([]model.Movie,error){
	return dao.QueryMovie(movieid)
}

