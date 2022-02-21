<a name="API"></a>
# API


<a name="gzwVa"></a>
# User
<a name="oIpNw"></a>
### `/user/password` POST`修改密码



| 请求参数 | 类型 | 说明 |
| --- | --- | --- |
| username | 必选 | 登录名 |
| old_password | 必选 | 旧密码 |
| new_password | 必选 | 新密码 |




| 返回参数 | 说明 |
| --- | --- |
| status | 状态码 |
| data | 返回消息 |

| status | data | 说明 |
| --- | --- | --- |
| `false` | `"旧密码不匹配"` | `不匹配`​ |
| `false` | `"密码应超过六位"` | <br />  `new_password`太短 |
| `true` | `"成功"` | `loginName`<br /> 与 `old_password`<br /> 匹配 |

<a name="dC9Nu"></a>
### `/user/changetxt` POST`修改签名
| 请求参数 | 类型 | 说明 |
| --- | --- | --- |
| txt | 必选 | 修改内容 |

| status | data | 说明 |
| --- | --- | --- |
| `false` | `"修改失败"` | 还没有修改失败过 |
| `true` | `"成功"` | 修改成功 |


<br />
<br />

<a name="3bf1f9ce"></a>
### `/login` POST`登陆



| 请求参数 | 类型 | 说明 |
| --- | --- | --- |
| username | 必选 | 登录名（手机/邮箱） |
| password | 必选 | 密码 |

<a name="gjZfj"></a>
### `/register` POST`注册


- `application/x-www-form-urlencoded`
- 密码登录
| 请求参数 | 类型 | 说明 |
| --- | --- | --- |
| username | 必选 | 登录名（手机/邮箱） |
| password | 必选 | 密码 |

| 返回参数 | 说明 |
| --- | --- |
| status | 状态码 |
| data | 返回消息 |

| status | data | 说明 |
| --- | --- | --- |
| `false` | `"用户名已存在"` | `用户名重复` |
| `false` | `"密码应该大于6位"` | `密码应该大于6位` |
| `true` | `"注册成功"` | `成功` |

<a name="hemXH"></a>
## 评论
<a name="UitJn"></a>
### `/comment/:id` POST`对电影id为id发送短评论/一级评论
| 请求参数 | 类型 | 说明 |
| --- | --- | --- |
| txt | 必选 | 评论内容 |
| flag | 必选 | 是否看过此电影 |
| score | 必选 | 对电影评分 |

| status | data | 说明 |
| --- | --- | --- |
| `false` | `"id格式不对"` | `id不是int类型` |
| `ture` | `"成功"` | 评论成功 |



<a name="zqKv5"></a>
### `/comment/c/:id ` POST`对评论id为：id的评论发送二级评论
| 请求参数 | 类型 | 说明 |
| --- | --- | --- |
| txt | 必选 | 评论内容 |

| status | data | 说明 |
| --- | --- | --- |
| `false` | `"id格式不对"` | `id不是int类型` |
| `ture` | `"成功"` | 评论成功 |

<a name="G9dSK"></a>
## `/comment/d/:id `POST 删除id为：id的评论
| status | data | 说明 |
| --- | --- | --- |
| `ture` | `"成功"` | 删除成功 |
| `false` | `"评论不存在"` | 评论不存在或者删除出错 |



<a name="gAssp"></a>
# 主页
<a name="hPhzO"></a>
### `/view/homepage`主页页面  GET 



| status | data | 说明 |
| --- | --- | --- |
| ture | moviename：“电影名字”<br />socore：“电影评分” | 主页 |



<a name="C92B6"></a>
### `/view/homepage`搜索  POST 
| 请求参数 | 类型 | 说明 |
| --- | --- | --- |
| searchdata | 必选 | 搜索内容 |

| status | data | 说明 |
| --- | --- | --- |
| ture | "People": [<br />            {<br />                "name": "演员姓名<br />                "birthdate": "演员出生年月日",<br />                "birthplace": "出生日期",<br />                "job": "演员职业"，<br />            }]<br />"Movie": [<br />            {<br />                "moviename": "电影",<br />                "director": "导演",<br />                "screenwriter": "编剧",<br />                "starring": " 主演",<br />                "style": "风格",<br />                "area": "发行地区",<br />                "language": "汉语普通话",<br />                "length": 电影时间长度,<br />                "score": 评分<br />            }<br />​<br /> | 两个结构体集合，一个演员，一个电影 |



<a name="cQ0SY"></a>
## `/view/movie/:id`展示电影主页 GET
| status | data | 说明 |
| --- | --- | --- |
| ture | "Movie": [<br />            {<br />                "moviename": "电影名字",<br />                "director": "导演",<br />                "screenwriter": "编剧",<br />                "starring": "主演",<br />                "style": "类型",<br />                "area": "地区",<br />                "language": "语言",<br />                "releasedata": "发行日期,<br />                "length": 时间长度,<br />                "imdb": "IMDB号",<br />                "score": 评分<br />            }<br />        ],<br />        "Commentsid": [<br />            5 //影评id合集<br />        ] | 电影的详情与电影下评论的id集合<br />（利用id集合可以将全部评论展示） |

<a name="vDBDS"></a>
## `/view/people/:id`展示人物主页 GET
| status | data | 说明 |
| --- | --- | --- |
| ture |  {<br />                "name": "演员姓名<br />                "birthdate": "演员出生年月日",<br />                "birthplace": "出生日期",<br />                "job": "演员职业"，<br />"IMDb":"IMDB号"<br />            } | 演员信息 |

<a name="PQh8i"></a>
## `/view/user/:id`展示用户主页 GET
| status | data | 说明 |
| --- | --- | --- |
| ture | {<br />            "Id": ID,<br />            "Username": "用户名",<br />            "Time": "注册时间",<br />            "Txt": "签名"<br />        } | 主页 |



<a name="KlupK"></a>
## `/view/comment/:id `展示评论 GET
| status | data | 说明 |
| --- | --- | --- |
| ture | {<br />        "Fcomment"//这是一级评论: {<br />            "Username": "评论者用户名",<br />            "Moviename": 评论影评",<br />            "Txt": "评论文字",<br />            "CommentTime": "评论时间",<br />            "Flag": bool值，表示是否**看过此电影**<br />        },<br />        "Coment2s"//二级评论: [<br />            {<br />                "Username": "评论者用户名",<br />                "Txt": "文本内容",<br />                "Commenttime": "评论时间"<br />            }<br />        ]<br />    } | 对一级影评以及电影上id集合的评论展示 |

<a name="QOK1L"></a>
## 管理员功能
<a name="pA6Uj"></a>
##  /admin/cmovie POST 创建电影
| 请求参数 | 类型 | 说明 |
| --- | --- | --- |
| 电影结构体 | 必选 | 先放着 |



<a name="MeTIH"></a>
##  /admin/cpeople POST 创建人物
| 请求参数 | 类型 | 说明 |
| --- | --- | --- |
| 人物结构体 | 必选 | 先放着 |



<a name="uPUfx"></a>
# 小型中间件
<a name="Da993"></a>
## /ping POST 检测是否宕机
| status | data | 说明 |
| --- | --- | --- |
| ture | 成功 | 没有宕机 |

<a name="f2ea87de"></a>
# 一般规定

<br />如无特殊说明，则返回一个以下格式的 json：<br />

```javascript
{
    status: true, // true：成功， false：失败
    data: "" // 提示信息
}
```
