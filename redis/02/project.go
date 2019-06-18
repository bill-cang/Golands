package main

import "time"

const (
	SysTimeform      = "2006-01-02 15:04:05" //格式化时间格式
	SysTimeformShort = "2006-01-02"          //格式化日期格式
	IpLimitMax       = 1000                  //同一IP一天的抽奖上限
	UMAxDayNum       = 1000                  //同一用户抽奖次数上限
)

//时区:中国时区包含【Shanghai/Chongqing】
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")
//字符串密钥
var SignSecret = []byte("ckx-640529-SZ*0523")
//cook密钥
var CookieSecret = "lottery@SZ-0523"
//优惠券类型
var GiftTypes= [4]int{0,1,2,3}
