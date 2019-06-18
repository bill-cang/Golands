package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	usefulData="E:/假数据/20190615/A.txt"
	junkData="E:/假数据/20190615/B.txt"
)


func main() {
	//打开数据源文件
	file, _ := os.Open("E:/假数据/大数据文本/数据/kaifangX.txt")
	defer file.Close()
	reader := bufio.NewReader(file)

	//创建数据清洗后保存文件
	goodData, _ := os.OpenFile(usefulData, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0744)
	badData, _ := os.OpenFile(junkData, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0744)

	defer func() {
		goodData.Close()
		badData.Close()
	}()


	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		lineStr, _ := ConvertEncoding(string(lineBytes), "GBK", "UTF-8")
		//fmt.Println(dstStr)
		lineSlice := strings.Split(lineStr, ",")
		if len(lineSlice) > 2 && len(lineSlice[1]) == 18 {
			writer := bufio.NewWriter(goodData)
			writer.Write([]byte(lineStr+"\n"))
			writer.Flush()
			fmt.Println("A：",lineStr)
		}else {
			writer := bufio.NewWriter(badData)
			writer.Write([]byte(lineStr+"\n"))
			writer.Flush()
			fmt.Println("B：",lineStr)
		}

	}
}
