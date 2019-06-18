package main

import (
	"fmt"
	"sync"
)

func main() {
	var account  = 1000

	wg:=sync.WaitGroup{}
	mt:=sync.Mutex{}

	wg.Add(1)
	go func() {
		mt.Lock()
		account-=500
		mt.Unlock()
		fmt.Println("取款后账户余额：",account)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		mt.Lock()
		account+=100
		mt.Unlock()
		fmt.Println("存款后账户余额：",account)
		wg.Done()
	}()

	wg.Wait()

}
