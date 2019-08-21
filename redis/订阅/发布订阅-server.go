package main

import (
	"../config"
	"../pool"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
)

//redis 发布订阅
func main() {
	rdscon := pool.NewRedisPool(*config.RedisServer, *config.RedisPassword).Get()
	rdscon.Send("subscribe", "redisChat")
	rdscon.Flush()
	for {
		reply, err := redis.Values(rdscon.Receive())
		if err != nil {
			fmt.Println(err)
		}
		if len(reply) == 3 {
			if reflect.TypeOf(reply[2]).String() == "[]uint8" {
				fmt.Println(string(reply[2].([]byte)))
			}
		}
	}
}
