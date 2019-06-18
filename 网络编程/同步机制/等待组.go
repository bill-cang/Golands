package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 11; i++ {
		wg.Add(1)
		go func(num int) {
			res:=getFibonacci(num)
			fmt.Println(res)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func getFibonacci(n int) int {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	<-time.After(3 * time.Second)
	return x
}