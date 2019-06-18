package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetPrimeNumber() {
	var num int
	var put, flag string
nodeA:
	fmt.Println("请输入需要检测数字")
	fmt.Scanf("%s", &put)
	put = strings.Replace(put, " ", "", -1)
	num, err := strconv.Atoi(put)
	if err != nil {
		fmt.Println("请输入正确的阿拉伯数字！\n***********************")
		goto nodeA
	}
nodeB:
	fmt.Println("请输入计算模式：a：数字检测，b:范围检测")
	fmt.Scanf("%s", &flag)
	flag = strings.Replace(flag," ","",-1)
	if flag != "a" && flag != "b" {
		fmt.Println("请输入正确的计算模式！\n**************************")
		goto nodeB
	}
	//平方根函数
	var psg int
	switch flag {
	case "a":
		psg = int(math.Sqrt(float64(num))) + 1
		ret := findSon(psg, num)
		if ret {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	case "b":
		var someNum []int
		numFlag := num + 1
		for i := 0; i < numFlag; i++ {
			//切片追加值
			someNum = append(someNum, i)
		}

		//var retNum []int
		len := len(someNum)
		for j := 0; j < len; j++ {
			//遍历切片
			var numSon int = someNum[j]
			psg = int(math.Sqrt(float64(numSon))) + 1
			ret := findSon(psg, numSon)
			if ret {
				//retNum= append(retNum,numSon)
				fmt.Print(numSon, "\t")
			}
		}

	}
}

func findSon(psg int, num int) (ret bool) {
	for i := 2; i < psg; i++ {
		if num%i == 0 {
			ret = false
			return
		} else {
			ret = true
		}
	}
	return
}
