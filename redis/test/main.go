package main

import (
	"../config"
	"../pool"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var (
	redPool *redis.Pool
)

func main() {
	flag.Parse()
	redPool = pool.NewRedisPool(*config.RedisServer, *config.RedisPassword)
	conn := redPool.Get()
	defer redPool.Close()
	//redis操作
	v, err := conn.Do("SET", "pool", "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	v, err = redis.String(conn.Do("GET", "commander"))
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println(v)

}
