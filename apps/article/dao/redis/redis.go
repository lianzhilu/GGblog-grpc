package redis

import (
	"context"
	"fmt"
	"ggblog-grpc/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var Text int

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisConf.Addr,
		Password: "",                        // no password set
		DB:       config.RedisConf.DB,       // use default DB
		PoolSize: config.RedisConf.PoolSize, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func InitDataBase() {
	if err := initClient(); err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	// GetRank(ctx, -1, -1)
	// GetRank(ctx, -1, 5)
	// GetRank(ctx, 5, -1)
	fmt.Println(GetRank(ctx, 5, 1))
	// AddClicks(ctx, "1")
	// res, err := rdb.ZScore(ctx, "article_view", "1").Result()
	// fmt.Println(err == redis.Nil)
	// fmt.Println(res)

	// res, err := rdb.ZScore(ctx, "article_key", "k1").Result()
	// fmt.Println(err == redis.Nil)
	// fmt.Println(res)
}
