package main


import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DBconn *sqlx.DB
)

func NewMysql() {
	var err error

	fmt.Print("******")

	DBconn, err = sqlx.Open("mysql", "root:123456@tcp(192.168.233.131:3306)/mygin?charset=utf8")
	/*	fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "123456", "192.168.233.130:3306", "mygin"))*/
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	DBconn.SetMaxIdleConns(100)
	DBconn.SetMaxOpenConns(100)
	fmt.Println("[InitMysql] MySQL connected.")
}

func CloseMysql() {
	DBconn.Close()
}

