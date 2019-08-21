package main

import (
	"fmt"
	"sync"
)

type account struct {
	name  string
	mtx   sync.Mutex
	money float32
}

func (a *account) handleMoney(f float32) float32 {
	a.mtx.Lock()
	a.money += f
	a.mtx.Unlock()
	return a.money
}

func (a *account) getAccNAme() string {
	return a.name
}

func main() {
	wg:=sync.WaitGroup{}

	aut := account{name:"张三",money:100}
	//取钱
	wg.Add(1)
	go func() {
		money := aut.handleMoney(-500)
		fmt.Println("取款后余额：",money)
		wg.Done()
	}()

	//存钱
	wg.Add(1)
	go func() {
		money := aut.handleMoney(200)
		fmt.Println("存款后余额：",money)
		wg.Done()
	}()

	wg.Wait()
}
