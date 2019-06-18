package main

import (
	"fmt"
)

/*
·命令行键入一行诗句启动应用：idiom.exe -cmd start -poem 大王派我来巡山
·将诗句中的每个字丢入【模糊查询管道】
·另外再建立【精确查询管道】和【结束管道】，分别存储【成语】（大鹏展翅、占山为王、龟派气功...）和【结束指令】（fuckoff）
·定时20秒，20秒之后程序结束
·时钟每秒随机读入一条管道数据：
	如果是【模糊查询管道】：起协程进行模糊查询，并汇总数据在内存
	如果是【精确查询管道】：起协程进行精确查询，并汇总数据在内存
	如果是【结束指令】：停止查询，将内存中的数据持久化为json并退出；
*/

var (
	ch001 = make(chan string, 20)//输入关键字管道
	ch002 = make(chan string, 20)//查询成语结构体管道
	ch003 = make(chan string, 1)//命令结束管道

	downMap=map[string]Idiom{}
)

func main() {
	defer writeJsonToFile(downMap)

	lin := "一树梨花压海棠"
	for _, ch := range lin {
		keyword := fmt.Sprintf("%c", ch)
		ch001 <- keyword
	}

	go func() {
		for {
			select {
			case keyStr := <-ch001:
				GetKeyIdioms(keyStr)
			case idiom := <-ch002:
				GetIdiomBody(idiom)
			}
		}
	}()

	//阻塞主程序，直到查询完毕给出信号
	<-ch003
}
