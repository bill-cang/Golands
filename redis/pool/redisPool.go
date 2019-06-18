package pool

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

//声明一些全局变量
/*
type Pool struct {
	//Dial 是创建链接的方法
	Dial func() (redis.Conn, error)

	//TestOnBorrow 是一个测试链接可用性的方法
	TestOnBorrow func(c redis.Conn, t time.Time) error

	// 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
	MaxIdle int

	// 最大的激活连接数，表示同时最多有N个连接 ，为0事表示没有限制
	MaxActive int

	//最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
	IdleTimeout time.Duration

	// 当链接数达到最大后是否阻塞，如果不的话，达到最大后返回错误
	Wait bool
}
*/
//初始化一个pool
func NewRedisPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     30,
		MaxActive:   5,
		IdleTimeout: 180 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
