package main

import (
	"encoding/json"
	"fmt"
)

//定义接收结构体
type studentf struct {
	Name  string
	Age   int
	Score float32
	Hobby []string
}

const (
	str1 = "{\"Name\":\"张三\",\"Age\":12,\"Score\":95.5,\"Hobby\":[\"乒乓球\",\"绘画\"]}"
	str2 = "[{\"Name\":\"张三\",\"Age\":12,\"Score\":95.5,\"Hobby\":[\"乒乓球\",\"绘画\"]},{\"Name\":\"李四\",\"Age\":13,\"Score\":97,\"Hobby\":[\"篮球\",\"朗诵\"]}]"
	str3 = "{\"persons\":[{\"Name\":\"张三\",\"Age\":12,\"Score\":95.5,\"Hobby\":[\"乒乓球\",\"绘画\"]},{\"Name\":\"李四\",\"Age\":13,\"Score\":97,\"Hobby\":[\"篮球\",\"朗诵\"]}]}"
)

var (
	zhangsan, lisi studentf
	zhangsan2      interface{}
	stuSpl         []studentf
	mmp            map[string]interface{}
)

func main() {
	//1、结构体struct的反序列化，在我们看到str1是一个结构体时，最好用一个结构体去接受，用空接口接受也可以，但会得到一个map
	dome1()
	//2、切片的反序列化
	dmoe2()
	//3、map 反序列化
	dome3()

}

func dome3() {
	err := json.Unmarshal([]byte(str3), &mmp)
	if err == nil {
		fmt.Printf("%v\n", mmp)
	} else {
		fmt.Println("序列化失败，错误原因：", err)
	}
	fmt.Println("==================================================dmoe3")
}

func dmoe2() {
	err := json.Unmarshal([]byte(str2), &stuSpl)
	if err == nil {
		fmt.Printf("%v\n", stuSpl)
	} else {
		fmt.Println("序列化失败，错误原因：", err)
	}
	fmt.Println("==================================================dmoe2")
}

func dome1() {
	//1、输出结构体
	err := json.Unmarshal([]byte(str1), &zhangsan)
	if err == nil {
		fmt.Printf("%v\n", zhangsan)
	} else {
		fmt.Println("序列化失败，错误原因：", err)
	}
	//2、输出map
	err2 := json.Unmarshal([]byte(str1), &zhangsan2)
	if err == nil {
		fmt.Printf("%v\n", zhangsan2)
	} else {
		fmt.Println("序列化失败，错误原因：", err2)
	}
	fmt.Println("==================================================dmoe1")

}
