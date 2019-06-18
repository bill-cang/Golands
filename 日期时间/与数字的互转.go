package main

import (
	"fmt"
	"time"
)

var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

func main() {
	//当前时间转化字符串
	var _startDate int64 = time.Now().Unix()
	var startDate string = time.Unix(_startDate, 0).Format("2006-01-02 15:04:05")
	fmt.Println(startDate)

	//字符串时间转时间秒
	ti, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-07-26 16:07:00", SysTimeLocation)
	ux:=int(ti.Unix())
	fmt.Println(ux)

	var in int64 = 1532592420
	var sd string = time.Unix(in, 0).Format("2006-01-02 15:04:05")
	fmt.Println(sd)

}
