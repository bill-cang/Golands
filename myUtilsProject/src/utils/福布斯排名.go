package utils

import "fmt"

//定义一个结构体类型
type person struct {
	name  string
	money int
}

func GetPersonSort() {
	//定义一个切片用于保存结构体，切片解释一个不定长的数组
	var pa []person
	var name string
	var money int
	for {

		fmt.Print("姓名：")
		fmt.Scanf("%s", &name)
		if name == "over" {
			break
		}
		fmt.Print("财富：")
		fmt.Scanf("%d", &money)

		var newPerson person = person{name, money}
		pa = append(pa, newPerson)

		sortPersonArray(pa)
	}
	fmt.Println(pa)
}

func sortPersonArray(pa []person) {
	len := len(pa) - 1
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if pa[i].money < pa[j].money {
				pa[i],pa[j]=pa[j],pa[i]
			}

		}
	}
}
