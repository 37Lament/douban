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
		fmt.Println("create err: ", err)
		tool.RespErrorWithDate(ctx,"请检查输入是否正确")
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

func showmovie(ctx*gin.Context){
	movieid:=ctx.Param("id")
	movieid2,err := strconv.Atoi(movieid)
	if err != nil {
		fmt.Println("showmovie err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	movie1,err:=service.SelectMovie(movieid2)
	if err != nil {
		fmt.Println("SelectMovie: ", err)
		tool.RespErrorWithDate(ctx,"没有此电影")
		return
	}
	tool.RespSuccessfulWithDate(ctx,movie1)
}