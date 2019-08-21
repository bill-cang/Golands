package baseTypeConversion

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

//向下截取保留精度
func Float64SubLower(w int, f float64) float64 {
	fStr := strconv.FormatFloat(f, 'f', -1, 64)
	index := strings.Index(fStr, ".") + 1
	newStr := fStr[:index+w]
	f64, e := strconv.ParseFloat(newStr, 64)
	if e != nil {
		fmt.Printf("[Float64Precision] err = %v", e)
		return -0.0
	}
	return f64
}

//向上截取保留精度
func Float64SubUper(w int, f float64) float64 {
	t := 0.1
	for i := 1; i < w; i++ {
		t = t * 0.1
	}
	f = f + t
	return Float64SubLower(w,f)
}
