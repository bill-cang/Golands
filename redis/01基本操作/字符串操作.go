package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
	"utils"
)

func doSet() {
	err := client01.Set("test001", "2019年3月26日", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client01.Get("test001").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("test001", val)
	val2, err := client01.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func doStr()  {
	//设置一个有时间限制的键,支持最小时间1毫秒，单位纳秒
	err := client01.SetXX("a12", "lkio uuoiii", 30*time.Second).Err()
	utils.HandleError(err,"client01.SetXX")
	//<-time.NewTicker(6*time.Second).C
	a12 := client01.Get("a12")
	fmt.Println(a12)



	wg.Done()

}