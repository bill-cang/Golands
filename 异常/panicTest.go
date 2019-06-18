package main

import (
	"fmt"
	"sort"
)

var chQ = make(chan string, 2)
var chT = make(chan int, 1000)
var chE = make(chan int, 1000)
var sli = make([]int, 0)

func init() {
	for i := 0; i < 1000; i++ {
		chT <- i
	}
	close(chT)
}

func main() {
	go writeErr()

	for i := 0; i < 10; i++ {
		go calc()
	}
	<-chQ
	<-chQ
	defer func() {
		sort.Ints(sli)
		fmt.Println(sli, "\n", len(sli))
	}()
}

func writeErr() {
	for v := range chE {
		sli = append(sli, v)
	}
	chQ <- "exit"
}

func calc() {
	var tmp int
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println(err)
		}
		go calc()
	}()

	for tmp = range chT {
		if tmp%10 == 0 {
			chE <- tmp
			sf := fmt.Sprintf("%d is err", tmp)
			panic(sf)
		}
		if tmp == 999 {
			chQ <- "exit"
			close(chE)
		}
	}
}
