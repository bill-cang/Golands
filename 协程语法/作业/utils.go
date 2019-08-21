package main

import (
	"math"
	"math/rand"
	"time"
)

//获取一个n位随机数
func GetRandNumber(n int) int {
	//设置随机数动态种子
	//rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	//求出随机数的位数上限
	pow10 := math.Pow10(n)
	//获取随机数
	rand := rand.Intn(int(pow10))
	//检查随机数是否小于最小n位随机数
	if rand < int(pow10)/10 {
		//若小于最小随机数，获取随机数加最小n位随机数【保证获取n位数】
		rand = +int(pow10) / 10
	}
	return rand
}
//判断奇偶性
func IsOddEven(i int) bool {
	return i%2 == 0
}
//统计奇偶个数
func OddEvenCount(paSlice []int) map[bool]int {
	reMap := map[bool]int{true:0,false:0}
	for _, v := range paSlice {
		reMap[IsOddEven(v)]++
	}
	return reMap
}

//判断是否水仙花数
func IsDaffodil(i int,rule func(i int)bool) bool {
	if i < 100 {
		return false
	}
	b := rule(i)
	return b
}


//生成[start,end)之间的随机整数
func GetRandomNumber(start, end int) int {
	randPtr := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := start + randPtr.Intn((end - start))
	return ret
}