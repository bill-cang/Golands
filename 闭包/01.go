package main

import "fmt"
//https://www.cnblogs.com/hzhuxin/p/9199332.html

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			i,"\t",
			pos(i),
			neg(-2*i),
		)
	}
}
