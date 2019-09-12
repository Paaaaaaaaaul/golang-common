package base_server_user

import (
	"github.com/becent/commom"
	"github.com/becent/commom/base_server_sdk"
	json "github.com/json-iterator/go"
	"strconv"
)

type User struct {
	UserId     int64  `json:"userId"`
	OrgId      int    `json:"orgId"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Account    string `json:"account"`
	LoginPwd   string `json:"loginPwd"`
	TransPwd   string `json:"transPwd"`
	NickName   string `json:"nickName"`
	Avatar     string `json:"avatar"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	IdCard     string `json:"idCard"`
	Sex        int    `json:"sex"`
	BirthDay   string `json:"birthDay"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	Ext        string `json:"ext"`
}

// Register 注册用户
//
// 1、orgId必须大于0
// 2、account、phone、email至少要有一个有值
// 3、account有值时，password必须有值
// 4、code有值时，将会校验短信或者邮件验证码
// 5、其他字段非必填
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1002 账户已经注册
func Register(user *User, code string) (*User, *base_server_sdk.Error) {
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(user.OrgId)
	params["phone"] = user.Phone
	params["email"] = user.Email
	params["account"] = user.Account
	params["loginPwd"] = user.LoginPwd
	params["transPwd"] = user.TransPwd
	params["nickName"] = user.NickName
	params["avatar"] = user.Avatar
	params["sex"] = strconv.Itoa(user.Sex)
	params["birthDay"] = user.BirthDay
	params["ext"] = user.Ext
	params["code"] = code

	if params["orgId"] == "0" || (params["phone"] == "" && params["account"] == "" && params["email"] == "") || (params["account"] != "" && params["loginPwd"] == "") {
		return nil, base_server_sdk.ErrInvalidParams
	}

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "register", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, user); err != nil {
		common.ErrorLog("baseServerSdk_Register", params, "unmarshal user fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return user, nil
}

// LoginByAccount 通过账号进行登录
//
// 1、password有值时，会进行登录密码校验，为空则略过
//
// 异常返回：
// 1003 账号不存在
// 1004 用户已被冻结
// 1005 密码错误
func LoginByAccount(orgId int, account string, password string) (*User, *base_server_sdk.Error) {
	if orgId == 0 || account == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["account"] = account
	params["password"] = password

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "loginByAccount", params)
	if err != nil {
		return nil, err
	}

	user := &User{}
	if err := json.Unmarshal(data, user); err != nil {
		common.ErrorLog("baseServerSdk_LoginByAccount", params, "unmarshal user fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return user, nil
}

// LoginByPhone 通过手机号进行登录
//
// 1、code有值时，会进行短信校验，为空则略过
// 2、password有值时，会进行登录密码校验，为空则略过
//
// 异常返回：
// 1003 账号不存在
// 1004 用户已被冻结
// 1005 密码错误
// 1008 短信验证码错误
func LoginByPhone(orgId int, phone string, code string, password string) (*User, *base_server_sdk.Error) {
	if orgId == 0 || phone == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["phone"] = phone
	params["code"] = code
	params["password"] = password

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "loginByPhone", params)
	if err != nil {
		return nil, err
	}

	user := &User{}
	if err := json.Unmarshal(data, user); err != nil {
		common.ErrorLog("baseServerSdk_LoginByPhone", params, "unmarshal user fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return user, nil
}

// LoginByEmail 通过手机号进行登录
//
// 1、code有值时，会进行邮件校验，为空则略过
// 2、password有值时，会进行登录密码校验，为空则略过
//
// 异常返回：
// 1003 账号不存在
// 1004 用户已被冻结
// 1005 密码错误
// 1009 邮件验证码错误
func LoginByEmail(orgId int, email string, code string, password string) (*User, *base_server_sdk.Error) {
	if orgId == 0 || email == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["email"] = email
	params["code"] = code
	params["password"] = password

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "loginByEmail", params)
	if err != nil {
		return nil, err
	}

	user := &User{}
	if err := json.Unmarshal(data, user); err != nil {
		common.ErrorLog("baseServerSdk_LoginByEmail", params, "unmarshal user fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return user, nil
}

// GetUserInfo 获取用户信息
//
// 异常返回：
// 1003 用户不存在
func GetUserInfo(orgId int, userId int64) (*User, *base_server_sdk.Error) {
	if orgId == 0 || userId == 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "getUserInfo", params)
	if err != nil {
		return nil, err
	}

	user := &User{}
	if err := json.Unmarshal(data, user); err != nil {
		common.ErrorLog("baseServerSdk_LoginByEmail", params, "unmarshal user fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return user, nil
}

// GetBackLoginPwdByPhone 通过手机找回登录密码
//
// 1、code有值的话会进行短信校验，为空则忽略
// 2、password是要设置的新密码
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
// 1008 短信验证码错误
func GetBackLoginPwdByPhone(orgId int, phone, code, password string) *base_server_sdk.Error {
	if orgId == 0 || phone == "" || password == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["phone"] = phone
	params["code"] = code
	params["password"] = password

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "getBackLoginPwd", params)
	if err != nil {
		return err
	}

	return nil
}

// GetBackLoginPwdByEmail 通过邮箱找回登录密码
//
// 1、code有值的话会进行邮箱校验，为空则忽略
// 2、password是要设置的新密码
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
// 1009 邮箱验证码错误
func GetBackLoginPwdByEmail(orgId int, email, code, password string) *base_server_sdk.Error {
	if orgId == 0 || email == "" || password == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["email"] = email
	params["code"] = code
	params["password"] = password

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "getBackLoginPwd", params)
	if err != nil {
		return err
	}

	return nil
}

// GetBackTransPwdByPhone 通过手机找回支付密码
//
// 1、code有值的话会进行短信校验，为空则忽略
// 2、password是要设置的新密码
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
// 1008 短信验证码错误
func GetBackTransPwdByPhone(orgId int, phone, code, password string) *base_server_sdk.Error {
	if orgId == 0 || phone == "" || password == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["phone"] = phone
	params["code"] = code
	params["password"] = password

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "getBackTransPwd", params)
	if err != nil {
		return err
	}

	return nil
}

// GetBackTransPwdByEmail 通过邮箱找回支付密码
//
// 1、code有值的话会进行邮箱校验，为空则忽略
// 2、password是要设置的新密码
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
// 1009 邮箱验证码错误
func GetBackTransPwdByEmail(orgId int, email, code, password string) *base_server_sdk.Error {
	if orgId == 0 || email == "" || password == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["email"] = email
	params["code"] = code
	params["password"] = password

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "getBackTransPwd", params)
	if err != nil {
		return err
	}

	return nil
}

// UpdateLoginPwd 更新登录密码
//
// oldPassword有值则校验旧密码，为空则略过
// newPassword必须有值
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
func UpdateLoginPwd(orgId int, userId int64, oldPassword, newPassword string) *base_server_sdk.Error {
	if orgId == 0 || userId == 0 || newPassword == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["oldPassword"] = oldPassword
	params["newPassword"] = newPassword

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "updateLoginPwd", params)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTransPwd 更新交易密码
//
// oldPassword有值则校验旧密码，为空则略过
// newPassword必须有值
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
func UpdateTransPwd(orgId int, userId int64, oldPassword, newPassword string) *base_server_sdk.Error {
	if orgId == 0 || userId == 0 || newPassword == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["oldPassword"] = oldPassword
	params["newPassword"] = newPassword

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "updateTransPwd", params)
	if err != nil {
		return err
	}

	return nil
}

// AuthRealName 实名认证
//
// idCard有值的话，会进行身份证校验，为空则忽略
// 真实姓名为firstName+lastName
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1010 身份证认证失败
func AuthRealName(orgId int, userId int64, firstName, lastName, idCard string) *base_server_sdk.Error {
	if orgId == 0 || userId == 0 {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["firstName"] = firstName
	params["lastName"] = lastName
	params["idCard"] = idCard

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "authRealName", params)
	if err != nil {
		return err
	}

	return nil
}

type Sex int

const (
	_ Sex = iota
	Boy
	Girl
)

type UserFields map[string]string

func (u UserFields) SetNickName(nickName string) {
	u["nickName"] = nickName
}

func (u UserFields) SetAvatar(avatar string) {
	u["avatar"] = avatar
}

func (u UserFields) SetSex(sex Sex) {
	u["sex"] = strconv.Itoa(int(sex))
}

func (u UserFields) SetBirthDay(birthDay string) {
	u["birthDay"] = birthDay
}

func (u UserFields) SetExt(ext string) {
	u["ext"] = ext
}

// UpdateUserInfo 更新用户信息
//
// 可以更新的字段为：
// 1、nickName
// 2、avatar
// 3、sex
// 4、birthDay
// 5、ext
//
// 用例:
// info := make(base_server_user.UserFields)
// info.SetNickName("jak")
// info.SetSex(base_server_user.Boy)
//
// base_server_user.UpdateUserInfo(1, 1000, info)
//
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
func UpdateUserInfo(orgId int, userId int64, info UserFields) *base_server_sdk.Error {
	if orgId == 0 || userId == 0 {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	for key, val := range info {
		params[string(key)] = val
	}

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "updateUserInfo", params)
	if err != nil {
		return err
	}

	return nil
}

// BindAccount 绑定登录账号
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1006 重复绑定
func BindAccount(orgId int, userId int64, account, password string) *base_server_sdk.Error {
	if orgId == 0 || userId == 0 || account == "" || password == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["account"] = account
	params["password"] = password

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "bindAccount", params)
	if err != nil {
		return err
	}

	return nil
}

// BindPhone 绑定手机
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1006 重复绑定
func BindPhone(orgId int, userId int64, phone, code string) *base_server_sdk.Error {
	if orgId == 0 || userId == 0 || phone == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["phone"] = phone
	params["code"] = code

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "bindPhone", params)
	if err != nil {
		return err
	}

	return nil
}

// BindEmail 绑定邮箱
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1006 重复绑定
func BindEmail(orgId int, userId int64, email, code string) *base_server_sdk.Error {
	if orgId == 0 || userId == 0 || email == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["email"] = email
	params["code"] = code

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "bindEmail", params)
	if err != nil {
		return err
	}

	return nil
}

// AuthTransPwd 校验支付密码
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
func AuthTransPwd(orgId int, userId int64, password string) *base_server_sdk.Error {
	if orgId == 0 || userId == 0 || password == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["password"] = password

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "authTransPwd", params)
	if err != nil {
		return err
	}

	return nil
}