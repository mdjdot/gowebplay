package confs

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// RedisPool redis 池
var RedisPool *redis.Pool

// InitRedisPool 初始化 Redis 池
func InitRedisPool() {
	RedisPool = &redis.Pool{
		MaxIdle:     30,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				Logger.Fatalf("连接Redis错误：%v\n", err)
			}
			return conn, err
		},
	}
}
