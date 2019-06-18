package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	//设置随机数种子

}

func getNewRand() int {
	return rand.Intn(1000)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//assignment to entry in nil map=>map需要初始化才能使用，全局变量里面只是声明
	s := []int{}
	//fmt.Printf("%p\n",&tmap)
	for i := 0; i < 1000; i++ {
		newRand := getNewRand()
		s = append(s, newRand)
	}

	lens := len(s) - 1
	for i := 0; i < lens; i++ {
		for k := i + 1; k < lens; k++ {
			if s[i] == s[k] {
				fmt.Printf("重复随机数是%d\n", s[i])
			}
		}
	}
	fmt.Println(s)

}
