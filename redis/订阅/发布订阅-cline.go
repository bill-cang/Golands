package main

import (
	"../config"
	"../pool"
	"fmt"
	"utils"
)

func main() {
	rdscon := pool.NewRedisPool(*config.RedisServer, *config.RedisPassword).Get()
	rdscon.Send("publish","redisChat","go redis")
	rdscon.Flush()

	reply, err := rdscon.Receive()
	utils.HandleError(err,"redis.Values")
	fmt.Println(reply)
}
