package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"sync"
)

type RedisConf struct {
	Url  string
	Port int
	Pass string
}

func (r *RedisConf) confTesting() {
	if r.Url == "" {
		r.Url = "127.0.0.1"
	}
	if r.Port == 0 {
		r.Port = 6379
	}
}

func RedisInit(redisConf RedisConf) {
	(&redisConf).confTesting()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.Url, redisConf.Port),
		Password: redisConf.Pass,
		DB:       0,
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("连接 Redis 出错,err:", err)
		os.Exit(0)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go (&redisConf).writeToml(&wg)
	wg.Wait()
}
