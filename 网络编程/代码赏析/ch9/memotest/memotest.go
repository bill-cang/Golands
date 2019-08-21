// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 272.

// Package memotest provides common functions for
// testing various designs of the memo package.
package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

//网络请求，私有方法
//!+httpRequestBody
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//!-httpRequestBody

//使用一个公开成员将私有成员共享出去了
//对大量的函数是否应该公开做统一管理
var HTTPGetBody = httpGetBody
//var Fuck = fuck
//var Shit = shit
//var Damn = damn

//返回一个只读管道，其中充满了url
//让呆板的静态数据管道化，拥有了操作的灵活性
func getChanUrls() <-chan string {
	//创建一个放置url的管道
	chUrls := make(chan string)

	go func() {
		//遍历数据切片，将其中的url全部丢入管道chUrls中
		//为什么不直接返回切片而要返回管道呢？？？——管道的功能：同步异步，并发控制，机会均等，读写权限控制...
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			chUrls <- url
		}

		//关闭管道，避免死循环读取时（无法判断循环结束点而）死锁
		close(chUrls)
	}()

	//返回数据管道
	return chUrls
}

//【数据获取器】接口:
//为啥要搞接口？？？——当然是为了实现多态！！！（二级缓存版实现，	DB版实现、文件版实现...）
type IDataGetter interface {
	//給入一个【指标key】，获得对应的【数据或异常】
	Get(key string) (interface{}, error)
}

/*
//!+seq
	m := memo.New(httpGetBody)
//!-seq
*/

//同步执行多个url的数据获取
//【测试用例t】是为了打印错误日志的
//【数据获取器接口实例dataGetter】用于获取数据——只在乎dataGetter能返回数据（接口方法Get），具体实现方式不关心（多态的共性）
func DoGetUrlSync(t *testing.T, dataGetter IDataGetter) {
	//!+seq
	//循环读取管道内容（写入端结束时必须要close掉否则容易造成死锁）
	//每一个url的数据获取都没有起goroutine,所以它们是串行执行的（同步）
	for url := range getChanUrls() {
		//记录开始时间
		start := time.Now()

		//核心代码：调用【数据获取器接口实例dataGetter】获取数据
		value, err := dataGetter.Get(url)

		//处理错误和打印执行时长、数据长度
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
	//!-seq
}

/*
//!+conc
	m := memo.New(httpGetBody)
//!-conc
*/

//异步执行多个url的数据获取
func DoGetUrlAync(t *testing.T, dataGetter IDataGetter) {
	//!+conc
	var wg sync.WaitGroup
	for url := range getChanUrls() {
		wg.Add(1)
		//为每个url的数据获取开辟一条独立的goroutine
		go func(url string) {
			defer wg.Done()

			start := time.Now()
			//核心代码：调用dataGetter的Get方法获取数据
			value, err := dataGetter.Get(url)

			//打印错误日志，执行时长、数据长度
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))

		}(url)
	}
	wg.Wait()
	//!-conc
}
