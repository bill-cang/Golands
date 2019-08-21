package main

import "fmt"
//声明mmp的map类型全局变量
var mmp map[string]interface{}

func main() {
	//1：mmp只是定义未经初始化，不可以存放数据，err：assignment to entry in nil map【分配到nil映射中的条目】
	dmoe1()
	//2：初始化mmp,并存入数据
	dome2()
	//3：make声明+初始化(:=)
	dome3()
	//4:声明+初始化+直接赋值
	dome4()

	//map 的遍历
	dome5()
	//map 只遍历某一个键值
	dome6()
	//map 删除键值
	dmoe7()

}

func dmoe7() {
	delete(mmp, "age")
	fmt.Println(mmp)
	fmt.Println("==================================================dmoe6")
}

func dome6() {
	//1
	val, ok := mmp["name"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("键值不存在")
	}
	val2, ok2 := mmp["money"]
	if ok2 {
		fmt.Println(val2)
	} else {
		fmt.Println("键值不存在")
	}
	//2
	name:=mmp["name"]
	if name!=nil {
		fmt.Println(name)
	}
	fmt.Println("==================================================dmoe6")
}

func dome5() {
	for key, val := range mmp {
		fmt.Println(key, ":", val)
	}
	fmt.Println("==================================================dmoe5")

}

func dome4() {
	mmp4 := map[string]interface{}{"name": "布斯雷", "money": 2.4e+8}
	fmt.Println(mmp4)
	fmt.Println("==================================================dmoe4")
}

func dome3() {
	mmp3 := make(map[string]interface{})
	mmp3["name"] = "托尼马"
	mmp3["sex"] = 1
	fmt.Println(mmp3,len(mmp3))
	fmt.Println("==================================================dmoe3")
}

func dome2() {
	mmp = make(map[string]interface{})
	fmt.Println(&mmp, len(mmp))
	mmp["name"] = "杰克马"
	mmp["age"] = 45
	fmt.Println(&mmp, len(mmp))
	//map 取值
	name:=mmp["name"]
	fmt.Println(name)
	money:=mmp["money"]
	fmt.Println(money)
	fmt.Println("==================================================dmoe2")

}

func dmoe1() {
	fmt.Println(&mmp, len(mmp))
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("错误类型：", err)
			fmt.Println("==================================================dmoe1")
		}
	}()
	mmp["name"] = "张三"
	fmt.Println(mmp)
}
