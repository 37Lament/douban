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



func changePassword(ctx *gin.Context) {
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	//检验旧密码是否正确
	flag, err := service.IsPasswordCorrect(username, oldPassword)
	if err != nil {
		fmt.Println("judge password correct err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	if !flag {
		tool.RespErrorWithDate(ctx, "旧密码错误")
		return
	}
	if len(newPassword) < 6{
		tool.RespErrorWithDate(ctx,"密码应该大于6位")
		return
	}


	//修改新密码
	err = service.ChangePassword(username, newPassword)
	if err != nil {
		fmt.Println("change password err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	flag, err := service.IsPasswordCorrect(username, password)
	if err != nil {
		fmt.Println("judge password correct err: ", err)
		tool.RespInternalError(ctx)
		return
	}


	if !flag {
		tool.RespErrorWithDate(ctx, "密码错误")
		return
	}
	ctx.SetCookie("username", username, 600, "/", "", false, false)

	tool.RespSuccessfulWithDate(ctx,CreateToken(username))
}

func register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	timenow := time.Now().Format("2006-01-02 15:04:05")
	user := model.User{
		Username: username,
		Password: password,
		Time:timenow ,
	}

	flag, err := service.IsRepeatUsername(username)
	if err != nil {
		fmt.Println("judge repeat username err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	if flag {
		tool.RespErrorWithDate(ctx, "用户名已经存在")
		return
	}
	if len(password) < 6{
		tool.RespErrorWithDate(ctx,"密码应该大于6位")
		return
	}

	err = service.Register(user)
	if err != nil {
		fmt.Println("register err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

func showuser(ctx*gin.Context){


	peopleid:=ctx.Param("id")
	peopleid2,err := strconv.Atoi(peopleid)
	if err != nil {
		fmt.Println("show user err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	people,err:=service.SelectU(peopleid2)
	if err != nil {
		fmt.Println("Selectpuser: ", err)
		tool.RespInternalError(ctx)
		return
	}



	tool.RespSuccessfulWithDate(ctx,people)
}

func Txt(ctx*gin.Context)  {
	iUsername, _ :=ctx.Get("username")
	username := iUsername.(string)
	Txt1:=ctx.PostForm("txt")
	err:=service.ChangeTxt(Txt1,username)
	if err != nil {
		fmt.Println("change txt err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}
