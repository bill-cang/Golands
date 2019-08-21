package main

import (
	"fmt"
	"testing"
)

func TestIsOddEven(t *testing.T) {
	resMAp := map[bool]int{true: 0, false: 0}
	for i := 0; i < 100; i++ {
		bool := IsOddEven(i)
		resMAp[bool]++
		if bool {
			fmt.Printf("%v是偶数\n", i)
			continue
		}
		fmt.Printf("%v是奇数\n", i)
	}
	fmt.Printf("奇数%v,偶数%v\n", resMAp[false], resMAp[true])
}

func TestOddEvenCount(t *testing.T) {
	ceSlice := []int{}
	for i := 0; i < 100; i++ {
		ceSlice = append(ceSlice, i)
	}
	reMap := OddEvenCount(ceSlice)
	fmt.Println(reMap)
}

func TestGetRandNumber(t *testing.T) {
	ceMap := map[int]int{}
	for i := 0; i < 100; i++ {
		randNumber := GetRandNumber(3)
		ceMap[randNumber] = 1
		fmt.Println(randNumber)
	}

	fmt.Printf("不重复率 %v\n", len(ceMap))
}

func TestIsDaffodil(t *testing.T) {
	for i := 100; i < 1000; i++ {
		bo := IsDaffodil(i, func(i int) bool {
			a := i / 100
			b := i % 100 / 10
			c := i % 10
			return a*a*a+b*b*b+c*c*c == i
		})
		if bo {
			fmt.Println("水仙花数：",i)
		}
	}
}