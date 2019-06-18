package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxInt32)

	sp01:=make([]int,5)
	for i := 0; i < 5; i++ {
		sp01[i]=i
		fmt.Println(&sp01,"\t",sp01)
	}

	sp02:=make([]int,5)
	for i := 0; i < 5; i++ {
		sp02=append(sp02, i)
		fmt.Println(&sp02,"\t",sp02)
	}

	key:=fmt.Sprintf("u_%d",0)
	sp03:=[]interface{}{key}
	sp03 = append(sp03, 2.3)
	sp03 = append(sp03, 123)
	fmt.Printf("%T\t%v",sp03,sp03)
}
