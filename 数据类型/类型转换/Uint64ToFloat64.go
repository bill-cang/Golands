package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"unsafe"
)

func main() {
	i := uint64(123456789)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	i = uint64(binary.BigEndian.Uint64(b))
	fmt.Println(i)

	fmt.Println("Float64frombits:",math.Float64frombits(i))

	i2 := (*int64)(unsafe.Pointer(&i))
	fmt.Printf("i2 type:%T ,%v\n", i2, *i2)
	i2Str := strconv.FormatInt(*i2, 10)
	f2, err := strconv.ParseFloat(i2Str, 64)
	if nil != err{
		fmt.Println(err)
	}
	fmt.Printf("f2 type:%T ,%v\n", f2, f2)

	fmt.Println(fmt.Printf("%.1f",0.1234578))
	fmt.Println(5/100)


}
