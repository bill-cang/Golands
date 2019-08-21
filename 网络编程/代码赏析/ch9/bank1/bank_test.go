// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {

	//结束管道
	chanQuit := make(chan struct{})

	// Alice
	go func() {
		//存入200大洋后，立即进行余额查询
		Save(200)
		//time.Sleep(time.Nanosecond)
		fmt.Println("=", GetMoney())//300,应该引起ALice的困惑，多的100哪来的？？？

		//向结束管道中写入空的结构体
		chanQuit <- struct{}{}
	}()

	// Bob
	go func() {
		//存入100大洋后，立即进行余额查询
		Save(100)

		//向结束管道中写入空的结构体
		chanQuit <- struct{}{}
	}()

	//阻塞从结束管道中读两次内容
	// Wait for both transactions.
	<-chanQuit
	<-chanQuit

	//如果查询到的银行余额got与预期中的余额want不符，就报错
	if money, want := GetMoney(), 300; money != want {
		t.Errorf("GetMoney = %d, want %d", money, want)
	}

}
