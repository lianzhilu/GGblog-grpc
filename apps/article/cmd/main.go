package main

import (
	"ggblog-grpc/apps/article/dao/mysql"
	"ggblog-grpc/apps/article/dao/redis"
	"ggblog-grpc/apps/article/service"
	"ggblog-grpc/config"
	articlePb "ggblog-grpc/idl/pb/article"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config.InitConfig("../../../config/config.yaml")
	mysql.InitDataBase()
	redis.InitDataBase()
	server := grpc.NewServer()
	articlePb.RegisterArticleServiceServer(server, &service.ArticleServer{})
	defer server.Stop()
	listen, err := net.Listen("tcp", config.ArticleSrv.Addr)
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(listen)
}
