package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
	"utils"
)

var cacheInstance *RedisConn
var rdsLock sync.Mutex

type RedisConn struct {
	pool      *redis.Pool
	showDebug bool
}

func (rds *RedisConn) Do(commandName string,
	args ...interface{}) (reply interface{}, err error) {
	conn := rds.pool.Get()
	defer conn.Close()

	t1 := time.Now().UnixNano()
	reply, err = conn.Do(commandName, args...)
	if err != nil {
		e := conn.Err()
		if nil != e {
			utils.HandleError(e, "conn.Do")
		}
		var ag string
		for _, v := range args {
			ag += "\t" + v.(string)
		}
		utils.HandleError(err, "$Do "+commandName+ag)
	}
	t2 := time.Now().UnixNano()
	if rds.showDebug {
		fmt.Printf("\n[redis] [info] [long=%d ]cmd=%s, args=%v,err=%s",
			(t2-t1)/1000, commandName, args, err, )
	}

	return reply, err
}

func (rds *RedisConn) ShowDebug(b bool) {
	rds.showDebug = b
}

//单例模式
func InstanceCache() *RedisConn {
	if cacheInstance != nil {
		return cacheInstance
	}
	rdsLock.Lock()
	defer rdsLock.Unlock()
	if cacheInstance != nil {
		return cacheInstance
	}
	return NewCache()
}

func NewCache() *RedisConn {
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			conn, e = redis.Dial("tcp", fmt.Sprintf("%s:%d", RdsCache.Host, RdsCache.Port))
			if nil != e {
				utils.HandleError(e, "Dial")
				return nil, e
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         10000,
		MaxActive:       10000,
		IdleTimeout:     0,
		Wait:            false,
		MaxConnLifetime: 0,
	}
	instance := &RedisConn{
		pool: &pool,
	}
	cacheInstance = instance
	instance.ShowDebug(true)
	return instance
}
