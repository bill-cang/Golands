package main

import (
	"fmt"
	"strings"
)

//声明字符串变量
var str0, str1 string

func main() {
	//1、赋值（未赋值前str默认空字符串）
	dom1()
	//2、字符串比较
	domestr2()
	//3、包含
	bo := strings.Contains("abc", "a")
	fmt.Println(bo)

}

func domestr2() {
	//2、1：==比较,区分大小写
	str1 = "A为中华之崛起而撸码"
	if str0 != str1 {
		fmt.Println("字符串==比较区分大小写")
	}
	//2、2：Compare()比较,区分大小写,返回int 0时相同，1时不同[效率高于==]
	int01 := strings.Compare("A", "A")
	int02 := strings.Compare("a", "A")
	fmt.Println(int01, "\t", int02)
	//2.3：EqualFold ()比较，比较utf-8编码小写的条件下是否相等，不区分大小写
	bool01 := strings.EqualFold("A", "A")
	bool02 := strings.EqualFold("a", "A")
	fmt.Println(bool01, "\t", bool02)
	fmt.Println("==================================================dmoe2")

}

func dom1() {
	//赋值,未赋值前str默认空字符串
	if str0 == "" {
		fmt.Printf("声明后地址：%p,值：%v\n", &str0, str0)
	}
	str0 = "a为中华之崛起而撸码"
	fmt.Printf("赋值后地址：%p,值：%v\n", &str0, str0)
	fmt.Println("==================================================dmoe1")
}
