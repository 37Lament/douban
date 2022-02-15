package api

import (
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//
func CreateMovie(ctx*gin.Context) {
	movieName:=ctx.PostForm("moviename")
	director:=ctx.PostForm("director")
	screenwriter:=ctx.PostForm("screenwriter")
	starring:=ctx.PostForm("starring")
	style:=ctx.PostForm("style")
	area:=ctx.PostForm("area")
	language:=ctx.PostForm("language")
	length:=ctx.PostForm("length")
	iMDb:=ctx.PostForm("iMDb")
	releasedata:=ctx.PostForm("releasedata")
	length2,err := strconv.Atoi(length)
	if err != nil {
		fmt.Println("recharge rmb err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	movie1:=model.Movie{
		Name:         movieName,
		Director:     director,
		Screenwriter: screenwriter,
		Starring:     starring,
		Style:        style,
		Area:         area ,
		Language:     language,
		ReleaseData:  releasedata,
		Length:       length2,
		IMDb:         iMDb,
	}
	err = service.CreateMovie(movie1)
	if err != nil {
		fmt.Println("createMovie err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}
func Moviee(ctx*gin.Context){
	movie:=ctx.PostForm("moviename")
	movie1,err:=service.SelectMovie(movie)
	if err != nil {
		fmt.Println("SelectMovie: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx,movie1)
}