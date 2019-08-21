package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"utils"
)

func main() {
	//http://z.syasn.com/b/b948   http://z.syasn.com/b/b922
	resp, err := http.Get("http://z.syasn.com/b/b922")
	utils.HandleError(err, "")
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)



	fileSource := "E:/workspace/Goland/test/" + "b948"
	fileSource2 := "E:/workspace/Goland/test/" + "b9482"
	filePointer, err := os.OpenFile(fileSource, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0751)
	filePointer2, err := os.OpenFile(fileSource2, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0751)
	defer filePointer.Close()
	writer := bufio.NewWriter(filePointer)
	writer2 := bufio.NewWriter(filePointer2)

	for {
		line, _, err := reader.ReadLine()
		fmt.Println(len(line) ,line)

		//_, err = writer.Write(line)
		_, err = reader.WriteTo(writer2)

		if err == io.EOF  {
			break
		} else {
			utils.HandleError(err, "")
		}
		writer.Flush()
	}
}

//E:\workspace\Goland\test\ts
func CreateWriterFile(fileName string, bytes []byte) {
	//创建文件,给与创、写权限，421:读写操
	fileSource := "E:/workspace/Goland/test/" + fileName
	filePointer, err := os.OpenFile(fileSource, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0751)
	if err != nil {
		fmt.Println("文件创建失败！错误：", err)
		return
	}
	//创建成功挂起关闭文件流,在函数结束前执行
	defer filePointer.Close()
	//NewWriter创建一个以目标文件具有默认大小缓冲、写入w的*Writer。
	writer := bufio.NewWriter(filePointer)
	//写入器将内容写入缓冲。返回写入的字节数。
	size, err := writer.Write(bytes)
	//Flush方法将缓冲中的数据写入下层的io.Writer接口。缺少，数据将保留在缓冲区，并未写入io.Writer接口
	writer.Flush()
	if err == nil {
		fmt.Println("文件创建并写入成功！字节数：", size)
	} else {
		fmt.Println("文件创建并写入失败！错误：", err)
	}
}
