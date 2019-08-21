package main

import (
	"fmt"
	"time"
)

var SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")

func main() {
	//当前时间转化字符串
	var _startDate int64 = time.Now().Unix()
	var startDate string = time.Unix(_startDate, 0).Format("2006-01-02 15:04:05")
	fmt.Println("startDate:",startDate)

	//字符串时间转时间秒   1565863200
	ti, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-08-19 09:00:00", SysTimeLocation)
	ux:=int(ti.Unix())
	fmt.Println("ti:",ux,_startDate)


	var in int64 = 1566176400
	var sd string = time.Unix(in, 0).Format("2006-01-02 15:04:05")
	fmt.Println(sd)
	var in2 int64 = 1565344800
	var sd2 string = time.Unix(in2, 0).Format("2006-01-02 15:04:05")
	fmt.Println(sd2)

}

or uid = "600056342"
or uid = "32694"
or uid = "300022864"
or uid = "300010774"
