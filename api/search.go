package api

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Search(ctx*gin.Context){
	data:=ctx.PostForm("searchdata")
	data1,err:=service.Searchsth(data)
	if err != nil {
		fmt.Println("SelectMovie: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx,data1)
}
