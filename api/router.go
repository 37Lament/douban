package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()

	engine.POST("/register", register) //注册
	engine.POST("/login", login)       //登陆
	engine.POST("ping",Ping)//检查网络连接
	userGroup := engine.Group("/user")
	{
		userGroup.Use(Ping)//检查网络连接
		userGroup.Use(auth)
		userGroup.POST("/password", changePassword) //修改密码
		userGroup.POST("/excit",)
	}
	viewGroup:=engine.Group("/view")//游客模式可使用的
	{
		viewGroup.Use(Ping)//检查网络连接
		viewGroup.Use()

	}
	commentGroup:=engine.Group("/comment")
	{
		commentGroup.Use(Ping)//检查网络连接
		commentGroup.Use(auth)
		commentGroup.POST("/", addComment)
	}
	adminGroup:=engine.Group("/admin")
	{
		adminGroup.Use(Admin)//检查是不是管理员
		adminGroup.POST("/cmovie",CreateMovie)
	}

	//adGroup:=engine.Group("ad")//广告，还没做
	engine.Run()
}
