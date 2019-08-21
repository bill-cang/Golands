package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type beauty struct {
	Name   string
	Age    int
	Weight float32
	Bwh    []int
}

const fileAddress = "E:\\workspace\\Goland\\json正反序列化\\zzm.json"

var zzm beauty

func init() {
	zzm = beauty{"佐々木あき", 40, 48.8, []int{82, 58, 86}}
}

func main() {
	//创建并打开使用文件
	filePointer, err := os.Create(fileAddress)
	if err != nil {
		fmt.Println("创建文件失败，error：", err)
		return
	}
	//声明当函数结束时关闭文件读写
	defer filePointer.Close()
	encoder := json.NewEncoder(filePointer)
	err = encoder.Encode(zzm)
	if err != nil {
		fmt.Println("json文件输出失败,error：", err)
	} else {
		fmt.Println("json文件输出成功！")
	}

}
