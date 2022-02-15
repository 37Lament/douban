package service

import (
	"douban/dao"
	"douban/model"
)

func CreateMovie(movie model.Movie) error {
	return dao.InsertMovie(movie)
}

func SelectMovie(movieName string)([]model.Movie,error){
	return dao.QueryMovie(movieName)
}