package api

import (
	"douban/dao"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Home(ctx*gin.Context){

	data1,err:=dao.Homep()
	if err != nil {
		fmt.Println("SelectMovie: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx,data1)
}
