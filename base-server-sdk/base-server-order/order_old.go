package base_server_order
//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/becent/golang-common/base-server-sdk"
//	"github.com/becent/golang-common/base-server-sdk/base-server-account"
//	"github.com/becent/golang-common/base-server-sdk/base-server-pay"
//	"strconv"
//)
//
////业务类型
//type BusinessType int
//
//const (
//	//商城
//	ORDER_TYPE_MALL = BusinessType(1)
//	//游戏
//	ORDER_TYPE_GAME = BusinessType(2)
//	//其他
//	ORDER_TYPE_OTHER = BusinessType(100)
//)
//
////订单详情参数
//type OrderDetailParam struct {
//	SellerId       int64  //卖家ID
//	SellerName     string //卖家名称
//	GoodsId        int64  //商品Id
//	GoodsInfo      string //商品信息
//	GoodsSkuId     int64  //skuId
//	GoodsSkuInfo   string //sku信息
//	Quantity       string //数量
//	UnitPrice      string //单价
//	TotalPrice     string //总价
//	BusinessStatus int    //业务状态
//	Remark         string //备注
//	Extra          string //额外字段
//}
//
////订单参数
//type OrderParam struct {
//	BusinessType       BusinessType        //业务状态
//	BuyerId            int64               //买家id
//	BuyerName          string              //买家名称
//	PaymentCurrency    string              //支付货币,CNY人民币
//	TotalPrice         string              //总价
//	PayableTotalPrice  string              //应付金额
//	Remark             string              //备注
//	OrderDetailParams  []*OrderDetailParam //订单详情
//	IsSubStock         int                 //是否减库存
//	NeedPay            int                 //是否需要支付
//	NeedAutoCancel     int                 //是否需要自动取消
//	AutoCancelSeconds  int                 //自动取消等待秒数
//	DeliverType        int                 //发货类型
//	NeedAutoDeliver    int                 //是否需要自动发货
//	NeedAutoConfirm    int                 //是否超时确定收货
//	AutoConfirmSeconds int                 //是否超时确定收货
//	InitStatus         int                 //初始状态,如果需要跳过订单步骤,可以设置此字段
//}
//
//type OrderStatus int
//
////订单状态
////1待付款、2待发货、3已发货、4已收货、5已取消、9超时取消、10超时确认收货
//const (
//	//待支付,订单的初始状态
//	ORDER_STATUS_WAIT_PAY = OrderStatus(1)
//	//已支付,待发货
//	ORDER_STATUS_WAIT_DELIVERY = OrderStatus(2)
//	//已发货
//	ORDER_STATUS_DELIVERED = OrderStatus(3)
//	//已收货
//	ORDER_STATUS_RECEIPT = OrderStatus(4)
//	//取消订单:未支付手动取消,未支付超时取消
//	ORDER_STATUS_CANCELED = OrderStatus(5)
//	//超时取消
//	ORDER_EXPIRE_CANCELED = OrderStatus(9)
//)
//
//type PaymentMethod string
//
////支付方式
//const (
//	//余额支付
//	BALANCE_PAYMENT = PaymentMethod("balance_payment")
//	//三方支付
//	THIRD_PARTY_PAYMENT = PaymentMethod("third_party_payment")
//	//代付
//	OTHERS_HELP_PAYMENT = PaymentMethod("others_help_payment")
//)
//
////订单
//type Order struct {
//	OrderId                int64         `gorm:"column:orderId;primary_key;AUTO_INCREMENT" json:"orderId"`    //主键递增
//	BusinessType           int           `gorm:"column:businessType" json:"businessType"`                     //业务类型
//	OrgId                  int           `gorm:"column:orgId" json:"orgId"`                                   //组织编号
//	OrderNo                int64         `gorm:"column:orderNo" json:"orderNo,string"`                        //订单号
//	BuyerId                int64         `gorm:"column:buyerId" json:"buyerId"`                               //买家ID
//	BuyerName              string        `gorm:"column:buyerName" json:"buyerName"`                           //买家名称
//	TotalPrice             string        `gorm:"column:totalPrice" json:"totalPrice"`                         //总价(优惠+卡券+账户余额支付+三方支付)
//	DiscountPrice          string        `gorm:"column:discountPrice" json:"discountPrice"`                   //优惠金额
//	PayableTotalPrice      string        `gorm:"column:payableTotalPrice" json:"payableTotalPrice"`           //实付价格(卡券+账户余额支付+三方支付)
//	BalancePaymentPrice    string        `gorm:"column:balancePaymentPrice" json:"balancePaymentPrice"`       //余额支付金额
//	ThirdPartyPaymentPrice string        `gorm:"column:thirdPartyPaymentPrice" json:"thirdPartyPaymentPrice"` //三方支付金额
//	CardTicketPaymentPrice string        `gorm:"column:cardTicketPaymentPrice" json:"cardTicketPaymentPrice"` //卡券支付金额
//	PaymentMethod          PaymentMethod `gorm:"column:paymentMethod" json:"paymentMethod"`                   //支付方式
//	PaymentCurrency        string        `gorm:"column:paymentCurrency" json:"paymentCurrency"`               //支付货币
//	Status                 OrderStatus   `gorm:"column:status" json:"status"`                                 //状态
//	CreateTime             int64         `gorm:"column:createTime" json:"createTime"`                         //创建时间
//	PayTime                int64         `gorm:"column:payTime" json:"payTime"`                               //支付时间
//	PayOrderNo             int64         `gorm:"column:payOrderNo" json:"payOrderNo,string"`                  //支付订单号
//	CancelTime             int64         `gorm:"column:cancelTime" json:"cancelTime"`                         //取消时间
//	IsSubStock             int           `gorm:"column:isSubStock" json:"isSubStock"`                         //是否减库存
//	DeliverOrderNo         int64         `gorm:"column:deliverOrderNo" json:"deliverOrderNo,string"`          //取消时间
//	DeliverTime            int64         `gorm:"column:deliverTime" json:"deliverTime"`                       //发货时间
//	DeliverType            int           `gorm:"column:deliverType" json:"deliverType"`                       //发货类型
//	ReceiptTime            int64         `gorm:"column:receiptTime" json:"receiptTime"`                       //确认收货时间
//	NeedAutoCancel         int           `gorm:"column:needAutoCancel" json:"needAutoCancel"`                 //是否自动超时取消
//	AutoCancelSeconds      int           `gorm:"column:autoCancelSeconds" json:"autoCancelSeconds"`           //超时取消等待秒数
//	NeedAutoDeliver        int           `gorm:"column:needAutoDeliver" json:"needAutoDeliver"`               //是否自动发货
//	NeedAutoConfirm        int           `gorm:"column:needAutoConfirm" json:"needAutoConfirm"`               //是否自动确认收货
//	AutoConfirmSeconds     int           `gorm:"column:AutoConfirmSeconds" json:"AutoConfirmSeconds"`         //超时自动确认秒数
//	Remark                 string        `gorm:"column:remark" json:"remark"`                                 //备注
//	Details                []OrderDetail `json:"details"`                                                     //订单详情
//}
//
////订单明细状态
//type OrderDetailStatus int
//
////6退款中、7取消退款、8已退款
//const (
//	//已收货
//	DETAIL_STATUS_RECEIPT = OrderDetailStatus(5)
//	//退款中
//	DETAIL_STATUS_REFUNDING = OrderDetailStatus(6)
//	//已退款
//	DETAIL_STATUS_REFUNDED = OrderDetailStatus(8)
//)
//
////订单详情
//type OrderDetail struct {
//	OrderId        int64             `gorm:"column:orderId;primary_key;AUTO_INCREMENT" json:"orderId"` //主键
//	OrgId          int               `gorm:"column:orgId" json:"orgId"`                                //组织id
//	OrderNo        int64             `gorm:"column:orderNo" json:"orderNo,string"`                     //订单号
//	MergeOrderNo   int64             `gorm:"column:mergeOrderNo" json:"mergeOrderNo,string"`           //合并订单号
//	BuyerId        int64             `gorm:"column:buyerId" json:"buyerId"`                            //买方ID
//	BuyerName      string            `gorm:"column:buyerName" json:"buyerName"`                        //买方名称
//	SellerId       int64             `gorm:"column:sellerId" json:"sellerId"`                          //卖方id
//	SellerName     string            `gorm:"column:sellerName" json:"sellerName"`                      //卖方名称
//	GoodsId        int64             `gorm:"column:goodsId" json:"goodsId"`                            //商品ID
//	GoodsInfo      string            `gorm:"column:goodsInfo" json:"goodsInfo"`                        //商品信息
//	GoodsSkuId     int64             `gorm:"column:goodsSkuId" json:"goodsSkuId"`                      //shuId
//	GoodsSkuInfo   string            `gorm:"column:goodsSkuInfo" json:"goodsSkuInfo"`                  //sku信息
//	UnitPrice      string            `gorm:"column:unitPrice" json:"unitPrice"`                        //单价
//	Quantity       string            `gorm:"column:quantity" json:"quantity"`                          //数量
//	TotalPrice     string            `gorm:"column:totalPrice" json:"totalPrice"`                      //总价
//	ServiceFee     string            `gorm:"column:serviceFee" json:"serviceFee"`                      //服务费
//	Status         OrderDetailStatus `gorm:"column:status" json:"status"`                              //状态
//	BusinessStatus int               `gorm:"column:businessStatus" json:"businessStatus"`              //业务状态
//	Remark         string            `gorm:"column:remark" json:"remark"`                              //备注
//	Extra          string            `gorm:"column:extra" json:"extra"`                                //扩展字段
//	CreateTime     int64             `gorm:"column:createTime" json:"createTime"`                      //创建时间
//	RefundTime     int64             `gorm:"column:refundTime" json:"refundTime"`                      //退款成功时间
//	RefundOrderNo  int64             `gorm:"column:refundOrderNo" json:"refundOrderNo,string"`         //退款订单号
//}
//
//type DeliverStatus int
//
////状态（1待发货 2已发货 3已签收 4发货失败）
//const (
//	WAIT_DELIVER = DeliverStatus(1)
//	DELIVERED    = DeliverStatus(2)
//	RECEIPTED    = DeliverStatus(3)
//	DELIVER_FAIL = DeliverStatus(4)
//)
//
////发货单
//type DeliverOrder struct {
//	DeliverId    int64  `gorm:"column:deliverId;primary_key;AUTO_INCREMENT" json:"deliverId"` //主键递增
//	OrderNo      int64  `gorm:"column:orderNo" json:"orderNo,string"`                         //发货单号
//	GoodsOrderNo int64  `gorm:"column:goodsOrderNo" json:"goodsOrderNo,string"`               //商品单号
//	DeliverType  int    `gorm:"column:deliverType" json:"deliverType"`                        //发货类型
//	Status       int    `gorm:"column:status" json:"status"`                                  //状态
//	DeliverInfos string `gorm:"column:deliverInfos" json:"deliverInfos"`                      //发货信息
//	CreateTime   int64  `gorm:"column:createTime" json:"createTime"`                          //创建时间
//}
//
////获取支付方式
//func GetPayMethods() ([]PaymentMethod, *base_server_sdk.Error) {
//	client := base_server_sdk.Instance
//	response, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "GetPayMethods", nil)
//	if err != nil {
//		return nil, err
//	}
//
//	var methods []PaymentMethod
//
//	err1 := json.Unmarshal(response, &methods)
//	if err1 != nil {
//		return nil, base_server_sdk.ErrServiceBusy
//	}
//
//	return methods, nil
//}
//
////获取三方支付方式
//func GetThirdPartyPayMethods(orgId int, userId int64, orderNo int64) ([]base_server_pay.PayMethod, *base_server_sdk.Error) {
//	request := map[string]string{}
//	request["orgId"] = strconv.Itoa(orgId)
//	request["userId"] = strconv.FormatInt(userId, 10)
//	request["orderNo"] = strconv.FormatInt(orderNo, 10)
//
//	client := base_server_sdk.Instance
//	response, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "GetThirdPartyPayMethods", request)
//	if err != nil {
//		return nil, err
//	}
//
//	var methods []base_server_pay.PayMethod
//
//	err1 := json.Unmarshal(response, &methods)
//	if err1 != nil {
//		return nil, base_server_sdk.ErrServiceBusy
//	}
//
//	return methods, nil
//}
//
////获取三方支付通道
//func GetThirdPartyPayChannels(orgId int, userId int64, orderNo int64, methodCode string) ([]base_server_pay.PayChannel, *base_server_sdk.Error) {
//	request := map[string]string{}
//	request["orgId"] = strconv.Itoa(orgId)
//	request["userId"] = strconv.FormatInt(userId, 10)
//	request["orderNo"] = strconv.FormatInt(orderNo, 10)
//	request["methodCode"] = methodCode
//
//	client := base_server_sdk.Instance
//	response, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "GetThirdPartyPayChannels", request)
//	if err != nil {
//		return nil, err
//	}
//
//	var methods []base_server_pay.PayChannel
//
//	err1 := json.Unmarshal(response, &methods)
//	if err1 != nil {
//		return nil, base_server_sdk.ErrServiceBusy
//	}
//
//	return methods, nil
//}
//
////创建订单
//func CreateOrder(orgId int, orderParams *OrderParam) (*Order, *base_server_sdk.Error) {
//	bytes, err := json.Marshal(orderParams)
//	if err != nil {
//		return nil, base_server_sdk.ErrServiceBusy
//	}
//
//	request := map[string]string{}
//	request["orgId"] = strconv.Itoa(orgId)
//	request["orderParams"] = string(bytes)
//
//	client := base_server_sdk.Instance
//	response, err1 := client.DoRequest(client.Hosts.OrderServerHost, "order", "CreateOrder", request)
//	if err != nil {
//		return nil, err1
//	}
//
//	var order Order
//
//	err2 := json.Unmarshal(response, &order)
//	if err2 != nil {
//		return nil, base_server_sdk.ErrServiceBusy
//	}
//
//	return &order, nil
//}
//
////获取订单信息
//func GetOrderInfo(orgId int, userId int64, orderNo int64) (*Order, *base_server_sdk.Error) {
//	request := map[string]string{}
//	request["orgId"] = strconv.Itoa(orgId)
//	request["userId"] = strconv.FormatInt(userId, 10)
//	request["orderNo"] = strconv.FormatInt(orderNo, 10)
//
//	client := base_server_sdk.Instance
//	response, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "GetOrderInfo", request)
//	if err != nil {
//		return nil, err
//	}
//
//	var order Order
//
//	err1 := json.Unmarshal(response, &order)
//	if err1 != nil {
//		fmt.Println(err1.Error())
//		return nil, base_server_sdk.ErrServiceBusy
//	}
//
//	return &order, nil
//}
//
////获取用户订单信息
//func GetUserOrders(orgId int, userId int64, pageNo int, pageSize int) ([]Order, *base_server_sdk.Error) {
//	request := map[string]string{}
//	request["orgId"] = strconv.Itoa(orgId)
//	request["userId"] = strconv.FormatInt(userId, 10)
//	request["pageNo"] = strconv.Itoa(pageNo)
//	request["pageSize"] = strconv.Itoa(pageSize)
//
//	client := base_server_sdk.Instance
//	response, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "GetUserOrders", request)
//	if err != nil {
//		return nil, err
//	}
//
//	var orders []Order
//
//	err1 := json.Unmarshal(response, &orders)
//	if err1 != nil {
//		return nil, base_server_sdk.ErrServiceBusy
//	}
//
//	return orders, nil
//}
//
////三方支付
//func ThirdPartyPay(orgId int, userId int64, orderNo int64, methodCode string, channelId string) (*base_server_pay.PayOrder, *base_server_sdk.Error) {
//	request := map[string]string{}
//	request["orgId"] = strconv.Itoa(orgId)
//	request["userId"] = strconv.FormatInt(userId, 10)
//	request["orderNo"] = strconv.FormatInt(orderNo, 10)
//	request["methodCode"] = methodCode
//	request["channelId"] = channelId
//
//	client := base_server_sdk.Instance
//	response, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "ThirdPartyPay", request)
//	if err != nil {
//		return nil, err
//	}
//
//	var order base_server_pay.PayOrder
//
//	err1 := json.Unmarshal(response, &order)
//	if err1 != nil {
//		return nil, base_server_sdk.ErrServiceBusy
//	}
//
//	return &order, nil
//}
//
////余额支付
//func BalancePay(orgId int, userId int64, orderNo int64) *base_server_sdk.Error {
//	request := map[string]string{}
//	request["orgId"] = strconv.Itoa(orgId)
//	request["userId"] = strconv.FormatInt(userId, 10)
//	request["orderNo"] = strconv.FormatInt(orderNo, 10)
//
//	client := base_server_sdk.Instance
//	_, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "BalancePay", request)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
////发货
//func DeliverOrderGoods(orgId int, userId int64, orderNo int64, deliverInfo string) *base_server_sdk.Error {
//	request := map[string]string{}
//	request["orgId"] = strconv.Itoa(orgId)
//	request["userId"] = strconv.FormatInt(userId, 10)
//	request["orderNo"] = strconv.FormatInt(orderNo, 10)
//	request["deliverInfo"] = deliverInfo
//
//	client := base_server_sdk.Instance
//	_, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "DeliverOrderGoods", request)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
////获取发货单
//func GetDeliverOrderByGoodsOrder(orgId int, userId int64, goodsOrderNo int64) (*DeliverOrder, *base_server_sdk.Error) {
//	request := map[string]string{}
//	request["orgId"] = strconv.Itoa(orgId)
//	request["userId"] = strconv.FormatInt(userId, 10)
//	request["goodsOrderNo"] = strconv.FormatInt(goodsOrderNo, 10)
//
//	client := base_server_sdk.Instance
//	response, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "GetDeliverOrderByGoodsOrder", request)
//	if err != nil {
//		return nil, err
//	}
//	fmt.Println(string(response))
//
//	var order DeliverOrder
//
//	err1 := json.Unmarshal(response, &order)
//	if err1 != nil {
//		return nil, base_server_sdk.ErrServiceBusy
//	}
//
//	return &order, nil
//}
//
////修改发货信息
//func ModifyDeliverInfo(orgId int, userId int64, orderNo int64, deliverInfo string) *base_server_sdk.Error {
//	request := map[string]string{}
//	request["orgId"] = strconv.Itoa(orgId)
//	request["userId"] = strconv.FormatInt(userId, 10)
//	request["orderNo"] = strconv.FormatInt(orderNo, 10)
//	request["deliverInfo"] = deliverInfo
//
//	client := base_server_sdk.Instance
//	_, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "ModifyDeliverInfo", request)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
////确认收货
//func ReceiptOrderGoods(orgId int, userId int64, orderNo int64) *base_server_sdk.Error {
//	request := map[string]string{}
//	request["orgId"] = strconv.Itoa(orgId)
//	request["userId"] = strconv.FormatInt(userId, 10)
//	request["orderNo"] = strconv.FormatInt(orderNo, 10)
//
//	client := base_server_sdk.Instance
//	_, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "ReceiptOrderGoods", request)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
////取消发货
//func CancelDeliverOrder(orgId int, userId int64, orderNo int64) *base_server_sdk.Error {
//	request := map[string]string{}
//	request["orgId"] = strconv.Itoa(orgId)
//	request["userId"] = strconv.FormatInt(userId, 10)
//	request["orderNo"] = strconv.FormatInt(orderNo, 10)
//
//	client := base_server_sdk.Instance
//	_, err := client.DoRequest(client.Hosts.OrderServerHost, "order", "CancelDeliverOrder", request)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
////业务更新
//func BusinessUpdate(orgId int, userId int64, orderNo int64, detailOrderNo int64, currentBusinessStatus int, nextBusinessStatus int, extra string,
//	accountOps []*base_server_account.TaskDetail) *base_server_sdk.Error {
//	accountOpsBytes, err := json.Marshal(accountOps)
//	if err != nil {
//		return base_server_sdk.ErrInvalidParams
//	}
//
//	request := map[string]string{}
//	request["orgId"] = strconv.Itoa(orgId)
//	request["userId"] = strconv.FormatInt(userId, 10)
//	request["orderNo"] = strconv.FormatInt(orderNo, 10)
//
//	request["detailOrderNo"] = strconv.FormatInt(detailOrderNo, 10)
//	request["currentBusinessStatus"] = strconv.Itoa(currentBusinessStatus)
//	request["extra"] = extra
//	request["accountOps"] = string(accountOpsBytes)
//	request["nextBusinessStatus"] = strconv.Itoa(nextBusinessStatus)
//
//	client := base_server_sdk.Instance
//	_, err1 := client.DoRequest(client.Hosts.OrderServerHost, "order", "BusinessUpdate", request)
//	if err1 != nil {
//		return err1
//	}
//
//	return nil
//}
