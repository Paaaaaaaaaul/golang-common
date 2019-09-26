package base_server_octopus

type BusinessId int

const (
	BusinessRegister        BusinessId = 1000 // 注册
	BusinessLogin           BusinessId = 1001 // 登录
	BusinessUpdatePhone     BusinessId = 1002 // 更新手机
	BusinessBindPhone       BusinessId = 1003 // 绑定手机
	BusinessUpdateEmail     BusinessId = 1004 // 更新邮箱
	BusinessBindEmail       BusinessId = 1005 // 绑定邮箱
	BusinessGetBackLoginPwd BusinessId = 1006 // 找回密码
)
