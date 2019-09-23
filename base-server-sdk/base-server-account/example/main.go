package main

import (
	"fmt"
	"github.com/becent/golang-common/base-server-sdk"
	"github.com/becent/golang-common/base-server-sdk/base-server-account"
	"time"
)

func main() {
	base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		OrgId:           8,
		AppId:           "10008",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts:           base_server_sdk.Hosts{
			AccountServerHost: "http://127.0.0.1:8081",
		},
		GRpcOnly:        false,
	})
	defer base_server_sdk.ReleaseBaseServerSdk()

	now := time.Now()
	defer func(now time.Time) {
		println(time.Since(now).String())
	}(now)

	// 创建账户
	account, err := base_server_account.CreateAccount(&base_server_account.Account{
		OrgId:        8,
		UserId:       100000,
		Currency:     "CC",
	})

	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("创建成功: {%v}\n", *account)
	}

	// 账户信息
	account, err = base_server_account.AccountInfo(8, 100000, "CC")

	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("账户信息: {%v}\n", *account)
	}

	// 账户状态更新
	err = base_server_account.UpdateStatus(8, 9, 2)
	if err != nil {
		println(err.String())
	} else {
		println("状态更新成功")
	}

	// 金额操作
	err = base_server_account.OperateAmount(8, 9, 1, "100")
	if err != nil {
		println(err.String())
	} else {
		println("操作成功")
	}

	// 账户日志列表
	logList, err := base_server_account.AccountLogList(8, 1, 2, "", 1, 10)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("账户列表： {%v}", *logList)
	}
}