package api

import (
	"douban/tool"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

//义生成结构体的私钥 key
var jwtKey = []byte("lxlxlxl666")
//Token 过期时间
var expireTime = time.Now().Add(24 * time.Hour).Unix()


type Token struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//生成token
func CreateToken(username string) string {
	jwtObj := Token{username , // 生成 token 的时候传值
		jwt.StandardClaims{
			ExpiresAt: expireTime, //过期时间
			Issuer:    "nibaxl",   // 签发人
		}}
	// 使用指定的签名方法创建签名对象
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtObj)
	// 使用指定的 secret 签名并获得完整的编码后的字符串 token
	tokenStr, _ := tokenObj.SignedString(jwtKey)
	return tokenStr
}

//验证token
func ParseToken(tokenString string) (*jwt.Token, *Token, error) {
	//创建Token实例
	token := &Token{}
	//验证Token
	jwtToken, err := jwt.ParseWithClaims(tokenString, token,
		func(token *jwt.Token) (i interface{}, err error) {
			return jwtKey, nil
		})
	return jwtToken, token, err
}

//获取用户id
func Getauth(ctx *gin.Context)  {
	tokenStr := ctx.Request.Header.Get("token")
	if tokenStr == "" {
		tool.RespErrorWithDate(ctx,"用户登录失效")
		ctx.Abort()
		return
	} else {
		token, mtoken, err := ParseToken(tokenStr)
		if err != nil || !token.Valid {
			tool.RespErrorWithDate(ctx,"验证失败")
			ctx.Abort()
			return
		}
		ctx.Set("username",mtoken.Username)
	}

}

