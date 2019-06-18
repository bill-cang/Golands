package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	sourceFile = "E:/workspace/Goland/文件操作/mzd.jpg"
	desFile    = "E:/workspace/Goland/文件操作/mzx.jpg"
)

func main() {
	domeI()
	domeII()
	domeIII()

}
//3、运用bufio缓存区读写
func domeIII() {
	//1、打开文件流
	surFilPtr, _ := os.OpenFile(sourceFile, os.O_RDONLY, 0754)
	desFilPtr, _ := os.OpenFile(desFile, os.O_WRONLY|os.O_CREATE, 0754)
	//挂起关闭文件流
	defer surFilPtr.Close()
	defer desFilPtr.Close()
	//创建读入器、写出器
	reader := bufio.NewReader(surFilPtr)
	writer := bufio.NewWriter(desFilPtr)
	//创建一块8字节的缓存区
	buffer := make([]byte, 8)
	for {
		//将文件数据依次读进缓存区
		_, err := reader.Read(buffer)
		if err == io.EOF {
			fmt.Println("拷贝成功！")
			//倒出缓存区数据到目标文件
			writer.Flush()
			return
		}
		//把缓存区的数据依次写入缓存区
		writer.Write(buffer)
	}
}

//2、利用ioutil快速读写完成拷贝
func domeII() {
	bytes, _ := ioutil.ReadFile(sourceFile)
	err := ioutil.WriteFile(desFile, bytes, 0754)
	if err != nil {
		fmt.Println("拷贝失败！错误：", err)
		return
	}
	fmt.Println("拷贝成功！")
}

//1、	io.Copy() 系统自带
func domeI() {
	//1、打开文件流
	surFilPtr, _ := os.OpenFile(sourceFile, os.O_RDONLY, 0754)
	desFilPtr, _ := os.OpenFile(desFile, os.O_WRONLY|os.O_CREATE, 0754)
	//挂起关闭文件流
	defer surFilPtr.Close()
	defer desFilPtr.Close()
	//实行拷贝
	_, err := io.Copy(desFilPtr, surFilPtr)
	if err != nil {
		fmt.Println("拷贝失败！错误：", err)
		return
	}
	fmt.Println("拷贝成功！")
}
