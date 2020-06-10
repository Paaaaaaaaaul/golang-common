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
			OrderServerHost: "http://localhost:5055",
		},
	})
	defer base_server_sdk.ReleaseBaseServerSdk()

	now := time.Now()
	defer func(now time.Time) {
		println(time.Since(now).String())
	}(now)


	findParams := &base_server_order.FindOrder{
		Order:     &base_server_order.Order{
			OrderId:    0,
			OrgId:      200,
			UserId:     0,
			OrderType:  0,
			OrderNo:    0,
			Status:     0,
			Remark:     "",
			CreateTime: 0,
			UpdateTime: 0,
			ExtStr1:    "",
			ExtStr2:    "",
			ExtStr3:    "",
			ExtStr4:    "",
			ExtStr5:    "",
			ExtStr6:    "",
			ExtStr7:    "",
			ExtStr8:    "",
			ExtStr9:    "",
			ExtStr10:   "",
			ExtStr11:   "",
			ExtStr12:   "",
			ExtStr13:   "",
			ExtStr14:   "",
			ExtStr15:   "",
			ExtStr16:   "",
			ExtStr17:   "",
			ExtStr18:   "",
			ExtStr19:   "",
			ExtStr20:   "",
			ExtStr21:   "",
			ExtStr22:   "",
			ExtStr23:   "",
			ExtStr24:   "",
			ExtStr25:   "",
			ExtStr26:   "",
			ExtStr27:   "",
			ExtStr28:   "",
			ExtStr29:   "",
			ExtStr30:   "",
			ExtBigStr1: "",
			ExtBigStr2: "",
			ExtBigStr3: "",
			ExtBigStr4: "",
			ExtBigStr5: "",
			ExtText1:   "",
			ExtText2:   "",
			ExtInt1:    0,
			ExtInt2:    0,
			ExtInt3:    0,
			ExtInt4:    0,
			ExtInt5:    0,
			ExtInt6:    0,
			ExtInt7:    0,
			ExtInt8:    0,
			ExtInt9:    0,
			ExtInt10:   0,
			ExtInt11:   0,
			ExtInt12:   0,
			ExtInt13:   0,
			ExtInt14:   0,
			ExtInt15:   0,
			ExtInt16:   0,
			ExtInt17:   0,
			ExtInt18:   0,
			ExtInt19:   0,
			ExtInt20:   0,
		},
		BeginTime: 0,
		EndTime:   0,
		Limit:     0,
		Page:      0,
	}

	var p []*base_server_order.FindOrder
	p = append(p, findParams)

	response, _ := base_server_order.Find(p)

	fmt.Println(response)


}
