package main

//斐波那契数列：1,1,2,3,5,8,13,21,34,55,89……
//方法一：每一项等于前两项之和
func Feibonacci1(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Feibonacci1(n-1) + Feibonacci1(n-2)
	}
}

//方法二:向后移动
func Feibonacci2(n int) int {
	X, Y := 1, 1
	for i := 0; i < n; i++ {
		X, Y = Y, X+Y
	}
	return X
}


