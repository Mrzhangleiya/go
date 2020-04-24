package utils

import (
	"github.com/gomodule/redigo/redis"
)

// 定义一个全局的pool
var Pool *redis.Pool

func init() {
	Pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:8080", redis.DialPassword("123"))
		},
		MaxIdle:     8,
		MaxActive:   0,   // 0表示没有限制
		IdleTimeout: 100, // 最大空闲时间
	}
}


