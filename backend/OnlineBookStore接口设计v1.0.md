<h2 align='center'>
	OnlineBookStore接口设计v1.0
</h2>

> 主题：接口设计
>
> 作者：章星明
>
> 时间：2020年4月18日

【TOC】：

[toc]

---

*用例请参考下图：*

![用例](http://tvax3.sinaimg.cn/large/006VDfrXly1gdyaycxzz2j31fx1lbwn7.jpg)

---

## 接口基本格式：

* 接口使用RESTful格式
* Response基本格式如下（JSON）：

| 字段名 | 类型     | 描述     | 例子  |
| ------ | -------- | -------- | ----- |
| code   | int      | 状态码   | 20000 |
| data   | json对象 | 数据     | null  |
| msg    | string   | 返回消息 | "ok"  |

> code:	200xx -> 正常返回，300xx -> 重定向，400xx -> 请求错误，500xx -> 系统错误


```json
{
    "code":	20000,
    "data": null,
    "msg": 	"ok",
}
```


* 测试端口：localhost:8000

---

## 买家

### 主页相关

#### 1.书名搜索

描述：使用用户输入的书名去搜索数据库中的书籍数据，并分页展示

方法：GET /api/v1/books/search

Query Parameters：

| 字段名 | 类型   | 描述              | 例子     |
| ------ | ------ | ----------------- | -------- |
| name   | string | 书名              | 高等数学 |
| page   | int    | 当前页码，初值为1 | 1        |
| items  | int    | 每一页书籍项目数  | 10       |

Response Body：

| 字段名   | 类型         | 描述            | 例子                    |
| -------- | ------------ | --------------- | ----------------------- |
| pages    | int          | 总分页数        | 10                      |
| items    | int          | 此页的项目数    | 2                       |
| item     | json对象数组 | 书籍介绍项目    |                         |
| title    | string       | 书名            | “高等数学”              |
| author   | string       | 作者            | “张宇昊”                |
| cover    | string       | 封面的URL       | “/static/cover/123.png” |
| price    | int          | 价格            | 36                      |
| salesnum | int          | 销量            | 12                      |
| descp    | string       | 详细描述图的URL | “/static/descp/123.png” |

```json
// GET /api/v1/books/search?name=高数&page=1&item=4
{
    "code": 20000,
    "data": {
        "pages": 	10,
        "items": 	2,
        "item":		[
            {
                "title":	"高等数学",
                "author":	"张宇昊",
                "cover":	"/static/cover/123.png",
                "price":	49,
                "salesnum":	12,
                "descp":	"/static/descp/123.png"
            },
            {
                "title":	"高等数学",
                "author":	"",
                "cover":	"/static/cover/124.png",
                "price":	34,
                "salesnum":	8,
                "descp"：	"/static/descp/124.png"
            }
        ]
    },
    "msg":	"ok"
}
```

#### ~~2.分类浏览~~

#### 3.推荐热销书籍

描述：系统向用户推荐的书籍

方法：GET /api/v1/books/recommend

Request Header：

| 字段名 | 类型   | 描述              | 例子     |
| ------ | ------ | ----------------- | -------- |
| page   | int    | 当前页码，初值为0 | 0        |
| items  | int    | 每一页书籍项目数  | 10       |

Response Body：

| 字段名   | 类型         | 描述            | 例子                    |
| -------- | ------------ | --------------- | ----------------------- |
| pages    | int          | 总分页数        | 10                      |
| items    | int          | 此页的项目数    | 2                       |
| item     | json对象数组 | 书籍介绍项目    |                         |
| title    | string       | 书名            | “高等数学”              |
| author   | string       | 作者            | “张宇昊”                |
| cover    | string       | 封面的URL       | “/static/cover/123.png” |
| price    | int          | 价格            | 36                      |
| salesnum | int          | 销量            | 12                      |
| descp    | string       | 详细描述图的URL | “/static/descp/123.png” |

```json
{
    "code": 20000,
    "data": {
        "pages": 	10,
        "items": 	2,
        "item":		[
            {
                "title":	"高等数学",
                "author":	"张宇昊",
                "cover":	"/static/cover/123.png",
                "price":	49,
                "salesnum":	12,
                "descp":	"/static/descp/123.png"
            },
            {
                "title":	"高等数学",
                "author":	"",
                "cover":	"/static/cover/124.png",
                "price":	34,
                "salesnum":	8,
                "descp"：	"/static/descp/124.png"
            }
        ]
    },
    "msg":	"ok"
}
```

#### ~~4.公告栏~~

描述：公告信息

方法：GET /api/v1/bulletin

Request Header：空







### 账号管理

#### 1.注册

描述：用户注册

方法:	POST /api/v1/user/register

Request Body：

| 字段名   | 类型   | 描述   | 例子     |
| -------- | ------ | ------ | -------- |
| username | string | 用户名 | woshizyh |
| password | string | 密码   | 12344    |

Response Body：

> code: 20000 -> 注册成功；40000 -> 用户名已注册

```json
{
	"code": 20000,
    "data": null,
    "msg": "注册成功"
}
{
	"code": 40000,
	"data": null,
	"msg": "用户已注册"
}
```

#### 2.登录

描述：用户登录

方法：POST /api/v1/user/login

Request Body：

| 字段名   | 类型   | 描述   | 例子     |
| -------- | ------ | ------ | -------- |
| username | string | 用户名 | woshizyh |
| password | string | 密码   | 12344    |

Response Body：

| 字段名   | 类型   | 描述   | 例子        |
| -------- | ------ | ------ | ----------- |
| id       | int    | 用户id | 12          |
| username | string | 用户名 | woshizyh    |
| token    | string | jwt    | e93fj8fb... |

> code:	20000 -> 正常；40001 -> 用户名或密码错误

```JSON
{
    "code": 20000,
    "data": {
        "id": 500,
        "username": "admin",
        "token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjUwMCwicmlkIjowLCJpYXQiOjE1MTI1NDQyOTksImV4cCI6MTUxMjYzMDY5OX0.eGrsrvwHm-tPsO9r_pxHIQ5i5L1kX9RX444uwnRGaIM"
    },
	"msg": "登录成功"
}
{
    "code": 40001,
    "data": null,
    "msg": "用户名或密码错误"
}
```

#### 3.修改密码

描述：修改用户密码

方法：PUT /api/v1/user/password

Request Header:

| 字段名  | 类型   | 描述       | 例子      |
| ------- | ------ | ---------- | --------- |
| user_id | int    | 用户的id号 | 12        |
| token   | string | jwt        | fe32af... |

Request Body:

| 字段名  | 类型   | 描述       | 例子      |
| ------------ | ------ | ----------- | -------- |
| password_old | string | 旧密码      | woshizyh |
| password_new | string | 新密码      | woshizxm |

Response Body:

> code:	20000-成功；40001-密码错误

```json
{
    "code": 20000,
    "data": null,
    "msg": "密码修改成功"
}
{
    "code": 40001,
    "data": null,
    "msg": "密码错误"
}
```

#### 4.修改用户名

描述：修改用户名

方法：PUT /api/v1/user/name

Request Header:

| 字段名  | 类型   | 描述       | 例子      |
| ------- | ------ | ---------- | --------- |
| user_id | int    | 用户的id号 | 12        |
| token   | string | jwt        | fe32af... |

Request Body:

| 字段名  		| 类型   | 描述       | 例子      |
| ------------ | ------ | ----------- | -------- |
| username_new | string | 新的用户名  | woshizxm |

Response Body:

> code:	20000-修改成功；40000-用户已注册

```json
{
  	"code": 20000,
  	"data": "woshizxm",
    "msg": "修改成功"
}
{
    "code": 40000,
    "data": null,
    "msg": "用户名已经注册"
}
```



### 收货地址管理

#### 1.增加收货地址

描述：增加一个收货地址

方法：POST /api/v1/user/address

Request Header:

| parameter | type   | description    | example     |
| --------- | ------ | -------------- | ----------- |
| user_id   | int    | 用户的id号     | 12          |
| token     | string | 用户的登录凭证 | 26374687234 |

Request Body:

| parameter | type   | description                                                  | example              |
| --------- | ------ | ------------------------------------------------------------ | -------------------- |
| realname  | string | 收货时候真实姓名                                             | 许沁宁               |
| address   | string | 收货地址 客户端将地址的省 市 县 详细地址 以空格间隔 拼接成字符串传给后台 | 江苏省 南京市 江宁区 |
| phone     | string | 手机号码                                                     | 13238274842          |

Response Body:

```json
{
    "code": 20000,
    "data": null,
    "msg": "收货地址添加成功"
}
```

#### 2.查看收货地址

描述：查看我的所有收货地址

方法：GET /api/v1/user/address

Request Header:

| parameter | type   | description    | example     |
| --------- | ------ | -------------- | ----------- |
| user_id   | int    | 用户的id号     | 12          |
| token     | string | 用户的登录凭证 | 26374687234 |

Response Body:

| parameter | type         | description                                                  | example              |
| --------- | ------------ | ------------------------------------------------------------ | -------------------- |
| realname  | string       | 收货时候真实姓名                                             | 许沁宁               |
| address   | string       | 收货地址 客户端将地址的省 市 县 详细地址 以空格间隔 拼接成字符串传给后台 | 江苏省 南京市 江宁区 |
| phone     | string       | 手机号码                                                     | 13238274842          |
| num       | int          | items的数量                                                  | 1                    |
| items     | json对象数组 | 收货地址数组                                                 |                      |
| id        | int          | 收货地址的id                                                 | 7                    |

```json
{
    "code": 20000,
    "data": {
        "num": 1,
        "items": [
            {
                "id":		7,
    			"realname":	"许沁宁",
    			"address":	"江苏省 南京市 江宁区 ***",
    			"phone":	"13247463548"
            },
        ]
    },
    "msg": "收货信息查询成功"
}
```

#### 3.删除收货地址

描述：删除一个收货地址

方法：DELETE /api/v1/user/address

Request Header:

| parameter | type   | description    | example     |
| --------- | ------ | -------------- | ----------- |
| user_id   | int    | 用户的id号     | 12          |
| token     | string | 用户的登录凭证 | 26374687234 |

Request Body:

| parameter | type | description     | example |
| --------- | ---- | --------------- | ------- |
| id        | int  | 地址address的id | 12      |

Response Body：

```json
{
    "code": 20000,
    "data": null,
    "msg": "ok"
}
```



### ~~购物车~~

#### 1.添加到购物车

#### 2.查看购物车内容

#### 3.从购物车中删除

#### 4、购物结算



---



## 卖家

### 仓库管理

#### 1、上传书籍

描述：卖家上传书籍信息

方法：POST /api/v1/user/books

Request Header:

| 字段名  | 类型   | 描述       | 例子      |
| ------- | ------ | ---------- | --------- |
| user_id | int    | 用户的id号 | 12        |
| token   | string | jwt        | fe32af... |

Request Body:

| 字段名 | 类型      | 描述     | 例子     |
| ------ | --------- | -------- | -------- |
| title  | string    | 书名     | “高数”   |
| author | string    | 作者     | “张宇昊” |
| cover  | form-data | 封面图片 |          |
| descp  | form-data | 描述图片 |          |
| price  | int       | 价格     | 45       |
| num    | int       | 数量     | 120      |
| kind   | string    | 图书类别 | “计算机” |

Response Body：

```json
{
    "code": 20000,
    "data":	null,
    "msg":	"ok"
}
```

#### 2、查看我的卖书

描述：查看卖家售卖的书籍

方法：GET /api/v1/user/books

Request Header:

| 字段名  | 类型   | 描述                         | 例子      |
| ------- | ------ | ---------------------------- | --------- |
| user_id | int    | 用户的id号                   | 12        |
| token   | string | jwt                          | fe32af... |
| page    | int    | 当前页码，初值为1（不是0！） | 1         |
| items   | int    | 每页的项目数                 | 8         |

Response Body：

| 字段名   | 类型         | 描述            | 例子                   |
| -------- | ------------ | --------------- | ---------------------- |
| pages    | int          | 总分页数        | 10                     |
| items    | int          | 此页的项目数    | 2                      |
| item     | json对象数组 | 书籍介绍项目    |                        |
| title    | string       | 书名            | “高等数学”             |
| author   | string       | 作者            | “张宇昊”               |
| cover    | string       | 封面的URL       | “static/cover/123.png” |
| price    | int          | 价格            | 36                     |
| salesnum | int          | 销量            | 12                     |
| descp    | string       | 详细描述图的URL | “static/descp/123.png” |
| num      | int          | 仓库里书的数量  | 107                    |
| kind     | string       | 书籍的类型      | “数学”                 |
| id       | int          | 书籍的id        | 21                     |

```json
{
    "code": 20000,
    "data": {
        "pages": 	10,
        "items": 	2,
        "item":		[
            {
                "id":		21,
                "title":	"高等数学",
                "author":	"张宇昊",
                "kind":		"数学",
                "cover":	"static/cover/123.png",
                "price":	49,
                "salesnum":	12,
                "num":		107,
                "descp":	"static/descp/123.png"
            },
            {
                "id":		22,
                "title":	"高等数学",
                "author":	"",
                "kind":		"数学",
                "cover":	"static/cover/124.png",
                "price":	34,
                "salesnum":	8,
                "num":		21,
                "descp"：	"static/descp/124.png"
            }
        ]
    },
    "msg":	"ok"
}
```

#### 3、修改书籍信息

描述：修改卖家的二手书信息

方法：PUT /api/v1/user/books

Request Header：

| 字段名  | 类型   | 描述       | 例子      |
| ------- | ------ | ---------- | --------- |
| user_id | int    | 用户的id号 | 12        |
| token   | string | jwt        | fe32af... |

Request Body:

| 字段名 | 类型   | 描述            | 例子                    |
| ------- | ------ | ---------- | --------- |
| title  | string | 书名            | “高等数学”              |
| author | string | 作者            | “张宇昊”                |
| cover  | form-data | 封面图片     |  |
| price  | int    | 价格            | 36                      |
| descp  | form-data | 描述图片 |            |
| num    | int    | 仓库里书的数量  | 107                     |
| kind   | string | 书籍的类型      | “数学”                  |
| id | int | 书籍的id | 21 |

> 上面的各字段除id外，都**不是必填**项目，不填的字段表示不变

Response Body:

```json
{
    "code":	20000,
    "data":	null,
    "msg":	"ok"
}
```

#### 4、删除书籍

描述：卖家删除自己的二手书

方法：DELETE /api/v1/user/books

Request Header:

| 字段名  | 类型   | 描述       | 例子      |
| ------- | ------ | ---------- | --------- |
| user_id | int    | 用户的id号 | 12        |
| token   | string | jwt        | fe32af... |

Request Body：

| 字段名 | 类型 | 描述     | 例子 |
| ------ | ---- | -------- | ---- |
| id     | int  | 书籍的id | 21   |

Response Body：

```json
{
    "code":	20000,
    "data":	null,
    "msg":	"ok"
}
```

### ~~订单管理~~

#### 1、查看我的订单

#### 2、订单处理-发货

---



## 管理员

### ~~主页相关~~

#### 1、图书分类设定

#### 2、公告内容设定