package main

import "fmt"

type student struct {
	name  string
	sex   string
	age   int
	score float64
}

func main() {

	var studentArray = []student{}
	stu1 := student{}
	stu1.name = "杰克马"
	stu1.sex = "男"
	stu1.age = 23
	stu1.score = 97.5
	studentArray = append(studentArray, stu1)
	fmt.Printf("studentArray的地址是:%p,值为:%v,长度:%d,容量:%d\n", &studentArray, studentArray, len(studentArray), cap(studentArray))

	stu2 := student{}
	stu2.name = "布斯雷"
	stu2.sex = "男"
	stu2.age = 23
	stu2.score = 92
	studentArray = append(studentArray, stu2)
	fmt.Printf("studentArray的地址是:%p,值为:%v,长度:%d,容量:%d\n", &studentArray, studentArray, len(studentArray), cap(studentArray))

	var studentArraySon3 = []student{}
	stu3 := student{}
	stu3.name = "董明珠"
	stu3.sex = "女"
	studentArraySon3 = append(studentArraySon3, stu3)

	var stu4 student = student{"刘弱西","男",30,97.2}
	studentArraySon3 = append(studentArraySon3, stu4)

	//切片的拼接 studentArraySon3...
	studentArray = append(studentArray, studentArraySon3...)
	fmt.Printf("studentArray的地址是:%p,值为:%v,长度:%d,容量:%d\n", &studentArray, studentArray, len(studentArray), cap(studentArray))


	//对切片或者数组的截取，不是值的拷贝，而是直接引用底层数组或切片的地址
	studentArray1 := studentArray[0:]
	fmt.Printf("%p\n%p\n",&studentArray1,&studentArray)
	fmt.Printf("%p\n%p\n",&studentArray1[0],&studentArray[0])//两个切片值的地址是同一地址

}
