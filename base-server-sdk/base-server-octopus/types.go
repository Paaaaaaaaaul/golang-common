package base_server_octopus

type BusinessId int

const (
	BusinessRegister        BusinessId = 1000 // 注册
	BusinessLogin           BusinessId = 1001 // 登录
	BusinessBindPhone       BusinessId = 1002 // 绑定手机
	BusinessUnBindPhone     BusinessId = 1003 // 解绑手机
	BusinessBindEmail       BusinessId = 1004 // 绑定邮箱
	BusinessUnBindEmail     BusinessId = 1005 // 解绑邮箱
	BusinessGetBackLoginPwd BusinessId = 1006 // 找回密码
	BusinessGetBackTransPwd BusinessId = 1007 // 找回支付密码

	BusinessPublishAd          BusinessId = 2001 //otc发布广告
	BusinessBandBankCard       BusinessId = 2002 //otc绑定银行卡
	BusinessBandQRCode         BusinessId = 2003 //otc绑定支付二维码
	BusinessPayConfirmBySeller BusinessId = 2004 //otc商家确认收款
	BusinessSell               BusinessId = 2005 //otc 卖出
)

func GetBusinessId(id int) (BusinessId, error) {
	return BusinessId(id), nil
}
