package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const fileAdres string = "E:\\workspace\\Goland\\json正反序列化\\"

var fileName, content string

func main() {
	//准备创建的文件名和数据
	fileName = "ceShi.txt"
	content = "江雪\n千山鸟飞绝，\n万径人踪灭。\n孤舟蓑笠翁，\n独钓寒江雪。"

	//写入
	CreateWriterFile(fileName, content)
	//读取
	BufReadFile(fileName)

}

/*
os.O_WRONLY | os.O_CREATE | O_EXCL  【如果已经存在，则失败】
os.O_WRONLY | os.O_CREATE   【如果已经存在，会覆盖写，不会清空原来的文件，而是从头直接覆盖写】
os.O_WRONLY | os.O_CREATE | os.O_APPEND  【如果已经存在，则在尾部添加写】
*/
func CreateWriterFile(fileName string, content string) {
	//创建文件,给与创、写权限，421:读写操
	filePointer, err := os.OpenFile(fileAdres+fileName, os.O_CREATE|os.O_WRONLY, 0751)
	if err != nil {
		fmt.Println("文件创建失败！错误：", err)
		return
	}
	//创建成功挂起关闭文件流,在函数结束前执行
	defer filePointer.Close()
	//NewWriter创建一个以目标文件具有默认大小缓冲、写入w的*Writer。
	writer := bufio.NewWriter(filePointer)
	//写入器将内容写入缓冲。返回写入的字节数。
	size, err := writer.Write([]byte(content))
	//Flush方法将缓冲中的数据写入下层的io.Writer接口。缺少，数据将保留在缓冲区，并未写入io.Writer接口
	writer.Flush()
	if err == nil {
		fmt.Println("文件创建并写入成功！字节数：", size)
	} else {
		fmt.Println("文件创建并写入失败！错误：", err)
	}
}

//!!!修改打开方式：os.O_RDONLY
func BufReadFile(fileName string) {
	//创建文件,给与创、写权限，421:读写操
	filePointer, err := os.OpenFile(fileAdres+fileName, os.O_RDONLY, 0751)
	if err != nil {
		fmt.Println("文件创建失败！错误：", err)
		return
	}
	//创建成功挂起关闭文件流,在函数结束前执行
	defer filePointer.Close()

	//filePointer.Seek()
	reader := bufio.NewReader(filePointer)
	var readString string
	for {
		readString, err = reader.ReadString('\n')
		if err == nil {
			fmt.Print(readString)
		} else if err == io.EOF {
			break
		} else {
			fmt.Println("读取失败，错误：", err)
			break
		}
	}
}
