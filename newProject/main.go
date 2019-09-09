package main

import (
	"flag"
	"github.com/becent/commom/newProject/config"
	"github.com/becent/commom/newProject/data"
	"github.com/becent/commom/newProject/exception"
	"github.com/becent/commom/newProject/gRpcHandler"
	"github.com/becent/commom/newProject/handler"
	"github.com/becent/commom/newProject/model"
	"github.com/becent/commom/newProject/router"
	"github.com/becent/commom/newProject/service"
	"os"
	"strings"
)

var (
	projectName = flag.String("projectName", "helloWorld", "project name")
)

type Func func(string) error

func main() {
	flag.Parse()

	// 创建项目文件
	if err := os.Mkdir(*projectName, 755); err != nil {
		println(err.Error())
		return
	}

	// 创建main文件
	file, err := os.OpenFile(*projectName+"/main.go", os.O_CREATE|os.O_RDWR, 755)
	if err != nil {
		println(err.Error())
		return
	}
	if _, err := file.WriteString(strings.Replace(main_temple, "{{projectName}}", *projectName, -1)); err != nil {
		println(err.Error())
		return
	}

	funcs := make([]Func, 0)
	funcs = append(funcs, []Func{
		config.G_config,           // 创建config文件夹
		data.G_data,               // 创建data目录
		exception.G_exception,     // 创建异常目录
		handler.G_handler,         // 创建handler
		gRpcHandler.G_gRpcHandler, // 创建gRpcHandler
		model.G_model,             // 创建model
		router.G_router,           // 创建router
		service.G_service,         // 创建service
	}...)

	for _, f := range funcs {
		if err := f(*projectName); err != nil {
			println(err.Error())
			return
		}
	}

}

var main_temple = `package main

import (
	"context"
	"github.com/becent/commom"
	"{{projectName}}/config"
	"{{projectName}}/router"
	"github.com/judwhite/go-svc/svc"
	"google.golang.org/grpc"
	"math"
	"net/http"
	"time"
)

type Service struct {
	gRpcSvr    *grpc.Server
	httpServer *http.Server
}

func (s *Service) Init(env svc.Environment) error {
	config.InitConfig()

	// init log
	common.ConfigLogger(
		config.CURMODE,
		config.GetConfig("system", "app_name"),
		config.GetConfig("logs", "dir"),
		config.GetConfig("logs", "file_name"),
		config.GetConfigInt("logs", "keep_days"),
		config.GetConfigInt("logs", "rate_hours"),
	)

	// init mysql
	dbInfo := config.GetSection("dbInfo")
	for name, info := range dbInfo {
		if err := common.AddDB(
			name,
			info,
			config.GetConfigInt("mysql", "maxConn"),
			config.GetConfigInt("mysql", "idleConn"),
			time.Hour*time.Duration(config.GetConfigInt("mysql", "maxLeftTime"))); err != nil {
			return err
		}
	}

	// init redis
	if err := common.AddRedisInstance(
		"",
		config.GetConfig("redis", "addr"),
		config.GetConfig("redis", "port"),
		config.GetConfig("redis", "password"),
		config.GetConfigInt("redis", "db_num")); err != nil {
		return err
	}

	//

	return nil
}

func (s *Service) Start() error {
	// launch http server here...
	s.httpServer = &http.Server{
		Addr:        ":" + config.GetConfig("system", "http_listen_port"),
		Handler:     router.NewGinEngine(),
		ReadTimeout: time.Second * 5,
	}

	go func() {
		// Service connections
		if err := s.httpServer.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

	// launch gRpc service here
	var err error
	s.gRpcSvr, err = router.NewGRpcEngine().Run(
		":"+config.GetConfig("system", "gRpc_listen_port"),
		grpc.MaxRecvMsgSize(math.MaxInt32))
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Stop() error {
	// stop gRpc server
	s.gRpcSvr.GracefulStop()
	common.InfoLog("SystemStop", nil, "gRpc server graceful stop")

	// stop http server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		println("Server Shutdown:", err)
	}
	common.InfoLog("SystemStop", nil, "http server graceful stop")

	// release source here
	common.ReleaseMysqlDBPool()
	common.ReleaseRedisPool()

	return nil
}

func main() {
	if err := svc.Run(&Service{}); err != nil {
		println(err.Error())
	}
}

`
