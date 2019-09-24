package base_server_octopus

import (
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
	"encoding/json"
)

type InitCaptchaRes struct {
	Success    int8   `json:"success"`
	CaptchaID  string `json:"gt"`
	Challenge  string `json:"challenge"`
	NewCaptcha int    `json:"new_captcha"`
}

//初始化验证码
//业务码请查看types.
//{
//"success": 0/1, //标识是否走本地验证
//"gt": "极验账户密钥",
//"challenge": "验证码唯一id",
//"new_captcha": 0/1 //标识是否走本地验证
//}
func InitCaptcha(orgId int, businessId BusinessId, account string, ip string) (*InitCaptchaRes, *base_server_sdk.Error) {
	if orgId == 0 || account == "" || ip == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["businessId"] = strconv.Itoa(int(businessId))
	params["account"] = account
	params["ip"] = ip

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.OctopusServerHost, "captcha", "initCaptcha", params)
	if err != nil {
		return nil, err
	}
	var resp InitCaptchaRes
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, base_server_sdk.ErrServiceBusy
	}
	return &resp, nil
}

//校验验证码
func VerifyCaptcha(orgId int, businessId BusinessId, account string, ip string, challenge, validate, seccode string) (bool, *base_server_sdk.Error) {
	if orgId == 0 || account == "" || ip == "" || challenge == "" || validate == "" || seccode == "" {
		return false, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["businessId"] = strconv.Itoa(int(businessId))
	params["account"] = account
	params["ip"] = ip
	params["challenge"] = challenge
	params["validate"] = validate
	params["seccode"] = seccode

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.OctopusServerHost, "captcha", "verifyCaptcha", params)
	if err != nil {
		return false, err
	}
	return true, nil
}