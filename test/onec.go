package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	once = &sync.Once{}
	dd int
)

func main() {
	for i := 0; i < 10; i++ {
		tmp:=testOnce()
		fmt.Println(tmp)
	}
	time.Sleep(1 * time.Second)
}

func testOnce() int {
	once.Do(func() {
		dd = +1
	})
	return dd
}
