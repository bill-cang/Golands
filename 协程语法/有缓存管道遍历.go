package main

import (
	"fmt"
)

var chtt = make(chan string ,20)

func main() {
	go setChtt()
	for {
		v,e:=<-chtt
		if !e {
			break
		}
		fmt.Println(v,'\t',len(chtt))
	}

}

func setChtt() {
	var ss  = "123456789"
	for _, v := range ss {
		spf := fmt.Sprintf("%c", v)
		chtt <- spf
	}
	close(chtt)

}
