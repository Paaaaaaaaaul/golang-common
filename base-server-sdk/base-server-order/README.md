# base_server_order 接口说明文档

## 初始化base_server_sdk
```go
base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		AppId:           "10002",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			OrderServerHost: "http://localhost:8083",
		},
})


// ....

defer base_server_sdk.ReleaseBaseServerSdk()
```

## 相关model
```go
//业务类型
type BusinessType int

const (
	//商城
	ORDER_TYPE_MALL = BusinessType(1)
	//游戏
	ORDER_TYPE_GAME = BusinessType(2)
	//其他
	ORDER_TYPE_OTHER = BusinessType(100)
)

//订单详情参数
type OrderDetailParam struct {
	SellerId       int64  //卖家ID
	SellerName     string //卖家名称
	GoodsId        int64  //商品Id
	GoodsInfo      string //商品信息
	GoodsSkuId     int64  //skuId
	GoodsSkuInfo   string //sku信息
	Quantity       string //数量
	UnitPrice      string //单价
	TotalPrice     string //总价
	BusinessStatus int    //业务状态
	Remark         string //备注
	Extra          string //额外字段
}

//订单参数
type OrderParam struct {
	BusinessType       BusinessType        //业务状态
	BuyerId            int64               //买家id
	BuyerName          string              //买家名称
	PaymentCurrency    string              //支付货币,CNY人民币
	TotalPrice         string              //总价
	PayableTotalPrice  string              //应付金额
	Remark             string              //备注
	OrderDetailParams  []*OrderDetailParam //订单详情
	IsSubStock         int                 //是否减库存
	NeedPay            int                 //是否需要支付
	NeedAutoCancel     int                 //是否需要自动取消
	AutoCancelSeconds  int                 //自动取消等待秒数
	DeliverType        int                 //发货类型
	NeedAutoDeliver    int                 //是否需要自动发货
	NeedAutoConfirm    int                 //是否超时确定收货
	AutoConfirmSeconds int                 //是否超时确定收货
	InitStatus         int                 //初始状态,如果需要跳过订单步骤,可以设置此字段
}

type OrderStatus int

//订单状态
//1待付款、2待发货、3已发货、4已收货、5已取消、9超时取消、10超时确认收货
const (
	//待支付,订单的初始状态
	ORDER_STATUS_WAIT_PAY = OrderStatus(1)
	//已支付,待发货
	ORDER_STATUS_WAIT_DELIVERY = OrderStatus(2)
	//已发货
	ORDER_STATUS_DELIVERED = OrderStatus(3)
	//已收货
	ORDER_STATUS_RECEIPT = OrderStatus(4)
	//取消订单:未支付手动取消,未支付超时取消
	ORDER_STATUS_CANCELED = OrderStatus(5)
	//超时取消
	ORDER_EXPIRE_CANCELED = OrderStatus(9)
)

type PaymentMethod string

//支付方式
const (
	//余额支付
	BALANCE_PAYMENT = PaymentMethod("balance_payment")
	//三方支付
	THIRD_PARTY_PAYMENT = PaymentMethod("third_party_payment")
	//代付
	OTHERS_HELP_PAYMENT = PaymentMethod("others_help_payment")
)

//订单
type Order struct {
	OrderId                int64         `json:"orderId"`    //主键递增
	BusinessType           int           `json:"businessType"`                     //业务类型
	OrgId                  int           `json:"orgId"`                                   //组织编号
	OrderNo                int64         `json:"orderNo,string"`                        //订单号
	BuyerId                int64         `json:"buyerId"`                               //买家ID
	BuyerName              string        `json:"buyerName"`                           //买家名称
	TotalPrice             string        `json:"totalPrice"`                         //总价(优惠+卡券+账户余额支付+三方支付)
	DiscountPrice          string        `json:"discountPrice"`                   //优惠金额
	PayableTotalPrice      string        `json:"payableTotalPrice"`           //实付价格(卡券+账户余额支付+三方支付)
	BalancePaymentPrice    string        `json:"balancePaymentPrice"`       //余额支付金额
	ThirdPartyPaymentPrice string        `json:"thirdPartyPaymentPrice"` //三方支付金额
	CardTicketPaymentPrice string        `json:"cardTicketPaymentPrice"` //卡券支付金额
	PaymentMethod          PaymentMethod `json:"paymentMethod"`                   //支付方式
	PaymentCurrency        string        `json:"paymentCurrency"`               //支付货币
	Status                 OrderStatus   `json:"status"`                                 //状态
	CreateTime             int64         `json:"createTime"`                         //创建时间
	PayTime                int64         `json:"payTime"`                               //支付时间
	PayOrderNo             int64         `json:"payOrderNo,string"`                  //支付订单号
	CancelTime             int64         `json:"cancelTime"`                         //取消时间
	IsSubStock             int           `json:"isSubStock"`                         //是否减库存
	DeliverOrderNo         int64         `json:"deliverOrderNo,string"`          //取消时间
	DeliverTime            int64         `json:"deliverTime"`                       //发货时间
	DeliverType            int           `json:"deliverType"`                       //发货类型
	ReceiptTime            int64         `json:"receiptTime"`                       //确认收货时间
	NeedAutoCancel         int           `json:"needAutoCancel"`                 //是否自动超时取消
	AutoCancelSeconds      int           `json:"autoCancelSeconds"`           //超时取消等待秒数
	NeedAutoDeliver        int           `json:"needAutoDeliver"`               //是否自动发货
	NeedAutoConfirm        int           `json:"needAutoConfirm"`               //是否自动确认收货
	AutoConfirmSeconds     int           `json:"AutoConfirmSeconds"`         //超时自动确认秒数
	Remark                 string        `json:"remark"`                                 //备注
	Details                []OrderDetail `json:"details"`                                                     //订单详情
}

//订单明细状态
type OrderDetailStatus int

//6退款中、7取消退款、8已退款
const (
	//已收货
	DETAIL_STATUS_RECEIPT = OrderDetailStatus(5)
	//退款中
	DETAIL_STATUS_REFUNDING = OrderDetailStatus(6)
	//已退款
	DETAIL_STATUS_REFUNDED = OrderDetailStatus(8)
)

//订单详情
type OrderDetail struct {
	OrderId        int64             `json:"orderId"` //主键
	OrgId          int               `json:"orgId"`                                //组织id
	OrderNo        int64             `json:"orderNo,string"`                     //订单号
	MergeOrderNo   int64             `json:"mergeOrderNo,string"`           //合并订单号
	BuyerId        int64             `json:"buyerId"`                            //买方ID
	BuyerName      string            `json:"buyerName"`                        //买方名称
	SellerId       int64             `json:"sellerId"`                          //卖方id
	SellerName     string            `json:"sellerName"`                      //卖方名称
	GoodsId        int64             `json:"goodsId"`                            //商品ID
	GoodsInfo      string            `json:"goodsInfo"`                        //商品信息
	GoodsSkuId     int64             `json:"goodsSkuId"`                      //shuId
	GoodsSkuInfo   string            `json:"goodsSkuInfo"`                  //sku信息
	UnitPrice      string            `json:"unitPrice"`                        //单价
	Quantity       string            `json:"quantity"`                          //数量
	TotalPrice     string            `json:"totalPrice"`                      //总价
	ServiceFee     string            `json:"serviceFee"`                      //服务费
	Status         OrderDetailStatus `json:"status"`                              //状态
	BusinessStatus int               `json:"businessStatus"`              //业务状态
	Remark         string            `json:"remark"`                              //备注
	Extra          string            `json:"extra"`                                //扩展字段
	CreateTime     int64             `json:"createTime"`                      //创建时间
	RefundTime     int64             `json:"refundTime"`                      //退款成功时间
	RefundOrderNo  int64             `json:"refundOrderNo,string"`         //退款订单号
}

type DeliverStatus int

//状态（1待发货 2已发货 3已签收 4发货失败）
const (
	WAIT_DELIVER = DeliverStatus(1)
	DELIVERED    = DeliverStatus(2)
	RECEIPTED    = DeliverStatus(3)
	DELIVER_FAIL = DeliverStatus(4)
)

//发货单
type DeliverOrder struct {
	DeliverId    int64  `json:"deliverId"` //主键递增
	OrderNo      int64  `json:"orderNo,string"`                         //发货单号
	GoodsOrderNo int64  `json:"goodsOrderNo,string"`               //商品单号
	DeliverType  int    `json:"deliverType"`                        //发货类型
	Status       int    `json:"status"`                                  //状态
	DeliverInfos string `json:"deliverInfos"`                      //发货信息
	CreateTime   int64  `json:"createTime"`                          //创建时间
}
```

## 相关错误码
```go
	SERVICE_BUSY = &exception.Exception{
		Code:    1000,
		Message: "服务繁忙",
	}

	DB_ERROR = &exception.Exception{
		Code:    1002,
		Message: "数据库异常",
	}

	STOCK_ERROR = &exception.Exception{
		Code:    2001,
		Message: "库存异常",
	}
	ORDER_STATUS_ERROR = &exception.Exception{
		Code:    2002,
		Message: "订单状态错误",
	}
	GOODS_OFF_SHELF = &exception.Exception{
		Code:    2003,
		Message: "商品已下架",
	}
	PAY_ERROR = &exception.Exception{
		Code:    2004,
		Message: "支付异常",
	}

	NOT_SUPPORT_PAYMENT_METHOD = &exception.Exception{
		Code:    2005,
		Message: "不支持的支付方式",
	}

	ACCOUNT_ERROR = &exception.Exception{
		Code:    2006,
		Message: "账户操作错误",
	}

	SIGN_ERROR = &exception.Exception{
		Code:    2007,
		Message: "签名错误",
	}

	BUSINESS_STATUS_ERROR = &exception.Exception{
		Code:    2008,
		Message: "业务状态错误",
	}

	DELIVER_STATUS_ERROR = &exception.Exception{
		Code:    2010,
		Message: "发货状态错误",
	}
	PERMISSION_DENIED_ERROR = &exception.Exception{
		Code:    2011,
		Message: "权限错误",
	}
```


## 接口说明

- 创建订单

```
func CreateOrder(orgId int, orderParams *OrderParam) (*Order, *base_server_sdk.Error)
```

```go
1. orgId必须大于0
2. orderParams参考OrderParam结构体注释

异常返回:
1000 服务繁忙
1001 参数异常
```

- 获取支付方式

func GetPayMethods() ([]PaymentMethod, *base_server_sdk.Error)

```go

异常返回:
1000 服务繁忙

```

- 获取三方支付方式

func GetThirdPartyPayMethods(orgId int, userId int64, orderNo int64) ([]base_server_pay.PayMethod, *base_server_sdk.Error)

```go
1. userId,买家用户ID
2. orderNo,createOrder接口返回的订单号

异常返回:
1000 服务繁忙
2002 订单状态错误
2004 支付接口异常
1002 数据异常
2011 权限错误

```

- 获取三方支付通道

func GetThirdPartyPayChannels(orgId int, userId int64, orderNo int64, methodCode string) ([]base_server_pay.PayChannel, *base_server_sdk.Error)

```go
1. userId,买家用户ID
2. orderNo,createOrder接口返回的订单号
3. methodCode,GetThirdPartyPayMethods接口返回的methodCode字段

异常返回:
1000 服务繁忙
2002 订单状态错误
2004 支付接口异常
1002 数据异常
2011 权限错误
```

- 获取订单信息

func GetOrderInfo(orgId int, userId int64, orderNo int64) (*Order, *base_server_sdk.Error)

```go
1. userId,买家用户ID
2. orderNo,订单号

异常返回:
1002 数据异常
2011 权限错误
```

- 获取用户订单
func GetUserOrders(orgId int, userId int64, pageNo int, pageSize int) ([]Order, *base_server_sdk.Error)

```go
1. userId,买家用户ID
2. orderNo,订单号
3. pageNo, 页码,0开始
4. pageSize 页大小

异常返回:
1002 数据异常
2011 权限错误
```

- 三方支付下单

func ThirdPartyPay(orgId int, userId int64, orderNo int64, methodCode string, channelId string) (*base_server_pay.PayOrder, *base_server_sdk.Error)

```go
1. userId,买家用户ID
2. orderNo,订单号
2. methodCode,三方支付方式查询接口返回
2. channelId,三方支付接口通道查询接口返回

异常返回:
1000 服务繁忙
2002 订单状态错误
2004 支付接口异常
1002 数据异常
2011 权限错误
```

- 余额支付

func BalancePay(orgId int, userId int64, orderNo int64) *base_server_sdk.Error

```go
1. userId,买家用户ID
2. orderNo,订单号

异常返回:
1000 服务繁忙
2002 订单状态错误
2006 账号操作异常
1002 数据异常
2011 权限错误
```

- 发货

func DeliverOrderGoods(orgId int, userId int64, orderNo int64, deliverInfo string) *base_server_sdk.Error
```go
1. userId,买家用户ID
2. orderNo,订单号
3. deliverInfo,发货信息

异常返回:
1000 服务繁忙
2002 订单状态错误
1002 数据异常
2011 权限错误
```

- 获取发货信息

func GetDeliverOrderByGoodsOrder(orgId int, userId int64, orderNo int64) (*DeliverOrder, *base_server_sdk.Error)

```go
1. userId,买家用户ID
2. orderNo,订单号

异常返回:
1000 服务繁忙
2002 订单状态错误
1002 数据异常
2011 权限错误
```

- 修改发货信息

func ModifyDeliverInfo(orgId int, userId int64, orderNo int64, deliverInfo string) *base_server_sdk.Error

```go
1. userId,买家用户ID
2. orderNo,订单号
3. deliverInfo,发货信息

异常返回:
1000 服务繁忙
2002 订单状态错误
1002 数据异常
2011 权限错误
```

- 确认收货

func ReceiptOrderGoods(orgId int, userId int64, orderNo int64) *base_server_sdk.Error

```go
1. userId,买家用户ID
2. orderNo,订单号

异常返回:
1000 服务繁忙
2002 订单状态错误
1002 数据异常
2011 权限错误
```

- 取消发货

func CancelDeliverOrder(orgId int, userId int64, orderNo int64) *base_server_sdk.Error

```go
1. userId,买家用户ID
2. orderNo,订单号

异常返回:
1000 服务繁忙
2002 订单状态错误
1002 数据异常
2011 权限错误
```

- 业务更新

func BusinessUpdate(orgId int, userId int64, orderNo int64, detailOrderNo int64, currentBusinessStatus int, nextBusinessStatus int, extra string,
	accountOps []*base_server_account.TaskDetail) *base_server_sdk.Error

```go
1. userId,买家用户ID
2. orderNo,订单号
3. detailOrderNo,详情单号
4. currentBusinessStatus,当前业务状态
5. nextBusinessStatus,下一个业务状态
6. extra,额外的业务信息，通常用json字符串
7. accountOps,业务更新伴随的账户操作

异常返回:
1000 服务繁忙
2002 订单状态错误
1002 数据异常
2011 权限错误
2008 业务状态错误
```