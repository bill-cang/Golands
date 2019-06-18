package utils

import (
	"fmt"
)

//判断year是否闰年
func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Printf("%d年是闰年\n", year)
		return true
	} else {
		return false
	}
}

//求起止年份之间有多少个闰年
//startYear, endYear int 两个整型参数
func GetLeapYears(startYear, endYear int) (leapYears int) {
	for i := startYear; i < endYear+1; i++ {
		if IsLeapYear(i) {
			//leapYears=leapYears+1
			leapYears++
		}
	}
	return
}

//输入今天是今年的第几天，得到当前日期
func GetDate(thisYear, thisYearDays int) (month, date int) {
	monthDaysNormal := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	monthDaysLeap := [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	var monthDays [12]int
	if IsLeapYear(thisYear) {
		monthDays = monthDaysLeap
	} else {
		monthDays = monthDaysNormal
	}

	//将monthDays中的天数逐个相加，一旦大于thisYearDays，代表当前元素就是月份，大多少就是多少号，正好相等则说明是这个月最后一天
	var temp int
	for i := 0; i < 12; i++ {
		//temp = temp + monthDays[i]
		temp += monthDays[i]

		if temp >= thisYearDays{
			month = i+1
			date = thisYearDays-(temp-monthDays[i])

			//结束循环继续执行循环后的内容
			//break
			//返回结果，函数结束
			//return month,date
			return
		}

	}

	return 0, 0
}



