# base_server_account 接口说明文档

## 初始化base_server_sdk
```go
base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
    OrgId:           8,
    AppId:           "10008",
    AppSecretKey:    "hiojklsankldlksdnlsdasd",
    RequestTimeout:  5 * time.Second,
    IdleConnTimeout: 10 * time.Minute,
    Hosts: base_server_sdk.Hosts{
        AccountServerHost: "http://127.0.0.1:8081",
    },
    GRpcOnly: false,
})

// ....

defer base_server_sdk.ReleaseBaseServerSdk()
```

## 相关model
```go
type Account struct {
	AccountId    int64  `json:"accountId"`
	OrgId        int    `json:"orgId"`
	UserId       int64  `json:"userId"`
	Currency     string `json:"currency"`
	AvailAmount  string `json:"availAmount"`
	FreezeAmount string `json:"freezeAmount"`
	Status       int    `json:"status"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
}

type LogList struct {
	LogId      int64	`json:"logId"`
	UserId     int64	`json:"userId"`
	Currency   string	`json:"currency"`
	LogType    int		`json:"logType"`
	Amount     string	`json:"amount"`
	CreateTime int64	`json:"createTime"`
}

type base_server_sdk.Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
```

## 相关错误码
```go
1001 参数错误
2001 账户已存在
2002 账户创建失败
2003 账户不存在
2004 更新状态失败
1009 BC操作失败
2005 账户可用增加失败
2007 可用余额不足
2008 解冻失败
2009 账户可用减少失败
2010 账户冻结减少失败
2011 账户日志创建失败
2004 更新状态失败
```

## 接口说明

- 创建账户

func CreateAccount(account *Account) (*Account, *base_server_sdk.Error)

```go
注意:
1. orgId必须大于0

异常错误:
1001 参数错误
2001 账户已存在
2002 账户创建失败
```

- 账户信息

func AccountInfo(orgId int, userId int64, currency string) (*Account, *base_server_sdk.Error)

```go
异常错误:
1001 参数错误
2003 账户不存在
```

- 状态变更

func UpdateStatus(orgId int, accountId int64, status int) *base_server_sdk.Error

```go
异常错误:
1001 参数错误
2003 账户不存在
2004 更新状态失败
```

- 金额操作

func OperateAmount(orgId int, accountId int64, opType int, amount string) *base_server_sdk.Error

```go
类型枚举:
1	//可用-加
2	//可用-减
3	//冻结-加
4	//冻结-减
5	//解冻-冻结进可用

异常错误:
1001 参数错误
2003 账户不存在
1009 BC操作失败
2005 账户可用增加失败
2007 可用余额不足
2008 解冻失败
2009 账户可用减少失败
2010 账户冻结减少失败
2011 账户日志创建失败
```

- 账户日志列表

func AccountLogList(orgId int, userId int64, opType int, currency string, page, limit int) (*[]LogList, *base_server_sdk.Error)

```go
类型枚举:
1	//可用-加
2	//可用-减
3	//冻结-加
4	//冻结-减
5	//解冻-冻结进可用

异常错误:
1001 参数错误
2003 账户不存在
2004 更新状态失败
```