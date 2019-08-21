package main

import (
	"fmt"
	"utils/baseTypeConversion"
)

func main() {
	var fl float64 = 876.123456789234567890
	flower := baseTypeConversion.Float64SubLower(3,fl)
	fmt.Println(flower)
	fuper:=baseTypeConversion.Float64SubUper(3,fl)
	fmt.Println(fuper)


}
