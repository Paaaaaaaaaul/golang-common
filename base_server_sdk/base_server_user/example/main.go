package main

import (
	"fmt"
	"github.com/becent/commom/base_server_sdk"
	"github.com/becent/commom/base_server_sdk/base_server_user"
	"time"
)

func main() {
	base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		OrgId:           5,
		AppId:           "10002",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			UserServerHost: "http://127.0.0.1:8081",
		},
		GRpcOnly: false,
	})
	defer base_server_sdk.ReleaseBaseServerSdk()

	now := time.Now()
	defer func(now time.Time) {
		println(time.Since(now).String())
	}(now)

	// 注册用户
	user, err := base_server_user.Register(&base_server_user.User{
		OrgId:    5,
		Phone:    "13560487593",
		LoginPwd: "123456",
		NickName: "song",
		Avatar:   "shuai.png",
		Ext:      "123",
	}, "")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("注册成功：[%v]\n", *user)
	}

	// 通过手机找回登录密码
	err = base_server_user.GetBackLoginPwdByPhone(5, "13560487593", "", "654321")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("找回登录密码成功\n")
	}

	// 登录
	user, err = base_server_user.LoginByPhone(5, "13560487593", "", "654321")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("登录成功：[%v]\n", *user)
	}

	// 获取用户信息
	user, err = base_server_user.GetUserInfo(5, user.UserId)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("获取用户信息成功：[%v]\n", *user)
	}

	// 修改登录密码
	err = base_server_user.UpdateLoginPwd(user.OrgId, user.UserId, "654321", "123456")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("登录密码修改成功\n")
	}

	// 再次登录
	user, err = base_server_user.LoginByPhone(5, "13560487593", "", "654321")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("登录成功：[%v]\n", *user)
	}

	// 改变密码再次登录
	user, err = base_server_user.LoginByPhone(5, "13560487593", "", "123456")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("登录成功：[%v]\n", *user)
	}

	// 实名认证
	if err = base_server_user.AuthRealName(5, user.UserId, "song", "liang", "360721199001040204"); err != nil {
		println(err.String())
	} else {
		println("实名认证成功")
	}

	// 获取用户信息
	user, err = base_server_user.GetUserInfo(5, user.UserId)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("获取用户信息成功：[%v]\n", *user)
	}

	// 通过手机找回交易密码
	err = base_server_user.GetBackTransPwdByPhone(5, "13560487593", "", "asdqwe")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("找回交易密码成功\n")
	}

	// 验证交易密码
	if err = base_server_user.AuthTransPwd(5, user.UserId, "123456"); err != nil {
		println(err.String())
	} else {
		println("验证交易密码成功")
	}

	// 验证交易密码
	if err = base_server_user.AuthTransPwd(5, user.UserId, "asdqwe"); err != nil {
		println(err.String())
	} else {
		println("验证交易密码成功")
	}

	// 更新交易密码
	if err = base_server_user.UpdateTransPwd(5, user.UserId, "asdqwe", "123465"); err != nil {
		println(err.String())
	} else {
		println("更新交易密码成功")
	}

	info := make(base_server_user.UserFields)
	info.SetNickName("新昵称")
	info.SetBirthDay("2000-10-10")
	info.SetAvatar("new.png")
	info.SetExt("456")
	info.SetSex(base_server_user.Boy)
	if err = base_server_user.UpdateUserInfo(5, user.UserId, info); err != nil {
		println(err.String())
	} else {
		println("更新用户信息成功")
	}

	// 获取用户信息
	user, err = base_server_user.GetUserInfo(5, user.UserId)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("获取用户信息成功：[%v]\n", *user)
	}
}
