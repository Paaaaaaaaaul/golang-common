package common

import (
	"testing"
)

func TestGET(t *testing.T) {
	body, err := GET("https://s-api.xyhj.io/v1/w/zh/:controller/:action")
	if err != nil {
		t.Fatalf(err.Error())
	}

	println(string(body))
}

func TestPOST(t *testing.T) {
	body, err := POST("https://s-api.xyhj.io/v1/w/zh/:controller/:action", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	println(string(body))
}
