package api

import (
	"douban/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)
//查看登陆状态
func auth(ctx *gin.Context) {
	username, err := ctx.Cookie("username")//检查cookie
	if err != nil {
		tool.RespErrorWithDate(ctx, "请登陆后进行操作")
		ctx.Abort()
	}

	ctx.Set("username", username)
	ctx.Next()
}
func Admin(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("username")
	if  err != nil {
			value := cookie.Value
			if value !="lxl" {
				tool.RespErrorWithDate(ctx,"你不是管理员")
				ctx.Abort()
			}
	}
		tool.RespSuccessfulWithDate(ctx,"管理员登陆成功")
		ctx.Next()
}


func Ping(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK,gin.H{
		"info": "ping",
	})
}