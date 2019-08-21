package main

import (
	"fmt"
	"time"
)

/*
·开辟三条协程，A每秒生成一个三位随机数，B输出该值的奇偶性，C输出奇数和偶数的总个数；
·当生成一个水仙花数时，程序结束；
·尽量封装和抽取能复用的部分，并对它们做单元测试；
*/
var (
	chAB = make(chan int)
	chBC = make(chan [2]int)
	chMI = make(chan string)

	cun [2]int
)

func main() {

	go Arontine()
	go Broutine()
	go Croutine()

	<-chMI
}

func Croutine() {
	for {
		if chBC != nil {
			temp := <-chBC
			fmt.Printf("Croutine统计：奇数：%v，偶数：%v\n", temp[1], temp[0])
		}
	}
}

func Broutine() {
	var jj, oo int
	for {
		if chAB != nil {
			n := <-chAB
			if n%2 == 0 {
				fmt.Printf("Broutine判断%v是偶数\n", n)
				oo++
				cun[0] = oo
				chBC <- cun
				continue
			}
			jj++
			cun[1] = jj
			chBC <- cun
			fmt.Printf("Broutine判断%v是奇数\n", n)
		}
	}

}

func Arontine() {
	daffodil := []int{}
	var cun int = 0
	for {
		rand := GetRandNumber(3)
		chAB <- rand
		fmt.Printf("Arontine生3产位随机数%v\n", rand)
		time.Sleep(time.Second * 1)
		bo := IsDaffodil(rand, func(i int) bool {
			a := i / 100
			b := i % 100 / 10
			c := i % 10
			return a*a*a+b*b*b+c*c*c == i
		})
		if bo {
			cun++
			fmt.Println("发现水仙花数===============================", rand)
			daffodil = append(daffodil, rand)
			if cun == 4 {
				chMI <- "over"
				fmt.Println(daffodil)
			}
		}
	}

}
