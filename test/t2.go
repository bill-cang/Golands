package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	chExit = make(chan string)
	ch001  = make(chan string)
	spl    = make([]*string, 10)
)

func main() {
	go sameWay()

	if _, ok := <-ch001; ok {
		readSplic()
		chExit <- "test over!"
	}

	<-chExit
}

func sameWay() {
	defer fmt.Println("============sameWay结束后再去查看内存是否还存在！")

	var chsw = make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			var tmp = strconv.Itoa(i)
			time.Sleep(1 * time.Second)
			spl[i] = &tmp
			fmt.Printf("【存】地址：%v\t值：%v\n", &tmp, tmp)
		}
	}()
	time.Sleep(5 * time.Second)
	go func() {
		readSplic()
		ch001 <- "sameWay over!"
		chsw <- "sameWay over!"
	}()
	<-chsw
}

func readSplic() {
	le := len(spl) - 1
	for i := 0; i < le; i++ {
		fmt.Printf("【取】地址：%v\t", spl[i])
		fmt.Printf("值：%v\n", *spl[i])
		time.Sleep(1 * time.Second)
	}
}
