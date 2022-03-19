package oauth

import (
	"douban/tool"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const ClientID = "02258769147d2a5fd211"
const Clientsecrets ="cea7bc57313314b56e3096eacca93bc450c3b546"


type Conf struct {
	ClientId     string // 对应: Client ID
	ClientSecret string // 对应: Client Secret
	RedirectUrl  string // 对应: callback URL
}

var conf = Conf{
	ClientId:     "02258769147d2a5fd211",
	ClientSecret: "cea7bc57313314b56e3096eacca93bc450c3b546",
	RedirectUrl:  "http://47.106.178.100/token",
}


type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段下面没用到
	Scope       string `json:"scope"`      // 这个字段下面也没用到
}

// 认证并获取用户信息
func Oauth(ctx*gin.Context) {

	var err error
	// 获取 code
	var code = ctx.Query("code")

	// 通过 code, 获取 token
	var tokenAuthUrl = GetTokenAuthUrl(code)
	var token *Token
	if token, err = GetToken(tokenAuthUrl); err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	}

	// 通过token，获取用户信息
	var userInfo map[string]interface{}
	if userInfo, err = GetUserInfo(token); err != nil {
		fmt.Println("获取用户信息失败，错误信息为:", err)
		tool.RespInternalError(ctx)
		return
	}

	//  将用户信息返回前端
	var userInfoBytes []byte
	if userInfoBytes, err = json.Marshal(userInfo); err != nil {
		fmt.Println("在将用户信息(map)转为用户信息([]byte)时发生错误，错误信息为:", err)
		tool.RespInternalError(ctx)
		return
	}
	ctx.Set("Content-Type", "application/json")
	 tool.RespSuccessfulWithDate(ctx,userInfoBytes);
}

// 通过code获取token认证url
func GetTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		conf.ClientId, conf.ClientSecret, code,
	)
}

// 获取 token
func GetToken(url string) (*Token, error) {

	// 形成请求
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}

	// 将响应体解析为 token，并返回
	var token Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

// 获取用户信息
func GetUserInfo(token *Token) (map[string]interface{}, error) {

	// 形成请求
	var userInfoUrl = "https://api.github.com/user"	// github用户信息获取接口
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	// 将响应的数据写入 userInfo 中，并返回
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}

