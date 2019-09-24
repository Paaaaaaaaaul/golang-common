## base_server_sdk

## 初始化base_server_sdk
```go
base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
    AppId:           "10000",
    AppSecretKey:    "hiojklsankldlksdnlsdasd",
    RequestTimeout:  5 * time.Second,
    IdleConnTimeout: 10 * time.Minute,
    Hosts: base_server_sdk.Hosts{
        UserServerHost: "http://127.0.0.1:8081",
        AccountServerHost: "http://127.0.0.1:8082",
        StatisticServerHost: "http://127.0.0.1:8083",
        OctopusServerHost: "http://127.0.0.1:8084",
    },
    GRpcOnly: false,
})

// ...
err := base_server_octopus.SendEmailCode(5, 1000, "xxxx@qq.com", "zh")

defer base_server_sdk.ReleaseBaseServerSdk()
```

## 测试环境
- http：http://127.0.0.1:8081
- grpc：127.0.0.1:18081

## 相关错误码
```go
1000    服务繁忙
1001    参数错误
1002    未找到邮件模板
1003    验证码发送太频繁
1004    发送邮件失败
1005    验证码检验失败
1006    无最新校验记录
1007    未找到短信模板
1008    发送短信失败
1009    实名认证失败
1010    生成GA密钥失败
1011    检验GA失败
1012    验证码初始化失败
1013    验证码校验失败
```

## 相关业务码
```go
1000  注册
1001  登录
1002  更新手机
1003  绑定手机
1004  更新邮箱
1005  绑定邮箱
1006  找回密码
```


## 邮件服务

**发送邮件验证码**
- 请求地址：base_server_octopus.SendEmailCode
-  参数
||~参数||~是否必填||~类型||~说明||
|| orgId || 是  || int || 白标id，仅当平台超管时可以使用 ||
|| businessId|| 是  || int|| 业务id ||
|| email|| 是  || string || 邮箱 ||
|| lang|| 是  || string || 语言 zh en ko ||


**校验邮件验证码**
- 请求地址：base_server_octopus.VerifyEmailCode
-  参数
||~参数||~是否必填||~类型||~说明||
|| orgId || 是  || int || 白标id，仅当平台超管时可以使用 ||
|| businessId|| 是  || int|| 业务id ||
|| email|| 是  || string || 邮箱 ||
|| code|| 是  || string || 验证码 ||


**校验上次邮件验证码是否通过**
- 请求地址：base_server_octopus.CheckLastEmailVerifyResult
-  参数
||~参数||~是否必填||~类型||~说明||
|| orgId || 是  || int || 白标id，仅当平台超管时可以使用 ||
|| businessId|| 是  || int|| 业务id ||
|| email|| 是  || string || 邮箱 ||


## 短信服务

**发送短信验证码**
- 请求地址：base_server_octopus.SendSimCode
-  参数
||~参数||~是否必填||~类型||~说明||
|| orgId || 是  || int || 白标id，仅当平台超管时可以使用 ||
|| businessId|| 是  || int|| 业务id ||
|| phone|| 是  || string || 手机号||
|| lang|| 是  || string || 语言 zh en ko ||


**校验短信验证码**
- 请求地址：base_server_octopus.VerifySimCode
-  参数
||~参数||~是否必填||~类型||~说明||
|| orgId || 是  || int || 白标id，仅当平台超管时可以使用 ||
|| businessId|| 是  || int|| 业务id ||
|| phone|| 是  || string || phone||
|| code|| 是  || string || 验证码 ||


**校验上次短信验证码是否通过**
- 请求地址：base_server_octopus.CheckLastSimVerifyResult
-  参数
||~参数||~是否必填||~类型||~说明||
|| orgId || 是  || int || 白标id，仅当平台超管时可以使用 ||
|| businessId|| 是  || int|| 业务id ||
|| phone|| 是  || string || 手机号||


## 实名验证

**实名验证**
- 请求地址：idcard/authRealName
-  参数
||~参数||~是否必填||~类型||~说明||
|| orgId || 是  || int || 白标id，仅当平台超管时可以使用 ||
|| name|| 是  || int|| 姓名 ||
|| cardNo|| 是  || string || 身份证号||


## 谷歌验证

**谷歌验证初始化（获取密钥）**
- 请求地址：ga/generateGa
-  参数
||~参数||~是否必填||~类型||~说明||
|| orgId || 是  || int || 白标id，仅当平台超管时可以使用 ||
|| account||  是 || string || 账户 ||
- 返回
```go
{
"qrCode": "...", // 二维码链接
"secretKey": "base64encodestring" //密钥
}
```

**校验code**
- 请求地址：ga/verifyGa
-  参数
||~参数||~是否必填||~类型||~说明||
|| orgId || 是  || int || 白标id，仅当平台超管时可以使用 ||
|| secret||  是 || string || 密钥 ||
|| gaCode||  是 || string || 验证码 ||


## 验证服务

**验证码初始化**
- 请求地址：captcha/initCaptcha
-  参数
||~参数||~是否必填||~类型||~说明||
|| account || 是  || string || 账户  ||
|| clientType||  是 || string || 客户端类型 web andriod ios ||
|| ip||  是 || string || ip地址 ||
- 返回
```go
{
"success": 0/1, //标识是否走本地验证
"gt": "极验账户密钥",
"challenge": "验证码唯一id",
"new_captcha": 0/1 //标识是否走本地验证
}
```

**服务端校验验证码
- 请求地址：captcha/verifyCaptcha
-  参数
||~参数||~是否必填||~类型||~说明||
|| account || 是  || string || 账户  ||
|| clientType||  是 || string || 客户端类型 web andriod ios ||
|| ip||  是 || string || ip地址 ||
|| challenge||  是 || string || 验证码id ||
|| validate||  是 || string || 检验字符串 ||
|| seccode||  是 || string || 验证等级 ||
