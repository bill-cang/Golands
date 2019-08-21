package main

import (
	"fmt"
	"strings"
)

func main() {

	//Compare，比较两个字符串是否相同，区分大小写，相等返回0，不相等返回1【效率高于==】
	int1 := strings.Compare("abc", "abc")
	int0 := strings.Compare("abc", "Abc")
	fmt.Println(int0, int1)

	//EqualFold,比较两个字符串是否相同，不区分大小写，返回boolean
	bo0 := strings.EqualFold("abc", "abc")
	bo1 := strings.EqualFold("abc", "Abc")
	fmt.Println(bo0, bo1)

	//Contains，判断字符串包含，str0中是否包含str1，返回Boolean。【包含连续，底层还是调用了Index函数】
	bo0 = strings.Contains("abc", "ab")
	bo1 = strings.Contains("abc", "ac")
	fmt.Println(bo0, bo1)

	//ContainsAny，判断字符串s是否包含字符串chars中的任一字符。
	boo0 := strings.ContainsAny("abc", "ab")
	boo1 := strings.ContainsAny("abc", "ac")
	boo2 := strings.ContainsAny("abc", "av")
	boo3 := strings.ContainsAny("abc", "mv")
	boo4 := strings.ContainsAny("任重道远", "重任")
	fmt.Println(boo0, boo1, boo2, boo3, boo4)

	//ContainsRune，判断字符串s是否包含rune【rune 类型，代表一个 UTF-8 字符。当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型实际是一个 int32。】
	bool1 := strings.ContainsRune("食品安全", '全')
	bool2 := strings.ContainsRune("食品安全", '>')
	fmt.Println(bool1, bool2)

	//Split，分割字符串形成字符串切片
	split0 := strings.Split("happy ,new ,year", ",")
	fmt.Println(split0)

	//Replace，字符串替换，想要替换几次就为最后参数赋值，当参数小于0时，代表为替换所有
	str0 := strings.Replace("hello word!", "o", "A", 1)
	str1 := strings.Replace("hello word!", "o", "A", -1)
	fmt.Println(str0, str1)

	//Index，判断字符串出现的下标位置，-1代表不存在。
	index0 := strings.Index("hello word!", "d")
	index1 := strings.Index("hello word!", "D")
	fmt.Println(index0, index1, )
	//LastIndex，判断字符串最后一次出现的下标位置
	index2 := strings.LastIndex("hello word!", "o")
	fmt.Println(index2)

	//join,将字符串切片以特定的符号拼接成一个字符串
	join := strings.Join([]string{"a", "b", "c", "d"}, "&")
	fmt.Println(join)

	//count 统计字符串出现的次数
	count := strings.Count("hellow word!", "o")
	fmt.Println(count)

	//trim 首位去空格，制表符……
	trim := strings.Trim("	hello word !	", " \t")
	fmt.Println(trim)
	//TrimSpace,首位去空格
	space := strings.TrimSpace("    hello word !    ")
	fmt.Println(space)
	//trimFunc，自定义首位去符号
	trimFunc := strings.TrimFunc(",#there is a gril in my hreart!%^", func(r rune) bool {
		if r == ',' || r == '#' || r == '%' || r == '^' {
			return true
		}
		return false
	})
	fmt.Println(trimFunc)

	//ToUpper ,字符串转大写
	upper := strings.ToUpper("THERE IS A GRIL IN MY HREART!")
	fmt.Println(upper)
	//ToLower，字符串转小写
	lower := strings.ToLower("There Is A Gril In My Hreart!")
	fmt.Println(lower)

	//Title，每个单词首字母大写，ToTitle全大写
	title0 := strings.Title("there is a gril in my hreart!")
	title1 := strings.ToTitle("there is a gril in my hreart!")
	fmt.Println(title0,title1)

	//Repeat,复制拼接源字符串扩大n倍
	repeat := strings.Repeat("there is a gril in my hreart!", 3)
	fmt.Println(repeat)

}
