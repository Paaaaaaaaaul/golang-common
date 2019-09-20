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

func InitCaptcha(orgId int, account string, ip string) (*InitCaptchaRes, *base_server_sdk.Error) {
	if orgId == 0 || account == "" || ip == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
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

func VerifyCaptcha(orgId int, account string, ip string, challenge, validate, seccode string) (bool, *base_server_sdk.Error) {
	if orgId == 0 || account == "" || ip == "" || challenge == "" || validate == "" || seccode == "" {
		return false, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["account"] = account
	params["ip"] = ip
	params["challenge"] = challenge
	params["validate"] = validate
	params["seccode"] = seccode

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.OctopusServerHost, "captcha", "initCaptcha", params)
	if err != nil {
		return false, err
	}

	return true, nil
}