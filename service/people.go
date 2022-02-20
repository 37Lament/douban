package service

import (
	"douban/dao"
	"douban/model"
)

func Createp(people1 model.People) error {
	return dao.InsertP(people1)
}

func SelectP(peopleid int)([]model.People,error){
	return dao.QueryP(peopleid)
}