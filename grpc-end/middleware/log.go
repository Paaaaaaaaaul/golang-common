package middleware

import (
	"github.com/becent/golang-common"
	"github.com/becent/golang-common/grpc-end"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger(c *grpc_end.GRpcContext) {
	now := time.Now()

	c.Next()

	in := c.GetRequest()

	logResp := append([]byte{}, c.GetResponse().Data...)
	if len(logResp) > 1024*2 {
		logResp = logResp[:1024*2]
		logResp = append(logResp, []byte("...")...)
	}

	logrus.WithFields(logrus.Fields{
		"appName":      c.GetAppName(),
		"controller":   in.Controller,
		"action":       in.Action,
		"param":        in.Params,
		"header":       in.Header,
		"response":     string(logResp),
		"responseSize": len(c.GetResponse().Data),
		"useTime":      time.Since(now).String(),
		"logId":        common.GetGorouterIDFlag(),
	}).Info()
}
