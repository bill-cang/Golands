package main

import (
	"fmt"
	"io/ioutil"
)

const (
	fileAdresI string = "E:\\workspace\\Goland\\json正反序列化\\"
	contentI string = "江雪\n千山鸟飞绝，\n万径人踪灭。\n孤舟蓑笠翁，\n独钓寒江雪。"
)

func main() {
	//写入
	iouWrite("ce2.text",contentI)
	//读出
	iouRead("ce2.text")

}

func iouRead(fileName string) {
	bytes, err := ioutil.ReadFile(fileAdresI + fileName)
	if err == nil {
		content := string(bytes)
		fmt.Println(content)
	} else {
		fmt.Println("读取文件失败！错误：", err)
	}
}

//ioutil 的写入模式：os.O_WRONLY|os.O_CREATE|os.O_TRUNC【无文件时创建文件写入，有文件时删除源文件覆盖写入】
func iouWrite(fileName string , content string) {
	err := ioutil.WriteFile(fileAdresI+fileName, []byte(content), 0751)
	if err == nil {
		fmt.Println("写入成功！")
	} else {
		fmt.Println("写入失败，错误：", err)
	}
}

