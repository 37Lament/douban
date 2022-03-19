package api

import (
	"douban/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)




//管理员鉴权
func Admin(ctx *gin.Context) {
	iUsername, flag := ctx.Get("username")
	username := iUsername.(string)
	if flag==false{
		tool.RespErrorWithDate(ctx,"请登录后再尝试")
		ctx.Abort()
	}
	if username !="lxl" {
			tool.RespErrorWithDate(ctx,"你不是管理员")
			ctx.Abort()
	}
}


//重定向


//检查是否宕机
func Ping(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK,gin.H{
		"info": "OK",
	})
}
