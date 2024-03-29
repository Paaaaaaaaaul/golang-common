package base_server_pay

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/becent/golang-common"
	"github.com/becent/golang-common/base-server-sdk"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

//接口版本
const VERSION = "1.0.1"
const VERSION_102 = "1.0.2"

//货币
const (
	CURRENCY_CC   = "CC"
	CURRENCY_USD  = "USD"
	CURRENCY_USDT = "USDT"
	CURRENCY_EOS  = "EOS"
	CURRENCY_XRP  = "XRP"
	CURRENCY_BTC  = "BTC"
	CURRENCY_BCH  = "BCH"
	CURRENCY_ETH  = "ETH"
)

//账户余额查询返回
type Account struct {
	MchId        string `json:"mch_id"`
	Currency     string `json:"currency"`
	AvailAmount  string `json:"avail_amount"`
	FreezeAmount string `json:"freeze_amount"`
}

//查询可用支付方式返回
type PayMethod struct {
	MethodCode           string `json:"methodCode"`
	MethodName           string `json:"methodName"`
	MinSingleOrderAmount string `json:"minSingleOrderAmount"`
	MaxSingleOrderAmount string `json:"maxSingleOrderAmount"`
}

//支付通道
type PayChannel struct {
	ChannelId            string `json:"channelId"`
	ChannelDesc          string `json:"channelDesc"`
	MinSingleOrderAmount string `json:"minSingleOrderAmount"`
	MaxSingleOrderAmount string `json:"maxSingleOrderAmount"`
}

//支付订单状态
const (
	PAY_STATUS_SUCCESS = "SUCCESS"
	PAY_STATUS_PAYING  = "PAYING"
	PAY_STATUS_FAIL    = "FAIL"

	WITHDRAW_STATUS_SUCCESS = "SUCCESS"
	WITHDRAW_STATUS_PAYING  = "PAYING"
	WITHDRAW_STATUS_FAIL    = "FAIL"
)

//钱包类型
const (
	WALLET_TYPE = "erc20"
)

//支付订单
type PayOrder struct {
	PayMethod      string `json:"pay_method"`
	PayChannel     string `json:"pay_channel"`
	MchId          string `json:"mch_id"`
	TransactionId  string `json:"transaction_id"`
	OutTradeNo     string `json:"out_trade_no"`
	NonceStr       string `json:"nonce_str"`
	SignType       string `json:"sign_type"`
	Detail         string `json:"detail"`
	Attach         string `json:"attach"`
	SpbillCreateIp string `json:"spbill_create_ip"`
	NotifyUrl      string `json:"notify_url"`
	UserOutFee     string `json:"user_out_fee"`
	UserOutType    string `json:"user_out_type"`
	MchInFee       string `json:"mch_in_fee"`
	MchInType      string `json:"mch_in_type"`
	TimeStart      int64  `json:"time_start"`
	TimeExpire     int64  `json:"time_expire"`
	TimeEnd        int64  `json:"time_end"`
	TradeStatus    string `json:"trade_status"`
	CodeContent    string `json:"code_content"`
	CodePage       string `json:"code_page"`
	Sign           string `json:"sign"`
	Version        string `json:"version"`
	ExchangeRate   string `json:"exchange_rate"`
}

//提现银行账户信息
type BankAccount struct {
	BankName       string
	BankUserName   string
	BankUserPhone  string
	BankBranchName string
	BankCardNo     string
	BankProvince   string
	BankCity       string
	QrCode         string
	QrCodeImgType  string
}

//数字货币账户
type BitcoinAccount struct {
	Address string
	UserId  string
}

//代付订单
type WithdrawOrder struct {
	PayMethod      string `json:"pay_method"`
	PayChannel     string `json:"pay_channel"`
	MchId          string `json:"mch_id"`
	TransactionId  string `json:"transaction_id"`
	OutTradeNo     string `json:"out_trade_no"`
	NonceStr       string `json:"nonce_str"`
	SignType       string `json:"sign_type"`
	Detail         string `json:"detail"`
	Attach         string `json:"attach"`
	SpbillCreateIp string `json:"spbill_create_ip"`
	NotifyUrl      string `json:"notify_url"`
	MchOutFee      string `json:"mch_out_fee"`
	MchOutType     string `json:"mch_out_type"`
	UserInType     string `json:"user_in_type"`
	UserInFee      string `json:"user_in_fee"`
	ExchangeRate   string `json:"exchange_rate"`
	TimeStart      string `json:"time_start"`
	TimeExpire     string `json:"time_expire"`
	TimeEnd        string `json:"time_end"`
	TradeStatus    string `json:"trade_status"`
	Sign           string `json:"sign"`
	Version        string `json:"version"`
}

//发送Aipay请求
func SendAipayRequest(controller string, action string, signKey string, params map[string]string, apiHost string) (res []byte, err *base_server_sdk.Error) {
	client := base_server_sdk.Instance
	if apiHost == "" {
		apiHost = client.Hosts.AiPayServerHost
	}

	params["sign_type"] = "MD5"
	params["sign"] = GenerateAipaySignature(params, signKey)

	response, err := client.DoRequest(apiHost, controller, action, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//生成签名
func GenerateAipaySignature(params map[string]string, signKey string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	plainText := ""
	for _, k := range keys {
		if k == "sign" {
			continue
		}
		if params[k] != "" {
			plainText += fmt.Sprintf("%v=%v&", k, params[k])
		}
	}

	plainText += "key=" + signKey

	fmt.Println(plainText)
	c := md5.New()
	c.Write([]byte(plainText))
	cipherStr := c.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}

//查询余额
func QueryBalance(mchId string, currency string, signKey string, version string, apiHost string) (*Account, *base_server_sdk.Error) {
	request := map[string]string{}
	request["mch_id"] = mchId
	request["nonce_str"] = strconv.FormatInt(time.Now().Unix(), 10)
	request["currency"] = currency
	request["version"] = version

	response, err := SendAipayRequest("account", "query", signKey, request, apiHost)
	if err != nil {
		return nil, err
	}

	var account Account
	if err := json.Unmarshal(response, &account); err != nil {
		common.ErrorLog("baseServerCommon_QueryBalance", request, "unmarshal fail: "+string(response))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return &account, nil
}

//查询可用支付方式
func SelectPayMethods(mchId string, userOutType string, userOutFee string, signKey string, version string, apiHost string) ([]PayMethod, *base_server_sdk.Error) {
	request := map[string]string{}
	request["mch_id"] = mchId
	request["nonce_str"] = strconv.FormatInt(time.Now().Unix(), 10)
	request["user_out_type"] = userOutType
	request["user_out_fee"] = userOutFee
	request["version"] = version

	response, err := SendAipayRequest("pay", "selectMethod", signKey, request, apiHost)
	if err != nil {
		return nil, err
	}

	var payMethods []PayMethod
	if err := json.Unmarshal(response, &payMethods); err != nil {
		common.ErrorLog("baseServerCommon_SelectPayMethods", request, "unmarshal fail: "+string(response))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return payMethods, nil
}

//查询可用通道
func SelectPayChannels(mchId string, payMethod string, userOutFee string, userOutType string, mchInType string, signKey string, version string, apiHost string) ([]PayChannel, *base_server_sdk.Error) {
	request := map[string]string{}
	request["mch_id"] = mchId
	request["nonce_str"] = strconv.FormatInt(time.Now().Unix(), 10)
	request["user_out_type"] = userOutType
	request["user_out_fee"] = userOutFee
	request["pay_method"] = payMethod
	request["mch_in_type"] = mchInType
	request["version"] = version

	response, err := SendAipayRequest("pay", "selectChannel", signKey, request, apiHost)
	if err != nil {
		return nil, err
	}

	var payChannels []PayChannel
	if err := json.Unmarshal(response, &payChannels); err != nil {
		common.ErrorLog("baseServerCommon_SelectPayChannels", request, "unmarshal fail: "+string(response))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return payChannels, nil
}

//提交支付订单
func SubmitPayOrder(mchId string,
	payMethod string,
	payChannel string,
	outTradeNo string,
	userOutFee string,
	userOutType string,
	mchInType string,
	detail string,
	attach string,
	notifyUrl string,
	signKey string,
	version string,
	apiHost string) (*PayOrder, *base_server_sdk.Error) {

	request := map[string]string{}
	request["mch_id"] = mchId
	request["nonce_str"] = strconv.FormatInt(time.Now().Unix(), 10)
	request["user_out_type"] = userOutType
	request["user_out_fee"] = userOutFee
	request["pay_method"] = payMethod
	request["mch_in_type"] = mchInType
	request["pay_channel"] = payChannel
	request["out_trade_no"] = outTradeNo
	request["detail"] = detail
	request["attach"] = attach
	request["notify_url"] = notifyUrl
	request["version"] = version

	response, err := SendAipayRequest("pay", "submit", signKey, request, apiHost)
	if err != nil {
		return nil, err
	}

	var order PayOrder
	if err := json.Unmarshal(response, &order); err != nil {
		fmt.Println(err.Error())
		common.ErrorLog("baseServerCommon_SubmitPayOrder", request, "unmarshal fail: "+string(response))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return &order, nil
}

//查询支付订单
func QueryPayOrder(mchId string, outTradeNo string, transactionId string, signKey string, version string, apiHost string) (*PayOrder, *base_server_sdk.Error) {
	request := map[string]string{}
	request["mch_id"] = mchId
	request["nonce_str"] = strconv.FormatInt(time.Now().Unix(), 10)
	request["transaction_id"] = transactionId
	request["out_trade_no"] = outTradeNo
	request["version"] = version

	response, err := SendAipayRequest("pay", "query", signKey, request, apiHost)
	if err != nil {
		return nil, err
	}

	var order PayOrder
	if err := json.Unmarshal(response, &order); err != nil {
		common.ErrorLog("baseServerCommon_QueryPayOrder", request, "unmarshal fail: "+string(response))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return &order, nil
}

//生成聚合支付链接
func GenerateUnionPayUrl(mchId string, currency string, userId string, reqTime string,
	amount string, notifyUrl string, redirectUrl string, attach string, signKey string, version string, apiHost string) (string, *base_server_sdk.Error) {
	request := map[string]string{}
	request["mch_id"] = mchId
	request["nonce_str"] = strconv.FormatInt(time.Now().Unix(), 10)
	request["currency"] = currency
	request["userId"] = userId
	request["time"] = reqTime
	request["amount"] = amount
	request["notify_url"] = notifyUrl
	request["redirect_url"] = redirectUrl
	request["sessionId"] = RandString(18)
	request["attach"] = attach
	request["version"] = version

	response, err := SendAipayRequest("union", "generateUrl", signKey, request, apiHost)
	if err != nil {
		return "", err
	}

	var url map[string]string
	if err := json.Unmarshal(response, &url); err != nil {
		common.ErrorLog("baseServerCommon_GenerateUnionPayUrl", request, "unmarshal fail: "+string(response))
		return "", base_server_sdk.ErrServiceBusy
	}

	return url["unionPayUrl"], nil
}

//查询可用提现方式
func SelectWithdrawMethods(mchId string, userInType string, signKey string, version string, apiHost string) ([]PayMethod, *base_server_sdk.Error) {
	request := map[string]string{}
	request["mch_id"] = mchId
	request["nonce_str"] = strconv.FormatInt(time.Now().Unix(), 10)
	request["user_in_type"] = userInType
	request["version"] = version

	response, err := SendAipayRequest("withdraw", "selectMethod", signKey, request, apiHost)
	if err != nil {
		return nil, err
	}

	var payMethods []PayMethod
	if err := json.Unmarshal(response, &payMethods); err != nil {
		common.ErrorLog("baseServerCommon_SelectWithdrawMethods", request, "unmarshal fail: "+string(response))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return payMethods, nil
}

//查询可用提现通道
func SelectWithdrawChannels(mchId string,
	payMethod string,
	userInFee string,
	userInType string,
	mchOutType string,
	mchOutFee string,
	signKey string,
	walletType string,
	version string,
	apiHost string) ([]PayChannel, *base_server_sdk.Error) {

	request := map[string]string{}
	request["mch_id"] = mchId
	request["nonce_str"] = strconv.FormatInt(time.Now().Unix(), 10)
	request["user_in_fee"] = userInFee
	request["user_in_type"] = userInType
	request["pay_method"] = payMethod
	request["mch_out_type"] = mchOutType
	request["mch_out_fee"] = mchOutFee
	request["version"] = version
	request["walletType"] = walletType

	response, err := SendAipayRequest("withdraw", "selectChannel", signKey, request, apiHost)
	if err != nil {
		return nil, err
	}

	var payChannels []PayChannel
	if err := json.Unmarshal(response, &payChannels); err != nil {
		common.ErrorLog("baseServerCommon_SelectWithdrawChannels", request, "unmarshal fail: "+string(response))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return payChannels, nil
}

//提交支付订单
func SubmitWithdrawOrder(mchId string,
	payMethod string,
	payChannel string,
	outTradeNo string,
	userInFee string,
	userInType string,
	mchOutType string,
	mchOutFee string,
	detail string,
	attach string,
	notifyUrl string,
	signKey string,
	bankAccount *BankAccount,
	bitcoinAccount *BitcoinAccount,
	version string,
	apiHost string) (*WithdrawOrder, *base_server_sdk.Error) {

	request := map[string]string{}
	request["mch_id"] = mchId
	request["nonce_str"] = strconv.FormatInt(time.Now().Unix(), 10)
	request["user_in_type"] = userInType
	request["user_in_fee"] = userInFee
	request["pay_method"] = payMethod
	request["mch_out_type"] = mchOutType
	request["mch_out_fee"] = mchOutFee
	request["pay_channel"] = payChannel
	request["out_trade_no"] = outTradeNo
	request["detail"] = detail
	request["attach"] = attach
	request["notify_url"] = notifyUrl
	request["version"] = version

	request["bank_name"] = bankAccount.BankName
	request["bank_user_name"] = bankAccount.BankUserName
	request["bank_user_phone"] = bankAccount.BankUserPhone
	request["bank_branch_name"] = bankAccount.BankBranchName
	request["bank_card_no"] = bankAccount.BankCardNo
	request["bank_province"] = bankAccount.BankProvince
	request["bank_city"] = bankAccount.BankCity
	request["qr_code"] = bankAccount.QrCode
	request["qr_code_img_type"] = bankAccount.QrCodeImgType

	request["receive_address"] = bitcoinAccount.Address
	request["user_id"] = bitcoinAccount.UserId

	response, err := SendAipayRequest("withdraw", "submit", signKey, request, apiHost)
	if err != nil {
		return nil, err
	}

	var order WithdrawOrder
	if err := json.Unmarshal(response, &order); err != nil {
		fmt.Println(err.Error())
		common.ErrorLog("baseServerCommon_SubmitWithdrawOrder", request, "unmarshal fail: "+string(response))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return &order, nil
}

//查询支付订单
func QueryWithdrawOrder(mchId string,
	outTradeNo string,
	transactionId string,
	signKey string,
	version string,
	apiHost string) (*WithdrawOrder, *base_server_sdk.Error) {

	request := map[string]string{}
	request["mch_id"] = mchId
	request["nonce_str"] = strconv.FormatInt(time.Now().Unix(), 10)
	request["transaction_id"] = transactionId
	request["out_trade_no"] = outTradeNo
	request["version"] = version

	response, err := SendAipayRequest("withdraw", "query", signKey, request, apiHost)
	if err != nil {
		return nil, err
	}

	var order WithdrawOrder
	if err := json.Unmarshal(response, &order); err != nil {
		common.ErrorLog("baseServerCommon_QueryWithdrawOrder", request, "unmarshal fail: "+string(response))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return &order, nil
}

//支付回调
type PayNotify struct {
	PayMethod      string `json:"pay_method"`
	MchId          string `json:"mch_id"`
	TransactionId  string `json:"transaction_id"`
	OutTradeNo     string `json:"out_trade_no"`
	PayChannel     string `json:"pay_channel"`
	NonceStr       string `json:"nonce_str"`
	TradeStatus    string `json:"trade_status"`
	SignType       string `json:"sign_type"`
	UserOutFee     string `json:"user_out_fee"`
	UserOutType    string `json:"user_out_type"`
	TimeStart      string `json:"time_start"`
	TimeExpire     string `json:"time_expire"`
	TimeEnd        string `json:"time_end"`
	Detail         string `json:"detail"`
	SpbillCreateIp string `json:"spbill_create_ip"`
	Attach         string `json:"attach"`
	NotifyUrl      string `json:"notify_url"`
	Version        string `json:"version"`
	UserId         string `json:"userId"`
	ExchangeRate   string `json:"exchange_rate"`
	MchInFee       string `json:"mch_in_fee"`
	MchInType      string `json:"mch_in_type"`
	Sign           string `json:"sign"`
}

//是否支付成功
func (notify *PayNotify) IsSuccess() bool {
	if notify.TradeStatus == PAY_STATUS_SUCCESS {
		return true
	}
	return false
}

//是否支付失败
func (notify *PayNotify) IsFail() bool {
	if notify.TradeStatus == PAY_STATUS_FAIL {
		return true
	}
	return false
}

//是否支付中
func (notify *PayNotify) IsPaying() bool {
	if notify.TradeStatus == PAY_STATUS_PAYING {
		return true
	}
	return false
}

const (
	NOTIFY_RESPONSE_SUCCESS = "SUCCESS"
	NOTIFY_RESPONSE_FAIL    = "FAIL"
)

//验证签名
func (notify *PayNotify) VerifySignature(mchId string, signKey string) bool {
	params := map[string]string{
		"pay_method":       notify.PayMethod,
		"mch_id":           notify.MchId,
		"transaction_id":   notify.TransactionId,
		"out_trade_no":     notify.OutTradeNo,
		"pay_channel":      notify.PayChannel,
		"nonce_str":        notify.NonceStr,
		"trade_status":     notify.TradeStatus,
		"sign_type":        notify.SignType,
		"user_out_fee":     notify.UserOutFee,
		"user_out_type":    notify.UserOutType,
		"time_start":       notify.TimeStart,
		"time_expire":      notify.TimeExpire,
		"time_end":         notify.TimeEnd,
		"detail":           notify.Detail,
		"spbill_create_ip": notify.SpbillCreateIp,
		"attach":           notify.Attach,
		"notify_url":       notify.NotifyUrl,
		"version":          notify.Version,
		"userId":           notify.UserId,
		"exchange_rate":    notify.ExchangeRate,
		"mch_in_fee":       notify.MchInFee,
		"mch_in_type":      notify.MchInType,
	}

	mySign := GenerateAipaySignature(params, signKey)
	if notify.Sign != mySign {
		return false
	}

	return true
}

// RandString 生成随机字符串
func RandString(len int) string {
	var r = rand.New(rand.NewSource(time.Now().Unix()))

	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
