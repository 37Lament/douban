package service

import (
	"douban/dao"
)

func Searchsth(data string) ( dao.Data1, error) {
	return dao.Searching(data)
}

