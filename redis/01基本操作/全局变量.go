package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

var (
	client01 *redis.Client
	wg sync.WaitGroup
)

func init() {
	client01 = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client01.Ping().Result()
	if err != nil {
		fmt.Println("!!!严重！redis链接失败！")
	}

}
