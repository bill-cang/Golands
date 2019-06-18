package config

import "flag"

var (
	RedisServer   = flag.String("redisServer", ":6379", "")
	RedisPassword = flag.String("redisPassword", "123456", "")
)
