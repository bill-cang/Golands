package main

import (
	"fmt"
)

//闭包函数
var heros = [...]string{"关羽", "张飞", "赵云", "马超", "黄忠"}

func main() {

	liubeiOrder := getHero(0)
	//fmt.Printf("liubeiOrder的数据类型是%T",liubeiOrder)
	var zhugeOrder func() string = getHero(3)

	for i := 0; i < 11; i++ {
		theHero1 := liubeiOrder()
		fmt.Printf("刘备第%d次调用了大将：%s",i+1, theHero1)
		fmt.Println()
		theHero2 := zhugeOrder()
		fmt.Printf("孔明第%d次调用了大将：%s",i+1, theHero2)
		fmt.Println("\n**************")

		//break
	}

}

func getHero(index int) func() string {
	//fmt.Println("外层inderx地址：",&index,"值为",index)
	return func() string {
		//fmt.Println("内层inderx地址：",&index,"值为",index)
		index++
		if index > len(heros)-1 {
			index = 0
		}
		return heros[index]
	}
}
