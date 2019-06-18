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

var (
	sourceData  = "E:/假数据/清洗数据/A.txt"
	provinceStr = []string{"北京市:11", "天津市:12", "河北省:13", "山西省:14", "内蒙古自治区:15", "辽宁省:21", "吉林省:22", "黑龙江省:23", "上海市:31", "江苏省:32", "浙江省:33", "安徽省:34", "福建省:35", "江西省:36", "山东省:37", "河南省:41", "湖北省:42", "湖南省:43", "广东省:44", "广西壮族自治区:45", "海南省:46", "重庆市:50", "四川省:51", "贵州省:52", "云南省:53", "西藏自治区:54", "陕西省:61", "甘肃省:62", "青海省:63", "宁夏回族自治区:64", "新疆维吾尔自治区:65", "台湾省:71", "香港特别行政区:81", "澳门特别行政区:91"}
	//全局等待组
	wg = sync.WaitGroup{}

	bgtime int64
	edtime int64
)

type province struct {
	name string
	code string
	fAdr string
	chda chan string
}

func init()  {
	bgtime=time.Now().Unix()
}

func main30() {
	//读取源文件
	file, _ := os.Open(sourceData)
	defer func() {
		edtime=time.Now().Unix()
		fmt.Println("共耗时(秒)：",bgtime-edtime)
		file.Close()
	}()
	reader := bufio.NewReader(file)

	tpMAp := make(map[string]province)

	//给每个省份开辟一条协程
	for _, v := range provinceStr {
		vsp := strings.Split(v, ":")
		name := vsp[0]
		code := vsp[1]
		nFilAddr := "E:/假数据/分省整理/fun1/" + name + ".txt"
		//依次创建每个省的对象
		pne := new(province)
		pne.name = name
		pne.code = code
		pne.fAdr = nFilAddr
		pne.chda = make(chan string,10)

		tpMAp[code] = *pne

		//为这个省开辟一条协程
		wg.Add(1)
		go writeToReport(pne)

	}

	//向每个省份的通道写入数据
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			for _, val := range tpMAp {
				close(val.chda)
				fmt.Println(val.name,"管道关闭")
			}
			break
		}

		lineStr := string(lineBytes)
		beg := strings.Index(lineStr, ",")
		//lineSlice := strings.Split(string(lineBytes), ",")
		key := lineStr[beg+1 : beg+3]

		if val, ok := tpMAp[key]; ok {
			//fmt.Println(val.name,":",len(val.chda))
			val.chda <- lineStr
		}
	}

	wg.Wait()
}

func writeToReport(pne *province) {
	file, _ := os.OpenFile(pne.fAdr, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 744)
	defer func() {
		file.Close()
		wg.Done()
	}()
	writer := bufio.NewWriter(file)
	//循环读取管道内容，直到管道关闭。
	for val := range pne.chda {
		writer.WriteString(val + "\n")
		fmt.Println("***",pne.name, ":", len(pne.chda))
	}
	writer.Flush()
}
