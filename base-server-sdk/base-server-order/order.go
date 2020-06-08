package base_server_order

import (
	"encoding/json"
	"fmt"
	base_server_sdk "github.com/becent/golang-common/base-server-sdk"
)

type Order struct {
	OrderId    int64  `gorm:"column:orderId;primary_key;AUTO_INCREMENT" json:"orderId"` // 订单id
	OrgId      int    `gorm:"column:orgId" json:"orgId"`                                // 项目id
	UserId     int64  `gorm:"column:userId" json:"userId"`                              // 用户id
	OrderType  int    `gorm:"column:orderType" json:"orderType"`                        // 订单类型 1:合约 2:otc ..
	OrderNo    int64  `gorm:"column:orderNo" json:"orderNo"`                            // 订单号
	Status     int    `gorm:"column:status" json:"status"`                              // 状态
	Remark     string `gorm:"column:remark" json:"remark"`                              // 备注
	CreateTime int64  `gorm:"column:createTime" json:"createTime"`                      // 创建时间
	UpdateTime int64  `gorm:"column:updateTime" json:"updateTime"`                      // 更新时间
	ExtStr1    string `gorm:"column:extStr1" json:"extStr1"`                            // 扩展字段STR1 255
	ExtStr2    string `gorm:"column:extStr2" json:"extStr2"`                            // 扩展字段STR2 255
	ExtStr3    string `gorm:"column:extStr3" json:"extStr3"`                            // 扩展字段STR3 255
	ExtStr4    string `gorm:"column:extStr4" json:"extStr4"`                            // 扩展字段STR4 255
	ExtStr5    string `gorm:"column:extStr5" json:"extStr5"`                            // 扩展字段STR5 255
	ExtStr6    string `gorm:"column:extStr6" json:"extStr6"`                            // 扩展字段STR6 255
	ExtStr7    string `gorm:"column:extStr7" json:"extStr7"`                            // 扩展字段STR7 255
	ExtStr8    string `gorm:"column:extStr8" json:"extStr8"`                            // 扩展字段STR8 255
	ExtStr9    string `gorm:"column:extStr9" json:"extStr9"`                            // 扩展字段STR9 255
	ExtStr10   string `gorm:"column:extStr10" json:"extStr10"`                          // 扩展字段STR10 255
	ExtStr11   string `gorm:"column:extStr11" json:"extStr11"`                          // 扩展字段STR11 255
	ExtStr12   string `gorm:"column:extStr12" json:"extStr12"`                          // 扩展字段STR12 255
	ExtStr13   string `gorm:"column:extStr13" json:"extStr13"`                          // 扩展字段STR13 255
	ExtStr14   string `gorm:"column:extStr14" json:"extStr14"`                          // 扩展字段STR14 255
	ExtStr15   string `gorm:"column:extStr15" json:"extStr15"`                          // 扩展字段STR15 255
	ExtStr16   string `gorm:"column:extStr16" json:"extStr16"`                          // 扩展字段STR16 255
	ExtStr17   string `gorm:"column:extStr17" json:"extStr17"`                          // 扩展字段STR17 255
	ExtStr18   string `gorm:"column:extStr18" json:"extStr18"`                          // 扩展字段STR18 255
	ExtStr19   string `gorm:"column:extStr19" json:"extStr19"`                          // 扩展字段STR19 255
	ExtStr20   string `gorm:"column:extStr20" json:"extStr20"`                          // 扩展字段STR20 255
	ExtStr21   string `gorm:"column:extStr21" json:"extStr21"`                          // 扩展字段STR21 255
	ExtStr22   string `gorm:"column:extStr22" json:"extStr22"`                          // 扩展字段STR22 255
	ExtStr23   string `gorm:"column:extStr23" json:"extStr23"`                          // 扩展字段STR23 255
	ExtStr24   string `gorm:"column:extStr24" json:"extStr24"`                          // 扩展字段STR24 255
	ExtStr25   string `gorm:"column:extStr25" json:"extStr25"`                          // 扩展字段STR25 255
	ExtStr26   string `gorm:"column:extStr26" json:"extStr26"`                          // 扩展字段STR25 255
	ExtStr27   string `gorm:"column:extStr27" json:"extStr27"`                          // 扩展字段STR27 255
	ExtStr28   string `gorm:"column:extStr28" json:"extStr28"`                          // 扩展字段STR28 255
	ExtStr29   string `gorm:"column:extStr29" json:"extStr29"`                          // 扩展字段STR29 255
	ExtStr30   string `gorm:"column:extStr30" json:"extStr30"`                          // 扩展字段STR30 255
	ExtBigStr1 string `gorm:"column:extBigStr1" json:"extBigStr1"`                      // 扩展字段STR30 1024
	ExtBigStr2 string `gorm:"column:extBigStr2" json:"extBigStr2"`                      // 扩展字段STR30 1024
	ExtBigStr3 string `gorm:"column:extBigStr3" json:"extBigStr3"`                      // 扩展字段STR30 1024
	ExtBigStr4 string `gorm:"column:extBigStr4" json:"extBigStr4"`                      // 扩展字段STR30 1024
	ExtBigStr5 string `gorm:"column:extBigStr5" json:"extBigStr5"`                      // 扩展字段STR30 1024
	ExtText1   string `gorm:"column:extText1" json:"extText1"`                          // 扩展字段Text1
	ExtText2   string `gorm:"column:extText2" json:"extText1"`                          // 扩展字段Text2
	ExtInt1    int64  `gorm:"column:extInt1" json:"extInt1"`                            // 扩展字段INT1
	ExtInt2    int64  `gorm:"column:extInt2" json:"extInt2"`                            // 扩展字段INT2
	ExtInt3    int64  `gorm:"column:extInt3" json:"extInt3"`                            // 扩展字段INT3
	ExtInt4    int64  `gorm:"column:extInt4" json:"extInt4"`                            // 扩展字段INT4
	ExtInt5    int64  `gorm:"column:extInt5" json:"extInt5"`                            // 扩展字段INT5
	ExtInt6    int64  `gorm:"column:extInt6" json:"extInt6"`                            // 扩展字段INT6
	ExtInt7    int64  `gorm:"column:extInt7" json:"extInt7"`                            // 扩展字段INT7
	ExtInt8    int64  `gorm:"column:extInt8" json:"extInt8"`                            // 扩展字段INT8
	ExtInt9    int64  `gorm:"column:extInt9" json:"extInt9"`                            // 扩展字段INT9
	ExtInt10   int64  `gorm:"column:extInt10" json:"extInt10"`                          // 扩展字段INT10
	ExtInt11   int64  `gorm:"column:extInt11" json:"extInt11"`                          // 扩展字段INT11
	ExtInt12   int64  `gorm:"column:extInt12" json:"extInt12"`                          // 扩展字段INT12
	ExtInt13   int64  `gorm:"column:extInt13" json:"extInt13"`                          // 扩展字段INT13
	ExtInt14   int64  `gorm:"column:extInt14" json:"extInt14"`                          // 扩展字段INT14
	ExtInt15   int64  `gorm:"column:extInt15" json:"extInt15"`                          // 扩展字段INT15
	ExtInt16   int64  `gorm:"column:extInt16" json:"extInt16"`                          // 扩展字段INT16
	ExtInt17   int64  `gorm:"column:extInt17" json:"extInt17"`                          // 扩展字段INT17
	ExtInt18   int64  `gorm:"column:extInt18" json:"extInt18"`                          // 扩展字段INT18
	ExtInt19   int64  `gorm:"column:extInt19" json:"extInt19"`                          // 扩展字段INT19
	ExtInt20   int64  `gorm:"column:extInt20" json:"extInt20"`                          // 扩展字段INT20
}

//创建订单
type CreateOrder struct {
	Order          *Order               `json:"order"`
	OperateAmounts []*TaskOperateAmount `json:"operateAmounts"`
	BeforeCheck    *BeforeCheck         `json:"beforeCheck"`
}

//更新订单
type UpdateOrder struct {
	Order          *Order
	OperateAmounts []*TaskOperateAmount `json:"operateAmounts"`
	BeforeCheck    *BeforeCheck         `json:"beforeCheck"`
	AfterAction    *AfterAction         `json:"afterAction"`
}

//更新,新增订单
type SaveOrUpdateOrder struct {
	SaveOrders   []*CreateOrder `json:"saveOrders"`
	UpdateOrders []*UpdateOrder `json:"updateOrders"`
}

//查询订单
type FindOrder struct {
	Order *Order `json:"order"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}

//查询订单结果
type FindOrderRs struct {
	FindParams *FindOrder `json:"findParams"`
	Orders     []*Order   `json:"order"`
}

//创建/更新 前置检查
type BeforeCheck struct {
	Status int `json:"status"`
}

//后置操作
type AfterAction struct {
	CreateOrders []*CreateOrder `json:"createOrder"`
}

//自定义查询参数
type FindByCustomParams struct {
	Select string        `json:"select"`
	Where  string        `json:"where"`
	Params []interface{} `json:"params"`
	Order  string        `json:"order"`
	Group  string        `json:"group"`
	Having string        `json:"having"`
	Result interface{}   `json:"result"`
}

//账户操作
type TaskOperateAmount struct {
	OpType        int    `json:"opType"`
	BsType        int    `json:"bsType"`
	AccountId     int64  `json:"accountId"`
	AllowNegative int    `json:"allowNegative"`
	Amount        string `json:"amount"`
	UserId        int64  `json:"userId"`
	Currency      string `json:"currency"`
	AccountType   int    `json:"accountType"`
	Detail        string `json:"detail"`
	Ext           string `json:"ext"`
}

//创建订单
func Create(cos []*CreateOrder) (map[string]bool, *base_server_sdk.Error) {
	request := map[string]string{}
	t, _ := json.Marshal(cos)
	request["orders"] = string(t)

	client := base_server_sdk.Instance
	response, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "create", request)
	if err != nil {
		return nil, err
	}

	var rs map[string]bool

	err1 := json.Unmarshal(response, &rs)
	if err1 != nil {
		fmt.Println(err1.Error())
		return nil, base_server_sdk.ErrServiceBusy
	}

	return rs, nil
}

//更新订单
func Update(uos []*UpdateOrder) (map[string]bool, *base_server_sdk.Error) {
	request := map[string]string{}
	t, _ := json.Marshal(uos)
	request["orders"] = string(t)

	client := base_server_sdk.Instance
	response, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "update", request)
	if err != nil {
		return nil, err
	}

	var rs map[string]bool

	err1 := json.Unmarshal(response, &rs)
	if err1 != nil {
		fmt.Println(err1.Error())
		return nil, base_server_sdk.ErrServiceBusy
	}

	return rs, nil
}

//创建/更新
func CreateUpdateBoth(suo  []*SaveOrUpdateOrder) (map[string]bool, *base_server_sdk.Error) {
	request := map[string]string{}
	t, _ := json.Marshal(suo)
	request["orders"] = string(t)

	client := base_server_sdk.Instance
	response, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "createUpdateBoth", request)
	if err != nil {
		return nil, err
	}

	var rs map[string]bool

	err1 := json.Unmarshal(response, &rs)
	if err1 != nil {
		fmt.Println(err1.Error())
		return nil, base_server_sdk.ErrServiceBusy
	}

	return rs, nil
}

//查询订单
func Find(fo []*FindOrder) ([]*FindOrderRs, *base_server_sdk.Error) {
	request := map[string]string{}
	t, _ := json.Marshal(fo)
	request["orders"] = string(t)

	client := base_server_sdk.Instance
	response, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "find", request)
	if err != nil {
		return nil, err
	}

	var data []*FindOrderRs

	err1 := json.Unmarshal(response, &data)
	if err1 != nil {
		fmt.Println(err1.Error())
		return nil, base_server_sdk.ErrServiceBusy
	}

	return data, nil
}

//自定义查询
func FindByCustom(fcp *FindByCustomParams) ([]map[string]interface{}, *base_server_sdk.Error) {
	request := map[string]string{}
	t, _ := json.Marshal(fcp)
	request["params"] = string(t)

	client := base_server_sdk.Instance
	response, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "findByCustom", request)
	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}

	err1 := json.Unmarshal(response, &data)
	if err1 != nil {
		fmt.Println(err1.Error())
		return nil, base_server_sdk.ErrServiceBusy
	}

	return data, nil
}
