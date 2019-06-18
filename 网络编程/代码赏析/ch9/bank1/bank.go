// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

//存钱管道：从中可以源源不断地拿出善款
var chanSave = make(chan int) // send amount to deposit

//最新总额管道：出纳员会源源不断地向管道中写入最新的总金额
var chanTotalMoney = make(chan int) // receive balance

//存钱方法：往存钱管道中丢一笔钱
func Save(amount int) {
	chanSave <- amount
}

//获取银行余额
func GetMoney() int       {
	return <-chanTotalMoney
}

//出纳员：以光速进行资金的入账和更新
func teller() {
	//基金最新余额
	var money int // money is confined to teller goroutine

	for {
		select {
		//源源不断地从存钱通道中拿出善款
		case amount := <-chanSave:
			money += amount

		//源源不断地将最新的总金额写入余额管道
		case chanTotalMoney <- money:
		}
	}
}

//数据初始化
func init() {
	go teller() // start the monitor goroutine
}

//!-
