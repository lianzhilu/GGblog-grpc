package main

import (
	"ggblog-grpc/apps/gateway/router"
	"ggblog-grpc/apps/gateway/rpc"
	"ggblog-grpc/config"
)

func main() {
	config.InitConfig("../../../config/config.yaml")
	rpc.Init()
	router.InitRouter()
}
