package main

import (
	"encoding/json"
	"fmt"
	"github.com/becent/golang-common/base-server-sdk/end"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) DoRequest(ctx context.Context, in *grpc_end.Request) (*grpc_end.Response, error) {
	println("get grpc request!!!")

	retData := make(map[string]interface{})
	retData["success"] = true
	retData["data"] = "call grpc success"
	buf, _ := json.Marshal(retData)

	return &grpc_end.Response{
		Success: true,
		Data:    buf,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		fmt.Println(err)
		return
	}

	ser := grpc.NewServer()
	grpc_end.RegisterEndServer(ser, &Server{})
	ser.Serve(lis)
}
