package main

import (
	"fmt"
	"time"
	"utils"
)

func main() {

	daySecondes := 3600 * 24
	yearSeconds := int(3600 * 24 * 365.2425)

	//Now()返回的是本地化的时间local time
	//UTC Universal Time Coordinated国际协调时间
	//北京时间比UTC快8小时
	//now := time.Now()
	now := time.Date(2018,time.July,26,7,0,0,0,time.UTC)
	elapsedSeconds := int(now.Unix())

	//算出总共逝去了多少天
	elapsedDays := elapsedSeconds/daySecondes + 1
	fmt.Println("elapsedDays=", elapsedDays)

	/*
	算出今年过去了多少天{
	·逝去了多少年
	·逝去的整年中一共包含多少天：365*n+闰年天数
	·求出今年过去的天数
	}
	*/
	//逝去了多少年
	elapsedYears := elapsedSeconds / yearSeconds
	fmt.Println("elapsedYears=", elapsedYears)

	//今年以前一共包含多少天
	thisYear := 1970 + elapsedYears
	beforeThisYearDays := 365*elapsedYears + utils.GetLeapYears(1970, thisYear-1)
	//求出今年过去的天数
	thisYearDays := elapsedDays - beforeThisYearDays
	fmt.Println("thisYearDays=", thisYearDays)

	//
	month, date := utils.GetDate(thisYear, thisYearDays)
	fmt.Printf("当前是%d年%d月%d日", thisYear, month, date)

}

