package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name  string
	Age   int
	Score float32
	Hobby []string
}

var (
	zhangSan01, liSi01 student
	stuSpl01           []student
	mmp01              map[string]interface{}
)

func init() {
	zhangSan01 = student{"张三", 12, 95.5, []string{"乒乓球", "绘画"}}
	liSi01 = student{"李四", 13, 97, []string{"篮球", "朗诵"}}

	stuSpl01 = append(stuSpl01, zhangSan01, liSi01)

	mmp01 = make(map[string]interface{})
	mmp01["persons"]= stuSpl01

}

func main() {
	//结构体的序列化
	if bytes, e := json.Marshal(zhangSan01); e == nil {
		fmt.Println(string(bytes))
	} else {
		fmt.Println("序列化失败，错误原因：：", e)
	}
	//切片序列化
	if bytes, e := json.Marshal(stuSpl01); e == nil {
		fmt.Println(string(bytes))
	} else {
		fmt.Println("序列化失败，错误原因：：", e)
	}
	//map序列化
	if bytes, e := json.Marshal(mmp01); e == nil {
		fmt.Println(string(bytes))
	} else {
		fmt.Println("序列化失败，错误原因：：", e)
	}

}
