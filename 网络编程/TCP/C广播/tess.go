package main

import "fmt"

func main() {
	var tsMap =make(map[int]string)
	tsMap[0]="A"
	tsMap[1]="B"
	tsMap[2]="C"
	tsMap[3]="D"

	fmt.Println(tsMap[5])

	for key:=range tsMap{
		fmt.Println("***",tsMap[key])
	}
}
