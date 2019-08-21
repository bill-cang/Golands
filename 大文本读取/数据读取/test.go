package main

import (
	"bufio"
	"os"
	"strconv"
	"time"
)

func maint() {

	file, _ := os.OpenFile("d:/t1.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0754)
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		time.Sleep(1*time.Second)
		if i==5 {
			defer file.Close()//在第五秒挂起关闭
		}
	}

	for i := 0; i < 10; i++ {
		time.Sleep(1*time.Second)
		_, _ = writer.WriteString(strconv.Itoa(i))
	}
	writer.Flush()

}
