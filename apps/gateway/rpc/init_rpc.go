package rpc

import (
	"ggblog-grpc/config"
	articlePb "ggblog-grpc/idl/pb/article"
	userPb "ggblog-grpc/idl/pb/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	UserClient    userPb.UserServiceClient
	ArticleClient articlePb.ArticleServiceClient
)

func Init() {
	initClient()
}

func initClient() {
	userConn, _ := connectServer("user")
	UserClient = userPb.NewUserServiceClient(userConn)

	articleConn, _ := connectServer("article")
	ArticleClient = articlePb.NewArticleServiceClient(articleConn)
}

func connectServer(name string) (conn *grpc.ClientConn, err error) {
	if name == "user" {
		conn, err = grpc.NewClient(
			config.UserSrv.Addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
	} else if name == "article" {
		conn, err = grpc.NewClient(
			config.ArticleSrv.Addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
	}

	return
}
