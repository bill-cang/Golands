package main

import "fmt"
//九九乘法表
func nineNine() {
	for j := 1; j < 10; j++ {
		for i := 1; i < j+1; i++ {
			fmt.Print(j, "*", i, "=", i*j, "\t")
		}
		fmt.Println()
	}
}

func dom2()  {
	fmt.Println("thIs is dom2")
	fmt.Println("thIs is dom2")
	fmt.Println("thIs is dom2")
	fmt.Println("thIs is dom2")
}