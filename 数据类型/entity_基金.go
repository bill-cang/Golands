package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Share struct {
	code      string //股票代码
	name      string //名称
	fundnum   string //持有基金家数
	total     string //持股总量（万股）
	change    string //持股变化和上季度比
	totalcap  string //持股总市值(万元)
	accrate   string //流通市值占比（%）
	changesta string //遇上季变化
	time      string //时间
}



func GetHttpData(url string) (ShaMap map[string]Share) {
	//打开请求网络
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("网络请求失败，err=", err)
		return
	}
	//挂起关闭网络数据流
	defer resp.Body.Close()
	//从reposBody中解析获取数据字节
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("解析网络数据失败，err=", err)
		return
	}
	ShaMap, err = unMarShare(bytes)
	if err != nil {
		fmt.Println("json反序列化失败，err=", err)
		return
	}
	return
}

