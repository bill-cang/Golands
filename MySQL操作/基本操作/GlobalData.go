package main

import (
	"bytes"
	"github.com/jmoiron/sqlx"
	"os"
	"runtime"
	"strconv"
)

/*定义全局的参数
*/
var (
	DB       *sqlx.DB
	e        error
	Chkp     chan *kfperson                //数据存取管道
	A_bad    *os.File                      //失败数据保存
	Chbd     chan *kfperson                //脏数据存取管道
)

type kfperson struct {
	Id     int    `DB:"id"`
	Name   string `DB:"name"`
	Idcard string `DB:"idcard"`
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
