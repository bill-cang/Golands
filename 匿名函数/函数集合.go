package main

import "fmt"

func main() {
	//求体积
	fn:= func(length, width , height int) int {
		return length*width*height
	}
	tiji := fn(3, 6, 5)
	fmt.Println(tiji)

	//匿名函数集合
	fns:=[](func(x int) int){
		func(x int) int {return x+1},
		func(x int) int {return x+2},
		func(x int) int {return x+3},
	}

	fmt.Println(fns[0](1))
	fmt.Println(fns[1](1))
	fmt.Println(fns[2](1))
}