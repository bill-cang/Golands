package main

import "fmt"

type DiscountActivity struct {
	Id          string `xorm:"not null pk comment('id') VARCHAR(50)"`
	DisName     string `xorm:"not null comment('活动名称') VARCHAR(128)"`
	DisToken    string `xorm:"not null comment('活动币种') VARCHAR(50)"`
	BuyToken    string `xorm:"comment('置换币种') VARCHAR(50)"`
	DisPool     string `xorm:"not null default 0.0000000000 comment('奖金池数量') DECIMAL(32,10)"`
	BuyAccSum   int    `xorm:"not null default 0 comment('已购人数') INT(11)"`
	BuyTokenSum string `xorm:"not null default 0.0000000000 comment('置换总额') DECIMAL(32,10)"`
	ByAccMax    string `xorm:"not null default 0.0000000000 comment('个人最大认购') DECIMAL(32,10)"`
	BuyAccMin   string `xorm:"not null default 0.0000000000 comment('个人最小认购') DECIMAL(32,10)"`
	Enable      int    `xorm:"not null default 0 comment('有效状态 1：有效，0：无效') INT(11)"`
	StartTime   string `xorm:"not null default '0' comment('开始时间') VARCHAR(50)"`
	EndTime     string `xorm:"not null default '0' comment('结束时间') VARCHAR(50)"`
	UpdateTime  string `xorm:"not null default '0' comment('创建/更新时间') VARCHAR(50)"`
}

func init() {
	NewMysql()
}


func main() {
	fmt.Println("DBconn:",DBconn)
	query := "select * from DISCOUNT where id = ?"
	param := "1"
	rows, err := DBconn.Query(query, param)
	HandleError(err,"777")
	fmt.Println(rows.Columns())

	activs := make([]DiscountActivity, 0)
	for rows.Next() {
		var disc DiscountActivity
		_ = rows.Scan(&disc.Id, &disc.Dis_name, &disc.Dis_token, &disc.Dis_pool, &disc.Buy_acc_sum, &disc.Buy_token_sum,
			&disc.Start_time, &disc.End_time,&disc.Created_time)
		activs = append(activs, disc)
	}

	fmt.Println(activs)

}
