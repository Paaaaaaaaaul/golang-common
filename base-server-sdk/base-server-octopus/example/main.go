package main

import (
	"fmt"
	"github.com/becent/golang-common/base-server-sdk"
	"github.com/becent/golang-common/base-server-sdk/base-server-octopus"
	"time"
)

func main() {
	base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		OrgId:           5,
		AppId:           "10002",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			OctopusServerHost: "127.0.0.1:18081",
		},
		GRpcOnly: true,
	})
	defer base_server_sdk.ReleaseBaseServerSdk()

	now := time.Now()
	defer func(now time.Time) {
		println(time.Since(now).String())
	}(now)

	// email
	//err := base_server_octopus.SendEmailCode(5, 1000, "xxx@qq.com", "zh")
	//if err != nil {
	//	println(err.String())
	//}

	// sim
	//err := base_server_octopus.SendSimCode(5,1000, "130xxxx1234", "zh")
	//res, err := base_server_octopus.VerifySimCode(5,1000, "130xxxx1234", "284770")

	// captcha
	//res, err := base_server_octopus.InitCaptcha(5, "130xxxx1234", "284770")

	// ga
	//res, err := base_server_octopus.GenerateGa(5, "130xxxx1234")

	// idCard
	res, err := base_server_octopus.AuthRealName(5, "张三", "010203201909201234")

	if err != nil {
		println(err.String())
	}
	fmt.Printf("res:%+v \n", res)
}
