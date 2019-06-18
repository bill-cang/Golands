package main

import (
	"./model"
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"io"
	"os"
	"strings"
	"sync"
	"time"
	"utils"
)

var SourceField = "E:/假数据/清洗数据/A.txt"
var chPeo = make(chan *models.PersonMsg, 1000)
var engine *xorm.Engine
var wt sync.WaitGroup
var chE = make(chan string, 1000)

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql",
		"root:123456@tcp(127.0.0.1:3306)/card?charset=utf8")
	if nil != err {
		utils.HandleError(err, "NewEngine")
		os.Exit(-1)
	}
	engine.SetMaxIdleConns(1000)
	engine.SetMaxOpenConns(2000)
}

func main() {
	go writeErr()

	wt.Add(100)
	for i := 0; i < 100; i++ {
		go insertPerson()
	}

	file, _ := os.Open(SourceField)
	reader := bufio.NewReader(file)
	readMaxFile(reader)

	wt.Wait()
}

func writeErr() {
	file, e := os.OpenFile("E:/假数据/清洗数据/A2.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0754)
	defer file.Close()
	if e != nil {
		utils.HandleError(e,"OpenFile")
		os.Exit(-2)
	}
	writer := bufio.NewWriter(file)
	for v := range chE {
		_, e := writer.WriteString(v+"\n")
		utils.HandleError(e,"write A2")
		writer.Flush()
	}
}

func readMaxFile(reader *bufio.Reader) {
	var dstStr string
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println(dstStr)
			chE <- dstStr
			go readMaxFile(reader)
		}
	}()

	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			time.Sleep(200*time.Millisecond)
			close(chE)
			break
		}

		//dstStr, _ := ConvertEncoding(string(lineBytes), "GBK", "UTF-8")
		dstStr = string(lineBytes)
		t := strings.Split(dstStr, ",")

		peo := [12]string{}
		for i, v := range t {
			peo[i] = v
		}

		var sex int
		if peo[2] == "F" {
			sex = 1
		}

		msg := &models.PersonMsg{
			peo[0],
			peo[1],
			sex,
			peo[4],
			peo[6],
			peo[7],
			peo[9],
		}
		chPeo <- msg
	}
	close(chPeo)
}

func insertPerson() {
	for v := range chPeo {
		_, e := engine.Insert(v)
		if e != nil {
			utils.HandleError(e, "Insert")
			fmt.Println(v)
		}
	}
	wt.Done()
}
