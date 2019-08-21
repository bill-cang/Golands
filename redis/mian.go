package main

import (
	"./config"
	"./pool"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"utils"
)

func main() {
	redPo := pool.NewRedisPool(*config.RedisServer, *config.RedisPassword)
	conn := redPo.Get()
	//dome1(conn)
	_, _ = conn.Do("SETBIT", "t03", 4200000000, 1)
	_, _ = conn.Do("SETBIT", "t03", 1, 1)
	reply, err := redis.Int64(conn.Do("BITCOUNT", "t03"))
	utils.HandleError(err, "SETBIT")
	fmt.Println("t03 :", reply)
	reply, _ = redis.Int64(conn.Do("GETBIT", "t03",4200000000))
	fmt.Println("t03 4200000000:", reply)
	conn.Close()
}

func dome1(conn redis.Conn) {
	for i := 0; i < 100; i++ {
		if i%10 == 0 {
			continue
		}
		reply, err := redis.Int64(conn.Do("SETBIT", "t02", i, 1))
		utils.HandleError(err, "SETBIT")
		fmt.Println("t02 :", reply)
	}
	reply, err := redis.Int64(conn.Do("BITCOUNT", "t02"))
	utils.HandleError(err, "BITCOUNT")
	fmt.Println("reply:", reply)
}
