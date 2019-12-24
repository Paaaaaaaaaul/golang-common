package main

import (
	"fmt"
	"github.com/becent/golang-common/base-server-sdk"
	"github.com/becent/golang-common/base-server-sdk/base-server-order"
	"time"
)

func main() {
	//	od := `{"AutoConfirmSeconds":0,"autoCancelSeconds":20,"balancePaymentPrice":"0","businessType":1,"buyerId":1,"buyerName":"buyerName","cancelTime":0,"cardTicketPaymentPrice":"0","createTime":1576737587,"deliverOrderNo":"0","deliverTime":0,"deliverType":0,"details":[{"businessStatus":1,"buyerId":1,"buyerName":
	//"buyerName","createTime":1576737587,"extra":"extra","goodsId":1,"goodsInfo":"goodsInfo","goodsSkuId":123,"goodsSkuInfo":"skuInfo","mergeOrderNo":"1157673758719523344","orderId":71,"orderNo":"1157673758719623344","orgId":1,"quantity":"10000000000000000000","refundOrderNo":"0","refundTime":0,"remark":
	//"remark","sellerId":1,"sellerName":"sellerName","serviceFee":"0","status":5,"totalPrice":"10000000000000000000","unitPrice":"10000000000000000000"}],"discountPrice":"0","isSubStock":2,"needAutoCancel":2,"needAutoConfirm":2,"needAutoDeliver":2,"orderId":71,"orderNo":"1157673758719523344","orgId":1,"payOrderNo"
	//:"0","payTime":0,"payableTotalPrice":"10000000000000000000","paymentCurrency":"CC","paymentMethod":"","receiptTime":0,"remark":"待支付","status":1,"thirdPartyPaymentPrice":"0","totalPrice":"10000000000000000000"}
	//`
	//	var o base_server_order.Order
	//	err1 := json.Unmarshal([]byte(od), &o)
	//	fmt.Println(err1)
	//	fmt.Println(o)
	//	os.Exit(-1)
	base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		AppId:           "10002",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			OrderServerHost: "http://localhost:8083",
		},
	})
	defer base_server_sdk.ReleaseBaseServerSdk()

	now := time.Now()
	defer func(now time.Time) {
		println(time.Since(now).String())
	}(now)

	response, err := base_server_order.GetPayMethods()
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("获取支付方式成功：[%v]\n", response)
	}

	var orderDetails []*base_server_order.OrderDetailParam
	orderDetails = append(orderDetails, &base_server_order.OrderDetailParam{
		SellerId:       1,
		SellerName:     "sellerName",
		GoodsId:        1,
		GoodsInfo:      "goodsInfo",
		GoodsSkuId:     123,
		GoodsSkuInfo:   "skuInfo",
		Quantity:       "10000000000000000000",
		UnitPrice:      "10000000000000000000",
		TotalPrice:     "10000000000000000000",
		ServiceFee:     "0",
		BusinessStatus: 1,
		Remark:         "remark",
		Extra:          "extra",
	})

	response3, err := base_server_order.CreateOrder(1, &base_server_order.OrderParam{
		BusinessType:       1,
		BuyerId:            1,
		BuyerName:          "buyerName",
		PaymentCurrency:    "CC",
		TotalPrice:         "10000000000000000000",
		PayableTotalPrice:  "10000000000000000000",
		Remark:             "123",
		OrderDetailParams:  orderDetails,
		IsSubStock:         2,
		NeedAutoCancel:     2,
		AutoCancelSeconds:  20,
		NeedAutoDeliver:    2,
		NeedAutoConfirm:    2,
		AutoConfirmSeconds: 2,
	})

	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("创建订单成功：[%v]\n", response3)
	}
	fmt.Println(response3)
	response1, err := base_server_order.GetThirdPartyPayMethods(1, 1, response3.OrderNo)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("获取三方支付方式成功：[%v]\n", response1)
	}

	response2, err := base_server_order.GetThirdPartyPayChannels(1, 1, response3.OrderNo, "alipay_qr_code_cny")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("获取三方支付通道成功：[%v]\n", response2)
	}

	response4, err := base_server_order.GetOrderInfo(1, 1, response3.OrderNo)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("获取订单信息成功：[%v]\n", response4)
	}

	response5, err := base_server_order.GetUserOrders(1, 1, 0, 10)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("获取用户订单成功：[%v]\n", response5)
	}

	response6, err := base_server_order.ThirdPartyPay(1, 1, response3.OrderNo, "alipay_qr_code_cny", "5")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("三方支付下单成功：[%v]\n", response6)
	}

	err = base_server_order.BalancePay(1, 1, response3.OrderNo)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("余额支付成功：[%v]\n", "")
	}

	err = base_server_order.DeliverOrderGoods(1, 1, response3.OrderNo, "deliverInfo")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("发货成功：[%v]\n", "")
	}

	response7, err := base_server_order.GetDeliverOrderByGoodsOrder(1, 1, response3.OrderNo)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("查询发货单：[%v]\n", response7)
	}

	err = base_server_order.ModifyDeliverInfo(1, 1, response3.OrderNo, "新发货信息")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("修改发货单成功：[%v]\n", "")
	}

	err = base_server_order.ReceiptOrderGoods(1, 1, response3.OrderNo)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("确认收货成功：[%v]\n", "")
	}

	err = base_server_order.CancelDeliverOrder(1, 1, response3.OrderNo)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("取消发货成功：[%v]\n", "")
	}

	err = base_server_order.BusinessUpdate(1, 1, response3.OrderNo, response4.Details[0].OrderNo, 1, 2, "新extra", nil)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("业务更新成功：[%v]\n", "")
	}

}
