package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

/*
*/

type Provice struct {
	//身份证号的前两位
	Id   string
	Name string
	//黑龙江.txt
	File *os.File
	//本省文件的数据管道
	chanData chan string
}

var (
	bgtime2 int64
	edtime2 int64
)

func init() {
	bgtime2 = time.Now().Unix()
}

func main3() {
	//创建34个省份的实例
	pMap := make(map[string]*Provice)
	ps := []string{"北京市:11", "天津市:12", "河北省:13", "山西省:14", "内蒙古自治区:15", "辽宁省:21", "吉林省:22", "黑龙江省:23", "上海市:31", "江苏省:32", "浙江省:33", "安徽省:34", "福建省:35", "江西省:36", "山东省:37", "河南省:41", "湖北省:42", "湖南省:43", "广东省:44", "广西壮族自治区:45", "海南省:46", "重庆市:50", "四川省:51", "贵州省:52", "云南省:53", "西藏自治区:54", "陕西省:61", "甘肃省:62", "青海省:63", "宁夏回族自治区:64", "新疆维吾尔自治区:65", "台湾省:71", "香港特别行政区:81", "澳门特别行政区:91"}
	for _, p := range ps {
		name := p[:len(p)-2]
		id := p[len(p)-2:]
		province := Provice{Id: id, Name: name}
		pMap[id] = &province

		//为每个省份打开一个文件
		file, _ := os.OpenFile("E:/假数据/分省整理/fun2/"+province.Name+".txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		province.File = file
		//defer file.Close()

		//创建每个省的数据管道，并各起一条协程
		province.chanData = make(chan string)
		fmt.Println("管道已经创建")
	}

	for _, province := range pMap {
		fmt.Println(province.chanData)
	}

	//创建并打开34个文件
	//创建每个省的数据管道，并各起一条协程
	for _, province := range pMap {
		wg03.Add(1)
		go writeFile(province)
	}

	//读入优质数据
	file, _ := os.Open("E:/假数据/清洗数据/A.txt")
	defer func() {
		edtime2 = time.Now().Unix()
		fmt.Println("共耗时(秒)：", bgtime2-edtime2)
		file.Close()
	}()
	reader := bufio.NewReader(file)

	//逐行判断身份证的前两位
	for {
		lineBytes, _, err := reader.ReadLine()

		//读取完毕时，关闭所有数据管道
		if err == io.EOF {
			for _, province := range pMap {
				close(province.chanData)
				fmt.Println("管道已关闭")
			}
			break
		}

		//拿出省份ID
		lineStr := string(lineBytes)
		fieldsSlice := strings.Split(lineStr, ",")
		id := fieldsSlice[1][0:2]

		//对号入座，写入相应的管道
		if province, ok := pMap[id]; ok {
			province.chanData <- (lineStr + "\n")
		}

	}

	//阻塞等待协程结束
	wg03.Wait()

}

var wg03 sync.WaitGroup

func writeFile(province *Provice) {
	writer := bufio.NewWriter(province.File)
	//死循环读取管道，管道关闭时循环结束
	for lineStr := range province.chanData {
		writer.WriteString(lineStr)
		fmt.Print(province.Name, "写入", lineStr)
	}
	writer.Flush()
	province.File.Close()
	//标记协程结束
	wg03.Done()
}
