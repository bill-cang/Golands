package baseTypeConversion

import (
	"fmt"
	"strconv"
	"unsafe"
)

func Uint64ToFloat64(i uint64) float64 {
	i2 := (*int64)(unsafe.Pointer(&i))
	i2Str := strconv.FormatInt(*i2, 10)
	f2, err := strconv.ParseFloat(i2Str, 64)
	if nil != err {
		fmt.Println(err)
		return 0
	}
	return f2
}

func Uint64ToInt64(ui uint64) int64 {
	i := (*int64)(unsafe.Pointer(&ui))
	return *i
}

