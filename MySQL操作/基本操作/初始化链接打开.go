package main

import (
	"bufio"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io"
	"os"
	"strings"
	"time"

	//执行mysql包的sql驱动
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//建立数据可链接
	DB, e = sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/gowk")
	HandleError(e, "sqlx.Connect")

	/*	var kfp =make([]kfperson,0)
		DB.Select(&kfp, "select id,name,idcard from t_bigData_kf where name like ?;", "东%")
		fmt.Println(kfp)*/
	//
	_, e = os.Stat("E:\\workspace\\Goland\\MySQL操作\\基本操作\\insData.mark")
	if e == nil {
		fmt.Println("数据已初始化====================！！！")
		return
	}

	//必要时建表
	_, e = DB.Exec("create table if not exists t_bigData_kf2(id int primary key auto_increment,name varchar(30),idcard char(18));")
	HandleError(e, "DB.Exec create table")
	fmt.Println("数据表已创建")

	//初始存取数据管道
	Chkp = make(chan *kfperson, 800)
	Chbd = make(chan *kfperson, 300)
	//开启1000条协程写入数据,
	for i := 0; i < 80; i++ {
		go insertKftable()
	}
	//准备一个处理脏数据的协程
	//go writeBadTxt()

	//创建一个失败数据Txt
	A_bad, _ = os.OpenFile("E:/假数据/清洗数据/A_bad.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0754)
	defer A_bad.Close()

	//读取大数据文本
	file, e := os.Open("E:/假数据/清洗数据/A.txt")
	HandleError(e, "os.Open")
	defer file.Close()
	//以缓存的方式读出，区别于ioutil.ReadFile()一次性读入文件到内存，不适合大文本的读取
	reader := bufio.NewReader(file)
	//kfps := new(kfperson)   ///这里曾经埋葬了我一夜！！！！！！
	for {
		linStr, e := reader.ReadString('\n')
		if e == io.EOF {
			//close(Chbd)
			//close(Chkp)
			break
		}
		HandleError(e, "reader.ReadString")
		linSplic := strings.Split(linStr, ",")

		kfps := new(kfperson)
		kfps.Name, kfps.Idcard = linSplic[0], linSplic[1]

		Chkp <- kfps
		time.Sleep(time.Nanosecond / 10)
	}

	//创建初始化标记【保证初始化只做一次】
	_, e = os.Create("E:\\workspace\\Goland\\MySQL操作\\基本操作\\insData.mark")
	if e == nil {
		fmt.Println("***初始化标记创建成功！")
	}

}

func writeBadTxt() {
	writer := bufio.NewWriter(A_bad)
	for ps := range Chbd {
		writer.WriteString(ps.Name + "," + ps.Idcard + "\n")
		//fmt.Println("***************************************", "Gid:", GetGID())
	}
	writer.Flush()
	close(Chbd)
}

func insertKftable() {
	defer func() {
		close(Chkp)
		DB.Close()
	}()

	for ps := range Chkp {
		//fmt.Println("============================", len(Chkp))

		DB.Exec("insert into t_bigData_kf2(name,idcard) values(?,?);", ps.Name, ps.Idcard)
/*		if e == nil {
			fmt.Println("Chkp:", len(Chkp), "\t", "Gid:", GetGID())
		} else {
			Chbd <- ps
		}*/

	}

}

//e := db.Select(&kfpeople, "select id,name,idcard from kfperson where name like ?;", name)
func main() {

	fmt.Println(DB)

	for {
		var scan string
		fmt.Print("请输入需要查询的姓名或字：")
		fmt.Scanf("%v", &scan)
		//scan="%"+scan+"%"

		if scan == "exit" {
			break
		}

		if val, ok := CacheMap[scan]; ok {
			fmt.Println(*val.con)
			val.cou += 1
			val.uptime = time.Now().UnixNano()
			//ReadPersons(val)
			continue
		}

		var kfp = make([]kfperson, 0)
		DB.Select(&kfp, "select id,name,idcard from t_bigData_kf where name like ?;", "%"+scan+"%")
		fmt.Println(kfp)
		cache := cache{con: &kfp, uptime: time.Now().UnixNano(), cou: 1}

		CacheMap[scan] = cache
		if len(CacheMap) > 3 {
			clrCacheMap(&CacheMap)
		}
	}
}

func ReadPersons(tepSlice []kfperson) {
	for _, val := range tepSlice {
		fmt.Printf("姓名：%v，身份证：%v\n", val.Name, val.Idcard)
	}
}
