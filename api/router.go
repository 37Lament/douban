package api

import (
	"douban/oauth"
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()

	engine.POST("/register", register) //注册
	engine.POST("/login", login)       //登陆

	userGroup := engine.Group("/user")
	{
		//userGroup.Use(auth)//检测登陆状态
		userGroup.Use(Getauth)
		userGroup.POST("/password", changePassword) //修改密码
		userGroup.POST("/changetxt",Txt)
	}
	viewGroup:=engine.Group("/view")//游客模式可使用的
	{
		viewGroup.GET("/homepage",Home)
		viewGroup.POST("/search",Search)
		viewGroup.POST("/")
		viewGroup.GET("/movie/:id",showmovie)
		viewGroup.GET("/people/:id",showpeople)
		viewGroup.GET("/user/:id",showuser)
		viewGroup.GET("/comment/:id",showCommenTool)
	}
	commentGroup:=engine.Group("/comment")
	{
		commentGroup.Use(Getauth)
		commentGroup.POST("/movie/:id", addComment)
		commentGroup.POST("/c/:id",addlevelComment)
		commentGroup.POST("/d/:id",deletComment)
	}
	adminGroup:=engine.Group("/admin")
	{
		adminGroup.Use(Getauth)
		adminGroup.Use(Admin)
		adminGroup.POST("/test")
		adminGroup.POST("/cmovie",CreateMovie)
		adminGroup.POST("/cpeople",Createpeople)
	}
	tokenGroup:=engine.Group("/token")
	{

		tokenGroup.GET("/",oauth.Oauth)
	}

	//adGroup:=engine.Group("ad")//广告，还没做
	engine.Run()
}
