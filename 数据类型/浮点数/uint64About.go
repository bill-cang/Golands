package main

import "fmt"

func main() {

	fmt.Println("12345"[0:4])
	fmt.Println(fmt.Sprintf("%.2f",29999.996))

	a:=2
	var c float64 = 1000
	var d float64 = 100000000
	fmt.Println(uint64(a)*uint64(c*d))
}
