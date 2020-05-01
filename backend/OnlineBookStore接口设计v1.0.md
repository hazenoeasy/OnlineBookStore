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
| id       | int          | 书籍的id        | 1                       |

```json
// GET /api/v1/books/search?name=高数&page=1&item=4
{
    "code": 20000,
    "data": {
        "pages": 	10,
        "items": 	2,
        "item":		[
            {
                "id":		2,
                "title":	"高等数学",
                "author":	"张宇昊",
                "cover":	"/static/cover/123.png",
                "price":	49,
                "salesnum":	12,
                "descp":	"/static/descp/123.png"
            },
            {
                "id":		3,
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

#### ~~3.推荐热销书籍~~

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
| id       | int          | 书籍id          | 1                       |

```json
{
    "code": 20000,
    "data": {
        "pages": 	10,
        "items": 	2,
        "item":		[
            {
                "id":		2,
                "title":	"高等数学",
                "author":	"张宇昊",
                "cover":	"/static/cover/123.png",
                "price":	49,
                "salesnum":	12,
                "descp":	"/static/descp/123.png"
            },
            {
                "id":		3,
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



### 购物车

#### 1.添加到购物车

描述：将书籍加入用户的车中

方法：POST /api/v1/user/cart

Request Header:

| 字段名  | 类型   | 描述       | 例子      |
| ------- | ------ | ---------- | --------- |
| user_id | int    | 用户的id号 | 12        |
| token   | string | jwt        | fe32af... |

Request Body:

| 字段名 | 类型   | 描述             | 例子             |
| ------ | ------ | ---------------- | ---------------- |
| id     | int    | 书籍的id         | 12               |
| name   | string | 书名             | 高等数学         |
| price  | int    | 图书单价         | 45               |
| cover  | string | 封面url          | static/323xx.png |
| num    | int    | 加入购物车的数量 | 1                |

Response Body：

```json
{
    "code":	20000,
    "data":	null,
    "msg":	"ok"
}
```



#### 2.查看购物车内容

描述：用户查看自己的购物车

方法：GET /api/v1/user/cart

Request Header:

| 字段名  | 类型   | 描述       | 例子      |
| ------- | ------ | ---------- | --------- |
| user_id | int    | 用户的id号 | 12        |
| token   | string | jwt        | fe32af... |

Response Body:

| 字段名 | 类型         | 描述             | 例子             |
| ------ | ------------ | ---------------- | ---------------- |
| id     | int          | 书籍的id         | 12               |
| name   | string       | 书名             | 高等数学         |
| price  | int          | 图书单价         | 45               |
| cover  | string       | 封面url          | static/323xx.png |
| num    | int          | 加入购物车的数量 | 1                |
| item   | json对象数组 | 购物车内的商品   |                  |
| items  | int          | item字段的长度   |                  |

```json
{
    "code":	20000,
    "data":	{
        items:	2,
        item:	[
            {
                "id":		23,
                "name":		"高等数学",
                "price":	45,
                "cover":	"static/323cover23.png",
                "num":		1
            },
            {
                "id":		24,
                "name":		"高等数学2",
                "price":	45,
                "cover":	"static/323cover24.png",
                "num":		1
            }
        ]
    },
    "msg":	"ok"
}
```



#### 3.从购物车中删除

描述：用户从购物车中删除某个商品

方法：DEL /api/v1/user/cart

Request Header：

| 字段名  | 类型   | 描述       | 例子      |
| ------- | ------ | ---------- | --------- |
| user_id | int    | 用户的id号 | 12        |
| token   | string | jwt        | fe32af... |

Request Body：

| 字段名  | 类型 | 描述       | 例子 |
| ------- | ---- | ---------- | ---- |
| book_id | int  | 二手书的id | 12   |

Response Body：

```json
{
    "code":	20000,
    "data":	null,
    "msg":	"ok"
}
```



#### 4、购物结算

描述：用户结算购物车中的商品，并生成订单

方法：POST /api/v1/user/order

Request Header:

| 字段名  | 类型   | 描述       | 例子      |
| ------- | ------ | ---------- | --------- |
| user_id | int    | 用户的id号 | 12        |
| token   | string | jwt        | fe32af... |

Request Body：

| 字段名 | 类型   | 描述                     | 例子      |
| ------ | ------ | ------------------------ | --------- |
| addr_id | int | 收货地址的id | 12 |
| id | int | 书籍的id | 1 |
| num | int | 书籍的数量 | 1 |

Response Body：

```json
{
    "code":	20000,
    "data":	null,
    "msg":	"ok"
}
```





#### 5、付款（测试版为假付款）

描述：虚假的订单付款

方法：PUT /api/v1/user/order

Request Header:

| 字段名  | 类型   | 描述       | 例子      |
| ------- | ------ | ---------- | --------- |
| user_id | int    | 用户的id号 | 12        |
| token   | string | jwt        | fe32af... |

Request Body:

| 字段名   | 类型 | 描述     | 例子 |
| -------- | ---- | -------- | ---- |
| order_id | int  | 订单的id | 12   |

Response Body：

```json
{
    "code":	20000,
    "data":	null,
    "msg":	"ok"
}
```



### 订单管理

#### 1.查看我的订单

描述：查看我的购物订单

方法：GET /api/v1/user/order

Request Header:

| 字段名  | 类型   | 描述              | 例子      |
| ------- | ------ | ----------------- | --------- |
| user_id | int    | 用户的id号        | 12        |
| token   | string | jwt               | fe32af... |
| page    | int    | 订单页数，初值为1 | 1         |
| items   | int    | 每页订单数量      | 4         |

Response Body：

| 字段名  | 类型     | 描述         | 例子     |
| ------- | -------- | ------------ | -------- |
| items   | int      | 订单项的数量 | 2        |
| item    | json对象 | 订单项       |          |
| id      | int      | 订单id       | 123      |
| status  | int      | 描述订单状态 | 1        |
| title   | string   | 书名         | “高数”   |
| price   | int      | 订单价格     | 98       |
| num     | int      | 书的购买量   | 1        |
| express | string   | 快递号       | “230230” |

**订单状态**：

>  0-未付款；1-用户已经付款，但卖家未发货；2-卖家已发货；3-订单已结束

```json
{
    "code": 20000,
    "data":	{
     	"items":	2,
        "item":		[
            {
                "id":		123,
                "status":	1,
                "title":	"高数",
                "price":	44,
                "num":		1,
                "express":	""
            },
            {
                "id":		134,
                "status":	2,
                "title":	"高数2",
                "price":	40,
                "num":		1,
                "express":	"1298823910012"
            }
        ]
    },
    "msg":	"ok"
}
```



####  ~~2.退/换商品~~

#### 3.确认收货

描述：用户确认收货，完成订单处理

方法：PUT /api/v1/user/commodity

Request Header:

| 字段名  | 类型   | 描述       | 例子      |
| ------- | ------ | ---------- | --------- |
| user_id | int    | 用户的id号 | 12        |
| token   | string | jwt        | fe32af... |

Request Body:

| 字段名   | 类型 | 描述     | 例子 |
| -------- | ---- | -------- | ---- |
| order_id | int  | 订单的id | 12   |

Response Body：

```json
{
    "code":	20000,
    "data":	null,
    "msg":	"ok"
}
```



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

### 订单管理

#### 1、卖家查看订单

描述：商家查看自己的订单（包括需要处理的和已经结束的订单）

方法：GET /api/v1/user/seller/order

Request Header:

| 字段名  | 类型   | 描述              | 例子      |
| ------- | ------ | ----------------- | --------- |
| user_id | int    | 用户的id号        | 12        |
| token   | string | jwt               | fe32af... |
| page    | int    | 订单页数，初值为1 | 1         |
| items   | int    | 每页订单数量      | 4         |
| status  | int    | 订单状态          | 1         |

**订单状态**：

>  0-未付款；1-用户已经付款，但卖家未发货；2-卖家已发货；3-订单已结束

Response Body：

| 字段名         | 类型     | 描述         | 例子          |
| -------------- | -------- | ------------ | ------------- |
| items          | int      | 订单项的数量 | 2             |
| item           | json对象 | 订单项       |               |
| id             | int      | 订单id       | 123           |
| status         | int      | 描述订单状态 | 1             |
| title          | string   | 书名         | “高数”        |
| price          | int      | 订单价格     | 98            |
| num            | int      | 书的购买量   | 1             |
| express        | string   | 快递号       | “230230”      |
| consumer_name  | string   | 买家姓名     | “bxh”         |
| consumer_addr  | string   | 买家地址     | “河北省涞”    |
| consumer_phone | string   | 买家手机号码 | “17248429312” |

```json
{
    "code": 20000,
    "data":	{
     	"items":	2,
        "item":		[
            {
                "id":				123,
                "status":			1,
                "title":			"高数",
                "price":			44,
                "num":				1,
                "express":			"",
                "consumer_name":	"bxh",
                "consumer_addr":	"河北省",
                "consumer_phone":	"1482918219"
            },
            {
                "id":		134,
                "status":	2,
                "title":	"高数2",
                "price":	40,
                "num":		1,
                "express":	"1298823910012",
                "consumer_name":	"bxh",
                "consumer_addr":	"河北省",
                "consumer_phone":	"1482918219"
            }
        ]
    },
    "msg":	"ok"
}
```


#### 2、订单处理-发货

描述：商家发货

方法：POST /api/v1/user/commodity

Request Header:

| 字段名  | 类型   | 描述       | 例子      |
| ------- | ------ | ---------- | --------- |
| user_id | int    | 用户的id号 | 12        |
| token   | string | jwt        | fe32af... |

Request Body:

| 字段名   | 类型 | 描述     | 例子 |
| -------- | ---- | -------- | ---- |
| order_id | int  | 订单的id | 12   |

Response Body：

```json
{
    "code":	20000,
    "data":	null,
    "msg":	"ok"
}
```



### ~~数据统计~~



---



## 管理员

### ~~主页相关~~

#### 1、图书分类设定

#### 2、公告内容设定



### ~~后台数据统计~~

