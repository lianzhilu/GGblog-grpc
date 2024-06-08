package rpc

import (
	"ggblog-grpc/config"
	"ggblog-grpc/idl/pb/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	UserClient user.UserServiceClient
)

func Init() {
	initClient()
}

func initClient() {
	conn, _ := connectServer()
	UserClient = user.NewUserServiceClient(conn)
}

func connectServer() (conn *grpc.ClientConn, err error) {
	conn, err = grpc.NewClient(
		config.UserSrv.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return
}
