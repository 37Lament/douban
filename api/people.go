package api

import (
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)




func Createpeople(ctx*gin.Context) {
	name:=ctx.PostForm("name")
	birthdate:=ctx.PostForm("birthdate")
	birthplace:=ctx.PostForm("birthplace")
	job:=ctx.PostForm("job")
	iMDb:=ctx.PostForm("imdb")

	people1:=model.People{
		Name: name,
		Birthdate: birthdate,
		Birthplace: birthplace,
		Job: job,
		IMDb:iMDb,
	}
	err:= service.Createp(people1)
	if err != nil {
		fmt.Println("createPeople err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}

func showpeople(ctx*gin.Context){


	peopleid:=ctx.Param("id")
	peopleid2,err := strconv.Atoi(peopleid)
	if err != nil {
		fmt.Println("showpeople err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	people,err:=service.SelectP(peopleid2)
	if err != nil {
		fmt.Println("Selectpeople: ", err)
		tool.RespInternalError(ctx)
		return
	}



	tool.RespSuccessfulWithDate(ctx,people)
}



