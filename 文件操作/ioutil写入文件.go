package main

import (
	"io/ioutil"
	"strconv"
)

func main() {
	fiAdd:="E:/workspace/Goland/文件操作/001.txt"
	for i := 0; i < 100; i++ {
		bytes:=[]byte(strconv.Itoa(i)+"\n")
		ioutil.WriteFile(fiAdd,bytes,0754)
	}
}
