package main

import (
	"ggblog-grpc/apps/user/dao/mysql"
	"ggblog-grpc/apps/user/service"
	"ggblog-grpc/config"
	userPb "ggblog-grpc/idl/pb/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config.InitConfig("../../../config/config.yaml")
	mysql.InitDataBase()
	server := grpc.NewServer()
	userPb.RegisterUserServiceServer(server, &service.UserServer{})

	defer server.Stop()
	listen, err := net.Listen("tcp", config.UserSrv.Addr)
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(listen)
}
