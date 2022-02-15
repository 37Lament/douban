package model

//电影结构体
type Movie struct{
	Name         string  `form:"moviname" json:"moviname" binding:"required"`         //名字
	Director     string  `form:"director" json:"director" binding:"required"`         //导演
	Screenwriter string  `form:"screenwriter" json:"screenwriter" binding:"required"` //编剧
	Starring     string  `form:"starring" json:"starring" binding:"required"`         //主演明星
	Style        string  `form:"style" json:"style" binding:"required"`               //风格
	Area         string  `form:"area" json:"area" binding:"required"`                 //地区
	Language     string  `form:"language" json:"language" binding:"required"`         //语言
	ReleaseData  string  `form:"releasedata" json:"releasedata" binding:"required"`   //发行日期
	Length       int     `form:"length" json:"length" binding:"required"`             //时长
	IMDb         string  `form:"imdb" json:"imdb" binding:"required"`                 //IMDb码
	Score        float32 `form:"score" json:"score" binding:"required"`               //分数
}
