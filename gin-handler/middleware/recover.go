package middleware

import (
	"fmt"
	"runtime"

	common "github.com/becent/golang-common"
	"github.com/becent/golang-common/gin-handler"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Recover the first gin'middleware to handle request,
// it catch panic exception of this request, make sure the system healthy and strong,
// and panic exception will be logged to log file.
//
// defer() and recover() will be take about 20ns loss every request, but it still necessary.
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logId := common.GetGorouterIDFlag()
			log.WithFields(log.Fields{
				"app":   c.GetString(gin_handler.KEY_APPNAME),
				"logId": logId,
				"stack": stack(),
			}).Error(err)

			c.JSON(200, gin_handler.EResponse{
				GatewayRet: false,
				Success:    false,
				Err: gin_handler.Error{
					Code:    1,
					Message: fmt.Sprintf("panic logId:%d", logId),
				},
			})
		}
	}()

	c.Next()
}

func stack() string {
	var buf [2 << 10]byte
	return string(buf[:runtime.Stack(buf[:], true)])
}
