package base_server_octopus

import (
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
	"encoding/json"
)

type GenerateGaRes struct {
	QrCode string `json:"qrCode"`
	SecretKey string `json:"secretKey"`
}

func GenerateGa(orgId int, email string, phone string) (*GenerateGaRes, *base_server_sdk.Error) {
	if orgId == 0 || (email == "" && phone == "") {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["email"] = email
	params["phone"] = phone

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.OctopusServerHost, "ga", "generateGa", params)
	if err != nil {
		return nil, err
	}
	var resp GenerateGaRes
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, base_server_sdk.ErrServiceBusy
	}
	return &resp, nil
}

func VerifyGa(orgId int, secret string, gaCode string) (bool, *base_server_sdk.Error) {
	if orgId == 0 || secret == "" || gaCode== "" {
		return false, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["secret"] = secret
	params["gaCode"] = gaCode

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.OctopusServerHost, "ga", "verifyGa", params)
	if err != nil {
		return false, err
	}

	return true, nil
}
