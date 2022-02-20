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
		//userGroup.Use(auth)//检测登陆状态
		userGroup.Use(Getauth)
		userGroup.POST("/password", changePassword) //修改密码
		userGroup.POST("/changetxt",Txt)
	}
	viewGroup:=engine.Group("/view")//游客模式可使用的
	{
		viewGroup.GET("/homepage",Home)
		viewGroup.Use(Ping)//检查网络连接
		viewGroup.Use(point)
		viewGroup.POST("/search",Search)
		viewGroup.POST("/")
		viewGroup.GET("/movie/:id",showmovie)
		viewGroup.GET("/people/:id",showpeople)
		viewGroup.GET("/user/:id",showuser)
		viewGroup.GET("/comment/:id",showCommenTool)
	}
	commentGroup:=engine.Group("/comment")
	{
		commentGroup.Use(Ping)//检查网络连接
		commentGroup.Use(Getauth)
		commentGroup.POST("/movie/:id", addComment)
		commentGroup.POST("/c/:id",addlevelComment)
		commentGroup.POST("/d/:id",deletComment)
	}
	adminGroup:=engine.Group("/admin")
	{
		adminGroup.Use(Ping)
		adminGroup.Use(Getauth)
		adminGroup.Use(Admin)
		adminGroup.POST("/test")
		adminGroup.POST("/cmovie",CreateMovie)
		adminGroup.POST("/cpeople",Createpeople)
	}

	//adGroup:=engine.Group("ad")//广告，还没做
	engine.Run()
}
