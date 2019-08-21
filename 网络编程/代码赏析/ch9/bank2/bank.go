// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 262.

// Package bank provides a concurrency-safe bank with one account.
package bank

import (
	"time"
	"fmt"
)

//!+
var (
	//信号量，控制着最大并发数
	//sema       = make(chan struct{}, 5) // a binary semaphore guarding totalMoney
	sema       = make(chan struct{}, 1) // a binary semaphore guarding totalMoney
	totalMoney int
)

func Save(amount int) {
	sema <- struct{}{} // acquire token
	totalMoney = totalMoney + amount
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("存入%d完毕\n",amount)
	<-sema // release token
}

func GetMoney() int {
	sema <- struct{}{} // acquire token
	b := totalMoney
	<-sema // release token
	return b
}

//!-
