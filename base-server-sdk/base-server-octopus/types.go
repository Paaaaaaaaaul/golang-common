package base_server_octopus

import (
	"errors"
)

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

	BusinessPublishAd BusinessId = 2001 			//otc发布广告
	BusinessBandBankCard BusinessId = 2002 			//otc绑定银行卡
	BusinessBandQRCode BusinessId = 2003 			//otc绑定支付二维码
	BusinessPayConfirmBySeller BusinessId = 2004 	//otc商家确认收款
	BusinessSell BusinessId = 2005 					//otc 卖出
)

var businessIdMap = map[int]BusinessId{
	1000: BusinessRegister,
	1001: BusinessLogin,
	1002: BusinessBindPhone,
	1003: BusinessUnBindPhone,
	1004: BusinessBindEmail,
	1005: BusinessUnBindEmail,
	1006: BusinessGetBackLoginPwd,
	1007: BusinessGetBackTransPwd,
	2001: BusinessPublishAd,
	2002: BusinessBandBankCard,
	2003: BusinessBandQRCode,
	2004: BusinessPayConfirmBySeller,
	2005: BusinessSell,
}

func GetBusinessId(id int) (BusinessId, error) {
	if _, ok := businessIdMap[id]; !ok {
		return 0, errors.New("invalid business type")
	}

	return businessIdMap[id], nil
}
