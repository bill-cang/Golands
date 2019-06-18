package main

import "fmt"

func main() {

	var animal [3]string = [3]string{}
	//数组的追加，地址不变
	animal[0] = "dog"
	fmt.Printf("animal的地址是:%p,值为:%v,长度:%d,容量:%d\n", &animal, animal, len(animal), cap(animal))
	animal[1] = "car"
	fmt.Printf("animal的地址是:%p,值为:%v,长度:%d,容量:%d\n", &animal, animal, len(animal), cap(animal))

	//数组的切片 含头不含尾
	var numAry = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sec0 := numAry[3:6]
	sec1 := numAry[:6]
	sec2 := numAry[7:]
	sec3 := numAry[:]
	fmt.Print(sec0, "\n", sec1, "\n", sec2, "\n", sec3, "\n")
	fmt.Print("***********\n")
	fmt.Printf("%p\n%p\n%p\n%p\n", &sec0, &sec1, &sec2, &sec3)
	fmt.Print("***********切片对数组的截取不是值的拷贝，而是直接引用数组底层地址\n")
	/*切片对数组的截取不是值的拷贝，而是直接引用数组底层地址,当某一切片操作修改自身元素值时，实际上对
	地址值的修改，其他的相关切片数组的对应的地址的值也随即发生改变*/
	fmt.Printf("%p,%p,%p,%p\n", &numAry[3], &numAry[0], &numAry[7], &numAry[0])
	fmt.Printf("%p,%p,%p,%p\n", &sec0[0], &sec1[0], &sec2[0], &sec3[0])

	sec1[3] = 444
	sec2[0] = 777
	fmt.Print(numAry, "\n", sec0, "\n", sec1, "\n", sec2, "\n", sec3, "\n")

	/*基于同一数组或者切片的切片，新的切片追加元素突破了原有的容量【cap】的上限，就会拷贝该切片的值
	到新的地址，脱离原数组，互不影响*/
	fmt.Printf("扩容前：sec1的cap：%d,sec1的Pointer:%p\n", cap(sec1), &sec1[0])
	sec1 = append(sec1, 11, 22, 33, 44, 55, 66)
	fmt.Printf("扩容后：sec1的cap：%d,sec1的Pointer:%p\n", cap(sec1), &sec1[0])//Pointer发生改变
	sec1[0]=10000
	fmt.Print(numAry,"\t",sec3,"\t",sec1)//脱离原数组不再影响

}
