---
title: mallshop
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# mallshop

Base URLs:

* <a href="http://dev-cn.your-api-server.com">开发环境: http://dev-cn.your-api-server.com</a>

# Authentication

# 用户

## POST 登录-用户名/密码

POST /u/v1/base/login

> Body 请求参数

```json
{
  "mobile": "string",
  "password": "string",
  "captcha": "string",
  "captcha_id": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Content-Type|header|string| 是 |none|
|body|body|object| 否 |none|
|» mobile|body|string| 是 |手机号码|
|» password|body|string| 是 |密码|
|» captcha|body|string| 是 |验证码|
|» captcha_id|body|string| 是 |验证码id|

> 返回示例

> 200 Response

```json
{
  "id": 0,
  "nick_name": "string",
  "token": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|integer|true|none||用户id|
|» nick_name|string|true|none||昵称|
|» token|string|true|none||jwt的token|

## POST 修改个人信息

POST /u/v1/base/info/modify

> Body 请求参数

```json
{
  "name": "string",
  "gender": "string",
  "birthday": "string",
  "email": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 | title|none|
|» name|body|string| 是 ||姓名|
|» gender|body|string| 是 ||性别|
|» birthday|body|string| 是 ||生日|
|» email|body|string| 是 ||邮箱|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 发送短信

POST /u/v1/sms

> Body 请求参数

```json
{
  "mobile": "string",
  "type": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» mobile|body|string| 否 ||手机号码|
|» type|body|number| 否 ||1 表示注册 2表示找回密码|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取验证码

GET /u/v1/base/captcha

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 用户列表

GET /u/v1/user/list

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 登录-短信验证码

POST /base/v1/login-sms

> Body 请求参数

```json
{
  "mobile": "string",
  "code": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 | title|none|
|» mobile|body|string| 是 ||手机号码|
|» code|body|string| 是 ||短信验证码|

> 返回示例

> 200 Response

```json
{
  "id": 0,
  "nick_name": "string",
  "token": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|number|true|none||none|
|» nick_name|string|true|none||none|
|» token|string|true|none||none|

## POST 注册

POST /u/v1/base/register

> Body 请求参数

```json
{
  "code": "string",
  "mobile": "string",
  "password": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» code|body|string| 是 | 短信验证码|短信验证码|
|» mobile|body|string| 是 | 手机号码|手机号码|
|» password|body|string| 是 | 密码|密码|

> 返回示例

> 200 Response

```json
{
  "id": 0,
  "nick_name": "string",
  "token": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|integer|false|none|用户id|用户id|
|» nick_name|string|false|none|昵称|昵称|
|» token|string|false|none|jwt的token|jwt的token|

# 商品

## PUT 修改商品信息

PUT /g/v1/goods/68

> Body 请求参数

```json
{
  "name": "string",
  "shop_price": 0,
  "goods_brief": "string",
  "desc": "string",
  "ship_free": true,
  "images": [
    "string"
  ],
  "front_image": "string",
  "stocks": 0,
  "market_price": 0,
  "goods_sn": "string",
  "category": 0,
  "brand": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» name|body|string| 是 ||商品名称|
|» shop_price|body|number| 是 ||价格|
|» goods_brief|body|string| 是 ||简介|
|» desc|body|string| 是 ||详情|
|» ship_free|body|boolean| 是 ||是否免运费|
|» images|body|[string]| 是 ||none|
|» front_image|body|string| 是 ||封面图|
|» stocks|body|integer| 是 ||库存|
|» market_price|body|number| 是 ||市场价格|
|» goods_sn|body|string| 是 ||商品编号|
|» category|body|integer| 是 ||分类|
|» brand|body|integer| 是 ||品牌id|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PATCH 修改商品状态

PATCH /v1/goods/1

> Body 请求参数

```json
{
  "sale": true,
  "new": true,
  "hot": true
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» sale|body|boolean| 否 ||是否上架|
|» new|body|boolean| 否 ||是否新品|
|» hot|body|boolean| 否 ||是否热销|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除商品

DELETE /g/v1/goods/1

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 商品详情

GET /g/v1/goods/1

> 返回示例

> 200 Response

```json
{
  "id": 0,
  "name": "string",
  "shop_price": 0,
  "goods_brief": "string",
  "ship_free": true,
  "images": [
    "string"
  ],
  "front_image": "string",
  "desc": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|number|false|none||none|
|» name|string|false|none||none|
|» shop_price|number|false|none||none|
|» goods_brief|string|false|none||none|
|» ship_free|boolean|false|none||none|
|» images|[string]|false|none||none|
|» front_image|string|false|none||none|
|» desc|string|true|none||none|

## GET 商品列表

GET /g/v1/goods

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|pmin|query|string| 否 ||最低价格|
|pmax|query|string| 否 ||最高价格|
|ih|query|string| 否 ||是否热门|
|in|query|string| 否 ||是否最新|
|it|query|string| 否 ||是否tab栏|
|c|query|string| 否 ||分类的id|
|pn|query|string| 否 ||页码|
|pnum|query|string| 否 ||每页的数量|
|q|query|string| 否 ||搜索关键词|
|b|query|string| 否 ||品牌id|

> 返回示例

> 200 Response

```json
{
  "data": [
    {
      "brand": {
        "id": 0,
        "logo": "string",
        "name": "string"
      },
      "ctegory": {
        "id": 0,
        "name": "string"
      },
      "desc": "string",
      "desc_images": [
        "string"
      ],
      "front_image": "string",
      "goods_brief": "string",
      "id": 0,
      "images": [
        "string"
      ],
      "is_hot": true,
      "is_new": true,
      "name": "string",
      "on_sale": true,
      "ship_free": true,
      "shop_price": 0
    }
  ],
  "total": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[object]|false|none||none|
|»» brand|object|true|none||none|
|»»» id|number|false|none||none|
|»»» logo|string|false|none||none|
|»»» name|string|false|none||none|
|»» ctegory|object|true|none||none|
|»»» id|number|false|none||none|
|»»» name|string|false|none||none|
|»» desc|string|true|none||none|
|»» desc_images|[string]|true|none||none|
|»» front_image|string|true|none||none|
|»» goods_brief|string|true|none||none|
|»» id|number|true|none||none|
|»» images|[string]|true|none||none|
|»» is_hot|boolean|true|none||none|
|»» is_new|boolean|true|none||none|
|»» name|string|true|none||none|
|»» on_sale|boolean|true|none||none|
|»» ship_free|boolean|true|none||none|
|»» shop_price|number|true|none||none|
|» total|number|false|none||none|

## POST 添加商品

POST /g/v1/goods

> Body 请求参数

```json
{
  "name": "string",
  "shop_price": 0,
  "goods_brief": "string",
  "ship_free": true,
  "images": [
    "string"
  ],
  "front_image": "string",
  "stocks": 0,
  "market_price": 0,
  "goods_sn": "string",
  "category": 0,
  "desc_images": [
    "string"
  ],
  "brand": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» name|body|string| 是 ||none|
|» shop_price|body|number| 是 ||none|
|» goods_brief|body|string| 是 ||none|
|» ship_free|body|boolean| 是 ||是否免运费|
|» images|body|[string]| 是 ||none|
|» front_image|body|string| 是 ||none|
|» stocks|body|integer| 是 ||库存|
|» market_price|body|number| 是 ||none|
|» goods_sn|body|string| 是 ||none|
|» category|body|integer| 是 ||none|
|» desc_images|body|[string]| 是 ||商品详情图片|
|» brand|body|integer| 是 ||品牌id|

> 返回示例

> 200 Response

```json
{
  "id": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|integer|false|none||none|

## GET 获取商品库存

GET /g/v1/goods/68/stocks

> 返回示例

> 200 Response

```json
{
  "stocks": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» stocks|integer|false|none||none|

# 商品分类

## PUT 修改分类

PUT /g/v1/categorys/239

> Body 请求参数

```json
{
  "name": "string",
  "level": 0,
  "parent": 0,
  "is_tab": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» name|body|string| 否 ||none|
|» level|body|integer| 否 ||none|
|» parent|body|integer| 否 ||none|
|» is_tab|body|integer| 否 ||1 表示设置tab 0表示不设置|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

*title*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|

## DELETE 删除分类

DELETE /g/v1/categorys/239

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 商品分类列表

GET /g/v1/categorys

> 返回示例

> 200 Response

```json
[
  {
    "id": 0,
    "name": "string",
    "level": 0,
    "is_tab": true,
    "parent_category": null
  }
]
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|number|true|none||none|
|» name|string|true|none||none|
|» level|integer|true|none||none|
|» is_tab|boolean|true|none||none|
|» parent_category|null|true|none||none|

## POST 添加分类

POST /g/v1/categorys

> Body 请求参数

```json
{
  "name": "string",
  "level": 0,
  "parent": 0,
  "is_tab": true
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 | title|none|
|» name|body|string| 是 ||名称|
|» level|body|integer| 是 ||1表示"一级" 2表示"二级"，3表示三级|
|» parent|body|integer| 是 ||父类目|
|» is_tab|body|boolean| 是 ||是否显示在tab页|

> 返回示例

> 200 Response

```json
{
  "id": 0,
  "name": "string",
  "category_type": 0,
  "parent_category": 0,
  "field_9": "string",
  "is_tab": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

*title*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|integer|true|none||id|
|» name|string|true|none||名称|
|» category_type|integer|true|none||1表示"一级" 2表示"二级"，3表示三级|
|» parent_category|integer|true|none||父类目|
|» field_9|string|true|none||是否显示在tab页|
|» is_tab|string|true|none||none|

## GET 获取分类品牌

GET /g/v1/categorybrands/130366

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取分类详情

GET /g/v1/categorys/110

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 品牌

## PUT 修改品牌

PUT /g/v1/brands/1

> Body 请求参数

```json
{
  "name": "string",
  "logo": "http://example.com"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» name|body|string| 否 ||none|
|» logo|body|string(uri)| 否 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

*title*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|

## DELETE 删除品牌

DELETE /g/v1/brands/1

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 品牌列表

GET /g/v1/brands

> 返回示例

> 200 Response

```json
[
  {
    "id": 0,
    "name": "string",
    "image": "string"
  }
]
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|number|false|none||none|
|» name|string|false|none||none|
|» image|string|false|none||none|

## POST 添加品牌

POST /g/v1/brands

> Body 请求参数

```json
{
  "name": "string",
  "logo": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» name|body|string| 否 ||none|
|» logo|body|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "id": 0,
  "name": "string",
  "image": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|number|false|none||none|
|» name|string|false|none||none|
|» image|string|false|none||none|

# 品牌分类

## PUT 修改品牌分类

PUT /g/v1/categorybrands/1

> Body 请求参数

```json
{
  "category_id": 0,
  "brand_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» category_id|body|integer| 是 ||类别id|
|» brand_id|body|integer| 是 ||品牌id|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

*title*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|

## DELETE 删除品牌分类

DELETE /g/v1/categorybrands/1

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 品牌分类列表

GET /g/v1/categorybrands

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|category_id|query|string| 否 ||分类的id|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 添加品牌分类

POST /g/v1/categorybrands

> Body 请求参数

```json
{
  "category_id": 0,
  "brand_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» category_id|body|integer| 是 ||分类id|
|» brand_id|body|integer| 是 ||品牌id|

> 返回示例

> 200 Response

```json
{
  "id": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|integer|true|none||记录id|

# 用户留言

## POST 添加留言

POST /up/v1/message

> Body 请求参数

```json
{
  "type": 0,
  "subject": "string",
  "message": "string",
  "file": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» type|body|integer| 是 ||留言类型: 1(留言),2(投诉),3(询问),4(售后),5(求购)|
|» subject|body|string| 是 ||主题|
|» message|body|string| 是 ||留言内容|
|» file|body|string| 是 ||上传的文件|

> 返回示例

> 200 Response

```json
{
  "id": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|integer|false|none||记录id|

## GET 留言列表

GET /up/v1/message

> 返回示例

> 200 Response

```json
[
  {
    "file": "string",
    "id": 0,
    "message": "string",
    "subject": "string",
    "type": 0,
    "user_id": 0
  }
]
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» file|string|false|none||none|
|» id|integer|false|none||none|
|» message|string|false|none||none|
|» subject|string|false|none||none|
|» type|integer|false|none||none|
|» user_id|integer|false|none||none|

# 用户地址

## DELETE 删除地址

DELETE /up/v1/address/1

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 更新地址

PUT /up/v1/address/1

> Body 请求参数

```json
{
  "province": "string",
  "city": "string",
  "district": "string",
  "address": "string",
  "signer_name": "string",
  "signer_mobile": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» province|body|string| 否 ||none|
|» city|body|string| 否 ||none|
|» district|body|string| 否 ||none|
|» address|body|string| 否 ||none|
|» signer_name|body|string| 否 ||none|
|» signer_mobile|body|string| 否 ||none|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 地址列表

GET /up/v1/address

> 返回示例

> 200 Response

```json
[
  {
    "address": "string",
    "city": "string",
    "district": "string",
    "id": 0,
    "province": "string",
    "signer_mobile": "string",
    "signer_name": "string",
    "user_id": 0
  }
]
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» address|string|false|none||none|
|» city|string|false|none||none|
|» district|string|false|none||none|
|» id|integer|false|none||none|
|» province|string|false|none||none|
|» signer_mobile|string|false|none||none|
|» signer_name|string|false|none||none|
|» user_id|integer|false|none||none|

## POST 添加地址

POST /up/v1/address

> Body 请求参数

```json
{
  "province": "string",
  "city": "string",
  "district": "string",
  "address": "string",
  "signer_name": "string",
  "signer_mobile": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» province|body|string| 否 ||省份|
|» city|body|string| 否 ||城市|
|» district|body|string| 否 ||地区|
|» address|body|string| 否 ||具体地址|
|» signer_name|body|string| 否 ||收货人|
|» signer_mobile|body|string| 否 ||收货人电话|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 用户收藏

## DELETE 取消收藏

DELETE /up/userfav

> Body 请求参数

```json
{
  "goods": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» goods|body|integer| 否 ||商品id|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 添加收藏

POST /up/v1/userfav

> Body 请求参数

```json
{
  "goods": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» goods|body|integer| 是 ||商品id|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 订单

## POST 支付接口

POST /o/v1/alipay/return

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 新建订单

POST /o/v1/order

> Body 请求参数

```json
{
  "address": "string",
  "name": "string",
  "mobile": "string",
  "post": "string",
  "items": [
    {
      "goods": 0,
      "nums": 0
    }
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» address|body|string| 否 ||none|
|» name|body|string| 否 ||none|
|» mobile|body|string| 否 ||none|
|» post|body|string| 否 ||none|
|» items|body|[object]| 否 ||none|
|»» goods|body|integer| 否 ||none|
|»» nums|body|integer| 否 ||none|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 订单列表

GET /o/v1/order

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 订单详情

GET /o/v1/order/18

> 返回示例

> 200 Response

```json
{
  "address": "string",
  "goods": [
    {
      "id": 0,
      "image": "string",
      "name": "string",
      "nums": 0,
      "price": 0
    }
  ],
  "id": 0,
  "mobile": "string",
  "name": "string",
  "pay_type": "string",
  "post": "string",
  "status": "string",
  "total": 0,
  "user": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» address|string|false|none||收货人地址|
|» goods|[object]|false|none||none|
|»» id|number|true|none||商品id|
|»» image|string|true|none||商品图片|
|»» name|string|true|none||商品名称|
|»» nums|number|true|none||数量|
|»» price|number|true|none||商品价格|
|» id|number|false|none||订单id|
|» mobile|string|false|none||收货人电话|
|» name|string|false|none||收货人姓名|
|» pay_type|string|false|none||alipay：支付宝 wechat：微信|
|» post|string|false|none||备注|
|» status|string|false|none||订单状态|
|» total|number|false|none||订单总金额|
|» user|integer|true|none||用户id|

# 购物车

## PUT 修改购物车信息

PUT /o/v1/shop_cart/1

> Body 请求参数

```json
{
  "nums": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» nums|body|integer| 否 ||数量|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除购物车记录

DELETE /o/v1/shop_cart/1

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 查看购物车记录

GET /o/v1/shop_cart

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 添加商品

POST /o/v1/shop_cart

> Body 请求参数

```json
{
  "goods": 0,
  "nums": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» goods|body|integer| 否 ||商品id|
|» nums|body|integer| 否 ||数量|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 轮播图

## PUT 修改轮播图

PUT /g/v1/banners/7

> Body 请求参数

```json
{
  "image": "string",
  "type": "string",
  "category_id": 0,
  "index": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» image|body|string| 否 ||none|
|» type|body|string| 否 ||none|
|» category_id|body|integer| 否 ||none|
|» index|body|integer| 否 ||none|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除轮播图

DELETE /g/v1/banners/1

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 添加轮播图

POST /g/v1/banners

> Body 请求参数

```json
{
  "image": "string",
  "index": 0,
  "url": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Content-Type|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» image|body|string| 否 ||none|
|» index|body|integer| 否 ||none|
|» url|body|string| 是 ||none|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 轮播图列表

GET /g/v1/banners

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

