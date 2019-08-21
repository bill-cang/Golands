package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type beautyI struct {
	Name   string
	Age    int
	Weight float32
	Bwh    []int
}

const fileAddressI = "E:\\workspace\\Goland\\json正反序列化\\zzm.json"

var zzmx beautyI

func main() {
	//打开文件
	file, err := os.Open(fileAddressI)
	if err != nil {
		fmt.Println("打开文件失败，错误：", err)
		return
	}
	//声明文件读取在函数结束时关闭
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&zzmx)
	if err == nil {
		fmt.Printf("json数据读取成功：%#v", zzmx)
	} else {
		fmt.Println("json数据读取失败，错误：", err)
	}

}
