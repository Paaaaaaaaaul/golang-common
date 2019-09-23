package base_server_account

import (
	"encoding/json"
	common "github.com/becent/golang-common"
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
)

type Account struct {
	AccountId    int64  `json:"accountId"`
	OrgId        int    `json:"orgId"`
	UserId       int64  `json:"userId"`
	Currency     string `json:"currency"`
	AvailAmount  string `json:"availAmount"`
	FreezeAmount string `json:"freezeAmount"`
	Status       int    `json:"status"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
}

type LogList struct {
	LogId      int64	`json:"logId"`
	UserId     int64	`json:"userId"`
	Currency   string	`json:"currency"`
	LogType    int		`json:"logType"`
	Amount     string	`json:"amount"`
	CreateTime int64	`json:"createTime"`
}

//  POST CreateAccount 创建账户
//
//	注意:
//	1. orgId必须大于0
//
//	异常错误:
//	1001 参数错误
//	2001 账户已存在
//	2002 账户创建失败
func CreateAccount(account *Account) (*Account, *base_server_sdk.Error) {
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(account.OrgId)
	params["userId"] = strconv.FormatInt(account.UserId, 10)
	params["currency"] = account.Currency

	if params["orgId"] == "0" || params["userId"] == "0" || params["currency"] == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.AccountServerHost, "account", "createAccount", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, account); err != nil {
		common.ErrorLog("baseServerSdk_CreateAccount", params, "unmarshal account fail" + string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}
	return account, nil
}

//	账户信息
//	POST account/updateStatus
//
//	异常错误:
//	1001 参数错误
//	2003 账户不存在
func AccountInfo(orgId int, userId int64, currency string) (*Account, *base_server_sdk.Error) {
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["currency"] = currency

	if params["orgId"] == "0" || params["userId"] == "0" || params["currency"] == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.AccountServerHost, "account", "accountInfo", params)
	if err != nil {
		return nil, err
	}

	account := &Account{}
	if err := json.Unmarshal(data, account); err != nil {
		common.ErrorLog("baseServerSdk_AccountInfo", params, "unmarshal account fail" + string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}
	return account, nil
}

//	状态变更
//	POST account/updateStatus
//
//	异常错误:
//	1001 参数错误
//	2003 账户不存在
//	2004 更新状态失败
func UpdateStatus(orgId int, accountId int64, status int) *base_server_sdk.Error {
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["accountId"] = strconv.FormatInt(accountId, 10)
	params["status"] = strconv.Itoa(status)

	if params["orgId"] == "0" || params["accountId"] == "0" || params["status"] == "0" {
		return base_server_sdk.ErrInvalidParams
	}

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.AccountServerHost, "account", "updateStatus", params)
	if err != nil {
		return err
	}
	return nil
}

//  金额操作
//  POST account/operateAmount
//	类型枚举:
//	1	//可用-加
//	2	//可用-减
//	3	//冻结-加
//	4	//冻结-减
//	5	//解冻-冻结进可用
//
//	异常错误:
//	1001 参数错误
//	2003 账户不存在
//	1009 BC操作失败
//	2005 账户可用增加失败
//	2007 可用余额不足
//	2008 解冻失败
//	2009 账户可用减少失败
//	2010 账户冻结减少失败
//	2011 账户日志创建失败
func OperateAmount(orgId int, accountId int64, opType int, amount string) *base_server_sdk.Error {
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["accountId"] = strconv.FormatInt(accountId, 10)
	params["opType"] = strconv.Itoa(opType)
	params["amount"] = amount

	if params["orgId"] == "0" || params["accountId"] == "0" || params["opType"] == "0" || amount == "" {
		return base_server_sdk.ErrInvalidParams
	}

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.AccountServerHost, "account", "operateAmount", params)
	if err != nil {
		return err
	}
	return nil
}

// 账户日志列表
// post account/accountLogList
//
//类型枚举:
//1	//可用-加
//2	//可用-减
//3	//冻结-加
//4	//冻结-减
//5	//解冻-冻结进可用
//
//异常错误:
//1001 参数错误
//2003 账户不存在
//2004 更新状态失败
func AccountLogList(orgId int, userId int64, opType int, currency string, page, limit int) (*[]LogList, *base_server_sdk.Error) {
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["opType"] = strconv.Itoa(opType)
	params["currency"] = currency
	params["page"] = strconv.Itoa(page)
	params["limit"] = strconv.Itoa(limit)

	if params["orgId"] == "0" || params["userId"] == "0" || params["opType"] == "0" || params["page"] <= "0" || params["limit"] <= "0" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.AccountServerHost, "account", "accountLogList", params)
	if err != nil {
		return nil, err
	}

	logList := &[]LogList{}
	if err := json.Unmarshal(data, logList); err != nil {
		common.ErrorLog("baseServerSdk_CreateAccount", params, "unmarshal account fail" + string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}
	return logList, nil
}