package main

import (
	"../config"
	"../pool"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sync"
	"utils"
)

const (
	moneyPool = "10000" //奖金池现金
)

var (
	wg    sync.WaitGroup
	redPo *redis.Pool
	mx    sync.Mutex
)

func init() {
	flag.Parse()
	redPo = pool.NewRedisPool(*config.RedisServer, *config.RedisPassword)
	redCon := redPo.Get()
	defer redCon.Close()
	redCon.Do("set", "moneyPool", moneyPool) //将奖金池设置到redis
}

func main() {
	defer func() {
		redCon := redPo.Get()
		redCon.Close()
		i, _ := redis.Int(redCon.Do("get", "commander"))
		fmt.Printf("moneyPool 剩余现金 %v", i)
	}()

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go handleMoney()
	}
	wg.Wait()
}

func handleMoney() {
	redCon := redPo.Get()
	defer redCon.Close()

	for {
		b := decrMoneyPool(redCon)
		if b {
			break
		}
	}
	wg.Done()
}

func decrMoneyPool(conn redis.Conn) bool {
	mx.Lock()
	reply, err := redis.Int(conn.Do("decrby", "moneyPool", 1))
	if reply < 1 {
		mx.Unlock() //这里一定要释放锁，否则调用的线程无法结束
		return true
	}
	mx.Unlock()
	utils.HandleError(err, "Do get")
	fmt.Println("reply", reply)
	return false
}
