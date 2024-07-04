package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var key string = "article_view"

func AddClicks(ctx context.Context, id int) float64 {
	if rdb == nil {
		fmt.Println("redis数据库未初始化")
		return 0.0
	}
	idString := strconv.Itoa(id)
	score, _ := rdb.ZScore(ctx, key, idString).Result()
	score++
	rdb.ZAdd(ctx, key, &redis.Z{Score: score, Member: idString})
	newScore, _ := rdb.ZScore(ctx, key, idString).Result()
	return newScore
}

func GetRank(ctx context.Context, pageSize int, pageNum int) []string {
	var start, stop int64
	if pageSize == -1 {
		start = 0
		stop = -1
	} else if pageSize > 0 && pageNum == -1 {
		start = 0
		stop = int64(pageSize) - 1
	} else {
		start = (int64(pageNum) - 1) * int64(pageSize)
		stop = int64(pageNum)*int64(pageSize) - 1
	}

	fmt.Println("start: ", start)
	fmt.Println("stop: ", stop)

	ranks, _ := rdb.ZRevRange(ctx, key, start, stop).Result()
	return ranks
}
