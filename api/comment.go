package api

import (
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var flag3 = false

func addComment(ctx *gin.Context){
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	txt := ctx.PostForm("txt")
	postMovie:=ctx.PostForm("moviename")
	flag1:=ctx.PostForm("flag")
	flag2,err:=strconv.Atoi(flag1)
	if err != nil {
		fmt.Println("flag string to int err: ", err)
		tool.RespErrorWithDate(ctx, "传入flag有误")
		return
	}
	if flag2==1 {
		flag3 = true
	}
	comment := model.Comment{
		Flag: 		flag3,
		Txt:         txt,
		Username:    username,
		CommentTime: time.Now(),
		MovieName:  postMovie,
	}
	err = service.AddComment(comment)
	if err != nil {
		fmt.Println("add comment err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}
