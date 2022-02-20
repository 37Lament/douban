package api

import (
	"douban/dao"
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var flag3 = false


//进行评论
func addComment(ctx *gin.Context){
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	postid:=ctx.Param("id")
	postid2,err := strconv.Atoi(postid)
	if err != nil {
		fmt.Println("addComment err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	moviename1,err:=dao.SelectMoviebyId(postid2)
	if err!=nil {
		fmt.Println("SelectMoviebyId err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	score:=ctx.PostForm("score")
	score2,err:=strconv.Atoi(score)
	if err != nil {
		fmt.Println("addComment err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	txt := ctx.PostForm("txt")
	flag1:=ctx.PostForm("flag")
	flag2,err:=strconv.Atoi(flag1)
	if err != nil {
		fmt.Println("flag string to int err: ", err)
		tool.RespErrorWithDate(ctx, "传入flag有误")
		return
	}
	if flag2==1 {
		flag3 = true
	} else {flag3=false}
	timenow := time.Now().Format("2006-01-02 15:04:05")
	comment := model.Comment{
		Flag: 		flag3,
		Txt:         txt,
		Username:    username,
		CommentTime: timenow,
		MovieName:  moviename1.Name,
		Score: score2,
	}
	err = service.AddComment(comment)
	if err != nil {
		fmt.Println("add comment err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

//对评论进行评论
func addlevelComment(ctx *gin.Context)  {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	postid:=ctx.Param("id")
	postid2,err := strconv.Atoi(postid)
	if err != nil {
		fmt.Println("addComment err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	txt := ctx.PostForm("txt")
	timenow := time.Now().Format("2006-01-02 15:04:05")
	comment := model.Comment{
		Txt:         txt,
		Username:    username,
		CommentTime: timenow,
		Postid: postid2,
	}
	err = service.AddlevelComment(comment)
	if err != nil {
		fmt.Println("add comment err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}


//展示评论
func showCommenTool(ctx *gin.Context){
	postid:=ctx.Param("id")
	postid2,err := strconv.Atoi(postid)
	if err != nil {
		fmt.Println("addComment err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	data,err:=dao.ShowComment(postid2)
	if err!=nil {
		fmt.Println("showCommen err ",err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx,data)
}

//对评论进行删除
func deletComment(ctx*gin.Context)  {
	id:=ctx.Param("id")
	id2,err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("deletComment err: ", err)
		tool.RespErrorWithDate(ctx,"写入出错")
		return
	}
	err = dao.DeletC(id2)
	if err!=nil {
		fmt.Println("deletCommen err ",err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}