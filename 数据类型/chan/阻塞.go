package main

import (
	"fmt"
	"sync"
	"time"
)

var(
	ch001 chan int//写入管道
	ch002 chan int//判断结束管道
	wg sync.WaitGroup
)

func init() {
	//初始化
	ch001=make(chan int,100)
	ch002=make(chan int,100)
}

func main() {

	//开始10条协程等待读取管道信息
	for i := 0; i < 10; i++ {
		go reCh()
	}

	//只需要一个等待组
	wg.Add(1)
	go func() {
		//循环执行判断是否读取完毕，读取完毕，关闭写入管道
		for  {
			time.Sleep(time.Second*30)//节约性能间歇性调用
			if len(ch002)==100 {
				close(ch001)
				break
			}
		}
		close(ch002)
		wg.Done()
	}()

	for i := 0; i < 100; i++ {
		ch001<-i
	}

	wg.Wait()
}

func reCh() {
	for val := range ch001 {
		fmt.Println(val)
		ch002<-val
		time.Sleep(1*time.Second)
	}
}
