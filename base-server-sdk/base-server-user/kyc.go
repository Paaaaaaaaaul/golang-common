package base_server_user

import (
	"encoding/json"
	common "github.com/becent/golang-common"
	base_server_sdk "github.com/becent/golang-common/base-server-sdk"
	"strconv"
)

type UserKyc struct {
	KycId       int64  `gorm:"column:kycId;primary_key;AUTO_INCREMENT" json:"kycId"` // KYC id
	OrgId       int    `gorm:"column:orgId" json:"orgId"`                            // 项目id
	UserId      int64  `gorm:"column:userId" json:"userId"`                          // 用户id
	Nationality string `gorm:"column:nationality" json:"nationality"`                // 国籍
	CertType    int    `gorm:"column:certType" json:"certType"`                      // 认证类型 1:个人 2:企业
	CertIdType  int    `gorm:"column:certIdType" json:"certIdType"`                  // 认证证件类型 1:身份证 2:护照 3:营业执照
	CertLevel   int    `gorm:"column:certLevel" json:"certLevel"`                    // 认证等级 1:初级 2:高级 3:公户
	CertStat    int    `gorm:"column:certStat" json:"certStat"`                      // 认证状态 1:未认证  2:认证中 3:认证失败 4:认证成功
	CertNo      string `gorm:"column:certNo" json:"certNo"`                          // 认证号码 身份证号/社会统一信用代码
	CertName    string `gorm:"column:certName" json:"certName"`                      // 认证名称 姓名/企业名称
	ImgFront    string `gorm:"column:imgFront" json:"imgFront"`                      // 身份证人像面
	ImgBack     string `gorm:"column:imgBack" json:"imgBack"`                        // 身份证国徽面
	ImgHandheld string `gorm:"column:imgHandheld" json:"imgHandheld"`                // 身份证手持照
	ImgLicense  string `gorm:"column:imgLicense" json:"imgLicense"`                  // 企业营业执照
	FailReason  string `gorm:"column:failReason" json:"failReason"`                  // 审核失败原因
	CreateTime  int64  `gorm:"column:createTime" json:"createTime"`                  // 创建时间
	UpdateTime  int64  `gorm:"column:updateTime" json:"updateTime"`                  // 更新时间
}

const (
	CERT_TYPE_PERSON     = 1 //个人
	CERT_TYPE_ENTERPRISE = 2 //企业

	CERT_LEVEL_BASIC      = 1 //初级
	CERT_LEVEL_ADVANCED   = 2 //高级
	CERT_LEVEL_ENTERPRISE = 3 //企业公户

	CERT_STAT_NO   = 1 //未认证
	CERT_STAT_ING  = 2 //认证中
	CERT_STAT_FAIL = 3 //失败
	CERT_STAT_YES  = 4 //通过

	CERT_ID_TYPE_ID_CARD  = 1 //身份证
	CERT_ID_TYPE_PASSPORT = 2 //护照
	CERT_ID_TYPE_LICENSE  = 3 //营业执照

	KYC_CERT_STAT_BASIC_YES = 1			//初级完成
	KYC_CERT_STAT_ADVANCE_AUDIT = 2		//高级审核中
	KYC_CERT_STAT_ADVANCE_YES = 3		//高级完成
	KYC_CERT_STAT_ENTERPRISE_AUDIT = 4	//企业审核
	KYC_CERT_STAT_ENTERPRISE_YES = 5	//企业完成

)



func Apply() (int64, *base_server_sdk.Error) {
	params := make(map[string]string)
	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "reserveUserId", params)
	if err != nil {
		return 0, err
	}

	ret := make(map[string]string)
	if err := json.Unmarshal(data, &ret); err != nil {
		common.ErrorLog("baseServerSdk_ReserveUserId", params, "unmarshal data fail: "+string(data))
		return 0, base_server_sdk.ErrServiceBusy
	}

	userId, e := strconv.ParseInt(ret["userId"], 10, 64)
	if e != nil {
		return 0, base_server_sdk.ErrServiceBusy
	}

	return userId, nil
}
